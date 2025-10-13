// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlgorithmVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlgorithmVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlgorithmVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlgorithmVersionResponseBody) *CreateAlgorithmVersionResponse
	GetBody() *CreateAlgorithmVersionResponseBody
}

type CreateAlgorithmVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlgorithmVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlgorithmVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlgorithmVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlgorithmVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlgorithmVersionResponse) GetBody() *CreateAlgorithmVersionResponseBody {
	return s.Body
}

func (s *CreateAlgorithmVersionResponse) SetHeaders(v map[string]*string) *CreateAlgorithmVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateAlgorithmVersionResponse) SetStatusCode(v int32) *CreateAlgorithmVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlgorithmVersionResponse) SetBody(v *CreateAlgorithmVersionResponseBody) *CreateAlgorithmVersionResponse {
	s.Body = v
	return s
}

func (s *CreateAlgorithmVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
