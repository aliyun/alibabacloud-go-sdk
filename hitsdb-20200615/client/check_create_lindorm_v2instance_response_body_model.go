// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreateLindormV2InstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CheckCreateLindormV2InstanceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CheckCreateLindormV2InstanceResponseBody
	GetCode() *string
	SetMsg(v string) *CheckCreateLindormV2InstanceResponseBody
	GetMsg() *string
	SetRequestId(v string) *CheckCreateLindormV2InstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckCreateLindormV2InstanceResponseBody
	GetSuccess() *bool
}

type CheckCreateLindormV2InstanceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Msg                *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckCreateLindormV2InstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCreateLindormV2InstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCreateLindormV2InstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CheckCreateLindormV2InstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckCreateLindormV2InstanceResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CheckCreateLindormV2InstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCreateLindormV2InstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckCreateLindormV2InstanceResponseBody) SetAccessDeniedDetail(v string) *CheckCreateLindormV2InstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckCreateLindormV2InstanceResponseBody) SetCode(v string) *CheckCreateLindormV2InstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CheckCreateLindormV2InstanceResponseBody) SetMsg(v string) *CheckCreateLindormV2InstanceResponseBody {
	s.Msg = &v
	return s
}

func (s *CheckCreateLindormV2InstanceResponseBody) SetRequestId(v string) *CheckCreateLindormV2InstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCreateLindormV2InstanceResponseBody) SetSuccess(v bool) *CheckCreateLindormV2InstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CheckCreateLindormV2InstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
