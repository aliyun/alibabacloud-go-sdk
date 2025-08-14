// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueryVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQueryVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQueryVariableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQueryVariableResponseBody) *DeleteQueryVariableResponse
	GetBody() *DeleteQueryVariableResponseBody
}

type DeleteQueryVariableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQueryVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQueryVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueryVariableResponse) GoString() string {
	return s.String()
}

func (s *DeleteQueryVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQueryVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQueryVariableResponse) GetBody() *DeleteQueryVariableResponseBody {
	return s.Body
}

func (s *DeleteQueryVariableResponse) SetHeaders(v map[string]*string) *DeleteQueryVariableResponse {
	s.Headers = v
	return s
}

func (s *DeleteQueryVariableResponse) SetStatusCode(v int32) *DeleteQueryVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQueryVariableResponse) SetBody(v *DeleteQueryVariableResponseBody) *DeleteQueryVariableResponse {
	s.Body = v
	return s
}

func (s *DeleteQueryVariableResponse) Validate() error {
	return dara.Validate(s)
}
