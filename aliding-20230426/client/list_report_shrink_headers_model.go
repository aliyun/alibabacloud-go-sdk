// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListReportShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListReportShrinkHeaders
	GetAccountContextShrink() *string
}

type ListReportShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListReportShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListReportShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListReportShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListReportShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListReportShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListReportShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListReportShrinkHeaders) SetAccountContextShrink(v string) *ListReportShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListReportShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
