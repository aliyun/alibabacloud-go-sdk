// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserProfilePathRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *SetUserProfilePathRulesRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupSetId(v string) *SetUserProfilePathRulesRequest
	GetAppInstanceGroupSetId() *string
	SetRegionId(v string) *SetUserProfilePathRulesRequest
	GetRegionId() *string
	SetUserProfilePathRules(v string) *SetUserProfilePathRulesRequest
	GetUserProfilePathRules() *string
	SetUserProfileRuleType(v string) *SetUserProfilePathRulesRequest
	GetUserProfileRuleType() *string
}

type SetUserProfilePathRulesRequest struct {
	// The ID of a single app instance group. You must specify either this parameter or AppInstanceGroupSetId. If the cloud application or cloud browser app instance group belongs to an enabled set, you must configure the rules for the set.
	//
	// example:
	//
	// aig-0bi83v****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of an enabled app instance group set. You must specify either this parameter or AppInstanceGroupId. You can configure rules for up to 200 member app instance groups at a time.
	//
	// example:
	//
	// set-3jm9d0abc00example
	AppInstanceGroupSetId *string `json:"AppInstanceGroupSetId,omitempty" xml:"AppInstanceGroupSetId,omitempty"`
	// The region ID. When you query or configure rules for a specific app instance group, specify the region where the target app instance group resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The JSON array string that contains the complete list of rules. Each item uses `blackPath` to specify the excluded path and `whitePaths` to specify the relative paths that still require roaming within the excluded directory. The path type `file` indicates a file, and `folder` indicates a directory.
	//
	// example:
	//
	// [{"blackPath":{"path":"AppData\\\\Local","type":"folder"},"whitePaths":[{"path":".\\\\Microsoft\\\\Office","type":"folder"}]}]
	UserProfilePathRules *string `json:"UserProfilePathRules,omitempty" xml:"UserProfilePathRules,omitempty"`
	// The type of rule to configure. Set this parameter to `AppInstanceGroup` when you configure rules for an app instance group or an app instance group set.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppInstanceGroup
	UserProfileRuleType *string `json:"UserProfileRuleType,omitempty" xml:"UserProfileRuleType,omitempty"`
}

func (s SetUserProfilePathRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetUserProfilePathRulesRequest) GoString() string {
	return s.String()
}

func (s *SetUserProfilePathRulesRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *SetUserProfilePathRulesRequest) GetAppInstanceGroupSetId() *string {
	return s.AppInstanceGroupSetId
}

func (s *SetUserProfilePathRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetUserProfilePathRulesRequest) GetUserProfilePathRules() *string {
	return s.UserProfilePathRules
}

func (s *SetUserProfilePathRulesRequest) GetUserProfileRuleType() *string {
	return s.UserProfileRuleType
}

func (s *SetUserProfilePathRulesRequest) SetAppInstanceGroupId(v string) *SetUserProfilePathRulesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *SetUserProfilePathRulesRequest) SetAppInstanceGroupSetId(v string) *SetUserProfilePathRulesRequest {
	s.AppInstanceGroupSetId = &v
	return s
}

func (s *SetUserProfilePathRulesRequest) SetRegionId(v string) *SetUserProfilePathRulesRequest {
	s.RegionId = &v
	return s
}

func (s *SetUserProfilePathRulesRequest) SetUserProfilePathRules(v string) *SetUserProfilePathRulesRequest {
	s.UserProfilePathRules = &v
	return s
}

func (s *SetUserProfilePathRulesRequest) SetUserProfileRuleType(v string) *SetUserProfilePathRulesRequest {
	s.UserProfileRuleType = &v
	return s
}

func (s *SetUserProfilePathRulesRequest) Validate() error {
	return dara.Validate(s)
}
