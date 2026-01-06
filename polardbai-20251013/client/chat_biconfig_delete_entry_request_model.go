// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigDeleteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIConfigDeleteEntryRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIConfigDeleteEntryRequest
	GetAuthType() *string
	SetDbName(v string) *ChatBIConfigDeleteEntryRequest
	GetDbName() *string
	SetId(v string) *ChatBIConfigDeleteEntryRequest
	GetId() *string
	SetInstanceName(v string) *ChatBIConfigDeleteEntryRequest
	GetInstanceName() *string
}

type ChatBIConfigDeleteEntryRequest struct {
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
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ChatBIConfigDeleteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigDeleteEntryRequest) GoString() string {
	return s.String()
}

func (s *ChatBIConfigDeleteEntryRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIConfigDeleteEntryRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIConfigDeleteEntryRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIConfigDeleteEntryRequest) GetId() *string {
	return s.Id
}

func (s *ChatBIConfigDeleteEntryRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIConfigDeleteEntryRequest) SetAuthMessage(v string) *ChatBIConfigDeleteEntryRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIConfigDeleteEntryRequest) SetAuthType(v string) *ChatBIConfigDeleteEntryRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIConfigDeleteEntryRequest) SetDbName(v string) *ChatBIConfigDeleteEntryRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIConfigDeleteEntryRequest) SetId(v string) *ChatBIConfigDeleteEntryRequest {
	s.Id = &v
	return s
}

func (s *ChatBIConfigDeleteEntryRequest) SetInstanceName(v string) *ChatBIConfigDeleteEntryRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIConfigDeleteEntryRequest) Validate() error {
	return dara.Validate(s)
}
