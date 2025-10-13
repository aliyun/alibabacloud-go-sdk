// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlgorithmVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlgorithmVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlgorithmVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlgorithmVersionResponseBody) *DeleteAlgorithmVersionResponse
	GetBody() *DeleteAlgorithmVersionResponseBody
}

type DeleteAlgorithmVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlgorithmVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlgorithmVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlgorithmVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlgorithmVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlgorithmVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlgorithmVersionResponse) GetBody() *DeleteAlgorithmVersionResponseBody {
	return s.Body
}

func (s *DeleteAlgorithmVersionResponse) SetHeaders(v map[string]*string) *DeleteAlgorithmVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlgorithmVersionResponse) SetStatusCode(v int32) *DeleteAlgorithmVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlgorithmVersionResponse) SetBody(v *DeleteAlgorithmVersionResponseBody) *DeleteAlgorithmVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteAlgorithmVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
