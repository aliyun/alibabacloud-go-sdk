// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeByUrlShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetNodeByUrlShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetNodeByUrlShrinkHeaders
	GetAccountContextShrink() *string
}

type GetNodeByUrlShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetNodeByUrlShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByUrlShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetNodeByUrlShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetNodeByUrlShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetNodeByUrlShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNodeByUrlShrinkHeaders) SetAccountContextShrink(v string) *GetNodeByUrlShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetNodeByUrlShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
