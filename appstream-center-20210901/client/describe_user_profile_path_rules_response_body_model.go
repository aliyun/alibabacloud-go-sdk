// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserProfilePathRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeUserProfilePathRulesResponseBody
	GetCode() *string
	SetData(v *DescribeUserProfilePathRulesResponseBodyData) *DescribeUserProfilePathRulesResponseBody
	GetData() *DescribeUserProfilePathRulesResponseBodyData
	SetMessage(v string) *DescribeUserProfilePathRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeUserProfilePathRulesResponseBody
	GetRequestId() *string
}

type DescribeUserProfilePathRulesResponseBody struct {
	// The error code returned when the request fails.
	//
	// example:
	//
	// InvalidParameter.RuleType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The path rules and user profile roaming configurations.
	Data *DescribeUserProfilePathRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// The parameter RuleType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// 01A1CEEB-CDFF-18F9-B2BF-9997600D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserProfilePathRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeUserProfilePathRulesResponseBody) GetData() *DescribeUserProfilePathRulesResponseBodyData {
	return s.Data
}

func (s *DescribeUserProfilePathRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeUserProfilePathRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserProfilePathRulesResponseBody) SetCode(v string) *DescribeUserProfilePathRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBody) SetData(v *DescribeUserProfilePathRulesResponseBodyData) *DescribeUserProfilePathRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBody) SetMessage(v string) *DescribeUserProfilePathRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBody) SetRequestId(v string) *DescribeUserProfilePathRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserProfilePathRulesResponseBodyData struct {
	// The app instance group ID corresponding to the rule. When querying set rules, this is the primary app instance group ID.
	//
	// example:
	//
	// aig-0bi83v****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The app instance group set ID corresponding to the rule. This is returned when querying set rules.
	//
	// example:
	//
	// set-3jm9d0abc00example
	AppInstanceGroupSetId *string `json:"AppInstanceGroupSetId,omitempty" xml:"AppInstanceGroupSetId,omitempty"`
	// The user profile roaming storage domain ID bound to the app instance group. This field may be omitted if no domain is bound.
	//
	// example:
	//
	// domain-existing
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// Not enabled. The current query result does not provide the file system ID. The example is used to illustrate the field status.
	//
	// example:
	//
	// Not applicable (not enabled)
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The user profile roaming switch. A value of `true` indicates that the feature is enabled, and a value of `false` indicates that the feature is disabled.
	ProfileFollowSwitch *bool `json:"ProfileFollowSwitch,omitempty" xml:"ProfileFollowSwitch,omitempty"`
	// The remote storage path for user profile roaming. If this path is not separately configured for the app instance group, the app instance group ID is returned.
	//
	// example:
	//
	// aig-0bi83v****
	RemoteStoragePath *string `json:"RemoteStoragePath,omitempty" xml:"RemoteStoragePath,omitempty"`
	// The path rule type. Valid values:
	//
	// - `Default`: Queries system default rules. Both target IDs can be omitted.
	//
	// - `AppInstanceGroup`: Queries the rules of an app instance group or an app instance group set.
	//
	// example:
	//
	// AppInstanceGroup
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The list of user profile roaming path rules.
	Rules []*DescribeUserProfilePathRulesResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The storage quota for user profile roaming, in GB. This field may be omitted if it is not configured.
	//
	// example:
	//
	// 30
	UserProfileSize *int64 `json:"UserProfileSize,omitempty" xml:"UserProfileSize,omitempty"`
	// The user profile roaming switch. The value is the same as ProfileFollowSwitch. A value of `true` indicates that the feature is enabled, and a value of `false` indicates that the feature is disabled.
	UserProfileSwitch *bool `json:"UserProfileSwitch,omitempty" xml:"UserProfileSwitch,omitempty"`
}

func (s DescribeUserProfilePathRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesResponseBodyData) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *DescribeUserProfilePathRulesResponseBodyData) GetAppInstanceGroupSetId() *string {
	return s.AppInstanceGroupSetId
}

func (s *DescribeUserProfilePathRulesResponseBodyData) GetDomainId() *string {
	return s.DomainId
}

func (s *DescribeUserProfilePathRulesResponseBodyData) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeUserProfilePathRulesResponseBodyData) GetProfileFollowSwitch() *bool {
	return s.ProfileFollowSwitch
}

