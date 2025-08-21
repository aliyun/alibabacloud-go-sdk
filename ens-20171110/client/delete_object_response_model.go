// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteObjectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteObjectResponseBody) *DeleteObjectResponse
	GetBody() *DeleteObjectResponseBody
}

type DeleteObjectResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteObjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteObjectResponse) GetBody() *DeleteObjectResponseBody {
	return s.Body
}

func (s *DeleteObjectResponse) SetHeaders(v map[string]*string) *DeleteObjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteObjectResponse) SetStatusCode(v int32) *DeleteObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteObjectResponse) SetBody(v *DeleteObjectResponseBody) *DeleteObjectResponse {
	s.Body = v
	return s
}

func (s *DeleteObjectResponse) Validate() error {
	return dara.Validate(s)
}
