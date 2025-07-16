// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetWorkspaceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetWorkspaceShrinkHeaders
	GetAccountContextShrink() *string
}

type GetWorkspaceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetWorkspaceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetWorkspaceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetWorkspaceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetWorkspaceShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetWorkspaceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWorkspaceShrinkHeaders) SetAccountContextShrink(v string) *GetWorkspaceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetWorkspaceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
