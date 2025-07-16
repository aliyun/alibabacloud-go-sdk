// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetNodesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetNodesShrinkHeaders
	GetAccountContextShrink() *string
}

type GetNodesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetNodesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetNodesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetNodesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetNodesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetNodesShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetNodesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNodesShrinkHeaders) SetAccountContextShrink(v string) *GetNodesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetNodesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
