// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageGroupUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMessageGroupUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMessageGroupUserResponse
	GetStatusCode() *int32
	SetBody(v *ListMessageGroupUserResponseBody) *ListMessageGroupUserResponse
	GetBody() *ListMessageGroupUserResponseBody
}

type ListMessageGroupUserResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessageGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessageGroupUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupUserResponse) GoString() string {
	return s.String()
}

func (s *ListMessageGroupUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMessageGroupUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMessageGroupUserResponse) GetBody() *ListMessageGroupUserResponseBody {
	return s.Body
}

func (s *ListMessageGroupUserResponse) SetHeaders(v map[string]*string) *ListMessageGroupUserResponse {
	s.Headers = v
	return s
}

func (s *ListMessageGroupUserResponse) SetStatusCode(v int32) *ListMessageGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageGroupUserResponse) SetBody(v *ListMessageGroupUserResponseBody) *ListMessageGroupUserResponse {
	s.Body = v
	return s
}

func (s *ListMessageGroupUserResponse) Validate() error {
	return dara.Validate(s)
}
