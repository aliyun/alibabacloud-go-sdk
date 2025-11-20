// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNewestInnerGroupsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetNewestInnerGroupsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetNewestInnerGroupsHeadersAccountContext) *GetNewestInnerGroupsHeaders
	GetAccountContext() *GetNewestInnerGroupsHeadersAccountContext
}

type GetNewestInnerGroupsHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetNewestInnerGroupsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetNewestInnerGroupsHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetNewestInnerGroupsHeaders) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetNewestInnerGroupsHeaders) GetAccountContext() *GetNewestInnerGroupsHeadersAccountContext {
	return s.AccountContext
}

func (s *GetNewestInnerGroupsHeaders) SetCommonHeaders(v map[string]*string) *GetNewestInnerGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNewestInnerGroupsHeaders) SetAccountContext(v *GetNewestInnerGroupsHeadersAccountContext) *GetNewestInnerGroupsHeaders {
	s.AccountContext = v
	return s
}

func (s *GetNewestInnerGroupsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNewestInnerGroupsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetNewestInnerGroupsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetNewestInnerGroupsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetNewestInnerGroupsHeadersAccountContext) SetAccountId(v string) *GetNewestInnerGroupsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetNewestInnerGroupsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
