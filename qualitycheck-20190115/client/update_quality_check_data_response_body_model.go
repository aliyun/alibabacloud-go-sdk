// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityCheckDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateQualityCheckDataResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateQualityCheckDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateQualityCheckDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateQualityCheckDataResponseBody
	GetSuccess() *bool
}

type UpdateQualityCheckDataResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateQualityCheckDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityCheckDataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateQualityCheckDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateQualityCheckDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQualityCheckDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateQualityCheckDataResponseBody) SetCode(v string) *UpdateQualityCheckDataResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateQualityCheckDataResponseBody) SetMessage(v string) *UpdateQualityCheckDataResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateQualityCheckDataResponseBody) SetRequestId(v string) *UpdateQualityCheckDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQualityCheckDataResponseBody) SetSuccess(v bool) *UpdateQualityCheckDataResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateQualityCheckDataResponseBody) Validate() error {
	return dara.Validate(s)
}
