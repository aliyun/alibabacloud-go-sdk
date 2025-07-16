// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateMultiDimTableShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateMultiDimTableShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateMultiDimTableShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateMultiDimTableShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateMultiDimTableShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateMultiDimTableShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateMultiDimTableShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMultiDimTableShrinkHeaders) SetAccountContextShrink(v string) *UpdateMultiDimTableShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateMultiDimTableShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
