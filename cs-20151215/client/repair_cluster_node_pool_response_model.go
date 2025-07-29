// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepairClusterNodePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RepairClusterNodePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RepairClusterNodePoolResponse
	GetStatusCode() *int32
	SetBody(v *RepairClusterNodePoolResponseBody) *RepairClusterNodePoolResponse
	GetBody() *RepairClusterNodePoolResponseBody
}

type RepairClusterNodePoolResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RepairClusterNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RepairClusterNodePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s RepairClusterNodePoolResponse) GoString() string {
	return s.String()
}

func (s *RepairClusterNodePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RepairClusterNodePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RepairClusterNodePoolResponse) GetBody() *RepairClusterNodePoolResponseBody {
	return s.Body
}

func (s *RepairClusterNodePoolResponse) SetHeaders(v map[string]*string) *RepairClusterNodePoolResponse {
	s.Headers = v
	return s
}

func (s *RepairClusterNodePoolResponse) SetStatusCode(v int32) *RepairClusterNodePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *RepairClusterNodePoolResponse) SetBody(v *RepairClusterNodePoolResponseBody) *RepairClusterNodePoolResponse {
	s.Body = v
	return s
}

func (s *RepairClusterNodePoolResponse) Validate() error {
	return dara.Validate(s)
}
