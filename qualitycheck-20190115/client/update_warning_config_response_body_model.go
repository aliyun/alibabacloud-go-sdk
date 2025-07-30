// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarningConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateWarningConfigResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateWarningConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWarningConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWarningConfigResponseBody
	GetSuccess() *bool
}

type UpdateWarningConfigResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWarningConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWarningConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWarningConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWarningConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWarningConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWarningConfigResponseBody) SetCode(v string) *UpdateWarningConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWarningConfigResponseBody) SetMessage(v string) *UpdateWarningConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWarningConfigResponseBody) SetRequestId(v string) *UpdateWarningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWarningConfigResponseBody) SetSuccess(v bool) *UpdateWarningConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWarningConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
