// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScannerTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScannerTaskId(v string) *DeleteScannerTaskRequest
	GetScannerTaskId() *string
}

type DeleteScannerTaskRequest struct {
	// The unique identifier of the scan task. This is the TaskId returned by CreateTargetScanTask or the ScannerTaskId returned by ListScanTasksByTarget. If the task does not exist or belongs to another tenant, a 400 error is returned without exposing whether the resource exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// task-abc123def4567
	ScannerTaskId *string `json:"ScannerTaskId,omitempty" xml:"ScannerTaskId,omitempty"`
}

func (s DeleteScannerTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScannerTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteScannerTaskRequest) GetScannerTaskId() *string {
	return s.ScannerTaskId
}

func (s *DeleteScannerTaskRequest) SetScannerTaskId(v string) *DeleteScannerTaskRequest {
	s.ScannerTaskId = &v
	return s
}

func (s *DeleteScannerTaskRequest) Validate() error {
	return dara.Validate(s)
}
