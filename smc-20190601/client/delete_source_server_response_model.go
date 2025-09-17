// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSourceServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSourceServerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSourceServerResponseBody) *DeleteSourceServerResponse
	GetBody() *DeleteSourceServerResponseBody
}

type DeleteSourceServerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSourceServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSourceServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteSourceServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSourceServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSourceServerResponse) GetBody() *DeleteSourceServerResponseBody {
	return s.Body
}

func (s *DeleteSourceServerResponse) SetHeaders(v map[string]*string) *DeleteSourceServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteSourceServerResponse) SetStatusCode(v int32) *DeleteSourceServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSourceServerResponse) SetBody(v *DeleteSourceServerResponseBody) *DeleteSourceServerResponse {
	s.Body = v
	return s
}

func (s *DeleteSourceServerResponse) Validate() error {
	return dara.Validate(s)
}
