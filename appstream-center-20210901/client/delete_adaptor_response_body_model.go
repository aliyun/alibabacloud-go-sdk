// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdaptorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAdaptorResponseBody
	GetCode() *string
	SetData(v bool) *DeleteAdaptorResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteAdaptorResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAdaptorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAdaptorResponseBody
	GetSuccess() *bool
}

type DeleteAdaptorResponseBody struct {
	// The error code. Returned when a request fails. A successful response may not contain this field.
	//
	// example:
	//
	// InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation was successful. A value of true indicates that the operation was successful.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message or result prompt. Used to describe the cause when a request fails. A successful response may not contain this field.
	//
	// example:
	//
	// Internal error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can provide this value for troubleshooting.
	//
	// example:
	//
	// 01A1A962-5503-107B-AC8F-EBA85DF905AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates that the request was successful. To determine whether an asynchronous task is completed, query the task status separately.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAdaptorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdaptorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAdaptorResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAdaptorResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteAdaptorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAdaptorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAdaptorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAdaptorResponseBody) SetCode(v string) *DeleteAdaptorResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAdaptorResponseBody) SetData(v bool) *DeleteAdaptorResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAdaptorResponseBody) SetMessage(v string) *DeleteAdaptorResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAdaptorResponseBody) SetRequestId(v string) *DeleteAdaptorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAdaptorResponseBody) SetSuccess(v bool) *DeleteAdaptorResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAdaptorResponseBody) Validate() error {
	return dara.Validate(s)
}
