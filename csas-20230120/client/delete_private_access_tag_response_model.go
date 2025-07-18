// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateAccessTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrivateAccessTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrivateAccessTagResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrivateAccessTagResponseBody) *DeletePrivateAccessTagResponse
	GetBody() *DeletePrivateAccessTagResponseBody
}

type DeletePrivateAccessTagResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivateAccessTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivateAccessTagResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateAccessTagResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrivateAccessTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrivateAccessTagResponse) GetBody() *DeletePrivateAccessTagResponseBody {
	return s.Body
}

func (s *DeletePrivateAccessTagResponse) SetHeaders(v map[string]*string) *DeletePrivateAccessTagResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivateAccessTagResponse) SetStatusCode(v int32) *DeletePrivateAccessTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivateAccessTagResponse) SetBody(v *DeletePrivateAccessTagResponseBody) *DeletePrivateAccessTagResponse {
	s.Body = v
	return s
}

func (s *DeletePrivateAccessTagResponse) Validate() error {
	return dara.Validate(s)
}
