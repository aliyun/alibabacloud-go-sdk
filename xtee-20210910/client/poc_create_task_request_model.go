// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPocCreateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDateFormat(v string) *PocCreateTaskRequest
	GetDateFormat() *string
	SetLang(v string) *PocCreateTaskRequest
	GetLang() *string
	SetRegId(v string) *PocCreateTaskRequest
	GetRegId() *string
	SetTaskName(v string) *PocCreateTaskRequest
	GetTaskName() *string
	SetToken(v string) *PocCreateTaskRequest
	GetToken() *string
}

type PocCreateTaskRequest struct {
	// Date format
	//
	// example:
	//
	// yyyyMMdd
	DateFormat *string `json:"DateFormat,omitempty" xml:"DateFormat,omitempty"`
	// Set the language type for request and response messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// Task name.
	//
	// example:
	//
	// o32d1pktx4t
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// Task token.
	//
	// This parameter is required.
	//
	// example:
	//
	// 68ce949aff6be7f1201eb2f9095ac95f
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s PocCreateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s PocCreateTaskRequest) GoString() string {
	return s.String()
}

func (s *PocCreateTaskRequest) GetDateFormat() *string {
	return s.DateFormat
}

func (s *PocCreateTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *PocCreateTaskRequest) GetRegId() *string {
	return s.RegId
}

func (s *PocCreateTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *PocCreateTaskRequest) GetToken() *string {
	return s.Token
}

func (s *PocCreateTaskRequest) SetDateFormat(v string) *PocCreateTaskRequest {
	s.DateFormat = &v
	return s
}

func (s *PocCreateTaskRequest) SetLang(v string) *PocCreateTaskRequest {
	s.Lang = &v
	return s
}

func (s *PocCreateTaskRequest) SetRegId(v string) *PocCreateTaskRequest {
	s.RegId = &v
	return s
}

func (s *PocCreateTaskRequest) SetTaskName(v string) *PocCreateTaskRequest {
	s.TaskName = &v
	return s
}

func (s *PocCreateTaskRequest) SetToken(v string) *PocCreateTaskRequest {
	s.Token = &v
	return s
}

func (s *PocCreateTaskRequest) Validate() error {
	return dara.Validate(s)
}
