// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRCInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopRCInstancesResponseBody
	GetRequestId() *string
}

type StopRCInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 481BC3B1-7069-5D37-9B6C-21757F8F9FB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRCInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRCInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StopRCInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRCInstancesResponseBody) SetRequestId(v string) *StopRCInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRCInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
