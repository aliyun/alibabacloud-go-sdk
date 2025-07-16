// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserLatestPlanHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserLatestPlanHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetUserLatestPlanHeadersAccountContext) *GetUserLatestPlanHeaders
	GetAccountContext() *GetUserLatestPlanHeadersAccountContext
}

type GetUserLatestPlanHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetUserLatestPlanHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetUserLatestPlanHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserLatestPlanHeaders) GoString() string {
	return s.String()
}

func (s *GetUserLatestPlanHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserLatestPlanHeaders) GetAccountContext() *GetUserLatestPlanHeadersAccountContext {
	return s.AccountContext
}

func (s *GetUserLatestPlanHeaders) SetCommonHeaders(v map[string]*string) *GetUserLatestPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserLatestPlanHeaders) SetAccountContext(v *GetUserLatestPlanHeadersAccountContext) *GetUserLatestPlanHeaders {
	s.AccountContext = v
	return s
}

func (s *GetUserLatestPlanHeaders) Validate() error {
	return dara.Validate(s)
}

type GetUserLatestPlanHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetUserLatestPlanHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetUserLatestPlanHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetUserLatestPlanHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetUserLatestPlanHeadersAccountContext) SetAccountId(v string) *GetUserLatestPlanHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetUserLatestPlanHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
