// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardInValidMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStandardInValidMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStandardInValidMappingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStandardInValidMappingResponseBody) *DeleteStandardInValidMappingResponse
	GetBody() *DeleteStandardInValidMappingResponseBody
}

type DeleteStandardInValidMappingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStandardInValidMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStandardInValidMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardInValidMappingResponse) GoString() string {
	return s.String()
}

func (s *DeleteStandardInValidMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStandardInValidMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStandardInValidMappingResponse) GetBody() *DeleteStandardInValidMappingResponseBody {
	return s.Body
}

func (s *DeleteStandardInValidMappingResponse) SetHeaders(v map[string]*string) *DeleteStandardInValidMappingResponse {
	s.Headers = v
	return s
}

func (s *DeleteStandardInValidMappingResponse) SetStatusCode(v int32) *DeleteStandardInValidMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStandardInValidMappingResponse) SetBody(v *DeleteStandardInValidMappingResponseBody) *DeleteStandardInValidMappingResponse {
	s.Body = v
	return s
}

func (s *DeleteStandardInValidMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
