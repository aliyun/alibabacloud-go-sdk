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
	// The process ID (PID) of the application.
	//
	// Log on to the ARMS console. In the left-side navigation pane, choose **Application Monitoring*	- > **Application List**. On the Application List page, click the name of an application. The URL in the address bar contains the PID of the application. The PID is indicated in the pid=xxx format. The PID is usually percent encoded as xxx%40xxx. You must modify this value to remove the percent encoding. For example, if the PID in the URL is xxx%4074xxx, you must replace %40 with an at sign (@) to obtain xxx@74xxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2n80plglh@745eddxxx
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The settings of Application Monitoring.
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
	// The key of the settings that you want to modify. For more information about the supported settings, see the following sections.
	//
	// 	- Trace sampling settings
	//
	// 	- Agent switch settings
	//
	// 	- Threshold settings
	//
	// 	- Advanced settings
	//
	// 	- Thread settings
	//
	// 	- Memory snapshot settings
	//
	// 	- URL convergence settings
	//
	// 	- Business log association settings
	//
	// example:
	//
	// sampling.enable
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the settings that you want to modify. For more information about the supported settings, see the following sections.
	//
	// 	- Trace sampling settings
	//
	// 	- Agent switch settings
	//
	// 	- Threshold settings
	//
	// 	- Advanced settings
	//
	// 	- Thread settings
	//
	// 	- Memory snapshot settings
	//
	// 	- URL convergence settings
	//
	// 	- Business log association settings
	//
	// example:
	//
	// true
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
