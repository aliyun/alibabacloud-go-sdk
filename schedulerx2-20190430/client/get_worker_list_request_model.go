// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetWorkerListRequest
	GetGroupId() *string
	SetNamespace(v string) *GetWorkerListRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *GetWorkerListRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *GetWorkerListRequest
	GetRegionId() *string
}

type GetWorkerListRequest struct {
	// The ID of the permission group.
	//
	// This parameter is required.
	//
	// example:
	//
	// usercenter
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the namespace. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetWorkerListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerListRequest) GoString() string {
	return s.String()
}

func (s *GetWorkerListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetWorkerListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetWorkerListRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *GetWorkerListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWorkerListRequest) SetGroupId(v string) *GetWorkerListRequest {
	s.GroupId = &v
	return s
}

func (s *GetWorkerListRequest) SetNamespace(v string) *GetWorkerListRequest {
	s.Namespace = &v
	return s
}

func (s *GetWorkerListRequest) SetNamespaceSource(v string) *GetWorkerListRequest {
	s.NamespaceSource = &v
	return s
}

func (s *GetWorkerListRequest) SetRegionId(v string) *GetWorkerListRequest {
	s.RegionId = &v
	return s
}

func (s *GetWorkerListRequest) Validate() error {
	return dara.Validate(s)
}
