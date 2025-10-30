// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApprovalProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalType(v int32) *UpdateApprovalProcessRequest
	GetApprovalType() *int32
	SetDescription(v string) *UpdateApprovalProcessRequest
	GetDescription() *string
	SetEventLabel(v string) *UpdateApprovalProcessRequest
	GetEventLabel() *string
	SetExternalConfig(v string) *UpdateApprovalProcessRequest
	GetExternalConfig() *string
	SetMatchSchemaConfigs(v *UpdateApprovalProcessRequestMatchSchemaConfigs) *UpdateApprovalProcessRequest
	GetMatchSchemaConfigs() *UpdateApprovalProcessRequestMatchSchemaConfigs
	SetMatchSchemas(v *UpdateApprovalProcessRequestMatchSchemas) *UpdateApprovalProcessRequest
	GetMatchSchemas() *UpdateApprovalProcessRequestMatchSchemas
	SetProcessId(v string) *UpdateApprovalProcessRequest
	GetProcessId() *string
	SetProcessName(v string) *UpdateApprovalProcessRequest
	GetProcessName() *string
	SetProcessNodes(v [][]*string) *UpdateApprovalProcessRequest
	GetProcessNodes() [][]*string
}

type UpdateApprovalProcessRequest struct {
	ApprovalType       *int32                                          `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	Description        *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	EventLabel         *string                                         `json:"EventLabel,omitempty" xml:"EventLabel,omitempty"`
	ExternalConfig     *string                                         `json:"ExternalConfig,omitempty" xml:"ExternalConfig,omitempty"`
	MatchSchemaConfigs *UpdateApprovalProcessRequestMatchSchemaConfigs `json:"MatchSchemaConfigs,omitempty" xml:"MatchSchemaConfigs,omitempty" type:"Struct"`
	MatchSchemas       *UpdateApprovalProcessRequestMatchSchemas       `json:"MatchSchemas,omitempty" xml:"MatchSchemas,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// approval-process-f16bf74b2b29****
	ProcessId    *string     `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ProcessName  *string     `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessNodes [][]*string `json:"ProcessNodes,omitempty" xml:"ProcessNodes,omitempty" type:"Repeated"`
}

func (s UpdateApprovalProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequest) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequest) GetApprovalType() *int32 {
	return s.ApprovalType
}

func (s *UpdateApprovalProcessRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApprovalProcessRequest) GetEventLabel() *string {
	return s.EventLabel
}

func (s *UpdateApprovalProcessRequest) GetExternalConfig() *string {
	return s.ExternalConfig
}

func (s *UpdateApprovalProcessRequest) GetMatchSchemaConfigs() *UpdateApprovalProcessRequestMatchSchemaConfigs {
	return s.MatchSchemaConfigs
}

func (s *UpdateApprovalProcessRequest) GetMatchSchemas() *UpdateApprovalProcessRequestMatchSchemas {
	return s.MatchSchemas
}

func (s *UpdateApprovalProcessRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *UpdateApprovalProcessRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *UpdateApprovalProcessRequest) GetProcessNodes() [][]*string {
	return s.ProcessNodes
}

func (s *UpdateApprovalProcessRequest) SetApprovalType(v int32) *UpdateApprovalProcessRequest {
	s.ApprovalType = &v
	return s
}

func (s *UpdateApprovalProcessRequest) SetDescription(v string) *UpdateApprovalProcessRequest {
	s.Description = &v
	return s
}

func (s *UpdateApprovalProcessRequest) SetEventLabel(v string) *UpdateApprovalProcessRequest {
	s.EventLabel = &v
	return s
}

func (s *UpdateApprovalProcessRequest) SetExternalConfig(v string) *UpdateApprovalProcessRequest {
	s.ExternalConfig = &v
	return s
}

func (s *UpdateApprovalProcessRequest) SetMatchSchemaConfigs(v *UpdateApprovalProcessRequestMatchSchemaConfigs) *UpdateApprovalProcessRequest {
	s.MatchSchemaConfigs = v
	return s
}

func (s *UpdateApprovalProcessRequest) SetMatchSchemas(v *UpdateApprovalProcessRequestMatchSchemas) *UpdateApprovalProcessRequest {
	s.MatchSchemas = v
	return s
}

func (s *UpdateApprovalProcessRequest) SetProcessId(v string) *UpdateApprovalProcessRequest {
	s.ProcessId = &v
	return s
}

func (s *UpdateApprovalProcessRequest) SetProcessName(v string) *UpdateApprovalProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *UpdateApprovalProcessRequest) SetProcessNodes(v [][]*string) *UpdateApprovalProcessRequest {
	s.ProcessNodes = v
	return s
}

