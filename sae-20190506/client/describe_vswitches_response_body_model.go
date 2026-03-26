// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeVSwitchesResponseBody
	GetCode() *string
	SetData(v []*DescribeVSwitchesResponseBodyData) *DescribeVSwitchesResponseBody
	GetData() []*DescribeVSwitchesResponseBodyData
	SetErrorCode(v string) *DescribeVSwitchesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeVSwitchesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeVSwitchesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeVSwitchesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeVSwitchesResponseBody
	GetTraceId() *string
}

type DescribeVSwitchesResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*DescribeVSwitchesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2D7E6B28-27C9-5EEA-B0A3-FC0D2F5B2263
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 0a001a0817594572118465681d2341
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeVSwitchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeVSwitchesResponseBody) GetData() []*DescribeVSwitchesResponseBodyData {
	return s.Data
}

func (s *DescribeVSwitchesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeVSwitchesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeVSwitchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVSwitchesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeVSwitchesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeVSwitchesResponseBody) SetCode(v string) *DescribeVSwitchesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetData(v []*DescribeVSwitchesResponseBodyData) *DescribeVSwitchesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetErrorCode(v string) *DescribeVSwitchesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetMessage(v string) *DescribeVSwitchesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetRequestId(v string) *DescribeVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetSuccess(v bool) *DescribeVSwitchesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetTraceId(v string) *DescribeVSwitchesResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVSwitchesResponseBodyData struct {
	// example:
	//
	// 246
	AvailableIpAddressCount *string `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vsw-2vc18rvxe1hk5dcsyy650
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// test
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// vpcid。
	//
	// example:
	//
	// vpc-2ze7b7cfxmxypkodyg5io
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-beijing-l
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyData) GetAvailableIpAddressCount() *string {
	return s.AvailableIpAddressCount
}

func (s *DescribeVSwitchesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeVSwitchesResponseBodyData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchesResponseBodyData) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchesResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchesResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVSwitchesResponseBodyData) SetAvailableIpAddressCount(v string) *DescribeVSwitchesResponseBodyData {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyData) SetStatus(v string) *DescribeVSwitchesResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyData) SetVSwitchId(v string) *DescribeVSwitchesResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyData) SetVSwitchName(v string) *DescribeVSwitchesResponseBodyData {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyData) SetVpcId(v string) *DescribeVSwitchesResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyData) SetZoneId(v string) *DescribeVSwitchesResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
