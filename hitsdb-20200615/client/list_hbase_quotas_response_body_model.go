// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHBaseQuotasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListHBaseQuotasResponseBody
	GetAccessDeniedDetail() *string
	SetRawData(v map[string]interface{}) *ListHBaseQuotasResponseBody
	GetRawData() map[string]interface{}
	SetRequestId(v string) *ListHBaseQuotasResponseBody
	GetRequestId() *string
}

type ListHBaseQuotasResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RawData            map[string]interface{} `json:"RawData,omitempty" xml:"RawData,omitempty"`
	RequestId          *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHBaseQuotasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHBaseQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListHBaseQuotasResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListHBaseQuotasResponseBody) GetRawData() map[string]interface{} {
	return s.RawData
}

func (s *ListHBaseQuotasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHBaseQuotasResponseBody) SetAccessDeniedDetail(v string) *ListHBaseQuotasResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListHBaseQuotasResponseBody) SetRawData(v map[string]interface{}) *ListHBaseQuotasResponseBody {
	s.RawData = v
	return s
}

func (s *ListHBaseQuotasResponseBody) SetRequestId(v string) *ListHBaseQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHBaseQuotasResponseBody) Validate() error {
	return dara.Validate(s)
}