func (s *UpdateApprovalProcessRequest) Validate() error {
	if s.MatchSchemaConfigs != nil {
		if err := s.MatchSchemaConfigs.Validate(); err != nil {
			return err
		}
	}
	if s.MatchSchemas != nil {
		if err := s.MatchSchemas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApprovalProcessRequestMatchSchemaConfigs struct {
	AppUninstallSchemaConfig       *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig       `json:"AppUninstallSchemaConfig,omitempty" xml:"AppUninstallSchemaConfig,omitempty" type:"Struct"`
	DeviceRegistrationSchemaConfig *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig `json:"DeviceRegistrationSchemaConfig,omitempty" xml:"DeviceRegistrationSchemaConfig,omitempty" type:"Struct"`
	DlpSendSchemaConfig            *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig            `json:"DlpSendSchemaConfig,omitempty" xml:"DlpSendSchemaConfig,omitempty" type:"Struct"`
	DomainBlacklistSchemaConfig    *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig    `json:"DomainBlacklistSchemaConfig,omitempty" xml:"DomainBlacklistSchemaConfig,omitempty" type:"Struct"`
	DomainWhitelistSchemaConfig    *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig    `json:"DomainWhitelistSchemaConfig,omitempty" xml:"DomainWhitelistSchemaConfig,omitempty" type:"Struct"`
	EndpointHardeningSchemaConfig  *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig  `json:"EndpointHardeningSchemaConfig,omitempty" xml:"EndpointHardeningSchemaConfig,omitempty" type:"Struct"`
	PeripheralBlockSchemaConfig    *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig    `json:"PeripheralBlockSchemaConfig,omitempty" xml:"PeripheralBlockSchemaConfig,omitempty" type:"Struct"`
	SoftwareBlockSchemaConfig      *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig      `json:"SoftwareBlockSchemaConfig,omitempty" xml:"SoftwareBlockSchemaConfig,omitempty" type:"Struct"`
	SoftwareHardeningSchemaConfig  *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig  `json:"SoftwareHardeningSchemaConfig,omitempty" xml:"SoftwareHardeningSchemaConfig,omitempty" type:"Struct"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigs) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) GetAppUninstallSchemaConfig() *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig {
	return s.AppUninstallSchemaConfig
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) GetDeviceRegistrationSchemaConfig() *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig {
	return s.DeviceRegistrationSchemaConfig
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) GetDlpSendSchemaConfig() *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig {
	return s.DlpSendSchemaConfig
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) GetDomainBlacklistSchemaConfig() *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig {
	return s.DomainBlacklistSchemaConfig
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) GetDomainWhitelistSchemaConfig() *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig {
	return s.DomainWhitelistSchemaConfig
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) GetEndpointHardeningSchemaConfig() *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig {
	return s.EndpointHardeningSchemaConfig
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) GetPeripheralBlockSchemaConfig() *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig {
	return s.PeripheralBlockSchemaConfig
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) GetSoftwareBlockSchemaConfig() *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig {
	return s.SoftwareBlockSchemaConfig
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) GetSoftwareHardeningSchemaConfig() *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig {
	return s.SoftwareHardeningSchemaConfig
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) SetAppUninstallSchemaConfig(v *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig) *UpdateApprovalProcessRequestMatchSchemaConfigs {
	s.AppUninstallSchemaConfig = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) SetDeviceRegistrationSchemaConfig(v *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig) *UpdateApprovalProcessRequestMatchSchemaConfigs {
	s.DeviceRegistrationSchemaConfig = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) SetDlpSendSchemaConfig(v *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig) *UpdateApprovalProcessRequestMatchSchemaConfigs {
	s.DlpSendSchemaConfig = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) SetDomainBlacklistSchemaConfig(v *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig) *UpdateApprovalProcessRequestMatchSchemaConfigs {
	s.DomainBlacklistSchemaConfig = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) SetDomainWhitelistSchemaConfig(v *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig) *UpdateApprovalProcessRequestMatchSchemaConfigs {
	s.DomainWhitelistSchemaConfig = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) SetEndpointHardeningSchemaConfig(v *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig) *UpdateApprovalProcessRequestMatchSchemaConfigs {
	s.EndpointHardeningSchemaConfig = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) SetPeripheralBlockSchemaConfig(v *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig) *UpdateApprovalProcessRequestMatchSchemaConfigs {
	s.PeripheralBlockSchemaConfig = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) SetSoftwareBlockSchemaConfig(v *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig) *UpdateApprovalProcessRequestMatchSchemaConfigs {
	s.SoftwareBlockSchemaConfig = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) SetSoftwareHardeningSchemaConfig(v *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig) *UpdateApprovalProcessRequestMatchSchemaConfigs {
	s.SoftwareHardeningSchemaConfig = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigs) Validate() error {
	if s.AppUninstallSchemaConfig != nil {
		if err := s.AppUninstallSchemaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DeviceRegistrationSchemaConfig != nil {
		if err := s.DeviceRegistrationSchemaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DlpSendSchemaConfig != nil {
		if err := s.DlpSendSchemaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DomainBlacklistSchemaConfig != nil {
		if err := s.DomainBlacklistSchemaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DomainWhitelistSchemaConfig != nil {
		if err := s.DomainWhitelistSchemaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.EndpointHardeningSchemaConfig != nil {
		if err := s.EndpointHardeningSchemaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PeripheralBlockSchemaConfig != nil {
		if err := s.PeripheralBlockSchemaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SoftwareBlockSchemaConfig != nil {
		if err := s.SoftwareBlockSchemaConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SoftwareHardeningSchemaConfig != nil {
		if err := s.SoftwareHardeningSchemaConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig struct {
	ExternalProcessId *string                                                                           `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	SchemaId          *string                                                                           `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig) GetFieldMap() []*UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig) SetExternalProcessId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig) SetFieldMap(v []*UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap) *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig) SetSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfig) Validate() error {
	if s.FieldMap != nil {
		for _, item := range s.FieldMap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap) SetDisplayField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap) SetSystemField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsAppUninstallSchemaConfigFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig struct {
	ExternalProcessId *string                                                                                 `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	SchemaId          *string                                                                                 `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig) GetFieldMap() []*UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig) SetExternalProcessId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig) SetFieldMap(v []*UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap) *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig) SetSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfig) Validate() error {
	if s.FieldMap != nil {
		for _, item := range s.FieldMap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap) SetDisplayField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap) SetSystemField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDeviceRegistrationSchemaConfigFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig struct {
	ExternalProcessId *string                                                                      `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	SchemaId          *string                                                                      `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig) GetFieldMap() []*UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig) SetExternalProcessId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig) SetFieldMap(v []*UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap) *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig) SetSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfig) Validate() error {
	if s.FieldMap != nil {
		for _, item := range s.FieldMap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap) SetDisplayField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap) SetSystemField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDlpSendSchemaConfigFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig struct {
	ExternalProcessId *string                                                                              `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	SchemaId          *string                                                                              `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig) GetFieldMap() []*UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig) SetExternalProcessId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig) SetFieldMap(v []*UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap) *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig) SetSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfig) Validate() error {
	if s.FieldMap != nil {
		for _, item := range s.FieldMap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap) SetDisplayField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap) SetSystemField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainBlacklistSchemaConfigFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig struct {
	ExternalProcessId *string                                                                              `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	SchemaId          *string                                                                              `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig) GetFieldMap() []*UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig) SetExternalProcessId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig) SetFieldMap(v []*UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap) *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig) SetSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfig) Validate() error {
	if s.FieldMap != nil {
		for _, item := range s.FieldMap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap) SetDisplayField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap) SetSystemField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsDomainWhitelistSchemaConfigFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig struct {
	ExternalProcessId *string                                                                                `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	SchemaId          *string                                                                                `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig) GetFieldMap() []*UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig) SetExternalProcessId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig) SetFieldMap(v []*UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap) *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig) SetSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfig) Validate() error {
	if s.FieldMap != nil {
		for _, item := range s.FieldMap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap) SetDisplayField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap) SetSystemField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsEndpointHardeningSchemaConfigFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig struct {
	ExternalProcessId *string                                                                              `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	SchemaId          *string                                                                              `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig) GetFieldMap() []*UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig) SetExternalProcessId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig) SetFieldMap(v []*UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap) *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig) SetSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfig) Validate() error {
	if s.FieldMap != nil {
		for _, item := range s.FieldMap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap) SetDisplayField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap) SetSystemField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsPeripheralBlockSchemaConfigFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig struct {
	ExternalProcessId *string                                                                            `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	SchemaId          *string                                                                            `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig) GetFieldMap() []*UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig) SetExternalProcessId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig) SetFieldMap(v []*UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap) *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig) SetSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfig) Validate() error {
	if s.FieldMap != nil {
		for _, item := range s.FieldMap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap) SetDisplayField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap) SetSystemField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareBlockSchemaConfigFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig struct {
	ExternalProcessId *string                                                                                `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	SchemaId          *string                                                                                `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig) GetFieldMap() []*UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig) SetExternalProcessId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig) SetFieldMap(v []*UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap) *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig) SetSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfig) Validate() error {
	if s.FieldMap != nil {
		for _, item := range s.FieldMap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap) SetDisplayField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap) SetSystemField(v string) *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemaConfigsSoftwareHardeningSchemaConfigFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessRequestMatchSchemas struct {
	// example:
	//
	// approval-schema-090134f1ebff****
	AppUninstallSchemaId *string `json:"AppUninstallSchemaId,omitempty" xml:"AppUninstallSchemaId,omitempty"`
	// example:
	//
	// approval-schema-090134f1ebff****
	DeviceRegistrationSchemaId *string `json:"DeviceRegistrationSchemaId,omitempty" xml:"DeviceRegistrationSchemaId,omitempty"`
	// example:
	//
	// approval-schema-090134f1ebff****
	DlpSendSchemaId *string `json:"DlpSendSchemaId,omitempty" xml:"DlpSendSchemaId,omitempty"`
	// example:
	//
	// approval-schema-090134f1ebff****
	DomainBlacklistSchemaId *string `json:"DomainBlacklistSchemaId,omitempty" xml:"DomainBlacklistSchemaId,omitempty"`
	// example:
	//
	// approval-schema-090134f1ebff****
	DomainWhitelistSchemaId   *string `json:"DomainWhitelistSchemaId,omitempty" xml:"DomainWhitelistSchemaId,omitempty"`
	EndpointHardeningSchemaId *string `json:"EndpointHardeningSchemaId,omitempty" xml:"EndpointHardeningSchemaId,omitempty"`
	// example:
	//
	// approval-schema-090134f1ebff****
	PeripheralBlockSchemaId *string `json:"PeripheralBlockSchemaId,omitempty" xml:"PeripheralBlockSchemaId,omitempty"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SoftwareBlockSchemaId     *string `json:"SoftwareBlockSchemaId,omitempty" xml:"SoftwareBlockSchemaId,omitempty"`
	SoftwareHardeningSchemaId *string `json:"SoftwareHardeningSchemaId,omitempty" xml:"SoftwareHardeningSchemaId,omitempty"`
}

func (s UpdateApprovalProcessRequestMatchSchemas) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessRequestMatchSchemas) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessRequestMatchSchemas) GetAppUninstallSchemaId() *string {
	return s.AppUninstallSchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemas) GetDeviceRegistrationSchemaId() *string {
	return s.DeviceRegistrationSchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemas) GetDlpSendSchemaId() *string {
	return s.DlpSendSchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemas) GetDomainBlacklistSchemaId() *string {
	return s.DomainBlacklistSchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemas) GetDomainWhitelistSchemaId() *string {
	return s.DomainWhitelistSchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemas) GetEndpointHardeningSchemaId() *string {
	return s.EndpointHardeningSchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemas) GetPeripheralBlockSchemaId() *string {
	return s.PeripheralBlockSchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemas) GetSoftwareBlockSchemaId() *string {
	return s.SoftwareBlockSchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemas) GetSoftwareHardeningSchemaId() *string {
	return s.SoftwareHardeningSchemaId
}

func (s *UpdateApprovalProcessRequestMatchSchemas) SetAppUninstallSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemas {
	s.AppUninstallSchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemas) SetDeviceRegistrationSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemas {
	s.DeviceRegistrationSchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemas) SetDlpSendSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemas {
	s.DlpSendSchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemas) SetDomainBlacklistSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemas {
	s.DomainBlacklistSchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemas) SetDomainWhitelistSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemas {
	s.DomainWhitelistSchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemas) SetEndpointHardeningSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemas {
	s.EndpointHardeningSchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemas) SetPeripheralBlockSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemas {
	s.PeripheralBlockSchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemas) SetSoftwareBlockSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemas {
	s.SoftwareBlockSchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemas) SetSoftwareHardeningSchemaId(v string) *UpdateApprovalProcessRequestMatchSchemas {
	s.SoftwareHardeningSchemaId = &v
	return s
}

func (s *UpdateApprovalProcessRequestMatchSchemas) Validate() error {
	return dara.Validate(s)
}
