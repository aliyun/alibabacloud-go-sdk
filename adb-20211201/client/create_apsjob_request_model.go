// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAPSJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApsJobName(v string) *CreateAPSJobRequest
	GetApsJobName() *string
	SetDbList(v string) *CreateAPSJobRequest
	GetDbList() *string
	SetDestinationEndpointInstanceID(v string) *CreateAPSJobRequest
	GetDestinationEndpointInstanceID() *string
	SetDestinationEndpointPassword(v string) *CreateAPSJobRequest
	GetDestinationEndpointPassword() *string
	SetDestinationEndpointUserName(v string) *CreateAPSJobRequest
	GetDestinationEndpointUserName() *string
	SetPartitionList(v string) *CreateAPSJobRequest
	GetPartitionList() *string
	SetRegionId(v string) *CreateAPSJobRequest
	GetRegionId() *string
	SetSourceEndpointInstanceID(v string) *CreateAPSJobRequest
	GetSourceEndpointInstanceID() *string
	SetSourceEndpointPassword(v string) *CreateAPSJobRequest
	GetSourceEndpointPassword() *string
	SetSourceEndpointRegion(v string) *CreateAPSJobRequest
	GetSourceEndpointRegion() *string
	SetSourceEndpointUserName(v string) *CreateAPSJobRequest
	GetSourceEndpointUserName() *string
	SetTargetTableMode(v string) *CreateAPSJobRequest
	GetTargetTableMode() *string
}

type CreateAPSJobRequest struct {
	// The name of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// aps-xxxxx
	ApsJobName *string `json:"ApsJobName,omitempty" xml:"ApsJobName,omitempty"`
	// The list of synchronization objects.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"EntireInstance":true}
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The instance ID of the destination database.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-xxxxx
	DestinationEndpointInstanceID *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	// The password of the destination instance.
	//
	// example:
	//
	// ******
	DestinationEndpointPassword *string `json:"DestinationEndpointPassword,omitempty" xml:"DestinationEndpointPassword,omitempty"`
	// The username of the destination instance.
	//
	// example:
	//
	// ******
	DestinationEndpointUserName *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	// The list of partitions.
	//
	// example:
	//
	// {}
	PartitionList *string `json:"PartitionList,omitempty" xml:"PartitionList,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the source database.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-xxxxx
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The password of the source endpoint.
	//
	// example:
	//
	// ******
	SourceEndpointPassword *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	// The region of the source database.
	//
	// example:
	//
	// cn-beijing
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The username of the source endpoint.
	//
	// example:
	//
	// xxxx
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	// The destination table mode.
	//
	// example:
	//
	// 1
	TargetTableMode *string `json:"TargetTableMode,omitempty" xml:"TargetTableMode,omitempty"`
}

func (s CreateAPSJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAPSJobRequest) GoString() string {
	return s.String()
}

func (s *CreateAPSJobRequest) GetApsJobName() *string {
	return s.ApsJobName
}

func (s *CreateAPSJobRequest) GetDbList() *string {
	return s.DbList
}

func (s *CreateAPSJobRequest) GetDestinationEndpointInstanceID() *string {
	return s.DestinationEndpointInstanceID
}

func (s *CreateAPSJobRequest) GetDestinationEndpointPassword() *string {
	return s.DestinationEndpointPassword
}

func (s *CreateAPSJobRequest) GetDestinationEndpointUserName() *string {
	return s.DestinationEndpointUserName
}

func (s *CreateAPSJobRequest) GetPartitionList() *string {
	return s.PartitionList
}

func (s *CreateAPSJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAPSJobRequest) GetSourceEndpointInstanceID() *string {
	return s.SourceEndpointInstanceID
}

func (s *CreateAPSJobRequest) GetSourceEndpointPassword() *string {
	return s.SourceEndpointPassword
}

func (s *CreateAPSJobRequest) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *CreateAPSJobRequest) GetSourceEndpointUserName() *string {
	return s.SourceEndpointUserName
}

func (s *CreateAPSJobRequest) GetTargetTableMode() *string {
	return s.TargetTableMode
}

func (s *CreateAPSJobRequest) SetApsJobName(v string) *CreateAPSJobRequest {
	s.ApsJobName = &v
	return s
}

func (s *CreateAPSJobRequest) SetDbList(v string) *CreateAPSJobRequest {
	s.DbList = &v
	return s
}

func (s *CreateAPSJobRequest) SetDestinationEndpointInstanceID(v string) *CreateAPSJobRequest {
	s.DestinationEndpointInstanceID = &v
	return s
}

func (s *CreateAPSJobRequest) SetDestinationEndpointPassword(v string) *CreateAPSJobRequest {
	s.DestinationEndpointPassword = &v
	return s
}

func (s *CreateAPSJobRequest) SetDestinationEndpointUserName(v string) *CreateAPSJobRequest {
	s.DestinationEndpointUserName = &v
	return s
}

func (s *CreateAPSJobRequest) SetPartitionList(v string) *CreateAPSJobRequest {
	s.PartitionList = &v
	return s
}

func (s *CreateAPSJobRequest) SetRegionId(v string) *CreateAPSJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAPSJobRequest) SetSourceEndpointInstanceID(v string) *CreateAPSJobRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *CreateAPSJobRequest) SetSourceEndpointPassword(v string) *CreateAPSJobRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *CreateAPSJobRequest) SetSourceEndpointRegion(v string) *CreateAPSJobRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *CreateAPSJobRequest) SetSourceEndpointUserName(v string) *CreateAPSJobRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *CreateAPSJobRequest) SetTargetTableMode(v string) *CreateAPSJobRequest {
	s.TargetTableMode = &v
	return s
}

func (s *CreateAPSJobRequest) Validate() error {
	return dara.Validate(s)
}
