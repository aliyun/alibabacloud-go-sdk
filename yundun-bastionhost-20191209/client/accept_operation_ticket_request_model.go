// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptOperationTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *AcceptOperationTicketRequest
	GetComment() *string
	SetEffectCount(v string) *AcceptOperationTicketRequest
	GetEffectCount() *string
	SetEffectEndTime(v string) *AcceptOperationTicketRequest
	GetEffectEndTime() *string
	SetEffectStartTime(v string) *AcceptOperationTicketRequest
	GetEffectStartTime() *string
	SetInstanceId(v string) *AcceptOperationTicketRequest
	GetInstanceId() *string
	SetOperationTicketId(v string) *AcceptOperationTicketRequest
	GetOperationTicketId() *string
	SetProjectId(v int64) *AcceptOperationTicketRequest
	GetProjectId() *int64
	SetRegionId(v string) *AcceptOperationTicketRequest
	GetRegionId() *string
}

type AcceptOperationTicketRequest struct {
	// The approval remarks.
	//
	// example:
	//
	// Comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The limit on the number of logons allowed. Valid values:
	//
	// - **0**: No limit on the number of logons. The O&M engineer can log on an unlimited number of times during the validity period.
	//
	// - **1**: The O&M engineer can log on only once during the validity period.
	//
	// > - The logon limit for database asset O&M approval can only be set to no limit on the number of logons in Settings.
	//
	// > - If an empty character string is passed, the default value is no limit on the number of logons.
	//
	// example:
	//
	// 1
	EffectCount *string `json:"EffectCount,omitempty" xml:"EffectCount,omitempty"`
	// The end time of the validity period for logon (in seconds, UNIX timestamp format).
	//
	// example:
	//
	// 1679393152
	EffectEndTime *string `json:"EffectEndTime,omitempty" xml:"EffectEndTime,omitempty"`
	// The start time of the validity period for logon (in seconds, UNIX timestamp format).
	//
	// example:
	//
	// 1685600242
	EffectStartTime *string `json:"EffectStartTime,omitempty" xml:"EffectStartTime,omitempty"`
	// The ID of the bastion host instance.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the O&M request to approve.
	//
	// > You can call the [ListOperationTickets](https://help.aliyun.com/document_detail/2584313.html) operation to query all OperationTicketId values that require approval.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	OperationTicketId *string `json:"OperationTicketId,omitempty" xml:"OperationTicketId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AcceptOperationTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s AcceptOperationTicketRequest) GoString() string {
	return s.String()
}

func (s *AcceptOperationTicketRequest) GetComment() *string {
	return s.Comment
}

func (s *AcceptOperationTicketRequest) GetEffectCount() *string {
	return s.EffectCount
}

func (s *AcceptOperationTicketRequest) GetEffectEndTime() *string {
	return s.EffectEndTime
}

func (s *AcceptOperationTicketRequest) GetEffectStartTime() *string {
	return s.EffectStartTime
}

func (s *AcceptOperationTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AcceptOperationTicketRequest) GetOperationTicketId() *string {
	return s.OperationTicketId
}

func (s *AcceptOperationTicketRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AcceptOperationTicketRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AcceptOperationTicketRequest) SetComment(v string) *AcceptOperationTicketRequest {
	s.Comment = &v
	return s
}

func (s *AcceptOperationTicketRequest) SetEffectCount(v string) *AcceptOperationTicketRequest {
	s.EffectCount = &v
	return s
}

func (s *AcceptOperationTicketRequest) SetEffectEndTime(v string) *AcceptOperationTicketRequest {
	s.EffectEndTime = &v
	return s
}

func (s *AcceptOperationTicketRequest) SetEffectStartTime(v string) *AcceptOperationTicketRequest {
	s.EffectStartTime = &v
	return s
}

func (s *AcceptOperationTicketRequest) SetInstanceId(v string) *AcceptOperationTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *AcceptOperationTicketRequest) SetOperationTicketId(v string) *AcceptOperationTicketRequest {
	s.OperationTicketId = &v
	return s
}

func (s *AcceptOperationTicketRequest) SetProjectId(v int64) *AcceptOperationTicketRequest {
	s.ProjectId = &v
	return s
}

func (s *AcceptOperationTicketRequest) SetRegionId(v string) *AcceptOperationTicketRequest {
	s.RegionId = &v
	return s
}

func (s *AcceptOperationTicketRequest) Validate() error {
	return dara.Validate(s)
}
