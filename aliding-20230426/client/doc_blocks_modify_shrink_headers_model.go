// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksModifyShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DocBlocksModifyShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DocBlocksModifyShrinkHeaders
	GetAccountContextShrink() *string
}

type DocBlocksModifyShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DocBlocksModifyShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksModifyShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DocBlocksModifyShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DocBlocksModifyShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DocBlocksModifyShrinkHeaders) SetCommonHeaders(v map[string]*string) *DocBlocksModifyShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DocBlocksModifyShrinkHeaders) SetAccountContextShrink(v string) *DocBlocksModifyShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DocBlocksModifyShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
