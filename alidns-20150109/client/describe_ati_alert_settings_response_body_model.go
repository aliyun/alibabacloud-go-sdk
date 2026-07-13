// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiAlertSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) *DescribeAtiAlertSettingsResponseBody
	GetAccessDeniedDetail() *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail
	SetAlertConfig(v *DescribeAtiAlertSettingsResponseBodyAlertConfig) *DescribeAtiAlertSettingsResponseBody
	GetAlertConfig() *DescribeAtiAlertSettingsResponseBodyAlertConfig
	SetAlertGroup(v *DescribeAtiAlertSettingsResponseBodyAlertGroup) *DescribeAtiAlertSettingsResponseBody
	GetAlertGroup() *DescribeAtiAlertSettingsResponseBodyAlertGroup
	SetRequestId(v string) *DescribeAtiAlertSettingsResponseBody
	GetRequestId() *string
}

type DescribeAtiAlertSettingsResponseBody struct {
	// The details about the access denial. This field is returned only when the RAM authentication fails.
	AccessDeniedDetail *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	AlertConfig        *DescribeAtiAlertSettingsResponseBodyAlertConfig        `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Struct"`
	AlertGroup         *DescribeAtiAlertSettingsResponseBodyAlertGroup         `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAtiAlertSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAlertSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAtiAlertSettingsResponseBody) GetAccessDeniedDetail() *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeAtiAlertSettingsResponseBody) GetAlertConfig() *DescribeAtiAlertSettingsResponseBodyAlertConfig {
	return s.AlertConfig
}

func (s *DescribeAtiAlertSettingsResponseBody) GetAlertGroup() *DescribeAtiAlertSettingsResponseBodyAlertGroup {
	return s.AlertGroup
}

func (s *DescribeAtiAlertSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAtiAlertSettingsResponseBody) SetAccessDeniedDetail(v *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) *DescribeAtiAlertSettingsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBody) SetAlertConfig(v *DescribeAtiAlertSettingsResponseBodyAlertConfig) *DescribeAtiAlertSettingsResponseBody {
	s.AlertConfig = v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBody) SetAlertGroup(v *DescribeAtiAlertSettingsResponseBodyAlertGroup) *DescribeAtiAlertSettingsResponseBody {
	s.AlertGroup = v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBody) SetRequestId(v string) *DescribeAtiAlertSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.AlertConfig != nil {
		if err := s.AlertConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AlertGroup != nil {
		if err := s.AlertGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// UpdateRspDomainServerProhibitStatusForGatewayOte
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorization principal.
	//
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The ID of the authorization principal owner.
	//
	// example:
	//
	// 10469733312XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encrypted diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaNIARXXXXUQwNjE0LUQzN0XXXXVEQy1BQzExLTMzXXXXNTkxRjk1Ng==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The reason for the authentication failure. Valid values:
	//
	// - ExplicitDeny: explicit deny.
	//
	// - ImplicitDeny: implicit deny.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeAtiAlertSettingsResponseBodyAlertConfig struct {
	AlertConfig []*DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
}

func (s DescribeAtiAlertSettingsResponseBodyAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAlertSettingsResponseBodyAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertConfig) GetAlertConfig() []*DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig {
	return s.AlertConfig
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertConfig) SetAlertConfig(v []*DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig) *DescribeAtiAlertSettingsResponseBodyAlertConfig {
	s.AlertConfig = v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertConfig) Validate() error {
	if s.AlertConfig != nil {
		for _, item := range s.AlertConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig struct {
	DingtalkNotice *bool   `json:"DingtalkNotice,omitempty" xml:"DingtalkNotice,omitempty"`
	EmailNotice    *bool   `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
	NoticeType     *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	SmsNotice      *bool   `json:"SmsNotice,omitempty" xml:"SmsNotice,omitempty"`
}

func (s DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig) GetDingtalkNotice() *bool {
	return s.DingtalkNotice
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig) GetEmailNotice() *bool {
	return s.EmailNotice
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig) GetNoticeType() *string {
	return s.NoticeType
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig) GetSmsNotice() *bool {
	return s.SmsNotice
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig) SetDingtalkNotice(v bool) *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig {
	s.DingtalkNotice = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig) SetEmailNotice(v bool) *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig {
	s.EmailNotice = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig) SetNoticeType(v string) *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig {
	s.NoticeType = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig) SetSmsNotice(v bool) *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig {
	s.SmsNotice = &v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertConfigAlertConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAtiAlertSettingsResponseBodyAlertGroup struct {
	AlertGroup []*string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty" type:"Repeated"`
}

func (s DescribeAtiAlertSettingsResponseBodyAlertGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAlertSettingsResponseBodyAlertGroup) GoString() string {
	return s.String()
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertGroup) GetAlertGroup() []*string {
	return s.AlertGroup
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertGroup) SetAlertGroup(v []*string) *DescribeAtiAlertSettingsResponseBodyAlertGroup {
	s.AlertGroup = v
	return s
}

func (s *DescribeAtiAlertSettingsResponseBodyAlertGroup) Validate() error {
	return dara.Validate(s)
}
