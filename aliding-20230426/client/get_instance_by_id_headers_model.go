// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceByIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetInstanceByIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetInstanceByIdHeadersAccountContext) *GetInstanceByIdHeaders
	GetAccountContext() *GetInstanceByIdHeadersAccountContext
}

type GetInstanceByIdHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetInstanceByIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetInstanceByIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceByIdHeaders) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetInstanceByIdHeaders) GetAccountContext() *GetInstanceByIdHeadersAccountContext {
	return s.AccountContext
}

func (s *GetInstanceByIdHeaders) SetCommonHeaders(v map[string]*string) *GetInstanceByIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstanceByIdHeaders) SetAccountContext(v *GetInstanceByIdHeadersAccountContext) *GetInstanceByIdHeaders {
	s.AccountContext = v
	return s
}

func (s *GetInstanceByIdHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceByIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetInstanceByIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceByIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetInstanceByIdHeadersAccountContext) SetAccountId(v string) *GetInstanceByIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetInstanceByIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
