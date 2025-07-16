// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordTextShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryCloudRecordTextShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryCloudRecordTextShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryCloudRecordTextShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryCloudRecordTextShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordTextShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryCloudRecordTextShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryCloudRecordTextShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryCloudRecordTextShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCloudRecordTextShrinkHeaders) SetAccountContextShrink(v string) *QueryCloudRecordTextShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryCloudRecordTextShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
