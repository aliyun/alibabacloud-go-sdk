// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDashboardListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DashboardListResponseBody
	GetCode() *int32
	SetData(v string) *DashboardListResponseBody
	GetData() *string
	SetMessage(v string) *DashboardListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DashboardListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DashboardListResponseBody
	GetSuccess() *bool
}

type DashboardListResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DashboardListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DashboardListResponseBody) GoString() string {
	return s.String()
}

func (s *DashboardListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DashboardListResponseBody) GetData() *string {
	return s.Data
}

func (s *DashboardListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DashboardListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DashboardListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DashboardListResponseBody) SetCode(v int32) *DashboardListResponseBody {
	s.Code = &v
	return s
}

func (s *DashboardListResponseBody) SetData(v string) *DashboardListResponseBody {
	s.Data = &v
	return s
}

func (s *DashboardListResponseBody) SetMessage(v string) *DashboardListResponseBody {
	s.Message = &v
	return s
}

func (s *DashboardListResponseBody) SetRequestId(v string) *DashboardListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DashboardListResponseBody) SetSuccess(v bool) *DashboardListResponseBody {
	s.Success = &v
	return s
}

func (s *DashboardListResponseBody) Validate() error {
	return dara.Validate(s)
}
