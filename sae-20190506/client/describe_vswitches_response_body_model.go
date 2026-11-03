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
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*DescribeVSwitchesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code. Valid values:
	//
	// - Successful request: The **ErrorCode*	- field is not returned.
	//
	// - Failed request: The **ErrorCode*	- field is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - **success**: The request is successful.
	//
	// - Specific error code: The request is abnormal.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2D7E6B28-27C9-5EEA-B0A3-FC0D2F5B2263
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the creation or update is successful. Valid values:
	//
	// - Successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID.
	//
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
	// The number of available IP addresses.
	//
	// example:
	//
	// 246
	AvailableIpAddressCount *string `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// The batch status. Valid values:
	//
	// - Available: Available.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-2vc18rvxe1hk5dcsyy650
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// test
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2ze7b7cfxmxypkodyg5io
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
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
