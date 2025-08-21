// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEpnInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEpnInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEpnInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEpnInstanceResponseBody) *DeleteEpnInstanceResponse
	GetBody() *DeleteEpnInstanceResponseBody
}

type DeleteEpnInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEpnInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEpnInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEpnInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteEpnInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEpnInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEpnInstanceResponse) GetBody() *DeleteEpnInstanceResponseBody {
	return s.Body
}

func (s *DeleteEpnInstanceResponse) SetHeaders(v map[string]*string) *DeleteEpnInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteEpnInstanceResponse) SetStatusCode(v int32) *DeleteEpnInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEpnInstanceResponse) SetBody(v *DeleteEpnInstanceResponseBody) *DeleteEpnInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteEpnInstanceResponse) Validate() error {
	return dara.Validate(s)
}
