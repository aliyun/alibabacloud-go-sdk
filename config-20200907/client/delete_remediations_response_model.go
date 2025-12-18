// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRemediationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRemediationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRemediationsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRemediationsResponseBody) *DeleteRemediationsResponse
	GetBody() *DeleteRemediationsResponseBody
}

type DeleteRemediationsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRemediationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRemediationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRemediationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteRemediationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRemediationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRemediationsResponse) GetBody() *DeleteRemediationsResponseBody {
	return s.Body
}

func (s *DeleteRemediationsResponse) SetHeaders(v map[string]*string) *DeleteRemediationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteRemediationsResponse) SetStatusCode(v int32) *DeleteRemediationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRemediationsResponse) SetBody(v *DeleteRemediationsResponseBody) *DeleteRemediationsResponse {
	s.Body = v
	return s
}

func (s *DeleteRemediationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
