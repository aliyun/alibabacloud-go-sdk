// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepublishModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PrepublishModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PrepublishModelResponse
	GetStatusCode() *int32
	SetBody(v *PrepublishModelResponseBody) *PrepublishModelResponse
	GetBody() *PrepublishModelResponseBody
}

type PrepublishModelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrepublishModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrepublishModelResponse) String() string {
	return dara.Prettify(s)
}

func (s PrepublishModelResponse) GoString() string {
	return s.String()
}

func (s *PrepublishModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PrepublishModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PrepublishModelResponse) GetBody() *PrepublishModelResponseBody {
	return s.Body
}

func (s *PrepublishModelResponse) SetHeaders(v map[string]*string) *PrepublishModelResponse {
	s.Headers = v
	return s
}

func (s *PrepublishModelResponse) SetStatusCode(v int32) *PrepublishModelResponse {
	s.StatusCode = &v
	return s
}

func (s *PrepublishModelResponse) SetBody(v *PrepublishModelResponseBody) *PrepublishModelResponse {
	s.Body = v
	return s
}

func (s *PrepublishModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
