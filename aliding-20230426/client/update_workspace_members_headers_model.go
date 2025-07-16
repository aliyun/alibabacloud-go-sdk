// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceMembersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateWorkspaceMembersHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateWorkspaceMembersHeadersAccountContext) *UpdateWorkspaceMembersHeaders
	GetAccountContext() *UpdateWorkspaceMembersHeadersAccountContext
}

type UpdateWorkspaceMembersHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateWorkspaceMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateWorkspaceMembersHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceMembersHeaders) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateWorkspaceMembersHeaders) GetAccountContext() *UpdateWorkspaceMembersHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateWorkspaceMembersHeaders) SetCommonHeaders(v map[string]*string) *UpdateWorkspaceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateWorkspaceMembersHeaders) SetAccountContext(v *UpdateWorkspaceMembersHeadersAccountContext) *UpdateWorkspaceMembersHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateWorkspaceMembersHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkspaceMembersHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateWorkspaceMembersHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateWorkspaceMembersHeadersAccountContext) SetAccountId(v string) *UpdateWorkspaceMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateWorkspaceMembersHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
