// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VideoGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VideoGenerationResponse
	GetStatusCode() *int32
	SetBody(v *VideoGenerationResponseBody) *VideoGenerationResponse
	GetBody() *VideoGenerationResponseBody
}

type VideoGenerationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VideoGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s VideoGenerationResponse) GoString() string {
	return s.String()
}

func (s *VideoGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VideoGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VideoGenerationResponse) GetBody() *VideoGenerationResponseBody {
	return s.Body
}

func (s *VideoGenerationResponse) SetHeaders(v map[string]*string) *VideoGenerationResponse {
	s.Headers = v
	return s
}

func (s *VideoGenerationResponse) SetStatusCode(v int32) *VideoGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoGenerationResponse) SetBody(v *VideoGenerationResponseBody) *VideoGenerationResponse {
	s.Body = v
	return s
}

func (s *VideoGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
