// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentifyModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *ListIdentifyModelsRequest
	GetFeatureType() *int32
	SetFilterAuditModel(v bool) *ListIdentifyModelsRequest
	GetFilterAuditModel() *bool
	SetLang(v string) *ListIdentifyModelsRequest
	GetLang() *string
	SetTemplateId(v int64) *ListIdentifyModelsRequest
	GetTemplateId() *int64
}

type ListIdentifyModelsRequest struct {
	// A reserved feature type parameter. This parameter is not used to filter results in the current query. You do not need to specify this parameter for column encryption calls.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// Specifies whether to return only detection models supported in audit scenarios. If this parameter is set to true, only audit-supported models such as phone number, ID card, and bank card models are returned. You do not need to specify this parameter for column encryption scenarios.
	FilterAuditModel *bool `json:"FilterAuditModel,omitempty" xml:"FilterAuditModel,omitempty"`
	// The language of the response description. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the data classification template. You can obtain this value from Items[].Id in the response of the ListTemplates operation. The template must be consistent with the selected sensitivity level and detection models.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListIdentifyModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIdentifyModelsRequest) GoString() string {
	return s.String()
}

func (s *ListIdentifyModelsRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *ListIdentifyModelsRequest) GetFilterAuditModel() *bool {
	return s.FilterAuditModel
}

func (s *ListIdentifyModelsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListIdentifyModelsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListIdentifyModelsRequest) SetFeatureType(v int32) *ListIdentifyModelsRequest {
	s.FeatureType = &v
	return s
}

func (s *ListIdentifyModelsRequest) SetFilterAuditModel(v bool) *ListIdentifyModelsRequest {
	s.FilterAuditModel = &v
	return s
}

func (s *ListIdentifyModelsRequest) SetLang(v string) *ListIdentifyModelsRequest {
	s.Lang = &v
	return s
}

func (s *ListIdentifyModelsRequest) SetTemplateId(v int64) *ListIdentifyModelsRequest {
	s.TemplateId = &v
	return s
}

func (s *ListIdentifyModelsRequest) Validate() error {
	return dara.Validate(s)
}
