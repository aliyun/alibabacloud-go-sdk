// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebReportTopIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeWebReportTopIpResponseBodyDataList) *DescribeWebReportTopIpResponseBody
	GetDataList() []*DescribeWebReportTopIpResponseBodyDataList
	SetRequestId(v string) *DescribeWebReportTopIpResponseBody
	GetRequestId() *string
}

type DescribeWebReportTopIpResponseBody struct {
	DataList  []*DescribeWebReportTopIpResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebReportTopIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebReportTopIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebReportTopIpResponseBody) GetDataList() []*DescribeWebReportTopIpResponseBodyDataList {
	return s.DataList
}

func (s *DescribeWebReportTopIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebReportTopIpResponseBody) SetDataList(v []*DescribeWebReportTopIpResponseBodyDataList) *DescribeWebReportTopIpResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeWebReportTopIpResponseBody) SetRequestId(v string) *DescribeWebReportTopIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebReportTopIpResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebReportTopIpResponseBodyDataList struct {
	AreaId   *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	Count    *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Isp      *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeWebReportTopIpResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebReportTopIpResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeWebReportTopIpResponseBodyDataList) GetAreaId() *string {
	return s.AreaId
}

func (s *DescribeWebReportTopIpResponseBodyDataList) GetCount() *int64 {
	return s.Count
}

func (s *DescribeWebReportTopIpResponseBodyDataList) GetIsp() *string {
	return s.Isp
}

func (s *DescribeWebReportTopIpResponseBodyDataList) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWebReportTopIpResponseBodyDataList) SetAreaId(v string) *DescribeWebReportTopIpResponseBodyDataList {
	s.AreaId = &v
	return s
}

func (s *DescribeWebReportTopIpResponseBodyDataList) SetCount(v int64) *DescribeWebReportTopIpResponseBodyDataList {
	s.Count = &v
	return s
}

func (s *DescribeWebReportTopIpResponseBodyDataList) SetIsp(v string) *DescribeWebReportTopIpResponseBodyDataList {
	s.Isp = &v
	return s
}

func (s *DescribeWebReportTopIpResponseBodyDataList) SetSourceIp(v string) *DescribeWebReportTopIpResponseBodyDataList {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebReportTopIpResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
