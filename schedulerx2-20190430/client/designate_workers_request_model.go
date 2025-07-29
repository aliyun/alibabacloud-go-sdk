// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDesignateWorkersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesignateType(v int32) *DesignateWorkersRequest
	GetDesignateType() *int32
	SetGroupId(v string) *DesignateWorkersRequest
	GetGroupId() *string
	SetJobId(v int64) *DesignateWorkersRequest
	GetJobId() *int64
	SetLabels(v string) *DesignateWorkersRequest
	GetLabels() *string
	SetNamespace(v string) *DesignateWorkersRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *DesignateWorkersRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *DesignateWorkersRequest
	GetRegionId() *string
	SetTransferable(v bool) *DesignateWorkersRequest
	GetTransferable() *bool
	SetWorkers(v string) *DesignateWorkersRequest
	GetWorkers() *string
}

type DesignateWorkersRequest struct {
	// The type of the machines to be designated. Valid values: 1 and 2. The value 1 specifies the worker type. The value 2 specifies the label type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	// The application group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hxm.test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 144153
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The designated `labels`. Specify the value of the parameter in a `JSON` string.
	//
	// example:
	//
	// ["gray"]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The unique identifier (UID) of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a06d5ea-f576-4326-842c-fb14ea043d8d
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to allow a failover.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Transferable *bool `json:"Transferable,omitempty" xml:"Transferable,omitempty"`
	// The designated machines. Specify the value of the parameter in a JSON string.
	//
	// example:
	//
	// ["127.0.0.1","127.0.0.2"]
	Workers *string `json:"Workers,omitempty" xml:"Workers,omitempty"`
}

func (s DesignateWorkersRequest) String() string {
	return dara.Prettify(s)
}

func (s DesignateWorkersRequest) GoString() string {
	return s.String()
}

func (s *DesignateWorkersRequest) GetDesignateType() *int32 {
	return s.DesignateType
}

func (s *DesignateWorkersRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DesignateWorkersRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *DesignateWorkersRequest) GetLabels() *string {
	return s.Labels
}

func (s *DesignateWorkersRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DesignateWorkersRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *DesignateWorkersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DesignateWorkersRequest) GetTransferable() *bool {
	return s.Transferable
}

func (s *DesignateWorkersRequest) GetWorkers() *string {
	return s.Workers
}

func (s *DesignateWorkersRequest) SetDesignateType(v int32) *DesignateWorkersRequest {
	s.DesignateType = &v
	return s
}

func (s *DesignateWorkersRequest) SetGroupId(v string) *DesignateWorkersRequest {
	s.GroupId = &v
	return s
}

func (s *DesignateWorkersRequest) SetJobId(v int64) *DesignateWorkersRequest {
	s.JobId = &v
	return s
}

func (s *DesignateWorkersRequest) SetLabels(v string) *DesignateWorkersRequest {
	s.Labels = &v
	return s
}

func (s *DesignateWorkersRequest) SetNamespace(v string) *DesignateWorkersRequest {
	s.Namespace = &v
	return s
}

func (s *DesignateWorkersRequest) SetNamespaceSource(v string) *DesignateWorkersRequest {
	s.NamespaceSource = &v
	return s
}

func (s *DesignateWorkersRequest) SetRegionId(v string) *DesignateWorkersRequest {
	s.RegionId = &v
	return s
}

func (s *DesignateWorkersRequest) SetTransferable(v bool) *DesignateWorkersRequest {
	s.Transferable = &v
	return s
}

func (s *DesignateWorkersRequest) SetWorkers(v string) *DesignateWorkersRequest {
	s.Workers = &v
	return s
}

func (s *DesignateWorkersRequest) Validate() error {
	return dara.Validate(s)
}
