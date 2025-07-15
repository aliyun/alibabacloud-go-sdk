// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpsecServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIpsecServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIpsecServerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIpsecServerResponseBody) *DeleteIpsecServerResponse
	GetBody() *DeleteIpsecServerResponseBody
}

type DeleteIpsecServerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIpsecServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIpsecServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpsecServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpsecServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIpsecServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIpsecServerResponse) GetBody() *DeleteIpsecServerResponseBody {
	return s.Body
}

func (s *DeleteIpsecServerResponse) SetHeaders(v map[string]*string) *DeleteIpsecServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpsecServerResponse) SetStatusCode(v int32) *DeleteIpsecServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpsecServerResponse) SetBody(v *DeleteIpsecServerResponseBody) *DeleteIpsecServerResponse {
	s.Body = v
	return s
}

func (s *DeleteIpsecServerResponse) Validate() error {
	return dara.Validate(s)
}
