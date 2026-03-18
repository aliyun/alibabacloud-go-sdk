// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantArticlesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListTextbookAssistantArticlesResponseBodyData) *ListTextbookAssistantArticlesResponseBody
	GetData() []*ListTextbookAssistantArticlesResponseBodyData
	SetErrCode(v string) *ListTextbookAssistantArticlesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListTextbookAssistantArticlesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ListTextbookAssistantArticlesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListTextbookAssistantArticlesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTextbookAssistantArticlesResponseBody
	GetSuccess() *bool
}

type ListTextbookAssistantArticlesResponseBody struct {
	Data []*ListTextbookAssistantArticlesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 70412360-4272-571A-827D-84C2C07C450F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTextbookAssistantArticlesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticlesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticlesResponseBody) GetData() []*ListTextbookAssistantArticlesResponseBodyData {
	return s.Data
}

func (s *ListTextbookAssistantArticlesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListTextbookAssistantArticlesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListTextbookAssistantArticlesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTextbookAssistantArticlesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTextbookAssistantArticlesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTextbookAssistantArticlesResponseBody) SetData(v []*ListTextbookAssistantArticlesResponseBodyData) *ListTextbookAssistantArticlesResponseBody {
	s.Data = v
	return s
}

func (s *ListTextbookAssistantArticlesResponseBody) SetErrCode(v string) *ListTextbookAssistantArticlesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTextbookAssistantArticlesResponseBody) SetErrMessage(v string) *ListTextbookAssistantArticlesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTextbookAssistantArticlesResponseBody) SetHttpStatusCode(v int32) *ListTextbookAssistantArticlesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTextbookAssistantArticlesResponseBody) SetRequestId(v string) *ListTextbookAssistantArticlesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTextbookAssistantArticlesResponseBody) SetSuccess(v bool) *ListTextbookAssistantArticlesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTextbookAssistantArticlesResponseBody) Validate() error {
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

type ListTextbookAssistantArticlesResponseBodyData struct {
	// example:
	//
	// 0c05700d4d9411efbe6e0c42a106bb02
	ArticleId *string `json:"articleId,omitempty" xml:"articleId,omitempty"`
}

func (s ListTextbookAssistantArticlesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantArticlesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticlesResponseBodyData) GetArticleId() *string {
	return s.ArticleId
}

func (s *ListTextbookAssistantArticlesResponseBodyData) SetArticleId(v string) *ListTextbookAssistantArticlesResponseBodyData {
	s.ArticleId = &v
	return s
}

func (s *ListTextbookAssistantArticlesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
