// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHangUpDoubleCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HangUpDoubleCallResponseBody
	GetCode() *string
	SetMessage(v string) *HangUpDoubleCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *HangUpDoubleCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HangUpDoubleCallResponseBody
	GetSuccess() *bool
}

type HangUpDoubleCallResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HangUpDoubleCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HangUpDoubleCallResponseBody) GoString() string {
	return s.String()
}

func (s *HangUpDoubleCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *HangUpDoubleCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HangUpDoubleCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HangUpDoubleCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HangUpDoubleCallResponseBody) SetCode(v string) *HangUpDoubleCallResponseBody {
	s.Code = &v
	return s
}

func (s *HangUpDoubleCallResponseBody) SetMessage(v string) *HangUpDoubleCallResponseBody {
	s.Message = &v
	return s
}

func (s *HangUpDoubleCallResponseBody) SetRequestId(v string) *HangUpDoubleCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HangUpDoubleCallResponseBody) SetSuccess(v bool) *HangUpDoubleCallResponseBody {
	s.Success = &v
	return s
}

func (s *HangUpDoubleCallResponseBody) Validate() error {
	return dara.Validate(s)
}
