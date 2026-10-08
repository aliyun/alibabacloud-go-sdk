// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBuildPackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildPackList(v *ListBuildPackResponseBodyBuildPackList) *ListBuildPackResponseBody
	GetBuildPackList() *ListBuildPackResponseBodyBuildPackList
	SetCode(v int32) *ListBuildPackResponseBody
	GetCode() *int32
	SetMessage(v string) *ListBuildPackResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBuildPackResponseBody
	GetRequestId() *string
}

type ListBuildPackResponseBody struct {
	BuildPackList *ListBuildPackResponseBodyBuildPackList `json:"BuildPackList,omitempty" xml:"BuildPackList,omitempty" type:"Struct"`
	// code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4FD4-*************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBuildPackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBuildPackResponseBody) GoString() string {
	return s.String()
}

func (s *ListBuildPackResponseBody) GetBuildPackList() *ListBuildPackResponseBodyBuildPackList {
	return s.BuildPackList
}

func (s *ListBuildPackResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListBuildPackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBuildPackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBuildPackResponseBody) SetBuildPackList(v *ListBuildPackResponseBodyBuildPackList) *ListBuildPackResponseBody {
	s.BuildPackList = v
	return s
}

func (s *ListBuildPackResponseBody) SetCode(v int32) *ListBuildPackResponseBody {
	s.Code = &v
	return s
}

func (s *ListBuildPackResponseBody) SetMessage(v string) *ListBuildPackResponseBody {
	s.Message = &v
	return s
}

func (s *ListBuildPackResponseBody) SetRequestId(v string) *ListBuildPackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBuildPackResponseBody) Validate() error {
	if s.BuildPackList != nil {
		if err := s.BuildPackList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBuildPackResponseBodyBuildPackList struct {
	BuildPack []*ListBuildPackResponseBodyBuildPackListBuildPack `json:"BuildPack,omitempty" xml:"BuildPack,omitempty" type:"Repeated"`
}

func (s ListBuildPackResponseBodyBuildPackList) String() string {
	return dara.Prettify(s)
}

func (s ListBuildPackResponseBodyBuildPackList) GoString() string {
	return s.String()
}

func (s *ListBuildPackResponseBodyBuildPackList) GetBuildPack() []*ListBuildPackResponseBodyBuildPackListBuildPack {
	return s.BuildPack
}

func (s *ListBuildPackResponseBodyBuildPackList) SetBuildPack(v []*ListBuildPackResponseBodyBuildPackListBuildPack) *ListBuildPackResponseBodyBuildPackList {
	s.BuildPack = v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackList) Validate() error {
	if s.BuildPack != nil {
		for _, item := range s.BuildPack {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBuildPackResponseBodyBuildPackListBuildPack struct {
	ConfigId           *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Disabled           *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	Feature            *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	ImageId            *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	MultipleTenant     *bool   `json:"MultipleTenant,omitempty" xml:"MultipleTenant,omitempty"`
	PackVersion        *string `json:"PackVersion,omitempty" xml:"PackVersion,omitempty"`
	PandoraDesc        *string `json:"PandoraDesc,omitempty" xml:"PandoraDesc,omitempty"`
	PandoraDownloadUrl *string `json:"PandoraDownloadUrl,omitempty" xml:"PandoraDownloadUrl,omitempty"`
	PandoraVersion     *string `json:"PandoraVersion,omitempty" xml:"PandoraVersion,omitempty"`
	PluginInfo         *string `json:"PluginInfo,omitempty" xml:"PluginInfo,omitempty"`
	ScriptName         *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	ScriptVersion      *string `json:"ScriptVersion,omitempty" xml:"ScriptVersion,omitempty"`
	SupportFeatures    *string `json:"SupportFeatures,omitempty" xml:"SupportFeatures,omitempty"`
	TengineDownloadUrl *string `json:"TengineDownloadUrl,omitempty" xml:"TengineDownloadUrl,omitempty"`
	TengineImageId     *string `json:"TengineImageId,omitempty" xml:"TengineImageId,omitempty"`
	TomcatDesc         *string `json:"TomcatDesc,omitempty" xml:"TomcatDesc,omitempty"`
	TomcatDownloadUrl  *string `json:"TomcatDownloadUrl,omitempty" xml:"TomcatDownloadUrl,omitempty"`
	TomcatPath         *string `json:"TomcatPath,omitempty" xml:"TomcatPath,omitempty"`
	TomcatVersion      *string `json:"TomcatVersion,omitempty" xml:"TomcatVersion,omitempty"`
	WithTengine        *bool   `json:"WithTengine,omitempty" xml:"WithTengine,omitempty"`
}

func (s ListBuildPackResponseBodyBuildPackListBuildPack) String() string {
	return dara.Prettify(s)
}

func (s ListBuildPackResponseBodyBuildPackListBuildPack) GoString() string {
	return s.String()
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetDisabled() *bool {
	return s.Disabled
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetFeature() *string {
	return s.Feature
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetImageId() *string {
	return s.ImageId
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetMultipleTenant() *bool {
	return s.MultipleTenant
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetPackVersion() *string {
	return s.PackVersion
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetPandoraDesc() *string {
	return s.PandoraDesc
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetPandoraDownloadUrl() *string {
	return s.PandoraDownloadUrl
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetPandoraVersion() *string {
	return s.PandoraVersion
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetPluginInfo() *string {
	return s.PluginInfo
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetScriptName() *string {
	return s.ScriptName
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetScriptVersion() *string {
	return s.ScriptVersion
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetSupportFeatures() *string {
	return s.SupportFeatures
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetTengineDownloadUrl() *string {
	return s.TengineDownloadUrl
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetTengineImageId() *string {
	return s.TengineImageId
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetTomcatDesc() *string {
	return s.TomcatDesc
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetTomcatDownloadUrl() *string {
	return s.TomcatDownloadUrl
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetTomcatPath() *string {
	return s.TomcatPath
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetTomcatVersion() *string {
	return s.TomcatVersion
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) GetWithTengine() *bool {
	return s.WithTengine
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetConfigId(v int64) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.ConfigId = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetDisabled(v bool) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.Disabled = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetFeature(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.Feature = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetImageId(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.ImageId = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetMultipleTenant(v bool) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.MultipleTenant = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetPackVersion(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.PackVersion = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetPandoraDesc(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.PandoraDesc = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetPandoraDownloadUrl(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.PandoraDownloadUrl = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetPandoraVersion(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.PandoraVersion = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetPluginInfo(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.PluginInfo = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetScriptName(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.ScriptName = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetScriptVersion(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.ScriptVersion = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetSupportFeatures(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.SupportFeatures = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetTengineDownloadUrl(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.TengineDownloadUrl = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetTengineImageId(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.TengineImageId = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetTomcatDesc(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.TomcatDesc = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetTomcatDownloadUrl(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.TomcatDownloadUrl = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetTomcatPath(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.TomcatPath = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetTomcatVersion(v string) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.TomcatVersion = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) SetWithTengine(v bool) *ListBuildPackResponseBodyBuildPackListBuildPack {
	s.WithTengine = &v
	return s
}

func (s *ListBuildPackResponseBodyBuildPackListBuildPack) Validate() error {
	return dara.Validate(s)
}
