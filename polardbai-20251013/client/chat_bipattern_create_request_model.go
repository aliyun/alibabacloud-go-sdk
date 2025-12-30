// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ChatBIPatternCreateRequest
	GetDbName() *string
	SetInstanceName(v string) *ChatBIPatternCreateRequest
	GetInstanceName() *string
	SetTableNameSuffix(v string) *ChatBIPatternCreateRequest
	GetTableNameSuffix() *string
}

type ChatBIPatternCreateRequest struct {
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
	// example:
	//
	// 空字符串
	TableNameSuffix *string `json:"TableNameSuffix,omitempty" xml:"TableNameSuffix,omitempty"`
}

func (s ChatBIPatternCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternCreateRequest) GoString() string {
	return s.String()
}

func (s *ChatBIPatternCreateRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIPatternCreateRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIPatternCreateRequest) GetTableNameSuffix() *string {
	return s.TableNameSuffix
}

func (s *ChatBIPatternCreateRequest) SetDbName(v string) *ChatBIPatternCreateRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIPatternCreateRequest) SetInstanceName(v string) *ChatBIPatternCreateRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIPatternCreateRequest) SetTableNameSuffix(v string) *ChatBIPatternCreateRequest {
	s.TableNameSuffix = &v
	return s
}

func (s *ChatBIPatternCreateRequest) Validate() error {
	return dara.Validate(s)
}
