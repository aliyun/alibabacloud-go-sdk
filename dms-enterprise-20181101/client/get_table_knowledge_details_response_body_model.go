// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableKnowledgeDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TableKnowledgeVO) *GetTableKnowledgeDetailsResponseBody
	GetData() *TableKnowledgeVO
	SetErrorCode(v string) *GetTableKnowledgeDetailsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTableKnowledgeDetailsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetTableKnowledgeDetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableKnowledgeDetailsResponseBody
	GetSuccess() *bool
}

type GetTableKnowledgeDetailsResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Data *TableKnowledgeVO `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s GetTableKnowledgeDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableKnowledgeDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableKnowledgeDetailsResponseBody) GetData() *TableKnowledgeVO {
	return s.Data
}

func (s *GetTableKnowledgeDetailsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTableKnowledgeDetailsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTableKnowledgeDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableKnowledgeDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableKnowledgeDetailsResponseBody) SetData(v *TableKnowledgeVO) *GetTableKnowledgeDetailsResponseBody {
	s.Data = v
	return s
}

func (s *GetTableKnowledgeDetailsResponseBody) SetErrorCode(v string) *GetTableKnowledgeDetailsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTableKnowledgeDetailsResponseBody) SetErrorMessage(v string) *GetTableKnowledgeDetailsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTableKnowledgeDetailsResponseBody) SetRequestId(v string) *GetTableKnowledgeDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableKnowledgeDetailsResponseBody) SetSuccess(v bool) *GetTableKnowledgeDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableKnowledgeDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
