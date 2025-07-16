// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByOrgIdAndStaffIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserIdByOrgIdAndStaffIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetUserIdByOrgIdAndStaffIdHeadersAccountContext) *GetUserIdByOrgIdAndStaffIdHeaders
	GetAccountContext() *GetUserIdByOrgIdAndStaffIdHeadersAccountContext
}

type GetUserIdByOrgIdAndStaffIdHeaders struct {
	CommonHeaders  map[string]*string                               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetUserIdByOrgIdAndStaffIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetUserIdByOrgIdAndStaffIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOrgIdAndStaffIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByOrgIdAndStaffIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserIdByOrgIdAndStaffIdHeaders) GetAccountContext() *GetUserIdByOrgIdAndStaffIdHeadersAccountContext {
	return s.AccountContext
}

func (s *GetUserIdByOrgIdAndStaffIdHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByOrgIdAndStaffIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdHeaders) SetAccountContext(v *GetUserIdByOrgIdAndStaffIdHeadersAccountContext) *GetUserIdByOrgIdAndStaffIdHeaders {
	s.AccountContext = v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdHeaders) Validate() error {
	return dara.Validate(s)
}

type GetUserIdByOrgIdAndStaffIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetUserIdByOrgIdAndStaffIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOrgIdAndStaffIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetUserIdByOrgIdAndStaffIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetUserIdByOrgIdAndStaffIdHeadersAccountContext) SetAccountId(v string) *GetUserIdByOrgIdAndStaffIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
