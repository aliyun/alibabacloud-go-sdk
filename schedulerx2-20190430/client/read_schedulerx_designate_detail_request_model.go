// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadSchedulerxDesignateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesignateType(v int32) *ReadSchedulerxDesignateDetailRequest
	GetDesignateType() *int32
	SetGroupId(v string) *ReadSchedulerxDesignateDetailRequest
	GetGroupId() *string
	SetJobId(v int64) *ReadSchedulerxDesignateDetailRequest
	GetJobId() *int64
	SetNamespace(v string) *ReadSchedulerxDesignateDetailRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *ReadSchedulerxDesignateDetailRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *ReadSchedulerxDesignateDetailRequest
	GetRegionId() *string
}

type ReadSchedulerxDesignateDetailRequest struct {
	// example:
	//
	// 1
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.defalutGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 368
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReadSchedulerxDesignateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateDetailRequest) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateDetailRequest) GetDesignateType() *int32 {
	return s.DesignateType
}

func (s *ReadSchedulerxDesignateDetailRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ReadSchedulerxDesignateDetailRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ReadSchedulerxDesignateDetailRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ReadSchedulerxDesignateDetailRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *ReadSchedulerxDesignateDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReadSchedulerxDesignateDetailRequest) SetDesignateType(v int32) *ReadSchedulerxDesignateDetailRequest {
	s.DesignateType = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailRequest) SetGroupId(v string) *ReadSchedulerxDesignateDetailRequest {
	s.GroupId = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailRequest) SetJobId(v int64) *ReadSchedulerxDesignateDetailRequest {
	s.JobId = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailRequest) SetNamespace(v string) *ReadSchedulerxDesignateDetailRequest {
	s.Namespace = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailRequest) SetNamespaceSource(v string) *ReadSchedulerxDesignateDetailRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailRequest) SetRegionId(v string) *ReadSchedulerxDesignateDetailRequest {
	s.RegionId = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailRequest) Validate() error {
	return dara.Validate(s)
}
