// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlgorithmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlgorithmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlgorithmResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlgorithmResponseBody) *DeleteAlgorithmResponse
	GetBody() *DeleteAlgorithmResponseBody
}

type DeleteAlgorithmResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlgorithmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlgorithmResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlgorithmResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlgorithmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlgorithmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlgorithmResponse) GetBody() *DeleteAlgorithmResponseBody {
	return s.Body
}

func (s *DeleteAlgorithmResponse) SetHeaders(v map[string]*string) *DeleteAlgorithmResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlgorithmResponse) SetStatusCode(v int32) *DeleteAlgorithmResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlgorithmResponse) SetBody(v *DeleteAlgorithmResponseBody) *DeleteAlgorithmResponse {
	s.Body = v
	return s
}

func (s *DeleteAlgorithmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
