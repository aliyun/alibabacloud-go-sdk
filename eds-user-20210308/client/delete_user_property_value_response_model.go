// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserPropertyValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserPropertyValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserPropertyValueResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserPropertyValueResponseBody) *DeleteUserPropertyValueResponse
	GetBody() *DeleteUserPropertyValueResponseBody
}

type DeleteUserPropertyValueResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserPropertyValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserPropertyValueResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserPropertyValueResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserPropertyValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserPropertyValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserPropertyValueResponse) GetBody() *DeleteUserPropertyValueResponseBody {
	return s.Body
}

func (s *DeleteUserPropertyValueResponse) SetHeaders(v map[string]*string) *DeleteUserPropertyValueResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserPropertyValueResponse) SetStatusCode(v int32) *DeleteUserPropertyValueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserPropertyValueResponse) SetBody(v *DeleteUserPropertyValueResponseBody) *DeleteUserPropertyValueResponse {
	s.Body = v
	return s
}

func (s *DeleteUserPropertyValueResponse) Validate() error {
	return dara.Validate(s)
}
