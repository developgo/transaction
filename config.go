package transactor

// Transactor
// Config
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

// import	"errors"

const trialLimit int = 2000000000
const trialStop int = 64
const permitError int64 = -9223372036854775806

type errorCodes int

const (
	ErrOk errorCodes = 200
)
const (
	ErrCodeUnitExist errorCodes = 400 + iota
	ErrCodeAccountExist
	ErrCodeAccountNotExist
	ErrCodeAccountNotEmpty
	ErrCodeAccountNotStop
)

// Error types
const (
	ErrorTypeExist    string = `Exist`
	ErrorTypeNotExist string = `Not exist`
	ErrorTypeNotEmpty string = `Not empty`
	ErrorTypeNotStop  string = `Not stop`
)

// Error level
const (
	ErrLayerAccount = 100 * iota
	ErrLayerUnit
	ErrLayerTransaction
)

// Error level
const (
	ErrLevelError   string = `ERROR`
	ErrLevelWarning string = `WARNING`
	ErrLevelNotice  string = `NOTICE`
)

// Error level
const (
	ErrMsgUnitExist       string = `This unit already exists`
	ErrMsgAccountExist    string = `This account already exists`
	ErrMsgAccountNotExist string = `This account already not exists`
	ErrMsgAccountNotEmpty string = `Account is not empty`
	ErrMsgAccountNotStop  string = `Account does not stop`
)

// Error_UnitExist := errors.New("This unit already exists")
