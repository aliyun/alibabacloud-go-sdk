// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRIUtilizationDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeductedInstanceId(v string) *QueryRIUtilizationDetailRequest
	GetDeductedInstanceId() *string
	SetEndTime(v string) *QueryRIUtilizationDetailRequest
	GetEndTime() *string
	SetInstanceSpec(v string) *QueryRIUtilizationDetailRequest
	GetInstanceSpec() *string
	SetPageNum(v int32) *QueryRIUtilizationDetailRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryRIUtilizationDetailRequest
	GetPageSize() *int32
	SetRICommodityCode(v string) *QueryRIUtilizationDetailRequest
	GetRICommodityCode() *string
	SetRIInstanceId(v string) *QueryRIUtilizationDetailRequest
	GetRIInstanceId() *string
	SetStartTime(v string) *QueryRIUtilizationDetailRequest
	GetStartTime() *string
}

type QueryRIUtilizationDetailRequest struct {
	// The ID of the deducted instance. If this parameter is left empty, the details of all instances are returned.
	//
	// example:
	//
	// jsdgfsdhgsdjh
	DeductedInstanceId *string `json:"DeductedInstanceId,omitempty" xml:"DeductedInstanceId,omitempty"`
	// The end time. Format: YYYY-MM-DD HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-05-23 12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance type corresponding to the reserved instance.
	//
	// example:
	//
	// Instancetyp
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Valid values:
	//
	// - ecsRi: ECS reserved instances. This is the default value.
	//
	// - scu_bag: Storage capacity unit (SCU) plans.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecsRi
	RICommodityCode *string `json:"RICommodityCode,omitempty" xml:"RICommodityCode,omitempty"`
	// The instance ID of the reserved instance for the filtered query. If this parameter is left empty, the details of all used reserved instances are returned.
	//
	// example:
	//
	// dsudfgdsjh
	RIInstanceId *string `json:"RIInstanceId,omitempty" xml:"RIInstanceId,omitempty"`
	// The start time. Format: YYYY-MM-DD HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-05-23 12:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryRIUtilizationDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRIUtilizationDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRIUtilizationDetailRequest) GetDeductedInstanceId() *string {
	return s.DeductedInstanceId
}

func (s *QueryRIUtilizationDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryRIUtilizationDetailRequest) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *QueryRIUtilizationDetailRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryRIUtilizationDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRIUtilizationDetailRequest) GetRICommodityCode() *string {
	return s.RICommodityCode
}

func (s *QueryRIUtilizationDetailRequest) GetRIInstanceId() *string {
	return s.RIInstanceId
}

func (s *QueryRIUtilizationDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryRIUtilizationDetailRequest) SetDeductedInstanceId(v string) *QueryRIUtilizationDetailRequest {
	s.DeductedInstanceId = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetEndTime(v string) *QueryRIUtilizationDetailRequest {
	s.EndTime = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetInstanceSpec(v string) *QueryRIUtilizationDetailRequest {
	s.InstanceSpec = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetPageNum(v int32) *QueryRIUtilizationDetailRequest {
	s.PageNum = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetPageSize(v int32) *QueryRIUtilizationDetailRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetRICommodityCode(v string) *QueryRIUtilizationDetailRequest {
	s.RICommodityCode = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetRIInstanceId(v string) *QueryRIUtilizationDetailRequest {
	s.RIInstanceId = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetStartTime(v string) *QueryRIUtilizationDetailRequest {
	s.StartTime = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) Validate() error {
	return dara.Validate(s)
}
