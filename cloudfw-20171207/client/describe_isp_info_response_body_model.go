// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeIspInfoResponseBodyDataList) *DescribeIspInfoResponseBody
	GetDataList() []*DescribeIspInfoResponseBodyDataList
	SetRequestId(v string) *DescribeIspInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeIspInfoResponseBody
	GetTotalCount() *int32
}

type DescribeIspInfoResponseBody struct {
	DataList []*DescribeIspInfoResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIspInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIspInfoResponseBody) GetDataList() []*DescribeIspInfoResponseBodyDataList {
	return s.DataList
}

func (s *DescribeIspInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIspInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIspInfoResponseBody) SetDataList(v []*DescribeIspInfoResponseBodyDataList) *DescribeIspInfoResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeIspInfoResponseBody) SetRequestId(v string) *DescribeIspInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIspInfoResponseBody) SetTotalCount(v int32) *DescribeIspInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIspInfoResponseBody) Validate() error {
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

type DescribeIspInfoResponseBodyDataList struct {
	// example:
	//
	// isp-dhyw2lxfpc****
	IspId   *int32  `json:"IspId,omitempty" xml:"IspId,omitempty"`
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s DescribeIspInfoResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspInfoResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeIspInfoResponseBodyDataList) GetIspId() *int32 {
	return s.IspId
}

func (s *DescribeIspInfoResponseBodyDataList) GetIspName() *string {
	return s.IspName
}

func (s *DescribeIspInfoResponseBodyDataList) SetIspId(v int32) *DescribeIspInfoResponseBodyDataList {
	s.IspId = &v
	return s
}

func (s *DescribeIspInfoResponseBodyDataList) SetIspName(v string) *DescribeIspInfoResponseBodyDataList {
	s.IspName = &v
	return s
}

func (s *DescribeIspInfoResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
