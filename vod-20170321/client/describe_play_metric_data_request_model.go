// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribePlayMetricDataRequest
	GetAppName() *string
	SetBeginTs(v string) *DescribePlayMetricDataRequest
	GetBeginTs() *string
	SetDefinition(v string) *DescribePlayMetricDataRequest
	GetDefinition() *string
	SetEndTs(v string) *DescribePlayMetricDataRequest
	GetEndTs() *string
	SetExperienceLevel(v string) *DescribePlayMetricDataRequest
	GetExperienceLevel() *string
	SetItemConfigs(v string) *DescribePlayMetricDataRequest
	GetItemConfigs() *string
	SetMetricType(v string) *DescribePlayMetricDataRequest
	GetMetricType() *string
	SetNetwork(v string) *DescribePlayMetricDataRequest
	GetNetwork() *string
	SetOs(v string) *DescribePlayMetricDataRequest
	GetOs() *string
	SetSdkVersion(v string) *DescribePlayMetricDataRequest
	GetSdkVersion() *string
	SetTerminalType(v string) *DescribePlayMetricDataRequest
	GetTerminalType() *string
}

type DescribePlayMetricDataRequest struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	BeginTs    *string `json:"BeginTs,omitempty" xml:"BeginTs,omitempty"`
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// This parameter is required.
	EndTs           *string `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	ExperienceLevel *string `json:"ExperienceLevel,omitempty" xml:"ExperienceLevel,omitempty"`
	ItemConfigs     *string `json:"ItemConfigs,omitempty" xml:"ItemConfigs,omitempty"`
	// This parameter is required.
	MetricType   *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Network      *string `json:"Network,omitempty" xml:"Network,omitempty"`
	Os           *string `json:"Os,omitempty" xml:"Os,omitempty"`
	SdkVersion   *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
}

func (s DescribePlayMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayMetricDataRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribePlayMetricDataRequest) GetBeginTs() *string {
	return s.BeginTs
}

func (s *DescribePlayMetricDataRequest) GetDefinition() *string {
	return s.Definition
}

func (s *DescribePlayMetricDataRequest) GetEndTs() *string {
	return s.EndTs
}

func (s *DescribePlayMetricDataRequest) GetExperienceLevel() *string {
	return s.ExperienceLevel
}

func (s *DescribePlayMetricDataRequest) GetItemConfigs() *string {
	return s.ItemConfigs
}

func (s *DescribePlayMetricDataRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribePlayMetricDataRequest) GetNetwork() *string {
	return s.Network
}

func (s *DescribePlayMetricDataRequest) GetOs() *string {
	return s.Os
}

func (s *DescribePlayMetricDataRequest) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *DescribePlayMetricDataRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribePlayMetricDataRequest) SetAppName(v string) *DescribePlayMetricDataRequest {
	s.AppName = &v
	return s
}

func (s *DescribePlayMetricDataRequest) SetBeginTs(v string) *DescribePlayMetricDataRequest {
	s.BeginTs = &v
	return s
}

func (s *DescribePlayMetricDataRequest) SetDefinition(v string) *DescribePlayMetricDataRequest {
	s.Definition = &v
	return s
}

func (s *DescribePlayMetricDataRequest) SetEndTs(v string) *DescribePlayMetricDataRequest {
	s.EndTs = &v
	return s
}

func (s *DescribePlayMetricDataRequest) SetExperienceLevel(v string) *DescribePlayMetricDataRequest {
	s.ExperienceLevel = &v
	return s
}

func (s *DescribePlayMetricDataRequest) SetItemConfigs(v string) *DescribePlayMetricDataRequest {
	s.ItemConfigs = &v
	return s
}

func (s *DescribePlayMetricDataRequest) SetMetricType(v string) *DescribePlayMetricDataRequest {
	s.MetricType = &v
	return s
}

func (s *DescribePlayMetricDataRequest) SetNetwork(v string) *DescribePlayMetricDataRequest {
	s.Network = &v
	return s
}

func (s *DescribePlayMetricDataRequest) SetOs(v string) *DescribePlayMetricDataRequest {
	s.Os = &v
	return s
}

func (s *DescribePlayMetricDataRequest) SetSdkVersion(v string) *DescribePlayMetricDataRequest {
	s.SdkVersion = &v
	return s
}

func (s *DescribePlayMetricDataRequest) SetTerminalType(v string) *DescribePlayMetricDataRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribePlayMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
