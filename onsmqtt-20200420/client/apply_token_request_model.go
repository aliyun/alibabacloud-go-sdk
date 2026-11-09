// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v string) *ApplyTokenRequest
	GetActions() *string
	SetExpireTime(v int64) *ApplyTokenRequest
	GetExpireTime() *int64
	SetInstanceId(v string) *ApplyTokenRequest
	GetInstanceId() *string
	SetResources(v string) *ApplyTokenRequest
	GetResources() *string
}

type ApplyTokenRequest struct {
	// The permission type of the token. Valid values:
	//
	// - **R**: read-only permissions. The client can only subscribe to the specified topics and cannot publish messages to the topics.
	//
	// - **W**: write-only permissions. The client can only publish messages to the specified topics and cannot subscribe to the topics.
	//
	// - **R,W**: read and write permissions. The client can publish messages to and subscribe to the specified topics. Separate **R*	- and **W*	- with a comma (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// R
	Actions *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The expiration time of the token. The value is a millisecond timestamp. The minimum effective period is 60 seconds, and the maximum effective period is 30 days. If the specified value exceeds 30 days, the API operation does not return an error, but the actual effective period is 30 days.
	//
	// Example: If the effective period of the token is 60 seconds and the current system time in milliseconds is 1609434061000, set this parameter to **1609434121000**, which is the sum of 1609434061000 and 60 × 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1609434121000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The instance ID must match the actual instance ID used by the client. You can obtain the instance ID on the **Instance Details*	- page in the [console](https://mqtt.console.aliyun.com/).
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The resource name, which is the topic of ApsaraMQ for MQTT. Separate multiple topics with commas (,). Each token can be used to perform operations on up to 100 resources. If multiple topics are specified, they must be sorted in lexicographical order.
	//
	// The resource parameter registered when you apply for a token supports the MQTT wildcard syntax, including the plus sign (+) for single-level wildcards and the number sign (#) for multi-level wildcards.
	//
	// For example, if you set **Resources*	- to "Topic1/+" when you apply for a token, the client can perform operations on any topic in the format of "Topic1/xxx". If you set **Resources*	- to "Topic1/#", the client can perform operations on any multi-level topic in the format of "Topic1/xxx/xxx/xxx".
	//
	// > ApsaraMQ for MQTT supports matching child topics. You do not need to create child topics in the console. You can directly configure the settings in the message sending and receiving code. The forward slash (/) in a topic name is used to separate the parent topic and child topics. For more information, see [Terms](https://help.aliyun.com/document_detail/42420.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// TopicA/+
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s ApplyTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyTokenRequest) GoString() string {
	return s.String()
}

func (s *ApplyTokenRequest) GetActions() *string {
	return s.Actions
}

func (s *ApplyTokenRequest) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ApplyTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ApplyTokenRequest) GetResources() *string {
	return s.Resources
}

func (s *ApplyTokenRequest) SetActions(v string) *ApplyTokenRequest {
	s.Actions = &v
	return s
}

func (s *ApplyTokenRequest) SetExpireTime(v int64) *ApplyTokenRequest {
	s.ExpireTime = &v
	return s
}

func (s *ApplyTokenRequest) SetInstanceId(v string) *ApplyTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *ApplyTokenRequest) SetResources(v string) *ApplyTokenRequest {
	s.Resources = &v
	return s
}

func (s *ApplyTokenRequest) Validate() error {
	return dara.Validate(s)
}
