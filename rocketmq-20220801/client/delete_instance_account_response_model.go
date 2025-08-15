// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceAccountResponseBody) *DeleteInstanceAccountResponse
	GetBody() *DeleteInstanceAccountResponseBody
}

type DeleteInstanceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceAccountResponse) GetBody() *DeleteInstanceAccountResponseBody {
	return s.Body
}

func (s *DeleteInstanceAccountResponse) SetHeaders(v map[string]*string) *DeleteInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceAccountResponse) SetStatusCode(v int32) *DeleteInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceAccountResponse) SetBody(v *DeleteInstanceAccountResponseBody) *DeleteInstanceAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceAccountResponse) Validate() error {
	return dara.Validate(s)
}
