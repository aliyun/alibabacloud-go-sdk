// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenIpAccessSrcStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeOpenIpAccessSrcStatResponseBody
	GetRequestId() *string
	SetStatList(v []*DescribeOpenIpAccessSrcStatResponseBodyStatList) *DescribeOpenIpAccessSrcStatResponseBody
	GetStatList() []*DescribeOpenIpAccessSrcStatResponseBodyStatList
	SetTotalCount(v int32) *DescribeOpenIpAccessSrcStatResponseBody
	GetTotalCount() *int32
}

type DescribeOpenIpAccessSrcStatResponseBody struct {
	// example:
	//
	// 0DC783F1-B3A7-578D-8A63-*****
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatList  []*DescribeOpenIpAccessSrcStatResponseBodyStatList `json:"StatList,omitempty" xml:"StatList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOpenIpAccessSrcStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenIpAccessSrcStatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenIpAccessSrcStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpenIpAccessSrcStatResponseBody) GetStatList() []*DescribeOpenIpAccessSrcStatResponseBodyStatList {
	return s.StatList
}

func (s *DescribeOpenIpAccessSrcStatResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOpenIpAccessSrcStatResponseBody) SetRequestId(v string) *DescribeOpenIpAccessSrcStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenIpAccessSrcStatResponseBody) SetStatList(v []*DescribeOpenIpAccessSrcStatResponseBodyStatList) *DescribeOpenIpAccessSrcStatResponseBody {
	s.StatList = v
	return s
}

func (s *DescribeOpenIpAccessSrcStatResponseBody) SetTotalCount(v int32) *DescribeOpenIpAccessSrcStatResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOpenIpAccessSrcStatResponseBody) Validate() error {
	if s.StatList != nil {
		for _, item := range s.StatList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOpenIpAccessSrcStatResponseBodyStatList struct {
	// example:
	//
	// 12
	AbnormalSrcIp *int64 `json:"AbnormalSrcIp,omitempty" xml:"AbnormalSrcIp,omitempty"`
	// example:
	//
	// HTTP
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// example:
	//
	// 21
	NormalSrcIp *int64 `json:"NormalSrcIp,omitempty" xml:"NormalSrcIp,omitempty"`
	// example:
	//
	// 6163
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeOpenIpAccessSrcStatResponseBodyStatList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenIpAccessSrcStatResponseBodyStatList) GoString() string {
	return s.String()
}

func (s *DescribeOpenIpAccessSrcStatResponseBodyStatList) GetAbnormalSrcIp() *int64 {
	return s.AbnormalSrcIp
}

func (s *DescribeOpenIpAccessSrcStatResponseBodyStatList) GetApp() *string {
	return s.App
}

func (s *DescribeOpenIpAccessSrcStatResponseBodyStatList) GetNormalSrcIp() *int64 {
	return s.NormalSrcIp
}

func (s *DescribeOpenIpAccessSrcStatResponseBodyStatList) GetPort() *string {
	return s.Port
}

func (s *DescribeOpenIpAccessSrcStatResponseBodyStatList) SetAbnormalSrcIp(v int64) *DescribeOpenIpAccessSrcStatResponseBodyStatList {
	s.AbnormalSrcIp = &v
	return s
}

func (s *DescribeOpenIpAccessSrcStatResponseBodyStatList) SetApp(v string) *DescribeOpenIpAccessSrcStatResponseBodyStatList {
	s.App = &v
	return s
}

func (s *DescribeOpenIpAccessSrcStatResponseBodyStatList) SetNormalSrcIp(v int64) *DescribeOpenIpAccessSrcStatResponseBodyStatList {
	s.NormalSrcIp = &v
	return s
}

func (s *DescribeOpenIpAccessSrcStatResponseBodyStatList) SetPort(v string) *DescribeOpenIpAccessSrcStatResponseBodyStatList {
	s.Port = &v
	return s
}

func (s *DescribeOpenIpAccessSrcStatResponseBodyStatList) Validate() error {
	return dara.Validate(s)
}
