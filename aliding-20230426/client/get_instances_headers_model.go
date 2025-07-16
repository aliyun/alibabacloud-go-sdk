// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetInstancesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetInstancesHeadersAccountContext) *GetInstancesHeaders
	GetAccountContext() *GetInstancesHeadersAccountContext
}

type GetInstancesHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetInstancesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetInstancesHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesHeaders) GoString() string {
	return s.String()
}

func (s *GetInstancesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetInstancesHeaders) GetAccountContext() *GetInstancesHeadersAccountContext {
	return s.AccountContext
}

func (s *GetInstancesHeaders) SetCommonHeaders(v map[string]*string) *GetInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstancesHeaders) SetAccountContext(v *GetInstancesHeadersAccountContext) *GetInstancesHeaders {
	s.AccountContext = v
	return s
}

func (s *GetInstancesHeaders) Validate() error {
	return dara.Validate(s)
}

type GetInstancesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetInstancesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetInstancesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetInstancesHeadersAccountContext) SetAccountId(v string) *GetInstancesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetInstancesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
