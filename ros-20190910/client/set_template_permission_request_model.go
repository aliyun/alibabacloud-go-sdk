// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTemplatePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *SetTemplatePermissionRequest
	GetAccountIds() []*string
	SetShareOption(v string) *SetTemplatePermissionRequest
	GetShareOption() *string
	SetTemplateId(v string) *SetTemplatePermissionRequest
	GetTemplateId() *string
	SetTemplateVersion(v string) *SetTemplatePermissionRequest
	GetTemplateVersion() *string
	SetVersionOption(v string) *SetTemplatePermissionRequest
	GetVersionOption() *string
}

type SetTemplatePermissionRequest struct {
	// The Alibaba Cloud accounts with or from which you want to share or unshare the template.\\
	//
	// Valid values of N: 1, 2, 3, 4, and 5.
	//
	// > - This parameter cannot be set to the ID of the Alibaba Cloud account that owns the template, or the RAM users of this Alibaba Cloud account.
	//
	// > - When ShareOption is set to CancelSharing, you can unshare the template from all the specified Alibaba Cloud accounts by using an asterisk (\\*).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The sharing option.
	//
	// Valid values:
	//
	// 	- ShareToAccounts: shares the template with other Alibaba Cloud accounts.
	//
	// 	- CancelSharing: unshares the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ShareToAccounts
	ShareOption *string `json:"ShareOption,omitempty" xml:"ShareOption,omitempty"`
	// The ID of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The version of the shared template. This parameter takes effect only if you set ShareOption to ShareToAccounts and set VersionOption to Specified.
	//
	// Valid values: v1 to v100.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The version option for the shared template. This parameter takes effect only if you set ShareOption to ShareToAccounts.
	//
	// Valid values:
	//
	// 	- AllVersions (default): shares all versions of the template.
	//
	// 	- Latest: shares only the latest version of template. When the version of the template is updated, ROS updates the shared version to the latest version.
	//
	// 	- Current: shares only the current version of the template. When the version of the template is updated, ROS does not update the shared version.
	//
	// 	- Specified: shares only the specified version of the template.
	//
	// example:
	//
	// Specified
	VersionOption *string `json:"VersionOption,omitempty" xml:"VersionOption,omitempty"`
}

func (s SetTemplatePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetTemplatePermissionRequest) GoString() string {
	return s.String()
}

func (s *SetTemplatePermissionRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *SetTemplatePermissionRequest) GetShareOption() *string {
	return s.ShareOption
}

func (s *SetTemplatePermissionRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SetTemplatePermissionRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *SetTemplatePermissionRequest) GetVersionOption() *string {
	return s.VersionOption
}

func (s *SetTemplatePermissionRequest) SetAccountIds(v []*string) *SetTemplatePermissionRequest {
	s.AccountIds = v
	return s
}

func (s *SetTemplatePermissionRequest) SetShareOption(v string) *SetTemplatePermissionRequest {
	s.ShareOption = &v
	return s
}

func (s *SetTemplatePermissionRequest) SetTemplateId(v string) *SetTemplatePermissionRequest {
	s.TemplateId = &v
	return s
}

func (s *SetTemplatePermissionRequest) SetTemplateVersion(v string) *SetTemplatePermissionRequest {
	s.TemplateVersion = &v
	return s
}

func (s *SetTemplatePermissionRequest) SetVersionOption(v string) *SetTemplatePermissionRequest {
	s.VersionOption = &v
	return s
}

func (s *SetTemplatePermissionRequest) Validate() error {
	return dara.Validate(s)
}
