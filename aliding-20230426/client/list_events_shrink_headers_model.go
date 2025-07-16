// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListEventsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListEventsShrinkHeaders
	GetAccountContextShrink() *string
}

type ListEventsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListEventsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListEventsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListEventsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListEventsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListEventsShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListEventsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEventsShrinkHeaders) SetAccountContextShrink(v string) *ListEventsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListEventsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
