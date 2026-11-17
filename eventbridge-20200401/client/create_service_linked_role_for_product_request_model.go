// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleForProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductName(v string) *CreateServiceLinkedRoleForProductRequest
	GetProductName() *string
}

type CreateServiceLinkedRoleForProductRequest struct {
	// The name of the cloud service or the name of the service-linked role associated with the cloud service.
	//
	// - AliyunServiceRoleForEventBridgeSendToFC: Performs event delivery to Function Compute.
	//
	// - AliyunServiceRoleForEventBridgeSendToSMS: Performs event delivery as text messages to Alibaba Cloud Short Message Service.
	//
	// - AliyunServiceRoleForEventBridgeSendToDirectMail: Performs event delivery as emails to Alibaba Cloud Direct Mail.
	//
	// - AliyunServiceRoleForEventBridgeSourceRocketMQ: Integrates consumed messages from Message Queue for RocketMQ instances into EventBridge.
	//
	// - AliyunServiceRoleForEventBridgeSourceMNS: Integrates Alibaba Cloud Message Service (MNS) into EventBridge.
	//
	// - AliyunServiceRoleForEventBridgeConnectVPC: Allows EventBridge to access your VPC.
	//
	// - AliyunServiceRoleForEventBridgeSourceActionTrail: Retrieves operation records from ActionTrail as event sources.
	//
	// - AliyunServiceRoleForEventBridgeSourceRabbitMQ: Integrates Message Queue for RabbitMQ instances into EventBridge.
	//
	// - AliyunServiceRoleForEventBridgeSendToRabbitMQ: Performs event delivery of EventBridge events to Message Queue for RabbitMQ instances.
	//
	// - AliyunServiceRoleForEventBridgeSendToRocketMQ: Performs event delivery of EventBridge events to Message Queue for RocketMQ instances.
	//
	// - AliyunServiceRoleForEventBridgeSourceCMS: Integrates Alibaba Cloud monitoring (CMS) into EventBridge.
	//
	// - AliyunServiceRoleForEventBridgeSendToKafka: Performs event delivery of EventBridge events to Message Queue for Apache Kafka clusters.
	//
	// - AliyunServiceRoleForEventBridgeSourceKafka: Integrates Message Queue for Apache Kafka into EventBridge.
	//
	// - AliyunServiceRoleForEventBridgeSendToRDS: Performs event delivery of EventBridge events to Relational Database Service (RDS) instances.
	//
	// - AliyunServiceRoleForEventBridgeSendToSAE: Performs event delivery of EventBridge events to Serverless App DPI engine (SAE) applications.
	//
	// - AliyunServiceRoleForEventBridgeSourceMqtt: Integrates Micro MSMQ for MQTT into EventBridge.
	//
	// - AliyunServiceRoleForEventBridgeSourceSLS: Integrates Simple Log Service into EventBridge.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunServiceRoleForEventBridgeSendToMNS
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s CreateServiceLinkedRoleForProductRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleForProductRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForProductRequest) GetProductName() *string {
	return s.ProductName
}

func (s *CreateServiceLinkedRoleForProductRequest) SetProductName(v string) *CreateServiceLinkedRoleForProductRequest {
	s.ProductName = &v
	return s
}

func (s *CreateServiceLinkedRoleForProductRequest) Validate() error {
	return dara.Validate(s)
}
