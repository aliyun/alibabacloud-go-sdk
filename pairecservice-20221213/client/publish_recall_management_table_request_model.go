// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRecallManagementTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PublishRecallManagementTableRequest
	GetInstanceId() *string
	SetMode(v string) *PublishRecallManagementTableRequest
	GetMode() *string
	SetPartition(v map[string]*string) *PublishRecallManagementTableRequest
	GetPartition() map[string]*string
	SetPartitions(v map[string]*string) *PublishRecallManagementTableRequest
	GetPartitions() map[string]*string
	SetSkipThresholdCheck(v bool) *PublishRecallManagementTableRequest
	GetSkipThresholdCheck() *bool
	SetStartTimestampMs(v int64) *PublishRecallManagementTableRequest
	GetStartTimestampMs() *int64
}

type PublishRecallManagementTableRequest struct {
	// **The instance ID.**
	//
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// **The enumeration value for the table synchronization mode**:
	//
	// - `Merge`: Merge.
	//
	// example:
	//
	// Merge
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The partition of the table to be published.
	Partition map[string]*string `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The partition of the table to be published.
	Partitions map[string]*string `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
	// **Specifies whether to skip the threshold check when synchronizing the table.**
	//
	// example:
	//
	// false
	SkipThresholdCheck *bool `json:"SkipThresholdCheck,omitempty" xml:"SkipThresholdCheck,omitempty"`
	// The start time of incremental data playback (UNIX milliseconds).
	//
	// example:
	//
	// 1792512000000
	StartTimestampMs *int64 `json:"StartTimestampMs,omitempty" xml:"StartTimestampMs,omitempty"`
}

func (s PublishRecallManagementTableRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishRecallManagementTableRequest) GoString() string {
	return s.String()
}

func (s *PublishRecallManagementTableRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PublishRecallManagementTableRequest) GetMode() *string {
	return s.Mode
}

func (s *PublishRecallManagementTableRequest) GetPartition() map[string]*string {
	return s.Partition
}

func (s *PublishRecallManagementTableRequest) GetPartitions() map[string]*string {
	return s.Partitions
}

func (s *PublishRecallManagementTableRequest) GetSkipThresholdCheck() *bool {
	return s.SkipThresholdCheck
}

func (s *PublishRecallManagementTableRequest) GetStartTimestampMs() *int64 {
	return s.StartTimestampMs
}

func (s *PublishRecallManagementTableRequest) SetInstanceId(v string) *PublishRecallManagementTableRequest {
	s.InstanceId = &v
	return s
}

func (s *PublishRecallManagementTableRequest) SetMode(v string) *PublishRecallManagementTableRequest {
	s.Mode = &v
	return s
}

func (s *PublishRecallManagementTableRequest) SetPartition(v map[string]*string) *PublishRecallManagementTableRequest {
	s.Partition = v
	return s
}

func (s *PublishRecallManagementTableRequest) SetPartitions(v map[string]*string) *PublishRecallManagementTableRequest {
	s.Partitions = v
	return s
}

func (s *PublishRecallManagementTableRequest) SetSkipThresholdCheck(v bool) *PublishRecallManagementTableRequest {
	s.SkipThresholdCheck = &v
	return s
}

func (s *PublishRecallManagementTableRequest) SetStartTimestampMs(v int64) *PublishRecallManagementTableRequest {
	s.StartTimestampMs = &v
	return s
}

func (s *PublishRecallManagementTableRequest) Validate() error {
	return dara.Validate(s)
}
