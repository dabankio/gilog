package ilog

// LeveledLogger Level(args...) Levelf(t,args...)
type LeveledLogger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Panic(args ...interface{})
	Panicf(template string, args ...interface{})
}

// NamedLogger has a name
type NamedLogger interface {
	Named(name string)
}

// WithLeveledLogger Levelw(msg,args...)
type WithLeveledLogger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
}

// FullLogger with sub logger
type FullLogger interface {
	LeveledLogger
	NamedLogger
	WithLeveledLogger
	With(keysAndValues ...interface{}) FullLogger
}
