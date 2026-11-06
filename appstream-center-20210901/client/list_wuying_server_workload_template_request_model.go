// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWuyingServerWorkloadTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWuyingServerWorkloadTemplateRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWuyingServerWorkloadTemplateRequest
	GetNextToken() *string
	SetVirtualNodePoolId(v string) *ListWuyingServerWorkloadTemplateRequest
	GetVirtualNodePoolId() *string
}

type ListWuyingServerWorkloadTemplateRequest struct {
	// A compatibility field. The current version does not read this parameter. The operation returns a fixed maximum of the first 100 templates.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A compatibility field. The current version does not read this parameter and does not support token-based pagination.
	//
	// example:
	//
	// unused
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the virtual node pool. This parameter is **required**. The node pool must belong to the current account. You can call [ListWuyingServer](~~ListWuyingServer~~) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vnp-bp1234567890abcde
	VirtualNodePoolId *string `json:"VirtualNodePoolId,omitempty" xml:"VirtualNodePoolId,omitempty"`
}

func (s ListWuyingServerWorkloadTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerWorkloadTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListWuyingServerWorkloadTemplateRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWuyingServerWorkloadTemplateRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWuyingServerWorkloadTemplateRequest) GetVirtualNodePoolId() *string {
	return s.VirtualNodePoolId
}

func (s *ListWuyingServerWorkloadTemplateRequest) SetMaxResults(v int32) *ListWuyingServerWorkloadTemplateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWuyingServerWorkloadTemplateRequest) SetNextToken(v string) *ListWuyingServerWorkloadTemplateRequest {
	s.NextToken = &v
	return s
}

func (s *ListWuyingServerWorkloadTemplateRequest) SetVirtualNodePoolId(v string) *ListWuyingServerWorkloadTemplateRequest {
	s.VirtualNodePoolId = &v
	return s
}

func (s *ListWuyingServerWorkloadTemplateRequest) Validate() error {
	return dara.Validate(s)
}
