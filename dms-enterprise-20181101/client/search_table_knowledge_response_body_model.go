// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTableKnowledgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SearchTableKnowledgeResponseBodyData) *SearchTableKnowledgeResponseBody
	GetData() *SearchTableKnowledgeResponseBodyData
	SetErrorCode(v string) *SearchTableKnowledgeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SearchTableKnowledgeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SearchTableKnowledgeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchTableKnowledgeResponseBody
	GetSuccess() *bool
}

type SearchTableKnowledgeResponseBody struct {
	Data *SearchTableKnowledgeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchTableKnowledgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTableKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTableKnowledgeResponseBody) GetData() *SearchTableKnowledgeResponseBodyData {
	return s.Data
}

func (s *SearchTableKnowledgeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchTableKnowledgeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SearchTableKnowledgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTableKnowledgeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchTableKnowledgeResponseBody) SetData(v *SearchTableKnowledgeResponseBodyData) *SearchTableKnowledgeResponseBody {
	s.Data = v
	return s
}

func (s *SearchTableKnowledgeResponseBody) SetErrorCode(v string) *SearchTableKnowledgeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchTableKnowledgeResponseBody) SetErrorMessage(v string) *SearchTableKnowledgeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchTableKnowledgeResponseBody) SetRequestId(v string) *SearchTableKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTableKnowledgeResponseBody) SetSuccess(v bool) *SearchTableKnowledgeResponseBody {
	s.Success = &v
	return s
}

func (s *SearchTableKnowledgeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchTableKnowledgeResponseBodyData struct {
	// example:
	//
	// SQL修复结果...
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// f63a6eed-0e3c-4564-8533-b1295db8d6ff
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s SearchTableKnowledgeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchTableKnowledgeResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchTableKnowledgeResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *SearchTableKnowledgeResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *SearchTableKnowledgeResponseBodyData) SetContent(v string) *SearchTableKnowledgeResponseBodyData {
	s.Content = &v
	return s
}

func (s *SearchTableKnowledgeResponseBodyData) SetSessionId(v string) *SearchTableKnowledgeResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *SearchTableKnowledgeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
