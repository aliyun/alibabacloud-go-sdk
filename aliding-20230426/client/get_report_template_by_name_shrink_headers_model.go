// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportTemplateByNameShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetReportTemplateByNameShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetReportTemplateByNameShrinkHeaders
	GetAccountContextShrink() *string
}

type GetReportTemplateByNameShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetReportTemplateByNameShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateByNameShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetReportTemplateByNameShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetReportTemplateByNameShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetReportTemplateByNameShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetReportTemplateByNameShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetReportTemplateByNameShrinkHeaders) SetAccountContextShrink(v string) *GetReportTemplateByNameShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetReportTemplateByNameShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
