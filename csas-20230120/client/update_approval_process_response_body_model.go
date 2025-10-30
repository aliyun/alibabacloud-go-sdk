// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApprovalProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProcess(v *UpdateApprovalProcessResponseBodyProcess) *UpdateApprovalProcessResponseBody
	GetProcess() *UpdateApprovalProcessResponseBodyProcess
	SetRequestId(v string) *UpdateApprovalProcessResponseBody
	GetRequestId() *string
}

type UpdateApprovalProcessResponseBody struct {
	Process *UpdateApprovalProcessResponseBodyProcess `json:"Process,omitempty" xml:"Process,omitempty" type:"Struct"`
	// example:
	//
	// 7E9D7ACD-53D5-56EF-A913-79D148D06299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApprovalProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBody) GetProcess() *UpdateApprovalProcessResponseBodyProcess {
	return s.Process
}

func (s *UpdateApprovalProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApprovalProcessResponseBody) SetProcess(v *UpdateApprovalProcessResponseBodyProcess) *UpdateApprovalProcessResponseBody {
	s.Process = v
	return s
}

func (s *UpdateApprovalProcessResponseBody) SetRequestId(v string) *UpdateApprovalProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBody) Validate() error {
	if s.Process != nil {
		if err := s.Process.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApprovalProcessResponseBodyProcess struct {
	AppUninstallPolicies *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies `json:"AppUninstallPolicies,omitempty" xml:"AppUninstallPolicies,omitempty" type:"Struct"`
	ApprovalType         *int32                                                        `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// example:
	//
	// 2022-07-11 15:31:39
	CreateTime                 *string                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description                *string                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceRegistrationPolicies *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies `json:"DeviceRegistrationPolicies,omitempty" xml:"DeviceRegistrationPolicies,omitempty" type:"Struct"`
	DlpSendPolicies            *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies            `json:"DlpSendPolicies,omitempty" xml:"DlpSendPolicies,omitempty" type:"Struct"`
	DomainBlacklistPolicies    *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies    `json:"DomainBlacklistPolicies,omitempty" xml:"DomainBlacklistPolicies,omitempty" type:"Struct"`
	DomainWhitelistPolicies    *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies    `json:"DomainWhitelistPolicies,omitempty" xml:"DomainWhitelistPolicies,omitempty" type:"Struct"`
	EndpointHardeningPolicies  *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies  `json:"EndpointHardeningPolicies,omitempty" xml:"EndpointHardeningPolicies,omitempty" type:"Struct"`
	EventLabel                 *string                                                             `json:"EventLabel,omitempty" xml:"EventLabel,omitempty"`
	ExternalConfig             *string                                                             `json:"ExternalConfig,omitempty" xml:"ExternalConfig,omitempty"`
	PeripheraBlockPolicies     *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies     `json:"PeripheraBlockPolicies,omitempty" xml:"PeripheraBlockPolicies,omitempty" type:"Struct"`
	// example:
	//
	// approval-process-2677fcf063f5****
	ProcessId                 *string                                                            `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ProcessName               *string                                                            `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessNodes              [][]*UpdateApprovalProcessResponseBodyProcessProcessNodes          `json:"ProcessNodes,omitempty" xml:"ProcessNodes,omitempty" type:"Repeated"`
	SoftwareBlockPolicies     *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies     `json:"SoftwareBlockPolicies,omitempty" xml:"SoftwareBlockPolicies,omitempty" type:"Struct"`
	SoftwareHardeningPolicies *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies `json:"SoftwareHardeningPolicies,omitempty" xml:"SoftwareHardeningPolicies,omitempty" type:"Struct"`
}

func (s UpdateApprovalProcessResponseBodyProcess) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcess) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetAppUninstallPolicies() *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies {
	return s.AppUninstallPolicies
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetApprovalType() *int32 {
	return s.ApprovalType
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetDescription() *string {
	return s.Description
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetDeviceRegistrationPolicies() *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	return s.DeviceRegistrationPolicies
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetDlpSendPolicies() *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies {
	return s.DlpSendPolicies
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetDomainBlacklistPolicies() *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	return s.DomainBlacklistPolicies
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetDomainWhitelistPolicies() *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	return s.DomainWhitelistPolicies
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetEndpointHardeningPolicies() *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	return s.EndpointHardeningPolicies
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetEventLabel() *string {
	return s.EventLabel
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetExternalConfig() *string {
	return s.ExternalConfig
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetPeripheraBlockPolicies() *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies {
	return s.PeripheraBlockPolicies
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetProcessId() *string {
	return s.ProcessId
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetProcessName() *string {
	return s.ProcessName
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetProcessNodes() [][]*UpdateApprovalProcessResponseBodyProcessProcessNodes {
	return s.ProcessNodes
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetSoftwareBlockPolicies() *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	return s.SoftwareBlockPolicies
}

func (s *UpdateApprovalProcessResponseBodyProcess) GetSoftwareHardeningPolicies() *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	return s.SoftwareHardeningPolicies
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetAppUninstallPolicies(v *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) *UpdateApprovalProcessResponseBodyProcess {
	s.AppUninstallPolicies = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetApprovalType(v int32) *UpdateApprovalProcessResponseBodyProcess {
	s.ApprovalType = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetCreateTime(v string) *UpdateApprovalProcessResponseBodyProcess {
	s.CreateTime = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetDescription(v string) *UpdateApprovalProcessResponseBodyProcess {
	s.Description = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetDeviceRegistrationPolicies(v *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) *UpdateApprovalProcessResponseBodyProcess {
	s.DeviceRegistrationPolicies = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetDlpSendPolicies(v *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) *UpdateApprovalProcessResponseBodyProcess {
	s.DlpSendPolicies = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetDomainBlacklistPolicies(v *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) *UpdateApprovalProcessResponseBodyProcess {
	s.DomainBlacklistPolicies = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetDomainWhitelistPolicies(v *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) *UpdateApprovalProcessResponseBodyProcess {
	s.DomainWhitelistPolicies = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetEndpointHardeningPolicies(v *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) *UpdateApprovalProcessResponseBodyProcess {
	s.EndpointHardeningPolicies = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetEventLabel(v string) *UpdateApprovalProcessResponseBodyProcess {
	s.EventLabel = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetExternalConfig(v string) *UpdateApprovalProcessResponseBodyProcess {
	s.ExternalConfig = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetPeripheraBlockPolicies(v *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) *UpdateApprovalProcessResponseBodyProcess {
	s.PeripheraBlockPolicies = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetProcessId(v string) *UpdateApprovalProcessResponseBodyProcess {
	s.ProcessId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetProcessName(v string) *UpdateApprovalProcessResponseBodyProcess {
	s.ProcessName = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetProcessNodes(v [][]*UpdateApprovalProcessResponseBodyProcessProcessNodes) *UpdateApprovalProcessResponseBodyProcess {
	s.ProcessNodes = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetSoftwareBlockPolicies(v *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) *UpdateApprovalProcessResponseBodyProcess {
	s.SoftwareBlockPolicies = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) SetSoftwareHardeningPolicies(v *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) *UpdateApprovalProcessResponseBodyProcess {
	s.SoftwareHardeningPolicies = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcess) Validate() error {
	if s.AppUninstallPolicies != nil {
		if err := s.AppUninstallPolicies.Validate(); err != nil {
			return err
		}
	}
	if s.DeviceRegistrationPolicies != nil {
		if err := s.DeviceRegistrationPolicies.Validate(); err != nil {
			return err
		}
	}
	if s.DlpSendPolicies != nil {
		if err := s.DlpSendPolicies.Validate(); err != nil {
			return err
		}
	}
	if s.DomainBlacklistPolicies != nil {
		if err := s.DomainBlacklistPolicies.Validate(); err != nil {
			return err
		}
	}
	if s.DomainWhitelistPolicies != nil {
		if err := s.DomainWhitelistPolicies.Validate(); err != nil {
			return err
		}
	}
	if s.EndpointHardeningPolicies != nil {
		if err := s.EndpointHardeningPolicies.Validate(); err != nil {
			return err
		}
	}
	if s.PeripheraBlockPolicies != nil {
		if err := s.PeripheraBlockPolicies.Validate(); err != nil {
			return err
		}
	}
	if s.SoftwareBlockPolicies != nil {
		if err := s.SoftwareBlockPolicies.Validate(); err != nil {
			return err
		}
	}
	if s.SoftwareHardeningPolicies != nil {
		if err := s.SoftwareHardeningPolicies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies struct {
	ExternalProcessId *string                                                                 `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	PolicyIds         []*string                                                               `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) GetFieldMap() []*UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) SetExternalProcessId(v string) *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) SetFieldMap(v []*UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) SetPolicyIds(v []*string) *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies {
	s.PolicyIds = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) SetSchemaId(v string) *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) Validate() error {
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

type UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) SetDisplayField(v string) *UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) SetSystemField(v string) *UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies struct {
	ExternalProcessId *string                                                                       `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	PolicyIds         []*string                                                                     `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetFieldMap() []*UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) SetExternalProcessId(v string) *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) SetFieldMap(v []*UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) SetPolicyIds(v []*string) *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	s.PolicyIds = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) SetSchemaId(v string) *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) Validate() error {
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

type UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) SetDisplayField(v string) *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) SetSystemField(v string) *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessDlpSendPolicies struct {
	ExternalProcessId *string                                                            `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	PolicyIds         []*string                                                          `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) GetFieldMap() []*UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) SetExternalProcessId(v string) *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) SetFieldMap(v []*UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) SetPolicyIds(v []*string) *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies {
	s.PolicyIds = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) SetSchemaId(v string) *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) Validate() error {
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

type UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) SetDisplayField(v string) *UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) SetSystemField(v string) *UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies struct {
	ExternalProcessId *string                                                                    `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	PolicyIds         []*string                                                                  `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetFieldMap() []*UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) SetExternalProcessId(v string) *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) SetFieldMap(v []*UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) SetPolicyIds(v []*string) *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	s.PolicyIds = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) SetSchemaId(v string) *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) Validate() error {
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

type UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) SetDisplayField(v string) *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) SetSystemField(v string) *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies struct {
	ExternalProcessId *string                                                                    `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	PolicyIds         []*string                                                                  `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetFieldMap() []*UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) SetExternalProcessId(v string) *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) SetFieldMap(v []*UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) SetPolicyIds(v []*string) *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	s.PolicyIds = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) SetSchemaId(v string) *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) Validate() error {
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

type UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) SetDisplayField(v string) *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) SetSystemField(v string) *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies struct {
	ExternalProcessId *string                                                                      `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	PolicyIds         []*string                                                                    `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	SchemaId          *string                                                                      `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetFieldMap() []*UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) SetExternalProcessId(v string) *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) SetFieldMap(v []*UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) SetPolicyIds(v []*string) *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	s.PolicyIds = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) SetSchemaId(v string) *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) Validate() error {
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

type UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) SetDisplayField(v string) *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) SetSystemField(v string) *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies struct {
	ExternalProcessId *string                                                                   `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	PolicyIds         []*string                                                                 `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) GetFieldMap() []*UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) SetExternalProcessId(v string) *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) SetFieldMap(v []*UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap) *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) SetPolicyIds(v []*string) *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies {
	s.PolicyIds = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) SetSchemaId(v string) *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) Validate() error {
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

type UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap) SetDisplayField(v string) *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap) SetSystemField(v string) *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessProcessNodes struct {
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessProcessNodes) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessProcessNodes) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessProcessNodes) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *UpdateApprovalProcessResponseBodyProcessProcessNodes) GetUsername() *string {
	return s.Username
}

