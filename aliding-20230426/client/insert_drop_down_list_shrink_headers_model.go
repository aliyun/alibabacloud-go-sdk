// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertDropDownListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsertDropDownListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *InsertDropDownListShrinkHeaders
	GetAccountContextShrink() *string
}

type InsertDropDownListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s InsertDropDownListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsertDropDownListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *InsertDropDownListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsertDropDownListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *InsertDropDownListShrinkHeaders) SetCommonHeaders(v map[string]*string) *InsertDropDownListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertDropDownListShrinkHeaders) SetAccountContextShrink(v string) *InsertDropDownListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *InsertDropDownListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
