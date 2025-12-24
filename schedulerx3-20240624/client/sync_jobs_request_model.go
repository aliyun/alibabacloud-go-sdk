// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v []*int64) *SyncJobsRequest
	GetJobIds() []*int64
	SetOriginalAppName(v string) *SyncJobsRequest
	GetOriginalAppName() *string
	SetOriginalClusterId(v string) *SyncJobsRequest
	GetOriginalClusterId() *string
	SetTargetAppName(v string) *SyncJobsRequest
	GetTargetAppName() *string
	SetTargetClusterId(v string) *SyncJobsRequest
	GetTargetClusterId() *string
}

type SyncJobsRequest struct {
	// This parameter is required.
	JobIds []*int64 `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test-source
	OriginalAppName *string `json:"OriginalAppName,omitempty" xml:"OriginalAppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-00ed7f0d742b1
	OriginalClusterId *string `json:"OriginalClusterId,omitempty" xml:"OriginalClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-target
	TargetAppName *string `json:"TargetAppName,omitempty" xml:"TargetAppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-008bfb1541111
	TargetClusterId *string `json:"TargetClusterId,omitempty" xml:"TargetClusterId,omitempty"`
}

func (s SyncJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncJobsRequest) GoString() string {
	return s.String()
}

func (s *SyncJobsRequest) GetJobIds() []*int64 {
	return s.JobIds
}

func (s *SyncJobsRequest) GetOriginalAppName() *string {
	return s.OriginalAppName
}

func (s *SyncJobsRequest) GetOriginalClusterId() *string {
	return s.OriginalClusterId
}

func (s *SyncJobsRequest) GetTargetAppName() *string {
	return s.TargetAppName
}

func (s *SyncJobsRequest) GetTargetClusterId() *string {
	return s.TargetClusterId
}

func (s *SyncJobsRequest) SetJobIds(v []*int64) *SyncJobsRequest {
	s.JobIds = v
	return s
}

func (s *SyncJobsRequest) SetOriginalAppName(v string) *SyncJobsRequest {
	s.OriginalAppName = &v
	return s
}

func (s *SyncJobsRequest) SetOriginalClusterId(v string) *SyncJobsRequest {
	s.OriginalClusterId = &v
	return s
}

func (s *SyncJobsRequest) SetTargetAppName(v string) *SyncJobsRequest {
	s.TargetAppName = &v
	return s
}

func (s *SyncJobsRequest) SetTargetClusterId(v string) *SyncJobsRequest {
	s.TargetClusterId = &v
	return s
}

func (s *SyncJobsRequest) Validate() error {
	return dara.Validate(s)
}
