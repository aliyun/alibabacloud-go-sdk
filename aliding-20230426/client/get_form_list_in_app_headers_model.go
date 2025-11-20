// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormListInAppHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetFormListInAppHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetFormListInAppHeadersAccountContext) *GetFormListInAppHeaders
	GetAccountContext() *GetFormListInAppHeadersAccountContext
}

type GetFormListInAppHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetFormListInAppHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetFormListInAppHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetFormListInAppHeaders) GoString() string {
	return s.String()
}

func (s *GetFormListInAppHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetFormListInAppHeaders) GetAccountContext() *GetFormListInAppHeadersAccountContext {
	return s.AccountContext
}

func (s *GetFormListInAppHeaders) SetCommonHeaders(v map[string]*string) *GetFormListInAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormListInAppHeaders) SetAccountContext(v *GetFormListInAppHeadersAccountContext) *GetFormListInAppHeaders {
	s.AccountContext = v
	return s
}

func (s *GetFormListInAppHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFormListInAppHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetFormListInAppHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetFormListInAppHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetFormListInAppHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetFormListInAppHeadersAccountContext) SetAccountId(v string) *GetFormListInAppHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetFormListInAppHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
