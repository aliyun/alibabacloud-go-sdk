// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAuthConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigList(v []*ModifyInstanceAuthConfigRequestConfigList) *ModifyInstanceAuthConfigRequest
	GetConfigList() []*ModifyInstanceAuthConfigRequestConfigList
	SetInstanceName(v string) *ModifyInstanceAuthConfigRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyInstanceAuthConfigRequest
	GetRegionId() *string
}

type ModifyInstanceAuthConfigRequest struct {
	// The ID of the RDS Supabase instance.
	ConfigList []*ModifyInstanceAuthConfigRequestConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The operation that you want to perform. Set the value to **ModifyInstanceAuthConfig**.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceAuthConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAuthConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAuthConfigRequest) GetConfigList() []*ModifyInstanceAuthConfigRequestConfigList {
	return s.ConfigList
}

func (s *ModifyInstanceAuthConfigRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceAuthConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceAuthConfigRequest) SetConfigList(v []*ModifyInstanceAuthConfigRequestConfigList) *ModifyInstanceAuthConfigRequest {
	s.ConfigList = v
	return s
}

func (s *ModifyInstanceAuthConfigRequest) SetInstanceName(v string) *ModifyInstanceAuthConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceAuthConfigRequest) SetRegionId(v string) *ModifyInstanceAuthConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceAuthConfigRequest) Validate() error {
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyInstanceAuthConfigRequestConfigList struct {
	// The name of the configuration item. Valid values:
	//
	// - **GOTRUE_EXTERNAL_EMAIL_ENABLED**: Enables external email addresses.
	//
	// - **GOTRUE_SITE_URL**: The website URL displayed in emails sent by the AI application.
	//
	// - **GOTRUE_SMTP_PORT**: The port of the SMTP service provider.
	//
	// - **GOTRUE_SMTP_SENDER_NAME**: The name of the email sender.
	//
	// - **GOTRUE_SMTP_USER**: The username of the SMTP service provider.
	//
	// - **GOTRUE_SMTP_PASS**: The key of the SMTP service provider.
	//
	// - **GOTRUE_SMTP_ADMIN_EMAIL**: The email address of the SMTP service provider.
	//
	// - **GOTRUE_SMTP_HOST**: The host address of the SMTP service provider.
	//
	// - **GOTRUE_MAILER_AUTOCONFIRM**: Specifies whether automatic confirmation is enabled.
	//
	// - **GOTRUE_MAILER_OTP_EXP**: The validity period of the one-time password (OTP), in seconds.
	//
	// - **GOTRUE_MAILER_OTP_LENGTH**: The length of the verification code for the one-time password (OTP). The value must be an integer greater than or equal to 6.
	//
	// example:
	//
	// GOTRUE_SITE_URL
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// http://8.152. XXX.XXX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyInstanceAuthConfigRequestConfigList) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAuthConfigRequestConfigList) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAuthConfigRequestConfigList) GetName() *string {
	return s.Name
}

func (s *ModifyInstanceAuthConfigRequestConfigList) GetValue() *string {
	return s.Value
}

func (s *ModifyInstanceAuthConfigRequestConfigList) SetName(v string) *ModifyInstanceAuthConfigRequestConfigList {
	s.Name = &v
	return s
}

func (s *ModifyInstanceAuthConfigRequestConfigList) SetValue(v string) *ModifyInstanceAuthConfigRequestConfigList {
	s.Value = &v
	return s
}

func (s *ModifyInstanceAuthConfigRequestConfigList) Validate() error {
	return dara.Validate(s)
}
