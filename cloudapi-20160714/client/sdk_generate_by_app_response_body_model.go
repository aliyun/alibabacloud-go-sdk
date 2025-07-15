// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdkGenerateByAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadLink(v string) *SdkGenerateByAppResponseBody
	GetDownloadLink() *string
	SetRequestId(v string) *SdkGenerateByAppResponseBody
	GetRequestId() *string
}

type SdkGenerateByAppResponseBody struct {
	// example:
	//
	// UEsDBBQACAAIADdwnFQAAAAAAAAAAAAAAAA2AAAAQ0FTREtfSkFWQV8xMjI3NDY2NjY0MzM0MTMzXzE2NTExMjU3MD......
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// example:
	//
	// 61A16D46-EC04-5288-8A18-811B0F536CC2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SdkGenerateByAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SdkGenerateByAppResponseBody) GoString() string {
	return s.String()
}

func (s *SdkGenerateByAppResponseBody) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *SdkGenerateByAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SdkGenerateByAppResponseBody) SetDownloadLink(v string) *SdkGenerateByAppResponseBody {
	s.DownloadLink = &v
	return s
}

func (s *SdkGenerateByAppResponseBody) SetRequestId(v string) *SdkGenerateByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *SdkGenerateByAppResponseBody) Validate() error {
	return dara.Validate(s)
}
