// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnsSaleControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnsSaleControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnsSaleControlResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEnsSaleControlResponseBody) *DeleteEnsSaleControlResponse
	GetBody() *DeleteEnsSaleControlResponseBody
}

type DeleteEnsSaleControlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnsSaleControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnsSaleControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsSaleControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnsSaleControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnsSaleControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnsSaleControlResponse) GetBody() *DeleteEnsSaleControlResponseBody {
	return s.Body
}

func (s *DeleteEnsSaleControlResponse) SetHeaders(v map[string]*string) *DeleteEnsSaleControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnsSaleControlResponse) SetStatusCode(v int32) *DeleteEnsSaleControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnsSaleControlResponse) SetBody(v *DeleteEnsSaleControlResponseBody) *DeleteEnsSaleControlResponse {
	s.Body = v
	return s
}

func (s *DeleteEnsSaleControlResponse) Validate() error {
	return dara.Validate(s)
}
