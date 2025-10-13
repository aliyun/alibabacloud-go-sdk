// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlgorithmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlgorithmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlgorithmResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlgorithmResponseBody) *CreateAlgorithmResponse
	GetBody() *CreateAlgorithmResponseBody
}

type CreateAlgorithmResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlgorithmResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlgorithmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlgorithmResponse) GetBody() *CreateAlgorithmResponseBody {
	return s.Body
}

func (s *CreateAlgorithmResponse) SetHeaders(v map[string]*string) *CreateAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *CreateAlgorithmResponse) SetStatusCode(v int32) *CreateAlgorithmResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlgorithmResponse) SetBody(v *CreateAlgorithmResponseBody) *CreateAlgorithmResponse {
	s.Body = v
	return s
}

func (s *CreateAlgorithmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
