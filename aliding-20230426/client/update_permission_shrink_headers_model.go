// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePermissionShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdatePermissionShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdatePermissionShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdatePermissionShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdatePermissionShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdatePermissionShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePermissionShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdatePermissionShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdatePermissionShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdatePermissionShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePermissionShrinkHeaders) SetAccountContextShrink(v string) *UpdatePermissionShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdatePermissionShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
