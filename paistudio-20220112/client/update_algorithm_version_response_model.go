// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlgorithmVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlgorithmVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlgorithmVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlgorithmVersionResponseBody) *UpdateAlgorithmVersionResponse
	GetBody() *UpdateAlgorithmVersionResponseBody
}

type UpdateAlgorithmVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlgorithmVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlgorithmVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlgorithmVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlgorithmVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlgorithmVersionResponse) GetBody() *UpdateAlgorithmVersionResponseBody {
	return s.Body
}

func (s *UpdateAlgorithmVersionResponse) SetHeaders(v map[string]*string) *UpdateAlgorithmVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlgorithmVersionResponse) SetStatusCode(v int32) *UpdateAlgorithmVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlgorithmVersionResponse) SetBody(v *UpdateAlgorithmVersionResponseBody) *UpdateAlgorithmVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateAlgorithmVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
