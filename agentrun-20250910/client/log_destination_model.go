// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogDestination interface {
	dara.Model
	String() string
	GoString() string
	SetSlsLogDestination(v *SLSLogDestination) *LogDestination
	GetSlsLogDestination() *SLSLogDestination
}

type LogDestination struct {
	// 阿里云日志服务（SLS）的日志目标配置
	SlsLogDestination *SLSLogDestination `json:"slsLogDestination,omitempty" xml:"slsLogDestination,omitempty"`
}

func (s LogDestination) String() string {
	return dara.Prettify(s)
}

func (s LogDestination) GoString() string {
	return s.String()
}

func (s *LogDestination) GetSlsLogDestination() *SLSLogDestination {
	return s.SlsLogDestination
}

func (s *LogDestination) SetSlsLogDestination(v *SLSLogDestination) *LogDestination {
	s.SlsLogDestination = v
	return s
}

func (s *LogDestination) Validate() error {
	if s.SlsLogDestination != nil {
		if err := s.SlsLogDestination.Validate(); err != nil {
			return err
		}
	}
	return nil
}
