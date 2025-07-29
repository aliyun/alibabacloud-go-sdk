// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetAppGroupRequest
	GetGroupId() *string
	SetNamespace(v string) *GetAppGroupRequest
	GetNamespace() *string
	SetRegionId(v string) *GetAppGroupRequest
	GetRegionId() *string
}

type GetAppGroupRequest struct {
	// The ID of the application. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the namespace. You can obtain the ID of the namespace on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAppGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppGroupRequest) GoString() string {
	return s.String()
}

func (s *GetAppGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetAppGroupRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetAppGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAppGroupRequest) SetGroupId(v string) *GetAppGroupRequest {
	s.GroupId = &v
	return s
}

func (s *GetAppGroupRequest) SetNamespace(v string) *GetAppGroupRequest {
	s.Namespace = &v
	return s
}

func (s *GetAppGroupRequest) SetRegionId(v string) *GetAppGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetAppGroupRequest) Validate() error {
	return dara.Validate(s)
}
