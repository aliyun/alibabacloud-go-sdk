// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoInfosResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoInfosResponseBody) *GetVideoInfosResponse
	GetBody() *GetVideoInfosResponseBody
}

type GetVideoInfosResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoInfosResponse) GoString() string {
	return s.String()
}

func (s *GetVideoInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoInfosResponse) GetBody() *GetVideoInfosResponseBody {
	return s.Body
}

func (s *GetVideoInfosResponse) SetHeaders(v map[string]*string) *GetVideoInfosResponse {
	s.Headers = v
	return s
}

func (s *GetVideoInfosResponse) SetStatusCode(v int32) *GetVideoInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoInfosResponse) SetBody(v *GetVideoInfosResponseBody) *GetVideoInfosResponse {
	s.Body = v
	return s
}

func (s *GetVideoInfosResponse) Validate() error {
	return dara.Validate(s)
}
