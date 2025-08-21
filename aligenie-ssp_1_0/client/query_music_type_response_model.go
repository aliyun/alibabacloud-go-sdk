// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMusicTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMusicTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMusicTypeResponse
	GetStatusCode() *int32
	SetBody(v *QueryMusicTypeResponseBody) *QueryMusicTypeResponse
	GetBody() *QueryMusicTypeResponseBody
}

type QueryMusicTypeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMusicTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMusicTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMusicTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMusicTypeResponse) GetBody() *QueryMusicTypeResponseBody {
	return s.Body
}

func (s *QueryMusicTypeResponse) SetHeaders(v map[string]*string) *QueryMusicTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryMusicTypeResponse) SetStatusCode(v int32) *QueryMusicTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMusicTypeResponse) SetBody(v *QueryMusicTypeResponseBody) *QueryMusicTypeResponse {
	s.Body = v
	return s
}

func (s *QueryMusicTypeResponse) Validate() error {
	return dara.Validate(s)
}
