// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupMachineGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceGroupMachineGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceGroupMachineGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceGroupMachineGroupResponseBody) *GetResourceGroupMachineGroupResponse
	GetBody() *GetResourceGroupMachineGroupResponseBody
}

type GetResourceGroupMachineGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupMachineGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupMachineGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceGroupMachineGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceGroupMachineGroupResponse) GetBody() *GetResourceGroupMachineGroupResponseBody {
	return s.Body
}

func (s *GetResourceGroupMachineGroupResponse) SetHeaders(v map[string]*string) *GetResourceGroupMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupMachineGroupResponse) SetStatusCode(v int32) *GetResourceGroupMachineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupMachineGroupResponse) SetBody(v *GetResourceGroupMachineGroupResponseBody) *GetResourceGroupMachineGroupResponse {
	s.Body = v
	return s
}

func (s *GetResourceGroupMachineGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
