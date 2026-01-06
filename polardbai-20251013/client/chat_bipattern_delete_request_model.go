// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIPatternDeleteRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIPatternDeleteRequest
	GetAuthType() *string
	SetDbName(v string) *ChatBIPatternDeleteRequest
	GetDbName() *string
	SetInstanceName(v string) *ChatBIPatternDeleteRequest
	GetInstanceName() *string
	SetTableName(v string) *ChatBIPatternDeleteRequest
	GetTableName() *string
}

type ChatBIPatternDeleteRequest struct {
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
	// polar4ai_nl2sql_pattern
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ChatBIPatternDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternDeleteRequest) GoString() string {
	return s.String()
}

func (s *ChatBIPatternDeleteRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIPatternDeleteRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIPatternDeleteRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIPatternDeleteRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIPatternDeleteRequest) GetTableName() *string {
	return s.TableName
}

func (s *ChatBIPatternDeleteRequest) SetAuthMessage(v string) *ChatBIPatternDeleteRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIPatternDeleteRequest) SetAuthType(v string) *ChatBIPatternDeleteRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIPatternDeleteRequest) SetDbName(v string) *ChatBIPatternDeleteRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIPatternDeleteRequest) SetInstanceName(v string) *ChatBIPatternDeleteRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIPatternDeleteRequest) SetTableName(v string) *ChatBIPatternDeleteRequest {
	s.TableName = &v
	return s
}

func (s *ChatBIPatternDeleteRequest) Validate() error {
	return dara.Validate(s)
}
