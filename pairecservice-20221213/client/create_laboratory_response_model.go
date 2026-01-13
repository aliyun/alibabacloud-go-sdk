// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLaboratoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLaboratoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLaboratoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateLaboratoryResponseBody) *CreateLaboratoryResponse
	GetBody() *CreateLaboratoryResponseBody
}

type CreateLaboratoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLaboratoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *CreateLaboratoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLaboratoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLaboratoryResponse) GetBody() *CreateLaboratoryResponseBody {
	return s.Body
}

func (s *CreateLaboratoryResponse) SetHeaders(v map[string]*string) *CreateLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *CreateLaboratoryResponse) SetStatusCode(v int32) *CreateLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLaboratoryResponse) SetBody(v *CreateLaboratoryResponseBody) *CreateLaboratoryResponse {
	s.Body = v
	return s
}

func (s *CreateLaboratoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
