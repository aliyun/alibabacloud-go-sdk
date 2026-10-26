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
	// The reserved feature type parameter. This parameter is not used to filter results in the current query. You do not need to specify this parameter for column encryption calls.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the response description. Valid values: zh: Chinese. en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the data classification template. You can obtain the ID from the Items[].Id parameter in the response of the ListTemplates operation. The template must be consistent with the selected sensitivity level and detection model.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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
