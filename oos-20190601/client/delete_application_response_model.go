// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationResponseBody) *DeleteApplicationResponse
	GetBody() *DeleteApplicationResponseBody
}

type DeleteApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationResponse) GetBody() *DeleteApplicationResponseBody {
	return s.Body
}

func (s *DeleteApplicationResponse) SetHeaders(v map[string]*string) *DeleteApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationResponse) SetStatusCode(v int32) *DeleteApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationResponse) SetBody(v *DeleteApplicationResponseBody) *DeleteApplicationResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationResponse) Validate() error {
	return dara.Validate(s)
}
