// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReportDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryReportDetailHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryReportDetailHeadersAccountContext) *QueryReportDetailHeaders
	GetAccountContext() *QueryReportDetailHeadersAccountContext
}

type QueryReportDetailHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryReportDetailHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryReportDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryReportDetailHeaders) GoString() string {
	return s.String()
}

func (s *QueryReportDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryReportDetailHeaders) GetAccountContext() *QueryReportDetailHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryReportDetailHeaders) SetCommonHeaders(v map[string]*string) *QueryReportDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryReportDetailHeaders) SetAccountContext(v *QueryReportDetailHeadersAccountContext) *QueryReportDetailHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryReportDetailHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryReportDetailHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryReportDetailHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryReportDetailHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryReportDetailHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryReportDetailHeadersAccountContext) SetAccountId(v string) *QueryReportDetailHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryReportDetailHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
