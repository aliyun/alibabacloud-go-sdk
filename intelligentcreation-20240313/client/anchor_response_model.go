// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnchorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetAnchorCategory(v string) *AnchorResponse
	GetAnchorCategory() *string
	SetAnchorId(v string) *AnchorResponse
	GetAnchorId() *string
	SetAnchorMaterialName(v string) *AnchorResponse
	GetAnchorMaterialName() *string
	SetAnchorType(v string) *AnchorResponse
	GetAnchorType() *string
	SetCoverHeight(v int32) *AnchorResponse
	GetCoverHeight() *int32
	SetCoverRate(v string) *AnchorResponse
	GetCoverRate() *string
	SetCoverThumbnailUrl(v string) *AnchorResponse
	GetCoverThumbnailUrl() *string
	SetCoverUrl(v string) *AnchorResponse
	GetCoverUrl() *string
	SetCoverWeight(v int32) *AnchorResponse
	GetCoverWeight() *int32
	SetDigitalHumanType(v string) *AnchorResponse
	GetDigitalHumanType() *string
	SetGender(v string) *AnchorResponse
	GetGender() *string
	SetResourceTypeDesc(v string) *AnchorResponse
	GetResourceTypeDesc() *string
	SetStatus(v string) *AnchorResponse
	GetStatus() *string
	SetSupportBgChange(v int32) *AnchorResponse
	GetSupportBgChange() *int32
	SetUseScene(v string) *AnchorResponse
	GetUseScene() *string
}

type AnchorResponse struct {
	AnchorCategory     *string `json:"anchorCategory,omitempty" xml:"anchorCategory,omitempty"`
	AnchorId           *string `json:"anchorId,omitempty" xml:"anchorId,omitempty"`
	AnchorMaterialName *string `json:"anchorMaterialName,omitempty" xml:"anchorMaterialName,omitempty"`
	AnchorType         *string `json:"anchorType,omitempty" xml:"anchorType,omitempty"`
	CoverHeight        *int32  `json:"coverHeight,omitempty" xml:"coverHeight,omitempty"`
	CoverRate          *string `json:"coverRate,omitempty" xml:"coverRate,omitempty"`
	CoverThumbnailUrl  *string `json:"coverThumbnailUrl,omitempty" xml:"coverThumbnailUrl,omitempty"`
	CoverUrl           *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	CoverWeight        *int32  `json:"coverWeight,omitempty" xml:"coverWeight,omitempty"`
	DigitalHumanType   *string `json:"digitalHumanType,omitempty" xml:"digitalHumanType,omitempty"`
	Gender             *string `json:"gender,omitempty" xml:"gender,omitempty"`
	ResourceTypeDesc   *string `json:"resourceTypeDesc,omitempty" xml:"resourceTypeDesc,omitempty"`
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
	SupportBgChange    *int32  `json:"supportBgChange,omitempty" xml:"supportBgChange,omitempty"`
	UseScene           *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
}

func (s AnchorResponse) String() string {
	return dara.Prettify(s)
}

func (s AnchorResponse) GoString() string {
	return s.String()
}

func (s *AnchorResponse) GetAnchorCategory() *string {
	return s.AnchorCategory
}

func (s *AnchorResponse) GetAnchorId() *string {
	return s.AnchorId
}

func (s *AnchorResponse) GetAnchorMaterialName() *string {
	return s.AnchorMaterialName
}

func (s *AnchorResponse) GetAnchorType() *string {
	return s.AnchorType
}

func (s *AnchorResponse) GetCoverHeight() *int32 {
	return s.CoverHeight
}

func (s *AnchorResponse) GetCoverRate() *string {
	return s.CoverRate
}

func (s *AnchorResponse) GetCoverThumbnailUrl() *string {
	return s.CoverThumbnailUrl
}

func (s *AnchorResponse) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *AnchorResponse) GetCoverWeight() *int32 {
	return s.CoverWeight
}

func (s *AnchorResponse) GetDigitalHumanType() *string {
	return s.DigitalHumanType
}

func (s *AnchorResponse) GetGender() *string {
	return s.Gender
}

func (s *AnchorResponse) GetResourceTypeDesc() *string {
	return s.ResourceTypeDesc
}

func (s *AnchorResponse) GetStatus() *string {
	return s.Status
}

func (s *AnchorResponse) GetSupportBgChange() *int32 {
	return s.SupportBgChange
}

func (s *AnchorResponse) GetUseScene() *string {
	return s.UseScene
}

func (s *AnchorResponse) SetAnchorCategory(v string) *AnchorResponse {
	s.AnchorCategory = &v
	return s
}

func (s *AnchorResponse) SetAnchorId(v string) *AnchorResponse {
	s.AnchorId = &v
	return s
}

func (s *AnchorResponse) SetAnchorMaterialName(v string) *AnchorResponse {
	s.AnchorMaterialName = &v
	return s
}

func (s *AnchorResponse) SetAnchorType(v string) *AnchorResponse {
	s.AnchorType = &v
	return s
}

func (s *AnchorResponse) SetCoverHeight(v int32) *AnchorResponse {
	s.CoverHeight = &v
	return s
}

func (s *AnchorResponse) SetCoverRate(v string) *AnchorResponse {
	s.CoverRate = &v
	return s
}

func (s *AnchorResponse) SetCoverThumbnailUrl(v string) *AnchorResponse {
	s.CoverThumbnailUrl = &v
	return s
}

func (s *AnchorResponse) SetCoverUrl(v string) *AnchorResponse {
	s.CoverUrl = &v
	return s
}

func (s *AnchorResponse) SetCoverWeight(v int32) *AnchorResponse {
	s.CoverWeight = &v
	return s
}

func (s *AnchorResponse) SetDigitalHumanType(v string) *AnchorResponse {
	s.DigitalHumanType = &v
	return s
}

func (s *AnchorResponse) SetGender(v string) *AnchorResponse {
	s.Gender = &v
	return s
}

func (s *AnchorResponse) SetResourceTypeDesc(v string) *AnchorResponse {
	s.ResourceTypeDesc = &v
	return s
}

func (s *AnchorResponse) SetStatus(v string) *AnchorResponse {
	s.Status = &v
	return s
}

func (s *AnchorResponse) SetSupportBgChange(v int32) *AnchorResponse {
	s.SupportBgChange = &v
	return s
}

func (s *AnchorResponse) SetUseScene(v string) *AnchorResponse {
	s.UseScene = &v
	return s
}

func (s *AnchorResponse) Validate() error {
	return dara.Validate(s)
}
