// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *SourceConfig
	GetLogstore() *string
	SetStartTime(v int64) *SourceConfig
	GetStartTime() *int64
}

type SourceConfig struct {
	// example:
	//
	// my-sls-logstore-name
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// 1704790317
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s SourceConfig) String() string {
	return dara.Prettify(s)
}

func (s SourceConfig) GoString() string {
	return s.String()
}

func (s *SourceConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *SourceConfig) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SourceConfig) SetLogstore(v string) *SourceConfig {
	s.Logstore = &v
	return s
}

func (s *SourceConfig) SetStartTime(v int64) *SourceConfig {
	s.StartTime = &v
	return s
}

func (s *SourceConfig) Validate() error {
	return dara.Validate(s)
}
