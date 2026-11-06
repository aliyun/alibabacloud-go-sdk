// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdaptorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAdaptorResponseBody
	GetCode() *string
	SetData(v bool) *UpdateAdaptorResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateAdaptorResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAdaptorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAdaptorResponseBody
	GetSuccess() *bool
}

type UpdateAdaptorResponseBody struct {
	// The error code. This field is returned when the request fails. A successful response may not contain this field.
	//
	// example:
	//
	// InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation was processed successfully. A value of true indicates that the operation was successful.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error or result message. This field describes the cause when the request fails. A successful response may not contain this field.
	//
	// example:
	//
	// Internal error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can provide this value when troubleshooting issues.
	//
	// example:
	//
	// 01A1A962-5503-107B-AC8F-EBA85DF905AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was processed successfully. A value of true indicates that the request was processed successfully. To check whether an asynchronous task is completed, query the task status separately.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAdaptorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdaptorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAdaptorResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAdaptorResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateAdaptorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAdaptorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAdaptorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAdaptorResponseBody) SetCode(v string) *UpdateAdaptorResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAdaptorResponseBody) SetData(v bool) *UpdateAdaptorResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAdaptorResponseBody) SetMessage(v string) *UpdateAdaptorResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAdaptorResponseBody) SetRequestId(v string) *UpdateAdaptorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAdaptorResponseBody) SetSuccess(v bool) *UpdateAdaptorResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAdaptorResponseBody) Validate() error {
	return dara.Validate(s)
}
