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
	// example:
	//
	// F4D7C841-A1C9-50B4-88B7-F5****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Config-000****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// gtm-cn-wwo3a3h****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
