// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTemplateCacheResponseBody
	GetCode() *string
	SetMaxResults(v int32) *ListTemplateCacheResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListTemplateCacheResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTemplateCacheResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTemplateCacheResponseBody
	GetRequestId() *string
	SetTemplateCaches(v []*PublicTemplateCache) *ListTemplateCacheResponseBody
	GetTemplateCaches() []*PublicTemplateCache
}

type ListTemplateCacheResponseBody struct {
	Code           *string                `json:"code,omitempty" xml:"code,omitempty"`
	MaxResults     *int32                 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Message        *string                `json:"message,omitempty" xml:"message,omitempty"`
	NextToken      *string                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId      *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TemplateCaches []*PublicTemplateCache `json:"templateCaches,omitempty" xml:"templateCaches,omitempty" type:"Repeated"`
}

func (s ListTemplateCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateCacheResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateCacheResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTemplateCacheResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplateCacheResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTemplateCacheResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplateCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplateCacheResponseBody) GetTemplateCaches() []*PublicTemplateCache {
	return s.TemplateCaches
}

func (s *ListTemplateCacheResponseBody) SetCode(v string) *ListTemplateCacheResponseBody {
	s.Code = &v
	return s
}

func (s *ListTemplateCacheResponseBody) SetMaxResults(v int32) *ListTemplateCacheResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateCacheResponseBody) SetMessage(v string) *ListTemplateCacheResponseBody {
	s.Message = &v
	return s
}

func (s *ListTemplateCacheResponseBody) SetNextToken(v string) *ListTemplateCacheResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplateCacheResponseBody) SetRequestId(v string) *ListTemplateCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateCacheResponseBody) SetTemplateCaches(v []*PublicTemplateCache) *ListTemplateCacheResponseBody {
	s.TemplateCaches = v
	return s
}

func (s *ListTemplateCacheResponseBody) Validate() error {
	if s.TemplateCaches != nil {
		for _, item := range s.TemplateCaches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
