// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiDimTableFieldShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateMultiDimTableFieldShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateMultiDimTableFieldShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateMultiDimTableFieldShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateMultiDimTableFieldShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiDimTableFieldShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateMultiDimTableFieldShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateMultiDimTableFieldShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateMultiDimTableFieldShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateMultiDimTableFieldShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMultiDimTableFieldShrinkHeaders) SetAccountContextShrink(v string) *CreateMultiDimTableFieldShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateMultiDimTableFieldShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
