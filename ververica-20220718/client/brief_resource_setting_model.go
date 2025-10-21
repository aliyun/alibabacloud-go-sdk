// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBriefResourceSetting interface {
	dara.Model
	String() string
	GoString() string
	SetBatchResourceSetting(v *BatchResourceSetting) *BriefResourceSetting
	GetBatchResourceSetting() *BatchResourceSetting
	SetFlinkConf(v map[string]interface{}) *BriefResourceSetting
	GetFlinkConf() map[string]interface{}
	SetStreamingResourceSetting(v *StreamingResourceSetting) *BriefResourceSetting
	GetStreamingResourceSetting() *StreamingResourceSetting
}

type BriefResourceSetting struct {
	BatchResourceSetting     *BatchResourceSetting     `json:"batchResourceSetting,omitempty" xml:"batchResourceSetting,omitempty"`
	FlinkConf                map[string]interface{}    `json:"flinkConf,omitempty" xml:"flinkConf,omitempty"`
	StreamingResourceSetting *StreamingResourceSetting `json:"streamingResourceSetting,omitempty" xml:"streamingResourceSetting,omitempty"`
}

func (s BriefResourceSetting) String() string {
	return dara.Prettify(s)
}

func (s BriefResourceSetting) GoString() string {
	return s.String()
}

func (s *BriefResourceSetting) GetBatchResourceSetting() *BatchResourceSetting {
	return s.BatchResourceSetting
}

func (s *BriefResourceSetting) GetFlinkConf() map[string]interface{} {
	return s.FlinkConf
}

func (s *BriefResourceSetting) GetStreamingResourceSetting() *StreamingResourceSetting {
	return s.StreamingResourceSetting
}

func (s *BriefResourceSetting) SetBatchResourceSetting(v *BatchResourceSetting) *BriefResourceSetting {
	s.BatchResourceSetting = v
	return s
}

func (s *BriefResourceSetting) SetFlinkConf(v map[string]interface{}) *BriefResourceSetting {
	s.FlinkConf = v
	return s
}

func (s *BriefResourceSetting) SetStreamingResourceSetting(v *StreamingResourceSetting) *BriefResourceSetting {
	s.StreamingResourceSetting = v
	return s
}

func (s *BriefResourceSetting) Validate() error {
	if s.BatchResourceSetting != nil {
		if err := s.BatchResourceSetting.Validate(); err != nil {
			return err
		}
	}
	if s.StreamingResourceSetting != nil {
		if err := s.StreamingResourceSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}
