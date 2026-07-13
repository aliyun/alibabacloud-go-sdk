// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpToolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListMcpToolsRequest
	GetId() *string
	SetInstanceId(v string) *ListMcpToolsRequest
	GetInstanceId() *string
}

type ListMcpToolsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// RUNNING
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListMcpToolsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpToolsRequest) GoString() string {
	return s.String()
}

func (s *ListMcpToolsRequest) GetId() *string {
	return s.Id
}

func (s *ListMcpToolsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMcpToolsRequest) SetId(v string) *ListMcpToolsRequest {
	s.Id = &v
	return s
}

func (s *ListMcpToolsRequest) SetInstanceId(v string) *ListMcpToolsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMcpToolsRequest) Validate() error {
	return dara.Validate(s)
}
