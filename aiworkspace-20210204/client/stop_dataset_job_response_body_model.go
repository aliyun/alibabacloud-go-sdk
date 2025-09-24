// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDatasetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopDatasetJobResponseBody
	GetRequestId() *string
}

type StopDatasetJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F620FFD3-FFDC-5873-A70C-6971CC45F467
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDatasetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDatasetJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopDatasetJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDatasetJobResponseBody) SetRequestId(v string) *StopDatasetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDatasetJobResponseBody) Validate() error {
	return dara.Validate(s)
}
