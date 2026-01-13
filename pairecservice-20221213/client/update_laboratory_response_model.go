// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLaboratoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLaboratoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLaboratoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLaboratoryResponseBody) *UpdateLaboratoryResponse
	GetBody() *UpdateLaboratoryResponseBody
}

type UpdateLaboratoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLaboratoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateLaboratoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLaboratoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLaboratoryResponse) GetBody() *UpdateLaboratoryResponseBody {
	return s.Body
}

func (s *UpdateLaboratoryResponse) SetHeaders(v map[string]*string) *UpdateLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateLaboratoryResponse) SetStatusCode(v int32) *UpdateLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLaboratoryResponse) SetBody(v *UpdateLaboratoryResponseBody) *UpdateLaboratoryResponse {
	s.Body = v
	return s
}

func (s *UpdateLaboratoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
