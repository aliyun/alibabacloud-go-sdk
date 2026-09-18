// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScannerTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteScannerTaskResponseBody
	GetRequestId() *string
}

type DeleteScannerTaskResponseBody struct {
	// The unique identifier of the request, used for troubleshooting and log tracing.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4C95-9C63-B7AB7B5A9C8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScannerTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScannerTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScannerTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScannerTaskResponseBody) SetRequestId(v string) *DeleteScannerTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScannerTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
