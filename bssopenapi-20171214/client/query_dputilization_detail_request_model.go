// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDPUtilizationDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *QueryDPUtilizationDetailRequest
	GetCommodityCode() *string
	SetDeductedInstanceId(v string) *QueryDPUtilizationDetailRequest
	GetDeductedInstanceId() *string
	SetEndTime(v string) *QueryDPUtilizationDetailRequest
	GetEndTime() *string
	SetIncludeShare(v bool) *QueryDPUtilizationDetailRequest
	GetIncludeShare() *bool
	SetInstanceId(v string) *QueryDPUtilizationDetailRequest
	GetInstanceId() *string
	SetInstanceSpec(v string) *QueryDPUtilizationDetailRequest
	GetInstanceSpec() *string
	SetLastToken(v string) *QueryDPUtilizationDetailRequest
	GetLastToken() *string
	SetLimit(v int32) *QueryDPUtilizationDetailRequest
	GetLimit() *int32
	SetProdCode(v string) *QueryDPUtilizationDetailRequest
	GetProdCode() *string
	SetStartTime(v string) *QueryDPUtilizationDetailRequest
	GetStartTime() *string
}

type QueryDPUtilizationDetailRequest struct {
	// The code of the resource, such as ecsRi and scu_bag. If this parameter is specified, the ProdCode parameter does not take effect for the request.
	//
	// example:
	//
	// ecsRi
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The ID of the deducted instance. If this parameter is not specified, the details of all instances are returned.
	//
	// example:
	//
	// oss-123123
	DeductedInstanceId *string `json:"DeductedInstanceId,omitempty" xml:"DeductedInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the YYYY-MM-DD HH:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-05-23 12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to query the resource plan usage of linked accounts. Valid values:
	//
	// 	- true: queries the resource plan usage of linked accounts.
	//
	// 	- false: does not query the resource plan usage of linked accounts.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IncludeShare *bool `json:"IncludeShare,omitempty" xml:"IncludeShare,omitempty"`
	// The ID of the instance to query. If this parameter is not specified, the details of all used instances are returned.
	//
	// example:
	//
	// oss-123123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type of the instance.
	//
	// example:
	//
	// Instancetyp
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The token that is used to retrieve the next page of results. For the first query, set the value to null. For subsequent queries, set the value to the token that is obtained from the NextToken parameter.
	//
	// example:
	//
	// CAESF***zNTAw
	LastToken *string `json:"LastToken,omitempty" xml:"LastToken,omitempty"`
	// The number of entries to return on each page. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The code of the service. Example: ecs.
	//
	// example:
	//
	// oss
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The beginning of the time range to query. Specify the time in the YYYY-MM-DD HH:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-05-23 12:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryDPUtilizationDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDPUtilizationDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryDPUtilizationDetailRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QueryDPUtilizationDetailRequest) GetDeductedInstanceId() *string {
	return s.DeductedInstanceId
}

func (s *QueryDPUtilizationDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryDPUtilizationDetailRequest) GetIncludeShare() *bool {
	return s.IncludeShare
}

func (s *QueryDPUtilizationDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryDPUtilizationDetailRequest) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *QueryDPUtilizationDetailRequest) GetLastToken() *string {
	return s.LastToken
}

func (s *QueryDPUtilizationDetailRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *QueryDPUtilizationDetailRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *QueryDPUtilizationDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryDPUtilizationDetailRequest) SetCommodityCode(v string) *QueryDPUtilizationDetailRequest {
	s.CommodityCode = &v
	return s
}

func (s *QueryDPUtilizationDetailRequest) SetDeductedInstanceId(v string) *QueryDPUtilizationDetailRequest {
	s.DeductedInstanceId = &v
	return s
}

func (s *QueryDPUtilizationDetailRequest) SetEndTime(v string) *QueryDPUtilizationDetailRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDPUtilizationDetailRequest) SetIncludeShare(v bool) *QueryDPUtilizationDetailRequest {
	s.IncludeShare = &v
	return s
}

func (s *QueryDPUtilizationDetailRequest) SetInstanceId(v string) *QueryDPUtilizationDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryDPUtilizationDetailRequest) SetInstanceSpec(v string) *QueryDPUtilizationDetailRequest {
	s.InstanceSpec = &v
	return s
}

func (s *QueryDPUtilizationDetailRequest) SetLastToken(v string) *QueryDPUtilizationDetailRequest {
	s.LastToken = &v
	return s
}

func (s *QueryDPUtilizationDetailRequest) SetLimit(v int32) *QueryDPUtilizationDetailRequest {
	s.Limit = &v
	return s
}

func (s *QueryDPUtilizationDetailRequest) SetProdCode(v string) *QueryDPUtilizationDetailRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryDPUtilizationDetailRequest) SetStartTime(v string) *QueryDPUtilizationDetailRequest {
	s.StartTime = &v
	return s
}

func (s *QueryDPUtilizationDetailRequest) Validate() error {
	return dara.Validate(s)
}
