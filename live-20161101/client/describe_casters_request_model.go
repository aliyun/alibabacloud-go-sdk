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
	// The ID of the production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848012.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the ApsaraVideo Live console and click Production Studios in the left-side navigation pane.
	//
	// >
	//
	// 	- You can find the ID of the production studio in the Instance ID/Name column.
	//
	// 	- If you leave this parameter empty, the data of all production studios is returned.
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
	// 	- 0: the subscription billing method
	//
	// 	- 1: the pay-as-you-go billing method
	//
	// example:
	//
	// 0
	ChargeType *int32 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2016-06-29T11:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the production studio. Valid values:
	//
	// 	- 1: general mode
	//
	// 	- 3: lightweight carousel playback mode
	//
	// 	- 4: virtual studio
	//
	// 	- 6: playlist mode
	//
	// example:
	//
	// 1
	NormType *string `json:"NormType,omitempty" xml:"NormType,omitempty"`
	// Specifies whether to sort the production studios in ascending order based on the modification time.
	//
	// >  If you leave this parameter empty, the default value is used. Default value: false.
	//
	// example:
	//
	// false
	OrderByModifyAsc *string `json:"OrderByModifyAsc,omitempty" xml:"OrderByModifyAsc,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of templates to return on each page. If you leave this parameter empty, the default value is used. Default value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/2381067.html).
	//
	// example:
	//
	// rg-aekzw******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2016-06-29T09:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the production studio. Valid values:
	//
	// 	- 0: idle
	//
	// 	- 1: streaming
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
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
	return dara.Validate(s)
}

type DescribeCastersRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
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
