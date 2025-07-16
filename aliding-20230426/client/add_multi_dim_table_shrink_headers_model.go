// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMultiDimTableShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddMultiDimTableShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddMultiDimTableShrinkHeaders
	GetAccountContextShrink() *string
}

type AddMultiDimTableShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddMultiDimTableShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddMultiDimTableShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddMultiDimTableShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddMultiDimTableShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddMultiDimTableShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddMultiDimTableShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddMultiDimTableShrinkHeaders) SetAccountContextShrink(v string) *AddMultiDimTableShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddMultiDimTableShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
