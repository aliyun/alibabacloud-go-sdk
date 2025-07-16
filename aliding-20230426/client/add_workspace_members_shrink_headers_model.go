// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceMembersShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddWorkspaceMembersShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddWorkspaceMembersShrinkHeaders
	GetAccountContextShrink() *string
}

type AddWorkspaceMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddWorkspaceMembersShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceMembersShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddWorkspaceMembersShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddWorkspaceMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceMembersShrinkHeaders) SetAccountContextShrink(v string) *AddWorkspaceMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddWorkspaceMembersShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
