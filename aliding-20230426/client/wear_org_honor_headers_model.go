// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWearOrgHonorHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *WearOrgHonorHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *WearOrgHonorHeadersAccountContext) *WearOrgHonorHeaders
	GetAccountContext() *WearOrgHonorHeadersAccountContext
}

type WearOrgHonorHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *WearOrgHonorHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s WearOrgHonorHeaders) String() string {
	return dara.Prettify(s)
}

func (s WearOrgHonorHeaders) GoString() string {
	return s.String()
}

func (s *WearOrgHonorHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *WearOrgHonorHeaders) GetAccountContext() *WearOrgHonorHeadersAccountContext {
	return s.AccountContext
}

func (s *WearOrgHonorHeaders) SetCommonHeaders(v map[string]*string) *WearOrgHonorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WearOrgHonorHeaders) SetAccountContext(v *WearOrgHonorHeadersAccountContext) *WearOrgHonorHeaders {
	s.AccountContext = v
	return s
}

func (s *WearOrgHonorHeaders) Validate() error {
	return dara.Validate(s)
}

type WearOrgHonorHeadersAccountContext struct {
	// example:
	//
	// 1915607600538524
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s WearOrgHonorHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s WearOrgHonorHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *WearOrgHonorHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *WearOrgHonorHeadersAccountContext) SetAccountId(v string) *WearOrgHonorHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *WearOrgHonorHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
