// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceConfigRequest
	GetInstanceId() *string
	SetReason(v string) *ModifyInstanceConfigRequest
	GetReason() *string
	SetUserConfig(v string) *ModifyInstanceConfigRequest
	GetUserConfig() *string
}

type ModifyInstanceConfigRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The reason for the update.
	//
	// This parameter is required.
	//
	// example:
	//
	// for test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// User-defined configuration.
	//
	// example:
	//
	// dataCoord:\\n  segment:\\n    maxSize: 1024
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
}

func (s ModifyInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceConfigRequest) GetReason() *string {
	return s.Reason
}

func (s *ModifyInstanceConfigRequest) GetUserConfig() *string {
	return s.UserConfig
}

func (s *ModifyInstanceConfigRequest) SetInstanceId(v string) *ModifyInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetReason(v string) *ModifyInstanceConfigRequest {
	s.Reason = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetUserConfig(v string) *ModifyInstanceConfigRequest {
	s.UserConfig = &v
	return s
}

func (s *ModifyInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
