// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadFileResponse
	GetStatusCode() *int32
	SetBody(v *UploadFileResponseBody) *UploadFileResponse
	GetBody() *UploadFileResponseBody
}

type UploadFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadFileResponse) GoString() string {
	return s.String()
}

func (s *UploadFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadFileResponse) GetBody() *UploadFileResponseBody {
	return s.Body
}

func (s *UploadFileResponse) SetHeaders(v map[string]*string) *UploadFileResponse {
	s.Headers = v
	return s
}

func (s *UploadFileResponse) SetStatusCode(v int32) *UploadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadFileResponse) SetBody(v *UploadFileResponseBody) *UploadFileResponse {
	s.Body = v
	return s
}

func (s *UploadFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
