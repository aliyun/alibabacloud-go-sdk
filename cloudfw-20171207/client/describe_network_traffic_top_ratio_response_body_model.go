// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkTrafficTopRatioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCount(v int32) *DescribeNetworkTrafficTopRatioResponseBody
	GetDataCount() *int32
	SetDataList(v []*DescribeNetworkTrafficTopRatioResponseBodyDataList) *DescribeNetworkTrafficTopRatioResponseBody
	GetDataList() []*DescribeNetworkTrafficTopRatioResponseBodyDataList
	SetDataType(v string) *DescribeNetworkTrafficTopRatioResponseBody
	GetDataType() *string
	SetRequestId(v string) *DescribeNetworkTrafficTopRatioResponseBody
	GetRequestId() *string
}

type DescribeNetworkTrafficTopRatioResponseBody struct {
	// example:
	//
	// 1
	DataCount *int32                                                `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	DataList  []*DescribeNetworkTrafficTopRatioResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// in_src_ip
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// C05D58A1-28A9-563A-BB59-5F7D1867****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNetworkTrafficTopRatioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkTrafficTopRatioResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkTrafficTopRatioResponseBody) GetDataCount() *int32 {
	return s.DataCount
}

func (s *DescribeNetworkTrafficTopRatioResponseBody) GetDataList() []*DescribeNetworkTrafficTopRatioResponseBodyDataList {
	return s.DataList
}

func (s *DescribeNetworkTrafficTopRatioResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *DescribeNetworkTrafficTopRatioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkTrafficTopRatioResponseBody) SetDataCount(v int32) *DescribeNetworkTrafficTopRatioResponseBody {
	s.DataCount = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioResponseBody) SetDataList(v []*DescribeNetworkTrafficTopRatioResponseBodyDataList) *DescribeNetworkTrafficTopRatioResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeNetworkTrafficTopRatioResponseBody) SetDataType(v string) *DescribeNetworkTrafficTopRatioResponseBody {
	s.DataType = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioResponseBody) SetRequestId(v string) *DescribeNetworkTrafficTopRatioResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioResponseBody) Validate() error {
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

type DescribeNetworkTrafficTopRatioResponseBodyDataList struct {
	// example:
	//
	// test
	DataName *string `json:"DataName,omitempty" xml:"DataName,omitempty"`
	// example:
	//
	// 12
	DataValue *string `json:"DataValue,omitempty" xml:"DataValue,omitempty"`
}

func (s DescribeNetworkTrafficTopRatioResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkTrafficTopRatioResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeNetworkTrafficTopRatioResponseBodyDataList) GetDataName() *string {
	return s.DataName
}

func (s *DescribeNetworkTrafficTopRatioResponseBodyDataList) GetDataValue() *string {
	return s.DataValue
}

func (s *DescribeNetworkTrafficTopRatioResponseBodyDataList) SetDataName(v string) *DescribeNetworkTrafficTopRatioResponseBodyDataList {
	s.DataName = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioResponseBodyDataList) SetDataValue(v string) *DescribeNetworkTrafficTopRatioResponseBodyDataList {
	s.DataValue = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
