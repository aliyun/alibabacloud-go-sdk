// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearDataHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ClearDataHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ClearDataHeadersAccountContext) *ClearDataHeaders
	GetAccountContext() *ClearDataHeadersAccountContext
}

type ClearDataHeaders struct {
	CommonHeaders  map[string]*string              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ClearDataHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ClearDataHeaders) String() string {
	return dara.Prettify(s)
}

func (s ClearDataHeaders) GoString() string {
	return s.String()
}

func (s *ClearDataHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ClearDataHeaders) GetAccountContext() *ClearDataHeadersAccountContext {
	return s.AccountContext
}

func (s *ClearDataHeaders) SetCommonHeaders(v map[string]*string) *ClearDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ClearDataHeaders) SetAccountContext(v *ClearDataHeadersAccountContext) *ClearDataHeaders {
	s.AccountContext = v
	return s
}

func (s *ClearDataHeaders) Validate() error {
	return dara.Validate(s)
}

type ClearDataHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ClearDataHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ClearDataHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ClearDataHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ClearDataHeadersAccountContext) SetAccountId(v string) *ClearDataHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ClearDataHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
