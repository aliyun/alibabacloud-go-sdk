// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteCustomAgentResponseBodyData) *DeleteCustomAgentResponseBody
	GetData() *DeleteCustomAgentResponseBodyData
	SetErrorCode(v string) *DeleteCustomAgentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteCustomAgentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteCustomAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomAgentResponseBody
	GetSuccess() *bool
}

type DeleteCustomAgentResponseBody struct {
	Data *DeleteCustomAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DeleteCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomAgentResponseBody) GetData() *DeleteCustomAgentResponseBodyData {
	return s.Data
}

func (s *DeleteCustomAgentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteCustomAgentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomAgentResponseBody) SetData(v *DeleteCustomAgentResponseBodyData) *DeleteCustomAgentResponseBody {
	s.Data = v
	return s
}

func (s *DeleteCustomAgentResponseBody) SetErrorCode(v string) *DeleteCustomAgentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteCustomAgentResponseBody) SetErrorMessage(v string) *DeleteCustomAgentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteCustomAgentResponseBody) SetRequestId(v string) *DeleteCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomAgentResponseBody) SetSuccess(v bool) *DeleteCustomAgentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCustomAgentResponseBodyData struct {
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// example:
	//
	// Custom agent successfully deleted
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteCustomAgentResponseBodyData) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *DeleteCustomAgentResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomAgentResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomAgentResponseBodyData) SetCustomAgentId(v string) *DeleteCustomAgentResponseBodyData {
	s.CustomAgentId = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) SetMessage(v string) *DeleteCustomAgentResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) SetSuccess(v bool) *DeleteCustomAgentResponseBodyData {
	s.Success = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
