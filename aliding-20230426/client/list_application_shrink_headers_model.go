// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListApplicationShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListApplicationShrinkHeaders
	GetAccountContextShrink() *string
}

type ListApplicationShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListApplicationShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListApplicationShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListApplicationShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListApplicationShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListApplicationShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListApplicationShrinkHeaders) SetAccountContextShrink(v string) *ListApplicationShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListApplicationShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
