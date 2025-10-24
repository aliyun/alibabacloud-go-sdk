// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomViewResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomViewResponseBody) *DeleteCustomViewResponse
	GetBody() *DeleteCustomViewResponseBody
}

type DeleteCustomViewResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomViewResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomViewResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomViewResponse) GetBody() *DeleteCustomViewResponseBody {
	return s.Body
}

func (s *DeleteCustomViewResponse) SetHeaders(v map[string]*string) *DeleteCustomViewResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomViewResponse) SetStatusCode(v int32) *DeleteCustomViewResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomViewResponse) SetBody(v *DeleteCustomViewResponseBody) *DeleteCustomViewResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
