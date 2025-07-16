// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMineWorkspaceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMineWorkspaceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetMineWorkspaceShrinkHeaders
	GetAccountContextShrink() *string
}

type GetMineWorkspaceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetMineWorkspaceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMineWorkspaceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMineWorkspaceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetMineWorkspaceShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetMineWorkspaceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMineWorkspaceShrinkHeaders) SetAccountContextShrink(v string) *GetMineWorkspaceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetMineWorkspaceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
