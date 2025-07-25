// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceGroupMachineGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceGroupMachineGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceGroupMachineGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceGroupMachineGroupResponseBody) *DeleteResourceGroupMachineGroupResponse
	GetBody() *DeleteResourceGroupMachineGroupResponseBody
}

type DeleteResourceGroupMachineGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceGroupMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceGroupMachineGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupMachineGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceGroupMachineGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceGroupMachineGroupResponse) GetBody() *DeleteResourceGroupMachineGroupResponseBody {
	return s.Body
}

func (s *DeleteResourceGroupMachineGroupResponse) SetHeaders(v map[string]*string) *DeleteResourceGroupMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceGroupMachineGroupResponse) SetStatusCode(v int32) *DeleteResourceGroupMachineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceGroupMachineGroupResponse) SetBody(v *DeleteResourceGroupMachineGroupResponseBody) *DeleteResourceGroupMachineGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceGroupMachineGroupResponse) Validate() error {
	return dara.Validate(s)
}
