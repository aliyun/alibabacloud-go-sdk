// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordHistoryConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SetPasswordHistoryConfigurationRequest
	GetInstanceId() *string
	SetPasswordHistoryMaxRetention(v int32) *SetPasswordHistoryConfigurationRequest
	GetPasswordHistoryMaxRetention() *int32
	SetPasswordHistoryStatus(v string) *SetPasswordHistoryConfigurationRequest
	GetPasswordHistoryStatus() *string
}

type SetPasswordHistoryConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of recent passwords that can be retained. This parameter must be specified when PasswordHistoryStatus is set to enabled.
	//
	// example:
	//
	// 3
	PasswordHistoryMaxRetention *int32 `json:"PasswordHistoryMaxRetention,omitempty" xml:"PasswordHistoryMaxRetention,omitempty"`
	// Specifies whether to enable the password history feature. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// This parameter is required.
	//
	// example:
	//
	// enabled
	PasswordHistoryStatus *string `json:"PasswordHistoryStatus,omitempty" xml:"PasswordHistoryStatus,omitempty"`
}

func (s SetPasswordHistoryConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordHistoryConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordHistoryConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPasswordHistoryConfigurationRequest) GetPasswordHistoryMaxRetention() *int32 {
	return s.PasswordHistoryMaxRetention
}

func (s *SetPasswordHistoryConfigurationRequest) GetPasswordHistoryStatus() *string {
	return s.PasswordHistoryStatus
}

func (s *SetPasswordHistoryConfigurationRequest) SetInstanceId(v string) *SetPasswordHistoryConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPasswordHistoryConfigurationRequest) SetPasswordHistoryMaxRetention(v int32) *SetPasswordHistoryConfigurationRequest {
	s.PasswordHistoryMaxRetention = &v
	return s
}

func (s *SetPasswordHistoryConfigurationRequest) SetPasswordHistoryStatus(v string) *SetPasswordHistoryConfigurationRequest {
	s.PasswordHistoryStatus = &v
	return s
}

func (s *SetPasswordHistoryConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
