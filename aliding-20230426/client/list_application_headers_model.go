// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListApplicationHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListApplicationHeadersAccountContext) *ListApplicationHeaders
	GetAccountContext() *ListApplicationHeadersAccountContext
}

type ListApplicationHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListApplicationHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListApplicationHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationHeaders) GoString() string {
	return s.String()
}

func (s *ListApplicationHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListApplicationHeaders) GetAccountContext() *ListApplicationHeadersAccountContext {
	return s.AccountContext
}

func (s *ListApplicationHeaders) SetCommonHeaders(v map[string]*string) *ListApplicationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListApplicationHeaders) SetAccountContext(v *ListApplicationHeadersAccountContext) *ListApplicationHeaders {
	s.AccountContext = v
	return s
}

func (s *ListApplicationHeaders) Validate() error {
	return dara.Validate(s)
}

type ListApplicationHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListApplicationHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListApplicationHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListApplicationHeadersAccountContext) SetAccountId(v string) *ListApplicationHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListApplicationHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
