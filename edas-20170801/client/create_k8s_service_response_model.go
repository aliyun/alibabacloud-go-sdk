// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateK8sServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateK8sServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateK8sServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateK8sServiceResponseBody) *CreateK8sServiceResponse
	GetBody() *CreateK8sServiceResponseBody
}

type CreateK8sServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateK8sServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateK8sServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateK8sServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateK8sServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateK8sServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateK8sServiceResponse) GetBody() *CreateK8sServiceResponseBody {
	return s.Body
}

func (s *CreateK8sServiceResponse) SetHeaders(v map[string]*string) *CreateK8sServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateK8sServiceResponse) SetStatusCode(v int32) *CreateK8sServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateK8sServiceResponse) SetBody(v *CreateK8sServiceResponseBody) *CreateK8sServiceResponse {
	s.Body = v
	return s
}

func (s *CreateK8sServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
