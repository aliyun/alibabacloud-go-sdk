// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeLindormV2StreamEngineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpgradeLindormV2StreamEngineResponseBody
	GetAccessDeniedDetail() *string
	SetRequestId(v string) *UpgradeLindormV2StreamEngineResponseBody
	GetRequestId() *string
}

type UpgradeLindormV2StreamEngineResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeLindormV2StreamEngineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeLindormV2StreamEngineResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeLindormV2StreamEngineResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpgradeLindormV2StreamEngineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeLindormV2StreamEngineResponseBody) SetAccessDeniedDetail(v string) *UpgradeLindormV2StreamEngineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineResponseBody) SetRequestId(v string) *UpgradeLindormV2StreamEngineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineResponseBody) Validate() error {
	return dara.Validate(s)
}
