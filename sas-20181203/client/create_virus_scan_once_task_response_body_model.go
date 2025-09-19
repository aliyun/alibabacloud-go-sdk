// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirusScanOnceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVirusScanOnceTaskResponseBody
	GetRequestId() *string
}

type CreateVirusScanOnceTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F3B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVirusScanOnceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirusScanOnceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirusScanOnceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirusScanOnceTaskResponseBody) SetRequestId(v string) *CreateVirusScanOnceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirusScanOnceTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
