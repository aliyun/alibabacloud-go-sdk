// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaDocumentDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLumaDocumentDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLumaDocumentDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetLumaDocumentDownloadUrlResponseBody) *GetLumaDocumentDownloadUrlResponse
	GetBody() *GetLumaDocumentDownloadUrlResponseBody
}

type GetLumaDocumentDownloadUrlResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLumaDocumentDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLumaDocumentDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLumaDocumentDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetLumaDocumentDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLumaDocumentDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLumaDocumentDownloadUrlResponse) GetBody() *GetLumaDocumentDownloadUrlResponseBody {
	return s.Body
}

func (s *GetLumaDocumentDownloadUrlResponse) SetHeaders(v map[string]*string) *GetLumaDocumentDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetLumaDocumentDownloadUrlResponse) SetStatusCode(v int32) *GetLumaDocumentDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlResponse) SetBody(v *GetLumaDocumentDownloadUrlResponseBody) *GetLumaDocumentDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *GetLumaDocumentDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
