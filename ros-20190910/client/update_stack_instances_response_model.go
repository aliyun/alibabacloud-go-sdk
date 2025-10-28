// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStackInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStackInstancesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStackInstancesResponseBody) *UpdateStackInstancesResponse
	GetBody() *UpdateStackInstancesResponseBody
}

type UpdateStackInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStackInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackInstancesResponse) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStackInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStackInstancesResponse) GetBody() *UpdateStackInstancesResponseBody {
	return s.Body
}

func (s *UpdateStackInstancesResponse) SetHeaders(v map[string]*string) *UpdateStackInstancesResponse {
	s.Headers = v
	return s
}

func (s *UpdateStackInstancesResponse) SetStatusCode(v int32) *UpdateStackInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStackInstancesResponse) SetBody(v *UpdateStackInstancesResponseBody) *UpdateStackInstancesResponse {
	s.Body = v
	return s
}

func (s *UpdateStackInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
