// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaMarksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaMarksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaMarksResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaMarksResponseBody) *DeleteMediaMarksResponse
	GetBody() *DeleteMediaMarksResponseBody
}

type DeleteMediaMarksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaMarksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaMarksResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaMarksResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaMarksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaMarksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaMarksResponse) GetBody() *DeleteMediaMarksResponseBody {
	return s.Body
}

func (s *DeleteMediaMarksResponse) SetHeaders(v map[string]*string) *DeleteMediaMarksResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaMarksResponse) SetStatusCode(v int32) *DeleteMediaMarksResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaMarksResponse) SetBody(v *DeleteMediaMarksResponseBody) *DeleteMediaMarksResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaMarksResponse) Validate() error {
	return dara.Validate(s)
}
