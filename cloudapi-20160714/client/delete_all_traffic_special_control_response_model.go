// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAllTrafficSpecialControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAllTrafficSpecialControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAllTrafficSpecialControlResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAllTrafficSpecialControlResponseBody) *DeleteAllTrafficSpecialControlResponse
	GetBody() *DeleteAllTrafficSpecialControlResponseBody
}

type DeleteAllTrafficSpecialControlResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAllTrafficSpecialControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAllTrafficSpecialControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAllTrafficSpecialControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteAllTrafficSpecialControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAllTrafficSpecialControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAllTrafficSpecialControlResponse) GetBody() *DeleteAllTrafficSpecialControlResponseBody {
	return s.Body
}

func (s *DeleteAllTrafficSpecialControlResponse) SetHeaders(v map[string]*string) *DeleteAllTrafficSpecialControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteAllTrafficSpecialControlResponse) SetStatusCode(v int32) *DeleteAllTrafficSpecialControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAllTrafficSpecialControlResponse) SetBody(v *DeleteAllTrafficSpecialControlResponseBody) *DeleteAllTrafficSpecialControlResponse {
	s.Body = v
	return s
}

func (s *DeleteAllTrafficSpecialControlResponse) Validate() error {
	return dara.Validate(s)
}
