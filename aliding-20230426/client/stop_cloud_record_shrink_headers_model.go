// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCloudRecordShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StopCloudRecordShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *StopCloudRecordShrinkHeaders
	GetAccountContextShrink() *string
}

type StopCloudRecordShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s StopCloudRecordShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s StopCloudRecordShrinkHeaders) GoString() string {
	return s.String()
}

func (s *StopCloudRecordShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StopCloudRecordShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *StopCloudRecordShrinkHeaders) SetCommonHeaders(v map[string]*string) *StopCloudRecordShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopCloudRecordShrinkHeaders) SetAccountContextShrink(v string) *StopCloudRecordShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *StopCloudRecordShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
