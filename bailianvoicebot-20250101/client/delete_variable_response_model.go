// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVariableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVariableResponseBody) *DeleteVariableResponse
	GetBody() *DeleteVariableResponseBody
}

type DeleteVariableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVariableResponse) GoString() string {
	return s.String()
}

func (s *DeleteVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVariableResponse) GetBody() *DeleteVariableResponseBody {
	return s.Body
}

func (s *DeleteVariableResponse) SetHeaders(v map[string]*string) *DeleteVariableResponse {
	s.Headers = v
	return s
}

func (s *DeleteVariableResponse) SetStatusCode(v int32) *DeleteVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVariableResponse) SetBody(v *DeleteVariableResponseBody) *DeleteVariableResponse {
	s.Body = v
	return s
}

func (s *DeleteVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
