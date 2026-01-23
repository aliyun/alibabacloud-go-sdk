// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardValidMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStandardValidMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStandardValidMappingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStandardValidMappingResponseBody) *DeleteStandardValidMappingResponse
	GetBody() *DeleteStandardValidMappingResponseBody
}

type DeleteStandardValidMappingResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStandardValidMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStandardValidMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardValidMappingResponse) GoString() string {
	return s.String()
}

func (s *DeleteStandardValidMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStandardValidMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStandardValidMappingResponse) GetBody() *DeleteStandardValidMappingResponseBody {
	return s.Body
}

func (s *DeleteStandardValidMappingResponse) SetHeaders(v map[string]*string) *DeleteStandardValidMappingResponse {
	s.Headers = v
	return s
}

func (s *DeleteStandardValidMappingResponse) SetStatusCode(v int32) *DeleteStandardValidMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStandardValidMappingResponse) SetBody(v *DeleteStandardValidMappingResponseBody) *DeleteStandardValidMappingResponse {
	s.Body = v
	return s
}

func (s *DeleteStandardValidMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
