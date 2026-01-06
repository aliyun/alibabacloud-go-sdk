// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigQueryEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIConfigQueryEntriesRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIConfigQueryEntriesRequest
	GetAuthType() *string
	SetDbName(v string) *ChatBIConfigQueryEntriesRequest
	GetDbName() *string
	SetId(v int32) *ChatBIConfigQueryEntriesRequest
	GetId() *int32
	SetInstanceName(v string) *ChatBIConfigQueryEntriesRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *ChatBIConfigQueryEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ChatBIConfigQueryEntriesRequest
	GetPageSize() *int32
}

type ChatBIConfigQueryEntriesRequest struct {
	AuthMessage *string `json:"AuthMessage,omitempty" xml:"AuthMessage,omitempty"`
	AuthType    *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// 1, 2, 3
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
}

func (s ChatBIConfigQueryEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigQueryEntriesRequest) GoString() string {
	return s.String()
}

func (s *ChatBIConfigQueryEntriesRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIConfigQueryEntriesRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIConfigQueryEntriesRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIConfigQueryEntriesRequest) GetId() *int32 {
	return s.Id
}

func (s *ChatBIConfigQueryEntriesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIConfigQueryEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ChatBIConfigQueryEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ChatBIConfigQueryEntriesRequest) SetAuthMessage(v string) *ChatBIConfigQueryEntriesRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIConfigQueryEntriesRequest) SetAuthType(v string) *ChatBIConfigQueryEntriesRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIConfigQueryEntriesRequest) SetDbName(v string) *ChatBIConfigQueryEntriesRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIConfigQueryEntriesRequest) SetId(v int32) *ChatBIConfigQueryEntriesRequest {
	s.Id = &v
	return s
}

func (s *ChatBIConfigQueryEntriesRequest) SetInstanceName(v string) *ChatBIConfigQueryEntriesRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIConfigQueryEntriesRequest) SetPageNumber(v int32) *ChatBIConfigQueryEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ChatBIConfigQueryEntriesRequest) SetPageSize(v int32) *ChatBIConfigQueryEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ChatBIConfigQueryEntriesRequest) Validate() error {
	return dara.Validate(s)
}
