// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMemoryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMemoryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMemoryConfigResponse
	GetStatusCode() *int32
	SetBody(v *QueryMemoryConfigResponseBody) *QueryMemoryConfigResponse
	GetBody() *QueryMemoryConfigResponseBody
}

type QueryMemoryConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMemoryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMemoryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMemoryConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryMemoryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMemoryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMemoryConfigResponse) GetBody() *QueryMemoryConfigResponseBody {
	return s.Body
}

func (s *QueryMemoryConfigResponse) SetHeaders(v map[string]*string) *QueryMemoryConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryMemoryConfigResponse) SetStatusCode(v int32) *QueryMemoryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMemoryConfigResponse) SetBody(v *QueryMemoryConfigResponseBody) *QueryMemoryConfigResponse {
	s.Body = v
	return s
}

func (s *QueryMemoryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
