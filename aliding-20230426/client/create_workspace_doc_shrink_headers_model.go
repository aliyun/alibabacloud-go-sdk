// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceDocShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateWorkspaceDocShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateWorkspaceDocShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateWorkspaceDocShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateWorkspaceDocShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceDocShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceDocShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateWorkspaceDocShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateWorkspaceDocShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateWorkspaceDocShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateWorkspaceDocShrinkHeaders) SetAccountContextShrink(v string) *CreateWorkspaceDocShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateWorkspaceDocShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
