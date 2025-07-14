// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelReportShareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelReportShareResponseBody
	GetRequestId() *string
	SetResult(v bool) *CancelReportShareResponseBody
	GetResult() *bool
	SetSuccess(v bool) *CancelReportShareResponseBody
	GetSuccess() *bool
}

type CancelReportShareResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelReportShareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelReportShareResponseBody) GoString() string {
	return s.String()
}

func (s *CancelReportShareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelReportShareResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CancelReportShareResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelReportShareResponseBody) SetRequestId(v string) *CancelReportShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelReportShareResponseBody) SetResult(v bool) *CancelReportShareResponseBody {
	s.Result = &v
	return s
}

func (s *CancelReportShareResponseBody) SetSuccess(v bool) *CancelReportShareResponseBody {
	s.Success = &v
	return s
}

func (s *CancelReportShareResponseBody) Validate() error {
	return dara.Validate(s)
}
