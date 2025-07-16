// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiDimTableFieldShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteMultiDimTableFieldShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteMultiDimTableFieldShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteMultiDimTableFieldShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteMultiDimTableFieldShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableFieldShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableFieldShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteMultiDimTableFieldShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteMultiDimTableFieldShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteMultiDimTableFieldShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMultiDimTableFieldShrinkHeaders) SetAccountContextShrink(v string) *DeleteMultiDimTableFieldShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteMultiDimTableFieldShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
