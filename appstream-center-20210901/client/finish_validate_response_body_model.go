// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishValidateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FinishValidateResponseBody
	GetCode() *string
	SetData(v bool) *FinishValidateResponseBody
	GetData() *bool
	SetMessage(v string) *FinishValidateResponseBody
	GetMessage() *string
	SetRequestId(v string) *FinishValidateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FinishValidateResponseBody
	GetSuccess() *bool
}

type FinishValidateResponseBody struct {
	// The error code returned when an error occurs.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error or result message. This field describes the cause of a failed request and may not be included in a successful response.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values: true: The call was successful. false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FinishValidateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FinishValidateResponseBody) GoString() string {
	return s.String()
}

func (s *FinishValidateResponseBody) GetCode() *string {
	return s.Code
}

func (s *FinishValidateResponseBody) GetData() *bool {
	return s.Data
}

func (s *FinishValidateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FinishValidateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FinishValidateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FinishValidateResponseBody) SetCode(v string) *FinishValidateResponseBody {
	s.Code = &v
	return s
}

func (s *FinishValidateResponseBody) SetData(v bool) *FinishValidateResponseBody {
	s.Data = &v
	return s
}

func (s *FinishValidateResponseBody) SetMessage(v string) *FinishValidateResponseBody {
	s.Message = &v
	return s
}

func (s *FinishValidateResponseBody) SetRequestId(v string) *FinishValidateResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishValidateResponseBody) SetSuccess(v bool) *FinishValidateResponseBody {
	s.Success = &v
	return s
}

func (s *FinishValidateResponseBody) Validate() error {
	return dara.Validate(s)
}
