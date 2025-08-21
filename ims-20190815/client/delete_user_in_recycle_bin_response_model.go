// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserInRecycleBinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserInRecycleBinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserInRecycleBinResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserInRecycleBinResponseBody) *DeleteUserInRecycleBinResponse
	GetBody() *DeleteUserInRecycleBinResponseBody
}

type DeleteUserInRecycleBinResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserInRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserInRecycleBinResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserInRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserInRecycleBinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserInRecycleBinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserInRecycleBinResponse) GetBody() *DeleteUserInRecycleBinResponseBody {
	return s.Body
}

func (s *DeleteUserInRecycleBinResponse) SetHeaders(v map[string]*string) *DeleteUserInRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserInRecycleBinResponse) SetStatusCode(v int32) *DeleteUserInRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserInRecycleBinResponse) SetBody(v *DeleteUserInRecycleBinResponseBody) *DeleteUserInRecycleBinResponse {
	s.Body = v
	return s
}

func (s *DeleteUserInRecycleBinResponse) Validate() error {
	return dara.Validate(s)
}
