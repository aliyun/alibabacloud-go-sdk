// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeResourceAuthUserMappingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetComputeResourceAuthUserMappingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetComputeResourceAuthUserMappingsResponse
	GetStatusCode() *int32
	SetBody(v *GetComputeResourceAuthUserMappingsResponseBody) *GetComputeResourceAuthUserMappingsResponse
	GetBody() *GetComputeResourceAuthUserMappingsResponseBody
}

type GetComputeResourceAuthUserMappingsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetComputeResourceAuthUserMappingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetComputeResourceAuthUserMappingsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetComputeResourceAuthUserMappingsResponse) GoString() string {
	return s.String()
}

func (s *GetComputeResourceAuthUserMappingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetComputeResourceAuthUserMappingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetComputeResourceAuthUserMappingsResponse) GetBody() *GetComputeResourceAuthUserMappingsResponseBody {
	return s.Body
}

func (s *GetComputeResourceAuthUserMappingsResponse) SetHeaders(v map[string]*string) *GetComputeResourceAuthUserMappingsResponse {
	s.Headers = v
	return s
}

func (s *GetComputeResourceAuthUserMappingsResponse) SetStatusCode(v int32) *GetComputeResourceAuthUserMappingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetComputeResourceAuthUserMappingsResponse) SetBody(v *GetComputeResourceAuthUserMappingsResponseBody) *GetComputeResourceAuthUserMappingsResponse {
	s.Body = v
	return s
}

func (s *GetComputeResourceAuthUserMappingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
