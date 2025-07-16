// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceDocMembersShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddWorkspaceDocMembersShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddWorkspaceDocMembersShrinkHeaders
	GetAccountContextShrink() *string
}

type AddWorkspaceDocMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddWorkspaceDocMembersShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceDocMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceDocMembersShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddWorkspaceDocMembersShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddWorkspaceDocMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceDocMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceDocMembersShrinkHeaders) SetAccountContextShrink(v string) *AddWorkspaceDocMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddWorkspaceDocMembersShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
