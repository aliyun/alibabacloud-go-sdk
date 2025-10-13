// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupMachineGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceGroupMachineGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceGroupMachineGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceGroupMachineGroupsResponseBody) *ListResourceGroupMachineGroupsResponse
	GetBody() *ListResourceGroupMachineGroupsResponseBody
}

type ListResourceGroupMachineGroupsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceGroupMachineGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceGroupMachineGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupMachineGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupMachineGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceGroupMachineGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceGroupMachineGroupsResponse) GetBody() *ListResourceGroupMachineGroupsResponseBody {
	return s.Body
}

func (s *ListResourceGroupMachineGroupsResponse) SetHeaders(v map[string]*string) *ListResourceGroupMachineGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupMachineGroupsResponse) SetStatusCode(v int32) *ListResourceGroupMachineGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceGroupMachineGroupsResponse) SetBody(v *ListResourceGroupMachineGroupsResponseBody) *ListResourceGroupMachineGroupsResponse {
	s.Body = v
	return s
}

func (s *ListResourceGroupMachineGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
