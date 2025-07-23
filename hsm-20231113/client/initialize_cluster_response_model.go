// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitializeClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitializeClusterResponse
	GetStatusCode() *int32
	SetBody(v *InitializeClusterResponseBody) *InitializeClusterResponse
	GetBody() *InitializeClusterResponseBody
}

type InitializeClusterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s InitializeClusterResponse) GoString() string {
	return s.String()
}

func (s *InitializeClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitializeClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitializeClusterResponse) GetBody() *InitializeClusterResponseBody {
	return s.Body
}

func (s *InitializeClusterResponse) SetHeaders(v map[string]*string) *InitializeClusterResponse {
	s.Headers = v
	return s
}

func (s *InitializeClusterResponse) SetStatusCode(v int32) *InitializeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeClusterResponse) SetBody(v *InitializeClusterResponseBody) *InitializeClusterResponse {
	s.Body = v
	return s
}

func (s *InitializeClusterResponse) Validate() error {
	return dara.Validate(s)
}
