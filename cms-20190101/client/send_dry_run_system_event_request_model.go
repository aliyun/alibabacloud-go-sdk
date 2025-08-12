// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendDryRunSystemEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventContent(v string) *SendDryRunSystemEventRequest
	GetEventContent() *string
	SetEventName(v string) *SendDryRunSystemEventRequest
	GetEventName() *string
	SetGroupId(v string) *SendDryRunSystemEventRequest
	GetGroupId() *string
	SetProduct(v string) *SendDryRunSystemEventRequest
	GetProduct() *string
	SetRegionId(v string) *SendDryRunSystemEventRequest
	GetRegionId() *string
}

type SendDryRunSystemEventRequest struct {
	// The content of the system event.
	//
	// >  The value of this parameter is a JSON object. We recommend that you include the `product`, `resourceId`, and `regionId` fields in the JSON object.
	//
	// example:
	//
	// {"product":"CloudMonitor","resourceId":"acs:ecs:cn-hongkong:173651113438****:instance/{instanceId}","level":"CRITICAL","instanceName":"instanceName","regionId":"cn-hangzhou","name":"Agent_Status_Stopped","content":{"ipGroup":"0.0.0.0,0.0.0.1","tianjimonVersion":"1.2.11"},"status":"stopped"}
	EventContent *string `json:"EventContent,omitempty" xml:"EventContent,omitempty"`
	// The name of the system event.
	//
	// >  For more information, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Agent_Status_Stopped
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 123456
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the cloud service.
	//
	// >  For information about the Alibaba Cloud services that are supported by CloudMonitor, see [Supported cloud services and their system events](https://help.aliyun.com/document_detail/167388.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SendDryRunSystemEventRequest) String() string {
	return dara.Prettify(s)
}

func (s SendDryRunSystemEventRequest) GoString() string {
	return s.String()
}

func (s *SendDryRunSystemEventRequest) GetEventContent() *string {
	return s.EventContent
}

func (s *SendDryRunSystemEventRequest) GetEventName() *string {
	return s.EventName
}

func (s *SendDryRunSystemEventRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SendDryRunSystemEventRequest) GetProduct() *string {
	return s.Product
}

func (s *SendDryRunSystemEventRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SendDryRunSystemEventRequest) SetEventContent(v string) *SendDryRunSystemEventRequest {
	s.EventContent = &v
	return s
}

func (s *SendDryRunSystemEventRequest) SetEventName(v string) *SendDryRunSystemEventRequest {
	s.EventName = &v
	return s
}

func (s *SendDryRunSystemEventRequest) SetGroupId(v string) *SendDryRunSystemEventRequest {
	s.GroupId = &v
	return s
}

func (s *SendDryRunSystemEventRequest) SetProduct(v string) *SendDryRunSystemEventRequest {
	s.Product = &v
	return s
}

func (s *SendDryRunSystemEventRequest) SetRegionId(v string) *SendDryRunSystemEventRequest {
	s.RegionId = &v
	return s
}

func (s *SendDryRunSystemEventRequest) Validate() error {
	return dara.Validate(s)
}
