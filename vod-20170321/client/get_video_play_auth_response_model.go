// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoPlayAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoPlayAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoPlayAuthResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoPlayAuthResponseBody) *GetVideoPlayAuthResponse
	GetBody() *GetVideoPlayAuthResponseBody
}

type GetVideoPlayAuthResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoPlayAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoPlayAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPlayAuthResponse) GoString() string {
	return s.String()
}

func (s *GetVideoPlayAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoPlayAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoPlayAuthResponse) GetBody() *GetVideoPlayAuthResponseBody {
	return s.Body
}

func (s *GetVideoPlayAuthResponse) SetHeaders(v map[string]*string) *GetVideoPlayAuthResponse {
	s.Headers = v
	return s
}

func (s *GetVideoPlayAuthResponse) SetStatusCode(v int32) *GetVideoPlayAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoPlayAuthResponse) SetBody(v *GetVideoPlayAuthResponseBody) *GetVideoPlayAuthResponse {
	s.Body = v
	return s
}

func (s *GetVideoPlayAuthResponse) Validate() error {
	return dara.Validate(s)
}
