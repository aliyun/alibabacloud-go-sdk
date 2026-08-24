// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelVirusScanTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelVirusScanTasksResponseBody
	GetRequestId() *string
}

type CancelVirusScanTasksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelVirusScanTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelVirusScanTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CancelVirusScanTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelVirusScanTasksResponseBody) SetRequestId(v string) *CancelVirusScanTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelVirusScanTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