func (s *UpdateApprovalProcessResponseBodyProcessProcessNodes) SetSaseUserId(v string) *UpdateApprovalProcessResponseBodyProcessProcessNodes {
	s.SaseUserId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessProcessNodes) SetUsername(v string) *UpdateApprovalProcessResponseBodyProcessProcessNodes {
	s.Username = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessProcessNodes) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies struct {
	ExternalProcessId *string                                                                  `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	PolicyIds         []*string                                                                `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetFieldMap() []*UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) SetExternalProcessId(v string) *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) SetFieldMap(v []*UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) SetPolicyIds(v []*string) *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	s.PolicyIds = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) SetSchemaId(v string) *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) Validate() error {
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

type UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) SetDisplayField(v string) *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) SetSystemField(v string) *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies struct {
	ExternalProcessId *string                                                                      `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	PolicyIds         []*string                                                                    `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	SchemaId          *string                                                                      `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetFieldMap() []*UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap {
	return s.FieldMap
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) SetExternalProcessId(v string) *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) SetFieldMap(v []*UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	s.FieldMap = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) SetPolicyIds(v []*string) *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	s.PolicyIds = v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) SetSchemaId(v string) *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	s.SchemaId = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) Validate() error {
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

type UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) SetDisplayField(v string) *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) SetDisplayFieldValue(v string) *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) SetSystemField(v string) *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}
