// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeywordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKeywordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKeywordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKeywordResponseBody) *DeleteKeywordResponse
	GetBody() *DeleteKeywordResponseBody
}

type DeleteKeywordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeywordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeywordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeywordResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeywordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKeywordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKeywordResponse) GetBody() *DeleteKeywordResponseBody {
	return s.Body
}

func (s *DeleteKeywordResponse) SetHeaders(v map[string]*string) *DeleteKeywordResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeywordResponse) SetStatusCode(v int32) *DeleteKeywordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeywordResponse) SetBody(v *DeleteKeywordResponseBody) *DeleteKeywordResponse {
	s.Body = v
	return s
}

func (s *DeleteKeywordResponse) Validate() error {
	return dara.Validate(s)
}
