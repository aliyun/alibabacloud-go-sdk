// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomImageResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomImageResponseBody) *CreateCustomImageResponse
	GetBody() *CreateCustomImageResponseBody
}

type CreateCustomImageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomImageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomImageResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomImageResponse) GetBody() *CreateCustomImageResponseBody {
	return s.Body
}

func (s *CreateCustomImageResponse) SetHeaders(v map[string]*string) *CreateCustomImageResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomImageResponse) SetStatusCode(v int32) *CreateCustomImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomImageResponse) SetBody(v *CreateCustomImageResponseBody) *CreateCustomImageResponse {
	s.Body = v
	return s
}

func (s *CreateCustomImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
