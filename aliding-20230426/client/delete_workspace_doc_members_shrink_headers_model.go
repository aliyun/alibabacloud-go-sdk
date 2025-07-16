// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceDocMembersShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteWorkspaceDocMembersShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteWorkspaceDocMembersShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteWorkspaceDocMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteWorkspaceDocMembersShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceDocMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceDocMembersShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteWorkspaceDocMembersShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteWorkspaceDocMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteWorkspaceDocMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteWorkspaceDocMembersShrinkHeaders) SetAccountContextShrink(v string) *DeleteWorkspaceDocMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteWorkspaceDocMembersShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
