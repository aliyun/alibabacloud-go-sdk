// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFigureClusterConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAutoClustering(v bool) *FigureClusterConfig
	GetAutoClustering() *bool
	SetAutoGenerate(v bool) *FigureClusterConfig
	GetAutoGenerate() *bool
	SetEnabledFeatures(v []*string) *FigureClusterConfig
	GetEnabledFeatures() []*string
	SetMinEntityCount(v int64) *FigureClusterConfig
	GetMinEntityCount() *int64
}

type FigureClusterConfig struct {
	AutoClustering  *bool     `json:"AutoClustering,omitempty" xml:"AutoClustering,omitempty"`
	AutoGenerate    *bool     `json:"AutoGenerate,omitempty" xml:"AutoGenerate,omitempty"`
	EnabledFeatures []*string `json:"EnabledFeatures,omitempty" xml:"EnabledFeatures,omitempty" type:"Repeated"`
	MinEntityCount  *int64    `json:"MinEntityCount,omitempty" xml:"MinEntityCount,omitempty"`
}

func (s FigureClusterConfig) String() string {
	return dara.Prettify(s)
}

func (s FigureClusterConfig) GoString() string {
	return s.String()
}

func (s *FigureClusterConfig) GetAutoClustering() *bool {
	return s.AutoClustering
}

func (s *FigureClusterConfig) GetAutoGenerate() *bool {
	return s.AutoGenerate
}

func (s *FigureClusterConfig) GetEnabledFeatures() []*string {
	return s.EnabledFeatures
}

func (s *FigureClusterConfig) GetMinEntityCount() *int64 {
	return s.MinEntityCount
}

func (s *FigureClusterConfig) SetAutoClustering(v bool) *FigureClusterConfig {
	s.AutoClustering = &v
	return s
}

func (s *FigureClusterConfig) SetAutoGenerate(v bool) *FigureClusterConfig {
	s.AutoGenerate = &v
	return s
}

func (s *FigureClusterConfig) SetEnabledFeatures(v []*string) *FigureClusterConfig {
	s.EnabledFeatures = v
	return s
}

func (s *FigureClusterConfig) SetMinEntityCount(v int64) *FigureClusterConfig {
	s.MinEntityCount = &v
	return s
}

func (s *FigureClusterConfig) Validate() error {
	return dara.Validate(s)
}
