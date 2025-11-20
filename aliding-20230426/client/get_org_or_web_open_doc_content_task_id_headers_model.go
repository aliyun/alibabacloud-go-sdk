// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgOrWebOpenDocContentTaskIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetOrgOrWebOpenDocContentTaskIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetOrgOrWebOpenDocContentTaskIdHeadersAccountContext) *GetOrgOrWebOpenDocContentTaskIdHeaders
	GetAccountContext() *GetOrgOrWebOpenDocContentTaskIdHeadersAccountContext
}

type GetOrgOrWebOpenDocContentTaskIdHeaders struct {
	CommonHeaders  map[string]*string                                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetOrgOrWebOpenDocContentTaskIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetOrgOrWebOpenDocContentTaskIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetOrgOrWebOpenDocContentTaskIdHeaders) GoString() string {
	return s.String()
}

func (s *GetOrgOrWebOpenDocContentTaskIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetOrgOrWebOpenDocContentTaskIdHeaders) GetAccountContext() *GetOrgOrWebOpenDocContentTaskIdHeadersAccountContext {
	return s.AccountContext
}

func (s *GetOrgOrWebOpenDocContentTaskIdHeaders) SetCommonHeaders(v map[string]*string) *GetOrgOrWebOpenDocContentTaskIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdHeaders) SetAccountContext(v *GetOrgOrWebOpenDocContentTaskIdHeadersAccountContext) *GetOrgOrWebOpenDocContentTaskIdHeaders {
	s.AccountContext = v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrgOrWebOpenDocContentTaskIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetOrgOrWebOpenDocContentTaskIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetOrgOrWebOpenDocContentTaskIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetOrgOrWebOpenDocContentTaskIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetOrgOrWebOpenDocContentTaskIdHeadersAccountContext) SetAccountId(v string) *GetOrgOrWebOpenDocContentTaskIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
