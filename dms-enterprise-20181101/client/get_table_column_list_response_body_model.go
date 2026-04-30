// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ColumnKnowledgeVO) *GetTableColumnListResponseBody
	GetData() []*ColumnKnowledgeVO
	SetErrorCode(v string) *GetTableColumnListResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTableColumnListResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetTableColumnListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableColumnListResponseBody
	GetSuccess() *bool
}

type GetTableColumnListResponseBody struct {
	Data []*ColumnKnowledgeVO `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s GetTableColumnListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableColumnListResponseBody) GetData() []*ColumnKnowledgeVO {
	return s.Data
}

func (s *GetTableColumnListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTableColumnListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTableColumnListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableColumnListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableColumnListResponseBody) SetData(v []*ColumnKnowledgeVO) *GetTableColumnListResponseBody {
	s.Data = v
	return s
}

func (s *GetTableColumnListResponseBody) SetErrorCode(v string) *GetTableColumnListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTableColumnListResponseBody) SetErrorMessage(v string) *GetTableColumnListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTableColumnListResponseBody) SetRequestId(v string) *GetTableColumnListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableColumnListResponseBody) SetSuccess(v bool) *GetTableColumnListResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableColumnListResponseBody) Validate() error {
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
