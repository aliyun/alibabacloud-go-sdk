// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordMinutesUrlShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryRecordMinutesUrlShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryRecordMinutesUrlShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryRecordMinutesUrlShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryRecordMinutesUrlShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordMinutesUrlShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryRecordMinutesUrlShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryRecordMinutesUrlShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryRecordMinutesUrlShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRecordMinutesUrlShrinkHeaders) SetAccountContextShrink(v string) *QueryRecordMinutesUrlShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryRecordMinutesUrlShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
