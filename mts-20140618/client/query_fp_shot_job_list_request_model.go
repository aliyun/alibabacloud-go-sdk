// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFpShotJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndOfJobCreatedTimeRange(v string) *QueryFpShotJobListRequest
	GetEndOfJobCreatedTimeRange() *string
	SetJobIds(v string) *QueryFpShotJobListRequest
	GetJobIds() *string
	SetMaximumPageSize(v int64) *QueryFpShotJobListRequest
	GetMaximumPageSize() *int64
	SetNextPageToken(v string) *QueryFpShotJobListRequest
	GetNextPageToken() *string
	SetOwnerAccount(v string) *QueryFpShotJobListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryFpShotJobListRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *QueryFpShotJobListRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *QueryFpShotJobListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryFpShotJobListRequest
	GetResourceOwnerId() *int64
	SetStartOfJobCreatedTimeRange(v string) *QueryFpShotJobListRequest
	GetStartOfJobCreatedTimeRange() *string
	SetState(v string) *QueryFpShotJobListRequest
	GetState() *string
}

type QueryFpShotJobListRequest struct {
	// The end of the time range within which the jobs to be queried were created.
	//
	// 	- Specify the time in the ISO 8601 standard in the
	//
	// 	- YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-02-14T02:16:07Z
	EndOfJobCreatedTimeRange *string `json:"EndOfJobCreatedTimeRange,omitempty" xml:"EndOfJobCreatedTimeRange,omitempty"`
	// The ID of the media fingerprint analysis job that you want to query. To view the job ID, log on to the [ApsaraVideo Media Processing (MPS) console](https://mps.console.aliyun.com/overview), click **Tasks*	- in the left-side navigation pane, and then click the **Video DNA*	- tab on the Tasks page. You can query up to 10 media fingerprint analysis jobs at a time. Separate multiple job IDs with commas (,).
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// 	- Default value: **10**.
	//
	// 	- Valid values: **1 to 100**.
	//
	// example:
	//
	// 10
	MaximumPageSize *int64 `json:"MaximumPageSize,omitempty" xml:"MaximumPageSize,omitempty"`
	// The token that is used to retrieve the next page of the query results. You do not need to specify this parameter in the first request. The response to the first request contains this parameter, which you add to the next request.
	//
	// example:
	//
	// 16f01ad6175e4230ac42bb5182cd****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the MPS queue. To view the ID of the MPS queue, log on to the [MPS console](https://mps.console.aliyun.com/overview) and choose **Global Settings*	- > **Pipelines*	- in the left-side navigation pane.
	//
	// example:
	//
	// b11c171cced04565b1f38f1ecc39****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range within which the jobs to be queried were created.
	//
	// 	- Specify the time in the ISO 8601 standard in the
	//
	// 	- YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-12-22T03:48:05Z
	StartOfJobCreatedTimeRange *string `json:"StartOfJobCreatedTimeRange,omitempty" xml:"StartOfJobCreatedTimeRange,omitempty"`
	// The status of the jobs to be queried. Valid values:
	//
	// 	- **All**: all jobs.
	//
	// 	- **Queuing**: the jobs that are being queued.
	//
	// 	- **Analysing**: the jobs that are in progress.
	//
	// 	- **Fail**: failed jobs.
	//
	// 	- **Success**: successful jobs.
	//
	// example:
	//
	// All
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s QueryFpShotJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListRequest) GetEndOfJobCreatedTimeRange() *string {
	return s.EndOfJobCreatedTimeRange
}

func (s *QueryFpShotJobListRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *QueryFpShotJobListRequest) GetMaximumPageSize() *int64 {
	return s.MaximumPageSize
}

func (s *QueryFpShotJobListRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *QueryFpShotJobListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryFpShotJobListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryFpShotJobListRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QueryFpShotJobListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryFpShotJobListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryFpShotJobListRequest) GetStartOfJobCreatedTimeRange() *string {
	return s.StartOfJobCreatedTimeRange
}

func (s *QueryFpShotJobListRequest) GetState() *string {
	return s.State
}

func (s *QueryFpShotJobListRequest) SetEndOfJobCreatedTimeRange(v string) *QueryFpShotJobListRequest {
	s.EndOfJobCreatedTimeRange = &v
	return s
}

func (s *QueryFpShotJobListRequest) SetJobIds(v string) *QueryFpShotJobListRequest {
	s.JobIds = &v
	return s
}

func (s *QueryFpShotJobListRequest) SetMaximumPageSize(v int64) *QueryFpShotJobListRequest {
	s.MaximumPageSize = &v
	return s
}

func (s *QueryFpShotJobListRequest) SetNextPageToken(v string) *QueryFpShotJobListRequest {
	s.NextPageToken = &v
	return s
}

func (s *QueryFpShotJobListRequest) SetOwnerAccount(v string) *QueryFpShotJobListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryFpShotJobListRequest) SetOwnerId(v int64) *QueryFpShotJobListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryFpShotJobListRequest) SetPipelineId(v string) *QueryFpShotJobListRequest {
	s.PipelineId = &v
	return s
}

func (s *QueryFpShotJobListRequest) SetResourceOwnerAccount(v string) *QueryFpShotJobListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryFpShotJobListRequest) SetResourceOwnerId(v int64) *QueryFpShotJobListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryFpShotJobListRequest) SetStartOfJobCreatedTimeRange(v string) *QueryFpShotJobListRequest {
	s.StartOfJobCreatedTimeRange = &v
	return s
}

func (s *QueryFpShotJobListRequest) SetState(v string) *QueryFpShotJobListRequest {
	s.State = &v
	return s
}

func (s *QueryFpShotJobListRequest) Validate() error {
	return dara.Validate(s)
}
