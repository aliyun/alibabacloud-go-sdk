// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopScannerTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopScannerTaskResponseBody
	GetRequestId() *string
}

type StopScannerTaskResponseBody struct {
	// The unique identifier of the request, used for troubleshooting and log tracing.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4C95-9C63-B7AB7B5A9C8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopScannerTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopScannerTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopScannerTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopScannerTaskResponseBody) SetRequestId(v string) *StopScannerTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopScannerTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
