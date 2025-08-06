// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserVodStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *DescribeUserVodStatusResponseBody
	GetEnabled() *bool
	SetInDebt(v bool) *DescribeUserVodStatusResponseBody
	GetInDebt() *bool
	SetInDebtOverdue(v bool) *DescribeUserVodStatusResponseBody
	GetInDebtOverdue() *bool
	SetOnService(v bool) *DescribeUserVodStatusResponseBody
	GetOnService() *bool
	SetRequestId(v string) *DescribeUserVodStatusResponseBody
	GetRequestId() *string
}

type DescribeUserVodStatusResponseBody struct {
	Enabled       *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	InDebt        *bool   `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	InDebtOverdue *bool   `json:"InDebtOverdue,omitempty" xml:"InDebtOverdue,omitempty"`
	OnService     *bool   `json:"OnService,omitempty" xml:"OnService,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserVodStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserVodStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserVodStatusResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeUserVodStatusResponseBody) GetInDebt() *bool {
	return s.InDebt
}

func (s *DescribeUserVodStatusResponseBody) GetInDebtOverdue() *bool {
	return s.InDebtOverdue
}

func (s *DescribeUserVodStatusResponseBody) GetOnService() *bool {
	return s.OnService
}

func (s *DescribeUserVodStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserVodStatusResponseBody) SetEnabled(v bool) *DescribeUserVodStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeUserVodStatusResponseBody) SetInDebt(v bool) *DescribeUserVodStatusResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeUserVodStatusResponseBody) SetInDebtOverdue(v bool) *DescribeUserVodStatusResponseBody {
	s.InDebtOverdue = &v
	return s
}

func (s *DescribeUserVodStatusResponseBody) SetOnService(v bool) *DescribeUserVodStatusResponseBody {
	s.OnService = &v
	return s
}

func (s *DescribeUserVodStatusResponseBody) SetRequestId(v string) *DescribeUserVodStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserVodStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
