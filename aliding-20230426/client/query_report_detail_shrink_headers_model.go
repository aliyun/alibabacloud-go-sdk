// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReportDetailShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryReportDetailShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryReportDetailShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryReportDetailShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryReportDetailShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryReportDetailShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryReportDetailShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryReportDetailShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryReportDetailShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryReportDetailShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryReportDetailShrinkHeaders) SetAccountContextShrink(v string) *QueryReportDetailShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryReportDetailShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
