// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePermissionShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeletePermissionShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeletePermissionShrinkHeaders
	GetAccountContextShrink() *string
}

type DeletePermissionShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeletePermissionShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeletePermissionShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeletePermissionShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeletePermissionShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeletePermissionShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeletePermissionShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeletePermissionShrinkHeaders) SetAccountContextShrink(v string) *DeletePermissionShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeletePermissionShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
