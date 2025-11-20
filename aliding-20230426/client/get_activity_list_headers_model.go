// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetActivityListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetActivityListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetActivityListHeadersAccountContext) *GetActivityListHeaders
	GetAccountContext() *GetActivityListHeadersAccountContext
}

type GetActivityListHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetActivityListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetActivityListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetActivityListHeaders) GoString() string {
	return s.String()
}

func (s *GetActivityListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetActivityListHeaders) GetAccountContext() *GetActivityListHeadersAccountContext {
	return s.AccountContext
}

func (s *GetActivityListHeaders) SetCommonHeaders(v map[string]*string) *GetActivityListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetActivityListHeaders) SetAccountContext(v *GetActivityListHeadersAccountContext) *GetActivityListHeaders {
	s.AccountContext = v
	return s
}

func (s *GetActivityListHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetActivityListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetActivityListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetActivityListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetActivityListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetActivityListHeadersAccountContext) SetAccountId(v string) *GetActivityListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetActivityListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
