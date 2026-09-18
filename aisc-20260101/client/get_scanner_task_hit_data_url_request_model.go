// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScannerTaskHitDataUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScannerTaskId(v string) *GetScannerTaskHitDataUrlRequest
	GetScannerTaskId() *string
}

type GetScannerTaskHitDataUrlRequest struct {
	// The unique identifier of the scan task. This is the TaskId returned by CreateTargetScanTask or the ScannerTaskId returned by ListScanTasksByTarget. This parameter is registered as optional but is required in practice. An empty value returns HTTP status code 400. If the task does not exist or belongs to another tenant, HTTP status code 400 is returned uniformly to avoid exposing whether the resource exists.
	//
	// example:
	//
	// task-abc123def4567
	ScannerTaskId *string `json:"ScannerTaskId,omitempty" xml:"ScannerTaskId,omitempty"`
}

func (s GetScannerTaskHitDataUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScannerTaskHitDataUrlRequest) GoString() string {
	return s.String()
}

func (s *GetScannerTaskHitDataUrlRequest) GetScannerTaskId() *string {
	return s.ScannerTaskId
}

func (s *GetScannerTaskHitDataUrlRequest) SetScannerTaskId(v string) *GetScannerTaskHitDataUrlRequest {
	s.ScannerTaskId = &v
	return s
}

func (s *GetScannerTaskHitDataUrlRequest) Validate() error {
	return dara.Validate(s)
}
