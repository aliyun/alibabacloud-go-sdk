// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceMembersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteWorkspaceMembersHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteWorkspaceMembersHeadersAccountContext) *DeleteWorkspaceMembersHeaders
	GetAccountContext() *DeleteWorkspaceMembersHeadersAccountContext
}

type DeleteWorkspaceMembersHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteWorkspaceMembersHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteWorkspaceMembersHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceMembersHeaders) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteWorkspaceMembersHeaders) GetAccountContext() *DeleteWorkspaceMembersHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteWorkspaceMembersHeaders) SetCommonHeaders(v map[string]*string) *DeleteWorkspaceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteWorkspaceMembersHeaders) SetAccountContext(v *DeleteWorkspaceMembersHeadersAccountContext) *DeleteWorkspaceMembersHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteWorkspaceMembersHeaders) Validate() error {
	return dara.Validate(s)
}

type DeleteWorkspaceMembersHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteWorkspaceMembersHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceMembersHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteWorkspaceMembersHeadersAccountContext) SetAccountId(v string) *DeleteWorkspaceMembersHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteWorkspaceMembersHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
