// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableAllSheetsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMultiDimTableAllSheetsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetMultiDimTableAllSheetsShrinkHeaders
	GetAccountContextShrink() *string
}

type GetMultiDimTableAllSheetsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetMultiDimTableAllSheetsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllSheetsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllSheetsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMultiDimTableAllSheetsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetMultiDimTableAllSheetsShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetMultiDimTableAllSheetsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMultiDimTableAllSheetsShrinkHeaders) SetAccountContextShrink(v string) *GetMultiDimTableAllSheetsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetMultiDimTableAllSheetsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
