// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdcClassListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeCdcClassListResponseBodyData) *DescribeCdcClassListResponseBody
	GetData() *DescribeCdcClassListResponseBodyData
	SetRequestId(v string) *DescribeCdcClassListResponseBody
	GetRequestId() *string
}

type DescribeCdcClassListResponseBody struct {
	Data *DescribeCdcClassListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdcClassListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcClassListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdcClassListResponseBody) GetData() *DescribeCdcClassListResponseBodyData {
	return s.Data
}

func (s *DescribeCdcClassListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdcClassListResponseBody) SetData(v *DescribeCdcClassListResponseBodyData) *DescribeCdcClassListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCdcClassListResponseBody) SetRequestId(v string) *DescribeCdcClassListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdcClassListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdcClassListResponseBodyData struct {
	// array
	ClassCodeList []*DescribeCdcClassListResponseBodyDataClassCodeList `json:"ClassCodeList,omitempty" xml:"ClassCodeList,omitempty" type:"Repeated"`
}

func (s DescribeCdcClassListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcClassListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCdcClassListResponseBodyData) GetClassCodeList() []*DescribeCdcClassListResponseBodyDataClassCodeList {
	return s.ClassCodeList
}

func (s *DescribeCdcClassListResponseBodyData) SetClassCodeList(v []*DescribeCdcClassListResponseBodyDataClassCodeList) *DescribeCdcClassListResponseBodyData {
	s.ClassCodeList = v
	return s
}

func (s *DescribeCdcClassListResponseBodyData) Validate() error {
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

type DescribeCdcClassListResponseBodyDataClassCodeList struct {
	// example:
	//
	// polarx.n2.small.2e.cdc
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

func (s DescribeCdcClassListResponseBodyDataClassCodeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdcClassListResponseBodyDataClassCodeList) GoString() string {
	return s.String()
}

func (s *DescribeCdcClassListResponseBodyDataClassCodeList) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeCdcClassListResponseBodyDataClassCodeList) GetCpuCore() *string {
	return s.CpuCore
}

func (s *DescribeCdcClassListResponseBodyDataClassCodeList) GetMem() *string {
	return s.Mem
}

func (s *DescribeCdcClassListResponseBodyDataClassCodeList) SetClassCode(v string) *DescribeCdcClassListResponseBodyDataClassCodeList {
	s.ClassCode = &v
	return s
}

func (s *DescribeCdcClassListResponseBodyDataClassCodeList) SetCpuCore(v string) *DescribeCdcClassListResponseBodyDataClassCodeList {
	s.CpuCore = &v
	return s
}

func (s *DescribeCdcClassListResponseBodyDataClassCodeList) SetMem(v string) *DescribeCdcClassListResponseBodyDataClassCodeList {
	s.Mem = &v
	return s
}

func (s *DescribeCdcClassListResponseBodyDataClassCodeList) Validate() error {
	return dara.Validate(s)
}
