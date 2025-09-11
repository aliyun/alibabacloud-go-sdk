// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadClusterManagedCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadClusterManagedCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadClusterManagedCertResponse
	GetStatusCode() *int32
	SetBody(v *DownloadClusterManagedCertResponseBody) *DownloadClusterManagedCertResponse
	GetBody() *DownloadClusterManagedCertResponseBody
}

type DownloadClusterManagedCertResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadClusterManagedCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadClusterManagedCertResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadClusterManagedCertResponse) GoString() string {
	return s.String()
}

func (s *DownloadClusterManagedCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadClusterManagedCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadClusterManagedCertResponse) GetBody() *DownloadClusterManagedCertResponseBody {
	return s.Body
}

func (s *DownloadClusterManagedCertResponse) SetHeaders(v map[string]*string) *DownloadClusterManagedCertResponse {
	s.Headers = v
	return s
}

func (s *DownloadClusterManagedCertResponse) SetStatusCode(v int32) *DownloadClusterManagedCertResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadClusterManagedCertResponse) SetBody(v *DownloadClusterManagedCertResponseBody) *DownloadClusterManagedCertResponse {
	s.Body = v
	return s
}

func (s *DownloadClusterManagedCertResponse) Validate() error {
	return dara.Validate(s)
}
