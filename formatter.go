package logging

import (
	"time"
)

// Default key names for the default fields
const (
	defaultTimestampFormat = time.RFC3339
	FieldKeyTime           = "time"
	FieldKeyLevel          = "level"
	FieldKeyName           = "name"
	FieldKeyMsg            = "msg"
	FieldKeyFunc           = "func"
	FieldKeyFile           = "file"
	FieldKeyContext        = "context"
	FieldKeyLoggingError   = "logging_error"
)

type Formatter interface {
	Format(entry *Entry) ([]byte, error)
}

func prefixFieldClashes(data Fields, fieldMap FieldMap, reportCaller bool) {
	timeKey := fieldMap.resolve(FieldKeyTime)
	if t, ok := data[timeKey]; ok {
		data["fields."+timeKey] = t
		delete(data, timeKey)
	}

	msgKey := fieldMap.resolve(FieldKeyMsg)
	if m, ok := data[msgKey]; ok {
		data["fields."+msgKey] = m
		delete(data, msgKey)
	}

	levelKey := fieldMap.resolve(FieldKeyLevel)
	if l, ok := data[levelKey]; ok {
		data["fields."+levelKey] = l
		delete(data, levelKey)
	}

	vloggerErrKey := fieldMap.resolve(FieldKeyLoggingError)
	if l, ok := data[vloggerErrKey]; ok {
		data["fields."+vloggerErrKey] = l
		delete(data, vloggerErrKey)
	}

	// If reportCaller is not set, 'func' will not conflict.
	if reportCaller {
		funcKey := fieldMap.resolve(FieldKeyFunc)
		if l, ok := data[funcKey]; ok {
			data["fields."+funcKey] = l
		}
		fileKey := fieldMap.resolve(FieldKeyFile)
		if l, ok := data[fileKey]; ok {
			data["fields."+fileKey] = l
		}
	}
}
