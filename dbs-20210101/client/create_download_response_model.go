// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDownloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDownloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDownloadResponse
	GetStatusCode() *int32
	SetBody(v *CreateDownloadResponseBody) *CreateDownloadResponse
	GetBody() *CreateDownloadResponseBody
}

type CreateDownloadResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDownloadResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDownloadResponse) GoString() string {
	return s.String()
}

func (s *CreateDownloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDownloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDownloadResponse) GetBody() *CreateDownloadResponseBody {
	return s.Body
}

func (s *CreateDownloadResponse) SetHeaders(v map[string]*string) *CreateDownloadResponse {
	s.Headers = v
	return s
}

func (s *CreateDownloadResponse) SetStatusCode(v int32) *CreateDownloadResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDownloadResponse) SetBody(v *CreateDownloadResponseBody) *CreateDownloadResponse {
	s.Body = v
	return s
}

func (s *CreateDownloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
