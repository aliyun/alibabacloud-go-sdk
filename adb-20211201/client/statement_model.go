// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatement interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunUid(v int64) *Statement
	GetAliyunUid() *int64
	SetCode(v string) *Statement
	GetCode() *string
	SetCodeState(v string) *Statement
	GetCodeState() *string
	SetCodeType(v string) *Statement
	GetCodeType() *string
	SetEndTime(v int64) *Statement
	GetEndTime() *int64
	SetError(v string) *Statement
	GetError() *string
	SetHaveRows(v bool) *Statement
	GetHaveRows() *bool
	SetOutput(v string) *Statement
	GetOutput() *string
	SetResourceGroup(v string) *Statement
	GetResourceGroup() *string
	SetSessionId(v int64) *Statement
	GetSessionId() *int64
	SetStartTime(v int64) *Statement
	GetStartTime() *int64
	SetStatementId(v int64) *Statement
	GetStatementId() *int64
	SetTotalCount(v int64) *Statement
	GetTotalCount() *int64
}

type Statement struct {
	// example:
	//
	// 1111111
	AliyunUid *int64 `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// example:
	//
	// SELECT 	- FROM table
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Waiting
	CodeState *string `json:"CodeState,omitempty" xml:"CodeState,omitempty"`
	// example:
	//
	// SQL
	CodeType *string `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	// example:
	//
	// 1658987911000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// Disk is full
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// true
	HaveRows *bool `json:"HaveRows,omitempty" xml:"HaveRows,omitempty"`
	// example:
	//
	// Spark is running, the ouput is...
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// rg1
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// example:
	//
	// 10
	SessionId *int64 `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1658977911000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 100
	StatementId *int64 `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s Statement) String() string {
	return dara.Prettify(s)
}

func (s Statement) GoString() string {
	return s.String()
}

func (s *Statement) GetAliyunUid() *int64 {
	return s.AliyunUid
}

func (s *Statement) GetCode() *string {
	return s.Code
}

func (s *Statement) GetCodeState() *string {
	return s.CodeState
}

func (s *Statement) GetCodeType() *string {
	return s.CodeType
}

func (s *Statement) GetEndTime() *int64 {
	return s.EndTime
}

func (s *Statement) GetError() *string {
	return s.Error
}

func (s *Statement) GetHaveRows() *bool {
	return s.HaveRows
}

func (s *Statement) GetOutput() *string {
	return s.Output
}

func (s *Statement) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *Statement) GetSessionId() *int64 {
	return s.SessionId
}

func (s *Statement) GetStartTime() *int64 {
	return s.StartTime
}

func (s *Statement) GetStatementId() *int64 {
	return s.StatementId
}

func (s *Statement) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *Statement) SetAliyunUid(v int64) *Statement {
	s.AliyunUid = &v
	return s
}

func (s *Statement) SetCode(v string) *Statement {
	s.Code = &v
	return s
}

func (s *Statement) SetCodeState(v string) *Statement {
	s.CodeState = &v
	return s
}

func (s *Statement) SetCodeType(v string) *Statement {
	s.CodeType = &v
	return s
}

func (s *Statement) SetEndTime(v int64) *Statement {
	s.EndTime = &v
	return s
}

func (s *Statement) SetError(v string) *Statement {
	s.Error = &v
	return s
}

func (s *Statement) SetHaveRows(v bool) *Statement {
	s.HaveRows = &v
	return s
}

func (s *Statement) SetOutput(v string) *Statement {
	s.Output = &v
	return s
}

func (s *Statement) SetResourceGroup(v string) *Statement {
	s.ResourceGroup = &v
	return s
}

func (s *Statement) SetSessionId(v int64) *Statement {
	s.SessionId = &v
	return s
}

func (s *Statement) SetStartTime(v int64) *Statement {
	s.StartTime = &v
	return s
}

func (s *Statement) SetStatementId(v int64) *Statement {
	s.StatementId = &v
	return s
}

func (s *Statement) SetTotalCount(v int64) *Statement {
	s.TotalCount = &v
	return s
}

func (s *Statement) Validate() error {
	return dara.Validate(s)
}
