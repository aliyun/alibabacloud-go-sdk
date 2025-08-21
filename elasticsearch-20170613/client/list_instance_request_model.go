// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ListInstanceRequest
	GetDescription() *string
	SetEsVersion(v string) *ListInstanceRequest
	GetEsVersion() *string
	SetInstanceCategory(v string) *ListInstanceRequest
	GetInstanceCategory() *string
	SetInstanceId(v string) *ListInstanceRequest
	GetInstanceId() *string
	SetPage(v int32) *ListInstanceRequest
	GetPage() *int32
	SetPaymentType(v string) *ListInstanceRequest
	GetPaymentType() *string
	SetResourceGroupId(v string) *ListInstanceRequest
	GetResourceGroupId() *string
	SetSize(v int32) *ListInstanceRequest
	GetSize() *int32
	SetStatus(v string) *ListInstanceRequest
	GetStatus() *string
	SetTags(v string) *ListInstanceRequest
	GetTags() *string
	SetVpcId(v string) *ListInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *ListInstanceRequest
	GetZoneId() *string
}

type ListInstanceRequest struct {
	// cn-hangzhou-i
	//
	// example:
	//
	// aliyunes_test1
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// advanced
	//
	// example:
	//
	// 6.7_with_X-Pack
	EsVersion *string `json:"esVersion,omitempty" xml:"esVersion,omitempty"`
	// The number of data nodes.
	//
	// example:
	//
	// advanced
	InstanceCategory *string `json:"instanceCategory,omitempty" xml:"instanceCategory,omitempty"`
	// postpaid
	//
	// example:
	//
	// es-cn-v641a0ta3000g****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// [{"tagKey":"key1","tagValue":"value1"}]
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// Specifies whether to include dedicated master nodes. Valid values:
	//
	// 	- true: The files contain data that is dumped to the IA storage medium.
	//
	// 	- false: The files do not contain data that is dumped to the IA storage medium.
	//
	// example:
	//
	// postpaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// rg-aekzvowej3i****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// vpc-bp16k1dvzxtmagcva\\*\\*\\*\\*
	//
	// example:
	//
	// 10
	Size   *int32  `json:"size,omitempty" xml:"size,omitempty"`
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The header of the response.
	//
	// example:
	//
	// [{"tagKey":"key1","tagValue":"value1"}]
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// vpc-bp16k1dvzxtmagcva****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The returned data.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *ListInstanceRequest) GetEsVersion() *string {
	return s.EsVersion
}

func (s *ListInstanceRequest) GetInstanceCategory() *string {
	return s.InstanceCategory
}

func (s *ListInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstanceRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListInstanceRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceRequest) GetTags() *string {
	return s.Tags
}

func (s *ListInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListInstanceRequest) SetDescription(v string) *ListInstanceRequest {
	s.Description = &v
	return s
}

func (s *ListInstanceRequest) SetEsVersion(v string) *ListInstanceRequest {
	s.EsVersion = &v
	return s
}

func (s *ListInstanceRequest) SetInstanceCategory(v string) *ListInstanceRequest {
	s.InstanceCategory = &v
	return s
}

func (s *ListInstanceRequest) SetInstanceId(v string) *ListInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceRequest) SetPage(v int32) *ListInstanceRequest {
	s.Page = &v
	return s
}

func (s *ListInstanceRequest) SetPaymentType(v string) *ListInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *ListInstanceRequest) SetResourceGroupId(v string) *ListInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstanceRequest) SetSize(v int32) *ListInstanceRequest {
	s.Size = &v
	return s
}

func (s *ListInstanceRequest) SetStatus(v string) *ListInstanceRequest {
	s.Status = &v
	return s
}

func (s *ListInstanceRequest) SetTags(v string) *ListInstanceRequest {
	s.Tags = &v
	return s
}

func (s *ListInstanceRequest) SetVpcId(v string) *ListInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *ListInstanceRequest) SetZoneId(v string) *ListInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *ListInstanceRequest) Validate() error {
	return dara.Validate(s)
}
