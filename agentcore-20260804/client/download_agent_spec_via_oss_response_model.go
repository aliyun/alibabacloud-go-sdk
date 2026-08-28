// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadAgentSpecViaOssResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadAgentSpecViaOssResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadAgentSpecViaOssResponse
	GetStatusCode() *int32
	SetBody(v *DownloadAgentSpecViaOssResponseBody) *DownloadAgentSpecViaOssResponse
	GetBody() *DownloadAgentSpecViaOssResponseBody
}

type DownloadAgentSpecViaOssResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadAgentSpecViaOssResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadAgentSpecViaOssResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadAgentSpecViaOssResponse) GoString() string {
	return s.String()
}

func (s *DownloadAgentSpecViaOssResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadAgentSpecViaOssResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadAgentSpecViaOssResponse) GetBody() *DownloadAgentSpecViaOssResponseBody {
	return s.Body
}

func (s *DownloadAgentSpecViaOssResponse) SetHeaders(v map[string]*string) *DownloadAgentSpecViaOssResponse {
	s.Headers = v
	return s
}

func (s *DownloadAgentSpecViaOssResponse) SetStatusCode(v int32) *DownloadAgentSpecViaOssResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadAgentSpecViaOssResponse) SetBody(v *DownloadAgentSpecViaOssResponseBody) *DownloadAgentSpecViaOssResponse {
	s.Body = v
	return s
}

func (s *DownloadAgentSpecViaOssResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
