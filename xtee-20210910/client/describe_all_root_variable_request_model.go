// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllRootVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAllRootVariableRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeAllRootVariableRequest
	GetSourceIp() *string
	SetDeviceVariableIds(v string) *DescribeAllRootVariableRequest
	GetDeviceVariableIds() *string
	SetEventCode(v string) *DescribeAllRootVariableRequest
	GetEventCode() *string
	SetExpressionVariableIds(v string) *DescribeAllRootVariableRequest
	GetExpressionVariableIds() *string
	SetId(v int64) *DescribeAllRootVariableRequest
	GetId() *int64
	SetNativeVariableIds(v string) *DescribeAllRootVariableRequest
	GetNativeVariableIds() *string
	SetQueryVariableIds(v string) *DescribeAllRootVariableRequest
	GetQueryVariableIds() *string
	SetRegId(v string) *DescribeAllRootVariableRequest
	GetRegId() *string
	SetVelocityVariableIds(v string) *DescribeAllRootVariableRequest
	GetVelocityVariableIds() *string
}

type DescribeAllRootVariableRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Request source IP. No need to fill in, the system will automatically obtain it.
	//
	// example:
	//
	// 61.169.104.202
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// Device variable list
	//
	// example:
	//
	// [180]
	DeviceVariableIds *string `json:"deviceVariableIds,omitempty" xml:"deviceVariableIds,omitempty"`
	// Event code
	//
	// example:
	//
	// de_arqbuy7206
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Custom variable list
	//
	// example:
	//
	// [6780]
	ExpressionVariableIds *string `json:"expressionVariableIds,omitempty" xml:"expressionVariableIds,omitempty"`
	// Variable ID.
	//
	// example:
	//
	// 2557
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Event field list
	//
	// example:
	//
	// [106780]
	NativeVariableIds *string `json:"nativeVariableIds,omitempty" xml:"nativeVariableIds,omitempty"`
	// Custom query variable
	//
	// example:
	//
	// [2678]
	QueryVariableIds *string `json:"queryVariableIds,omitempty" xml:"queryVariableIds,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Custom cumulative variable
	//
	// example:
	//
	// [780]
	VelocityVariableIds *string `json:"velocityVariableIds,omitempty" xml:"velocityVariableIds,omitempty"`
}

func (s DescribeAllRootVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllRootVariableRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllRootVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAllRootVariableRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAllRootVariableRequest) GetDeviceVariableIds() *string {
	return s.DeviceVariableIds
}

func (s *DescribeAllRootVariableRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeAllRootVariableRequest) GetExpressionVariableIds() *string {
	return s.ExpressionVariableIds
}

func (s *DescribeAllRootVariableRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeAllRootVariableRequest) GetNativeVariableIds() *string {
	return s.NativeVariableIds
}

func (s *DescribeAllRootVariableRequest) GetQueryVariableIds() *string {
	return s.QueryVariableIds
}

func (s *DescribeAllRootVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAllRootVariableRequest) GetVelocityVariableIds() *string {
	return s.VelocityVariableIds
}

func (s *DescribeAllRootVariableRequest) SetLang(v string) *DescribeAllRootVariableRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAllRootVariableRequest) SetSourceIp(v string) *DescribeAllRootVariableRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAllRootVariableRequest) SetDeviceVariableIds(v string) *DescribeAllRootVariableRequest {
	s.DeviceVariableIds = &v
	return s
}

func (s *DescribeAllRootVariableRequest) SetEventCode(v string) *DescribeAllRootVariableRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeAllRootVariableRequest) SetExpressionVariableIds(v string) *DescribeAllRootVariableRequest {
	s.ExpressionVariableIds = &v
	return s
}

func (s *DescribeAllRootVariableRequest) SetId(v int64) *DescribeAllRootVariableRequest {
	s.Id = &v
	return s
}

func (s *DescribeAllRootVariableRequest) SetNativeVariableIds(v string) *DescribeAllRootVariableRequest {
	s.NativeVariableIds = &v
	return s
}

func (s *DescribeAllRootVariableRequest) SetQueryVariableIds(v string) *DescribeAllRootVariableRequest {
	s.QueryVariableIds = &v
	return s
}

func (s *DescribeAllRootVariableRequest) SetRegId(v string) *DescribeAllRootVariableRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAllRootVariableRequest) SetVelocityVariableIds(v string) *DescribeAllRootVariableRequest {
	s.VelocityVariableIds = &v
	return s
}

func (s *DescribeAllRootVariableRequest) Validate() error {
	return dara.Validate(s)
}
