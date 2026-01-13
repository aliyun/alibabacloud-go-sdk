// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLaboratoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLaboratoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLaboratoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLaboratoryResponseBody) *DeleteLaboratoryResponse
	GetBody() *DeleteLaboratoryResponseBody
}

type DeleteLaboratoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLaboratoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteLaboratoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLaboratoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLaboratoryResponse) GetBody() *DeleteLaboratoryResponseBody {
	return s.Body
}

func (s *DeleteLaboratoryResponse) SetHeaders(v map[string]*string) *DeleteLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteLaboratoryResponse) SetStatusCode(v int32) *DeleteLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLaboratoryResponse) SetBody(v *DeleteLaboratoryResponseBody) *DeleteLaboratoryResponse {
	s.Body = v
	return s
}

func (s *DeleteLaboratoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
