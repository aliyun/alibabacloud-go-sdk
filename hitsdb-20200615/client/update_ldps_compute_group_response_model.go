// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLdpsComputeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLdpsComputeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLdpsComputeGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLdpsComputeGroupResponseBody) *UpdateLdpsComputeGroupResponse
	GetBody() *UpdateLdpsComputeGroupResponseBody
}

type UpdateLdpsComputeGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLdpsComputeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLdpsComputeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLdpsComputeGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateLdpsComputeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLdpsComputeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLdpsComputeGroupResponse) GetBody() *UpdateLdpsComputeGroupResponseBody {
	return s.Body
}

func (s *UpdateLdpsComputeGroupResponse) SetHeaders(v map[string]*string) *UpdateLdpsComputeGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateLdpsComputeGroupResponse) SetStatusCode(v int32) *UpdateLdpsComputeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLdpsComputeGroupResponse) SetBody(v *UpdateLdpsComputeGroupResponseBody) *UpdateLdpsComputeGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateLdpsComputeGroupResponse) Validate() error {
	return dara.Validate(s)
}
