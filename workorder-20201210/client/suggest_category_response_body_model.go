// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuggestCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SuggestCategoryResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SuggestCategoryResponseBody
	GetCode() *string
	SetData(v []map[string]interface{}) *SuggestCategoryResponseBody
	GetData() []map[string]interface{}
	SetHttpStatusCode(v int32) *SuggestCategoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SuggestCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *SuggestCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuggestCategoryResponseBody
	GetSuccess() *bool
}

type SuggestCategoryResponseBody struct {
	AccessDeniedDetail *string                  `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode     *int32                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuggestCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuggestCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *SuggestCategoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SuggestCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *SuggestCategoryResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *SuggestCategoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SuggestCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SuggestCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuggestCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuggestCategoryResponseBody) SetAccessDeniedDetail(v string) *SuggestCategoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SuggestCategoryResponseBody) SetCode(v string) *SuggestCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *SuggestCategoryResponseBody) SetData(v []map[string]interface{}) *SuggestCategoryResponseBody {
	s.Data = v
	return s
}

func (s *SuggestCategoryResponseBody) SetHttpStatusCode(v int32) *SuggestCategoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SuggestCategoryResponseBody) SetMessage(v string) *SuggestCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *SuggestCategoryResponseBody) SetRequestId(v string) *SuggestCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuggestCategoryResponseBody) SetSuccess(v bool) *SuggestCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *SuggestCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
