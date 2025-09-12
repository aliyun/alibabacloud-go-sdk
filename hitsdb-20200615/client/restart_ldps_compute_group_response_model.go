// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartLdpsComputeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartLdpsComputeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartLdpsComputeGroupResponse
	GetStatusCode() *int32
	SetBody(v *RestartLdpsComputeGroupResponseBody) *RestartLdpsComputeGroupResponse
	GetBody() *RestartLdpsComputeGroupResponseBody
}

type RestartLdpsComputeGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartLdpsComputeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartLdpsComputeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartLdpsComputeGroupResponse) GoString() string {
	return s.String()
}

func (s *RestartLdpsComputeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartLdpsComputeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartLdpsComputeGroupResponse) GetBody() *RestartLdpsComputeGroupResponseBody {
	return s.Body
}

func (s *RestartLdpsComputeGroupResponse) SetHeaders(v map[string]*string) *RestartLdpsComputeGroupResponse {
	s.Headers = v
	return s
}

func (s *RestartLdpsComputeGroupResponse) SetStatusCode(v int32) *RestartLdpsComputeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartLdpsComputeGroupResponse) SetBody(v *RestartLdpsComputeGroupResponseBody) *RestartLdpsComputeGroupResponse {
	s.Body = v
	return s
}

func (s *RestartLdpsComputeGroupResponse) Validate() error {
	return dara.Validate(s)
}
