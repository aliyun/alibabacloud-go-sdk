// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIdsShrink(v string) *SyncJobsShrinkRequest
	GetJobIdsShrink() *string
	SetOriginalAppName(v string) *SyncJobsShrinkRequest
	GetOriginalAppName() *string
	SetOriginalClusterId(v string) *SyncJobsShrinkRequest
	GetOriginalClusterId() *string
	SetTargetAppName(v string) *SyncJobsShrinkRequest
	GetTargetAppName() *string
	SetTargetClusterId(v string) *SyncJobsShrinkRequest
	GetTargetClusterId() *string
}

type SyncJobsShrinkRequest struct {
	// This parameter is required.
	JobIdsShrink *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
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

func (s SyncJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SyncJobsShrinkRequest) GetJobIdsShrink() *string {
	return s.JobIdsShrink
}

func (s *SyncJobsShrinkRequest) GetOriginalAppName() *string {
	return s.OriginalAppName
}

func (s *SyncJobsShrinkRequest) GetOriginalClusterId() *string {
	return s.OriginalClusterId
}

func (s *SyncJobsShrinkRequest) GetTargetAppName() *string {
	return s.TargetAppName
}

func (s *SyncJobsShrinkRequest) GetTargetClusterId() *string {
	return s.TargetClusterId
}

func (s *SyncJobsShrinkRequest) SetJobIdsShrink(v string) *SyncJobsShrinkRequest {
	s.JobIdsShrink = &v
	return s
}

func (s *SyncJobsShrinkRequest) SetOriginalAppName(v string) *SyncJobsShrinkRequest {
	s.OriginalAppName = &v
	return s
}

func (s *SyncJobsShrinkRequest) SetOriginalClusterId(v string) *SyncJobsShrinkRequest {
	s.OriginalClusterId = &v
	return s
}

func (s *SyncJobsShrinkRequest) SetTargetAppName(v string) *SyncJobsShrinkRequest {
	s.TargetAppName = &v
	return s
}

func (s *SyncJobsShrinkRequest) SetTargetClusterId(v string) *SyncJobsShrinkRequest {
	s.TargetClusterId = &v
	return s
}

func (s *SyncJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