func (s *DescribeUserProfilePathRulesResponseBodyData) GetRemoteStoragePath() *string {
	return s.RemoteStoragePath
}

func (s *DescribeUserProfilePathRulesResponseBodyData) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeUserProfilePathRulesResponseBodyData) GetRules() []*DescribeUserProfilePathRulesResponseBodyDataRules {
	return s.Rules
}

func (s *DescribeUserProfilePathRulesResponseBodyData) GetUserProfileSize() *int64 {
	return s.UserProfileSize
}

func (s *DescribeUserProfilePathRulesResponseBodyData) GetUserProfileSwitch() *bool {
	return s.UserProfileSwitch
}

func (s *DescribeUserProfilePathRulesResponseBodyData) SetAppInstanceGroupId(v string) *DescribeUserProfilePathRulesResponseBodyData {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyData) SetAppInstanceGroupSetId(v string) *DescribeUserProfilePathRulesResponseBodyData {
	s.AppInstanceGroupSetId = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyData) SetDomainId(v string) *DescribeUserProfilePathRulesResponseBodyData {
	s.DomainId = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyData) SetFileSystemId(v string) *DescribeUserProfilePathRulesResponseBodyData {
	s.FileSystemId = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyData) SetProfileFollowSwitch(v bool) *DescribeUserProfilePathRulesResponseBodyData {
	s.ProfileFollowSwitch = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyData) SetRemoteStoragePath(v string) *DescribeUserProfilePathRulesResponseBodyData {
	s.RemoteStoragePath = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyData) SetRuleType(v string) *DescribeUserProfilePathRulesResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyData) SetRules(v []*DescribeUserProfilePathRulesResponseBodyDataRules) *DescribeUserProfilePathRulesResponseBodyData {
	s.Rules = v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyData) SetUserProfileSize(v int64) *DescribeUserProfilePathRulesResponseBodyData {
	s.UserProfileSize = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyData) SetUserProfileSwitch(v bool) *DescribeUserProfilePathRulesResponseBodyData {
	s.UserProfileSwitch = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyData) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserProfilePathRulesResponseBodyDataRules struct {
	// The file or directory excluded from user profile roaming.
	BlackPath *DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath `json:"BlackPath,omitempty" xml:"BlackPath,omitempty" type:"Struct"`
	// The list of files or directories that still participate in user profile roaming under the excluded directories.
	WhitePaths []*DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths `json:"WhitePaths,omitempty" xml:"WhitePaths,omitempty" type:"Repeated"`
}

func (s DescribeUserProfilePathRulesResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRules) GetBlackPath() *DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath {
	return s.BlackPath
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRules) GetWhitePaths() []*DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths {
	return s.WhitePaths
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRules) SetBlackPath(v *DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath) *DescribeUserProfilePathRulesResponseBodyDataRules {
	s.BlackPath = v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRules) SetWhitePaths(v []*DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths) *DescribeUserProfilePathRulesResponseBodyDataRules {
	s.WhitePaths = v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRules) Validate() error {
	if s.BlackPath != nil {
		if err := s.BlackPath.Validate(); err != nil {
			return err
		}
	}
	if s.WhitePaths != nil {
		for _, item := range s.WhitePaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath struct {
	// The excluded path relative to the user profile directory.
	//
	// example:
	//
	// AppData\\Local
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The path type. Valid values:
	//
	// - `file`: File.
	//
	// - `folder`: Folder.
	//
	// example:
	//
	// folder
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath) GetPath() *string {
	return s.Path
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath) GetType() *string {
	return s.Type
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath) SetPath(v string) *DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath {
	s.Path = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath) SetType(v string) *DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath {
	s.Type = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRulesBlackPath) Validate() error {
	return dara.Validate(s)
}

type DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths struct {
	// The exception path relative to BlackPath.Path.
	//
	// example:
	//
	// .\\Microsoft\\Office
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The path type. Valid values:
	//
	// - `file`: File.
	//
	// - `folder`: Folder.
	//
	// example:
	//
	// folder
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths) GetPath() *string {
	return s.Path
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths) GetType() *string {
	return s.Type
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths) SetPath(v string) *DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths {
	s.Path = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths) SetType(v string) *DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths {
	s.Type = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponseBodyDataRulesWhitePaths) Validate() error {
	return dara.Validate(s)
}
