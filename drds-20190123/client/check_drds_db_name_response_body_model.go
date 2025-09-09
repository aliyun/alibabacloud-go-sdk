// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDrdsDbNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckDrdsDbNameResponseBody
	GetRequestId() *string
	SetResult(v bool) *CheckDrdsDbNameResponseBody
	GetResult() *bool
	SetSuccess(v bool) *CheckDrdsDbNameResponseBody
	GetSuccess() *bool
}

type CheckDrdsDbNameResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CF38538C-68BD-4278-B58F-EDE96F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the DRDS database name is valid. Valid values: true: The database name is valid. false: the database name is invalid.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckDrdsDbNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDrdsDbNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDrdsDbNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDrdsDbNameResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CheckDrdsDbNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckDrdsDbNameResponseBody) SetRequestId(v string) *CheckDrdsDbNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDrdsDbNameResponseBody) SetResult(v bool) *CheckDrdsDbNameResponseBody {
	s.Result = &v
	return s
}

func (s *CheckDrdsDbNameResponseBody) SetSuccess(v bool) *CheckDrdsDbNameResponseBody {
	s.Success = &v
	return s
}

func (s *CheckDrdsDbNameResponseBody) Validate() error {
	return dara.Validate(s)
}
