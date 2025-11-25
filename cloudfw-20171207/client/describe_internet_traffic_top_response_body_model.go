// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetTrafficTopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCount(v int32) *DescribeInternetTrafficTopResponseBody
	GetDataCount() *int32
	SetDataList(v []*DescribeInternetTrafficTopResponseBodyDataList) *DescribeInternetTrafficTopResponseBody
	GetDataList() []*DescribeInternetTrafficTopResponseBodyDataList
	SetDataType(v string) *DescribeInternetTrafficTopResponseBody
	GetDataType() *string
	SetRequestId(v string) *DescribeInternetTrafficTopResponseBody
	GetRequestId() *string
}

type DescribeInternetTrafficTopResponseBody struct {
	// example:
	//
	// 10
	DataCount *int32                                            `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	DataList  []*DescribeInternetTrafficTopResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// in_src_ip
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// 15FCCC52-1E23-57AE-B5EF-3E00A3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInternetTrafficTopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetTrafficTopResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetTrafficTopResponseBody) GetDataCount() *int32 {
	return s.DataCount
}

func (s *DescribeInternetTrafficTopResponseBody) GetDataList() []*DescribeInternetTrafficTopResponseBodyDataList {
	return s.DataList
}

func (s *DescribeInternetTrafficTopResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *DescribeInternetTrafficTopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInternetTrafficTopResponseBody) SetDataCount(v int32) *DescribeInternetTrafficTopResponseBody {
	s.DataCount = &v
	return s
}

func (s *DescribeInternetTrafficTopResponseBody) SetDataList(v []*DescribeInternetTrafficTopResponseBodyDataList) *DescribeInternetTrafficTopResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeInternetTrafficTopResponseBody) SetDataType(v string) *DescribeInternetTrafficTopResponseBody {
	s.DataType = &v
	return s
}

func (s *DescribeInternetTrafficTopResponseBody) SetRequestId(v string) *DescribeInternetTrafficTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetTrafficTopResponseBody) Validate() error {
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

type DescribeInternetTrafficTopResponseBodyDataList struct {
	// example:
	//
	// US
	DataName *string `json:"DataName,omitempty" xml:"DataName,omitempty"`
	// example:
	//
	// 47.12
	DataValue *string `json:"DataValue,omitempty" xml:"DataValue,omitempty"`
	// example:
	//
	// false
	IsSubscribed *bool     `json:"IsSubscribed,omitempty" xml:"IsSubscribed,omitempty"`
	LabelList    []*string `json:"LabelList,omitempty" xml:"LabelList,omitempty" type:"Repeated"`
	// example:
	//
	// 27
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// example:
	//
	// 0
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
}

func (s DescribeInternetTrafficTopResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetTrafficTopResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) GetDataName() *string {
	return s.DataName
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) GetDataValue() *string {
	return s.DataValue
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) GetIsSubscribed() *bool {
	return s.IsSubscribed
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) GetLabelList() []*string {
	return s.LabelList
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) SetDataName(v string) *DescribeInternetTrafficTopResponseBodyDataList {
	s.DataName = &v
	return s
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) SetDataValue(v string) *DescribeInternetTrafficTopResponseBodyDataList {
	s.DataValue = &v
	return s
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) SetIsSubscribed(v bool) *DescribeInternetTrafficTopResponseBodyDataList {
	s.IsSubscribed = &v
	return s
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) SetLabelList(v []*string) *DescribeInternetTrafficTopResponseBodyDataList {
	s.LabelList = v
	return s
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) SetSessionCount(v int64) *DescribeInternetTrafficTopResponseBodyDataList {
	s.SessionCount = &v
	return s
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) SetTotalBytes(v int64) *DescribeInternetTrafficTopResponseBodyDataList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeInternetTrafficTopResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
