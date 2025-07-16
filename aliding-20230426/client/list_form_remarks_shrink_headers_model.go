// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFormRemarksShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListFormRemarksShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListFormRemarksShrinkHeaders
	GetAccountContextShrink() *string
}

type ListFormRemarksShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListFormRemarksShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListFormRemarksShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListFormRemarksShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListFormRemarksShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListFormRemarksShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListFormRemarksShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListFormRemarksShrinkHeaders) SetAccountContextShrink(v string) *ListFormRemarksShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListFormRemarksShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
