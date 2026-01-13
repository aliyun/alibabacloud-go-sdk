// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadApiCallDailyDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadUrl(v string) *DownloadApiCallDailyDetailResponseBody
	GetDownloadUrl() *string
	SetRequestId(v string) *DownloadApiCallDailyDetailResponseBody
	GetRequestId() *string
}

type DownloadApiCallDailyDetailResponseBody struct {
	// example:
	//
	// https://
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 37A6ED15-DB44-50B5-B0DF-9449EED8FDF4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DownloadApiCallDailyDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadApiCallDailyDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadApiCallDailyDetailResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DownloadApiCallDailyDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadApiCallDailyDetailResponseBody) SetDownloadUrl(v string) *DownloadApiCallDailyDetailResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *DownloadApiCallDailyDetailResponseBody) SetRequestId(v string) *DownloadApiCallDailyDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadApiCallDailyDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
