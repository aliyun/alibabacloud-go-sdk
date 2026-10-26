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
	// The reserved feature type parameter. This parameter is not used to filter results in the current query and does not need to be specified for column encryption calls.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// Specifies whether to restrict the results to recognition models supported in audit scenarios. If you set this parameter to true, models supported for auditing, such as mobile phone numbers, ID cards, and bank cards, are filtered. This parameter is optional for column encryption scenarios.
	FilterAuditModel *bool `json:"FilterAuditModel,omitempty" xml:"FilterAuditModel,omitempty"`
	// The language of the response description. Valid values: zh: Chinese. en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the data classification template. You can obtain the ID from Items[].Id in the response of the ListTemplates operation. The ID must be consistent with the selected sensitivity level and recognition model.
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
