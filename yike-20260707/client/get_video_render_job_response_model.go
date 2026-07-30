// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoRenderJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoRenderJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoRenderJobResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoRenderJobResponseBody) *GetVideoRenderJobResponse
	GetBody() *GetVideoRenderJobResponseBody
}

type GetVideoRenderJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoRenderJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoRenderJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoRenderJobResponse) GoString() string {
	return s.String()
}

func (s *GetVideoRenderJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoRenderJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoRenderJobResponse) GetBody() *GetVideoRenderJobResponseBody {
	return s.Body
}

func (s *GetVideoRenderJobResponse) SetHeaders(v map[string]*string) *GetVideoRenderJobResponse {
	s.Headers = v
	return s
}

func (s *GetVideoRenderJobResponse) SetStatusCode(v int32) *GetVideoRenderJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoRenderJobResponse) SetBody(v *GetVideoRenderJobResponseBody) *GetVideoRenderJobResponse {
	s.Body = v
	return s
}

func (s *GetVideoRenderJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
