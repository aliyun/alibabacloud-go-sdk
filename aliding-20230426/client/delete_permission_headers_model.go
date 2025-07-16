// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePermissionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeletePermissionHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeletePermissionHeadersAccountContext) *DeletePermissionHeaders
	GetAccountContext() *DeletePermissionHeadersAccountContext
}

type DeletePermissionHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeletePermissionHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeletePermissionHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeletePermissionHeaders) GoString() string {
	return s.String()
}

func (s *DeletePermissionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeletePermissionHeaders) GetAccountContext() *DeletePermissionHeadersAccountContext {
	return s.AccountContext
}

func (s *DeletePermissionHeaders) SetCommonHeaders(v map[string]*string) *DeletePermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeletePermissionHeaders) SetAccountContext(v *DeletePermissionHeadersAccountContext) *DeletePermissionHeaders {
	s.AccountContext = v
	return s
}

func (s *DeletePermissionHeaders) Validate() error {
	return dara.Validate(s)
}

type DeletePermissionHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeletePermissionHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeletePermissionHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeletePermissionHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeletePermissionHeadersAccountContext) SetAccountId(v string) *DeletePermissionHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeletePermissionHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
