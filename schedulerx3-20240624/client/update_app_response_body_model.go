// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateAppResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAppResponseBody
	GetSuccess() *bool
}

type UpdateAppResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 39AA91C1-7BB7-5934-B15B-FD8E706D76C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAppResponseBody) SetCode(v int32) *UpdateAppResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppResponseBody) SetMessage(v string) *UpdateAppResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppResponseBody) SetRequestId(v string) *UpdateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppResponseBody) SetSuccess(v bool) *UpdateAppResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAppResponseBody) Validate() error {
	return dara.Validate(s)
}
