// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscribeSmartAccessGatewayDiagnosisReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnoseResult(v *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) *DiscribeSmartAccessGatewayDiagnosisReportResponseBody
	GetDiagnoseResult() *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult
	SetRequestId(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBody
	GetRequestId() *string
}

type DiscribeSmartAccessGatewayDiagnosisReportResponseBody struct {
	// The diagnosis report of the Smart Access Gateway device.
	DiagnoseResult *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult `json:"DiagnoseResult,omitempty" xml:"DiagnoseResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D7D6E3AB-D41A-42E3-8D4E-97B145F4B7C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBody) GoString() string {
	return s.String()
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBody) GetDiagnoseResult() *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	return s.DiagnoseResult
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBody) SetDiagnoseResult(v *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) *DiscribeSmartAccessGatewayDiagnosisReportResponseBody {
	s.DiagnoseResult = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBody) SetRequestId(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBody) Validate() error {
	if s.DiagnoseResult != nil {
		if err := s.DiagnoseResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult struct {
	// The type of the Smart Access Gateway device.
	//
	// - **sag-1000**
	//
	// - **sag-100WM**
	//
	// example:
	//
	// sag-1000
	BoxType *string `json:"BoxType,omitempty" xml:"BoxType,omitempty"`
	// The software version that runs on the Smart Access Gateway device.
	//
	// example:
	//
	// 2.1.0
	BoxVersion *string `json:"BoxVersion,omitempty" xml:"BoxVersion,omitempty"`
	// The list of diagnosis results.
	Details []*DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The diagnosis ID.
	//
	// example:
	//
	// dia-sag42c3t703trh02olv5rf****
	DiagnoseId *string `json:"DiagnoseId,omitempty" xml:"DiagnoseId,omitempty"`
	// The timestamp when the diagnosis ended.
	//
	// example:
	//
	// 160274157
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of completed diagnosis items.
	//
	// example:
	//
	// 15
	FinishedNumber *int32 `json:"FinishedNumber,omitempty" xml:"FinishedNumber,omitempty"`
	// The ID of the Smart Access Gateway instance.
	//
	// example:
	//
	// sag-0nnteglltw6z4b***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The diagnosis result level.
	Level *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel `json:"Level,omitempty" xml:"Level,omitempty" type:"Struct"`
	// The monitoring version used by the Smart Access Gateway device.
	//
	// example:
	//
	// 2.0.2.9
	MonitorVersion *string `json:"MonitorVersion,omitempty" xml:"MonitorVersion,omitempty"`
	// The completion percentage of the diagnosis report.
	//
	// example:
	//
	// 100
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The status of uploading the diagnosis report to SLS.
	//
	// - **0**: The upload failed.
	//
	// - **1**: The upload was successful.
	//
	// example:
	//
	// 0
	ReportSLSSuccess *int32 `json:"ReportSLSSuccess,omitempty" xml:"ReportSLSSuccess,omitempty"`
	// The serial number of the Smart Access Gateway device.
	//
	// example:
	//
	// sag42c3****
	SN *string `json:"SN,omitempty" xml:"SN,omitempty"`
	// The timestamp when the diagnosis started.
	//
	// example:
	//
	// 160274157
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The diagnosis state.
	//
	// - **processing**: The diagnosis is in progress.
	//
	// - **finished**: The diagnosis is successful.
	//
	// - **failed**: The diagnosis failed.
	//
	// - **error**: An error occurred during the diagnosis.
	//
	// - **upload_to_sls_fail**: The diagnosis report failed to be uploaded.
	//
	// example:
	//
	// finished
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The overall statistics of diagnosis item results.
	Statistics *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	// The storage type.
	//
	// Value: **both**, which indicates that the diagnosis report is stored on both the Smart Access Gateway device and in Log Service (SLS).
	//
	// example:
	//
	// both
	StoreType *string `json:"StoreType,omitempty" xml:"StoreType,omitempty"`
	// The total number of diagnosis result entries.
	//
	// example:
	//
	// 15
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
	// The ID of the account to which the Smart Access Gateway instance belongs.
	//
	// example:
	//
	// 1688000000000000
	UId *string `json:"UId,omitempty" xml:"UId,omitempty"`
	// The type of user who initiated the diagnosis. Value: **user**.
	//
	// example:
	//
	// user
	UserLevel *string `json:"UserLevel,omitempty" xml:"UserLevel,omitempty"`
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) String() string {
	return dara.Prettify(s)
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GoString() string {
	return s.String()
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetBoxType() *string {
	return s.BoxType
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetBoxVersion() *string {
	return s.BoxVersion
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetDetails() []*DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails {
	return s.Details
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetDiagnoseId() *string {
	return s.DiagnoseId
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetEndTime() *int32 {
	return s.EndTime
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetFinishedNumber() *int32 {
	return s.FinishedNumber
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetLevel() *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel {
	return s.Level
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetMonitorVersion() *string {
	return s.MonitorVersion
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetPercent() *int32 {
	return s.Percent
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetReportSLSSuccess() *int32 {
	return s.ReportSLSSuccess
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetSN() *string {
	return s.SN
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetStartTime() *int32 {
	return s.StartTime
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetState() *string {
	return s.State
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetStatistics() *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics {
	return s.Statistics
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetStoreType() *string {
	return s.StoreType
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetUId() *string {
	return s.UId
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) GetUserLevel() *string {
	return s.UserLevel
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetBoxType(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.BoxType = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetBoxVersion(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.BoxVersion = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetDetails(v []*DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.Details = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetDiagnoseId(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.DiagnoseId = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetEndTime(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.EndTime = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetFinishedNumber(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.FinishedNumber = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetInstanceId(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.InstanceId = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetLevel(v *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.Level = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetMonitorVersion(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.MonitorVersion = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetPercent(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.Percent = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetReportSLSSuccess(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.ReportSLSSuccess = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetSN(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.SN = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetStartTime(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.StartTime = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetState(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.State = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetStatistics(v *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.Statistics = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetStoreType(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.StoreType = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetTotalNumber(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.TotalNumber = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetUId(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.UId = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) SetUserLevel(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult {
	s.UserLevel = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResult) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Level != nil {
		if err := s.Level.Validate(); err != nil {
			return err
		}
	}
	if s.Statistics != nil {
		if err := s.Statistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails struct {
	// The list of detailed information about diagnosis items.
	Items []*DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The statistics on the severity levels of diagnosis items under the current diagnosis type.
	Statistics *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
	// The diagnosis type.
	//
	// - **config**: SAG configuration.
	//
	// - **internet**: Internet quality.
	//
	// - **biz**: Service quality.
	//
	// example:
	//
	// config
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails) String() string {
	return dara.Prettify(s)
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails) GoString() string {
	return s.String()
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails) GetItems() []*DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems {
	return s.Items
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails) GetStatistics() *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics {
	return s.Statistics
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails) GetType() *string {
	return s.Type
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails) SetItems(v []*DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails {
	s.Items = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails) SetStatistics(v *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails {
	s.Statistics = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails) SetType(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails {
	s.Type = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetails) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Statistics != nil {
		if err := s.Statistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems struct {
	// The diagnosis result in Chinese.
	CN *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN `json:"CN,omitempty" xml:"CN,omitempty" type:"Struct"`
	// The diagnosis result in English.
	EN *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN `json:"EN,omitempty" xml:"EN,omitempty" type:"Struct"`
	// The timestamp when the diagnosis of the diagnosis item ended.
	//
	// example:
	//
	// 1602741570596
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the diagnosis item, which is the unique identifier of the diagnosis item.
	//
	// example:
	//
	// eccConfigCheck
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The severity level of the diagnosis result for the diagnosis item.
	//
	// - **error**: Critical.
	//
	// - **warning**: Warning.
	//
	// - **info**: Normal.
	//
	// example:
	//
	// error
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The timestamp when the diagnosis of the diagnosis item started.
	//
	// example:
	//
	// 1602741570567
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The diagnosis type to which the diagnosis item belongs.
	//
	// - **config**: SAG configuration.
	//
	// - **internet**: Internet quality.
	//
	// - **biz**: Service quality.
	//
	// example:
	//
	// config
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) String() string {
	return dara.Prettify(s)
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) GoString() string {
	return s.String()
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) GetCN() *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN {
	return s.CN
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) GetEN() *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN {
	return s.EN
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) GetItemName() *string {
	return s.ItemName
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) GetLevel() *string {
	return s.Level
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) GetType() *string {
	return s.Type
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) SetCN(v *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems {
	s.CN = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) SetEN(v *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems {
	s.EN = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) SetEndTime(v int64) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems {
	s.EndTime = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) SetItemName(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems {
	s.ItemName = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) SetLevel(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems {
	s.Level = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) SetStartTime(v int64) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems {
	s.StartTime = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) SetType(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems {
	s.Type = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItems) Validate() error {
	if s.CN != nil {
		if err := s.CN.Validate(); err != nil {
			return err
		}
	}
	if s.EN != nil {
		if err := s.EN.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN struct {
	// The diagnosis suggestion.
	//
	// example:
	//
	// Check the leased-line port link wiring first, and then go to the local Leased Line Management page for configuration.
	Advice []*string `json:"Advice,omitempty" xml:"Advice,omitempty" type:"Repeated"`
	// The diagnosis result.
	//
	// example:
	//
	// All leased-line ports have link anomalies or missing port IPs: Port 4.
	Details []*string `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The severity level of the diagnosis result for the diagnosis item.
	//
	// - **严重*	- (Critical): indicates that the issue of the diagnosis item may affect service running. We recommend that you handle the issue at the earliest opportunity.
	//
	// - **警告*	- (Warning): indicates that the diagnosis item has an issue. Handle the issue based on the suggestion.
	//
	// - **正常*	- (Normal): indicates that the diagnosis item is running as expected. No action is required.
	//
	// example:
	//
	// Normal
	ItemLevel *string `json:"ItemLevel,omitempty" xml:"ItemLevel,omitempty"`
	// The name of the diagnosis item.
	//
	// example:
	//
	// Leased line port configuration check
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The diagnosis type to which the diagnosis item belongs.
	//
	// - **配置*	- (Configuration): indicates the **SAG configuration*	- type.
	//
	// - **业务*	- (Service): indicates the **service quality*	- type.
	//
	// - **公网*	- (Internet): indicates the **Internet quality*	- type.
	//
	// example:
	//
	// Configuration
	ItemType *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) String() string {
	return dara.Prettify(s)
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) GoString() string {
	return s.String()
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) GetAdvice() []*string {
	return s.Advice
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) GetDetails() []*string {
	return s.Details
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) GetItemLevel() *string {
	return s.ItemLevel
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) GetItemName() *string {
	return s.ItemName
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) GetItemType() *string {
	return s.ItemType
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) SetAdvice(v []*string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN {
	s.Advice = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) SetDetails(v []*string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN {
	s.Details = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) SetItemLevel(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN {
	s.ItemLevel = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) SetItemName(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN {
	s.ItemName = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) SetItemType(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN {
	s.ItemType = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsCN) Validate() error {
	return dara.Validate(s)
}

type DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN struct {
	// The diagnosis suggestion.
	Advice []*string `json:"Advice,omitempty" xml:"Advice,omitempty" type:"Repeated"`
	// The diagnosis result.
	Details []*string `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The severity level of the diagnosis result for the diagnosis item.
	//
	// - **ERROR**: indicates that the issue of the diagnosis item may affect service running. We recommend that you handle the issue at the earliest opportunity.
	//
	// - **WARNING**: indicates that the diagnosis item has an issue. Handle the issue based on the suggestion.
	//
	// - **INFO**: indicates that the diagnosis item is running as expected. No action is required.
	//
	// example:
	//
	// ERROR
	ItemLevel *string `json:"ItemLevel,omitempty" xml:"ItemLevel,omitempty"`
	// The name of the diagnosis item.
	//
	// example:
	//
	// Express Connect Port Configuration
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The diagnosis type to which the diagnosis item belongs.
	//
	// - **Config**: indicates the **SAG configuration*	- type.
	//
	// - **Service**: indicates the **service quality*	- type.
	//
	// - **Internet**: indicates the **Internet quality*	- type.
	//
	// example:
	//
	// Config
	ItemType *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) String() string {
	return dara.Prettify(s)
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) GoString() string {
	return s.String()
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) GetAdvice() []*string {
	return s.Advice
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) GetDetails() []*string {
	return s.Details
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) GetItemLevel() *string {
	return s.ItemLevel
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) GetItemName() *string {
	return s.ItemName
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) GetItemType() *string {
	return s.ItemType
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) SetAdvice(v []*string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN {
	s.Advice = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) SetDetails(v []*string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN {
	s.Details = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) SetItemLevel(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN {
	s.ItemLevel = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) SetItemName(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN {
	s.ItemName = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) SetItemType(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN {
	s.ItemType = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsItemsEN) Validate() error {
	return dara.Validate(s)
}

type DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics struct {
	// The number of diagnosis items with the **Error*	- severity level.
	//
	// example:
	//
	// 5
	Error *int32 `json:"Error,omitempty" xml:"Error,omitempty"`
	// The number of diagnosis items with the **Info*	- severity level.
	//
	// example:
	//
	// 3
	Info *int32 `json:"Info,omitempty" xml:"Info,omitempty"`
	// The total number of diagnosis items under the current diagnosis type.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The number of diagnosis items with the **Warning*	- severity level.
	//
	// example:
	//
	// 2
	Warning *int32 `json:"Warning,omitempty" xml:"Warning,omitempty"`
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics) String() string {
	return dara.Prettify(s)
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics) GoString() string {
	return s.String()
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics) GetError() *int32 {
	return s.Error
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics) GetInfo() *int32 {
	return s.Info
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics) GetTotal() *int32 {
	return s.Total
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics) GetWarning() *int32 {
	return s.Warning
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics) SetError(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics {
	s.Error = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics) SetInfo(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics {
	s.Info = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics) SetTotal(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics {
	s.Total = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics) SetWarning(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics {
	s.Warning = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultDetailsStatistics) Validate() error {
	return dara.Validate(s)
}

type DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel struct {
	// The diagnosis result level for service quality.
	//
	// example:
	//
	// warning
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The diagnosis result level for SAG configuration.
	//
	// example:
	//
	// info
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The overall diagnosis result level.
	//
	// - **error**: Critical.
	//
	// - **warning**: Warning.
	//
	// - **info**: Normal.
	//
	// example:
	//
	// error
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel) String() string {
	return dara.Prettify(s)
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel) GoString() string {
	return s.String()
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel) GetBiz() *string {
	return s.Biz
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel) GetConfiguration() *string {
	return s.Configuration
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel) GetTotal() *string {
	return s.Total
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel) SetBiz(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel {
	s.Biz = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel) SetConfiguration(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel {
	s.Configuration = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel) SetTotal(v string) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel {
	s.Total = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultLevel) Validate() error {
	return dara.Validate(s)
}

type DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics struct {
	// The total number of diagnosis items with the **Error*	- severity level.
	//
	// example:
	//
	// 2
	Error *int32 `json:"Error,omitempty" xml:"Error,omitempty"`
	// The total number of diagnosis items with the **Info*	- severity level.
	//
	// example:
	//
	// 5
	Info *int32 `json:"Info,omitempty" xml:"Info,omitempty"`
	// The total number of all diagnosis items.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The total number of diagnosis items with the **Warning*	- severity level.
	//
	// example:
	//
	// 3
	Warning *int32 `json:"Warning,omitempty" xml:"Warning,omitempty"`
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics) String() string {
	return dara.Prettify(s)
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics) GoString() string {
	return s.String()
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics) GetError() *int32 {
	return s.Error
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics) GetInfo() *int32 {
	return s.Info
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics) GetTotal() *int32 {
	return s.Total
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics) GetWarning() *int32 {
	return s.Warning
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics) SetError(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics {
	s.Error = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics) SetInfo(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics {
	s.Info = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics) SetTotal(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics {
	s.Total = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics) SetWarning(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics {
	s.Warning = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponseBodyDiagnoseResultStatistics) Validate() error {
	return dara.Validate(s)
}
