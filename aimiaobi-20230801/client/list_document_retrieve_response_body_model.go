// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentRetrieveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDocumentRetrieveResponseBody
	GetCode() *string
	SetData(v []*ListDocumentRetrieveResponseBodyData) *ListDocumentRetrieveResponseBody
	GetData() []*ListDocumentRetrieveResponseBodyData
	SetHttpStatusCode(v int32) *ListDocumentRetrieveResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListDocumentRetrieveResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListDocumentRetrieveResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListDocumentRetrieveResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDocumentRetrieveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDocumentRetrieveResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListDocumentRetrieveResponseBody
	GetTotalCount() *int32
}

type ListDocumentRetrieveResponseBody struct {
	// example:
	//
	// successful
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// data
	Data []*ListDocumentRetrieveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 71
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// cEoBWREAXdxaOyjq/cqAbg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDocumentRetrieveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentRetrieveResponseBody) GoString() string {
	return s.String()
}

func (s *ListDocumentRetrieveResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDocumentRetrieveResponseBody) GetData() []*ListDocumentRetrieveResponseBodyData {
	return s.Data
}

func (s *ListDocumentRetrieveResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDocumentRetrieveResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocumentRetrieveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDocumentRetrieveResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocumentRetrieveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDocumentRetrieveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDocumentRetrieveResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDocumentRetrieveResponseBody) SetCode(v string) *ListDocumentRetrieveResponseBody {
	s.Code = &v
	return s
}

func (s *ListDocumentRetrieveResponseBody) SetData(v []*ListDocumentRetrieveResponseBodyData) *ListDocumentRetrieveResponseBody {
	s.Data = v
	return s
}

func (s *ListDocumentRetrieveResponseBody) SetHttpStatusCode(v int32) *ListDocumentRetrieveResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDocumentRetrieveResponseBody) SetMaxResults(v int32) *ListDocumentRetrieveResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDocumentRetrieveResponseBody) SetMessage(v string) *ListDocumentRetrieveResponseBody {
	s.Message = &v
	return s
}

func (s *ListDocumentRetrieveResponseBody) SetNextToken(v string) *ListDocumentRetrieveResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDocumentRetrieveResponseBody) SetRequestId(v string) *ListDocumentRetrieveResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDocumentRetrieveResponseBody) SetSuccess(v bool) *ListDocumentRetrieveResponseBody {
	s.Success = &v
	return s
}

func (s *ListDocumentRetrieveResponseBody) SetTotalCount(v int32) *ListDocumentRetrieveResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDocumentRetrieveResponseBody) Validate() error {
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

type ListDocumentRetrieveResponseBodyData struct {
	// example:
	//
	// 文章正文
	Essay *string `json:"Essay,omitempty" xml:"Essay,omitempty"`
	// example:
	//
	// 发布机构
	IssuingAuthority *string `json:"IssuingAuthority,omitempty" xml:"IssuingAuthority,omitempty"`
	// example:
	//
	// 文章链接
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// example:
	//
	// 2023-02-01
	PublicationDate *string `json:"PublicationDate,omitempty" xml:"PublicationDate,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListDocumentRetrieveResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentRetrieveResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDocumentRetrieveResponseBodyData) GetEssay() *string {
	return s.Essay
}

func (s *ListDocumentRetrieveResponseBodyData) GetIssuingAuthority() *string {
	return s.IssuingAuthority
}

func (s *ListDocumentRetrieveResponseBodyData) GetLink() *string {
	return s.Link
}

func (s *ListDocumentRetrieveResponseBodyData) GetPublicationDate() *string {
	return s.PublicationDate
}

func (s *ListDocumentRetrieveResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListDocumentRetrieveResponseBodyData) SetEssay(v string) *ListDocumentRetrieveResponseBodyData {
	s.Essay = &v
	return s
}

func (s *ListDocumentRetrieveResponseBodyData) SetIssuingAuthority(v string) *ListDocumentRetrieveResponseBodyData {
	s.IssuingAuthority = &v
	return s
}

func (s *ListDocumentRetrieveResponseBodyData) SetLink(v string) *ListDocumentRetrieveResponseBodyData {
	s.Link = &v
	return s
}

func (s *ListDocumentRetrieveResponseBodyData) SetPublicationDate(v string) *ListDocumentRetrieveResponseBodyData {
	s.PublicationDate = &v
	return s
}

func (s *ListDocumentRetrieveResponseBodyData) SetTitle(v string) *ListDocumentRetrieveResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListDocumentRetrieveResponseBodyData) Validate() error {
	return dara.Validate(s)
}
