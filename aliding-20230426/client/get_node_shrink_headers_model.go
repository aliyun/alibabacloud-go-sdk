// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetNodeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetNodeShrinkHeaders
	GetAccountContextShrink() *string
}

type GetNodeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetNodeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetNodeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetNodeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetNodeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetNodeShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetNodeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNodeShrinkHeaders) SetAccountContextShrink(v string) *GetNodeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetNodeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
