// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClusterConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClusterConfigsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClusterConfigsResponseBody) *UpdateClusterConfigsResponse
	GetBody() *UpdateClusterConfigsResponseBody
}

type UpdateClusterConfigsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClusterConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClusterConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterConfigsResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClusterConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClusterConfigsResponse) GetBody() *UpdateClusterConfigsResponseBody {
	return s.Body
}

func (s *UpdateClusterConfigsResponse) SetHeaders(v map[string]*string) *UpdateClusterConfigsResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterConfigsResponse) SetStatusCode(v int32) *UpdateClusterConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClusterConfigsResponse) SetBody(v *UpdateClusterConfigsResponseBody) *UpdateClusterConfigsResponse {
	s.Body = v
	return s
}

func (s *UpdateClusterConfigsResponse) Validate() error {
	return dara.Validate(s)
}
