// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteTrafficSequenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSiteTrafficSequenceResponseBody
	GetRequestId() *string
	SetTrafficSequences(v []*GetSiteTrafficSequenceResponseBodyTrafficSequences) *GetSiteTrafficSequenceResponseBody
	GetTrafficSequences() []*GetSiteTrafficSequenceResponseBodyTrafficSequences
}

type GetSiteTrafficSequenceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The traffic sequences of the site.
	TrafficSequences []*GetSiteTrafficSequenceResponseBodyTrafficSequences `json:"TrafficSequences,omitempty" xml:"TrafficSequences,omitempty" type:"Repeated"`
}

func (s GetSiteTrafficSequenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSiteTrafficSequenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteTrafficSequenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSiteTrafficSequenceResponseBody) GetTrafficSequences() []*GetSiteTrafficSequenceResponseBodyTrafficSequences {
	return s.TrafficSequences
}

func (s *GetSiteTrafficSequenceResponseBody) SetRequestId(v string) *GetSiteTrafficSequenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteTrafficSequenceResponseBody) SetTrafficSequences(v []*GetSiteTrafficSequenceResponseBodyTrafficSequences) *GetSiteTrafficSequenceResponseBody {
	s.TrafficSequences = v
	return s
}

func (s *GetSiteTrafficSequenceResponseBody) Validate() error {
	if s.TrafficSequences != nil {
		for _, item := range s.TrafficSequences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSiteTrafficSequenceResponseBodyTrafficSequences struct {
	// The list of site features associated with the traffic sequence.
	FunctionList []*GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList `json:"FunctionList,omitempty" xml:"FunctionList,omitempty" type:"Repeated"`
	// The order of the current sequence in the entire traffic sequence.
	//
	// example:
	//
	// 1
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The traffic sequence routing.
	//
	// example:
	//
	// /safe/ddos/basic
	Router *string `json:"Router,omitempty" xml:"Router,omitempty"`
	// The sequence code.
	//
	// example:
	//
	// ddos
	SequenceCode *string `json:"SequenceCode,omitempty" xml:"SequenceCode,omitempty"`
	// The sequence name.
	//
	// example:
	//
	// DDOS
	SequenceName *string `json:"SequenceName,omitempty" xml:"SequenceName,omitempty"`
}

func (s GetSiteTrafficSequenceResponseBodyTrafficSequences) String() string {
	return dara.Prettify(s)
}

func (s GetSiteTrafficSequenceResponseBodyTrafficSequences) GoString() string {
	return s.String()
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequences) GetFunctionList() []*GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList {
	return s.FunctionList
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequences) GetOrder() *string {
	return s.Order
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequences) GetRouter() *string {
	return s.Router
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequences) GetSequenceCode() *string {
	return s.SequenceCode
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequences) GetSequenceName() *string {
	return s.SequenceName
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequences) SetFunctionList(v []*GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList) *GetSiteTrafficSequenceResponseBodyTrafficSequences {
	s.FunctionList = v
	return s
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequences) SetOrder(v string) *GetSiteTrafficSequenceResponseBodyTrafficSequences {
	s.Order = &v
	return s
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequences) SetRouter(v string) *GetSiteTrafficSequenceResponseBodyTrafficSequences {
	s.Router = &v
	return s
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequences) SetSequenceCode(v string) *GetSiteTrafficSequenceResponseBodyTrafficSequences {
	s.SequenceCode = &v
	return s
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequences) SetSequenceName(v string) *GetSiteTrafficSequenceResponseBodyTrafficSequences {
	s.SequenceName = &v
	return s
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequences) Validate() error {
	if s.FunctionList != nil {
		for _, item := range s.FunctionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList struct {
	// The list of configurations for the site feature associated with the traffic sequence.
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	Configs []*GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The feature name.
	//
	// example:
	//
	// redirect_rules
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// Indicates whether the site has a corresponding configuration. Valid values:
	//
	// - true: The site has a corresponding configuration.
	//
	// - false: The site does not have a corresponding configuration.
	//
	// example:
	//
	// true
	HasConfig *bool `json:"HasConfig,omitempty" xml:"HasConfig,omitempty"`
}

func (s GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList) String() string {
	return dara.Prettify(s)
}

func (s GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList) GoString() string {
	return s.String()
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList) GetConfigs() []*GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs {
	return s.Configs
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList) GetHasConfig() *bool {
	return s.HasConfig
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList) SetConfigs(v []*GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs) *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList {
	s.Configs = v
	return s
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList) SetFunctionName(v string) *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList {
	s.FunctionName = &v
	return s
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList) SetHasConfig(v bool) *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList {
	s.HasConfig = &v
	return s
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionList) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs struct {
	// The configuration ID.
	//
	// example:
	//
	// 480339095269376
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. Valid values:
	//
	// - global: global configuration.
	//
	// - rule: rule configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
}

func (s GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs) GoString() string {
	return s.String()
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs) SetConfigId(v int64) *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs {
	s.ConfigId = &v
	return s
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs) SetConfigType(v string) *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs {
	s.ConfigType = &v
	return s
}

func (s *GetSiteTrafficSequenceResponseBodyTrafficSequencesFunctionListConfigs) Validate() error {
	return dara.Validate(s)
}
