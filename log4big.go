package log4big
import(
    "fmt"
    "time"
    "strings"
)

var level int = 0


mlevel:=make(map[string]int){
    "TRACE":0,
    "DEBUG":1,
    "INFO":2,
    "WARN":3,
    "ERROR":4,
    "FETAL":5,
}

func SetLevel(lvs string) error{
    if lvs ==""{
        return fmt.Errorf("log level is blank")
    }
    l:=strings.ToUpper(lvs)
    if v,ok:=mlevel[l];ok{
        level=v
    }else{
        return fmt.Errorf("log level setting error")
    }
    return nil
}

func SetLevelWithDefault(lv, defaultLv string) {
	err := SetLevel(lv)
	if err != nil {
		SetLevel(defaultLv)
		Warn("log level not valid. use default level: %s", defaultLv)
	}
}

// level: 0
func Trace(format string, v ...interface{}) {
	if level <= 0 {
		print_(" [TRACE] "+format, v...)
	}
}

// level: 1
func Debug(format string, v ...interface{}) {
	if level <= 1 {
		print_(" [DEBUG] "+format, v...)
	}
}

// level: 2
func Info(format string, v ...interface{}) {
	if level <= 2 {
		print_(" [INFO] "+format, v...)
	}
}

// level: 3
func Warn(format string, v ...interface{}) {
	if level <= 3 {
		print_(" [WARN] "+format, v...)
	}
}

// level: 4
func Error(format string, v ...interface{}) {
	if level <= 4 {
		print_(" [ERROR] "+format, v...)
	}
}

// level: 5
func Fetal(format string, v ...interface{}) {
	if level <= 5 {
		print_(" [FETAL] "+format, v...)
	}
}

func print_(format string, v ...interface{}) {
	fmt.Printf(time.Now().Format("2016/08/19 15:04:05")+format+"\n", v...)
}