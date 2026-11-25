// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHBaseQuotasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SetHBaseQuotasResponseBody
	GetAccessDeniedDetail() *string
	SetRawData(v map[string]interface{}) *SetHBaseQuotasResponseBody
	GetRawData() map[string]interface{}
	SetRequestId(v string) *SetHBaseQuotasResponseBody
	GetRequestId() *string
}

type SetHBaseQuotasResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RawData            map[string]interface{} `json:"RawData,omitempty" xml:"RawData,omitempty"`
	RequestId          *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetHBaseQuotasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetHBaseQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *SetHBaseQuotasResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SetHBaseQuotasResponseBody) GetRawData() map[string]interface{} {
	return s.RawData
}

func (s *SetHBaseQuotasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetHBaseQuotasResponseBody) SetAccessDeniedDetail(v string) *SetHBaseQuotasResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SetHBaseQuotasResponseBody) SetRawData(v map[string]interface{}) *SetHBaseQuotasResponseBody {
	s.RawData = v
	return s
}

func (s *SetHBaseQuotasResponseBody) SetRequestId(v string) *SetHBaseQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetHBaseQuotasResponseBody) Validate() error {
	return dara.Validate(s)
}
