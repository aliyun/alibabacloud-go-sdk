// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTempFileDownloadLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTempFileDownloadLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTempFileDownloadLinkResponse
	GetStatusCode() *int32
	SetBody(v *GetTempFileDownloadLinkResponseBody) *GetTempFileDownloadLinkResponse
	GetBody() *GetTempFileDownloadLinkResponseBody
}

type GetTempFileDownloadLinkResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTempFileDownloadLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTempFileDownloadLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTempFileDownloadLinkResponse) GoString() string {
	return s.String()
}

func (s *GetTempFileDownloadLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTempFileDownloadLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTempFileDownloadLinkResponse) GetBody() *GetTempFileDownloadLinkResponseBody {
	return s.Body
}

func (s *GetTempFileDownloadLinkResponse) SetHeaders(v map[string]*string) *GetTempFileDownloadLinkResponse {
	s.Headers = v
	return s
}

func (s *GetTempFileDownloadLinkResponse) SetStatusCode(v int32) *GetTempFileDownloadLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTempFileDownloadLinkResponse) SetBody(v *GetTempFileDownloadLinkResponseBody) *GetTempFileDownloadLinkResponse {
	s.Body = v
	return s
}

func (s *GetTempFileDownloadLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
