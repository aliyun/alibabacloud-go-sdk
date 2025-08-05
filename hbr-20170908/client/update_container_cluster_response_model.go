// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContainerClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContainerClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContainerClusterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContainerClusterResponseBody) *UpdateContainerClusterResponse
	GetBody() *UpdateContainerClusterResponseBody
}

type UpdateContainerClusterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContainerClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContainerClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContainerClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateContainerClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContainerClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContainerClusterResponse) GetBody() *UpdateContainerClusterResponseBody {
	return s.Body
}

func (s *UpdateContainerClusterResponse) SetHeaders(v map[string]*string) *UpdateContainerClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateContainerClusterResponse) SetStatusCode(v int32) *UpdateContainerClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContainerClusterResponse) SetBody(v *UpdateContainerClusterResponseBody) *UpdateContainerClusterResponse {
	s.Body = v
	return s
}

func (s *UpdateContainerClusterResponse) Validate() error {
	return dara.Validate(s)
}
