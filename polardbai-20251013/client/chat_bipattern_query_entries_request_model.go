// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternQueryEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ChatBIPatternQueryEntriesRequest
	GetDbName() *string
	SetId(v int32) *ChatBIPatternQueryEntriesRequest
	GetId() *int32
	SetInstanceName(v string) *ChatBIPatternQueryEntriesRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *ChatBIPatternQueryEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ChatBIPatternQueryEntriesRequest
	GetPageSize() *int32
	SetTableName(v string) *ChatBIPatternQueryEntriesRequest
	GetTableName() *string
}

type ChatBIPatternQueryEntriesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// polar4ai_nl2sql_pattern
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ChatBIPatternQueryEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternQueryEntriesRequest) GoString() string {
	return s.String()
}

func (s *ChatBIPatternQueryEntriesRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIPatternQueryEntriesRequest) GetId() *int32 {
	return s.Id
}

func (s *ChatBIPatternQueryEntriesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIPatternQueryEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ChatBIPatternQueryEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ChatBIPatternQueryEntriesRequest) GetTableName() *string {
	return s.TableName
}

func (s *ChatBIPatternQueryEntriesRequest) SetDbName(v string) *ChatBIPatternQueryEntriesRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIPatternQueryEntriesRequest) SetId(v int32) *ChatBIPatternQueryEntriesRequest {
	s.Id = &v
	return s
}

func (s *ChatBIPatternQueryEntriesRequest) SetInstanceName(v string) *ChatBIPatternQueryEntriesRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIPatternQueryEntriesRequest) SetPageNumber(v int32) *ChatBIPatternQueryEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ChatBIPatternQueryEntriesRequest) SetPageSize(v int32) *ChatBIPatternQueryEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ChatBIPatternQueryEntriesRequest) SetTableName(v string) *ChatBIPatternQueryEntriesRequest {
	s.TableName = &v
	return s
}

func (s *ChatBIPatternQueryEntriesRequest) Validate() error {
	return dara.Validate(s)
}
