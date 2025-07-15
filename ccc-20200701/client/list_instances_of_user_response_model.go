// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesOfUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstancesOfUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstancesOfUserResponse
	GetStatusCode() *int32
	SetBody(v *ListInstancesOfUserResponseBody) *ListInstancesOfUserResponse
	GetBody() *ListInstancesOfUserResponseBody
}

type ListInstancesOfUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesOfUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesOfUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesOfUserResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstancesOfUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstancesOfUserResponse) GetBody() *ListInstancesOfUserResponseBody {
	return s.Body
}

func (s *ListInstancesOfUserResponse) SetHeaders(v map[string]*string) *ListInstancesOfUserResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesOfUserResponse) SetStatusCode(v int32) *ListInstancesOfUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesOfUserResponse) SetBody(v *ListInstancesOfUserResponseBody) *ListInstancesOfUserResponse {
	s.Body = v
	return s
}

func (s *ListInstancesOfUserResponse) Validate() error {
	return dara.Validate(s)
}
