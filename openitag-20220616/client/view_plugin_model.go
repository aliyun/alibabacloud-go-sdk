// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewPlugin interface {
	dara.Model
	String() string
	GoString() string
	SetBindField(v string) *ViewPlugin
	GetBindField() *string
	SetConvertor(v string) *ViewPlugin
	GetConvertor() *string
	SetCorsProxy(v bool) *ViewPlugin
	GetCorsProxy() *bool
	SetDisplayOriImg(v bool) *ViewPlugin
	GetDisplayOriImg() *bool
	SetExif(v map[string]interface{}) *ViewPlugin
	GetExif() map[string]interface{}
	SetHide(v bool) *ViewPlugin
	GetHide() *bool
	SetPlugins(v map[string]interface{}) *ViewPlugin
	GetPlugins() map[string]interface{}
	SetRelationQuestionIds(v []*string) *ViewPlugin
	GetRelationQuestionIds() []*string
	SetType(v string) *ViewPlugin
	GetType() *string
	SetVisitInfo(v *ViewPluginVisitInfo) *ViewPlugin
	GetVisitInfo() *ViewPluginVisitInfo
}

type ViewPlugin struct {
	// Field mapping to a field in the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// url
	BindField *string `json:"BindField,omitempty" xml:"BindField,omitempty"`
	// Array transformation UDF.
	//
	// example:
	//
	// null
	Convertor *string `json:"Convertor,omitempty" xml:"Convertor,omitempty"`
	// Whether to handle cross-domain requests.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	CorsProxy *bool `json:"CorsProxy,omitempty" xml:"CorsProxy,omitempty"`
	// Whether to display the original image.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DisplayOriImg *bool `json:"DisplayOriImg,omitempty" xml:"DisplayOriImg,omitempty"`
	// Extra information.
	//
	// example:
	//
	// {}
	Exif map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// Whether to hide.
	//
	// example:
	//
	// true
	Hide *bool `json:"Hide,omitempty" xml:"Hide,omitempty"`
	// Nested widgets.
	//
	// example:
	//
	// []
	Plugins map[string]interface{} `json:"Plugins,omitempty" xml:"Plugins,omitempty"`
	// List of associated questions.
	RelationQuestionIds []*string `json:"RelationQuestionIds,omitempty" xml:"RelationQuestionIds,omitempty" type:"Repeated"`
	// Widget type. Valid values:
	//
	// - Image: Image.
	//
	// - Text: Text.
	//
	// - Video: Video.
	//
	// - Audio: Audio.
	//
	// This parameter is required.
	//
	// example:
	//
	// Text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Access information.
	VisitInfo *ViewPluginVisitInfo `json:"VisitInfo,omitempty" xml:"VisitInfo,omitempty" type:"Struct"`
}

func (s ViewPlugin) String() string {
	return dara.Prettify(s)
}

func (s ViewPlugin) GoString() string {
	return s.String()
}

func (s *ViewPlugin) GetBindField() *string {
	return s.BindField
}

func (s *ViewPlugin) GetConvertor() *string {
	return s.Convertor
}

func (s *ViewPlugin) GetCorsProxy() *bool {
	return s.CorsProxy
}

func (s *ViewPlugin) GetDisplayOriImg() *bool {
	return s.DisplayOriImg
}

func (s *ViewPlugin) GetExif() map[string]interface{} {
	return s.Exif
}

func (s *ViewPlugin) GetHide() *bool {
	return s.Hide
}

func (s *ViewPlugin) GetPlugins() map[string]interface{} {
	return s.Plugins
}

func (s *ViewPlugin) GetRelationQuestionIds() []*string {
	return s.RelationQuestionIds
}

func (s *ViewPlugin) GetType() *string {
	return s.Type
}

func (s *ViewPlugin) GetVisitInfo() *ViewPluginVisitInfo {
	return s.VisitInfo
}

func (s *ViewPlugin) SetBindField(v string) *ViewPlugin {
	s.BindField = &v
	return s
}

func (s *ViewPlugin) SetConvertor(v string) *ViewPlugin {
	s.Convertor = &v
	return s
}

func (s *ViewPlugin) SetCorsProxy(v bool) *ViewPlugin {
	s.CorsProxy = &v
	return s
}

func (s *ViewPlugin) SetDisplayOriImg(v bool) *ViewPlugin {
	s.DisplayOriImg = &v
	return s
}

func (s *ViewPlugin) SetExif(v map[string]interface{}) *ViewPlugin {
	s.Exif = v
	return s
}

func (s *ViewPlugin) SetHide(v bool) *ViewPlugin {
	s.Hide = &v
	return s
}

func (s *ViewPlugin) SetPlugins(v map[string]interface{}) *ViewPlugin {
	s.Plugins = v
	return s
}

func (s *ViewPlugin) SetRelationQuestionIds(v []*string) *ViewPlugin {
	s.RelationQuestionIds = v
	return s
}

func (s *ViewPlugin) SetType(v string) *ViewPlugin {
	s.Type = &v
	return s
}

func (s *ViewPlugin) SetVisitInfo(v *ViewPluginVisitInfo) *ViewPlugin {
	s.VisitInfo = v
	return s
}

func (s *ViewPlugin) Validate() error {
	if s.VisitInfo != nil {
		if err := s.VisitInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ViewPluginVisitInfo struct {
	// AFTS configuration.
	//
	// example:
	//
	// {"expiredTime":1}
	AftsConf map[string]interface{} `json:"aftsConf,omitempty" xml:"aftsConf,omitempty"`
	// OSS configuration.
	//
	// example:
	//
	// {"ossEndpoint":"","ossAk":"","ossAs":"","ossOwner":"","ossBucket":"","folder":"","expiredTime":""}
	OssConf map[string]interface{} `json:"ossConf,omitempty" xml:"ossConf,omitempty"`
}

func (s ViewPluginVisitInfo) String() string {
	return dara.Prettify(s)
}

func (s ViewPluginVisitInfo) GoString() string {
	return s.String()
}

func (s *ViewPluginVisitInfo) GetAftsConf() map[string]interface{} {
	return s.AftsConf
}

func (s *ViewPluginVisitInfo) GetOssConf() map[string]interface{} {
	return s.OssConf
}

func (s *ViewPluginVisitInfo) SetAftsConf(v map[string]interface{}) *ViewPluginVisitInfo {
	s.AftsConf = v
	return s
}

func (s *ViewPluginVisitInfo) SetOssConf(v map[string]interface{}) *ViewPluginVisitInfo {
	s.OssConf = v
	return s
}

func (s *ViewPluginVisitInfo) Validate() error {
	return dara.Validate(s)
}
