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
	SetSuccess(v bool) *StopDIJobResponseBody
	GetSuccess() *bool
}

type StopDIJobResponseBody struct {
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 92F778C7-8F00-53B1-AE1A-B3B17101247D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *StopDIJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopDIJobResponseBody) SetRequestId(v string) *StopDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDIJobResponseBody) SetSuccess(v bool) *StopDIJobResponseBody {
	s.Success = &v
	return s
}

func (s *StopDIJobResponseBody) Validate() error {
	return dara.Validate(s)
}
