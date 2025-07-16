// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSheetShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSheetShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetSheetShrinkHeaders
	GetAccountContextShrink() *string
}

type GetSheetShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetSheetShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSheetShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetSheetShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSheetShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetSheetShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetSheetShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSheetShrinkHeaders) SetAccountContextShrink(v string) *GetSheetShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetSheetShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
