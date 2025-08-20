// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadInstanceCACertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadUrl(v string) *DownloadInstanceCACertificateResponseBody
	GetDownloadUrl() *string
	SetRequestId(v string) *DownloadInstanceCACertificateResponseBody
	GetRequestId() *string
}

type DownloadInstanceCACertificateResponseBody struct {
	// The OSS URL of the downloaded certificate.
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9CCFAAB4-97B7-5800-B9F2-685EB596E3EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadInstanceCACertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadInstanceCACertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadInstanceCACertificateResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DownloadInstanceCACertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadInstanceCACertificateResponseBody) SetDownloadUrl(v string) *DownloadInstanceCACertificateResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *DownloadInstanceCACertificateResponseBody) SetRequestId(v string) *DownloadInstanceCACertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadInstanceCACertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
