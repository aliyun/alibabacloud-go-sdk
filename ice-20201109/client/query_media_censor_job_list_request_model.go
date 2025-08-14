// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaCensorJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndOfJobCreatedTimeRange(v string) *QueryMediaCensorJobListRequest
	GetEndOfJobCreatedTimeRange() *string
	SetJobIds(v string) *QueryMediaCensorJobListRequest
	GetJobIds() *string
	SetMaximumPageSize(v int64) *QueryMediaCensorJobListRequest
	GetMaximumPageSize() *int64
	SetNextPageToken(v string) *QueryMediaCensorJobListRequest
	GetNextPageToken() *string
	SetOwnerAccount(v string) *QueryMediaCensorJobListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryMediaCensorJobListRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *QueryMediaCensorJobListRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *QueryMediaCensorJobListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryMediaCensorJobListRequest
	GetResourceOwnerId() *int64
	SetStartOfJobCreatedTimeRange(v string) *QueryMediaCensorJobListRequest
	GetStartOfJobCreatedTimeRange() *string
	SetState(v string) *QueryMediaCensorJobListRequest
	GetState() *string
}

type QueryMediaCensorJobListRequest struct {
	// The end of the time range to query.
	//
	// 	- Specify the time in the ISO 8601 standard. The time must be in UTC.
	//
	// 	- Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2022-02-14T02:16:07Z
	EndOfJobCreatedTimeRange *string `json:"EndOfJobCreatedTimeRange,omitempty" xml:"EndOfJobCreatedTimeRange,omitempty"`
	// The IDs of the content moderation jobs. You can obtain the ID of a content moderation job from the response parameters of the SubmitMediaCensorJob operation. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// fa9c34be3bcf42919ac4d1775239****,78dc866518b843259669df58ed30****
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The number of entries per page.
	//
	// 	- Default value: **30**.
	//
	// 	- Valid values: **1 to 300**.
	//
	// example:
	//
	// 20
	MaximumPageSize *int64 `json:"MaximumPageSize,omitempty" xml:"MaximumPageSize,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// 79aff3eee82242e092899db5f669
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the ApsaraVideo Media Processing (MPS) queue to which the jobs were submitted.
	//
	// example:
	//
	// c5b30b7c0d0e4a0abde1d5f9e751****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query.
	//
	// 	- Specify the time in the ISO 8601 standard. The time must be in UTC.
	//
	// 	- Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2021-12-22T03:48:05Z
	StartOfJobCreatedTimeRange *string `json:"StartOfJobCreatedTimeRange,omitempty" xml:"StartOfJobCreatedTimeRange,omitempty"`
	// The state of the jobs that you want to query. Valid values:
	//
	// 	- **All**: all jobs.
	//
	// 	- **Queuing**: the jobs that are waiting in the queue.
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

func (s QueryMediaCensorJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListRequest) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListRequest) GetEndOfJobCreatedTimeRange() *string {
	return s.EndOfJobCreatedTimeRange
}

func (s *QueryMediaCensorJobListRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *QueryMediaCensorJobListRequest) GetMaximumPageSize() *int64 {
	return s.MaximumPageSize
}

func (s *QueryMediaCensorJobListRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *QueryMediaCensorJobListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryMediaCensorJobListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryMediaCensorJobListRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QueryMediaCensorJobListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryMediaCensorJobListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryMediaCensorJobListRequest) GetStartOfJobCreatedTimeRange() *string {
	return s.StartOfJobCreatedTimeRange
}

func (s *QueryMediaCensorJobListRequest) GetState() *string {
	return s.State
}

func (s *QueryMediaCensorJobListRequest) SetEndOfJobCreatedTimeRange(v string) *QueryMediaCensorJobListRequest {
	s.EndOfJobCreatedTimeRange = &v
	return s
}

func (s *QueryMediaCensorJobListRequest) SetJobIds(v string) *QueryMediaCensorJobListRequest {
	s.JobIds = &v
	return s
}

func (s *QueryMediaCensorJobListRequest) SetMaximumPageSize(v int64) *QueryMediaCensorJobListRequest {
	s.MaximumPageSize = &v
	return s
}

func (s *QueryMediaCensorJobListRequest) SetNextPageToken(v string) *QueryMediaCensorJobListRequest {
	s.NextPageToken = &v
	return s
}

func (s *QueryMediaCensorJobListRequest) SetOwnerAccount(v string) *QueryMediaCensorJobListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryMediaCensorJobListRequest) SetOwnerId(v int64) *QueryMediaCensorJobListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMediaCensorJobListRequest) SetPipelineId(v string) *QueryMediaCensorJobListRequest {
	s.PipelineId = &v
	return s
}

func (s *QueryMediaCensorJobListRequest) SetResourceOwnerAccount(v string) *QueryMediaCensorJobListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMediaCensorJobListRequest) SetResourceOwnerId(v int64) *QueryMediaCensorJobListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMediaCensorJobListRequest) SetStartOfJobCreatedTimeRange(v string) *QueryMediaCensorJobListRequest {
	s.StartOfJobCreatedTimeRange = &v
	return s
}

func (s *QueryMediaCensorJobListRequest) SetState(v string) *QueryMediaCensorJobListRequest {
	s.State = &v
	return s
}

func (s *QueryMediaCensorJobListRequest) Validate() error {
	return dara.Validate(s)
}
