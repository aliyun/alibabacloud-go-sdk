// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccessKeyResponseBody) *DeleteAccessKeyResponse
	GetBody() *DeleteAccessKeyResponseBody
}

type DeleteAccessKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessKeyResponse) GetBody() *DeleteAccessKeyResponseBody {
	return s.Body
}

func (s *DeleteAccessKeyResponse) SetHeaders(v map[string]*string) *DeleteAccessKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessKeyResponse) SetStatusCode(v int32) *DeleteAccessKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessKeyResponse) SetBody(v *DeleteAccessKeyResponseBody) *DeleteAccessKeyResponse {
	s.Body = v
	return s
}

func (s *DeleteAccessKeyResponse) Validate() error {
	return dara.Validate(s)
}
