// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySnapshotJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndOfJobCreatedTimeRange(v string) *QuerySnapshotJobListRequest
	GetEndOfJobCreatedTimeRange() *string
	SetMaximumPageSize(v int64) *QuerySnapshotJobListRequest
	GetMaximumPageSize() *int64
	SetNextPageToken(v string) *QuerySnapshotJobListRequest
	GetNextPageToken() *string
	SetOwnerAccount(v string) *QuerySnapshotJobListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QuerySnapshotJobListRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *QuerySnapshotJobListRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *QuerySnapshotJobListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySnapshotJobListRequest
	GetResourceOwnerId() *int64
	SetSnapshotJobIds(v string) *QuerySnapshotJobListRequest
	GetSnapshotJobIds() *string
	SetStartOfJobCreatedTimeRange(v string) *QuerySnapshotJobListRequest
	GetStartOfJobCreatedTimeRange() *string
	SetState(v string) *QuerySnapshotJobListRequest
	GetState() *string
}

type QuerySnapshotJobListRequest struct {
	// The snapshot configuration.
	//
	// example:
	//
	// 2014-01-12T12:00:00Z
	EndOfJobCreatedTimeRange *string `json:"EndOfJobCreatedTimeRange,omitempty" xml:"EndOfJobCreatedTimeRange,omitempty"`
	// The ID of the MPS queue to which the snapshot jobs that you want to query are submitted. To obtain the ID, you can log on to the **MPS console*	- and choose **Global Settings*	- > **Pipelines*	- in the left-side navigation pane.
	//
	// example:
	//
	// 30
	MaximumPageSize *int64 `json:"MaximumPageSize,omitempty" xml:"MaximumPageSize,omitempty"`
	// The end of the time range within which the creation time of snapshot jobs to be queried is.
	//
	// 	- Specify the time in the ISO 8601 standard in the
	//
	// 	- YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// cc6cbef8e8d5481ca536f5d2a466****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start time for taking snapshots. Unit: milliseconds.
	//
	// example:
	//
	// b11c171cced04565b1f38f1ecc39****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range within which the creation time of snapshot jobs to be queried is.
	//
	// 	- Specify the time in the ISO 8601 standard in the
	//
	// 	- YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 72dfa5e679ab4be9a3ed9974c736****
	SnapshotJobIds *string `json:"SnapshotJobIds,omitempty" xml:"SnapshotJobIds,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2014-01-10T12:00:00Z
	StartOfJobCreatedTimeRange *string `json:"StartOfJobCreatedTimeRange,omitempty" xml:"StartOfJobCreatedTimeRange,omitempty"`
	// The information about the snapshot jobs.
	//
	// example:
	//
	// Snapshoting
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s QuerySnapshotJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListRequest) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListRequest) GetEndOfJobCreatedTimeRange() *string {
	return s.EndOfJobCreatedTimeRange
}

func (s *QuerySnapshotJobListRequest) GetMaximumPageSize() *int64 {
	return s.MaximumPageSize
}

func (s *QuerySnapshotJobListRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *QuerySnapshotJobListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QuerySnapshotJobListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySnapshotJobListRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QuerySnapshotJobListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySnapshotJobListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySnapshotJobListRequest) GetSnapshotJobIds() *string {
	return s.SnapshotJobIds
}

func (s *QuerySnapshotJobListRequest) GetStartOfJobCreatedTimeRange() *string {
	return s.StartOfJobCreatedTimeRange
}

func (s *QuerySnapshotJobListRequest) GetState() *string {
	return s.State
}

func (s *QuerySnapshotJobListRequest) SetEndOfJobCreatedTimeRange(v string) *QuerySnapshotJobListRequest {
	s.EndOfJobCreatedTimeRange = &v
	return s
}

func (s *QuerySnapshotJobListRequest) SetMaximumPageSize(v int64) *QuerySnapshotJobListRequest {
	s.MaximumPageSize = &v
	return s
}

func (s *QuerySnapshotJobListRequest) SetNextPageToken(v string) *QuerySnapshotJobListRequest {
	s.NextPageToken = &v
	return s
}

func (s *QuerySnapshotJobListRequest) SetOwnerAccount(v string) *QuerySnapshotJobListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QuerySnapshotJobListRequest) SetOwnerId(v int64) *QuerySnapshotJobListRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySnapshotJobListRequest) SetPipelineId(v string) *QuerySnapshotJobListRequest {
	s.PipelineId = &v
	return s
}

func (s *QuerySnapshotJobListRequest) SetResourceOwnerAccount(v string) *QuerySnapshotJobListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySnapshotJobListRequest) SetResourceOwnerId(v int64) *QuerySnapshotJobListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySnapshotJobListRequest) SetSnapshotJobIds(v string) *QuerySnapshotJobListRequest {
	s.SnapshotJobIds = &v
	return s
}

func (s *QuerySnapshotJobListRequest) SetStartOfJobCreatedTimeRange(v string) *QuerySnapshotJobListRequest {
	s.StartOfJobCreatedTimeRange = &v
	return s
}

func (s *QuerySnapshotJobListRequest) SetState(v string) *QuerySnapshotJobListRequest {
	s.State = &v
	return s
}

func (s *QuerySnapshotJobListRequest) Validate() error {
	return dara.Validate(s)
}
