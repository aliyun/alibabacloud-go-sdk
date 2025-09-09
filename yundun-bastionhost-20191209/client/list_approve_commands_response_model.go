// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApproveCommandsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApproveCommandsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApproveCommandsResponse
	GetStatusCode() *int32
	SetBody(v *ListApproveCommandsResponseBody) *ListApproveCommandsResponse
	GetBody() *ListApproveCommandsResponseBody
}

type ListApproveCommandsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApproveCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApproveCommandsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApproveCommandsResponse) GoString() string {
	return s.String()
}

func (s *ListApproveCommandsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApproveCommandsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApproveCommandsResponse) GetBody() *ListApproveCommandsResponseBody {
	return s.Body
}

func (s *ListApproveCommandsResponse) SetHeaders(v map[string]*string) *ListApproveCommandsResponse {
	s.Headers = v
	return s
}

func (s *ListApproveCommandsResponse) SetStatusCode(v int32) *ListApproveCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApproveCommandsResponse) SetBody(v *ListApproveCommandsResponseBody) *ListApproveCommandsResponse {
	s.Body = v
	return s
}

func (s *ListApproveCommandsResponse) Validate() error {
	return dara.Validate(s)
}
