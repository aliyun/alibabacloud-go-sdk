// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectNotifyReceiversRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteProjectNotifyReceiversRequest
	GetInstanceId() *string
	SetReceiverIds(v string) *DeleteProjectNotifyReceiversRequest
	GetReceiverIds() *string
	SetRegionId(v string) *DeleteProjectNotifyReceiversRequest
	GetRegionId() *string
}

type DeleteProjectNotifyReceiversRequest struct {
	// The instance ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to retrieve this parameter.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the notification recipients to delete.
	//
	// > You can call the GetProject operation to retrieve this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReceiverIds *string `json:"ReceiverIds,omitempty" xml:"ReceiverIds,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteProjectNotifyReceiversRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectNotifyReceiversRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectNotifyReceiversRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteProjectNotifyReceiversRequest) GetReceiverIds() *string {
	return s.ReceiverIds
}

func (s *DeleteProjectNotifyReceiversRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteProjectNotifyReceiversRequest) SetInstanceId(v string) *DeleteProjectNotifyReceiversRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteProjectNotifyReceiversRequest) SetReceiverIds(v string) *DeleteProjectNotifyReceiversRequest {
	s.ReceiverIds = &v
	return s
}

func (s *DeleteProjectNotifyReceiversRequest) SetRegionId(v string) *DeleteProjectNotifyReceiversRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteProjectNotifyReceiversRequest) Validate() error {
	return dara.Validate(s)
}
