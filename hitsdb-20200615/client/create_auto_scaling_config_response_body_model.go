// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoScalingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAutoScalingConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateAutoScalingConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAutoScalingConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAutoScalingConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAutoScalingConfigResponseBody
	GetSuccess() *bool
}

type CreateAutoScalingConfigResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAutoScalingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoScalingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAutoScalingConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAutoScalingConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAutoScalingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAutoScalingConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAutoScalingConfigResponseBody) SetCode(v string) *CreateAutoScalingConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAutoScalingConfigResponseBody) SetHttpStatusCode(v int32) *CreateAutoScalingConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAutoScalingConfigResponseBody) SetMessage(v string) *CreateAutoScalingConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAutoScalingConfigResponseBody) SetRequestId(v string) *CreateAutoScalingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoScalingConfigResponseBody) SetSuccess(v bool) *CreateAutoScalingConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAutoScalingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
