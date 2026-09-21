// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentlessTaskCountBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTargetType(v int32) *GetAgentlessTaskCountBatchRequest
	GetTargetType() *int32
	SetUuidList(v []*string) *GetAgentlessTaskCountBatchRequest
	GetUuidList() []*string
}

type GetAgentlessTaskCountBatchRequest struct {
	// The detection object type. Valid values:
	//
	// - **1**: host snapshot
	//
	// - **2**: host image
	//
	// - **3**: user snapshot
	//
	// - **4**: user image
	//
	// - **5**: NAS file system
	//
	// - **6**: parallel sandbox
	//
	// - **7**: security fix
	//
	// example:
	//
	// 3
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The list of resource UUIDs to query. The list can contain 1 to 100 elements.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["3bb30859-b3b5-4f28-868f-b0892c98****"]
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s GetAgentlessTaskCountBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentlessTaskCountBatchRequest) GoString() string {
	return s.String()
}

func (s *GetAgentlessTaskCountBatchRequest) GetTargetType() *int32 {
	return s.TargetType
}

func (s *GetAgentlessTaskCountBatchRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *GetAgentlessTaskCountBatchRequest) SetTargetType(v int32) *GetAgentlessTaskCountBatchRequest {
	s.TargetType = &v
	return s
}

func (s *GetAgentlessTaskCountBatchRequest) SetUuidList(v []*string) *GetAgentlessTaskCountBatchRequest {
	s.UuidList = v
	return s
}

func (s *GetAgentlessTaskCountBatchRequest) Validate() error {
	return dara.Validate(s)
}
