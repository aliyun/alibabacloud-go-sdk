// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListServiceConfigsResponseBody
	GetCode() *int32
	SetData(v []*ListServiceConfigsResponseBodyData) *ListServiceConfigsResponseBody
	GetData() []*ListServiceConfigsResponseBodyData
	SetMsg(v string) *ListServiceConfigsResponseBody
	GetMsg() *string
	SetRequestId(v string) *ListServiceConfigsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListServiceConfigsResponseBody
	GetSuccess() *bool
}

type ListServiceConfigsResponseBody struct {
	// example:
	//
	// 400
	Code *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListServiceConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListServiceConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListServiceConfigsResponseBody) GetData() []*ListServiceConfigsResponseBodyData {
	return s.Data
}

func (s *ListServiceConfigsResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ListServiceConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListServiceConfigsResponseBody) SetCode(v int32) *ListServiceConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListServiceConfigsResponseBody) SetData(v []*ListServiceConfigsResponseBodyData) *ListServiceConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListServiceConfigsResponseBody) SetMsg(v string) *ListServiceConfigsResponseBody {
	s.Msg = &v
	return s
}

func (s *ListServiceConfigsResponseBody) SetRequestId(v string) *ListServiceConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceConfigsResponseBody) SetSuccess(v bool) *ListServiceConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListServiceConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListServiceConfigsResponseBodyData struct {
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// example:
	//
	// nickname_detection
	CopyFrom          *string                                              `json:"CopyFrom,omitempty" xml:"CopyFrom,omitempty"`
	CustomServiceConf *ListServiceConfigsResponseBodyDataCustomServiceConf `json:"CustomServiceConf,omitempty" xml:"CustomServiceConf,omitempty" type:"Struct"`
	// example:
	//
	// 2023-07-11 15:40:04
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// {}
	Option map[string]interface{} `json:"Option,omitempty" xml:"Option,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// nickname_detection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceDesc *string `json:"ServiceDesc,omitempty" xml:"ServiceDesc,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// plus
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// UID。
	//
	// example:
	//
	// 1674*****0071291
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UseStatus *string `json:"UseStatus,omitempty" xml:"UseStatus,omitempty"`
}

func (s ListServiceConfigsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListServiceConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponseBodyData) GetClassify() *string {
	return s.Classify
}

func (s *ListServiceConfigsResponseBodyData) GetCopyFrom() *string {
	return s.CopyFrom
}

func (s *ListServiceConfigsResponseBodyData) GetCustomServiceConf() *ListServiceConfigsResponseBodyDataCustomServiceConf {
	return s.CustomServiceConf
}

func (s *ListServiceConfigsResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceConfigsResponseBodyData) GetOption() map[string]interface{} {
	return s.Option
}

func (s *ListServiceConfigsResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListServiceConfigsResponseBodyData) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *ListServiceConfigsResponseBodyData) GetServiceDesc() *string {
	return s.ServiceDesc
}

func (s *ListServiceConfigsResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListServiceConfigsResponseBodyData) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServiceConfigsResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *ListServiceConfigsResponseBodyData) GetUseStatus() *string {
	return s.UseStatus
}

