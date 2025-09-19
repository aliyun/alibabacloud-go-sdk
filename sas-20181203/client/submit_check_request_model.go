// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScanRange(v string) *SubmitCheckRequest
	GetScanRange() *string
	SetTaskSource(v string) *SubmitCheckRequest
	GetTaskSource() *string
}

type SubmitCheckRequest struct {
	// The check items that are scanned. Valid values:
	//
	// 	- **FULL**: All check items are scanned.
	//
	// 	- **FULL**: Only the check items that are configured are scanned.
	//
	// example:
	//
	// POLICY
	ScanRange  *string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty"`
	TaskSource *string `json:"TaskSource,omitempty" xml:"TaskSource,omitempty"`
}

func (s SubmitCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCheckRequest) GoString() string {
	return s.String()
}

func (s *SubmitCheckRequest) GetScanRange() *string {
	return s.ScanRange
}

func (s *SubmitCheckRequest) GetTaskSource() *string {
	return s.TaskSource
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
