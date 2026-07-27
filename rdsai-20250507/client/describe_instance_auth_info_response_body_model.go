// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAuthInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeys(v *DescribeInstanceAuthInfoResponseBodyApiKeys) *DescribeInstanceAuthInfoResponseBody
	GetApiKeys() *DescribeInstanceAuthInfoResponseBodyApiKeys
	SetBranchName(v string) *DescribeInstanceAuthInfoResponseBody
	GetBranchName() *string
	SetConfigList(v []*DescribeInstanceAuthInfoResponseBodyConfigList) *DescribeInstanceAuthInfoResponseBody
	GetConfigList() []*DescribeInstanceAuthInfoResponseBodyConfigList
	SetInstanceName(v string) *DescribeInstanceAuthInfoResponseBody
	GetInstanceName() *string
	SetJwtSecret(v string) *DescribeInstanceAuthInfoResponseBody
	GetJwtSecret() *string
	SetRequestId(v string) *DescribeInstanceAuthInfoResponseBody
	GetRequestId() *string
}

type DescribeInstanceAuthInfoResponseBody struct {
	// API Keys。
	ApiKeys    *DescribeInstanceAuthInfoResponseBodyApiKeys `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Struct"`
	BranchName *string                                      `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The list of authentication configurations.
	ConfigList []*DescribeInstanceAuthInfoResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The instance ID of the AI application.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The JWT secret.
	//
	// example:
	//
	// g7jgepleljS8nxAwsOd2EDWkBWi7JcU1m2Gj****
	JwtSecret *string `json:"JwtSecret,omitempty" xml:"JwtSecret,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceAuthInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAuthInfoResponseBody) GetApiKeys() *DescribeInstanceAuthInfoResponseBodyApiKeys {
	return s.ApiKeys
}

func (s *DescribeInstanceAuthInfoResponseBody) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeInstanceAuthInfoResponseBody) GetConfigList() []*DescribeInstanceAuthInfoResponseBodyConfigList {
	return s.ConfigList
}

func (s *DescribeInstanceAuthInfoResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceAuthInfoResponseBody) GetJwtSecret() *string {
	return s.JwtSecret
}

func (s *DescribeInstanceAuthInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAuthInfoResponseBody) SetApiKeys(v *DescribeInstanceAuthInfoResponseBodyApiKeys) *DescribeInstanceAuthInfoResponseBody {
	s.ApiKeys = v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBody) SetBranchName(v string) *DescribeInstanceAuthInfoResponseBody {
	s.BranchName = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBody) SetConfigList(v []*DescribeInstanceAuthInfoResponseBodyConfigList) *DescribeInstanceAuthInfoResponseBody {
	s.ConfigList = v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBody) SetInstanceName(v string) *DescribeInstanceAuthInfoResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBody) SetJwtSecret(v string) *DescribeInstanceAuthInfoResponseBody {
	s.JwtSecret = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBody) SetRequestId(v string) *DescribeInstanceAuthInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBody) Validate() error {
	if s.ApiKeys != nil {
		if err := s.ApiKeys.Validate(); err != nil {
			return err
		}
	}
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

type DescribeInstanceAuthInfoResponseBodyApiKeys struct {
	// The ANON_KEY of Supabase.
	//
	// example:
	//
	// eyJ0eXAiOiJKV1QiLCJhbGciOiJIUz****J9.eyJpc3MiOiJzdXBhYmFzZSIsInJvbGUiOiJhbm9uIiwiaWF0IjoxNzU1Nzg1ODc1LCJleHAiOjEzMjY2NDI1ODc1fQ.EGNFdeWRZBsdB051EzQsBwvDJveC9IMEXWUCDLX****
	AnonKey   *string `json:"AnonKey,omitempty" xml:"AnonKey,omitempty"`
	E2bApiKey *string `json:"E2bApiKey,omitempty" xml:"E2bApiKey,omitempty"`
	// The SERVICE_ROLE_KEY of Supabase.
	//
	// example:
	//
	// eyJ0eXAiOiJKV1QiLCJhbGciOiJIUz****J9.eyJpc3MiOiJzdXBhYmFzZSIsInJvbGUiOiJzZXJ2aWNlX3JvbGUiLCJpYXQiOjE3NTU3ODU4NzUsImV4cCI6MTMyNjY0MjU4NzV9.oJt4UF8cpSDOvjW39IM4fLp2750rEvxFnkNqcVM****
	ServiceKey *string `json:"ServiceKey,omitempty" xml:"ServiceKey,omitempty"`
}

func (s DescribeInstanceAuthInfoResponseBodyApiKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAuthInfoResponseBodyApiKeys) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAuthInfoResponseBodyApiKeys) GetAnonKey() *string {
	return s.AnonKey
}

func (s *DescribeInstanceAuthInfoResponseBodyApiKeys) GetE2bApiKey() *string {
	return s.E2bApiKey
}

func (s *DescribeInstanceAuthInfoResponseBodyApiKeys) GetServiceKey() *string {
	return s.ServiceKey
}

func (s *DescribeInstanceAuthInfoResponseBodyApiKeys) SetAnonKey(v string) *DescribeInstanceAuthInfoResponseBodyApiKeys {
	s.AnonKey = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBodyApiKeys) SetE2bApiKey(v string) *DescribeInstanceAuthInfoResponseBodyApiKeys {
	s.E2bApiKey = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBodyApiKeys) SetServiceKey(v string) *DescribeInstanceAuthInfoResponseBodyApiKeys {
	s.ServiceKey = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBodyApiKeys) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAuthInfoResponseBodyConfigList struct {
	// The name of the configuration item. Valid values:
	//
	// - **GOTRUE_EXTERNAL_EMAIL_ENABLED**: external email enabled.
	//
	// - **GOTRUE_SITE_URL**: the website URL displayed when the AI application sends emails.
	//
	// - **GOTRUE_SMTP_PORT**: the port of the SMTP provider.
	//
	// - **GOTRUE_SMTP_SENDER_NAME**: the name of the email sender.
	//
	// - **GOTRUE_SMTP_USER**: the username of the SMTP provider.
	//
	// - **GOTRUE_SMTP_PASS**: the secret of the SMTP provider.
	//
	// - **GOTRUE_SMTP_ADMIN_EMAIL**: the email address of the SMTP provider.
	//
	// - **GOTRUE_SMTP_HOST**: the host address of the SMTP provider.
	//
	// - **GOTRUE_MAILER_AUTOCONFIRM**: specifies whether to enable autoconfirm.
	//
	// - **GOTRUE_MAILER_OTP_EXP**: the validity period of the one-time password (OTP). Unit: seconds.
	//
	// - **GOTRUE_MAILER_OTP_LENGTH**: the length of the one-time password (OTP) verification code. The value must be an integer greater than or equal to 6.
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

func (s DescribeInstanceAuthInfoResponseBodyConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAuthInfoResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAuthInfoResponseBodyConfigList) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceAuthInfoResponseBodyConfigList) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceAuthInfoResponseBodyConfigList) SetName(v string) *DescribeInstanceAuthInfoResponseBodyConfigList {
	s.Name = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBodyConfigList) SetValue(v string) *DescribeInstanceAuthInfoResponseBodyConfigList {
	s.Value = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponseBodyConfigList) Validate() error {
	return dara.Validate(s)
}
