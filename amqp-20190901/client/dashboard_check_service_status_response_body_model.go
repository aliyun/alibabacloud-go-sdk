// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDashboardCheckServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DashboardCheckServiceStatusResponseBody
	GetCode() *int32
	SetData(v string) *DashboardCheckServiceStatusResponseBody
	GetData() *string
	SetMessage(v string) *DashboardCheckServiceStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *DashboardCheckServiceStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DashboardCheckServiceStatusResponseBody
	GetSuccess() *bool
}

type DashboardCheckServiceStatusResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DashboardCheckServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DashboardCheckServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DashboardCheckServiceStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DashboardCheckServiceStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *DashboardCheckServiceStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DashboardCheckServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DashboardCheckServiceStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DashboardCheckServiceStatusResponseBody) SetCode(v int32) *DashboardCheckServiceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DashboardCheckServiceStatusResponseBody) SetData(v string) *DashboardCheckServiceStatusResponseBody {
	s.Data = &v
	return s
}

func (s *DashboardCheckServiceStatusResponseBody) SetMessage(v string) *DashboardCheckServiceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DashboardCheckServiceStatusResponseBody) SetRequestId(v string) *DashboardCheckServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DashboardCheckServiceStatusResponseBody) SetSuccess(v bool) *DashboardCheckServiceStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DashboardCheckServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
