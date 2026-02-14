// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OperateCustomAgentResponseBodyData) *OperateCustomAgentResponseBody
	GetData() *OperateCustomAgentResponseBodyData
	SetErrorCode(v string) *OperateCustomAgentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *OperateCustomAgentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *OperateCustomAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateCustomAgentResponseBody
	GetSuccess() *bool
}

type OperateCustomAgentResponseBody struct {
	Data *OperateCustomAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *OperateCustomAgentResponseBody) GetData() *OperateCustomAgentResponseBodyData {
	return s.Data
}

func (s *OperateCustomAgentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *OperateCustomAgentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *OperateCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateCustomAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateCustomAgentResponseBody) SetData(v *OperateCustomAgentResponseBodyData) *OperateCustomAgentResponseBody {
	s.Data = v
	return s
}

func (s *OperateCustomAgentResponseBody) SetErrorCode(v string) *OperateCustomAgentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OperateCustomAgentResponseBody) SetErrorMessage(v string) *OperateCustomAgentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *OperateCustomAgentResponseBody) SetRequestId(v string) *OperateCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateCustomAgentResponseBody) SetSuccess(v bool) *OperateCustomAgentResponseBody {
	s.Success = &v
	return s
}

func (s *OperateCustomAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OperateCustomAgentResponseBodyData struct {
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// example:
	//
	// Custom agent successfully released
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateCustomAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OperateCustomAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *OperateCustomAgentResponseBodyData) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *OperateCustomAgentResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *OperateCustomAgentResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *OperateCustomAgentResponseBodyData) SetCustomAgentId(v string) *OperateCustomAgentResponseBodyData {
	s.CustomAgentId = &v
	return s
}

func (s *OperateCustomAgentResponseBodyData) SetMessage(v string) *OperateCustomAgentResponseBodyData {
	s.Message = &v
	return s
}

func (s *OperateCustomAgentResponseBodyData) SetSuccess(v bool) *OperateCustomAgentResponseBodyData {
	s.Success = &v
	return s
}

func (s *OperateCustomAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
