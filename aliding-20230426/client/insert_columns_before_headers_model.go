// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertColumnsBeforeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsertColumnsBeforeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *InsertColumnsBeforeHeadersAccountContext) *InsertColumnsBeforeHeaders
	GetAccountContext() *InsertColumnsBeforeHeadersAccountContext
}

type InsertColumnsBeforeHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *InsertColumnsBeforeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s InsertColumnsBeforeHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsertColumnsBeforeHeaders) GoString() string {
	return s.String()
}

func (s *InsertColumnsBeforeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsertColumnsBeforeHeaders) GetAccountContext() *InsertColumnsBeforeHeadersAccountContext {
	return s.AccountContext
}

func (s *InsertColumnsBeforeHeaders) SetCommonHeaders(v map[string]*string) *InsertColumnsBeforeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertColumnsBeforeHeaders) SetAccountContext(v *InsertColumnsBeforeHeadersAccountContext) *InsertColumnsBeforeHeaders {
	s.AccountContext = v
	return s
}

func (s *InsertColumnsBeforeHeaders) Validate() error {
	return dara.Validate(s)
}

type InsertColumnsBeforeHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s InsertColumnsBeforeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s InsertColumnsBeforeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *InsertColumnsBeforeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *InsertColumnsBeforeHeadersAccountContext) SetAccountId(v string) *InsertColumnsBeforeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *InsertColumnsBeforeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
