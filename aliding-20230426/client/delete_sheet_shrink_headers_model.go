// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSheetShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteSheetShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteSheetShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteSheetShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteSheetShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteSheetShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSheetShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteSheetShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteSheetShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteSheetShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSheetShrinkHeaders) SetAccountContextShrink(v string) *DeleteSheetShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteSheetShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
