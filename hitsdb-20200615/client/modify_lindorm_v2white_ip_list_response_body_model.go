// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLindormV2WhiteIpListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyLindormV2WhiteIpListResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *ModifyLindormV2WhiteIpListResponseBody
	GetRequestId() *string
}

type ModifyLindormV2WhiteIpListResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLindormV2WhiteIpListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLindormV2WhiteIpListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2WhiteIpListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyLindormV2WhiteIpListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLindormV2WhiteIpListResponseBody) SetAccessDeniedDetail(v string) *ModifyLindormV2WhiteIpListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListResponseBody) SetRequestId(v string) *ModifyLindormV2WhiteIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListResponseBody) Validate() error {
	return dara.Validate(s)
}
