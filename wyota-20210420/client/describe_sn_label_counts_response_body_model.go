// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnLabelCountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSnLabelCountsResponseBody
	GetCode() *string
	SetData(v *DescribeSnLabelCountsResponseBodyData) *DescribeSnLabelCountsResponseBody
	GetData() *DescribeSnLabelCountsResponseBodyData
	SetMessage(v string) *DescribeSnLabelCountsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSnLabelCountsResponseBody
	GetRequestId() *string
}

type DescribeSnLabelCountsResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeSnLabelCountsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSnLabelCountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnLabelCountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnLabelCountsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSnLabelCountsResponseBody) GetData() *DescribeSnLabelCountsResponseBodyData {
	return s.Data
}

func (s *DescribeSnLabelCountsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSnLabelCountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnLabelCountsResponseBody) SetCode(v string) *DescribeSnLabelCountsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSnLabelCountsResponseBody) SetData(v *DescribeSnLabelCountsResponseBodyData) *DescribeSnLabelCountsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSnLabelCountsResponseBody) SetMessage(v string) *DescribeSnLabelCountsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSnLabelCountsResponseBody) SetRequestId(v string) *DescribeSnLabelCountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnLabelCountsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSnLabelCountsResponseBodyData struct {
	LabelCountDTOList []*DescribeSnLabelCountsResponseBodyDataLabelCountDTOList `json:"LabelCountDTOList,omitempty" xml:"LabelCountDTOList,omitempty" type:"Repeated"`
	TotalCount        *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSnLabelCountsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnLabelCountsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSnLabelCountsResponseBodyData) GetLabelCountDTOList() []*DescribeSnLabelCountsResponseBodyDataLabelCountDTOList {
	return s.LabelCountDTOList
}

func (s *DescribeSnLabelCountsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSnLabelCountsResponseBodyData) SetLabelCountDTOList(v []*DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) *DescribeSnLabelCountsResponseBodyData {
	s.LabelCountDTOList = v
	return s
}

func (s *DescribeSnLabelCountsResponseBodyData) SetTotalCount(v int32) *DescribeSnLabelCountsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeSnLabelCountsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSnLabelCountsResponseBodyDataLabelCountDTOList struct {
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) GoString() string {
	return s.String()
}

func (s *DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) GetCount() *string {
	return s.Count
}

func (s *DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) GetLabel() *string {
	return s.Label
}

func (s *DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) SetCount(v string) *DescribeSnLabelCountsResponseBodyDataLabelCountDTOList {
	s.Count = &v
	return s
}

func (s *DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) SetLabel(v string) *DescribeSnLabelCountsResponseBodyDataLabelCountDTOList {
	s.Label = &v
	return s
}

func (s *DescribeSnLabelCountsResponseBodyDataLabelCountDTOList) Validate() error {
	return dara.Validate(s)
}
