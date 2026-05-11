// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *CreateDownloadUrlResponseBody) *CreateDownloadUrlResponse
	GetBody() *CreateDownloadUrlResponseBody
}

type CreateDownloadUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDownloadUrlResponse) GetBody() *CreateDownloadUrlResponseBody {
	return s.Body
}

func (s *CreateDownloadUrlResponse) SetHeaders(v map[string]*string) *CreateDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateDownloadUrlResponse) SetStatusCode(v int32) *CreateDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDownloadUrlResponse) SetBody(v *CreateDownloadUrlResponseBody) *CreateDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *CreateDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
