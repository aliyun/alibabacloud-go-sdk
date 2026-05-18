// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*OneMetaKnowledgeBase) *ListKnowledgeBasesResponseBody
	GetData() []*OneMetaKnowledgeBase
	SetErrorCode(v string) *ListKnowledgeBasesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListKnowledgeBasesResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListKnowledgeBasesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListKnowledgeBasesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListKnowledgeBasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKnowledgeBasesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListKnowledgeBasesResponseBody
	GetTotalCount() *int64
}

type ListKnowledgeBasesResponseBody struct {
	Data []*OneMetaKnowledgeBase `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter Tag is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// NesLoKLEdIZrKhDT7I2gS****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKnowledgeBasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesResponseBody) GetData() []*OneMetaKnowledgeBase {
	return s.Data
}

func (s *ListKnowledgeBasesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListKnowledgeBasesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListKnowledgeBasesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBasesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgeBasesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListKnowledgeBasesResponseBody) SetData(v []*OneMetaKnowledgeBase) *ListKnowledgeBasesResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetErrorCode(v string) *ListKnowledgeBasesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetErrorMessage(v string) *ListKnowledgeBasesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetMaxResults(v int32) *ListKnowledgeBasesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetNextToken(v string) *ListKnowledgeBasesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetRequestId(v string) *ListKnowledgeBasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetSuccess(v bool) *ListKnowledgeBasesResponseBody {
	s.Success = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) SetTotalCount(v int64) *ListKnowledgeBasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKnowledgeBasesResponseBody) Validate() error {
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
