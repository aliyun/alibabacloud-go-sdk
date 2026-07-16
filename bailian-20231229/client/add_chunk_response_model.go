// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChunkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddChunkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddChunkResponse
	GetStatusCode() *int32
	SetBody(v *AddChunkResponseBody) *AddChunkResponse
	GetBody() *AddChunkResponseBody
}

type AddChunkResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddChunkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddChunkResponse) String() string {
	return dara.Prettify(s)
}

func (s AddChunkResponse) GoString() string {
	return s.String()
}

func (s *AddChunkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddChunkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddChunkResponse) GetBody() *AddChunkResponseBody {
	return s.Body
}

func (s *AddChunkResponse) SetHeaders(v map[string]*string) *AddChunkResponse {
	s.Headers = v
	return s
}

func (s *AddChunkResponse) SetStatusCode(v int32) *AddChunkResponse {
	s.StatusCode = &v
	return s
}

func (s *AddChunkResponse) SetBody(v *AddChunkResponseBody) *AddChunkResponse {
	s.Body = v
	return s
}

func (s *AddChunkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
