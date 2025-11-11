// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuditTermsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAuditTermsResponseBody
	GetCode() *string
	SetData(v []*ListAuditTermsResponseBodyData) *ListAuditTermsResponseBody
	GetData() []*ListAuditTermsResponseBodyData
	SetHttpStatusCode(v int32) *ListAuditTermsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListAuditTermsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAuditTermsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAuditTermsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAuditTermsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAuditTermsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListAuditTermsResponseBody
	GetTotalCount() *int32
}

type ListAuditTermsResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListAuditTermsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 77
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// x\\"x\\"x
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
	// 58
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuditTermsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuditTermsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuditTermsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAuditTermsResponseBody) GetData() []*ListAuditTermsResponseBodyData {
	return s.Data
}

func (s *ListAuditTermsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAuditTermsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuditTermsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAuditTermsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuditTermsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuditTermsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAuditTermsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAuditTermsResponseBody) SetCode(v string) *ListAuditTermsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAuditTermsResponseBody) SetData(v []*ListAuditTermsResponseBodyData) *ListAuditTermsResponseBody {
	s.Data = v
	return s
}

func (s *ListAuditTermsResponseBody) SetHttpStatusCode(v int32) *ListAuditTermsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAuditTermsResponseBody) SetMaxResults(v int32) *ListAuditTermsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAuditTermsResponseBody) SetMessage(v string) *ListAuditTermsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAuditTermsResponseBody) SetNextToken(v string) *ListAuditTermsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAuditTermsResponseBody) SetRequestId(v string) *ListAuditTermsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuditTermsResponseBody) SetSuccess(v bool) *ListAuditTermsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAuditTermsResponseBody) SetTotalCount(v int32) *ListAuditTermsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuditTermsResponseBody) Validate() error {
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

type ListAuditTermsResponseBodyData struct {
	ExceptionWord []*string `json:"ExceptionWord,omitempty" xml:"ExceptionWord,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 龘
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 龘(dá)
	SuggestWord *string `json:"SuggestWord,omitempty" xml:"SuggestWord,omitempty"`
	// example:
	//
	// 龙行龘龘出自四库本《玉篇》23龙部第8字，文字释义为群龙腾飞的样子，昂扬而热烈。
	TermsDesc *string `json:"TermsDesc,omitempty" xml:"TermsDesc,omitempty"`
	TermsName *string `json:"TermsName,omitempty" xml:"TermsName,omitempty"`
}

func (s ListAuditTermsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAuditTermsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAuditTermsResponseBodyData) GetExceptionWord() []*string {
	return s.ExceptionWord
}

func (s *ListAuditTermsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListAuditTermsResponseBodyData) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAuditTermsResponseBodyData) GetSuggestWord() *string {
	return s.SuggestWord
}

func (s *ListAuditTermsResponseBodyData) GetTermsDesc() *string {
	return s.TermsDesc
}

func (s *ListAuditTermsResponseBodyData) GetTermsName() *string {
	return s.TermsName
}

func (s *ListAuditTermsResponseBodyData) SetExceptionWord(v []*string) *ListAuditTermsResponseBodyData {
	s.ExceptionWord = v
	return s
}

func (s *ListAuditTermsResponseBodyData) SetId(v string) *ListAuditTermsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAuditTermsResponseBodyData) SetKeyword(v string) *ListAuditTermsResponseBodyData {
	s.Keyword = &v
	return s
}

func (s *ListAuditTermsResponseBodyData) SetSuggestWord(v string) *ListAuditTermsResponseBodyData {
	s.SuggestWord = &v
	return s
}

func (s *ListAuditTermsResponseBodyData) SetTermsDesc(v string) *ListAuditTermsResponseBodyData {
	s.TermsDesc = &v
	return s
}

func (s *ListAuditTermsResponseBodyData) SetTermsName(v string) *ListAuditTermsResponseBodyData {
	s.TermsName = &v
	return s
}

func (s *ListAuditTermsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
