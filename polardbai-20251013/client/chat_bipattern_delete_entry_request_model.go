// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternDeleteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIPatternDeleteEntryRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIPatternDeleteEntryRequest
	GetAuthType() *string
	SetDbName(v string) *ChatBIPatternDeleteEntryRequest
	GetDbName() *string
	SetId(v string) *ChatBIPatternDeleteEntryRequest
	GetId() *string
	SetInstanceName(v string) *ChatBIPatternDeleteEntryRequest
	GetInstanceName() *string
	SetTableName(v string) *ChatBIPatternDeleteEntryRequest
	GetTableName() *string
}

type ChatBIPatternDeleteEntryRequest struct {
	AuthMessage *string `json:"AuthMessage,omitempty" xml:"AuthMessage,omitempty"`
	AuthType    *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// polar4ai_nl2sql_pattern
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// polar4ai_nl2sql_pattern
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ChatBIPatternDeleteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternDeleteEntryRequest) GoString() string {
	return s.String()
}

func (s *ChatBIPatternDeleteEntryRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIPatternDeleteEntryRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIPatternDeleteEntryRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIPatternDeleteEntryRequest) GetId() *string {
	return s.Id
}

func (s *ChatBIPatternDeleteEntryRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIPatternDeleteEntryRequest) GetTableName() *string {
	return s.TableName
}

func (s *ChatBIPatternDeleteEntryRequest) SetAuthMessage(v string) *ChatBIPatternDeleteEntryRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIPatternDeleteEntryRequest) SetAuthType(v string) *ChatBIPatternDeleteEntryRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIPatternDeleteEntryRequest) SetDbName(v string) *ChatBIPatternDeleteEntryRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIPatternDeleteEntryRequest) SetId(v string) *ChatBIPatternDeleteEntryRequest {
	s.Id = &v
	return s
}

func (s *ChatBIPatternDeleteEntryRequest) SetInstanceName(v string) *ChatBIPatternDeleteEntryRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIPatternDeleteEntryRequest) SetTableName(v string) *ChatBIPatternDeleteEntryRequest {
	s.TableName = &v
	return s
}

func (s *ChatBIPatternDeleteEntryRequest) Validate() error {
	return dara.Validate(s)
}
