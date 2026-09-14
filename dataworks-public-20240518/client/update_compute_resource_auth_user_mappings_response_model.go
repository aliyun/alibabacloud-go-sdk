// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeResourceAuthUserMappingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateComputeResourceAuthUserMappingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateComputeResourceAuthUserMappingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateComputeResourceAuthUserMappingsResponseBody) *UpdateComputeResourceAuthUserMappingsResponse
	GetBody() *UpdateComputeResourceAuthUserMappingsResponseBody
}

type UpdateComputeResourceAuthUserMappingsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateComputeResourceAuthUserMappingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateComputeResourceAuthUserMappingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeResourceAuthUserMappingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateComputeResourceAuthUserMappingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateComputeResourceAuthUserMappingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateComputeResourceAuthUserMappingsResponse) GetBody() *UpdateComputeResourceAuthUserMappingsResponseBody {
	return s.Body
}

func (s *UpdateComputeResourceAuthUserMappingsResponse) SetHeaders(v map[string]*string) *UpdateComputeResourceAuthUserMappingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsResponse) SetStatusCode(v int32) *UpdateComputeResourceAuthUserMappingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsResponse) SetBody(v *UpdateComputeResourceAuthUserMappingsResponseBody) *UpdateComputeResourceAuthUserMappingsResponse {
	s.Body = v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
