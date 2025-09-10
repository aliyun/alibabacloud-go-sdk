// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaAlarmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmName(v string) *ListQuotaAlarmsRequest
	GetAlarmName() *string
	SetMaxResults(v int32) *ListQuotaAlarmsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListQuotaAlarmsRequest
	GetNextToken() *string
	SetProductCode(v string) *ListQuotaAlarmsRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *ListQuotaAlarmsRequest
	GetQuotaActionCode() *string
	SetQuotaDimensions(v []*ListQuotaAlarmsRequestQuotaDimensions) *ListQuotaAlarmsRequest
	GetQuotaDimensions() []*ListQuotaAlarmsRequestQuotaDimensions
}

type ListQuotaAlarmsRequest struct {
	// The name of the alert.
	//
	// example:
	//
	// rules
	AlarmName *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	// The maximum number of records that you want to return for the query.
	//
	// Valid values: 1 to 200. Default value: 30.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query.
	//
	// > An empty value indicates that the query starts from the beginning.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// >  To query the abbreviation of an Alibaba Cloud service name, call the [ListProducts](https://help.aliyun.com/document_detail/440554.html) operation and check the value of `ProductCode` in the response.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// >
	//
	// 	- To obtain the quota ID of a cloud service, call the [ListProductQuotas](https://help.aliyun.com/document_detail/440554.html) operation and check the value of `QuotaActionCode` in the response.
	//
	// 	- If you specify this parameter, you must specify `ProductCode`.
	//
	// example:
	//
	// q_hvnoqv
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The quota dimensions.
	QuotaDimensions []*ListQuotaAlarmsRequestQuotaDimensions `json:"QuotaDimensions,omitempty" xml:"QuotaDimensions,omitempty" type:"Repeated"`
}

func (s ListQuotaAlarmsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaAlarmsRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsRequest) GetAlarmName() *string {
	return s.AlarmName
}

func (s *ListQuotaAlarmsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQuotaAlarmsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQuotaAlarmsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListQuotaAlarmsRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListQuotaAlarmsRequest) GetQuotaDimensions() []*ListQuotaAlarmsRequestQuotaDimensions {
	return s.QuotaDimensions
}

func (s *ListQuotaAlarmsRequest) SetAlarmName(v string) *ListQuotaAlarmsRequest {
	s.AlarmName = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetMaxResults(v int32) *ListQuotaAlarmsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetNextToken(v string) *ListQuotaAlarmsRequest {
	s.NextToken = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetProductCode(v string) *ListQuotaAlarmsRequest {
	s.ProductCode = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetQuotaActionCode(v string) *ListQuotaAlarmsRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ListQuotaAlarmsRequest) SetQuotaDimensions(v []*ListQuotaAlarmsRequestQuotaDimensions) *ListQuotaAlarmsRequest {
	s.QuotaDimensions = v
	return s
}

func (s *ListQuotaAlarmsRequest) Validate() error {
	return dara.Validate(s)
}

type ListQuotaAlarmsRequestQuotaDimensions struct {
	// The key of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- This parameter is required if you set the `ProductCode` parameter to `ecs`, `ecs-spec`, `actiontrail`, or `ess`.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the dimension.
	//
	// >
	//
	// 	- The value range of N varies based on the number of dimensions that are supported by the related Alibaba Cloud service.
	//
	// 	- This parameter is required if you set the `ProductCode` parameter to `ecs`, `ecs-spec`, `actiontrail`, or `ess`.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQuotaAlarmsRequestQuotaDimensions) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaAlarmsRequestQuotaDimensions) GoString() string {
	return s.String()
}

func (s *ListQuotaAlarmsRequestQuotaDimensions) GetKey() *string {
	return s.Key
}

func (s *ListQuotaAlarmsRequestQuotaDimensions) GetValue() *string {
	return s.Value
}

func (s *ListQuotaAlarmsRequestQuotaDimensions) SetKey(v string) *ListQuotaAlarmsRequestQuotaDimensions {
	s.Key = &v
	return s
}

func (s *ListQuotaAlarmsRequestQuotaDimensions) SetValue(v string) *ListQuotaAlarmsRequestQuotaDimensions {
	s.Value = &v
	return s
}

func (s *ListQuotaAlarmsRequestQuotaDimensions) Validate() error {
	return dara.Validate(s)
}
