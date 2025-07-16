// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceDocMembersShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateWorkspaceDocMembersShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateWorkspaceDocMembersShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateWorkspaceDocMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateWorkspaceDocMembersShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceDocMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDocMembersShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateWorkspaceDocMembersShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateWorkspaceDocMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateWorkspaceDocMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateWorkspaceDocMembersShrinkHeaders) SetAccountContextShrink(v string) *UpdateWorkspaceDocMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateWorkspaceDocMembersShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
