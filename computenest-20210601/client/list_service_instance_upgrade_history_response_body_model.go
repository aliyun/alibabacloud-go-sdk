// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceUpgradeHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceInstanceUpgradeHistoryResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceInstanceUpgradeHistoryResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceInstanceUpgradeHistoryResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListServiceInstanceUpgradeHistoryResponseBody
	GetTotalCount() *int64
	SetUpgradeHistory(v []*ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) *ListServiceInstanceUpgradeHistoryResponseBody
	GetUpgradeHistory() []*ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory
}

type ListServiceInstanceUpgradeHistoryResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI41
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE3EDF4E-B3B1-19B6-BD01-30D4D00F6E5D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The upgrade history.
	UpgradeHistory []*ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory `json:"UpgradeHistory,omitempty" xml:"UpgradeHistory,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceUpgradeHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceUpgradeHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) GetUpgradeHistory() []*ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	return s.UpgradeHistory
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetMaxResults(v int32) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetNextToken(v string) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetRequestId(v string) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetTotalCount(v int64) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetUpgradeHistory(v []*ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.UpgradeHistory = v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory struct {
	// The time when the update ended.
	//
	// example:
	//
	// 2022-04-26T09:09:51Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The version before the upgrade.
	//
	// example:
	//
	// 1
	FromVersion *string `json:"FromVersion,omitempty" xml:"FromVersion,omitempty"`
	// The upgrade result.
	//
	// example:
	//
	// {\\"PreUpgradeExecutionId\\":\\"exec-123\\"}
	Results *string `json:"Results,omitempty" xml:"Results,omitempty"`
	// The time when the update started.
	//
	// example:
	//
	// 2022-04-26T08:09:51Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the update. Valid values:
	//
	// 	- upgrading: The service instance is being upgraded.
	//
	// 	- UpgradeSuccessful: The service instance is upgraded.
	//
	// 	- UpgradeFailed: The service instance failed to be upgraded.
	//
	// example:
	//
	// UpgradeFailed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version after the upgrade.
	//
	// example:
	//
	// 3
	ToVersion *string `json:"ToVersion,omitempty" xml:"ToVersion,omitempty"`
	// The update type.
	//
	// example:
	//
	// Upgrade
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the upgrade record.
	//
	// example:
	//
	// exec-123
	UpgradeHistoryId *string `json:"UpgradeHistoryId,omitempty" xml:"UpgradeHistoryId,omitempty"`
}

func (s ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) GetEndTime() *string {
	return s.EndTime
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) GetFromVersion() *string {
	return s.FromVersion
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) GetResults() *string {
	return s.Results
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) GetStartTime() *string {
	return s.StartTime
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) GetStatus() *string {
	return s.Status
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) GetToVersion() *string {
	return s.ToVersion
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) GetType() *string {
	return s.Type
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) GetUpgradeHistoryId() *string {
	return s.UpgradeHistoryId
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetEndTime(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.EndTime = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetFromVersion(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.FromVersion = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetResults(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.Results = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetStartTime(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.StartTime = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetStatus(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.Status = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetToVersion(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.ToVersion = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetType(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.Type = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetUpgradeHistoryId(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.UpgradeHistoryId = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) Validate() error {
	return dara.Validate(s)
}
