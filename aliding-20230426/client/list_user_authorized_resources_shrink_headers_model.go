// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAuthorizedResourcesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListUserAuthorizedResourcesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListUserAuthorizedResourcesShrinkHeaders
	GetAccountContextShrink() *string
}

type ListUserAuthorizedResourcesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListUserAuthorizedResourcesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthorizedResourcesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListUserAuthorizedResourcesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListUserAuthorizedResourcesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListUserAuthorizedResourcesShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListUserAuthorizedResourcesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUserAuthorizedResourcesShrinkHeaders) SetAccountContextShrink(v string) *ListUserAuthorizedResourcesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListUserAuthorizedResourcesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
