// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoScalingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAutoScalingConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteAutoScalingConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAutoScalingConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAutoScalingConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAutoScalingConfigResponseBody
	GetSuccess() *bool
}

type DeleteAutoScalingConfigResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAutoScalingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoScalingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoScalingConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAutoScalingConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAutoScalingConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAutoScalingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAutoScalingConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAutoScalingConfigResponseBody) SetCode(v string) *DeleteAutoScalingConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAutoScalingConfigResponseBody) SetHttpStatusCode(v int32) *DeleteAutoScalingConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAutoScalingConfigResponseBody) SetMessage(v string) *DeleteAutoScalingConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAutoScalingConfigResponseBody) SetRequestId(v string) *DeleteAutoScalingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutoScalingConfigResponseBody) SetSuccess(v bool) *DeleteAutoScalingConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAutoScalingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
