// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDashboardNl2sqlStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDashboardNl2sqlStatusResponseBody
	GetRequestId() *string
	SetResult(v []*string) *ModifyDashboardNl2sqlStatusResponseBody
	GetResult() []*string
	SetSuccess(v bool) *ModifyDashboardNl2sqlStatusResponseBody
	GetSuccess() *bool
}

type ModifyDashboardNl2sqlStatusResponseBody struct {
	// example:
	//
	// 46e537a5****
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDashboardNl2sqlStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDashboardNl2sqlStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDashboardNl2sqlStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDashboardNl2sqlStatusResponseBody) GetResult() []*string {
	return s.Result
}

func (s *ModifyDashboardNl2sqlStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDashboardNl2sqlStatusResponseBody) SetRequestId(v string) *ModifyDashboardNl2sqlStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDashboardNl2sqlStatusResponseBody) SetResult(v []*string) *ModifyDashboardNl2sqlStatusResponseBody {
	s.Result = v
	return s
}

func (s *ModifyDashboardNl2sqlStatusResponseBody) SetSuccess(v bool) *ModifyDashboardNl2sqlStatusResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDashboardNl2sqlStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
