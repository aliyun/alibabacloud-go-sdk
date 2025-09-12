// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLdpsComputeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLdpsComputeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLdpsComputeGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetLdpsComputeGroupResponseBody) *GetLdpsComputeGroupResponse
	GetBody() *GetLdpsComputeGroupResponseBody
}

type GetLdpsComputeGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLdpsComputeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLdpsComputeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLdpsComputeGroupResponse) GoString() string {
	return s.String()
}

func (s *GetLdpsComputeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLdpsComputeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLdpsComputeGroupResponse) GetBody() *GetLdpsComputeGroupResponseBody {
	return s.Body
}

func (s *GetLdpsComputeGroupResponse) SetHeaders(v map[string]*string) *GetLdpsComputeGroupResponse {
	s.Headers = v
	return s
}

func (s *GetLdpsComputeGroupResponse) SetStatusCode(v int32) *GetLdpsComputeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLdpsComputeGroupResponse) SetBody(v *GetLdpsComputeGroupResponseBody) *GetLdpsComputeGroupResponse {
	s.Body = v
	return s
}

func (s *GetLdpsComputeGroupResponse) Validate() error {
	return dara.Validate(s)
}
