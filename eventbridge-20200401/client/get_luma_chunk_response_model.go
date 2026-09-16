// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaChunkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLumaChunkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLumaChunkResponse
	GetStatusCode() *int32
	SetBody(v *GetLumaChunkResponseBody) *GetLumaChunkResponse
	GetBody() *GetLumaChunkResponseBody
}

type GetLumaChunkResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLumaChunkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLumaChunkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLumaChunkResponse) GoString() string {
	return s.String()
}

func (s *GetLumaChunkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLumaChunkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLumaChunkResponse) GetBody() *GetLumaChunkResponseBody {
	return s.Body
}

func (s *GetLumaChunkResponse) SetHeaders(v map[string]*string) *GetLumaChunkResponse {
	s.Headers = v
	return s
}

func (s *GetLumaChunkResponse) SetStatusCode(v int32) *GetLumaChunkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLumaChunkResponse) SetBody(v *GetLumaChunkResponseBody) *GetLumaChunkResponse {
	s.Body = v
	return s
}

func (s *GetLumaChunkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
