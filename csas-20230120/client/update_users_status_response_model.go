// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUsersStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUsersStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUsersStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUsersStatusResponseBody) *UpdateUsersStatusResponse
	GetBody() *UpdateUsersStatusResponseBody
}

type UpdateUsersStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUsersStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUsersStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUsersStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUsersStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUsersStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUsersStatusResponse) GetBody() *UpdateUsersStatusResponseBody {
	return s.Body
}

func (s *UpdateUsersStatusResponse) SetHeaders(v map[string]*string) *UpdateUsersStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUsersStatusResponse) SetStatusCode(v int32) *UpdateUsersStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUsersStatusResponse) SetBody(v *UpdateUsersStatusResponseBody) *UpdateUsersStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateUsersStatusResponse) Validate() error {
	return dara.Validate(s)
}
