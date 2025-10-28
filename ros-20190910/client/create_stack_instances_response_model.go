// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStackInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStackInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStackInstancesResponse
	GetStatusCode() *int32
	SetBody(v *CreateStackInstancesResponseBody) *CreateStackInstancesResponse
	GetBody() *CreateStackInstancesResponseBody
}

type CreateStackInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStackInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStackInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStackInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStackInstancesResponse) GetBody() *CreateStackInstancesResponseBody {
	return s.Body
}

func (s *CreateStackInstancesResponse) SetHeaders(v map[string]*string) *CreateStackInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateStackInstancesResponse) SetStatusCode(v int32) *CreateStackInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStackInstancesResponse) SetBody(v *CreateStackInstancesResponseBody) *CreateStackInstancesResponse {
	s.Body = v
	return s
}

func (s *CreateStackInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
