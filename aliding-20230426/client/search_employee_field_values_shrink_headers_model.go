// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEmployeeFieldValuesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SearchEmployeeFieldValuesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SearchEmployeeFieldValuesShrinkHeaders
	GetAccountContextShrink() *string
}

type SearchEmployeeFieldValuesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SearchEmployeeFieldValuesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SearchEmployeeFieldValuesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SearchEmployeeFieldValuesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SearchEmployeeFieldValuesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SearchEmployeeFieldValuesShrinkHeaders) SetCommonHeaders(v map[string]*string) *SearchEmployeeFieldValuesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchEmployeeFieldValuesShrinkHeaders) SetAccountContextShrink(v string) *SearchEmployeeFieldValuesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SearchEmployeeFieldValuesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
