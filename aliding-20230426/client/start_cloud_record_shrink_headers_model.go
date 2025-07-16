// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudRecordShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StartCloudRecordShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *StartCloudRecordShrinkHeaders
	GetAccountContextShrink() *string
}

type StartCloudRecordShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s StartCloudRecordShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkHeaders) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StartCloudRecordShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *StartCloudRecordShrinkHeaders) SetCommonHeaders(v map[string]*string) *StartCloudRecordShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartCloudRecordShrinkHeaders) SetAccountContextShrink(v string) *StartCloudRecordShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *StartCloudRecordShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
