// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChunkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChunkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChunkResponse
	GetStatusCode() *int32
	SetBody(v *GetChunkResponseBody) *GetChunkResponse
	GetBody() *GetChunkResponseBody
}

type GetChunkResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChunkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChunkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChunkResponse) GoString() string {
	return s.String()
}

func (s *GetChunkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChunkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChunkResponse) GetBody() *GetChunkResponseBody {
	return s.Body
}

func (s *GetChunkResponse) SetHeaders(v map[string]*string) *GetChunkResponse {
	s.Headers = v
	return s
}

func (s *GetChunkResponse) SetStatusCode(v int32) *GetChunkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChunkResponse) SetBody(v *GetChunkResponseBody) *GetChunkResponse {
	s.Body = v
	return s
}

func (s *GetChunkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
