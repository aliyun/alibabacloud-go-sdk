// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceMembersShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteWorkspaceMembersShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteWorkspaceMembersShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteWorkspaceMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteWorkspaceMembersShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceMembersShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteWorkspaceMembersShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteWorkspaceMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteWorkspaceMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteWorkspaceMembersShrinkHeaders) SetAccountContextShrink(v string) *DeleteWorkspaceMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteWorkspaceMembersShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
