// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIdListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetInstanceIdListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetInstanceIdListHeadersAccountContext) *GetInstanceIdListHeaders
	GetAccountContext() *GetInstanceIdListHeadersAccountContext
}

type GetInstanceIdListHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetInstanceIdListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetInstanceIdListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIdListHeaders) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetInstanceIdListHeaders) GetAccountContext() *GetInstanceIdListHeadersAccountContext {
	return s.AccountContext
}

func (s *GetInstanceIdListHeaders) SetCommonHeaders(v map[string]*string) *GetInstanceIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstanceIdListHeaders) SetAccountContext(v *GetInstanceIdListHeadersAccountContext) *GetInstanceIdListHeaders {
	s.AccountContext = v
	return s
}

func (s *GetInstanceIdListHeaders) Validate() error {
	return dara.Validate(s)
}

type GetInstanceIdListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetInstanceIdListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIdListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetInstanceIdListHeadersAccountContext) SetAccountId(v string) *GetInstanceIdListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetInstanceIdListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
