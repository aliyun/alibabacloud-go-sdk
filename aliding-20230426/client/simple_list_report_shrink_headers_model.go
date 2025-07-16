// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleListReportShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SimpleListReportShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SimpleListReportShrinkHeaders
	GetAccountContextShrink() *string
}

type SimpleListReportShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SimpleListReportShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SimpleListReportShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SimpleListReportShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SimpleListReportShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SimpleListReportShrinkHeaders) SetCommonHeaders(v map[string]*string) *SimpleListReportShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SimpleListReportShrinkHeaders) SetAccountContextShrink(v string) *SimpleListReportShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SimpleListReportShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
