// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContainerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContainerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContainerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContainerResponseBody) *UpdateContainerResponse
	GetBody() *UpdateContainerResponseBody
}

type UpdateContainerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContainerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContainerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContainerResponse) GoString() string {
	return s.String()
}

func (s *UpdateContainerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContainerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContainerResponse) GetBody() *UpdateContainerResponseBody {
	return s.Body
}

func (s *UpdateContainerResponse) SetHeaders(v map[string]*string) *UpdateContainerResponse {
	s.Headers = v
	return s
}

func (s *UpdateContainerResponse) SetStatusCode(v int32) *UpdateContainerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContainerResponse) SetBody(v *UpdateContainerResponseBody) *UpdateContainerResponse {
	s.Body = v
	return s
}

func (s *UpdateContainerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
