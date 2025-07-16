// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScenegroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateScenegroupHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateScenegroupHeadersAccountContext) *CreateScenegroupHeaders
	GetAccountContext() *CreateScenegroupHeadersAccountContext
}

type CreateScenegroupHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateScenegroupHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateScenegroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateScenegroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateScenegroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateScenegroupHeaders) GetAccountContext() *CreateScenegroupHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateScenegroupHeaders) SetCommonHeaders(v map[string]*string) *CreateScenegroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateScenegroupHeaders) SetAccountContext(v *CreateScenegroupHeadersAccountContext) *CreateScenegroupHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateScenegroupHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateScenegroupHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateScenegroupHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateScenegroupHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateScenegroupHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateScenegroupHeadersAccountContext) SetAccountId(v string) *CreateScenegroupHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateScenegroupHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
