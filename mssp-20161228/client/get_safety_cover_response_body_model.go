// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSafetyCoverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSafetyCoverResponseBody
	GetCode() *string
	SetData(v *GetSafetyCoverResponseBodyData) *GetSafetyCoverResponseBody
	GetData() *GetSafetyCoverResponseBodyData
	SetHttpStatusCode(v int32) *GetSafetyCoverResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSafetyCoverResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSafetyCoverResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSafetyCoverResponseBody
	GetSuccess() *bool
}

type GetSafetyCoverResponseBody struct {
	// API return code.
	//
	// example:
	//
	// 404
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data query result.
	Data *GetSafetyCoverResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Message of the response result.
	//
	// example:
	//
	// system error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 564f8bb9-df3c-42a0-877a-b35d48f66603
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful:
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSafetyCoverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSafetyCoverResponseBody) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSafetyCoverResponseBody) GetData() *GetSafetyCoverResponseBodyData {
	return s.Data
}

func (s *GetSafetyCoverResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSafetyCoverResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSafetyCoverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSafetyCoverResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSafetyCoverResponseBody) SetCode(v string) *GetSafetyCoverResponseBody {
	s.Code = &v
	return s
}

func (s *GetSafetyCoverResponseBody) SetData(v *GetSafetyCoverResponseBodyData) *GetSafetyCoverResponseBody {
	s.Data = v
	return s
}

func (s *GetSafetyCoverResponseBody) SetHttpStatusCode(v int32) *GetSafetyCoverResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSafetyCoverResponseBody) SetMessage(v string) *GetSafetyCoverResponseBody {
	s.Message = &v
	return s
}

func (s *GetSafetyCoverResponseBody) SetRequestId(v string) *GetSafetyCoverResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSafetyCoverResponseBody) SetSuccess(v bool) *GetSafetyCoverResponseBody {
	s.Success = &v
	return s
}

