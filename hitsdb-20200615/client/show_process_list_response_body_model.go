// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShowProcessListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ShowProcessListResponseBody
	GetAccessDeniedDetail() *string
	SetRawData(v map[string]interface{}) *ShowProcessListResponseBody
	GetRawData() map[string]interface{}
	SetRequestId(v string) *ShowProcessListResponseBody
	GetRequestId() *string
}

type ShowProcessListResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RawData            map[string]interface{} `json:"RawData,omitempty" xml:"RawData,omitempty"`
	RequestId          *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ShowProcessListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ShowProcessListResponseBody) GoString() string {
	return s.String()
}

func (s *ShowProcessListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ShowProcessListResponseBody) GetRawData() map[string]interface{} {
	return s.RawData
}

func (s *ShowProcessListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ShowProcessListResponseBody) SetAccessDeniedDetail(v string) *ShowProcessListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ShowProcessListResponseBody) SetRawData(v map[string]interface{}) *ShowProcessListResponseBody {
	s.RawData = v
	return s
}

func (s *ShowProcessListResponseBody) SetRequestId(v string) *ShowProcessListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShowProcessListResponseBody) Validate() error {
	return dara.Validate(s)
}
