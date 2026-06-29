// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateView interface {
	dara.Model
	String() string
	GoString() string
	SetFields(v []*TemplateViewFields) *TemplateView
	GetFields() []*TemplateViewFields
}

type TemplateView struct {
	// View List
	Fields []*TemplateViewFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
}

func (s TemplateView) String() string {
	return dara.Prettify(s)
}

func (s TemplateView) GoString() string {
	return s.String()
}

func (s *TemplateView) GetFields() []*TemplateViewFields {
	return s.Fields
}

func (s *TemplateView) SetFields(v []*TemplateViewFields) *TemplateView {
	s.Fields = v
	return s
}

func (s *TemplateView) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TemplateViewFields struct {
	// Whether to Display Original Image
	//
	// example:
	//
	// true
	DisplayOriImg *bool `json:"DisplayOriImg,omitempty" xml:"DisplayOriImg,omitempty"`
	// Associated Column Name
	//
	// example:
	//
	// url
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// View Type
	//
	// example:
	//
	// IMG
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Access Information
	VisitInfo *TemplateViewFieldsVisitInfo `json:"VisitInfo,omitempty" xml:"VisitInfo,omitempty" type:"Struct"`
}

func (s TemplateViewFields) String() string {
	return dara.Prettify(s)
}

func (s TemplateViewFields) GoString() string {
	return s.String()
}

func (s *TemplateViewFields) GetDisplayOriImg() *bool {
	return s.DisplayOriImg
}

func (s *TemplateViewFields) GetFieldName() *string {
	return s.FieldName
}

func (s *TemplateViewFields) GetType() *string {
	return s.Type
}

func (s *TemplateViewFields) GetVisitInfo() *TemplateViewFieldsVisitInfo {
	return s.VisitInfo
}

func (s *TemplateViewFields) SetDisplayOriImg(v bool) *TemplateViewFields {
	s.DisplayOriImg = &v
	return s
}

func (s *TemplateViewFields) SetFieldName(v string) *TemplateViewFields {
	s.FieldName = &v
	return s
}

func (s *TemplateViewFields) SetType(v string) *TemplateViewFields {
	s.Type = &v
	return s
}

func (s *TemplateViewFields) SetVisitInfo(v *TemplateViewFieldsVisitInfo) *TemplateViewFields {
	s.VisitInfo = v
	return s
}

func (s *TemplateViewFields) Validate() error {
	if s.VisitInfo != nil {
		if err := s.VisitInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TemplateViewFieldsVisitInfo struct {
	// Afts Configuration
	//
	// example:
	//
	// {"expiredTime":20}
	AftsConf map[string]interface{} `json:"AftsConf,omitempty" xml:"AftsConf,omitempty"`
	// OSS Configuration
	//
	// example:
	//
	// {"ossEndpoint":"***","ossAk":"***","ossAs":"***","ossOwner":"","ossBucket":""}
	OssConf map[string]interface{} `json:"OssConf,omitempty" xml:"OssConf,omitempty"`
}

func (s TemplateViewFieldsVisitInfo) String() string {
	return dara.Prettify(s)
}

func (s TemplateViewFieldsVisitInfo) GoString() string {
	return s.String()
}

func (s *TemplateViewFieldsVisitInfo) GetAftsConf() map[string]interface{} {
	return s.AftsConf
}

func (s *TemplateViewFieldsVisitInfo) GetOssConf() map[string]interface{} {
	return s.OssConf
}

func (s *TemplateViewFieldsVisitInfo) SetAftsConf(v map[string]interface{}) *TemplateViewFieldsVisitInfo {
	s.AftsConf = v
	return s
}

func (s *TemplateViewFieldsVisitInfo) SetOssConf(v map[string]interface{}) *TemplateViewFieldsVisitInfo {
	s.OssConf = v
	return s
}

func (s *TemplateViewFieldsVisitInfo) Validate() error {
	return dara.Validate(s)
}
