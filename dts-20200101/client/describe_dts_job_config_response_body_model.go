// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsJobConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameters(v []*DescribeDtsJobConfigResponseBodyParameters) *DescribeDtsJobConfigResponseBody
	GetParameters() []*DescribeDtsJobConfigResponseBodyParameters
	SetRequestId(v string) *DescribeDtsJobConfigResponseBody
	GetRequestId() *string
}

type DescribeDtsJobConfigResponseBody struct {
	Parameters []*DescribeDtsJobConfigResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDtsJobConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobConfigResponseBody) GetParameters() []*DescribeDtsJobConfigResponseBodyParameters {
	return s.Parameters
}

func (s *DescribeDtsJobConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDtsJobConfigResponseBody) SetParameters(v []*DescribeDtsJobConfigResponseBodyParameters) *DescribeDtsJobConfigResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeDtsJobConfigResponseBody) SetRequestId(v string) *DescribeDtsJobConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDtsJobConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobConfigResponseBodyParameters struct {
	CheckingCode   *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	DefaultValue   *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ForceRestart   *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	Modifiable     *string `json:"Modifiable,omitempty" xml:"Modifiable,omitempty"`
	Module         *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RecommendValue *string `json:"RecommendValue,omitempty" xml:"RecommendValue,omitempty"`
	RunningValue   *string `json:"RunningValue,omitempty" xml:"RunningValue,omitempty"`
	ValueType      *int32  `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s DescribeDtsJobConfigResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobConfigResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobConfigResponseBodyParameters) GetCheckingCode() *string {
	return s.CheckingCode
}

func (s *DescribeDtsJobConfigResponseBodyParameters) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeDtsJobConfigResponseBodyParameters) GetDescription() *string {
	return s.Description
}

func (s *DescribeDtsJobConfigResponseBodyParameters) GetForceRestart() *string {
	return s.ForceRestart
}

func (s *DescribeDtsJobConfigResponseBodyParameters) GetModifiable() *string {
	return s.Modifiable
}

func (s *DescribeDtsJobConfigResponseBodyParameters) GetModule() *string {
	return s.Module
}

func (s *DescribeDtsJobConfigResponseBodyParameters) GetName() *string {
	return s.Name
}

func (s *DescribeDtsJobConfigResponseBodyParameters) GetRecommendValue() *string {
	return s.RecommendValue
}

func (s *DescribeDtsJobConfigResponseBodyParameters) GetRunningValue() *string {
	return s.RunningValue
}

func (s *DescribeDtsJobConfigResponseBodyParameters) GetValueType() *int32 {
	return s.ValueType
}

func (s *DescribeDtsJobConfigResponseBodyParameters) SetCheckingCode(v string) *DescribeDtsJobConfigResponseBodyParameters {
	s.CheckingCode = &v
	return s
}

func (s *DescribeDtsJobConfigResponseBodyParameters) SetDefaultValue(v string) *DescribeDtsJobConfigResponseBodyParameters {
	s.DefaultValue = &v
	return s
}

func (s *DescribeDtsJobConfigResponseBodyParameters) SetDescription(v string) *DescribeDtsJobConfigResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *DescribeDtsJobConfigResponseBodyParameters) SetForceRestart(v string) *DescribeDtsJobConfigResponseBodyParameters {
	s.ForceRestart = &v
	return s
}

func (s *DescribeDtsJobConfigResponseBodyParameters) SetModifiable(v string) *DescribeDtsJobConfigResponseBodyParameters {
	s.Modifiable = &v
	return s
}

func (s *DescribeDtsJobConfigResponseBodyParameters) SetModule(v string) *DescribeDtsJobConfigResponseBodyParameters {
	s.Module = &v
	return s
}

func (s *DescribeDtsJobConfigResponseBodyParameters) SetName(v string) *DescribeDtsJobConfigResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *DescribeDtsJobConfigResponseBodyParameters) SetRecommendValue(v string) *DescribeDtsJobConfigResponseBodyParameters {
	s.RecommendValue = &v
	return s
}

func (s *DescribeDtsJobConfigResponseBodyParameters) SetRunningValue(v string) *DescribeDtsJobConfigResponseBodyParameters {
	s.RunningValue = &v
	return s
}

func (s *DescribeDtsJobConfigResponseBodyParameters) SetValueType(v int32) *DescribeDtsJobConfigResponseBodyParameters {
	s.ValueType = &v
	return s
}

func (s *DescribeDtsJobConfigResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}
