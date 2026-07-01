// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySendStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *QuerySendStatisticsRequest
	GetEndDate() *string
	SetIsGlobe(v int32) *QuerySendStatisticsRequest
	GetIsGlobe() *int32
	SetOwnerId(v int64) *QuerySendStatisticsRequest
	GetOwnerId() *int64
	SetPageIndex(v int32) *QuerySendStatisticsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *QuerySendStatisticsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *QuerySendStatisticsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySendStatisticsRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *QuerySendStatisticsRequest
	GetSignName() *string
	SetStartDate(v string) *QuerySendStatisticsRequest
	GetStartDate() *string
	SetTemplateType(v int32) *QuerySendStatisticsRequest
	GetTemplateType() *int32
}

type QuerySendStatisticsRequest struct {
	// The end date. The format is yyyyMMdd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20201003
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The destination scope of the messages. Valid values:
	//
	// - **1**: domestic messages.
	//
	// - **2**: international messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	IsGlobe *int32 `json:"IsGlobe,omitempty" xml:"IsGlobe,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The page size. Valid values: **1 to 50**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature.
	//
	// example:
	//
	// 阿里云平台
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The start date. The format is yyyyMMdd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20201002
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The template type. Valid values:
	//
	// - **0**: verification code.
	//
	// - **1**: notification message.
	//
	// - **2**: promotional message. (Enterprise customers only)
	//
	// - **3**: international message. (Enterprise customers only)
	//
	// - **7**: digital message.
	//
	// example:
	//
	// 0
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s QuerySendStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySendStatisticsRequest) GoString() string {
	return s.String()
}

func (s *QuerySendStatisticsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QuerySendStatisticsRequest) GetIsGlobe() *int32 {
	return s.IsGlobe
}

func (s *QuerySendStatisticsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySendStatisticsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *QuerySendStatisticsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySendStatisticsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySendStatisticsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySendStatisticsRequest) GetSignName() *string {
	return s.SignName
}

func (s *QuerySendStatisticsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QuerySendStatisticsRequest) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *QuerySendStatisticsRequest) SetEndDate(v string) *QuerySendStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetIsGlobe(v int32) *QuerySendStatisticsRequest {
	s.IsGlobe = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetOwnerId(v int64) *QuerySendStatisticsRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetPageIndex(v int32) *QuerySendStatisticsRequest {
	s.PageIndex = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetPageSize(v int32) *QuerySendStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetResourceOwnerAccount(v string) *QuerySendStatisticsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetResourceOwnerId(v int64) *QuerySendStatisticsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetSignName(v string) *QuerySendStatisticsRequest {
	s.SignName = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetStartDate(v string) *QuerySendStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *QuerySendStatisticsRequest) SetTemplateType(v int32) *QuerySendStatisticsRequest {
	s.TemplateType = &v
	return s
}

func (s *QuerySendStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
