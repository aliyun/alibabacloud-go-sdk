// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateFormDataHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateOrUpdateFormDataHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateOrUpdateFormDataHeadersAccountContext) *CreateOrUpdateFormDataHeaders
	GetAccountContext() *CreateOrUpdateFormDataHeadersAccountContext
}

type CreateOrUpdateFormDataHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateOrUpdateFormDataHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateOrUpdateFormDataHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateFormDataHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateOrUpdateFormDataHeaders) GetAccountContext() *CreateOrUpdateFormDataHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateOrUpdateFormDataHeaders) SetCommonHeaders(v map[string]*string) *CreateOrUpdateFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrUpdateFormDataHeaders) SetAccountContext(v *CreateOrUpdateFormDataHeadersAccountContext) *CreateOrUpdateFormDataHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateOrUpdateFormDataHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateFormDataHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateOrUpdateFormDataHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateFormDataHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateOrUpdateFormDataHeadersAccountContext) SetAccountId(v string) *CreateOrUpdateFormDataHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateOrUpdateFormDataHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
