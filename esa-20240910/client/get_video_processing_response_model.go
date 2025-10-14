// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoProcessingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoProcessingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoProcessingResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoProcessingResponseBody) *GetVideoProcessingResponse
	GetBody() *GetVideoProcessingResponseBody
}

type GetVideoProcessingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoProcessingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoProcessingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoProcessingResponse) GoString() string {
	return s.String()
}

func (s *GetVideoProcessingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoProcessingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoProcessingResponse) GetBody() *GetVideoProcessingResponseBody {
	return s.Body
}

func (s *GetVideoProcessingResponse) SetHeaders(v map[string]*string) *GetVideoProcessingResponse {
	s.Headers = v
	return s
}

func (s *GetVideoProcessingResponse) SetStatusCode(v int32) *GetVideoProcessingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoProcessingResponse) SetBody(v *GetVideoProcessingResponseBody) *GetVideoProcessingResponse {
	s.Body = v
	return s
}

func (s *GetVideoProcessingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
