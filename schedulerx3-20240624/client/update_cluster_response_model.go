// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClusterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClusterResponseBody) *UpdateClusterResponse
	GetBody() *UpdateClusterResponseBody
}

type UpdateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClusterResponse) GetBody() *UpdateClusterResponseBody {
	return s.Body
}

func (s *UpdateClusterResponse) SetHeaders(v map[string]*string) *UpdateClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterResponse) SetStatusCode(v int32) *UpdateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterResponse) SetBody(v *UpdateClusterResponseBody) *UpdateClusterResponse {
	s.Body = v
	return s
}

func (s *UpdateClusterResponse) Validate() error {
	return dara.Validate(s)
}
