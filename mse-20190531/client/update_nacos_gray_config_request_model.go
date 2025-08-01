// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosGrayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateNacosGrayConfigRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *UpdateNacosGrayConfigRequest
	GetAppName() *string
	SetContent(v string) *UpdateNacosGrayConfigRequest
	GetContent() *string
	SetDataId(v string) *UpdateNacosGrayConfigRequest
	GetDataId() *string
	SetGrayRule(v string) *UpdateNacosGrayConfigRequest
	GetGrayRule() *string
	SetGrayRuleName(v string) *UpdateNacosGrayConfigRequest
	GetGrayRuleName() *string
	SetGrayRulePriority(v int32) *UpdateNacosGrayConfigRequest
	GetGrayRulePriority() *int32
	SetGrayType(v string) *UpdateNacosGrayConfigRequest
	GetGrayType() *string
	SetGroup(v string) *UpdateNacosGrayConfigRequest
	GetGroup() *string
	SetInstanceId(v string) *UpdateNacosGrayConfigRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *UpdateNacosGrayConfigRequest
	GetNamespaceId() *string
	SetOpType(v string) *UpdateNacosGrayConfigRequest
	GetOpType() *string
	SetRegionId(v string) *UpdateNacosGrayConfigRequest
	GetRegionId() *string
	SetRequestPars(v string) *UpdateNacosGrayConfigRequest
	GetRequestPars() *string
	SetStopGray(v bool) *UpdateNacosGrayConfigRequest
	GetStopGray() *bool
}

type UpdateNacosGrayConfigRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// asdf
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// key=value1,value2
	GrayRule         *string `json:"GrayRule,omitempty" xml:"GrayRule,omitempty"`
	GrayRuleName     *string `json:"GrayRuleName,omitempty" xml:"GrayRuleName,omitempty"`
	GrayRulePriority *int32  `json:"GrayRulePriority,omitempty" xml:"GrayRulePriority,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Tags
	GrayType *string `json:"GrayType,omitempty" xml:"GrayType,omitempty"`
	// example:
	//
	// DEFAULT_GROUP
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-st2212****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 6cf708a5-****-89f2-3ba62c5ee9ba
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	OpType      *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// example:
	//
	// true
	StopGray *bool `json:"StopGray,omitempty" xml:"StopGray,omitempty"`
}

func (s UpdateNacosGrayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosGrayConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateNacosGrayConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateNacosGrayConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateNacosGrayConfigRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateNacosGrayConfigRequest) GetDataId() *string {
	return s.DataId
}

func (s *UpdateNacosGrayConfigRequest) GetGrayRule() *string {
	return s.GrayRule
}

func (s *UpdateNacosGrayConfigRequest) GetGrayRuleName() *string {
	return s.GrayRuleName
}

func (s *UpdateNacosGrayConfigRequest) GetGrayRulePriority() *int32 {
	return s.GrayRulePriority
}

func (s *UpdateNacosGrayConfigRequest) GetGrayType() *string {
	return s.GrayType
}

func (s *UpdateNacosGrayConfigRequest) GetGroup() *string {
	return s.Group
}

func (s *UpdateNacosGrayConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateNacosGrayConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateNacosGrayConfigRequest) GetOpType() *string {
	return s.OpType
}

func (s *UpdateNacosGrayConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateNacosGrayConfigRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *UpdateNacosGrayConfigRequest) GetStopGray() *bool {
	return s.StopGray
}

func (s *UpdateNacosGrayConfigRequest) SetAcceptLanguage(v string) *UpdateNacosGrayConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetAppName(v string) *UpdateNacosGrayConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetContent(v string) *UpdateNacosGrayConfigRequest {
	s.Content = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetDataId(v string) *UpdateNacosGrayConfigRequest {
	s.DataId = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetGrayRule(v string) *UpdateNacosGrayConfigRequest {
	s.GrayRule = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetGrayRuleName(v string) *UpdateNacosGrayConfigRequest {
	s.GrayRuleName = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetGrayRulePriority(v int32) *UpdateNacosGrayConfigRequest {
	s.GrayRulePriority = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetGrayType(v string) *UpdateNacosGrayConfigRequest {
	s.GrayType = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetGroup(v string) *UpdateNacosGrayConfigRequest {
	s.Group = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetInstanceId(v string) *UpdateNacosGrayConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetNamespaceId(v string) *UpdateNacosGrayConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetOpType(v string) *UpdateNacosGrayConfigRequest {
	s.OpType = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetRegionId(v string) *UpdateNacosGrayConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetRequestPars(v string) *UpdateNacosGrayConfigRequest {
	s.RequestPars = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) SetStopGray(v bool) *UpdateNacosGrayConfigRequest {
	s.StopGray = &v
	return s
}

func (s *UpdateNacosGrayConfigRequest) Validate() error {
	return dara.Validate(s)
}
