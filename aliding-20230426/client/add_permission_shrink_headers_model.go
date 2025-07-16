// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPermissionShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddPermissionShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddPermissionShrinkHeaders
	GetAccountContextShrink() *string
}

type AddPermissionShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddPermissionShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddPermissionShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddPermissionShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddPermissionShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddPermissionShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddPermissionShrinkHeaders) SetAccountContextShrink(v string) *AddPermissionShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddPermissionShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
