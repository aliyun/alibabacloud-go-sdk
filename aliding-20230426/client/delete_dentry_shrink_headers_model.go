// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDentryShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteDentryShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteDentryShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteDentryShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteDentryShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteDentryShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDentryShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteDentryShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteDentryShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteDentryShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDentryShrinkHeaders) SetAccountContextShrink(v string) *DeleteDentryShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteDentryShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
