// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUploadImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUploadImageResponse
	GetStatusCode() *int32
	SetBody(v *CreateUploadImageResponseBody) *CreateUploadImageResponse
	GetBody() *CreateUploadImageResponseBody
}

type CreateUploadImageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUploadImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUploadImageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadImageResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUploadImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUploadImageResponse) GetBody() *CreateUploadImageResponseBody {
	return s.Body
}

func (s *CreateUploadImageResponse) SetHeaders(v map[string]*string) *CreateUploadImageResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadImageResponse) SetStatusCode(v int32) *CreateUploadImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUploadImageResponse) SetBody(v *CreateUploadImageResponseBody) *CreateUploadImageResponse {
	s.Body = v
	return s
}

func (s *CreateUploadImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
