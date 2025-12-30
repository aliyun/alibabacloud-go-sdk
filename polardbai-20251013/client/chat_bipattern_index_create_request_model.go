// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternIndexCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ChatBIPatternIndexCreateRequest
	GetDbName() *string
	SetInstanceName(v string) *ChatBIPatternIndexCreateRequest
	GetInstanceName() *string
	SetPatternTableName(v string) *ChatBIPatternIndexCreateRequest
	GetPatternTableName() *string
	SetTableNameSuffix(v string) *ChatBIPatternIndexCreateRequest
	GetTableNameSuffix() *string
}

type ChatBIPatternIndexCreateRequest struct {
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
	// polar4ai_nl2sql_pattern, polar4ai_nl2sql_pattern_1
	PatternTableName *string `json:"PatternTableName,omitempty" xml:"PatternTableName,omitempty"`
	// example:
	//
	// 空字符串
	TableNameSuffix *string `json:"TableNameSuffix,omitempty" xml:"TableNameSuffix,omitempty"`
}

func (s ChatBIPatternIndexCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternIndexCreateRequest) GoString() string {
	return s.String()
}

func (s *ChatBIPatternIndexCreateRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIPatternIndexCreateRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIPatternIndexCreateRequest) GetPatternTableName() *string {
	return s.PatternTableName
}

func (s *ChatBIPatternIndexCreateRequest) GetTableNameSuffix() *string {
	return s.TableNameSuffix
}

func (s *ChatBIPatternIndexCreateRequest) SetDbName(v string) *ChatBIPatternIndexCreateRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIPatternIndexCreateRequest) SetInstanceName(v string) *ChatBIPatternIndexCreateRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIPatternIndexCreateRequest) SetPatternTableName(v string) *ChatBIPatternIndexCreateRequest {
	s.PatternTableName = &v
	return s
}

func (s *ChatBIPatternIndexCreateRequest) SetTableNameSuffix(v string) *ChatBIPatternIndexCreateRequest {
	s.TableNameSuffix = &v
	return s
}

func (s *ChatBIPatternIndexCreateRequest) Validate() error {
	return dara.Validate(s)
}
