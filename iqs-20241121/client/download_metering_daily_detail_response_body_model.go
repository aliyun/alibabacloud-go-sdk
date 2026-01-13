// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadMeteringDailyDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadUrl(v string) *DownloadMeteringDailyDetailResponseBody
	GetDownloadUrl() *string
	SetRequestId(v string) *DownloadMeteringDailyDetailResponseBody
	GetRequestId() *string
}

type DownloadMeteringDailyDetailResponseBody struct {
	// example:
	//
	// https://
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DownloadMeteringDailyDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadMeteringDailyDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadMeteringDailyDetailResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DownloadMeteringDailyDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadMeteringDailyDetailResponseBody) SetDownloadUrl(v string) *DownloadMeteringDailyDetailResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *DownloadMeteringDailyDetailResponseBody) SetRequestId(v string) *DownloadMeteringDailyDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadMeteringDailyDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
