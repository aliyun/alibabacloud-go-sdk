// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveDelayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDelayConfigList(v *ListLiveDelayConfigResponseBodyDelayConfigList) *ListLiveDelayConfigResponseBody
	GetDelayConfigList() *ListLiveDelayConfigResponseBodyDelayConfigList
	SetRequestId(v string) *ListLiveDelayConfigResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListLiveDelayConfigResponseBody
	GetTotal() *int32
}

type ListLiveDelayConfigResponseBody struct {
	DelayConfigList *ListLiveDelayConfigResponseBodyDelayConfigList `json:"DelayConfigList,omitempty" xml:"DelayConfigList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A3136B58-5876-4168-83CA-B562781981A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of stream delay configurations.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListLiveDelayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveDelayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveDelayConfigResponseBody) GetDelayConfigList() *ListLiveDelayConfigResponseBodyDelayConfigList {
	return s.DelayConfigList
}

func (s *ListLiveDelayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveDelayConfigResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListLiveDelayConfigResponseBody) SetDelayConfigList(v *ListLiveDelayConfigResponseBodyDelayConfigList) *ListLiveDelayConfigResponseBody {
	s.DelayConfigList = v
	return s
}

func (s *ListLiveDelayConfigResponseBody) SetRequestId(v string) *ListLiveDelayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveDelayConfigResponseBody) SetTotal(v int32) *ListLiveDelayConfigResponseBody {
	s.Total = &v
	return s
}

func (s *ListLiveDelayConfigResponseBody) Validate() error {
	if s.DelayConfigList != nil {
		if err := s.DelayConfigList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLiveDelayConfigResponseBodyDelayConfigList struct {
	DelayConfig []*ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig `json:"DelayConfig,omitempty" xml:"DelayConfig,omitempty" type:"Repeated"`
}

func (s ListLiveDelayConfigResponseBodyDelayConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveDelayConfigResponseBodyDelayConfigList) GoString() string {
	return s.String()
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigList) GetDelayConfig() []*ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig {
	return s.DelayConfig
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigList) SetDelayConfig(v []*ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) *ListLiveDelayConfigResponseBodyDelayConfigList {
	s.DelayConfig = v
	return s
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigList) Validate() error {
	if s.DelayConfig != nil {
		for _, item := range s.DelayConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig struct {
	App             *string `json:"App,omitempty" xml:"App,omitempty"`
	DelayTime       *string `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Stream          *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
	TaskTriggerMode *string `json:"TaskTriggerMode,omitempty" xml:"TaskTriggerMode,omitempty"`
}

func (s ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) String() string {
	return dara.Prettify(s)
}

func (s ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) GoString() string {
	return s.String()
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) GetApp() *string {
	return s.App
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) GetDelayTime() *string {
	return s.DelayTime
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) GetDomain() *string {
	return s.Domain
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) GetStream() *string {
	return s.Stream
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) GetTaskTriggerMode() *string {
	return s.TaskTriggerMode
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) SetApp(v string) *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig {
	s.App = &v
	return s
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) SetDelayTime(v string) *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig {
	s.DelayTime = &v
	return s
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) SetDomain(v string) *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig {
	s.Domain = &v
	return s
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) SetStream(v string) *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig {
	s.Stream = &v
	return s
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) SetTaskTriggerMode(v string) *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig {
	s.TaskTriggerMode = &v
	return s
}

func (s *ListLiveDelayConfigResponseBodyDelayConfigListDelayConfig) Validate() error {
	return dara.Validate(s)
}
