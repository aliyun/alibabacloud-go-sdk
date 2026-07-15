// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteRouteStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *BatchDeleteRouteStrategyRequest
	GetGroupId() *string
	SetJobIdList(v []*int64) *BatchDeleteRouteStrategyRequest
	GetJobIdList() []*int64
	SetNamespace(v string) *BatchDeleteRouteStrategyRequest
	GetNamespace() *string
	SetRegionId(v string) *BatchDeleteRouteStrategyRequest
	GetRegionId() *string
}

type BatchDeleteRouteStrategyRequest struct {
	// The ID of the Application. You can find the ID on the **Application Management*	- page in the Console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// A list of Job IDs.
	JobIdList []*int64 `json:"JobIdList,omitempty" xml:"JobIdList,omitempty" type:"Repeated"`
	// The ID of the Namespace. You can find the ID on the **Namespaces*	- page in the Console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the Region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchDeleteRouteStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteRouteStrategyRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteRouteStrategyRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *BatchDeleteRouteStrategyRequest) GetJobIdList() []*int64 {
	return s.JobIdList
}

func (s *BatchDeleteRouteStrategyRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDeleteRouteStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchDeleteRouteStrategyRequest) SetGroupId(v string) *BatchDeleteRouteStrategyRequest {
	s.GroupId = &v
	return s
}

func (s *BatchDeleteRouteStrategyRequest) SetJobIdList(v []*int64) *BatchDeleteRouteStrategyRequest {
	s.JobIdList = v
	return s
}

func (s *BatchDeleteRouteStrategyRequest) SetNamespace(v string) *BatchDeleteRouteStrategyRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteRouteStrategyRequest) SetRegionId(v string) *BatchDeleteRouteStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *BatchDeleteRouteStrategyRequest) Validate() error {
	return dara.Validate(s)
}
