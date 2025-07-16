// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFieldDefByUuidHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetFieldDefByUuidHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetFieldDefByUuidHeadersAccountContext) *GetFieldDefByUuidHeaders
	GetAccountContext() *GetFieldDefByUuidHeadersAccountContext
}

type GetFieldDefByUuidHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetFieldDefByUuidHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetFieldDefByUuidHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetFieldDefByUuidHeaders) GoString() string {
	return s.String()
}

func (s *GetFieldDefByUuidHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetFieldDefByUuidHeaders) GetAccountContext() *GetFieldDefByUuidHeadersAccountContext {
	return s.AccountContext
}

func (s *GetFieldDefByUuidHeaders) SetCommonHeaders(v map[string]*string) *GetFieldDefByUuidHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFieldDefByUuidHeaders) SetAccountContext(v *GetFieldDefByUuidHeadersAccountContext) *GetFieldDefByUuidHeaders {
	s.AccountContext = v
	return s
}

func (s *GetFieldDefByUuidHeaders) Validate() error {
	return dara.Validate(s)
}

type GetFieldDefByUuidHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetFieldDefByUuidHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetFieldDefByUuidHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetFieldDefByUuidHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetFieldDefByUuidHeadersAccountContext) SetAccountId(v string) *GetFieldDefByUuidHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetFieldDefByUuidHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
