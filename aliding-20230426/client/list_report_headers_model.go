// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListReportHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListReportHeadersAccountContext) *ListReportHeaders
	GetAccountContext() *ListReportHeadersAccountContext
}

type ListReportHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListReportHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListReportHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListReportHeaders) GoString() string {
	return s.String()
}

func (s *ListReportHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListReportHeaders) GetAccountContext() *ListReportHeadersAccountContext {
	return s.AccountContext
}

func (s *ListReportHeaders) SetCommonHeaders(v map[string]*string) *ListReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListReportHeaders) SetAccountContext(v *ListReportHeadersAccountContext) *ListReportHeaders {
	s.AccountContext = v
	return s
}

func (s *ListReportHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListReportHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListReportHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListReportHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListReportHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListReportHeadersAccountContext) SetAccountId(v string) *ListReportHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListReportHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
