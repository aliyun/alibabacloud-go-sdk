// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnarClassListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeColumnarClassListResponseBodyData) *DescribeColumnarClassListResponseBody
	GetData() *DescribeColumnarClassListResponseBodyData
	SetRequestId(v string) *DescribeColumnarClassListResponseBody
	GetRequestId() *string
}

type DescribeColumnarClassListResponseBody struct {
	Data *DescribeColumnarClassListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeColumnarClassListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarClassListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColumnarClassListResponseBody) GetData() *DescribeColumnarClassListResponseBodyData {
	return s.Data
}

func (s *DescribeColumnarClassListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeColumnarClassListResponseBody) SetData(v *DescribeColumnarClassListResponseBodyData) *DescribeColumnarClassListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeColumnarClassListResponseBody) SetRequestId(v string) *DescribeColumnarClassListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColumnarClassListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeColumnarClassListResponseBodyData struct {
	// object
	ClassCodeList []*DescribeColumnarClassListResponseBodyDataClassCodeList `json:"ClassCodeList,omitempty" xml:"ClassCodeList,omitempty" type:"Repeated"`
}

func (s DescribeColumnarClassListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarClassListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeColumnarClassListResponseBodyData) GetClassCodeList() []*DescribeColumnarClassListResponseBodyDataClassCodeList {
	return s.ClassCodeList
}

func (s *DescribeColumnarClassListResponseBodyData) SetClassCodeList(v []*DescribeColumnarClassListResponseBodyDataClassCodeList) *DescribeColumnarClassListResponseBodyData {
	s.ClassCodeList = v
	return s
}

func (s *DescribeColumnarClassListResponseBodyData) Validate() error {
	if s.ClassCodeList != nil {
		for _, item := range s.ClassCodeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeColumnarClassListResponseBodyDataClassCodeList struct {
	// example:
	//
	// polarx.n8.large.col
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// example:
	//
	// 2
	CpuCore *string `json:"CpuCore,omitempty" xml:"CpuCore,omitempty"`
	// example:
	//
	// 2048
	Mem *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeColumnarClassListResponseBodyDataClassCodeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarClassListResponseBodyDataClassCodeList) GoString() string {
	return s.String()
}

func (s *DescribeColumnarClassListResponseBodyDataClassCodeList) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeColumnarClassListResponseBodyDataClassCodeList) GetCpuCore() *string {
	return s.CpuCore
}

func (s *DescribeColumnarClassListResponseBodyDataClassCodeList) GetMem() *string {
	return s.Mem
}

func (s *DescribeColumnarClassListResponseBodyDataClassCodeList) SetClassCode(v string) *DescribeColumnarClassListResponseBodyDataClassCodeList {
	s.ClassCode = &v
	return s
}

func (s *DescribeColumnarClassListResponseBodyDataClassCodeList) SetCpuCore(v string) *DescribeColumnarClassListResponseBodyDataClassCodeList {
	s.CpuCore = &v
	return s
}

func (s *DescribeColumnarClassListResponseBodyDataClassCodeList) SetMem(v string) *DescribeColumnarClassListResponseBodyDataClassCodeList {
	s.Mem = &v
	return s
}

func (s *DescribeColumnarClassListResponseBodyDataClassCodeList) Validate() error {
	return dara.Validate(s)
}
