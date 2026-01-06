// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkProjectionInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkProjectionInfoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetDingtalkProjectionInfoHeadersAccountContext) *GetDingtalkProjectionInfoHeaders
	GetAccountContext() *GetDingtalkProjectionInfoHeadersAccountContext
}

type GetDingtalkProjectionInfoHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetDingtalkProjectionInfoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetDingtalkProjectionInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkProjectionInfoHeaders) GetAccountContext() *GetDingtalkProjectionInfoHeadersAccountContext {
	return s.AccountContext
}

func (s *GetDingtalkProjectionInfoHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkProjectionInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkProjectionInfoHeaders) SetAccountContext(v *GetDingtalkProjectionInfoHeadersAccountContext) *GetDingtalkProjectionInfoHeaders {
	s.AccountContext = v
	return s
}

func (s *GetDingtalkProjectionInfoHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkProjectionInfoHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetDingtalkProjectionInfoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionInfoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionInfoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetDingtalkProjectionInfoHeadersAccountContext) SetAccountId(v string) *GetDingtalkProjectionInfoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetDingtalkProjectionInfoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
