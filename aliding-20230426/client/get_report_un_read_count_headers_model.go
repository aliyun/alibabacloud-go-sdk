// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportUnReadCountHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetReportUnReadCountHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetReportUnReadCountHeadersAccountContext) *GetReportUnReadCountHeaders
	GetAccountContext() *GetReportUnReadCountHeadersAccountContext
}

type GetReportUnReadCountHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetReportUnReadCountHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetReportUnReadCountHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetReportUnReadCountHeaders) GoString() string {
	return s.String()
}

func (s *GetReportUnReadCountHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetReportUnReadCountHeaders) GetAccountContext() *GetReportUnReadCountHeadersAccountContext {
	return s.AccountContext
}

func (s *GetReportUnReadCountHeaders) SetCommonHeaders(v map[string]*string) *GetReportUnReadCountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetReportUnReadCountHeaders) SetAccountContext(v *GetReportUnReadCountHeadersAccountContext) *GetReportUnReadCountHeaders {
	s.AccountContext = v
	return s
}

func (s *GetReportUnReadCountHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetReportUnReadCountHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetReportUnReadCountHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetReportUnReadCountHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetReportUnReadCountHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetReportUnReadCountHeadersAccountContext) SetAccountId(v string) *GetReportUnReadCountHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetReportUnReadCountHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
