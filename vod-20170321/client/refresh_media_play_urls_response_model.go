// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshMediaPlayUrlsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshMediaPlayUrlsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshMediaPlayUrlsResponse
	GetStatusCode() *int32
	SetBody(v *RefreshMediaPlayUrlsResponseBody) *RefreshMediaPlayUrlsResponse
	GetBody() *RefreshMediaPlayUrlsResponseBody
}

type RefreshMediaPlayUrlsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshMediaPlayUrlsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshMediaPlayUrlsResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshMediaPlayUrlsResponse) GoString() string {
	return s.String()
}

func (s *RefreshMediaPlayUrlsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshMediaPlayUrlsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshMediaPlayUrlsResponse) GetBody() *RefreshMediaPlayUrlsResponseBody {
	return s.Body
}

func (s *RefreshMediaPlayUrlsResponse) SetHeaders(v map[string]*string) *RefreshMediaPlayUrlsResponse {
	s.Headers = v
	return s
}

func (s *RefreshMediaPlayUrlsResponse) SetStatusCode(v int32) *RefreshMediaPlayUrlsResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshMediaPlayUrlsResponse) SetBody(v *RefreshMediaPlayUrlsResponseBody) *RefreshMediaPlayUrlsResponse {
	s.Body = v
	return s
}

func (s *RefreshMediaPlayUrlsResponse) Validate() error {
	return dara.Validate(s)
}
