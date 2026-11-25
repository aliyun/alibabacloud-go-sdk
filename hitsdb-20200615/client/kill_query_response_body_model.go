// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *KillQueryResponseBody
	GetAccessDeniedDetail() *string
	SetRawData(v map[string]interface{}) *KillQueryResponseBody
	GetRawData() map[string]interface{}
	SetRequestId(v string) *KillQueryResponseBody
	GetRequestId() *string
}

type KillQueryResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RawData            map[string]interface{} `json:"RawData,omitempty" xml:"RawData,omitempty"`
	RequestId          *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KillQueryResponseBody) GoString() string {
	return s.String()
}

func (s *KillQueryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *KillQueryResponseBody) GetRawData() map[string]interface{} {
	return s.RawData
}

func (s *KillQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KillQueryResponseBody) SetAccessDeniedDetail(v string) *KillQueryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *KillQueryResponseBody) SetRawData(v map[string]interface{}) *KillQueryResponseBody {
	s.RawData = v
	return s
}

func (s *KillQueryResponseBody) SetRequestId(v string) *KillQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *KillQueryResponseBody) Validate() error {
	return dara.Validate(s)
}
