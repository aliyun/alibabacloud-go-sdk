// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleTrialAuthInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetModuleTrialAuthInfoResponseBodyData) *GetModuleTrialAuthInfoResponseBody
	GetData() *GetModuleTrialAuthInfoResponseBodyData
	SetRequestId(v string) *GetModuleTrialAuthInfoResponseBody
	GetRequestId() *string
}

type GetModuleTrialAuthInfoResponseBody struct {
	// The returned data.
	Data *GetModuleTrialAuthInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetModuleTrialAuthInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModuleTrialAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetModuleTrialAuthInfoResponseBody) GetData() *GetModuleTrialAuthInfoResponseBodyData {
	return s.Data
}

func (s *GetModuleTrialAuthInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModuleTrialAuthInfoResponseBody) SetData(v *GetModuleTrialAuthInfoResponseBodyData) *GetModuleTrialAuthInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetModuleTrialAuthInfoResponseBody) SetRequestId(v string) *GetModuleTrialAuthInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModuleTrialAuthInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModuleTrialAuthInfoResponseBodyData struct {
	// Indicates whether the user is qualified for the trial use. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CanTry *bool `json:"CanTry,omitempty" xml:"CanTry,omitempty"`
	// The code of the feature. Valid values:
	//
	// 	- **vulFix**: vulnerability fixing.
	//
	// 	- **cloudSiem**: threat analysis and response.
	//
	// example:
	//
	// vulFix
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The trial use record.
	TrialRecordList []*GetModuleTrialAuthInfoResponseBodyDataTrialRecordList `json:"TrialRecordList,omitempty" xml:"TrialRecordList,omitempty" type:"Repeated"`
}

func (s GetModuleTrialAuthInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetModuleTrialAuthInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModuleTrialAuthInfoResponseBodyData) GetCanTry() *bool {
	return s.CanTry
}

func (s *GetModuleTrialAuthInfoResponseBodyData) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetModuleTrialAuthInfoResponseBodyData) GetTrialRecordList() []*GetModuleTrialAuthInfoResponseBodyDataTrialRecordList {
	return s.TrialRecordList
}

func (s *GetModuleTrialAuthInfoResponseBodyData) SetCanTry(v bool) *GetModuleTrialAuthInfoResponseBodyData {
	s.CanTry = &v
	return s
}

func (s *GetModuleTrialAuthInfoResponseBodyData) SetModuleCode(v string) *GetModuleTrialAuthInfoResponseBodyData {
	s.ModuleCode = &v
	return s
}

func (s *GetModuleTrialAuthInfoResponseBodyData) SetTrialRecordList(v []*GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) *GetModuleTrialAuthInfoResponseBodyData {
	s.TrialRecordList = v
	return s
}

func (s *GetModuleTrialAuthInfoResponseBodyData) Validate() error {
	if s.TrialRecordList != nil {
		for _, item := range s.TrialRecordList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModuleTrialAuthInfoResponseBodyDataTrialRecordList struct {
	// The quota.
	//
	// example:
	//
	// 100
	AuthLimit *int64 `json:"AuthLimit,omitempty" xml:"AuthLimit,omitempty"`
	// The list of quotas. This parameter is available if the value of the ModuleCode parameter is cloudSiem. The value of this parameter consists of the log storage capacity for the threat analysis and response feature and the log data to add. Units: GB and GB-day.
	//
	// example:
	//
	// [1,100]
	AuthLimitList *string `json:"AuthLimitList,omitempty" xml:"AuthLimitList,omitempty"`
	// The end time of the trial use.
	//
	// example:
	//
	// 1679760000000
	GmtEnd *int64 `json:"GmtEnd,omitempty" xml:"GmtEnd,omitempty"`
	// The start time of the trial use.
	//
	// example:
	//
	// 1669824000000
	GmtStart *int64 `json:"GmtStart,omitempty" xml:"GmtStart,omitempty"`
	// The code of the feature. Valid values:
	//
	// 	- **vulFix**: vulnerability fixing.
	//
	// 	- **cloudSiem**: threat analysis and response.
	//
	// example:
	//
	// vulFix
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The status of the trial use. Valid values:
	//
	// 	- **1**: The feature is in trial use.
	//
	// 	- **0**: The trial use ends.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) String() string {
	return dara.Prettify(s)
}

func (s GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) GoString() string {
	return s.String()
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) GetAuthLimit() *int64 {
	return s.AuthLimit
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) GetAuthLimitList() *string {
	return s.AuthLimitList
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) GetGmtEnd() *int64 {
	return s.GmtEnd
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) GetGmtStart() *int64 {
	return s.GmtStart
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) GetStatus() *int32 {
	return s.Status
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) SetAuthLimit(v int64) *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList {
	s.AuthLimit = &v
	return s
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) SetAuthLimitList(v string) *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList {
	s.AuthLimitList = &v
	return s
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) SetGmtEnd(v int64) *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList {
	s.GmtEnd = &v
	return s
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) SetGmtStart(v int64) *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList {
	s.GmtStart = &v
	return s
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) SetModuleCode(v string) *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList {
	s.ModuleCode = &v
	return s
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) SetStatus(v int32) *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList {
	s.Status = &v
	return s
}

func (s *GetModuleTrialAuthInfoResponseBodyDataTrialRecordList) Validate() error {
	return dara.Validate(s)
}
