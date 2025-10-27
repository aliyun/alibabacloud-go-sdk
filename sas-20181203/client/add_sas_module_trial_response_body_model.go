// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSasModuleTrialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddSasModuleTrialResponseBodyData) *AddSasModuleTrialResponseBody
	GetData() *AddSasModuleTrialResponseBodyData
	SetRequestId(v string) *AddSasModuleTrialResponseBody
	GetRequestId() *string
}

type AddSasModuleTrialResponseBody struct {
	// The data returned.
	Data *AddSasModuleTrialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 09969D2C-4FAD-429E-BFBF-9A60DEF8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSasModuleTrialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSasModuleTrialResponseBody) GoString() string {
	return s.String()
}

func (s *AddSasModuleTrialResponseBody) GetData() *AddSasModuleTrialResponseBodyData {
	return s.Data
}

func (s *AddSasModuleTrialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSasModuleTrialResponseBody) SetData(v *AddSasModuleTrialResponseBodyData) *AddSasModuleTrialResponseBody {
	s.Data = v
	return s
}

func (s *AddSasModuleTrialResponseBody) SetRequestId(v string) *AddSasModuleTrialResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSasModuleTrialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddSasModuleTrialResponseBodyData struct {
	// The information about the trial use.
	TrialRecordList []*AddSasModuleTrialResponseBodyDataTrialRecordList `json:"TrialRecordList,omitempty" xml:"TrialRecordList,omitempty" type:"Repeated"`
}

func (s AddSasModuleTrialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddSasModuleTrialResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddSasModuleTrialResponseBodyData) GetTrialRecordList() []*AddSasModuleTrialResponseBodyDataTrialRecordList {
	return s.TrialRecordList
}

func (s *AddSasModuleTrialResponseBodyData) SetTrialRecordList(v []*AddSasModuleTrialResponseBodyDataTrialRecordList) *AddSasModuleTrialResponseBodyData {
	s.TrialRecordList = v
	return s
}

func (s *AddSasModuleTrialResponseBodyData) Validate() error {
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

type AddSasModuleTrialResponseBodyDataTrialRecordList struct {
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
	// 1638201599999
	GmtEnd *int64 `json:"GmtEnd,omitempty" xml:"GmtEnd,omitempty"`
	// The start time of the trial use.
	//
	// example:
	//
	// 1667232000000
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

func (s AddSasModuleTrialResponseBodyDataTrialRecordList) String() string {
	return dara.Prettify(s)
}

func (s AddSasModuleTrialResponseBodyDataTrialRecordList) GoString() string {
	return s.String()
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) GetAuthLimit() *int64 {
	return s.AuthLimit
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) GetAuthLimitList() *string {
	return s.AuthLimitList
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) GetGmtEnd() *int64 {
	return s.GmtEnd
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) GetGmtStart() *int64 {
	return s.GmtStart
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) GetStatus() *int32 {
	return s.Status
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) SetAuthLimit(v int64) *AddSasModuleTrialResponseBodyDataTrialRecordList {
	s.AuthLimit = &v
	return s
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) SetAuthLimitList(v string) *AddSasModuleTrialResponseBodyDataTrialRecordList {
	s.AuthLimitList = &v
	return s
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) SetGmtEnd(v int64) *AddSasModuleTrialResponseBodyDataTrialRecordList {
	s.GmtEnd = &v
	return s
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) SetGmtStart(v int64) *AddSasModuleTrialResponseBodyDataTrialRecordList {
	s.GmtStart = &v
	return s
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) SetModuleCode(v string) *AddSasModuleTrialResponseBodyDataTrialRecordList {
	s.ModuleCode = &v
	return s
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) SetStatus(v int32) *AddSasModuleTrialResponseBodyDataTrialRecordList {
	s.Status = &v
	return s
}

func (s *AddSasModuleTrialResponseBodyDataTrialRecordList) Validate() error {
	return dara.Validate(s)
}
