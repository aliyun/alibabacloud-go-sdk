// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSolutionInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSolutionInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSolutionInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSolutionInstanceAttributeResponseBody) *UpdateSolutionInstanceAttributeResponse
	GetBody() *UpdateSolutionInstanceAttributeResponseBody
}

type UpdateSolutionInstanceAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSolutionInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSolutionInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSolutionInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSolutionInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSolutionInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSolutionInstanceAttributeResponse) GetBody() *UpdateSolutionInstanceAttributeResponseBody {
	return s.Body
}

func (s *UpdateSolutionInstanceAttributeResponse) SetHeaders(v map[string]*string) *UpdateSolutionInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSolutionInstanceAttributeResponse) SetStatusCode(v int32) *UpdateSolutionInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSolutionInstanceAttributeResponse) SetBody(v *UpdateSolutionInstanceAttributeResponseBody) *UpdateSolutionInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateSolutionInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
