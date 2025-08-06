// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoConfigResponseBody) *GetVideoConfigResponse
	GetBody() *GetVideoConfigResponseBody
}

type GetVideoConfigResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoConfigResponse) GoString() string {
	return s.String()
}

func (s *GetVideoConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoConfigResponse) GetBody() *GetVideoConfigResponseBody {
	return s.Body
}

func (s *GetVideoConfigResponse) SetHeaders(v map[string]*string) *GetVideoConfigResponse {
	s.Headers = v
	return s
}

func (s *GetVideoConfigResponse) SetStatusCode(v int32) *GetVideoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoConfigResponse) SetBody(v *GetVideoConfigResponseBody) *GetVideoConfigResponse {
	s.Body = v
	return s
}

func (s *GetVideoConfigResponse) Validate() error {
	return dara.Validate(s)
}
