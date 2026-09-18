// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScannerTaskResultHtmlUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScannerTaskId(v string) *GetScannerTaskResultHtmlUrlRequest
	GetScannerTaskId() *string
}

type GetScannerTaskResultHtmlUrlRequest struct {
	// The unique identifier of the scan task. This is the TaskId returned by CreateTargetScanTask or the ScannerTaskId returned by ListScanTasksByTarget. If the task does not exist or belongs to another tenant, a 400 error is returned to avoid exposing whether the resource exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// task-abc123def4567
	ScannerTaskId *string `json:"ScannerTaskId,omitempty" xml:"ScannerTaskId,omitempty"`
}

func (s GetScannerTaskResultHtmlUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScannerTaskResultHtmlUrlRequest) GoString() string {
	return s.String()
}

func (s *GetScannerTaskResultHtmlUrlRequest) GetScannerTaskId() *string {
	return s.ScannerTaskId
}

func (s *GetScannerTaskResultHtmlUrlRequest) SetScannerTaskId(v string) *GetScannerTaskResultHtmlUrlRequest {
	s.ScannerTaskId = &v
	return s
}

func (s *GetScannerTaskResultHtmlUrlRequest) Validate() error {
	return dara.Validate(s)
}
