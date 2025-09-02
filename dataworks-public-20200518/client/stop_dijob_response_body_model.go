// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDIJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopDIJobResponseBody
	GetRequestId() *string
}

type StopDIJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 92F778C7-8F00-53B1-AE1A-B3B17101247D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDIJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopDIJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDIJobResponseBody) SetRequestId(v string) *StopDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDIJobResponseBody) Validate() error {
	return dara.Validate(s)
}
