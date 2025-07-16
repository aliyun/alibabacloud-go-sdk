// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRowsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteRowsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteRowsShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteRowsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteRowsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteRowsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteRowsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteRowsShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteRowsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteRowsShrinkHeaders) SetAccountContextShrink(v string) *DeleteRowsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteRowsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
