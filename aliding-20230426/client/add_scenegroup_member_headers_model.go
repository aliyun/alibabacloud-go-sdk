// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddScenegroupMemberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddScenegroupMemberHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddScenegroupMemberHeadersAccountContext) *AddScenegroupMemberHeaders
	GetAccountContext() *AddScenegroupMemberHeadersAccountContext
}

type AddScenegroupMemberHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddScenegroupMemberHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddScenegroupMemberHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddScenegroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *AddScenegroupMemberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddScenegroupMemberHeaders) GetAccountContext() *AddScenegroupMemberHeadersAccountContext {
	return s.AccountContext
}

func (s *AddScenegroupMemberHeaders) SetCommonHeaders(v map[string]*string) *AddScenegroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddScenegroupMemberHeaders) SetAccountContext(v *AddScenegroupMemberHeadersAccountContext) *AddScenegroupMemberHeaders {
	s.AccountContext = v
	return s
}

func (s *AddScenegroupMemberHeaders) Validate() error {
	return dara.Validate(s)
}

type AddScenegroupMemberHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddScenegroupMemberHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddScenegroupMemberHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddScenegroupMemberHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddScenegroupMemberHeadersAccountContext) SetAccountId(v string) *AddScenegroupMemberHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddScenegroupMemberHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
