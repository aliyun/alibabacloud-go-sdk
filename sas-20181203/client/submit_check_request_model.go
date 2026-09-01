// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceDirectoryAccountId(v int64) *SubmitCheckRequest
	GetResourceDirectoryAccountId() *int64
	SetScanRange(v string) *SubmitCheckRequest
	GetScanRange() *string
	SetTaskSource(v string) *SubmitCheckRequest
	GetTaskSource() *string
}

type SubmitCheckRequest struct {
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The scan range. Valid values:
	//
	// - **FULL**: scans all check items
	//
	// - **POLICY**: scans custom-configured check items
	//
	// example:
	//
	// POLICY
	ScanRange *string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty"`
	// The task source. Valid values:
	//
	// - **YAO_CHI**: Alibaba Cloud ApsaraDB console.
	//
	// example:
	//
	// YAO_CHI
	TaskSource *string `json:"TaskSource,omitempty" xml:"TaskSource,omitempty"`
}

func (s SubmitCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCheckRequest) GoString() string {
	return s.String()
}

func (s *SubmitCheckRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *SubmitCheckRequest) GetScanRange() *string {
	return s.ScanRange
}

func (s *SubmitCheckRequest) GetTaskSource() *string {
	return s.TaskSource
}

func (s *SubmitCheckRequest) SetResourceDirectoryAccountId(v int64) *SubmitCheckRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *SubmitCheckRequest) SetScanRange(v string) *SubmitCheckRequest {
	s.ScanRange = &v
	return s
}

func (s *SubmitCheckRequest) SetTaskSource(v string) *SubmitCheckRequest {
	s.TaskSource = &v
	return s
}

func (s *SubmitCheckRequest) Validate() error {
	return dara.Validate(s)
}
