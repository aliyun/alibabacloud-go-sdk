// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAlibabaStaffHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CheckAlibabaStaffHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CheckAlibabaStaffHeadersAccountContext) *CheckAlibabaStaffHeaders
	GetAccountContext() *CheckAlibabaStaffHeadersAccountContext
}

type CheckAlibabaStaffHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CheckAlibabaStaffHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CheckAlibabaStaffHeaders) String() string {
	return dara.Prettify(s)
}

func (s CheckAlibabaStaffHeaders) GoString() string {
	return s.String()
}

func (s *CheckAlibabaStaffHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CheckAlibabaStaffHeaders) GetAccountContext() *CheckAlibabaStaffHeadersAccountContext {
	return s.AccountContext
}

func (s *CheckAlibabaStaffHeaders) SetCommonHeaders(v map[string]*string) *CheckAlibabaStaffHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckAlibabaStaffHeaders) SetAccountContext(v *CheckAlibabaStaffHeadersAccountContext) *CheckAlibabaStaffHeaders {
	s.AccountContext = v
	return s
}

func (s *CheckAlibabaStaffHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckAlibabaStaffHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CheckAlibabaStaffHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CheckAlibabaStaffHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CheckAlibabaStaffHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CheckAlibabaStaffHeadersAccountContext) SetAccountId(v string) *CheckAlibabaStaffHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CheckAlibabaStaffHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
