// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLdpsComputeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLdpsComputeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLdpsComputeGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLdpsComputeGroupResponseBody) *DeleteLdpsComputeGroupResponse
	GetBody() *DeleteLdpsComputeGroupResponseBody
}

type DeleteLdpsComputeGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLdpsComputeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLdpsComputeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLdpsComputeGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteLdpsComputeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLdpsComputeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLdpsComputeGroupResponse) GetBody() *DeleteLdpsComputeGroupResponseBody {
	return s.Body
}

func (s *DeleteLdpsComputeGroupResponse) SetHeaders(v map[string]*string) *DeleteLdpsComputeGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteLdpsComputeGroupResponse) SetStatusCode(v int32) *DeleteLdpsComputeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLdpsComputeGroupResponse) SetBody(v *DeleteLdpsComputeGroupResponseBody) *DeleteLdpsComputeGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteLdpsComputeGroupResponse) Validate() error {
	return dara.Validate(s)
}
