// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoScalingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyAutoScalingConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyAutoScalingConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyAutoScalingConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyAutoScalingConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyAutoScalingConfigResponseBody
	GetSuccess() *bool
}

type ModifyAutoScalingConfigResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAutoScalingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyAutoScalingConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyAutoScalingConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyAutoScalingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAutoScalingConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAutoScalingConfigResponseBody) SetCode(v string) *ModifyAutoScalingConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) SetHttpStatusCode(v int32) *ModifyAutoScalingConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) SetMessage(v string) *ModifyAutoScalingConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) SetRequestId(v string) *ModifyAutoScalingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) SetSuccess(v bool) *ModifyAutoScalingConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
