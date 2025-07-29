// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadSchedulerxDesignateInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ReadSchedulerxDesignateInfoRequest
	GetGroupId() *string
	SetJobId(v int64) *ReadSchedulerxDesignateInfoRequest
	GetJobId() *int64
	SetNamespace(v string) *ReadSchedulerxDesignateInfoRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *ReadSchedulerxDesignateInfoRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *ReadSchedulerxDesignateInfoRequest
	GetRegionId() *string
}

type ReadSchedulerxDesignateInfoRequest struct {
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

func (s ReadSchedulerxDesignateInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateInfoRequest) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateInfoRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ReadSchedulerxDesignateInfoRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ReadSchedulerxDesignateInfoRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ReadSchedulerxDesignateInfoRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *ReadSchedulerxDesignateInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReadSchedulerxDesignateInfoRequest) SetGroupId(v string) *ReadSchedulerxDesignateInfoRequest {
	s.GroupId = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoRequest) SetJobId(v int64) *ReadSchedulerxDesignateInfoRequest {
	s.JobId = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoRequest) SetNamespace(v string) *ReadSchedulerxDesignateInfoRequest {
	s.Namespace = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoRequest) SetNamespaceSource(v string) *ReadSchedulerxDesignateInfoRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoRequest) SetRegionId(v string) *ReadSchedulerxDesignateInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoRequest) Validate() error {
	return dara.Validate(s)
}
