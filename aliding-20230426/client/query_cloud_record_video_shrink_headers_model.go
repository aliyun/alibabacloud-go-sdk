// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordVideoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryCloudRecordVideoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryCloudRecordVideoShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryCloudRecordVideoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryCloudRecordVideoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryCloudRecordVideoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryCloudRecordVideoShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryCloudRecordVideoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCloudRecordVideoShrinkHeaders) SetAccountContextShrink(v string) *QueryCloudRecordVideoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryCloudRecordVideoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
