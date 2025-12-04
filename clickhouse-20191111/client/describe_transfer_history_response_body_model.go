// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransferHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHistoryDetails(v *DescribeTransferHistoryResponseBodyHistoryDetails) *DescribeTransferHistoryResponseBody
	GetHistoryDetails() *DescribeTransferHistoryResponseBodyHistoryDetails
	SetRequestId(v string) *DescribeTransferHistoryResponseBody
	GetRequestId() *string
}

type DescribeTransferHistoryResponseBody struct {
	// The migration information.
	HistoryDetails *DescribeTransferHistoryResponseBodyHistoryDetails `json:"HistoryDetails,omitempty" xml:"HistoryDetails,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTransferHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransferHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransferHistoryResponseBody) GetHistoryDetails() *DescribeTransferHistoryResponseBodyHistoryDetails {
	return s.HistoryDetails
}

func (s *DescribeTransferHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTransferHistoryResponseBody) SetHistoryDetails(v *DescribeTransferHistoryResponseBodyHistoryDetails) *DescribeTransferHistoryResponseBody {
	s.HistoryDetails = v
	return s
}

func (s *DescribeTransferHistoryResponseBody) SetRequestId(v string) *DescribeTransferHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransferHistoryResponseBody) Validate() error {
	if s.HistoryDetails != nil {
		if err := s.HistoryDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTransferHistoryResponseBodyHistoryDetails struct {
	HistoryDetail []*DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail `json:"HistoryDetail,omitempty" xml:"HistoryDetail,omitempty" type:"Repeated"`
}

func (s DescribeTransferHistoryResponseBodyHistoryDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransferHistoryResponseBodyHistoryDetails) GoString() string {
	return s.String()
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetails) GetHistoryDetail() []*DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	return s.HistoryDetail
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetails) SetHistoryDetail(v []*DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) *DescribeTransferHistoryResponseBodyHistoryDetails {
	s.HistoryDetail = v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetails) Validate() error {
	if s.HistoryDetail != nil {
		for _, item := range s.HistoryDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail struct {
	// The amount of data that is migrated per minute.
	//
	// example:
	//
	// 10240
	BytesPerMinute *int64 `json:"BytesPerMinute,omitempty" xml:"BytesPerMinute,omitempty"`
	// The time window during which write operations are stopped.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 2025-02-08T00:00:00+08:00,2025-02-12T00:00:00+08:00
	DisableWriteWindows *string `json:"DisableWriteWindows,omitempty" xml:"DisableWriteWindows,omitempty"`
	// The number of parts that are migrated per minute.
	//
	// example:
	//
	// 1.24
	PartsPerMinute *float64 `json:"PartsPerMinute,omitempty" xml:"PartsPerMinute,omitempty"`
	// The progress of the data migration.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The control version of the source cluster.
	//
	// example:
	//
	// v1
	SourceControlVersion *string `json:"SourceControlVersion,omitempty" xml:"SourceControlVersion,omitempty"`
	// The ID of the source cluster.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	SourceDBCluster *string `json:"SourceDBCluster,omitempty" xml:"SourceDBCluster,omitempty"`
	// The status of the data migration task. Valid values:
	//
	// 	- **Finished**: The data migration task is complete.
	//
	// 	- **Processing**: The data migration task is in progress.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The running subtask.
	//
	// example:
	//
	// SubJob
	SubJob *string `json:"SubJob,omitempty" xml:"SubJob,omitempty"`
	// The subtask status.
	//
	// example:
	//
	// running
	SubJobStatus *string `json:"SubJobStatus,omitempty" xml:"SubJobStatus,omitempty"`
	// The control version of the destination cluster.
	//
	// example:
	//
	// v2
	TargetControlVersion *string `json:"TargetControlVersion,omitempty" xml:"TargetControlVersion,omitempty"`
	// The ID of the destination cluster.
	//
	// example:
	//
	// cc-bp13zkh9uw523****
	TargetDBCluster *string `json:"TargetDBCluster,omitempty" xml:"TargetDBCluster,omitempty"`
	// The amount of data that is not migrated.
	//
	// example:
	//
	// 102400
	UnsyncedBytes *int64 `json:"UnsyncedBytes,omitempty" xml:"UnsyncedBytes,omitempty"`
	// The number of parts that are not migrated.
	//
	// example:
	//
	// 1000
	UnsyncedParts *int64 `json:"UnsyncedParts,omitempty" xml:"UnsyncedParts,omitempty"`
}

func (s DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GoString() string {
	return s.String()
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetBytesPerMinute() *int64 {
	return s.BytesPerMinute
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetDisableWriteWindows() *string {
	return s.DisableWriteWindows
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetPartsPerMinute() *float64 {
	return s.PartsPerMinute
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetProgress() *string {
	return s.Progress
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetSourceControlVersion() *string {
	return s.SourceControlVersion
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetSourceDBCluster() *string {
	return s.SourceDBCluster
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetStatus() *string {
	return s.Status
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetSubJob() *string {
	return s.SubJob
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetSubJobStatus() *string {
	return s.SubJobStatus
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetTargetControlVersion() *string {
	return s.TargetControlVersion
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetTargetDBCluster() *string {
	return s.TargetDBCluster
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetUnsyncedBytes() *int64 {
	return s.UnsyncedBytes
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) GetUnsyncedParts() *int64 {
	return s.UnsyncedParts
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetBytesPerMinute(v int64) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.BytesPerMinute = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetDisableWriteWindows(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.DisableWriteWindows = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetPartsPerMinute(v float64) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.PartsPerMinute = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetProgress(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.Progress = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetSourceControlVersion(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.SourceControlVersion = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetSourceDBCluster(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.SourceDBCluster = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetStatus(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.Status = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetSubJob(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.SubJob = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetSubJobStatus(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.SubJobStatus = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetTargetControlVersion(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.TargetControlVersion = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetTargetDBCluster(v string) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.TargetDBCluster = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetUnsyncedBytes(v int64) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.UnsyncedBytes = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) SetUnsyncedParts(v int64) *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail {
	s.UnsyncedParts = &v
	return s
}

func (s *DescribeTransferHistoryResponseBodyHistoryDetailsHistoryDetail) Validate() error {
	return dara.Validate(s)
}
