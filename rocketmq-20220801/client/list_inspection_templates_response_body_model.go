// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInspectionTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListInspectionTemplatesResponseBodyData) *ListInspectionTemplatesResponseBody
	GetData() []*ListInspectionTemplatesResponseBodyData
	SetMaxResults(v int32) *ListInspectionTemplatesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListInspectionTemplatesResponseBody
	GetNextToken() *string
	SetTotal(v int32) *ListInspectionTemplatesResponseBody
	GetTotal() *int32
}

type ListInspectionTemplatesResponseBody struct {
	// example:
	//
	// []
	Data []*ListInspectionTemplatesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 2
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListInspectionTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInspectionTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInspectionTemplatesResponseBody) GetData() []*ListInspectionTemplatesResponseBodyData {
	return s.Data
}

func (s *ListInspectionTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInspectionTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInspectionTemplatesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListInspectionTemplatesResponseBody) SetData(v []*ListInspectionTemplatesResponseBodyData) *ListInspectionTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListInspectionTemplatesResponseBody) SetMaxResults(v int32) *ListInspectionTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInspectionTemplatesResponseBody) SetNextToken(v string) *ListInspectionTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInspectionTemplatesResponseBody) SetTotal(v int32) *ListInspectionTemplatesResponseBody {
	s.Total = &v
	return s
}

func (s *ListInspectionTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInspectionTemplatesResponseBodyData struct {
	// example:
	//
	// {"timeWindow":{"defaultLookback":"1h","maxLookbackSeconds":3888000},"thresholds":[]}
	Configuration *ListInspectionTemplatesResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// example:
	//
	// 检查消息堆积和消费健康度
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 队列健康巡检
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// []
	Params interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// Instance
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// example:
	//
	// queue-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListInspectionTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInspectionTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInspectionTemplatesResponseBodyData) GetConfiguration() *ListInspectionTemplatesResponseBodyDataConfiguration {
	return s.Configuration
}

func (s *ListInspectionTemplatesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListInspectionTemplatesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListInspectionTemplatesResponseBodyData) GetParams() interface{} {
	return s.Params
}

func (s *ListInspectionTemplatesResponseBodyData) GetScopeType() *string {
	return s.ScopeType
}

func (s *ListInspectionTemplatesResponseBodyData) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListInspectionTemplatesResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListInspectionTemplatesResponseBodyData) SetConfiguration(v *ListInspectionTemplatesResponseBodyDataConfiguration) *ListInspectionTemplatesResponseBodyData {
	s.Configuration = v
	return s
}

func (s *ListInspectionTemplatesResponseBodyData) SetDescription(v string) *ListInspectionTemplatesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyData) SetName(v string) *ListInspectionTemplatesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyData) SetParams(v interface{}) *ListInspectionTemplatesResponseBodyData {
	s.Params = v
	return s
}

func (s *ListInspectionTemplatesResponseBodyData) SetScopeType(v string) *ListInspectionTemplatesResponseBodyData {
	s.ScopeType = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyData) SetTemplateCode(v string) *ListInspectionTemplatesResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyData) SetVersion(v string) *ListInspectionTemplatesResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyData) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInspectionTemplatesResponseBodyDataConfiguration struct {
	// example:
	//
	// []
	Parameters []*ListInspectionTemplatesResponseBodyDataConfigurationParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
}

func (s ListInspectionTemplatesResponseBodyDataConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListInspectionTemplatesResponseBodyDataConfiguration) GoString() string {
	return s.String()
}

func (s *ListInspectionTemplatesResponseBodyDataConfiguration) GetParameters() []*ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	return s.Parameters
}

func (s *ListInspectionTemplatesResponseBodyDataConfiguration) SetParameters(v []*ListInspectionTemplatesResponseBodyDataConfigurationParameters) *ListInspectionTemplatesResponseBodyDataConfiguration {
	s.Parameters = v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfiguration) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInspectionTemplatesResponseBodyDataConfigurationParameters struct {
	// example:
	//
	// "1h"
	Default interface{} `json:"default,omitempty" xml:"default,omitempty"`
	// example:
	//
	// ["all","active"]
	Enum []*string `json:"enum,omitempty" xml:"enum,omitempty" type:"Repeated"`
	// example:
	//
	// duration
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// lookback
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 3888000
	MaxLookbackSeconds *int64 `json:"maxLookbackSeconds,omitempty" xml:"maxLookbackSeconds,omitempty"`
	// example:
	//
	// 86400
	Maximum *float64 `json:"maximum,omitempty" xml:"maximum,omitempty"`
	// example:
	//
	// 1
	Minimum *float64 `json:"minimum,omitempty" xml:"minimum,omitempty"`
	// example:
	//
	// 回看时间
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// seconds
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// false
	WholeDays *bool `json:"wholeDays,omitempty" xml:"wholeDays,omitempty"`
}

func (s ListInspectionTemplatesResponseBodyDataConfigurationParameters) String() string {
	return dara.Prettify(s)
}

func (s ListInspectionTemplatesResponseBodyDataConfigurationParameters) GoString() string {
	return s.String()
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) GetDefault() interface{} {
	return s.Default
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) GetEnum() []*string {
	return s.Enum
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) GetFormat() *string {
	return s.Format
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) GetKey() *string {
	return s.Key
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) GetMaxLookbackSeconds() *int64 {
	return s.MaxLookbackSeconds
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) GetMaximum() *float64 {
	return s.Maximum
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) GetMinimum() *float64 {
	return s.Minimum
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) GetName() *string {
	return s.Name
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) GetRequired() *bool {
	return s.Required
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) GetType() *string {
	return s.Type
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) GetUnit() *string {
	return s.Unit
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) GetWholeDays() *bool {
	return s.WholeDays
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) SetDefault(v interface{}) *ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	s.Default = v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) SetEnum(v []*string) *ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	s.Enum = v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) SetFormat(v string) *ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	s.Format = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) SetKey(v string) *ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	s.Key = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) SetMaxLookbackSeconds(v int64) *ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	s.MaxLookbackSeconds = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) SetMaximum(v float64) *ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	s.Maximum = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) SetMinimum(v float64) *ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	s.Minimum = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) SetName(v string) *ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	s.Name = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) SetRequired(v bool) *ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	s.Required = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) SetType(v string) *ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	s.Type = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) SetUnit(v string) *ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	s.Unit = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) SetWholeDays(v bool) *ListInspectionTemplatesResponseBodyDataConfigurationParameters {
	s.WholeDays = &v
	return s
}

func (s *ListInspectionTemplatesResponseBodyDataConfigurationParameters) Validate() error {
	return dara.Validate(s)
}
