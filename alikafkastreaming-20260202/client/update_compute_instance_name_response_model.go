// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeInstanceNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComputeInstanceNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComputeInstanceNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComputeInstanceNameResponseBody) *UpdateComputeInstanceNameResponse
	GetBody() *UpdateComputeInstanceNameResponseBody
}

type UpdateComputeInstanceNameResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeInstanceNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeInstanceNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComputeInstanceNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComputeInstanceNameResponse) GetBody() *UpdateComputeInstanceNameResponseBody {
	return s.Body
}

func (s *UpdateComputeInstanceNameResponse) SetHeaders(v map[string]*string) *UpdateComputeInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeInstanceNameResponse) SetStatusCode(v int32) *UpdateComputeInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeInstanceNameResponse) SetBody(v *UpdateComputeInstanceNameResponseBody) *UpdateComputeInstanceNameResponse {
	s.Body = v
	return s
}

func (s *UpdateComputeInstanceNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
