// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableKnowledgeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetTableKnowledgeInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTableKnowledgeInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetTableKnowledgeInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableKnowledgeInfoResponseBody
	GetSuccess() *bool
	SetTable(v *TableKnowledgeInfo) *GetTableKnowledgeInfoResponseBody
	GetTable() *TableKnowledgeInfo
}

type GetTableKnowledgeInfoResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the request ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The information about the table.
	//
	// example:
	//
	// bill_orders
	Table *TableKnowledgeInfo `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s GetTableKnowledgeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableKnowledgeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableKnowledgeInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTableKnowledgeInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTableKnowledgeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableKnowledgeInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableKnowledgeInfoResponseBody) GetTable() *TableKnowledgeInfo {
	return s.Table
}

func (s *GetTableKnowledgeInfoResponseBody) SetErrorCode(v string) *GetTableKnowledgeInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTableKnowledgeInfoResponseBody) SetErrorMessage(v string) *GetTableKnowledgeInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTableKnowledgeInfoResponseBody) SetRequestId(v string) *GetTableKnowledgeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableKnowledgeInfoResponseBody) SetSuccess(v bool) *GetTableKnowledgeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableKnowledgeInfoResponseBody) SetTable(v *TableKnowledgeInfo) *GetTableKnowledgeInfoResponseBody {
	s.Table = v
	return s
}

func (s *GetTableKnowledgeInfoResponseBody) Validate() error {
	if s.Table != nil {
		if err := s.Table.Validate(); err != nil {
			return err
		}
	}
	return nil
}
