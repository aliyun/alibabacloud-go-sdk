// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportTemplateByNameHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetReportTemplateByNameHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetReportTemplateByNameHeadersAccountContext) *GetReportTemplateByNameHeaders
	GetAccountContext() *GetReportTemplateByNameHeadersAccountContext
}

type GetReportTemplateByNameHeaders struct {
	CommonHeaders  map[string]*string                            `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetReportTemplateByNameHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetReportTemplateByNameHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateByNameHeaders) GoString() string {
	return s.String()
}

func (s *GetReportTemplateByNameHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetReportTemplateByNameHeaders) GetAccountContext() *GetReportTemplateByNameHeadersAccountContext {
	return s.AccountContext
}

func (s *GetReportTemplateByNameHeaders) SetCommonHeaders(v map[string]*string) *GetReportTemplateByNameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetReportTemplateByNameHeaders) SetAccountContext(v *GetReportTemplateByNameHeadersAccountContext) *GetReportTemplateByNameHeaders {
	s.AccountContext = v
	return s
}

func (s *GetReportTemplateByNameHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetReportTemplateByNameHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetReportTemplateByNameHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateByNameHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetReportTemplateByNameHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetReportTemplateByNameHeadersAccountContext) SetAccountId(v string) *GetReportTemplateByNameHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetReportTemplateByNameHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
