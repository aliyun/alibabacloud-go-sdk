// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsPathMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolarFsPathMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolarFsPathMappingResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolarFsPathMappingResponseBody) *DeletePolarFsPathMappingResponse
	GetBody() *DeletePolarFsPathMappingResponseBody
}

type DeletePolarFsPathMappingResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolarFsPathMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolarFsPathMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsPathMappingResponse) GoString() string {
	return s.String()
}

func (s *DeletePolarFsPathMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolarFsPathMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolarFsPathMappingResponse) GetBody() *DeletePolarFsPathMappingResponseBody {
	return s.Body
}

func (s *DeletePolarFsPathMappingResponse) SetHeaders(v map[string]*string) *DeletePolarFsPathMappingResponse {
	s.Headers = v
	return s
}

func (s *DeletePolarFsPathMappingResponse) SetStatusCode(v int32) *DeletePolarFsPathMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolarFsPathMappingResponse) SetBody(v *DeletePolarFsPathMappingResponseBody) *DeletePolarFsPathMappingResponse {
	s.Body = v
	return s
}

func (s *DeletePolarFsPathMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
