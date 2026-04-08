// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoggingConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetDestinations(v []*LogDestination) *LoggingConfiguration
	GetDestinations() []*LogDestination
}

type LoggingConfiguration struct {
	// 日志输出的目标配置列表
	Destinations []*LogDestination `json:"destinations" xml:"destinations" type:"Repeated"`
}

func (s LoggingConfiguration) String() string {
	return dara.Prettify(s)
}

func (s LoggingConfiguration) GoString() string {
	return s.String()
}

func (s *LoggingConfiguration) GetDestinations() []*LogDestination {
	return s.Destinations
}

func (s *LoggingConfiguration) SetDestinations(v []*LogDestination) *LoggingConfiguration {
	s.Destinations = v
	return s
}

func (s *LoggingConfiguration) Validate() error {
	if s.Destinations != nil {
		for _, item := range s.Destinations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
