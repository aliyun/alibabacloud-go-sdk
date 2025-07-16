// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceMembersShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateWorkspaceMembersShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateWorkspaceMembersShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateWorkspaceMembersShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateWorkspaceMembersShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceMembersShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceMembersShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateWorkspaceMembersShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateWorkspaceMembersShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateWorkspaceMembersShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateWorkspaceMembersShrinkHeaders) SetAccountContextShrink(v string) *UpdateWorkspaceMembersShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateWorkspaceMembersShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
