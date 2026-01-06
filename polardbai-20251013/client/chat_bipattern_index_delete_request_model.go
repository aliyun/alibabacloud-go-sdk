// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternIndexDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIPatternIndexDeleteRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIPatternIndexDeleteRequest
	GetAuthType() *string
	SetDbName(v string) *ChatBIPatternIndexDeleteRequest
	GetDbName() *string
	SetInstanceName(v string) *ChatBIPatternIndexDeleteRequest
	GetInstanceName() *string
	SetTableName(v string) *ChatBIPatternIndexDeleteRequest
	GetTableName() *string
}

type ChatBIPatternIndexDeleteRequest struct {
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
	// pc-2ze454l20me07****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pattern_index, pattern_index_1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ChatBIPatternIndexDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternIndexDeleteRequest) GoString() string {
	return s.String()
}

func (s *ChatBIPatternIndexDeleteRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIPatternIndexDeleteRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIPatternIndexDeleteRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIPatternIndexDeleteRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIPatternIndexDeleteRequest) GetTableName() *string {
	return s.TableName
}

func (s *ChatBIPatternIndexDeleteRequest) SetAuthMessage(v string) *ChatBIPatternIndexDeleteRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIPatternIndexDeleteRequest) SetAuthType(v string) *ChatBIPatternIndexDeleteRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIPatternIndexDeleteRequest) SetDbName(v string) *ChatBIPatternIndexDeleteRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIPatternIndexDeleteRequest) SetInstanceName(v string) *ChatBIPatternIndexDeleteRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIPatternIndexDeleteRequest) SetTableName(v string) *ChatBIPatternIndexDeleteRequest {
	s.TableName = &v
	return s
}

func (s *ChatBIPatternIndexDeleteRequest) Validate() error {
	return dara.Validate(s)
}
