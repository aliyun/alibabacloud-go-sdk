// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancesByIdListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetInstancesByIdListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetInstancesByIdListHeadersAccountContext) *GetInstancesByIdListHeaders
	GetAccountContext() *GetInstancesByIdListHeadersAccountContext
}

type GetInstancesByIdListHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetInstancesByIdListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetInstancesByIdListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesByIdListHeaders) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetInstancesByIdListHeaders) GetAccountContext() *GetInstancesByIdListHeadersAccountContext {
	return s.AccountContext
}

func (s *GetInstancesByIdListHeaders) SetCommonHeaders(v map[string]*string) *GetInstancesByIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstancesByIdListHeaders) SetAccountContext(v *GetInstancesByIdListHeadersAccountContext) *GetInstancesByIdListHeaders {
	s.AccountContext = v
	return s
}

func (s *GetInstancesByIdListHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstancesByIdListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetInstancesByIdListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesByIdListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetInstancesByIdListHeadersAccountContext) SetAccountId(v string) *GetInstancesByIdListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetInstancesByIdListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
