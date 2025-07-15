// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoResponseBody) *GetVideoResponse
	GetBody() *GetVideoResponseBody
}

type GetVideoResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoResponse) GoString() string {
	return s.String()
}

func (s *GetVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoResponse) GetBody() *GetVideoResponseBody {
	return s.Body
}

func (s *GetVideoResponse) SetHeaders(v map[string]*string) *GetVideoResponse {
	s.Headers = v
	return s
}

func (s *GetVideoResponse) SetStatusCode(v int32) *GetVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoResponse) SetBody(v *GetVideoResponseBody) *GetVideoResponse {
	s.Body = v
	return s
}

func (s *GetVideoResponse) Validate() error {
	return dara.Validate(s)
}
