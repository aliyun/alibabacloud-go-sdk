// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRelatedWorkspacesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetRelatedWorkspacesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetRelatedWorkspacesShrinkHeaders
	GetAccountContextShrink() *string
}

type GetRelatedWorkspacesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetRelatedWorkspacesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetRelatedWorkspacesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetRelatedWorkspacesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetRelatedWorkspacesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetRelatedWorkspacesShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetRelatedWorkspacesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRelatedWorkspacesShrinkHeaders) SetAccountContextShrink(v string) *GetRelatedWorkspacesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetRelatedWorkspacesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
