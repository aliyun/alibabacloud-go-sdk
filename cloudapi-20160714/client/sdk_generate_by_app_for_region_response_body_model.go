// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdkGenerateByAppForRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadLink(v string) *SdkGenerateByAppForRegionResponseBody
	GetDownloadLink() *string
	SetRequestId(v string) *SdkGenerateByAppForRegionResponseBody
	GetRequestId() *string
}

type SdkGenerateByAppForRegionResponseBody struct {
	// The code of the SDK by using the Base64 scheme. You can obtain the file by using the Base64 decoding scheme.
	//
	// example:
	//
	// UEsDBBQACAAIADdwnFQAAAAAAAAAAAAAAAA2AAAAQ0FTREtfSkFWQV8xMjI3NDY2NjY0MzM0MTMzXzE2NTExMjU3MD......
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CE5722A6-AE78-4741-A9B0-6C817D360510
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SdkGenerateByAppForRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SdkGenerateByAppForRegionResponseBody) GoString() string {
	return s.String()
}

func (s *SdkGenerateByAppForRegionResponseBody) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *SdkGenerateByAppForRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SdkGenerateByAppForRegionResponseBody) SetDownloadLink(v string) *SdkGenerateByAppForRegionResponseBody {
	s.DownloadLink = &v
	return s
}

func (s *SdkGenerateByAppForRegionResponseBody) SetRequestId(v string) *SdkGenerateByAppForRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SdkGenerateByAppForRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
