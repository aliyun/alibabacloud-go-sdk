// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateLabelResponse
	GetStatusCode() *int32
	SetBody(v *GenerateLabelResponseBody) *GenerateLabelResponse
	GetBody() *GenerateLabelResponseBody
}

type GenerateLabelResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateLabelResponse) GoString() string {
	return s.String()
}

func (s *GenerateLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateLabelResponse) GetBody() *GenerateLabelResponseBody {
	return s.Body
}

func (s *GenerateLabelResponse) SetHeaders(v map[string]*string) *GenerateLabelResponse {
	s.Headers = v
	return s
}

func (s *GenerateLabelResponse) SetStatusCode(v int32) *GenerateLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateLabelResponse) SetBody(v *GenerateLabelResponseBody) *GenerateLabelResponse {
	s.Body = v
	return s
}

func (s *GenerateLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
