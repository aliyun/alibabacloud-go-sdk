// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDesktopsResponseBody) *DeleteDesktopsResponse
	GetBody() *DeleteDesktopsResponseBody
}

type DeleteDesktopsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDesktopsResponse) GetBody() *DeleteDesktopsResponseBody {
	return s.Body
}

func (s *DeleteDesktopsResponse) SetHeaders(v map[string]*string) *DeleteDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDesktopsResponse) SetStatusCode(v int32) *DeleteDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDesktopsResponse) SetBody(v *DeleteDesktopsResponseBody) *DeleteDesktopsResponse {
	s.Body = v
	return s
}

func (s *DeleteDesktopsResponse) Validate() error {
	return dara.Validate(s)
}
