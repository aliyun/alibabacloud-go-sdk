// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTaskForCreateImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartTaskForCreateImageResponseBody
	GetCode() *string
	SetData(v string) *StartTaskForCreateImageResponseBody
	GetData() *string
	SetMessage(v string) *StartTaskForCreateImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartTaskForCreateImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartTaskForCreateImageResponseBody
	GetSuccess() *bool
}

type StartTaskForCreateImageResponseBody struct {
	// The business status code. A value of 200 indicates that the request is processed successfully. Other values indicate that the request fails.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned image ID.
	//
	// example:
	//
	// imgc-dabde****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned when an error occurs.
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartTaskForCreateImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartTaskForCreateImageResponseBody) GoString() string {
	return s.String()
}

func (s *StartTaskForCreateImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartTaskForCreateImageResponseBody) GetData() *string {
	return s.Data
}

func (s *StartTaskForCreateImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartTaskForCreateImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartTaskForCreateImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartTaskForCreateImageResponseBody) SetCode(v string) *StartTaskForCreateImageResponseBody {
	s.Code = &v
	return s
}

func (s *StartTaskForCreateImageResponseBody) SetData(v string) *StartTaskForCreateImageResponseBody {
	s.Data = &v
	return s
}

func (s *StartTaskForCreateImageResponseBody) SetMessage(v string) *StartTaskForCreateImageResponseBody {
	s.Message = &v
	return s
}

func (s *StartTaskForCreateImageResponseBody) SetRequestId(v string) *StartTaskForCreateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTaskForCreateImageResponseBody) SetSuccess(v bool) *StartTaskForCreateImageResponseBody {
	s.Success = &v
	return s
}

func (s *StartTaskForCreateImageResponseBody) Validate() error {
	return dara.Validate(s)
}
