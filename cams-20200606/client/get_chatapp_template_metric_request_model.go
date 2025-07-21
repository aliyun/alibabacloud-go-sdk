// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappTemplateMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetChatappTemplateMetricRequest
	GetCustSpaceId() *string
	SetEnd(v int64) *GetChatappTemplateMetricRequest
	GetEnd() *int64
	SetGranularity(v string) *GetChatappTemplateMetricRequest
	GetGranularity() *string
	SetIsvCode(v string) *GetChatappTemplateMetricRequest
	GetIsvCode() *string
	SetLanguage(v string) *GetChatappTemplateMetricRequest
	GetLanguage() *string
	SetOwnerId(v int64) *GetChatappTemplateMetricRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetChatappTemplateMetricRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetChatappTemplateMetricRequest
	GetResourceOwnerId() *int64
	SetStart(v int64) *GetChatappTemplateMetricRequest
	GetStart() *int64
	SetTemplateCode(v string) *GetChatappTemplateMetricRequest
	GetTemplateCode() *string
	SetTemplateType(v string) *GetChatappTemplateMetricRequest
	GetTemplateType() *string
}

type GetChatappTemplateMetricRequest struct {
	// The space ID of the RAM user within the ISV account.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1693407714687
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The granularity of the metric.
	//
	// Valid values:
	//
	// 	- DAILY
	//
	// 	- HALF_HOUR
	//
	// example:
	//
	// DAILY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The independent software vendor (ISV) verification code, which is used to verify whether the RAM user is authorized by the ISV account.
	//
	// example:
	//
	// skdi3kksloslikdkkdk
	IsvCode *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	// The template language.
	//
	// example:
	//
	// en
	Language             *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1693107714687
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The template code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca36bf5
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template type. If you do not specify this parameter, the default value WHATSAPP is used.
	//
	// Valid values:
	//
	// 	- VIBER
	//
	// 	- WHATSAPP
	//
	// example:
	//
	// WHATSAPP
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetChatappTemplateMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatappTemplateMetricRequest) GoString() string {
	return s.String()
}

func (s *GetChatappTemplateMetricRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetChatappTemplateMetricRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetChatappTemplateMetricRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *GetChatappTemplateMetricRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *GetChatappTemplateMetricRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetChatappTemplateMetricRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetChatappTemplateMetricRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetChatappTemplateMetricRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetChatappTemplateMetricRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetChatappTemplateMetricRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetChatappTemplateMetricRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetChatappTemplateMetricRequest) SetCustSpaceId(v string) *GetChatappTemplateMetricRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetEnd(v int64) *GetChatappTemplateMetricRequest {
	s.End = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetGranularity(v string) *GetChatappTemplateMetricRequest {
	s.Granularity = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetIsvCode(v string) *GetChatappTemplateMetricRequest {
	s.IsvCode = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetLanguage(v string) *GetChatappTemplateMetricRequest {
	s.Language = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetOwnerId(v int64) *GetChatappTemplateMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetResourceOwnerAccount(v string) *GetChatappTemplateMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetResourceOwnerId(v int64) *GetChatappTemplateMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetStart(v int64) *GetChatappTemplateMetricRequest {
	s.Start = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetTemplateCode(v string) *GetChatappTemplateMetricRequest {
	s.TemplateCode = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) SetTemplateType(v string) *GetChatappTemplateMetricRequest {
	s.TemplateType = &v
	return s
}

func (s *GetChatappTemplateMetricRequest) Validate() error {
	return dara.Validate(s)
}
