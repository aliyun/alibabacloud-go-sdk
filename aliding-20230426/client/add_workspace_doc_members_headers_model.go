// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceDocMembersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddWorkspaceDocMembersHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddWorkspaceDocMembersHeadersAccountContext) *AddWorkspaceDocMembersHeaders
	GetAccountContext() *AddWorkspaceDocMembersHeadersAccountContext
}

type AddWorkspaceDocMembersHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddWorkspaceDocMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddWorkspaceDocMembersHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceDocMembersHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddWorkspaceDocMembersHeaders) GetAccountContext() *AddWorkspaceDocMembersHeadersAccountContext {
	return s.AccountContext
}

func (s *AddWorkspaceDocMembersHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceDocMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceDocMembersHeaders) SetAccountContext(v *AddWorkspaceDocMembersHeadersAccountContext) *AddWorkspaceDocMembersHeaders {
	s.AccountContext = v
	return s
}

func (s *AddWorkspaceDocMembersHeaders) Validate() error {
	return dara.Validate(s)
}

type AddWorkspaceDocMembersHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddWorkspaceDocMembersHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceDocMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddWorkspaceDocMembersHeadersAccountContext) SetAccountId(v string) *AddWorkspaceDocMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddWorkspaceDocMembersHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
