// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTemplatesResponseBody
	GetCode() *string
	SetMaxResults(v int32) *ListTemplatesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListTemplatesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*PublicTemplate) *ListTemplatesResponseBody
	GetTemplates() []*PublicTemplate
}

type ListTemplatesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// eyJNYXhSZXN1bHRzIjoxMH0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B5AD8B54-4358-5F5B-ACAA-52F2016459C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of templates.
	Templates []*PublicTemplate `json:"templates,omitempty" xml:"templates,omitempty" type:"Repeated"`
}

func (s ListTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplatesResponseBody) GetTemplates() []*PublicTemplate {
	return s.Templates
}

func (s *ListTemplatesResponseBody) SetCode(v string) *ListTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTemplatesResponseBody) SetMaxResults(v int32) *ListTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesResponseBody) SetMessage(v string) *ListTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTemplatesResponseBody) SetNextToken(v string) *ListTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*PublicTemplate) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
