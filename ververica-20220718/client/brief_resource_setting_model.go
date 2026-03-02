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
	// The resource configuration for the deployment in batch mode. This parameter is required for a deployment in batch mode.
	BatchResourceSetting *BatchResourceSetting `json:"batchResourceSetting,omitempty" xml:"batchResourceSetting,omitempty"`
	// The Realtime Compute for Apache Flink configuration.
	//
	// example:
	//
	// “execution.checkpointing.interval: 180s”
	FlinkConf map[string]interface{} `json:"flinkConf,omitempty" xml:"flinkConf,omitempty"`
	// The resource configuration for the deployment in streaming mode. This parameter is required for a deployment in streaming mode.
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
