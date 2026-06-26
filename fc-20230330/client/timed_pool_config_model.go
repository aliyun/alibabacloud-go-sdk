// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTimedPoolConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveEndDate(v string) *TimedPoolConfig
	GetEffectiveEndDate() *string
	SetEffectiveStartDate(v string) *TimedPoolConfig
	GetEffectiveStartDate() *string
	SetElasticIntervals(v []*ElasticInterval) *TimedPoolConfig
	GetElasticIntervals() []*ElasticInterval
	SetTimeZone(v string) *TimedPoolConfig
	GetTimeZone() *string
}

type TimedPoolConfig struct {
	EffectiveEndDate   *string            `json:"effectiveEndDate,omitempty" xml:"effectiveEndDate,omitempty"`
	EffectiveStartDate *string            `json:"effectiveStartDate,omitempty" xml:"effectiveStartDate,omitempty"`
	ElasticIntervals   []*ElasticInterval `json:"elasticIntervals" xml:"elasticIntervals" type:"Repeated"`
	TimeZone           *string            `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s TimedPoolConfig) String() string {
	return dara.Prettify(s)
}

func (s TimedPoolConfig) GoString() string {
	return s.String()
}

func (s *TimedPoolConfig) GetEffectiveEndDate() *string {
	return s.EffectiveEndDate
}

func (s *TimedPoolConfig) GetEffectiveStartDate() *string {
	return s.EffectiveStartDate
}

func (s *TimedPoolConfig) GetElasticIntervals() []*ElasticInterval {
	return s.ElasticIntervals
}

func (s *TimedPoolConfig) GetTimeZone() *string {
	return s.TimeZone
}

func (s *TimedPoolConfig) SetEffectiveEndDate(v string) *TimedPoolConfig {
	s.EffectiveEndDate = &v
	return s
}

func (s *TimedPoolConfig) SetEffectiveStartDate(v string) *TimedPoolConfig {
	s.EffectiveStartDate = &v
	return s
}

func (s *TimedPoolConfig) SetElasticIntervals(v []*ElasticInterval) *TimedPoolConfig {
	s.ElasticIntervals = v
	return s
}

func (s *TimedPoolConfig) SetTimeZone(v string) *TimedPoolConfig {
	s.TimeZone = &v
	return s
}

func (s *TimedPoolConfig) Validate() error {
	if s.ElasticIntervals != nil {
		for _, item := range s.ElasticIntervals {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
