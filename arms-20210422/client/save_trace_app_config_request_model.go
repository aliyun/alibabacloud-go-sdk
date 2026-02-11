// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTraceAppConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPid(v string) *SaveTraceAppConfigRequest
	GetPid() *string
	SetSettings(v []*SaveTraceAppConfigRequestSettings) *SaveTraceAppConfigRequest
	GetSettings() []*SaveTraceAppConfigRequestSettings
}

type SaveTraceAppConfigRequest struct {
	// This parameter is required.
	Pid      *string                              `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Settings []*SaveTraceAppConfigRequestSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Repeated"`
}

func (s SaveTraceAppConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveTraceAppConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigRequest) GetPid() *string {
	return s.Pid
}

func (s *SaveTraceAppConfigRequest) GetSettings() []*SaveTraceAppConfigRequestSettings {
	return s.Settings
}

func (s *SaveTraceAppConfigRequest) SetPid(v string) *SaveTraceAppConfigRequest {
	s.Pid = &v
	return s
}

func (s *SaveTraceAppConfigRequest) SetSettings(v []*SaveTraceAppConfigRequestSettings) *SaveTraceAppConfigRequest {
	s.Settings = v
	return s
}

func (s *SaveTraceAppConfigRequest) Validate() error {
	if s.Settings != nil {
		for _, item := range s.Settings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveTraceAppConfigRequestSettings struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SaveTraceAppConfigRequestSettings) String() string {
	return dara.Prettify(s)
}

func (s SaveTraceAppConfigRequestSettings) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigRequestSettings) GetKey() *string {
	return s.Key
}

func (s *SaveTraceAppConfigRequestSettings) GetValue() *string {
	return s.Value
}

func (s *SaveTraceAppConfigRequestSettings) SetKey(v string) *SaveTraceAppConfigRequestSettings {
	s.Key = &v
	return s
}

func (s *SaveTraceAppConfigRequestSettings) SetValue(v string) *SaveTraceAppConfigRequestSettings {
	s.Value = &v
	return s
}

func (s *SaveTraceAppConfigRequestSettings) Validate() error {
	return dara.Validate(s)
}
