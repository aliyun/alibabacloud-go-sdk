// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTemplateCriterionIssuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetServiceTemplateCriterionIssuesResponseBody
	GetRequestId() *string
	SetRiskyTemplateCount(v int32) *GetServiceTemplateCriterionIssuesResponseBody
	GetRiskyTemplateCount() *int32
	SetTemplateCriterionIssueList(v []*GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) *GetServiceTemplateCriterionIssuesResponseBody
	GetTemplateCriterionIssueList() []*GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList
	SetTotalCriterionIssueCount(v int32) *GetServiceTemplateCriterionIssuesResponseBody
	GetTotalCriterionIssueCount() *int32
	SetTotalMandatoryCriterionIssueCount(v int32) *GetServiceTemplateCriterionIssuesResponseBody
	GetTotalMandatoryCriterionIssueCount() *int32
}

type GetServiceTemplateCriterionIssuesResponseBody struct {
	// example:
	//
	// A361BA9E-xxxx-xxxx-xxxx-C26E5180456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	RiskyTemplateCount         *int32                                                                     `json:"RiskyTemplateCount,omitempty" xml:"RiskyTemplateCount,omitempty"`
	TemplateCriterionIssueList []*GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList `json:"TemplateCriterionIssueList,omitempty" xml:"TemplateCriterionIssueList,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	TotalCriterionIssueCount *int32 `json:"TotalCriterionIssueCount,omitempty" xml:"TotalCriterionIssueCount,omitempty"`
	// example:
	//
	// 1
	TotalMandatoryCriterionIssueCount *int32 `json:"TotalMandatoryCriterionIssueCount,omitempty" xml:"TotalMandatoryCriterionIssueCount,omitempty"`
}

func (s GetServiceTemplateCriterionIssuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateCriterionIssuesResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateCriterionIssuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceTemplateCriterionIssuesResponseBody) GetRiskyTemplateCount() *int32 {
	return s.RiskyTemplateCount
}

func (s *GetServiceTemplateCriterionIssuesResponseBody) GetTemplateCriterionIssueList() []*GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList {
	return s.TemplateCriterionIssueList
}

func (s *GetServiceTemplateCriterionIssuesResponseBody) GetTotalCriterionIssueCount() *int32 {
	return s.TotalCriterionIssueCount
}

func (s *GetServiceTemplateCriterionIssuesResponseBody) GetTotalMandatoryCriterionIssueCount() *int32 {
	return s.TotalMandatoryCriterionIssueCount
}

func (s *GetServiceTemplateCriterionIssuesResponseBody) SetRequestId(v string) *GetServiceTemplateCriterionIssuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBody) SetRiskyTemplateCount(v int32) *GetServiceTemplateCriterionIssuesResponseBody {
	s.RiskyTemplateCount = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBody) SetTemplateCriterionIssueList(v []*GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) *GetServiceTemplateCriterionIssuesResponseBody {
	s.TemplateCriterionIssueList = v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBody) SetTotalCriterionIssueCount(v int32) *GetServiceTemplateCriterionIssuesResponseBody {
	s.TotalCriterionIssueCount = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBody) SetTotalMandatoryCriterionIssueCount(v int32) *GetServiceTemplateCriterionIssuesResponseBody {
	s.TotalMandatoryCriterionIssueCount = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList struct {
	CriterionIssues []*GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues `json:"CriterionIssues,omitempty" xml:"CriterionIssues,omitempty" type:"Repeated"`
	TemplateName    *string                                                                                   `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// http://service-private-info/xxx/ros/template/tpl-xxxx.json
	TemplateUrl *int32 `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
	// example:
	//
	// 3
	TotalCriterionIssueCount *int32 `json:"TotalCriterionIssueCount,omitempty" xml:"TotalCriterionIssueCount,omitempty"`
	// example:
	//
	// 1
	TotalMandatoryCriterionIssueCount *int32 `json:"TotalMandatoryCriterionIssueCount,omitempty" xml:"TotalMandatoryCriterionIssueCount,omitempty"`
}

func (s GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) GetCriterionIssues() []*GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues {
	return s.CriterionIssues
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) GetTemplateUrl() *int32 {
	return s.TemplateUrl
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) GetTotalCriterionIssueCount() *int32 {
	return s.TotalCriterionIssueCount
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) GetTotalMandatoryCriterionIssueCount() *int32 {
	return s.TotalMandatoryCriterionIssueCount
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) SetCriterionIssues(v []*GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues) *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList {
	s.CriterionIssues = v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) SetTemplateName(v string) *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList {
	s.TemplateName = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) SetTemplateUrl(v int32) *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList {
	s.TemplateUrl = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) SetTotalCriterionIssueCount(v int32) *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList {
	s.TotalCriterionIssueCount = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) SetTotalMandatoryCriterionIssueCount(v int32) *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList {
	s.TotalMandatoryCriterionIssueCount = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueList) Validate() error {
	return dara.Validate(s)
}

type GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues struct {
	ExtendInfo *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty" type:"Struct"`
	// example:
	//
	// Mandatory
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// $.Parameters.PayType
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// example:
	//
	// ParameterNeedAssociationProperty
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues) GetExtendInfo() *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo {
	return s.ExtendInfo
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues) GetLevel() *string {
	return s.Level
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues) GetPosition() *string {
	return s.Position
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues) GetType() *string {
	return s.Type
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues) SetExtendInfo(v *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo) *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues {
	s.ExtendInfo = v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues) SetLevel(v string) *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues {
	s.Level = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues) SetPosition(v string) *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues {
	s.Position = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues) SetType(v string) *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues {
	s.Type = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssues) Validate() error {
	return dara.Validate(s)
}

type GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo struct {
	// example:
	//
	// ChargeType
	AssociationProperty *string `json:"AssociationProperty,omitempty" xml:"AssociationProperty,omitempty"`
	// example:
	//
	// null
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// example:
	//
	// null
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
}

func (s GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo) GetAssociationProperty() *string {
	return s.AssociationProperty
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo) GetProperty() *string {
	return s.Property
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo) SetAssociationProperty(v string) *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo {
	s.AssociationProperty = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo) SetProperty(v string) *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo {
	s.Property = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo) SetPropertyValue(v string) *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo {
	s.PropertyValue = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponseBodyTemplateCriterionIssueListCriterionIssuesExtendInfo) Validate() error {
	return dara.Validate(s)
}
