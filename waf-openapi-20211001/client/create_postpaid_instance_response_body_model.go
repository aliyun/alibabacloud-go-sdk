// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostpaidInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreatePostpaidInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *CreatePostpaidInstanceResponseBody
	GetRequestId() *string
}

type CreatePostpaidInstanceResponseBody struct {
	// The ID of the WAF instance.
	//
	// example:
	//
	// waf_v2_public_cn-x0r****gr1i
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 73A4E786-8235-50C0-9631-87C8****4A36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePostpaidInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePostpaidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePostpaidInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreatePostpaidInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePostpaidInstanceResponseBody) SetInstanceId(v string) *CreatePostpaidInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreatePostpaidInstanceResponseBody) SetRequestId(v string) *CreatePostpaidInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePostpaidInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
