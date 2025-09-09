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
	SetRegionId(v string) *AcceptOperationTicketRequest
	GetRegionId() *string
}

type AcceptOperationTicketRequest struct {
	// The review description.
	//
	// example:
	//
	// O\\&M allowed
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The maximum number of logons allowed. Valid values:
	//
	// 	- **0**: The number of logons is unlimited. The O\\&M engineer can log on to the specified asset for unlimited times during the validity period.
	//
	// 	- **1**: The O\\&M engineer can log on to the specified asset only once during the validity period.
	//
	// > 	- You can set this parameter only to 0 if you review an O\\&M application on a database.
	//
	// > 	- If you do not specify this parameter, the default value 0 is used.
	//
	// example:
	//
	// 1
	EffectCount *string `json:"EffectCount,omitempty" xml:"EffectCount,omitempty"`
	// The end time of the validity period. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1679393152
	EffectEndTime *string `json:"EffectEndTime,omitempty" xml:"EffectEndTime,omitempty"`
	// The start time of the validity period. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1685600242
	EffectStartTime *string `json:"EffectStartTime,omitempty" xml:"EffectStartTime,omitempty"`
	// The ID of the bastion host.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the O\\&M application that you want to approve. You can call the ListOperationTickets operation to query the IDs of all O\\&M applications that require review.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	OperationTicketId *string `json:"OperationTicketId,omitempty" xml:"OperationTicketId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *AcceptOperationTicketRequest) SetRegionId(v string) *AcceptOperationTicketRequest {
	s.RegionId = &v
	return s
}

func (s *AcceptOperationTicketRequest) Validate() error {
	return dara.Validate(s)
}
