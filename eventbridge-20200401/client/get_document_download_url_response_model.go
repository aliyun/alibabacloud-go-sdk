// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentDownloadUrlResponseBody) *GetDocumentDownloadUrlResponse
	GetBody() *GetDocumentDownloadUrlResponseBody
}

type GetDocumentDownloadUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentDownloadUrlResponse) GetBody() *GetDocumentDownloadUrlResponseBody {
	return s.Body
}

func (s *GetDocumentDownloadUrlResponse) SetHeaders(v map[string]*string) *GetDocumentDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentDownloadUrlResponse) SetStatusCode(v int32) *GetDocumentDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentDownloadUrlResponse) SetBody(v *GetDocumentDownloadUrlResponseBody) *GetDocumentDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *GetDocumentDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
