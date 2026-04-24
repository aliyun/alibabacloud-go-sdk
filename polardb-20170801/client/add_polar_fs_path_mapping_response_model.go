// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPolarFsPathMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPolarFsPathMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPolarFsPathMappingResponse
	GetStatusCode() *int32
	SetBody(v *AddPolarFsPathMappingResponseBody) *AddPolarFsPathMappingResponse
	GetBody() *AddPolarFsPathMappingResponseBody
}

type AddPolarFsPathMappingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPolarFsPathMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPolarFsPathMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPolarFsPathMappingResponse) GoString() string {
	return s.String()
}

func (s *AddPolarFsPathMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPolarFsPathMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPolarFsPathMappingResponse) GetBody() *AddPolarFsPathMappingResponseBody {
	return s.Body
}

func (s *AddPolarFsPathMappingResponse) SetHeaders(v map[string]*string) *AddPolarFsPathMappingResponse {
	s.Headers = v
	return s
}

func (s *AddPolarFsPathMappingResponse) SetStatusCode(v int32) *AddPolarFsPathMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPolarFsPathMappingResponse) SetBody(v *AddPolarFsPathMappingResponseBody) *AddPolarFsPathMappingResponse {
	s.Body = v
	return s
}

func (s *AddPolarFsPathMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
