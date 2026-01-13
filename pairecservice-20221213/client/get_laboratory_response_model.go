// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLaboratoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLaboratoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLaboratoryResponse
	GetStatusCode() *int32
	SetBody(v *GetLaboratoryResponseBody) *GetLaboratoryResponse
	GetBody() *GetLaboratoryResponseBody
}

type GetLaboratoryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLaboratoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *GetLaboratoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLaboratoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLaboratoryResponse) GetBody() *GetLaboratoryResponseBody {
	return s.Body
}

func (s *GetLaboratoryResponse) SetHeaders(v map[string]*string) *GetLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *GetLaboratoryResponse) SetStatusCode(v int32) *GetLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLaboratoryResponse) SetBody(v *GetLaboratoryResponseBody) *GetLaboratoryResponse {
	s.Body = v
	return s
}

func (s *GetLaboratoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
