// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksDeleteShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DocBlocksDeleteShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DocBlocksDeleteShrinkHeaders
	GetAccountContextShrink() *string
}

type DocBlocksDeleteShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DocBlocksDeleteShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksDeleteShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DocBlocksDeleteShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DocBlocksDeleteShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DocBlocksDeleteShrinkHeaders) SetCommonHeaders(v map[string]*string) *DocBlocksDeleteShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DocBlocksDeleteShrinkHeaders) SetAccountContextShrink(v string) *DocBlocksDeleteShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DocBlocksDeleteShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
