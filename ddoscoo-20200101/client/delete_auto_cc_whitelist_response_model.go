// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoCcWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAutoCcWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAutoCcWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAutoCcWhitelistResponseBody) *DeleteAutoCcWhitelistResponse
	GetBody() *DeleteAutoCcWhitelistResponseBody
}

type DeleteAutoCcWhitelistResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutoCcWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutoCcWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoCcWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoCcWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAutoCcWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAutoCcWhitelistResponse) GetBody() *DeleteAutoCcWhitelistResponseBody {
	return s.Body
}

func (s *DeleteAutoCcWhitelistResponse) SetHeaders(v map[string]*string) *DeleteAutoCcWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoCcWhitelistResponse) SetStatusCode(v int32) *DeleteAutoCcWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoCcWhitelistResponse) SetBody(v *DeleteAutoCcWhitelistResponseBody) *DeleteAutoCcWhitelistResponse {
	s.Body = v
	return s
}

func (s *DeleteAutoCcWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
