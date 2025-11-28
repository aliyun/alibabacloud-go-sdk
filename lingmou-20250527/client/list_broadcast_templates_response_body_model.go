// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBroadcastTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBroadcastTemplatesResponseBody
	GetCode() *string
	SetData(v []*BroadcastTemplate) *ListBroadcastTemplatesResponseBody
	GetData() []*BroadcastTemplate
	SetHttpStatusCode(v int32) *ListBroadcastTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListBroadcastTemplatesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListBroadcastTemplatesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListBroadcastTemplatesResponseBody
	GetNextToken() *string
	SetPage(v int32) *ListBroadcastTemplatesResponseBody
	GetPage() *int32
	SetRequestId(v string) *ListBroadcastTemplatesResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListBroadcastTemplatesResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListBroadcastTemplatesResponseBody
	GetSuccess() *bool
}

type ListBroadcastTemplatesResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string              `json:"code,omitempty" xml:"code,omitempty"`
	Data []*BroadcastTemplate `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	MaxResults     *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// SUCCESS
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 7239F9E5-A4DB-55BA-B701-0CE47DBDB0A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListBroadcastTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBroadcastTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBroadcastTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBroadcastTemplatesResponseBody) GetData() []*BroadcastTemplate {
	return s.Data
}

func (s *ListBroadcastTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBroadcastTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBroadcastTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBroadcastTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBroadcastTemplatesResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *ListBroadcastTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBroadcastTemplatesResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListBroadcastTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBroadcastTemplatesResponseBody) SetCode(v string) *ListBroadcastTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListBroadcastTemplatesResponseBody) SetData(v []*BroadcastTemplate) *ListBroadcastTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListBroadcastTemplatesResponseBody) SetHttpStatusCode(v int32) *ListBroadcastTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBroadcastTemplatesResponseBody) SetMaxResults(v int32) *ListBroadcastTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBroadcastTemplatesResponseBody) SetMessage(v string) *ListBroadcastTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListBroadcastTemplatesResponseBody) SetNextToken(v string) *ListBroadcastTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBroadcastTemplatesResponseBody) SetPage(v int32) *ListBroadcastTemplatesResponseBody {
	s.Page = &v
	return s
}

func (s *ListBroadcastTemplatesResponseBody) SetRequestId(v string) *ListBroadcastTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBroadcastTemplatesResponseBody) SetSize(v int32) *ListBroadcastTemplatesResponseBody {
	s.Size = &v
	return s
}

func (s *ListBroadcastTemplatesResponseBody) SetSuccess(v bool) *ListBroadcastTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListBroadcastTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
