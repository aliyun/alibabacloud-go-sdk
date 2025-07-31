// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUdfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUdfResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUdfResponseBody) *DeleteUdfResponse
	GetBody() *DeleteUdfResponseBody
}

type DeleteUdfResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUdfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUdfResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdfResponse) GoString() string {
	return s.String()
}

func (s *DeleteUdfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUdfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUdfResponse) GetBody() *DeleteUdfResponseBody {
	return s.Body
}

func (s *DeleteUdfResponse) SetHeaders(v map[string]*string) *DeleteUdfResponse {
	s.Headers = v
	return s
}

func (s *DeleteUdfResponse) SetStatusCode(v int32) *DeleteUdfResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUdfResponse) SetBody(v *DeleteUdfResponseBody) *DeleteUdfResponse {
	s.Body = v
	return s
}

func (s *DeleteUdfResponse) Validate() error {
	return dara.Validate(s)
}
