// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartVirusScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartVirusScanTaskResponseBody
	GetRequestId() *string
	SetScanTaskId(v int64) *StartVirusScanTaskResponseBody
	GetScanTaskId() *int64
}

type StartVirusScanTaskResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// DAE17926-4ABE-4DBD-9600-DDCB9B200F35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the virus scan task.
	//
	// example:
	//
	// 282832
	ScanTaskId *int64 `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
}

func (s StartVirusScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartVirusScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartVirusScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartVirusScanTaskResponseBody) GetScanTaskId() *int64 {
	return s.ScanTaskId
}

func (s *StartVirusScanTaskResponseBody) SetRequestId(v string) *StartVirusScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartVirusScanTaskResponseBody) SetScanTaskId(v int64) *StartVirusScanTaskResponseBody {
	s.ScanTaskId = &v
	return s
}

func (s *StartVirusScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
