// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableSheetShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMultiDimTableSheetShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetMultiDimTableSheetShrinkHeaders
	GetAccountContextShrink() *string
}

type GetMultiDimTableSheetShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetMultiDimTableSheetShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableSheetShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableSheetShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMultiDimTableSheetShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetMultiDimTableSheetShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetMultiDimTableSheetShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMultiDimTableSheetShrinkHeaders) SetAccountContextShrink(v string) *GetMultiDimTableSheetShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetMultiDimTableSheetShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
