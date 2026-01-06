// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIConfigCreateRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIConfigCreateRequest
	GetAuthType() *string
	SetDbName(v string) *ChatBIConfigCreateRequest
	GetDbName() *string
	SetInstanceName(v string) *ChatBIConfigCreateRequest
	GetInstanceName() *string
}

type ChatBIConfigCreateRequest struct {
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
}

func (s ChatBIConfigCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigCreateRequest) GoString() string {
	return s.String()
}

func (s *ChatBIConfigCreateRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIConfigCreateRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIConfigCreateRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIConfigCreateRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIConfigCreateRequest) SetAuthMessage(v string) *ChatBIConfigCreateRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIConfigCreateRequest) SetAuthType(v string) *ChatBIConfigCreateRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIConfigCreateRequest) SetDbName(v string) *ChatBIConfigCreateRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIConfigCreateRequest) SetInstanceName(v string) *ChatBIConfigCreateRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIConfigCreateRequest) Validate() error {
	return dara.Validate(s)
}
