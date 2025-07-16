// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddWorkspaceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddWorkspaceShrinkHeaders
	GetAccountContextShrink() *string
}

type AddWorkspaceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddWorkspaceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddWorkspaceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddWorkspaceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddWorkspaceShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddWorkspaceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddWorkspaceShrinkHeaders) SetAccountContextShrink(v string) *AddWorkspaceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddWorkspaceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
