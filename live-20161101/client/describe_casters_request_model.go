// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCastersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCastersRequest
	GetCasterId() *string
	SetCasterName(v string) *DescribeCastersRequest
	GetCasterName() *string
	SetChargeType(v int32) *DescribeCastersRequest
	GetChargeType() *int32
	SetEndTime(v string) *DescribeCastersRequest
	GetEndTime() *string
	SetNormType(v string) *DescribeCastersRequest
	GetNormType() *string
	SetOrderByModifyAsc(v string) *DescribeCastersRequest
	GetOrderByModifyAsc() *string
	SetOwnerId(v int64) *DescribeCastersRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeCastersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeCastersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCastersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeCastersRequest
	GetResourceGroupId() *string
	SetStartTime(v string) *DescribeCastersRequest
	GetStartTime() *string
	SetStatus(v int32) *DescribeCastersRequest
	GetStatus() *int32
	SetTag(v []*DescribeCastersRequestTag) *DescribeCastersRequest
	GetTag() []*DescribeCastersRequestTag
}

type DescribeCastersRequest struct {
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster operation](https://help.aliyun.com/document_detail/2848012.html), check the CasterId parameter returned by the CreateCaster operation.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to **ApsaraVideo Live console > Production Studios > Cloud Production Studio*	- to view the ID.
	//
	// > - The production studio name in the production studio list on the Cloud Production Studio page is the production studio ID.
	//
	// > - If this parameter is left empty, the merged data of all production studios is returned by default.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The name of the production studio.
	//
	// example:
	//
	// liveCaster****
	CasterName *string `json:"CasterName,omitempty" xml:"CasterName,omitempty"`
	// The billing method. Valid values:
	//
	// - 0: PrePaid (subscription).
	//
	// - 1: PostPaid (pay-as-you-go).
	//
	// example:
	//
	// 0
	ChargeType *int32 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The end time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// example:
	//
	// 2016-06-29T11:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The specification type of the production studio. Valid values:
	//
	// - 1: general mode.
	//
	// - 3: lightweight playlist mode.
	//
	// - 4: virtual studio mode.
	//
	// - 6: playlist mode (new playlist mode production studio).
	//
	// example:
	//
	// 1
	NormType *string `json:"NormType,omitempty" xml:"NormType,omitempty"`
	// Specifies whether to sort the production studios in ascending order by modification time.
	//
	// Valid values: true (ascending order by modification time) | false (descending order by modification time, which is the default value).
	//
	// > If this parameter is not specified, the default value is "false".
	//
	// example:
	//
	// false
	OrderByModifyAsc *string `json:"OrderByModifyAsc,omitempty" xml:"OrderByModifyAsc,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. For more information about resource groups, see [What is a resource group](https://help.aliyun.com/document_detail/2381067.html).
	//
	// example:
	//
	// rg-aekzw******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// example:
	//
	// 2016-06-29T09:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status. Valid values:
	//
	// - 0: idle.
	//
	// - 1: streaming.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tag []*DescribeCastersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCastersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCastersRequest) GoString() string {
	return s.String()
}

func (s *DescribeCastersRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCastersRequest) GetCasterName() *string {
	return s.CasterName
}

func (s *DescribeCastersRequest) GetChargeType() *int32 {
	return s.ChargeType
}

func (s *DescribeCastersRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCastersRequest) GetNormType() *string {
	return s.NormType
}

func (s *DescribeCastersRequest) GetOrderByModifyAsc() *string {
	return s.OrderByModifyAsc
}

func (s *DescribeCastersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCastersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeCastersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCastersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCastersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCastersRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCastersRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCastersRequest) GetTag() []*DescribeCastersRequestTag {
	return s.Tag
}

func (s *DescribeCastersRequest) SetCasterId(v string) *DescribeCastersRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCastersRequest) SetCasterName(v string) *DescribeCastersRequest {
	s.CasterName = &v
	return s
}

func (s *DescribeCastersRequest) SetChargeType(v int32) *DescribeCastersRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeCastersRequest) SetEndTime(v string) *DescribeCastersRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCastersRequest) SetNormType(v string) *DescribeCastersRequest {
	s.NormType = &v
	return s
}

func (s *DescribeCastersRequest) SetOrderByModifyAsc(v string) *DescribeCastersRequest {
	s.OrderByModifyAsc = &v
	return s
}

func (s *DescribeCastersRequest) SetOwnerId(v int64) *DescribeCastersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCastersRequest) SetPageNum(v int32) *DescribeCastersRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeCastersRequest) SetPageSize(v int32) *DescribeCastersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCastersRequest) SetRegionId(v string) *DescribeCastersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCastersRequest) SetResourceGroupId(v string) *DescribeCastersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCastersRequest) SetStartTime(v string) *DescribeCastersRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCastersRequest) SetStatus(v int32) *DescribeCastersRequest {
	s.Status = &v
	return s
}

func (s *DescribeCastersRequest) SetTag(v []*DescribeCastersRequestTag) *DescribeCastersRequest {
	s.Tag = v
	return s
}

func (s *DescribeCastersRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCastersRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCastersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCastersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCastersRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCastersRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCastersRequestTag) SetKey(v string) *DescribeCastersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCastersRequestTag) SetValue(v string) *DescribeCastersRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeCastersRequestTag) Validate() error {
	return dara.Validate(s)
}
