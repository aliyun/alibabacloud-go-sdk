// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse
	GetBody() *DeleteServiceResponseBody
}

type DeleteServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceResponse) GetBody() *DeleteServiceResponseBody {
	return s.Body
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceResponse) SetStatusCode(v int32) *DeleteServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceResponse) SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
