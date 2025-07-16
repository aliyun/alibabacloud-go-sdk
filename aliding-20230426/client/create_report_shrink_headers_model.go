// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReportShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateReportShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateReportShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateReportShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateReportShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateReportShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateReportShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateReportShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateReportShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateReportShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateReportShrinkHeaders) SetAccountContextShrink(v string) *CreateReportShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateReportShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
