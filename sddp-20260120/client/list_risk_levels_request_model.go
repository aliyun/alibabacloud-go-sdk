// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRiskLevelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *ListRiskLevelsRequest
	GetFeatureType() *int32
	SetLang(v string) *ListRiskLevelsRequest
	GetLang() *string
	SetTemplateId(v int64) *ListRiskLevelsRequest
	GetTemplateId() *int64
}

type ListRiskLevelsRequest struct {
	FeatureType *int32  `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	TemplateId  *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListRiskLevelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRiskLevelsRequest) GoString() string {
	return s.String()
}

func (s *ListRiskLevelsRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *ListRiskLevelsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListRiskLevelsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListRiskLevelsRequest) SetFeatureType(v int32) *ListRiskLevelsRequest {
	s.FeatureType = &v
	return s
}

func (s *ListRiskLevelsRequest) SetLang(v string) *ListRiskLevelsRequest {
	s.Lang = &v
	return s
}

func (s *ListRiskLevelsRequest) SetTemplateId(v int64) *ListRiskLevelsRequest {
	s.TemplateId = &v
	return s
}

func (s *ListRiskLevelsRequest) Validate() error {
	return dara.Validate(s)
}
