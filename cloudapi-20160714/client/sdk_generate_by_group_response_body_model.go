// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdkGenerateByGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadLink(v string) *SdkGenerateByGroupResponseBody
	GetDownloadLink() *string
	SetRequestId(v string) *SdkGenerateByGroupResponseBody
	GetRequestId() *string
}

type SdkGenerateByGroupResponseBody struct {
	// example:
	//
	// http://oss-bucket/object
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// example:
	//
	// D0075BDA-8AED-5073-A70A-FE44E86AB20F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SdkGenerateByGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SdkGenerateByGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SdkGenerateByGroupResponseBody) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *SdkGenerateByGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SdkGenerateByGroupResponseBody) SetDownloadLink(v string) *SdkGenerateByGroupResponseBody {
	s.DownloadLink = &v
	return s
}

func (s *SdkGenerateByGroupResponseBody) SetRequestId(v string) *SdkGenerateByGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SdkGenerateByGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
