// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUploadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUploadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetUploadUrlResponseBody) *GetUploadUrlResponse
	GetBody() *GetUploadUrlResponseBody
}

type GetUploadUrlResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetUploadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUploadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUploadUrlResponse) GetBody() *GetUploadUrlResponseBody {
	return s.Body
}

func (s *GetUploadUrlResponse) SetHeaders(v map[string]*string) *GetUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetUploadUrlResponse) SetStatusCode(v int32) *GetUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadUrlResponse) SetBody(v *GetUploadUrlResponseBody) *GetUploadUrlResponse {
	s.Body = v
	return s
}

func (s *GetUploadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
