// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceDocMembersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteWorkspaceDocMembersHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteWorkspaceDocMembersHeadersAccountContext) *DeleteWorkspaceDocMembersHeaders
	GetAccountContext() *DeleteWorkspaceDocMembersHeadersAccountContext
}

type DeleteWorkspaceDocMembersHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteWorkspaceDocMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteWorkspaceDocMembersHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceDocMembersHeaders) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteWorkspaceDocMembersHeaders) GetAccountContext() *DeleteWorkspaceDocMembersHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteWorkspaceDocMembersHeaders) SetCommonHeaders(v map[string]*string) *DeleteWorkspaceDocMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteWorkspaceDocMembersHeaders) SetAccountContext(v *DeleteWorkspaceDocMembersHeadersAccountContext) *DeleteWorkspaceDocMembersHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteWorkspaceDocMembersHeaders) Validate() error {
	return dara.Validate(s)
}

type DeleteWorkspaceDocMembersHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteWorkspaceDocMembersHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceDocMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteWorkspaceDocMembersHeadersAccountContext) SetAccountId(v string) *DeleteWorkspaceDocMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteWorkspaceDocMembersHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
