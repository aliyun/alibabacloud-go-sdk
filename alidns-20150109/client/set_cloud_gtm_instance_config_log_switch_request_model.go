// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCloudGtmInstanceConfigLogSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SetCloudGtmInstanceConfigLogSwitchRequest
	GetClientToken() *string
	SetConfigId(v string) *SetCloudGtmInstanceConfigLogSwitchRequest
	GetConfigId() *string
	SetInstanceId(v string) *SetCloudGtmInstanceConfigLogSwitchRequest
	GetInstanceId() *string
	SetStatus(v string) *SetCloudGtmInstanceConfigLogSwitchRequest
	GetStatus() *string
}

type SetCloudGtmInstanceConfigLogSwitchRequest struct {
	// The client token that is used to ensure the idempotence of the request. Generate a unique value for this parameter. The client token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId of each API request may be different.
	//
	// example:
	//
	// F4D7C841-A1C9-50B4-88B7-F5****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance configuration. A and AAAA records can be configured for the same connected domain name and GTM instance. In this case, the GTM instance has two instance configurations. ConfigId uniquely identifies an instance configuration. To find the ConfigId for a domain name instance, call the [ListCloudGtmInstanceConfigs](https://help.aliyun.com/document_detail/2797349.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// Config-000****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the Global Traffic Manager (GTM) 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3h****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the network traffic analysis feature:
	//
	// - enable
	//
	// - disable
	//
	// This parameter is required.
	//
	// example:
	//
	// enable
	//
	// disable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetCloudGtmInstanceConfigLogSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCloudGtmInstanceConfigLogSwitchRequest) GoString() string {
	return s.String()
}

func (s *SetCloudGtmInstanceConfigLogSwitchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetCloudGtmInstanceConfigLogSwitchRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *SetCloudGtmInstanceConfigLogSwitchRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetCloudGtmInstanceConfigLogSwitchRequest) GetStatus() *string {
	return s.Status
}

func (s *SetCloudGtmInstanceConfigLogSwitchRequest) SetClientToken(v string) *SetCloudGtmInstanceConfigLogSwitchRequest {
	s.ClientToken = &v
	return s
}

func (s *SetCloudGtmInstanceConfigLogSwitchRequest) SetConfigId(v string) *SetCloudGtmInstanceConfigLogSwitchRequest {
	s.ConfigId = &v
	return s
}

func (s *SetCloudGtmInstanceConfigLogSwitchRequest) SetInstanceId(v string) *SetCloudGtmInstanceConfigLogSwitchRequest {
	s.InstanceId = &v
	return s
}

func (s *SetCloudGtmInstanceConfigLogSwitchRequest) SetStatus(v string) *SetCloudGtmInstanceConfigLogSwitchRequest {
	s.Status = &v
	return s
}

func (s *SetCloudGtmInstanceConfigLogSwitchRequest) Validate() error {
	return dara.Validate(s)
}
