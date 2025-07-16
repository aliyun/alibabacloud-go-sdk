// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksQueryShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DocBlocksQueryShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DocBlocksQueryShrinkHeaders
	GetAccountContextShrink() *string
}

type DocBlocksQueryShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DocBlocksQueryShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksQueryShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DocBlocksQueryShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DocBlocksQueryShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DocBlocksQueryShrinkHeaders) SetCommonHeaders(v map[string]*string) *DocBlocksQueryShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DocBlocksQueryShrinkHeaders) SetAccountContextShrink(v string) *DocBlocksQueryShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DocBlocksQueryShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
