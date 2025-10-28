// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteK8sServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteK8sServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteK8sServiceResponseBody) *DeleteK8sServiceResponse
	GetBody() *DeleteK8sServiceResponseBody
}

type DeleteK8sServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteK8sServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteK8sServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteK8sServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteK8sServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteK8sServiceResponse) GetBody() *DeleteK8sServiceResponseBody {
	return s.Body
}

func (s *DeleteK8sServiceResponse) SetHeaders(v map[string]*string) *DeleteK8sServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteK8sServiceResponse) SetStatusCode(v int32) *DeleteK8sServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteK8sServiceResponse) SetBody(v *DeleteK8sServiceResponseBody) *DeleteK8sServiceResponse {
	s.Body = v
	return s
}

func (s *DeleteK8sServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
