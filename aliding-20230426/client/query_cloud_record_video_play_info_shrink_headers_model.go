// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordVideoPlayInfoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryCloudRecordVideoPlayInfoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryCloudRecordVideoPlayInfoShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryCloudRecordVideoPlayInfoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryCloudRecordVideoPlayInfoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryCloudRecordVideoPlayInfoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryCloudRecordVideoPlayInfoShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryCloudRecordVideoPlayInfoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoShrinkHeaders) SetAccountContextShrink(v string) *QueryCloudRecordVideoPlayInfoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
