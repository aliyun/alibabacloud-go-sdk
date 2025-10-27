// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentVersionPublishResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetCurrentVersionPublishResponseBodyData) *GetCurrentVersionPublishResponseBody
	GetData() *GetCurrentVersionPublishResponseBodyData
	SetRequestId(v string) *GetCurrentVersionPublishResponseBody
	GetRequestId() *string
}

type GetCurrentVersionPublishResponseBody struct {
	// The data returned.
	Data *GetCurrentVersionPublishResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1383B0DB-D5D6-4B0C-9E6B-75939C8E67FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCurrentVersionPublishResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentVersionPublishResponseBody) GoString() string {
	return s.String()
}

func (s *GetCurrentVersionPublishResponseBody) GetData() *GetCurrentVersionPublishResponseBodyData {
	return s.Data
}

func (s *GetCurrentVersionPublishResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCurrentVersionPublishResponseBody) SetData(v *GetCurrentVersionPublishResponseBodyData) *GetCurrentVersionPublishResponseBody {
	s.Data = v
	return s
}

func (s *GetCurrentVersionPublishResponseBody) SetRequestId(v string) *GetCurrentVersionPublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCurrentVersionPublishResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCurrentVersionPublishResponseBodyData struct {
	// Indicates whether automatic upgrade is enabled. Valid values:
	//
	// 	- **1**: yes.
	//
	// 	- **0**: no.
	//
	// example:
	//
	// 1
	AutoUpgrade *int32 `json:"AutoUpgrade,omitempty" xml:"AutoUpgrade,omitempty"`
	// Indicates whether you can enable custom upgrade for the Security Center agent. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	BigCustomer *bool `json:"BigCustomer,omitempty" xml:"BigCustomer,omitempty"`
	// The version of the Security Center agent.
	//
	// example:
	//
	// 0.0.8
	CurVersion *string `json:"CurVersion,omitempty" xml:"CurVersion,omitempty"`
	// The timestamp when the Security Center agent was forcibly upgraded.
	//
	// example:
	//
	// 1732506308000
	ForceUpgradeTime *int64 `json:"ForceUpgradeTime,omitempty" xml:"ForceUpgradeTime,omitempty"`
	// Indicates whether the canary release policy is enabled. Valid values:
	//
	// 	- **1**: yes.
	//
	// 	- .**0**: no.
	//
	// example:
	//
	// 1
	GraySwitchStatus *int32 `json:"GraySwitchStatus,omitempty" xml:"GraySwitchStatus,omitempty"`
	// The latest version of the Security Center agent.
	//
	// example:
	//
	// 0.0.9
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The timestamp when the latest version of the Security Center agent was created.
	//
	// example:
	//
	// 1662639150000
	LatestVersionCreate *int64 `json:"LatestVersionCreate,omitempty" xml:"LatestVersionCreate,omitempty"`
	// The description of about the latest version.
	//
	// example:
	//
	// test
	LatestVersionDesc *string `json:"LatestVersionDesc,omitempty" xml:"LatestVersionDesc,omitempty"`
	// The publish status of the Security Center agent. Valid values:
	//
	// 	- **0**: not started.
	//
	// 	- **1**: publishing.
	//
	// 	- **2**: published.
	//
	// 	- **3**: publish suspended.
	//
	// 	- **4**: forcibly upgrading.
	//
	// example:
	//
	// 1
	PublishStatus *int32 `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// The destination version of the Security Center agent.
	//
	// example:
	//
	// 0.0.9
	UpgradeVersion *string `json:"UpgradeVersion,omitempty" xml:"UpgradeVersion,omitempty"`
}

func (s GetCurrentVersionPublishResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentVersionPublishResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCurrentVersionPublishResponseBodyData) GetAutoUpgrade() *int32 {
	return s.AutoUpgrade
}

func (s *GetCurrentVersionPublishResponseBodyData) GetBigCustomer() *bool {
	return s.BigCustomer
}

func (s *GetCurrentVersionPublishResponseBodyData) GetCurVersion() *string {
	return s.CurVersion
}

func (s *GetCurrentVersionPublishResponseBodyData) GetForceUpgradeTime() *int64 {
	return s.ForceUpgradeTime
}

func (s *GetCurrentVersionPublishResponseBodyData) GetGraySwitchStatus() *int32 {
	return s.GraySwitchStatus
}

func (s *GetCurrentVersionPublishResponseBodyData) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *GetCurrentVersionPublishResponseBodyData) GetLatestVersionCreate() *int64 {
	return s.LatestVersionCreate
}

func (s *GetCurrentVersionPublishResponseBodyData) GetLatestVersionDesc() *string {
	return s.LatestVersionDesc
}

func (s *GetCurrentVersionPublishResponseBodyData) GetPublishStatus() *int32 {
	return s.PublishStatus
}

func (s *GetCurrentVersionPublishResponseBodyData) GetUpgradeVersion() *string {
	return s.UpgradeVersion
}

func (s *GetCurrentVersionPublishResponseBodyData) SetAutoUpgrade(v int32) *GetCurrentVersionPublishResponseBodyData {
	s.AutoUpgrade = &v
	return s
}

func (s *GetCurrentVersionPublishResponseBodyData) SetBigCustomer(v bool) *GetCurrentVersionPublishResponseBodyData {
	s.BigCustomer = &v
	return s
}

func (s *GetCurrentVersionPublishResponseBodyData) SetCurVersion(v string) *GetCurrentVersionPublishResponseBodyData {
	s.CurVersion = &v
	return s
}

func (s *GetCurrentVersionPublishResponseBodyData) SetForceUpgradeTime(v int64) *GetCurrentVersionPublishResponseBodyData {
	s.ForceUpgradeTime = &v
	return s
}

func (s *GetCurrentVersionPublishResponseBodyData) SetGraySwitchStatus(v int32) *GetCurrentVersionPublishResponseBodyData {
	s.GraySwitchStatus = &v
	return s
}

func (s *GetCurrentVersionPublishResponseBodyData) SetLatestVersion(v string) *GetCurrentVersionPublishResponseBodyData {
	s.LatestVersion = &v
	return s
}

func (s *GetCurrentVersionPublishResponseBodyData) SetLatestVersionCreate(v int64) *GetCurrentVersionPublishResponseBodyData {
	s.LatestVersionCreate = &v
	return s
}

func (s *GetCurrentVersionPublishResponseBodyData) SetLatestVersionDesc(v string) *GetCurrentVersionPublishResponseBodyData {
	s.LatestVersionDesc = &v
	return s
}

func (s *GetCurrentVersionPublishResponseBodyData) SetPublishStatus(v int32) *GetCurrentVersionPublishResponseBodyData {
	s.PublishStatus = &v
	return s
}

func (s *GetCurrentVersionPublishResponseBodyData) SetUpgradeVersion(v string) *GetCurrentVersionPublishResponseBodyData {
	s.UpgradeVersion = &v
	return s
}

func (s *GetCurrentVersionPublishResponseBodyData) Validate() error {
	return dara.Validate(s)
}
