// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePermissionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdatePermissionHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdatePermissionHeadersAccountContext) *UpdatePermissionHeaders
	GetAccountContext() *UpdatePermissionHeadersAccountContext
}

type UpdatePermissionHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdatePermissionHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdatePermissionHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdatePermissionHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePermissionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdatePermissionHeaders) GetAccountContext() *UpdatePermissionHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdatePermissionHeaders) SetCommonHeaders(v map[string]*string) *UpdatePermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePermissionHeaders) SetAccountContext(v *UpdatePermissionHeadersAccountContext) *UpdatePermissionHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdatePermissionHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdatePermissionHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdatePermissionHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdatePermissionHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdatePermissionHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdatePermissionHeadersAccountContext) SetAccountId(v string) *UpdatePermissionHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdatePermissionHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
