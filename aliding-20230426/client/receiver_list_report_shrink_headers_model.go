// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiverListReportShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ReceiverListReportShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ReceiverListReportShrinkHeaders
	GetAccountContextShrink() *string
}

type ReceiverListReportShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ReceiverListReportShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ReceiverListReportShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ReceiverListReportShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ReceiverListReportShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ReceiverListReportShrinkHeaders) SetCommonHeaders(v map[string]*string) *ReceiverListReportShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReceiverListReportShrinkHeaders) SetAccountContextShrink(v string) *ReceiverListReportShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ReceiverListReportShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
