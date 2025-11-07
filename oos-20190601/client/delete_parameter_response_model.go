// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteParameterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteParameterResponseBody) *DeleteParameterResponse
	GetBody() *DeleteParameterResponseBody
}

type DeleteParameterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterResponse) GoString() string {
	return s.String()
}

func (s *DeleteParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteParameterResponse) GetBody() *DeleteParameterResponseBody {
	return s.Body
}

func (s *DeleteParameterResponse) SetHeaders(v map[string]*string) *DeleteParameterResponse {
	s.Headers = v
	return s
}

func (s *DeleteParameterResponse) SetStatusCode(v int32) *DeleteParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteParameterResponse) SetBody(v *DeleteParameterResponseBody) *DeleteParameterResponse {
	s.Body = v
	return s
}

func (s *DeleteParameterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
