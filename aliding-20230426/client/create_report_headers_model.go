// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateReportHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateReportHeadersAccountContext) *CreateReportHeaders
	GetAccountContext() *CreateReportHeadersAccountContext
}

type CreateReportHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateReportHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateReportHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateReportHeaders) GoString() string {
	return s.String()
}

func (s *CreateReportHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateReportHeaders) GetAccountContext() *CreateReportHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateReportHeaders) SetCommonHeaders(v map[string]*string) *CreateReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateReportHeaders) SetAccountContext(v *CreateReportHeadersAccountContext) *CreateReportHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateReportHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateReportHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateReportHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateReportHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateReportHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateReportHeadersAccountContext) SetAccountId(v string) *CreateReportHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateReportHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
