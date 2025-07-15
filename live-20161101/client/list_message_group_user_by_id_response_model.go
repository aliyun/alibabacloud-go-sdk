// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageGroupUserByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMessageGroupUserByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMessageGroupUserByIdResponse
	GetStatusCode() *int32
	SetBody(v *ListMessageGroupUserByIdResponseBody) *ListMessageGroupUserByIdResponse
	GetBody() *ListMessageGroupUserByIdResponseBody
}

type ListMessageGroupUserByIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessageGroupUserByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessageGroupUserByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMessageGroupUserByIdResponse) GoString() string {
	return s.String()
}

func (s *ListMessageGroupUserByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMessageGroupUserByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMessageGroupUserByIdResponse) GetBody() *ListMessageGroupUserByIdResponseBody {
	return s.Body
}

func (s *ListMessageGroupUserByIdResponse) SetHeaders(v map[string]*string) *ListMessageGroupUserByIdResponse {
	s.Headers = v
	return s
}

func (s *ListMessageGroupUserByIdResponse) SetStatusCode(v int32) *ListMessageGroupUserByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageGroupUserByIdResponse) SetBody(v *ListMessageGroupUserByIdResponseBody) *ListMessageGroupUserByIdResponse {
	s.Body = v
	return s
}

func (s *ListMessageGroupUserByIdResponse) Validate() error {
	return dara.Validate(s)
}
