// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListPermissionsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListPermissionsShrinkHeaders
	GetAccountContextShrink() *string
}

type ListPermissionsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListPermissionsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListPermissionsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListPermissionsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListPermissionsShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListPermissionsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListPermissionsShrinkHeaders) SetAccountContextShrink(v string) *ListPermissionsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListPermissionsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
