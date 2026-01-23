// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardMappingToInvalidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStandardMappingToInvalidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStandardMappingToInvalidResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStandardMappingToInvalidResponseBody) *UpdateStandardMappingToInvalidResponse
	GetBody() *UpdateStandardMappingToInvalidResponseBody
}

type UpdateStandardMappingToInvalidResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStandardMappingToInvalidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStandardMappingToInvalidResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardMappingToInvalidResponse) GoString() string {
	return s.String()
}

func (s *UpdateStandardMappingToInvalidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStandardMappingToInvalidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStandardMappingToInvalidResponse) GetBody() *UpdateStandardMappingToInvalidResponseBody {
	return s.Body
}

func (s *UpdateStandardMappingToInvalidResponse) SetHeaders(v map[string]*string) *UpdateStandardMappingToInvalidResponse {
	s.Headers = v
	return s
}

func (s *UpdateStandardMappingToInvalidResponse) SetStatusCode(v int32) *UpdateStandardMappingToInvalidResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStandardMappingToInvalidResponse) SetBody(v *UpdateStandardMappingToInvalidResponseBody) *UpdateStandardMappingToInvalidResponse {
	s.Body = v
	return s
}

func (s *UpdateStandardMappingToInvalidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
