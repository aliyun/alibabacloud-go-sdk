// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportUnReadCountShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetReportUnReadCountShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetReportUnReadCountShrinkHeaders
	GetAccountContextShrink() *string
}

type GetReportUnReadCountShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetReportUnReadCountShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetReportUnReadCountShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetReportUnReadCountShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetReportUnReadCountShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetReportUnReadCountShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetReportUnReadCountShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetReportUnReadCountShrinkHeaders) SetAccountContextShrink(v string) *GetReportUnReadCountShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetReportUnReadCountShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
