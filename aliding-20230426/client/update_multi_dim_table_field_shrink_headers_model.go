// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableFieldShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateMultiDimTableFieldShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateMultiDimTableFieldShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateMultiDimTableFieldShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateMultiDimTableFieldShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableFieldShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableFieldShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateMultiDimTableFieldShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateMultiDimTableFieldShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateMultiDimTableFieldShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMultiDimTableFieldShrinkHeaders) SetAccountContextShrink(v string) *UpdateMultiDimTableFieldShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateMultiDimTableFieldShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
