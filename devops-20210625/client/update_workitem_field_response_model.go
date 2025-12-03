// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkitemFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkitemFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkitemFieldResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkitemFieldResponseBody) *UpdateWorkitemFieldResponse
	GetBody() *UpdateWorkitemFieldResponseBody
}

type UpdateWorkitemFieldResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkitemFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkitemFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkitemFieldResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkitemFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkitemFieldResponse) GetBody() *UpdateWorkitemFieldResponseBody {
	return s.Body
}

func (s *UpdateWorkitemFieldResponse) SetHeaders(v map[string]*string) *UpdateWorkitemFieldResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkitemFieldResponse) SetStatusCode(v int32) *UpdateWorkitemFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkitemFieldResponse) SetBody(v *UpdateWorkitemFieldResponseBody) *UpdateWorkitemFieldResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkitemFieldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
