// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivyComputeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateLivyComputeResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateLivyComputeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateLivyComputeResponseBody
	GetRequestId() *string
}

type UpdateLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateLivyComputeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLivyComputeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateLivyComputeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLivyComputeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLivyComputeResponseBody) SetCode(v string) *UpdateLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLivyComputeResponseBody) SetMessage(v string) *UpdateLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLivyComputeResponseBody) SetRequestId(v string) *UpdateLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLivyComputeResponseBody) Validate() error {
	return dara.Validate(s)
}
