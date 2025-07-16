// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPermissionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddPermissionHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddPermissionHeadersAccountContext) *AddPermissionHeaders
	GetAccountContext() *AddPermissionHeadersAccountContext
}

type AddPermissionHeaders struct {
	CommonHeaders  map[string]*string                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddPermissionHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddPermissionHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionHeaders) GoString() string {
	return s.String()
}

func (s *AddPermissionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddPermissionHeaders) GetAccountContext() *AddPermissionHeadersAccountContext {
	return s.AccountContext
}

func (s *AddPermissionHeaders) SetCommonHeaders(v map[string]*string) *AddPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddPermissionHeaders) SetAccountContext(v *AddPermissionHeadersAccountContext) *AddPermissionHeaders {
	s.AccountContext = v
	return s
}

func (s *AddPermissionHeaders) Validate() error {
	return dara.Validate(s)
}

type AddPermissionHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddPermissionHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddPermissionHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddPermissionHeadersAccountContext) SetAccountId(v string) *AddPermissionHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddPermissionHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
