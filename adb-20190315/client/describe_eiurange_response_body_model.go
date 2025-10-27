// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEIURangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEIUInfo(v *DescribeEIURangeResponseBodyEIUInfo) *DescribeEIURangeResponseBody
	GetEIUInfo() *DescribeEIURangeResponseBodyEIUInfo
	SetRequestId(v string) *DescribeEIURangeResponseBody
	GetRequestId() *string
}

type DescribeEIURangeResponseBody struct {
	// The queried information about the number of EIUs.
	EIUInfo *DescribeEIURangeResponseBodyEIUInfo `json:"EIUInfo,omitempty" xml:"EIUInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D65A809F-34CE-4550-9BC1-0ED21ETG380
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEIURangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEIURangeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEIURangeResponseBody) GetEIUInfo() *DescribeEIURangeResponseBodyEIUInfo {
	return s.EIUInfo
}

func (s *DescribeEIURangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEIURangeResponseBody) SetEIUInfo(v *DescribeEIURangeResponseBodyEIUInfo) *DescribeEIURangeResponseBody {
	s.EIUInfo = v
	return s
}

func (s *DescribeEIURangeResponseBody) SetRequestId(v string) *DescribeEIURangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEIURangeResponseBody) Validate() error {
	if s.EIUInfo != nil {
		if err := s.EIUInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEIURangeResponseBodyEIUInfo struct {
	DefaultReservedNodeSize *string `json:"DefaultReservedNodeSize,omitempty" xml:"DefaultReservedNodeSize,omitempty"`
	// The suggested value for the number of EIUs.
	//
	// example:
	//
	// 2
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The queried range for the number of EIUs.
	EIURange []*int64 `json:"EIURange,omitempty" xml:"EIURange,omitempty" type:"Repeated"`
	// A reserved parameter.
	//
	// example:
	//
	// none
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// none
	MinValue              *string   `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ReservedNodeSizeRange []*string `json:"ReservedNodeSizeRange,omitempty" xml:"ReservedNodeSizeRange,omitempty" type:"Repeated"`
	// A reserved parameter.
	//
	// example:
	//
	// none
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
	// A reserved parameter.
	StorageResourceRange []*string `json:"StorageResourceRange,omitempty" xml:"StorageResourceRange,omitempty" type:"Repeated"`
}

func (s DescribeEIURangeResponseBodyEIUInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEIURangeResponseBodyEIUInfo) GoString() string {
	return s.String()
}

func (s *DescribeEIURangeResponseBodyEIUInfo) GetDefaultReservedNodeSize() *string {
	return s.DefaultReservedNodeSize
}

func (s *DescribeEIURangeResponseBodyEIUInfo) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeEIURangeResponseBodyEIUInfo) GetEIURange() []*int64 {
	return s.EIURange
}

func (s *DescribeEIURangeResponseBodyEIUInfo) GetMaxValue() *string {
	return s.MaxValue
}

func (s *DescribeEIURangeResponseBodyEIUInfo) GetMinValue() *string {
	return s.MinValue
}

func (s *DescribeEIURangeResponseBodyEIUInfo) GetReservedNodeSizeRange() []*string {
	return s.ReservedNodeSizeRange
}

func (s *DescribeEIURangeResponseBodyEIUInfo) GetStep() *string {
	return s.Step
}

func (s *DescribeEIURangeResponseBodyEIUInfo) GetStorageResourceRange() []*string {
	return s.StorageResourceRange
}

func (s *DescribeEIURangeResponseBodyEIUInfo) SetDefaultReservedNodeSize(v string) *DescribeEIURangeResponseBodyEIUInfo {
	s.DefaultReservedNodeSize = &v
	return s
}

func (s *DescribeEIURangeResponseBodyEIUInfo) SetDefaultValue(v string) *DescribeEIURangeResponseBodyEIUInfo {
	s.DefaultValue = &v
	return s
}

func (s *DescribeEIURangeResponseBodyEIUInfo) SetEIURange(v []*int64) *DescribeEIURangeResponseBodyEIUInfo {
	s.EIURange = v
	return s
}

func (s *DescribeEIURangeResponseBodyEIUInfo) SetMaxValue(v string) *DescribeEIURangeResponseBodyEIUInfo {
	s.MaxValue = &v
	return s
}

func (s *DescribeEIURangeResponseBodyEIUInfo) SetMinValue(v string) *DescribeEIURangeResponseBodyEIUInfo {
	s.MinValue = &v
	return s
}

func (s *DescribeEIURangeResponseBodyEIUInfo) SetReservedNodeSizeRange(v []*string) *DescribeEIURangeResponseBodyEIUInfo {
	s.ReservedNodeSizeRange = v
	return s
}

func (s *DescribeEIURangeResponseBodyEIUInfo) SetStep(v string) *DescribeEIURangeResponseBodyEIUInfo {
	s.Step = &v
	return s
}

func (s *DescribeEIURangeResponseBodyEIUInfo) SetStorageResourceRange(v []*string) *DescribeEIURangeResponseBodyEIUInfo {
	s.StorageResourceRange = v
	return s
}

func (s *DescribeEIURangeResponseBodyEIUInfo) Validate() error {
	return dara.Validate(s)
}
