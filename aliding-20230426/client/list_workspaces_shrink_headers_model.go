// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListWorkspacesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListWorkspacesShrinkHeaders
	GetAccountContextShrink() *string
}

type ListWorkspacesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListWorkspacesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListWorkspacesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListWorkspacesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListWorkspacesShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListWorkspacesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListWorkspacesShrinkHeaders) SetAccountContextShrink(v string) *ListWorkspacesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListWorkspacesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
