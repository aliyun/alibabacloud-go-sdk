// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIFileUploadCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIFileUploadCallbackRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIFileUploadCallbackRequest
	GetAuthType() *string
	SetCharacterSetName(v string) *ChatBIFileUploadCallbackRequest
	GetCharacterSetName() *string
	SetDbName(v string) *ChatBIFileUploadCallbackRequest
	GetDbName() *string
	SetFileName(v string) *ChatBIFileUploadCallbackRequest
	GetFileName() *string
	SetInstanceName(v string) *ChatBIFileUploadCallbackRequest
	GetInstanceName() *string
	SetTableName(v string) *ChatBIFileUploadCallbackRequest
	GetTableName() *string
	SetTableType(v string) *ChatBIFileUploadCallbackRequest
	GetTableType() *string
}

type ChatBIFileUploadCallbackRequest struct {
	AuthMessage *string `json:"AuthMessage,omitempty" xml:"AuthMessage,omitempty"`
	AuthType    *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// gbk, utf-8
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
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
	// 1778851887351348/record_1746670038342.xlsx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// MANAGED_TABLE
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s ChatBIFileUploadCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIFileUploadCallbackRequest) GoString() string {
	return s.String()
}

func (s *ChatBIFileUploadCallbackRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIFileUploadCallbackRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIFileUploadCallbackRequest) GetCharacterSetName() *string {
	return s.CharacterSetName
}

func (s *ChatBIFileUploadCallbackRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIFileUploadCallbackRequest) GetFileName() *string {
	return s.FileName
}

func (s *ChatBIFileUploadCallbackRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIFileUploadCallbackRequest) GetTableName() *string {
	return s.TableName
}

func (s *ChatBIFileUploadCallbackRequest) GetTableType() *string {
	return s.TableType
}

func (s *ChatBIFileUploadCallbackRequest) SetAuthMessage(v string) *ChatBIFileUploadCallbackRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIFileUploadCallbackRequest) SetAuthType(v string) *ChatBIFileUploadCallbackRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIFileUploadCallbackRequest) SetCharacterSetName(v string) *ChatBIFileUploadCallbackRequest {
	s.CharacterSetName = &v
	return s
}

func (s *ChatBIFileUploadCallbackRequest) SetDbName(v string) *ChatBIFileUploadCallbackRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIFileUploadCallbackRequest) SetFileName(v string) *ChatBIFileUploadCallbackRequest {
	s.FileName = &v
	return s
}

func (s *ChatBIFileUploadCallbackRequest) SetInstanceName(v string) *ChatBIFileUploadCallbackRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIFileUploadCallbackRequest) SetTableName(v string) *ChatBIFileUploadCallbackRequest {
	s.TableName = &v
	return s
}

func (s *ChatBIFileUploadCallbackRequest) SetTableType(v string) *ChatBIFileUploadCallbackRequest {
	s.TableType = &v
	return s
}

func (s *ChatBIFileUploadCallbackRequest) Validate() error {
	return dara.Validate(s)
}
