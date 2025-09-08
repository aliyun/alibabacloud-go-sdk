// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBindAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBindAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBindAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBindAccountResponseBody) *DeleteBindAccountResponse
	GetBody() *DeleteBindAccountResponseBody
}

type DeleteBindAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBindAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBindAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBindAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteBindAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBindAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBindAccountResponse) GetBody() *DeleteBindAccountResponseBody {
	return s.Body
}

func (s *DeleteBindAccountResponse) SetHeaders(v map[string]*string) *DeleteBindAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteBindAccountResponse) SetStatusCode(v int32) *DeleteBindAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBindAccountResponse) SetBody(v *DeleteBindAccountResponseBody) *DeleteBindAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteBindAccountResponse) Validate() error {
	return dara.Validate(s)
}
