// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchMemoryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PatchMemoryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PatchMemoryConfigResponse
	GetStatusCode() *int32
	SetBody(v *PatchMemoryConfigResponseBody) *PatchMemoryConfigResponse
	GetBody() *PatchMemoryConfigResponseBody
}

type PatchMemoryConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PatchMemoryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PatchMemoryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PatchMemoryConfigResponse) GoString() string {
	return s.String()
}

func (s *PatchMemoryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PatchMemoryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PatchMemoryConfigResponse) GetBody() *PatchMemoryConfigResponseBody {
	return s.Body
}

func (s *PatchMemoryConfigResponse) SetHeaders(v map[string]*string) *PatchMemoryConfigResponse {
	s.Headers = v
	return s
}

func (s *PatchMemoryConfigResponse) SetStatusCode(v int32) *PatchMemoryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PatchMemoryConfigResponse) SetBody(v *PatchMemoryConfigResponseBody) *PatchMemoryConfigResponse {
	s.Body = v
	return s
}

func (s *PatchMemoryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
