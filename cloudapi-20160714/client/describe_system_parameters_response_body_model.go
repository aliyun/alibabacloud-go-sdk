// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSystemParametersResponseBody
	GetRequestId() *string
	SetSystemParams(v *DescribeSystemParametersResponseBodySystemParams) *DescribeSystemParametersResponseBody
	GetSystemParams() *DescribeSystemParametersResponseBodySystemParams
}

type DescribeSystemParametersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0CCDF65E-6050-412D-AD68-FA3D9196836C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned information about system parameters. It is an array that consists of SystemParam data.
	SystemParams *DescribeSystemParametersResponseBodySystemParams `json:"SystemParams,omitempty" xml:"SystemParams,omitempty" type:"Struct"`
}

func (s DescribeSystemParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSystemParametersResponseBody) GetSystemParams() *DescribeSystemParametersResponseBodySystemParams {
	return s.SystemParams
}

func (s *DescribeSystemParametersResponseBody) SetRequestId(v string) *DescribeSystemParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemParametersResponseBody) SetSystemParams(v *DescribeSystemParametersResponseBodySystemParams) *DescribeSystemParametersResponseBody {
	s.SystemParams = v
	return s
}

func (s *DescribeSystemParametersResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSystemParametersResponseBodySystemParams struct {
	SystemParamItem []*DescribeSystemParametersResponseBodySystemParamsSystemParamItem `json:"SystemParamItem,omitempty" xml:"SystemParamItem,omitempty" type:"Repeated"`
}

func (s DescribeSystemParametersResponseBodySystemParams) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemParametersResponseBodySystemParams) GoString() string {
	return s.String()
}

func (s *DescribeSystemParametersResponseBodySystemParams) GetSystemParamItem() []*DescribeSystemParametersResponseBodySystemParamsSystemParamItem {
	return s.SystemParamItem
}

func (s *DescribeSystemParametersResponseBodySystemParams) SetSystemParamItem(v []*DescribeSystemParametersResponseBodySystemParamsSystemParamItem) *DescribeSystemParametersResponseBodySystemParams {
	s.SystemParamItem = v
	return s
}

func (s *DescribeSystemParametersResponseBodySystemParams) Validate() error {
	return dara.Validate(s)
}

type DescribeSystemParametersResponseBodySystemParamsSystemParamItem struct {
	// Examples
	//
	// example:
	//
	// 192.168.1.1
	DemoValue *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	// The description of a parameter.
	//
	// example:
	//
	// Client IP Address
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// CaClientIp
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// The type of the parameter.
	//
	// example:
	//
	// string
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s DescribeSystemParametersResponseBodySystemParamsSystemParamItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemParametersResponseBodySystemParamsSystemParamItem) GoString() string {
	return s.String()
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) GetDemoValue() *string {
	return s.DemoValue
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) GetDescription() *string {
	return s.Description
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) GetParamName() *string {
	return s.ParamName
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) GetParamType() *string {
	return s.ParamType
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) SetDemoValue(v string) *DescribeSystemParametersResponseBodySystemParamsSystemParamItem {
	s.DemoValue = &v
	return s
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) SetDescription(v string) *DescribeSystemParametersResponseBodySystemParamsSystemParamItem {
	s.Description = &v
	return s
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) SetParamName(v string) *DescribeSystemParametersResponseBodySystemParamsSystemParamItem {
	s.ParamName = &v
	return s
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) SetParamType(v string) *DescribeSystemParametersResponseBodySystemParamsSystemParamItem {
	s.ParamType = &v
	return s
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) Validate() error {
	return dara.Validate(s)
}
