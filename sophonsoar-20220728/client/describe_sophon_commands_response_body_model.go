// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSophonCommandsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeSophonCommandsResponseBodyData) *DescribeSophonCommandsResponseBody
	GetData() []*DescribeSophonCommandsResponseBodyData
	SetRequestId(v string) *DescribeSophonCommandsResponseBody
	GetRequestId() *string
}

type DescribeSophonCommandsResponseBody struct {
	// The commands.
	Data []*DescribeSophonCommandsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1E1EC464-3BD7-518F-9937-BCC12E6855FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSophonCommandsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSophonCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsResponseBody) GetData() []*DescribeSophonCommandsResponseBodyData {
	return s.Data
}

func (s *DescribeSophonCommandsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSophonCommandsResponseBody) SetData(v []*DescribeSophonCommandsResponseBodyData) *DescribeSophonCommandsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSophonCommandsResponseBody) SetRequestId(v string) *DescribeSophonCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSophonCommandsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSophonCommandsResponseBodyData struct {
	// The description of the command.
	//
	// example:
	//
	// This is a action of processing for WAF
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the command.
	//
	// example:
	//
	// WAF Process IP
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The name of the command.
	//
	// example:
	//
	// waf_process_ip_v2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter configurations.
	ParamConfig []*DescribeSophonCommandsResponseBodyDataParamConfig `json:"ParamConfig,omitempty" xml:"ParamConfig,omitempty" type:"Repeated"`
}

func (s DescribeSophonCommandsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSophonCommandsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeSophonCommandsResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeSophonCommandsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeSophonCommandsResponseBodyData) GetParamConfig() []*DescribeSophonCommandsResponseBodyDataParamConfig {
	return s.ParamConfig
}

func (s *DescribeSophonCommandsResponseBodyData) SetDescription(v string) *DescribeSophonCommandsResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyData) SetDisplayName(v string) *DescribeSophonCommandsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyData) SetName(v string) *DescribeSophonCommandsResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyData) SetParamConfig(v []*DescribeSophonCommandsResponseBodyDataParamConfig) *DescribeSophonCommandsResponseBodyData {
	s.ParamConfig = v
	return s
}

func (s *DescribeSophonCommandsResponseBodyData) Validate() error {
	if s.ParamConfig != nil {
		for _, item := range s.ParamConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSophonCommandsResponseBodyDataParamConfig struct {
	// The regular expression that is used to check the format of the parameter value. If the parameter is left empty, the check is not performed.
	//
	// example:
	//
	// [0-9]{4}\\.[0-9]{4}\\.[0-9]{4}\\.[0-9]{4}
	CheckField *string `json:"CheckField,omitempty" xml:"CheckField,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// ip
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// Indicates whether the parameter is required. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Necessary *bool `json:"Necessary,omitempty" xml:"Necessary,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 12.xx.xx.xx
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSophonCommandsResponseBodyDataParamConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeSophonCommandsResponseBodyDataParamConfig) GoString() string {
	return s.String()
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) GetCheckField() *string {
	return s.CheckField
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) GetField() *string {
	return s.Field
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) GetNecessary() *bool {
	return s.Necessary
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) GetValue() *string {
	return s.Value
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) SetCheckField(v string) *DescribeSophonCommandsResponseBodyDataParamConfig {
	s.CheckField = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) SetField(v string) *DescribeSophonCommandsResponseBodyDataParamConfig {
	s.Field = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) SetNecessary(v bool) *DescribeSophonCommandsResponseBodyDataParamConfig {
	s.Necessary = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) SetValue(v string) *DescribeSophonCommandsResponseBodyDataParamConfig {
	s.Value = &v
	return s
}

func (s *DescribeSophonCommandsResponseBodyDataParamConfig) Validate() error {
	return dara.Validate(s)
}
