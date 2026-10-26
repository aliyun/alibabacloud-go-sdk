// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProjectNotifyReceiverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *ModifyProjectNotifyReceiverRequest
	GetEmail() *string
	SetInstanceId(v string) *ModifyProjectNotifyReceiverRequest
	GetInstanceId() *string
	SetName(v string) *ModifyProjectNotifyReceiverRequest
	GetName() *string
	SetPhone(v string) *ModifyProjectNotifyReceiverRequest
	GetPhone() *string
	SetReceiverId(v int64) *ModifyProjectNotifyReceiverRequest
	GetReceiverId() *int64
	SetRegionId(v string) *ModifyProjectNotifyReceiverRequest
	GetRegionId() *string
}

type ModifyProjectNotifyReceiverRequest struct {
	// The new email address of the notification recipient.
	//
	// example:
	//
	// 1824***@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the bastion host instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host instance ID.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new name of the notification recipient. The name must be 1 to 128 characters in length, cannot start with a special character, and can contain only the following special characters: periods (.), underscores (_), hyphens (-), backslashes (\\), and spaces.
	//
	// > If this parameter is not specified, the name remains unchanged.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The new phone number of the notification recipient.
	//
	// example:
	//
	// 171****5421
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The ID of the notification recipient.
	//
	// > You can call the ListProjectNotifyReceivers operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReceiverId *int64 `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyProjectNotifyReceiverRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyProjectNotifyReceiverRequest) GoString() string {
	return s.String()
}

func (s *ModifyProjectNotifyReceiverRequest) GetEmail() *string {
	return s.Email
}

func (s *ModifyProjectNotifyReceiverRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyProjectNotifyReceiverRequest) GetName() *string {
	return s.Name
}

func (s *ModifyProjectNotifyReceiverRequest) GetPhone() *string {
	return s.Phone
}

func (s *ModifyProjectNotifyReceiverRequest) GetReceiverId() *int64 {
	return s.ReceiverId
}

func (s *ModifyProjectNotifyReceiverRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyProjectNotifyReceiverRequest) SetEmail(v string) *ModifyProjectNotifyReceiverRequest {
	s.Email = &v
	return s
}

func (s *ModifyProjectNotifyReceiverRequest) SetInstanceId(v string) *ModifyProjectNotifyReceiverRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyProjectNotifyReceiverRequest) SetName(v string) *ModifyProjectNotifyReceiverRequest {
	s.Name = &v
	return s
}

func (s *ModifyProjectNotifyReceiverRequest) SetPhone(v string) *ModifyProjectNotifyReceiverRequest {
	s.Phone = &v
	return s
}

func (s *ModifyProjectNotifyReceiverRequest) SetReceiverId(v int64) *ModifyProjectNotifyReceiverRequest {
	s.ReceiverId = &v
	return s
}

func (s *ModifyProjectNotifyReceiverRequest) SetRegionId(v string) *ModifyProjectNotifyReceiverRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyProjectNotifyReceiverRequest) Validate() error {
	return dara.Validate(s)
}
