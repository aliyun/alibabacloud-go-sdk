// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ChatBIConfigDeleteRequest
	GetDbName() *string
	SetInstanceName(v string) *ChatBIConfigDeleteRequest
	GetInstanceName() *string
}

type ChatBIConfigDeleteRequest struct {
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

func (s ChatBIConfigDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigDeleteRequest) GoString() string {
	return s.String()
}

func (s *ChatBIConfigDeleteRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIConfigDeleteRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIConfigDeleteRequest) SetDbName(v string) *ChatBIConfigDeleteRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIConfigDeleteRequest) SetInstanceName(v string) *ChatBIConfigDeleteRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIConfigDeleteRequest) Validate() error {
	return dara.Validate(s)
}
