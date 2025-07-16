// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *InvokeSkillResponseBody
	GetData() interface{}
	SetErrorCode(v string) *InvokeSkillResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *InvokeSkillResponseBody
	GetErrorMsg() *string
	SetFinished(v bool) *InvokeSkillResponseBody
	GetFinished() *bool
	SetMetadata(v interface{}) *InvokeSkillResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *InvokeSkillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvokeSkillResponseBody
	GetSuccess() *bool
}

type InvokeSkillResponseBody struct {
	Data      interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string     `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// if can be null:
	// false
	Finished *bool       `json:"Finished,omitempty" xml:"Finished,omitempty"`
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// RequestId
	//
	// example:
	//
	// 2715B4D3-A3FB-5FC7-AFA0-4471687B1EC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InvokeSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeSkillResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeSkillResponseBody) GetData() interface{} {
	return s.Data
}

func (s *InvokeSkillResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *InvokeSkillResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *InvokeSkillResponseBody) GetFinished() *bool {
	return s.Finished
}

func (s *InvokeSkillResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *InvokeSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeSkillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvokeSkillResponseBody) SetData(v interface{}) *InvokeSkillResponseBody {
	s.Data = v
	return s
}

func (s *InvokeSkillResponseBody) SetErrorCode(v string) *InvokeSkillResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InvokeSkillResponseBody) SetErrorMsg(v string) *InvokeSkillResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *InvokeSkillResponseBody) SetFinished(v bool) *InvokeSkillResponseBody {
	s.Finished = &v
	return s
}

func (s *InvokeSkillResponseBody) SetMetadata(v interface{}) *InvokeSkillResponseBody {
	s.Metadata = v
	return s
}

func (s *InvokeSkillResponseBody) SetRequestId(v string) *InvokeSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeSkillResponseBody) SetSuccess(v bool) *InvokeSkillResponseBody {
	s.Success = &v
	return s
}

func (s *InvokeSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
