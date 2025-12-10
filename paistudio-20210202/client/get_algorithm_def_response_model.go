// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmDefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlgorithmDefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlgorithmDefResponse
	GetStatusCode() *int32
	SetBody(v *GetAlgorithmDefResponseBody) *GetAlgorithmDefResponse
	GetBody() *GetAlgorithmDefResponseBody
}

type GetAlgorithmDefResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlgorithmDefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlgorithmDefResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmDefResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlgorithmDefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlgorithmDefResponse) GetBody() *GetAlgorithmDefResponseBody {
	return s.Body
}

func (s *GetAlgorithmDefResponse) SetHeaders(v map[string]*string) *GetAlgorithmDefResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmDefResponse) SetStatusCode(v int32) *GetAlgorithmDefResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgorithmDefResponse) SetBody(v *GetAlgorithmDefResponseBody) *GetAlgorithmDefResponse {
	s.Body = v
	return s
}

func (s *GetAlgorithmDefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
