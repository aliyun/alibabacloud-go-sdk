// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentCodePackageDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentCodePackageDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentCodePackageDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentCodePackageDownloadUrlResponseBody) *GetAgentCodePackageDownloadUrlResponse
	GetBody() *GetAgentCodePackageDownloadUrlResponseBody
}

type GetAgentCodePackageDownloadUrlResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentCodePackageDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentCodePackageDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentCodePackageDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentCodePackageDownloadUrlResponse) GetBody() *GetAgentCodePackageDownloadUrlResponseBody {
	return s.Body
}

func (s *GetAgentCodePackageDownloadUrlResponse) SetHeaders(v map[string]*string) *GetAgentCodePackageDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponse) SetStatusCode(v int32) *GetAgentCodePackageDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponse) SetBody(v *GetAgentCodePackageDownloadUrlResponseBody) *GetAgentCodePackageDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
