// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceDocMembersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateWorkspaceDocMembersHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateWorkspaceDocMembersHeadersAccountContext) *UpdateWorkspaceDocMembersHeaders
	GetAccountContext() *UpdateWorkspaceDocMembersHeadersAccountContext
}

type UpdateWorkspaceDocMembersHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateWorkspaceDocMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateWorkspaceDocMembersHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceDocMembersHeaders) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateWorkspaceDocMembersHeaders) GetAccountContext() *UpdateWorkspaceDocMembersHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateWorkspaceDocMembersHeaders) SetCommonHeaders(v map[string]*string) *UpdateWorkspaceDocMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateWorkspaceDocMembersHeaders) SetAccountContext(v *UpdateWorkspaceDocMembersHeadersAccountContext) *UpdateWorkspaceDocMembersHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateWorkspaceDocMembersHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkspaceDocMembersHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateWorkspaceDocMembersHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceDocMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateWorkspaceDocMembersHeadersAccountContext) SetAccountId(v string) *UpdateWorkspaceDocMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateWorkspaceDocMembersHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
