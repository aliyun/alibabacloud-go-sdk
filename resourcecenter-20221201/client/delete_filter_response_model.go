// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFilterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFilterResponseBody) *DeleteFilterResponse
	GetBody() *DeleteFilterResponseBody
}

type DeleteFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilterResponse) GoString() string {
	return s.String()
}

func (s *DeleteFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFilterResponse) GetBody() *DeleteFilterResponseBody {
	return s.Body
}

func (s *DeleteFilterResponse) SetHeaders(v map[string]*string) *DeleteFilterResponse {
	s.Headers = v
	return s
}

func (s *DeleteFilterResponse) SetStatusCode(v int32) *DeleteFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFilterResponse) SetBody(v *DeleteFilterResponseBody) *DeleteFilterResponse {
	s.Body = v
	return s
}

func (s *DeleteFilterResponse) Validate() error {
	return dara.Validate(s)
}
