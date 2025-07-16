// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordPermissionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddRecordPermissionHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddRecordPermissionHeadersAccountContext) *AddRecordPermissionHeaders
	GetAccountContext() *AddRecordPermissionHeadersAccountContext
}

type AddRecordPermissionHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddRecordPermissionHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddRecordPermissionHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddRecordPermissionHeaders) GoString() string {
	return s.String()
}

func (s *AddRecordPermissionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddRecordPermissionHeaders) GetAccountContext() *AddRecordPermissionHeadersAccountContext {
	return s.AccountContext
}

func (s *AddRecordPermissionHeaders) SetCommonHeaders(v map[string]*string) *AddRecordPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddRecordPermissionHeaders) SetAccountContext(v *AddRecordPermissionHeadersAccountContext) *AddRecordPermissionHeaders {
	s.AccountContext = v
	return s
}

func (s *AddRecordPermissionHeaders) Validate() error {
	return dara.Validate(s)
}

type AddRecordPermissionHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddRecordPermissionHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddRecordPermissionHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddRecordPermissionHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddRecordPermissionHeadersAccountContext) SetAccountId(v string) *AddRecordPermissionHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddRecordPermissionHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
