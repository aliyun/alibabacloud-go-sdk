// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndOfJobCreatedTimeRange(v string) *ListJobRequest
	GetEndOfJobCreatedTimeRange() *string
	SetMaximumPageSize(v int64) *ListJobRequest
	GetMaximumPageSize() *int64
	SetNextPageToken(v string) *ListJobRequest
	GetNextPageToken() *string
	SetOwnerAccount(v string) *ListJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListJobRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *ListJobRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *ListJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListJobRequest
	GetResourceOwnerId() *int64
	SetStartOfJobCreatedTimeRange(v string) *ListJobRequest
	GetStartOfJobCreatedTimeRange() *string
	SetState(v string) *ListJobRequest
	GetState() *string
}

type ListJobRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2014-01-11T12:00:00Z
	EndOfJobCreatedTimeRange *string `json:"EndOfJobCreatedTimeRange,omitempty" xml:"EndOfJobCreatedTimeRange,omitempty"`
	// The number of entries per page.
	//
	// 	- Default value: **10**.
	//
	// 	- Valid values: **1 to 100**.
	//
	// example:
	//
	// 10
	MaximumPageSize *int64 `json:"MaximumPageSize,omitempty" xml:"MaximumPageSize,omitempty"`
	// The token that is used to retrieve the next page of the query results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextPageToken.
	//
	// example:
	//
	// 16f01ad6175e4230ac42bb5182cd****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the MPS queue to which the job is submitted. To obtain the ID of an MPS queue, you can log on to the [MPS console](https://mps.console.aliyun.com/overview) and choose **Global Settings*	- > **MPS Queue and Callback*	- in the left-side navigation pane.
	//
	// example:
	//
	// 88c6ca184c0e424d5w5b665e2a12****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2014-01-10T12:00:00Z
	StartOfJobCreatedTimeRange *string `json:"StartOfJobCreatedTimeRange,omitempty" xml:"StartOfJobCreatedTimeRange,omitempty"`
	// The state of the transcoding job. Default value: **All**. Valid values:
	//
	// 	- **All**
	//
	// 	- **Submitted**
	//
	// 	- **Transcoding**
	//
	// 	- **TranscodeSuccess**
	//
	// 	- **TranscodeFail**
	//
	// 	- **TranscodeCancelled**
	//
	// example:
	//
	// All
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobRequest) GoString() string {
	return s.String()
}

func (s *ListJobRequest) GetEndOfJobCreatedTimeRange() *string {
	return s.EndOfJobCreatedTimeRange
}

func (s *ListJobRequest) GetMaximumPageSize() *int64 {
	return s.MaximumPageSize
}

func (s *ListJobRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListJobRequest) GetStartOfJobCreatedTimeRange() *string {
	return s.StartOfJobCreatedTimeRange
}

func (s *ListJobRequest) GetState() *string {
	return s.State
}

func (s *ListJobRequest) SetEndOfJobCreatedTimeRange(v string) *ListJobRequest {
	s.EndOfJobCreatedTimeRange = &v
	return s
}

func (s *ListJobRequest) SetMaximumPageSize(v int64) *ListJobRequest {
	s.MaximumPageSize = &v
	return s
}

func (s *ListJobRequest) SetNextPageToken(v string) *ListJobRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListJobRequest) SetOwnerAccount(v string) *ListJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListJobRequest) SetOwnerId(v int64) *ListJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListJobRequest) SetPipelineId(v string) *ListJobRequest {
	s.PipelineId = &v
	return s
}

func (s *ListJobRequest) SetResourceOwnerAccount(v string) *ListJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListJobRequest) SetResourceOwnerId(v int64) *ListJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListJobRequest) SetStartOfJobCreatedTimeRange(v string) *ListJobRequest {
	s.StartOfJobCreatedTimeRange = &v
	return s
}

func (s *ListJobRequest) SetState(v string) *ListJobRequest {
	s.State = &v
	return s
}

func (s *ListJobRequest) Validate() error {
	return dara.Validate(s)
}
