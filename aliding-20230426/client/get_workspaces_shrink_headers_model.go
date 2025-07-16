// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspacesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetWorkspacesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetWorkspacesShrinkHeaders
	GetAccountContextShrink() *string
}

type GetWorkspacesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetWorkspacesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetWorkspacesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetWorkspacesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetWorkspacesShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetWorkspacesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWorkspacesShrinkHeaders) SetAccountContextShrink(v string) *GetWorkspacesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetWorkspacesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
