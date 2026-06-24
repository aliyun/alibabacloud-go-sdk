// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDiagnosisSettingsResponseBody
	GetRequestId() *string
	SetResult(v *DescribeDiagnosisSettingsResponseBodyResult) *DescribeDiagnosisSettingsResponseBody
	GetResult() *DescribeDiagnosisSettingsResponseBodyResult
}

type DescribeDiagnosisSettingsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5E82B8A8-EED7-4557-A6E9-D1AD3E58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *DescribeDiagnosisSettingsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDiagnosisSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosisSettingsResponseBody) GetResult() *DescribeDiagnosisSettingsResponseBodyResult {
	return s.Result
}

func (s *DescribeDiagnosisSettingsResponseBody) SetRequestId(v string) *DescribeDiagnosisSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBody) SetResult(v *DescribeDiagnosisSettingsResponseBodyResult) *DescribeDiagnosisSettingsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDiagnosisSettingsResponseBodyResult struct {
	AuthorizationStatus  *bool   `json:"authorizationStatus,omitempty" xml:"authorizationStatus,omitempty"`
	DailyLimit           *int32  `json:"dailyLimit,omitempty" xml:"dailyLimit,omitempty"`
	DailyScheduleEnabled *bool   `json:"dailyScheduleEnabled,omitempty" xml:"dailyScheduleEnabled,omitempty"`
	DiagnosisMode        *string `json:"diagnosisMode,omitempty" xml:"diagnosisMode,omitempty"`
	// The scenario of intelligent O&M.
	//
	// example:
	//
	// Business Search
	Scene         *string   `json:"scene,omitempty" xml:"scene,omitempty"`
	SelectedItems []*string `json:"selectedItems,omitempty" xml:"selectedItems,omitempty" type:"Repeated"`
	// The timestamp when the intelligent O&M scenario was last updated.
	//
	// example:
	//
	// 1588994035385
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s DescribeDiagnosisSettingsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisSettingsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) GetAuthorizationStatus() *bool {
	return s.AuthorizationStatus
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) GetDailyLimit() *int32 {
	return s.DailyLimit
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) GetDailyScheduleEnabled() *bool {
	return s.DailyScheduleEnabled
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) GetDiagnosisMode() *string {
	return s.DiagnosisMode
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) GetScene() *string {
	return s.Scene
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) GetSelectedItems() []*string {
	return s.SelectedItems
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) SetAuthorizationStatus(v bool) *DescribeDiagnosisSettingsResponseBodyResult {
	s.AuthorizationStatus = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) SetDailyLimit(v int32) *DescribeDiagnosisSettingsResponseBodyResult {
	s.DailyLimit = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) SetDailyScheduleEnabled(v bool) *DescribeDiagnosisSettingsResponseBodyResult {
	s.DailyScheduleEnabled = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) SetDiagnosisMode(v string) *DescribeDiagnosisSettingsResponseBodyResult {
	s.DiagnosisMode = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) SetScene(v string) *DescribeDiagnosisSettingsResponseBodyResult {
	s.Scene = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) SetSelectedItems(v []*string) *DescribeDiagnosisSettingsResponseBodyResult {
	s.SelectedItems = v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) SetUpdateTime(v int64) *DescribeDiagnosisSettingsResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
