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
	// The error code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The maximum number of entries per page used in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. This parameter is empty if no more results are available.
	//
	// example:
	//
	// cae5f900-8b1d-4c0e-9c2a-1a2b3c4d5e6f
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B5AD8B54-4358-5F5B-ACAA-52F2016459C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of template cache objects.
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
