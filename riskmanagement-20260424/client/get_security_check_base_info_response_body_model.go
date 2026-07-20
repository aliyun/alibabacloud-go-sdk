// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityCheckBaseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSecurityCheckBaseInfoResponseBody
	GetCode() *string
	SetData(v *GetSecurityCheckBaseInfoResponseBodyData) *GetSecurityCheckBaseInfoResponseBody
	GetData() *GetSecurityCheckBaseInfoResponseBodyData
	SetMessage(v string) *GetSecurityCheckBaseInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSecurityCheckBaseInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSecurityCheckBaseInfoResponseBody
	GetSuccess() *bool
}

type GetSecurityCheckBaseInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSecurityCheckBaseInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6B57D35D-9DAC-5393-AE39-07697E37C2E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSecurityCheckBaseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityCheckBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityCheckBaseInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSecurityCheckBaseInfoResponseBody) GetData() *GetSecurityCheckBaseInfoResponseBodyData {
	return s.Data
}

func (s *GetSecurityCheckBaseInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecurityCheckBaseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecurityCheckBaseInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSecurityCheckBaseInfoResponseBody) SetCode(v string) *GetSecurityCheckBaseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecurityCheckBaseInfoResponseBody) SetData(v *GetSecurityCheckBaseInfoResponseBodyData) *GetSecurityCheckBaseInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetSecurityCheckBaseInfoResponseBody) SetMessage(v string) *GetSecurityCheckBaseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecurityCheckBaseInfoResponseBody) SetRequestId(v string) *GetSecurityCheckBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityCheckBaseInfoResponseBody) SetSuccess(v bool) *GetSecurityCheckBaseInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetSecurityCheckBaseInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecurityCheckBaseInfoResponseBodyData struct {
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// true
	TaskCompleted *bool `json:"TaskCompleted,omitempty" xml:"TaskCompleted,omitempty"`
}

func (s GetSecurityCheckBaseInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityCheckBaseInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSecurityCheckBaseInfoResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetSecurityCheckBaseInfoResponseBodyData) GetTaskCompleted() *bool {
	return s.TaskCompleted
}

func (s *GetSecurityCheckBaseInfoResponseBodyData) SetEnabled(v bool) *GetSecurityCheckBaseInfoResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetSecurityCheckBaseInfoResponseBodyData) SetTaskCompleted(v bool) *GetSecurityCheckBaseInfoResponseBodyData {
	s.TaskCompleted = &v
	return s
}

func (s *GetSecurityCheckBaseInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
