// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateWorkspaceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateWorkspaceShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateWorkspaceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateWorkspaceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateWorkspaceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateWorkspaceShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateWorkspaceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateWorkspaceShrinkHeaders) SetAccountContextShrink(v string) *CreateWorkspaceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateWorkspaceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
