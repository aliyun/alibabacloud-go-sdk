// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllSheetsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAllSheetsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetAllSheetsShrinkHeaders
	GetAccountContextShrink() *string
}

type GetAllSheetsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetAllSheetsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAllSheetsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetAllSheetsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAllSheetsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetAllSheetsShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetAllSheetsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAllSheetsShrinkHeaders) SetAccountContextShrink(v string) *GetAllSheetsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetAllSheetsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
