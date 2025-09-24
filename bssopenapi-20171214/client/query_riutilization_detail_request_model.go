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
	// The ID of the instance whose fees are deducted by using the RI. If this parameter is left empty, the usage details of all instances are queried.
	//
	// example:
	//
	// jsdgfsdhgsdjh
	DeductedInstanceId *string `json:"DeductedInstanceId,omitempty" xml:"DeductedInstanceId,omitempty"`
	// The time when the RI expires. Specify the time in the YYYY-MM-DD HH:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-05-23 12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance type of the RI.
	//
	// example:
	//
	// Instancetyp
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The code of the service to which the RI is applied. Default value: ecsRi. Valid values:
	//
	// 	- ecsRi: ECS RI.
	//
	// 	- scu_bag: storage capacity unit (SCU).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecsRi
	RICommodityCode *string `json:"RICommodityCode,omitempty" xml:"RICommodityCode,omitempty"`
	// The ID of the RI. If this parameter is left empty, the usage details of all RIs are queried.
	//
	// example:
	//
	// dsudfgdsjh
	RIInstanceId *string `json:"RIInstanceId,omitempty" xml:"RIInstanceId,omitempty"`
	// The time when the RI was created. Specify the time in the YYYY-MM-DD HH:mm:ss format.
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