func (s *GetSafetyCoverResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSafetyCoverResponseBodyData struct {
	// CFW protection coverage.
	CfwProtection *GetSafetyCoverResponseBodyDataCfwProtection `json:"CfwProtection,omitempty" xml:"CfwProtection,omitempty" type:"Struct"`
	// ECS protection coverage.
	EcsProtection *GetSafetyCoverResponseBodyDataEcsProtection `json:"EcsProtection,omitempty" xml:"EcsProtection,omitempty" type:"Struct"`
	// WAF protection coverage.
	WafProtection *GetSafetyCoverResponseBodyDataWafProtection `json:"WafProtection,omitempty" xml:"WafProtection,omitempty" type:"Struct"`
}

func (s GetSafetyCoverResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSafetyCoverResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBodyData) GetCfwProtection() *GetSafetyCoverResponseBodyDataCfwProtection {
	return s.CfwProtection
}

func (s *GetSafetyCoverResponseBodyData) GetEcsProtection() *GetSafetyCoverResponseBodyDataEcsProtection {
	return s.EcsProtection
}

func (s *GetSafetyCoverResponseBodyData) GetWafProtection() *GetSafetyCoverResponseBodyDataWafProtection {
	return s.WafProtection
}

func (s *GetSafetyCoverResponseBodyData) SetCfwProtection(v *GetSafetyCoverResponseBodyDataCfwProtection) *GetSafetyCoverResponseBodyData {
	s.CfwProtection = v
	return s
}

func (s *GetSafetyCoverResponseBodyData) SetEcsProtection(v *GetSafetyCoverResponseBodyDataEcsProtection) *GetSafetyCoverResponseBodyData {
	s.EcsProtection = v
	return s
}

func (s *GetSafetyCoverResponseBodyData) SetWafProtection(v *GetSafetyCoverResponseBodyDataWafProtection) *GetSafetyCoverResponseBodyData {
	s.WafProtection = v
	return s
}

func (s *GetSafetyCoverResponseBodyData) Validate() error {
	if s.CfwProtection != nil {
		if err := s.CfwProtection.Validate(); err != nil {
			return err
		}
	}
	if s.EcsProtection != nil {
		if err := s.EcsProtection.Validate(); err != nil {
			return err
		}
	}
	if s.WafProtection != nil {
		if err := s.WafProtection.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSafetyCoverResponseBodyDataCfwProtection struct {
	// Number of unprotected items.
	//
	// example:
	//
	// 5
	NoProtectionCount *int64 `json:"NoProtectionCount,omitempty" xml:"NoProtectionCount,omitempty"`
	// Number of protected items.
	//
	// example:
	//
	// 5
	ProtectionCount *int64 `json:"ProtectionCount,omitempty" xml:"ProtectionCount,omitempty"`
	// Year-over-year protection rate.
	//
	// example:
	//
	// 35.00
	ProtectionGrowthRate *string `json:"ProtectionGrowthRate,omitempty" xml:"ProtectionGrowthRate,omitempty"`
	// Protection rate.
	//
	// example:
	//
	// 50.00
	ProtectionRate *string `json:"ProtectionRate,omitempty" xml:"ProtectionRate,omitempty"`
	// Total quantity.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSafetyCoverResponseBodyDataCfwProtection) String() string {
	return dara.Prettify(s)
}

func (s GetSafetyCoverResponseBodyDataCfwProtection) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) GetNoProtectionCount() *int64 {
	return s.NoProtectionCount
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) GetProtectionCount() *int64 {
	return s.ProtectionCount
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) GetProtectionGrowthRate() *string {
	return s.ProtectionGrowthRate
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) GetProtectionRate() *string {
	return s.ProtectionRate
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetNoProtectionCount(v int64) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.NoProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetProtectionCount(v int64) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.ProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetProtectionGrowthRate(v string) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.ProtectionGrowthRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetProtectionRate(v string) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.ProtectionRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) SetTotalCount(v int64) *GetSafetyCoverResponseBodyDataCfwProtection {
	s.TotalCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataCfwProtection) Validate() error {
	return dara.Validate(s)
}

type GetSafetyCoverResponseBodyDataEcsProtection struct {
	// Number of unprotected items.
	//
	// example:
	//
	// 5
	NoProtectionCount *int64 `json:"NoProtectionCount,omitempty" xml:"NoProtectionCount,omitempty"`
	// Number of protected items.
	//
	// example:
	//
	// 5
	ProtectionCount *int64 `json:"ProtectionCount,omitempty" xml:"ProtectionCount,omitempty"`
	// Year-over-year growth in protection rate.
	//
	// example:
	//
	// 35.00
	ProtectionGrowthRate *string `json:"ProtectionGrowthRate,omitempty" xml:"ProtectionGrowthRate,omitempty"`
	// Protection rate.
	//
	// example:
	//
	// 50.00
	ProtectionRate *string `json:"ProtectionRate,omitempty" xml:"ProtectionRate,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSafetyCoverResponseBodyDataEcsProtection) String() string {
	return dara.Prettify(s)
}

func (s GetSafetyCoverResponseBodyDataEcsProtection) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) GetNoProtectionCount() *int64 {
	return s.NoProtectionCount
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) GetProtectionCount() *int64 {
	return s.ProtectionCount
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) GetProtectionGrowthRate() *string {
	return s.ProtectionGrowthRate
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) GetProtectionRate() *string {
	return s.ProtectionRate
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetNoProtectionCount(v int64) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.NoProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetProtectionCount(v int64) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.ProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetProtectionGrowthRate(v string) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.ProtectionGrowthRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetProtectionRate(v string) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.ProtectionRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) SetTotalCount(v int64) *GetSafetyCoverResponseBodyDataEcsProtection {
	s.TotalCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataEcsProtection) Validate() error {
	return dara.Validate(s)
}

type GetSafetyCoverResponseBodyDataWafProtection struct {
	// Number of unprotected items.
	//
	// example:
	//
	// 5
	NoProtectionCount *int64 `json:"NoProtectionCount,omitempty" xml:"NoProtectionCount,omitempty"`
	// Number of protected items.
	//
	// example:
	//
	// 5
	ProtectionCount *int64 `json:"ProtectionCount,omitempty" xml:"ProtectionCount,omitempty"`
	// Year-over-year growth in protection rate.
	//
	// example:
	//
	// 35.00
	ProtectionGrowthRate *string `json:"ProtectionGrowthRate,omitempty" xml:"ProtectionGrowthRate,omitempty"`
	// Protection rate.
	//
	// example:
	//
	// 50.00
	ProtectionRate *string `json:"ProtectionRate,omitempty" xml:"ProtectionRate,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSafetyCoverResponseBodyDataWafProtection) String() string {
	return dara.Prettify(s)
}

func (s GetSafetyCoverResponseBodyDataWafProtection) GoString() string {
	return s.String()
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) GetNoProtectionCount() *int64 {
	return s.NoProtectionCount
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) GetProtectionCount() *int64 {
	return s.ProtectionCount
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) GetProtectionGrowthRate() *string {
	return s.ProtectionGrowthRate
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) GetProtectionRate() *string {
	return s.ProtectionRate
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetNoProtectionCount(v int64) *GetSafetyCoverResponseBodyDataWafProtection {
	s.NoProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetProtectionCount(v int64) *GetSafetyCoverResponseBodyDataWafProtection {
	s.ProtectionCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetProtectionGrowthRate(v string) *GetSafetyCoverResponseBodyDataWafProtection {
	s.ProtectionGrowthRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetProtectionRate(v string) *GetSafetyCoverResponseBodyDataWafProtection {
	s.ProtectionRate = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) SetTotalCount(v int64) *GetSafetyCoverResponseBodyDataWafProtection {
	s.TotalCount = &v
	return s
}

func (s *GetSafetyCoverResponseBodyDataWafProtection) Validate() error {
	return dara.Validate(s)
}
