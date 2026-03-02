// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnterpriseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEnterpriseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEnterpriseResponse
	GetStatusCode() *int32
	SetBody(v *Enterprise) *UpdateEnterpriseResponse
	GetBody() *Enterprise
}

type UpdateEnterpriseResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Enterprise        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnterpriseResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnterpriseResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnterpriseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEnterpriseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEnterpriseResponse) GetBody() *Enterprise {
	return s.Body
}

func (s *UpdateEnterpriseResponse) SetHeaders(v map[string]*string) *UpdateEnterpriseResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnterpriseResponse) SetStatusCode(v int32) *UpdateEnterpriseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnterpriseResponse) SetBody(v *Enterprise) *UpdateEnterpriseResponse {
	s.Body = v
	return s
}

func (s *UpdateEnterpriseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
