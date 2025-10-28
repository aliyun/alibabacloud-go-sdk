// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoResourceOptimizeRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetAutoResourceOptimizeRulesResponseBody
	GetCode() *int64
	SetData(v *GetAutoResourceOptimizeRulesResponseBodyData) *GetAutoResourceOptimizeRulesResponseBody
	GetData() *GetAutoResourceOptimizeRulesResponseBodyData
	SetMessage(v string) *GetAutoResourceOptimizeRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAutoResourceOptimizeRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAutoResourceOptimizeRulesResponseBody
	GetSuccess() *bool
}

type GetAutoResourceOptimizeRulesResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetAutoResourceOptimizeRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutoResourceOptimizeRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoResourceOptimizeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoResourceOptimizeRulesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetAutoResourceOptimizeRulesResponseBody) GetData() *GetAutoResourceOptimizeRulesResponseBodyData {
	return s.Data
}

func (s *GetAutoResourceOptimizeRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAutoResourceOptimizeRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoResourceOptimizeRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAutoResourceOptimizeRulesResponseBody) SetCode(v int64) *GetAutoResourceOptimizeRulesResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBody) SetData(v *GetAutoResourceOptimizeRulesResponseBodyData) *GetAutoResourceOptimizeRulesResponseBody {
	s.Data = v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBody) SetMessage(v string) *GetAutoResourceOptimizeRulesResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBody) SetRequestId(v string) *GetAutoResourceOptimizeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBody) SetSuccess(v bool) *GetAutoResourceOptimizeRulesResponseBody {
	s.Success = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoResourceOptimizeRulesResponseBodyData struct {
	// The number of database instances for which the automatic fragment recycling feature is currently enabled.
	//
	// example:
	//
	// 1
	EnableAutoResourceOptimizeCount *int64 `json:"EnableAutoResourceOptimizeCount,omitempty" xml:"EnableAutoResourceOptimizeCount,omitempty"`
	// The database instances for which the automatic fragment recycling feature is currently enabled.
	EnableAutoResourceOptimizeList []*GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList `json:"EnableAutoResourceOptimizeList,omitempty" xml:"EnableAutoResourceOptimizeList,omitempty" type:"Repeated"`
	// The number of database instances for which the automatic fragment recycling feature is enabled and DAS Enterprise Edition is disabled.
	//
	// example:
	//
	// 1
	HasEnableRuleButNotDasProCount *int64 `json:"HasEnableRuleButNotDasProCount,omitempty" xml:"HasEnableRuleButNotDasProCount,omitempty"`
	// The database instances for which the automatic fragment recycling feature is enabled and DAS Enterprise Edition is disabled.
	//
	// >  Automatic fragment recycling tasks are run on this type of database instances only if DAS Enterprise Edition is enabled for the database instances again.
	HasEnableRuleButNotDasProList []*GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList `json:"HasEnableRuleButNotDasProList,omitempty" xml:"HasEnableRuleButNotDasProList,omitempty" type:"Repeated"`
	// The number of database instances that do not exist or for which the automatic fragment recycling feature has never been enabled.
	//
	// >  If a database instance does not exist, the instance has been released or the specified instance ID is invalid.
	//
	// example:
	//
	// 1
	NeverEnableAutoResourceOptimizeOrReleasedInstanceCount *int64 `json:"NeverEnableAutoResourceOptimizeOrReleasedInstanceCount,omitempty" xml:"NeverEnableAutoResourceOptimizeOrReleasedInstanceCount,omitempty"`
	// The database instances that do not exist or for which the automatic fragment recycling feature has never been enabled.
	NeverEnableAutoResourceOptimizeOrReleasedInstanceIdList []*string `json:"NeverEnableAutoResourceOptimizeOrReleasedInstanceIdList,omitempty" xml:"NeverEnableAutoResourceOptimizeOrReleasedInstanceIdList,omitempty" type:"Repeated"`
	// The number of database instances for which the automatic fragment recycling feature has been enabled.
	//
	// example:
	//
	// 3
	TotalAutoResourceOptimizeRulesCount *int64 `json:"TotalAutoResourceOptimizeRulesCount,omitempty" xml:"TotalAutoResourceOptimizeRulesCount,omitempty"`
	// The number of database instances for which the automatic fragment recycling feature was once enabled but is currently disabled.
	//
	// example:
	//
	// 1
	TurnOffAutoResourceOptimizeCount *int64 `json:"TurnOffAutoResourceOptimizeCount,omitempty" xml:"TurnOffAutoResourceOptimizeCount,omitempty"`
	// The database instances for which the automatic fragment recycling feature was once enabled but is currently disabled.
	TurnOffAutoResourceOptimizeList []*GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList `json:"TurnOffAutoResourceOptimizeList,omitempty" xml:"TurnOffAutoResourceOptimizeList,omitempty" type:"Repeated"`
}

func (s GetAutoResourceOptimizeRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAutoResourceOptimizeRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) GetEnableAutoResourceOptimizeCount() *int64 {
	return s.EnableAutoResourceOptimizeCount
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) GetEnableAutoResourceOptimizeList() []*GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList {
	return s.EnableAutoResourceOptimizeList
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) GetHasEnableRuleButNotDasProCount() *int64 {
	return s.HasEnableRuleButNotDasProCount
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) GetHasEnableRuleButNotDasProList() []*GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList {
	return s.HasEnableRuleButNotDasProList
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) GetNeverEnableAutoResourceOptimizeOrReleasedInstanceCount() *int64 {
	return s.NeverEnableAutoResourceOptimizeOrReleasedInstanceCount
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) GetNeverEnableAutoResourceOptimizeOrReleasedInstanceIdList() []*string {
	return s.NeverEnableAutoResourceOptimizeOrReleasedInstanceIdList
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) GetTotalAutoResourceOptimizeRulesCount() *int64 {
	return s.TotalAutoResourceOptimizeRulesCount
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) GetTurnOffAutoResourceOptimizeCount() *int64 {
	return s.TurnOffAutoResourceOptimizeCount
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) GetTurnOffAutoResourceOptimizeList() []*GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList {
	return s.TurnOffAutoResourceOptimizeList
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) SetEnableAutoResourceOptimizeCount(v int64) *GetAutoResourceOptimizeRulesResponseBodyData {
	s.EnableAutoResourceOptimizeCount = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) SetEnableAutoResourceOptimizeList(v []*GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) *GetAutoResourceOptimizeRulesResponseBodyData {
	s.EnableAutoResourceOptimizeList = v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) SetHasEnableRuleButNotDasProCount(v int64) *GetAutoResourceOptimizeRulesResponseBodyData {
	s.HasEnableRuleButNotDasProCount = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) SetHasEnableRuleButNotDasProList(v []*GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) *GetAutoResourceOptimizeRulesResponseBodyData {
	s.HasEnableRuleButNotDasProList = v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) SetNeverEnableAutoResourceOptimizeOrReleasedInstanceCount(v int64) *GetAutoResourceOptimizeRulesResponseBodyData {
	s.NeverEnableAutoResourceOptimizeOrReleasedInstanceCount = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) SetNeverEnableAutoResourceOptimizeOrReleasedInstanceIdList(v []*string) *GetAutoResourceOptimizeRulesResponseBodyData {
	s.NeverEnableAutoResourceOptimizeOrReleasedInstanceIdList = v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) SetTotalAutoResourceOptimizeRulesCount(v int64) *GetAutoResourceOptimizeRulesResponseBodyData {
	s.TotalAutoResourceOptimizeRulesCount = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) SetTurnOffAutoResourceOptimizeCount(v int64) *GetAutoResourceOptimizeRulesResponseBodyData {
	s.TurnOffAutoResourceOptimizeCount = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) SetTurnOffAutoResourceOptimizeList(v []*GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) *GetAutoResourceOptimizeRulesResponseBodyData {
	s.TurnOffAutoResourceOptimizeList = v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyData) Validate() error {
	if s.EnableAutoResourceOptimizeList != nil {
		for _, item := range s.EnableAutoResourceOptimizeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HasEnableRuleButNotDasProList != nil {
		for _, item := range s.HasEnableRuleButNotDasProList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TurnOffAutoResourceOptimizeList != nil {
		for _, item := range s.TurnOffAutoResourceOptimizeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList struct {
	// Indicates whether the automatic fragment recycling feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoDefragment *bool `json:"AutoDefragment,omitempty" xml:"AutoDefragment,omitempty"`
	// Indicates whether DAS Enterprise Edition is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DasProOn *bool `json:"DasProOn,omitempty" xml:"DasProOn,omitempty"`
	// The database instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The fragmentation rate of a single physical table for which the automatic fragment recycling feature is enabled.
	//
	// example:
	//
	// 0.2
	TableFragmentationRatio *float64 `json:"TableFragmentationRatio,omitempty" xml:"TableFragmentationRatio,omitempty"`
	// The minimum storage usage of a single physical table for which the automatic fragment recycling feature is enabled. Unit: GB.
	//
	// example:
	//
	// 10
	TableSpaceSize *float64 `json:"TableSpaceSize,omitempty" xml:"TableSpaceSize,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the database instance.
	//
	// example:
	//
	// 140692647406****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) String() string {
	return dara.Prettify(s)
}

func (s GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) GoString() string {
	return s.String()
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) GetAutoDefragment() *bool {
	return s.AutoDefragment
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) GetDasProOn() *bool {
	return s.DasProOn
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) GetTableFragmentationRatio() *float64 {
	return s.TableFragmentationRatio
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) GetTableSpaceSize() *float64 {
	return s.TableSpaceSize
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) GetUserId() *string {
	return s.UserId
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) SetAutoDefragment(v bool) *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList {
	s.AutoDefragment = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) SetDasProOn(v bool) *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList {
	s.DasProOn = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) SetInstanceId(v string) *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList {
	s.InstanceId = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) SetTableFragmentationRatio(v float64) *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList {
	s.TableFragmentationRatio = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) SetTableSpaceSize(v float64) *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList {
	s.TableSpaceSize = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) SetUserId(v string) *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList {
	s.UserId = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataEnableAutoResourceOptimizeList) Validate() error {
	return dara.Validate(s)
}

type GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList struct {
	// Indicates whether the automatic fragment recycling feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoDefragment *bool `json:"AutoDefragment,omitempty" xml:"AutoDefragment,omitempty"`
	// Indicates whether DAS Enterprise Edition is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DasProOn *bool `json:"DasProOn,omitempty" xml:"DasProOn,omitempty"`
	// The database instance ID.
	//
	// example:
	//
	// rm-2ze9xrhze0709****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The fragmentation rate of a single physical table for which the automatic fragment recycling feature is enabled.
	//
	// example:
	//
	// 0.2
	TableFragmentationRatio *float64 `json:"TableFragmentationRatio,omitempty" xml:"TableFragmentationRatio,omitempty"`
	// The minimum storage usage of a single physical table for which the automatic fragment recycling feature is enabled. Unit: GB.
	//
	// example:
	//
	// 10
	TableSpaceSize *float64 `json:"TableSpaceSize,omitempty" xml:"TableSpaceSize,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the database instance.
	//
	// example:
	//
	// 140692647406****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) String() string {
	return dara.Prettify(s)
}

func (s GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) GoString() string {
	return s.String()
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) GetAutoDefragment() *bool {
	return s.AutoDefragment
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) GetDasProOn() *bool {
	return s.DasProOn
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) GetTableFragmentationRatio() *float64 {
	return s.TableFragmentationRatio
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) GetTableSpaceSize() *float64 {
	return s.TableSpaceSize
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) GetUserId() *string {
	return s.UserId
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) SetAutoDefragment(v bool) *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList {
	s.AutoDefragment = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) SetDasProOn(v bool) *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList {
	s.DasProOn = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) SetInstanceId(v string) *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList {
	s.InstanceId = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) SetTableFragmentationRatio(v float64) *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList {
	s.TableFragmentationRatio = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) SetTableSpaceSize(v float64) *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList {
	s.TableSpaceSize = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) SetUserId(v string) *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList {
	s.UserId = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataHasEnableRuleButNotDasProList) Validate() error {
	return dara.Validate(s)
}

type GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList struct {
	// Indicates whether the automatic fragment recycling feature is enabled. Valid values:
	//
	// 	- **true**:
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoDefragment *bool `json:"AutoDefragment,omitempty" xml:"AutoDefragment,omitempty"`
	// Indicates whether DAS Enterprise Edition is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DasProOn *bool `json:"DasProOn,omitempty" xml:"DasProOn,omitempty"`
	// The database instance ID.
	//
	// example:
	//
	// rm-2vc54m2a6pd6p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The fragmentation rate of a single physical table for which the automatic fragment recycling feature is enabled.
	//
	// example:
	//
	// 0.2
	TableFragmentationRatio *float64 `json:"TableFragmentationRatio,omitempty" xml:"TableFragmentationRatio,omitempty"`
	// The minimum storage usage of a single physical table for which the automatic fragment recycling feature is enabled. Unit: GB.
	//
	// example:
	//
	// 10
	TableSpaceSize *float64 `json:"TableSpaceSize,omitempty" xml:"TableSpaceSize,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the database instance.
	//
	// example:
	//
	// 140692647406****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) String() string {
	return dara.Prettify(s)
}

func (s GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) GoString() string {
	return s.String()
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) GetAutoDefragment() *bool {
	return s.AutoDefragment
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) GetDasProOn() *bool {
	return s.DasProOn
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) GetTableFragmentationRatio() *float64 {
	return s.TableFragmentationRatio
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) GetTableSpaceSize() *float64 {
	return s.TableSpaceSize
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) GetUserId() *string {
	return s.UserId
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) SetAutoDefragment(v bool) *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList {
	s.AutoDefragment = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) SetDasProOn(v bool) *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList {
	s.DasProOn = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) SetInstanceId(v string) *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList {
	s.InstanceId = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) SetTableFragmentationRatio(v float64) *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList {
	s.TableFragmentationRatio = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) SetTableSpaceSize(v float64) *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList {
	s.TableSpaceSize = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) SetUserId(v string) *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList {
	s.UserId = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesResponseBodyDataTurnOffAutoResourceOptimizeList) Validate() error {
	return dara.Validate(s)
}
