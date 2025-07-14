// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddShareReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddShareReportResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddShareReportResponseBody
	GetResult() *bool
	SetSuccess(v bool) *AddShareReportResponseBody
	GetSuccess() *bool
}

type AddShareReportResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 05739b8e-3de0-4204-9669-7f04f02522b9
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

func (s AddShareReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddShareReportResponseBody) GoString() string {
	return s.String()
}

func (s *AddShareReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddShareReportResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddShareReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddShareReportResponseBody) SetRequestId(v string) *AddShareReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddShareReportResponseBody) SetResult(v bool) *AddShareReportResponseBody {
	s.Result = &v
	return s
}

func (s *AddShareReportResponseBody) SetSuccess(v bool) *AddShareReportResponseBody {
	s.Success = &v
	return s
}

func (s *AddShareReportResponseBody) Validate() error {
	return dara.Validate(s)
}
