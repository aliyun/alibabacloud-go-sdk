// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConcurrencyControlKeywordsFromSqlTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextRequest
	GetConsoleContext() *string
	SetInstanceId(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextRequest
	GetInstanceId() *string
	SetSqlText(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextRequest
	GetSqlText() *string
}

type GetSqlConcurrencyControlKeywordsFromSqlTextRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze5hpn2b99d2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The SQL statement based on which a throttling keyword string is to be generated.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM test where name = \\"das\\"
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
}

func (s GetSqlConcurrencyControlKeywordsFromSqlTextRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlKeywordsFromSqlTextRequest) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextRequest) GetSqlText() *string {
	return s.SqlText
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextRequest) SetConsoleContext(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextRequest) SetInstanceId(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextRequest) SetSqlText(v string) *GetSqlConcurrencyControlKeywordsFromSqlTextRequest {
	s.SqlText = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextRequest) Validate() error {
	return dara.Validate(s)
}
