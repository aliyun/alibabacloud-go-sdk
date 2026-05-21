// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpgradeStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetUpgradeStatusResponseBodyData) *GetUpgradeStatusResponseBody
	GetData() []*GetUpgradeStatusResponseBodyData
	SetErrorCode(v string) *GetUpgradeStatusResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetUpgradeStatusResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *GetUpgradeStatusResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *GetUpgradeStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUpgradeStatusResponseBody
	GetSuccess() *bool
}

type GetUpgradeStatusResponseBody struct {
	Data []*GetUpgradeStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D73E42D0-AA72-5880-B96F-548B43C84736
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUpgradeStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUpgradeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusResponseBody) GetData() []*GetUpgradeStatusResponseBodyData {
	return s.Data
}

func (s *GetUpgradeStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetUpgradeStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetUpgradeStatusResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetUpgradeStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUpgradeStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUpgradeStatusResponseBody) SetData(v []*GetUpgradeStatusResponseBodyData) *GetUpgradeStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetErrorCode(v string) *GetUpgradeStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetErrorMessage(v string) *GetUpgradeStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetHttpStatusCode(v string) *GetUpgradeStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetRequestId(v string) *GetUpgradeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) SetSuccess(v bool) *GetUpgradeStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetUpgradeStatusResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUpgradeStatusResponseBodyData struct {
	// example:
	//
	// {"UpgradePhase":"Upgrade","StatusCode":"LeaderRestartFailed","FlightingReport":{"SQLResultStat":{"TotalExecuted":9000"Exceptions":0,"Incorrect":1,"Slow":1,"DQLSpeedup":2}"SegmentFormatStat":[{"Database":"test_db","TableUsingSegmentCount":5},{"Database":"user_order","TableUsingSegmentCount":7}],"EstimatedUpgradeTime":{"StopInstance":30,"BackupData":120,"UpgradeInstance":600},"ReadPermissionCheckStat":[{"db":"tst0","user":"v4_300433463265624129","schema":"dim_db","table":"it_rpt_org_tree_info_partition_all"},{"db":"tst1","user":"v4_300433463265624129","schema":"dim_db","table":"it_rpt_org_tree_info_partition_all_2"},]},"UpgradingSteps":{"Stop":{"Status":"Success","StartTime":"2023-05-09T06:48:28.843Z","StopTime":"2023-05-09T06:48:28.843Z"},"Backup":{"Status":"Success","StartTime":"2023-05-09T06:48:28.843Z","StopTime":"2023-05-09T06:48:28.843Z"},"DoUpgrade":{"Status":"Success","StartTime":"2023-05-09T06:48:28.843Z","StopTime":"2023-05-09T06:48:28.843Z"},"Rollback":{"Status":"Success","StartTime":"2023-05-09T06:48:28.843Z","StopTime":"2023-05-09T06:48:28.843Z"}},"Instances":{"LeaderInstanceId":"hgxxx","LeaderInstanceName":"ERP instance","LeaderInstanceStatus":"Running","FollowerInstances":[{"InstanceId":"hgxxx","InstanceName":"TMSinstance","Status":"Running"},{"InstanceId":"hxxx""InstanceName":"WMS readonly","Status":"Unavailable"}]}}
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// 1.3.23
	FromVersion *string `json:"FromVersion,omitempty" xml:"FromVersion,omitempty"`
	// example:
	//
	// 2023-05-09T06:48:28.843Z
	PrepareFinishTime *string `json:"PrepareFinishTime,omitempty" xml:"PrepareFinishTime,omitempty"`
	// example:
	//
	// archived
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1.3.24
	ToVersion *string `json:"ToVersion,omitempty" xml:"ToVersion,omitempty"`
}

func (s GetUpgradeStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUpgradeStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUpgradeStatusResponseBodyData) GetDetails() *string {
	return s.Details
}

func (s *GetUpgradeStatusResponseBodyData) GetFromVersion() *string {
	return s.FromVersion
}

func (s *GetUpgradeStatusResponseBodyData) GetPrepareFinishTime() *string {
	return s.PrepareFinishTime
}

func (s *GetUpgradeStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetUpgradeStatusResponseBodyData) GetToVersion() *string {
	return s.ToVersion
}

func (s *GetUpgradeStatusResponseBodyData) SetDetails(v string) *GetUpgradeStatusResponseBodyData {
	s.Details = &v
	return s
}

func (s *GetUpgradeStatusResponseBodyData) SetFromVersion(v string) *GetUpgradeStatusResponseBodyData {
	s.FromVersion = &v
	return s
}

func (s *GetUpgradeStatusResponseBodyData) SetPrepareFinishTime(v string) *GetUpgradeStatusResponseBodyData {
	s.PrepareFinishTime = &v
	return s
}

func (s *GetUpgradeStatusResponseBodyData) SetStatus(v string) *GetUpgradeStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetUpgradeStatusResponseBodyData) SetToVersion(v string) *GetUpgradeStatusResponseBodyData {
	s.ToVersion = &v
	return s
}

func (s *GetUpgradeStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