func (s *ListServiceConfigsResponseBodyData) SetClassify(v string) *ListServiceConfigsResponseBodyData {
	s.Classify = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetCopyFrom(v string) *ListServiceConfigsResponseBodyData {
	s.CopyFrom = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetCustomServiceConf(v *ListServiceConfigsResponseBodyDataCustomServiceConf) *ListServiceConfigsResponseBodyData {
	s.CustomServiceConf = v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetGmtModified(v string) *ListServiceConfigsResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetOption(v map[string]interface{}) *ListServiceConfigsResponseBodyData {
	s.Option = v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetResourceType(v string) *ListServiceConfigsResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetServiceCode(v string) *ListServiceConfigsResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetServiceDesc(v string) *ListServiceConfigsResponseBodyData {
	s.ServiceDesc = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetServiceName(v string) *ListServiceConfigsResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetServiceType(v string) *ListServiceConfigsResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetUid(v string) *ListServiceConfigsResponseBodyData {
	s.Uid = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) SetUseStatus(v string) *ListServiceConfigsResponseBodyData {
	s.UseStatus = &v
	return s
}

func (s *ListServiceConfigsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListServiceConfigsResponseBodyDataCustomServiceConf struct {
	// example:
	//
	// audio_media_detection
	AudioService       *string                                                     `json:"AudioService,omitempty" xml:"AudioService,omitempty"`
	ImageService       []*string                                                   `json:"ImageService,omitempty" xml:"ImageService,omitempty" type:"Repeated"`
	KeywordFilterLibs  []*string                                                   `json:"KeywordFilterLibs,omitempty" xml:"KeywordFilterLibs,omitempty" type:"Repeated"`
	KeywordHitLibs     []*string                                                   `json:"KeywordHitLibs,omitempty" xml:"KeywordHitLibs,omitempty" type:"Repeated"`
	Rules              []*ListServiceConfigsResponseBodyDataCustomServiceConfRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	SimilarTextHitLibs []*string                                                   `json:"SimilarTextHitLibs,omitempty" xml:"SimilarTextHitLibs,omitempty" type:"Repeated"`
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConf) String() string {
	return dara.Prettify(s)
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConf) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) GetAudioService() *string {
	return s.AudioService
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) GetImageService() []*string {
	return s.ImageService
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) GetKeywordFilterLibs() []*string {
	return s.KeywordFilterLibs
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) GetKeywordHitLibs() []*string {
	return s.KeywordHitLibs
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) GetRules() []*ListServiceConfigsResponseBodyDataCustomServiceConfRules {
	return s.Rules
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) GetSimilarTextHitLibs() []*string {
	return s.SimilarTextHitLibs
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) SetAudioService(v string) *ListServiceConfigsResponseBodyDataCustomServiceConf {
	s.AudioService = &v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) SetImageService(v []*string) *ListServiceConfigsResponseBodyDataCustomServiceConf {
	s.ImageService = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) SetKeywordFilterLibs(v []*string) *ListServiceConfigsResponseBodyDataCustomServiceConf {
	s.KeywordFilterLibs = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) SetKeywordHitLibs(v []*string) *ListServiceConfigsResponseBodyDataCustomServiceConf {
	s.KeywordHitLibs = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) SetRules(v []*ListServiceConfigsResponseBodyDataCustomServiceConfRules) *ListServiceConfigsResponseBodyDataCustomServiceConf {
	s.Rules = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) SetSimilarTextHitLibs(v []*string) *ListServiceConfigsResponseBodyDataCustomServiceConf {
	s.SimilarTextHitLibs = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConf) Validate() error {
	return dara.Validate(s)
}

type ListServiceConfigsResponseBodyDataCustomServiceConfRules struct {
	ImageScanRule *ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule `json:"ImageScanRule,omitempty" xml:"ImageScanRule,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Index        *int32                                                                `json:"Index,omitempty" xml:"Index,omitempty"`
	TextScanRule *ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule `json:"TextScanRule,omitempty" xml:"TextScanRule,omitempty" type:"Struct"`
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConfRules) String() string {
	return dara.Prettify(s)
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConfRules) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRules) GetImageScanRule() *ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule {
	return s.ImageScanRule
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRules) GetIndex() *int32 {
	return s.Index
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRules) GetTextScanRule() *ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule {
	return s.TextScanRule
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRules) SetImageScanRule(v *ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule) *ListServiceConfigsResponseBodyDataCustomServiceConfRules {
	s.ImageScanRule = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRules) SetIndex(v int32) *ListServiceConfigsResponseBodyDataCustomServiceConfRules {
	s.Index = &v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRules) SetTextScanRule(v *ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule) *ListServiceConfigsResponseBodyDataCustomServiceConfRules {
	s.TextScanRule = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRules) Validate() error {
	return dara.Validate(s)
}

type ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule struct {
	Services []*string `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule) String() string {
	return dara.Prettify(s)
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule) GetServices() []*string {
	return s.Services
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule) SetServices(v []*string) *ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule {
	s.Services = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRulesImageScanRule) Validate() error {
	return dara.Validate(s)
}

type ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule struct {
	Services []*string `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule) String() string {
	return dara.Prettify(s)
}

func (s ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule) GoString() string {
	return s.String()
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule) GetServices() []*string {
	return s.Services
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule) SetServices(v []*string) *ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule {
	s.Services = v
	return s
}

func (s *ListServiceConfigsResponseBodyDataCustomServiceConfRulesTextScanRule) Validate() error {
	return dara.Validate(s)
}
