// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoPlayInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoPlayInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoPlayInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoPlayInfoResponseBody) *GetVideoPlayInfoResponse
	GetBody() *GetVideoPlayInfoResponseBody
}

type GetVideoPlayInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoPlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoPlayInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPlayInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVideoPlayInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoPlayInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoPlayInfoResponse) GetBody() *GetVideoPlayInfoResponseBody {
	return s.Body
}

func (s *GetVideoPlayInfoResponse) SetHeaders(v map[string]*string) *GetVideoPlayInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVideoPlayInfoResponse) SetStatusCode(v int32) *GetVideoPlayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoPlayInfoResponse) SetBody(v *GetVideoPlayInfoResponseBody) *GetVideoPlayInfoResponse {
	s.Body = v
	return s
}

func (s *GetVideoPlayInfoResponse) Validate() error {
	return dara.Validate(s)
}
