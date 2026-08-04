// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProcess(v *GetApprovalProcessResponseBodyProcess) *GetApprovalProcessResponseBody
	GetProcess() *GetApprovalProcessResponseBodyProcess
	SetRequestId(v string) *GetApprovalProcessResponseBody
	GetRequestId() *string
}

type GetApprovalProcessResponseBody struct {
	// The approval flow.
	Process *GetApprovalProcessResponseBodyProcess `json:"Process,omitempty" xml:"Process,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C81E0B4B-AAEB-5FDD-B27E-3F5AF7EBD7EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApprovalProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBody) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBody) GetProcess() *GetApprovalProcessResponseBodyProcess {
	return s.Process
}

func (s *GetApprovalProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApprovalProcessResponseBody) SetProcess(v *GetApprovalProcessResponseBodyProcess) *GetApprovalProcessResponseBody {
	s.Process = v
	return s
}

func (s *GetApprovalProcessResponseBody) SetRequestId(v string) *GetApprovalProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApprovalProcessResponseBody) Validate() error {
	if s.Process != nil {
		if err := s.Process.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApprovalProcessResponseBodyProcess struct {
	// The list of policies associated with endpoint uninstallation.
	AppUninstallPolicies *GetApprovalProcessResponseBodyProcessAppUninstallPolicies `json:"AppUninstallPolicies,omitempty" xml:"AppUninstallPolicies,omitempty" type:"Struct"`
	// The approval type.
	ApprovalType *int32 `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The time when the approval flow was created.
	//
	// example:
	//
	// 2022-10-25 10:44:09
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The approval flow description.
	//
	// example:
	//
	// 这是一个审批流程
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of policies associated with device registration.
	DeviceRegistrationPolicies *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies `json:"DeviceRegistrationPolicies,omitempty" xml:"DeviceRegistrationPolicies,omitempty" type:"Struct"`
	// The list of policies associated with file outgoing.
	DlpSendPolicies *GetApprovalProcessResponseBodyProcessDlpSendPolicies `json:"DlpSendPolicies,omitempty" xml:"DlpSendPolicies,omitempty" type:"Struct"`
	// The list of policies associated with the domain name blacklist.
	DomainBlacklistPolicies *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies `json:"DomainBlacklistPolicies,omitempty" xml:"DomainBlacklistPolicies,omitempty" type:"Struct"`
	// The list of policies associated with the domain name whitelist.
	DomainWhitelistPolicies *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies `json:"DomainWhitelistPolicies,omitempty" xml:"DomainWhitelistPolicies,omitempty" type:"Struct"`
	// The endpoint protection policies.
	EndpointHardeningPolicies *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies `json:"EndpointHardeningPolicies,omitempty" xml:"EndpointHardeningPolicies,omitempty" type:"Struct"`
	// The label.
	EventLabel *string `json:"EventLabel,omitempty" xml:"EventLabel,omitempty"`
	// The external configuration.
	ExternalConfig *string `json:"ExternalConfig,omitempty" xml:"ExternalConfig,omitempty"`
	// The list of policies associated with peripheral control.
	PeripheralBlockPolicies    *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies    `json:"PeripheralBlockPolicies,omitempty" xml:"PeripheralBlockPolicies,omitempty" type:"Struct"`
	PrivateAccessBlockPolicies *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies `json:"PrivateAccessBlockPolicies,omitempty" xml:"PrivateAccessBlockPolicies,omitempty" type:"Struct"`
	// The approval flow ID.
	//
	// example:
	//
	// approval-process-35ee09077ee9****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The approval flow name.
	//
	// example:
	//
	// 测试
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The list of approval nodes.
	ProcessNodes [][]*GetApprovalProcessResponseBodyProcessProcessNodes `json:"ProcessNodes,omitempty" xml:"ProcessNodes,omitempty" type:"Repeated"`
	// The list of policies associated with software blocking.
	SoftwareBlockPolicies *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies `json:"SoftwareBlockPolicies,omitempty" xml:"SoftwareBlockPolicies,omitempty" type:"Struct"`
	// The software hardening policies.
	SoftwareHardeningPolicies *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies `json:"SoftwareHardeningPolicies,omitempty" xml:"SoftwareHardeningPolicies,omitempty" type:"Struct"`
}

func (s GetApprovalProcessResponseBodyProcess) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcess) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcess) GetAppUninstallPolicies() *GetApprovalProcessResponseBodyProcessAppUninstallPolicies {
	return s.AppUninstallPolicies
}

func (s *GetApprovalProcessResponseBodyProcess) GetApprovalType() *int32 {
	return s.ApprovalType
}

func (s *GetApprovalProcessResponseBodyProcess) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetApprovalProcessResponseBodyProcess) GetDescription() *string {
	return s.Description
}

func (s *GetApprovalProcessResponseBodyProcess) GetDeviceRegistrationPolicies() *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	return s.DeviceRegistrationPolicies
}

func (s *GetApprovalProcessResponseBodyProcess) GetDlpSendPolicies() *GetApprovalProcessResponseBodyProcessDlpSendPolicies {
	return s.DlpSendPolicies
}

func (s *GetApprovalProcessResponseBodyProcess) GetDomainBlacklistPolicies() *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	return s.DomainBlacklistPolicies
}

func (s *GetApprovalProcessResponseBodyProcess) GetDomainWhitelistPolicies() *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	return s.DomainWhitelistPolicies
}

func (s *GetApprovalProcessResponseBodyProcess) GetEndpointHardeningPolicies() *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	return s.EndpointHardeningPolicies
}

func (s *GetApprovalProcessResponseBodyProcess) GetEventLabel() *string {
	return s.EventLabel
}

func (s *GetApprovalProcessResponseBodyProcess) GetExternalConfig() *string {
	return s.ExternalConfig
}

func (s *GetApprovalProcessResponseBodyProcess) GetPeripheralBlockPolicies() *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies {
	return s.PeripheralBlockPolicies
}

func (s *GetApprovalProcessResponseBodyProcess) GetPrivateAccessBlockPolicies() *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies {
	return s.PrivateAccessBlockPolicies
}

func (s *GetApprovalProcessResponseBodyProcess) GetProcessId() *string {
	return s.ProcessId
}

func (s *GetApprovalProcessResponseBodyProcess) GetProcessName() *string {
	return s.ProcessName
}

func (s *GetApprovalProcessResponseBodyProcess) GetProcessNodes() [][]*GetApprovalProcessResponseBodyProcessProcessNodes {
	return s.ProcessNodes
}

func (s *GetApprovalProcessResponseBodyProcess) GetSoftwareBlockPolicies() *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	return s.SoftwareBlockPolicies
}

func (s *GetApprovalProcessResponseBodyProcess) GetSoftwareHardeningPolicies() *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	return s.SoftwareHardeningPolicies
}

func (s *GetApprovalProcessResponseBodyProcess) SetAppUninstallPolicies(v *GetApprovalProcessResponseBodyProcessAppUninstallPolicies) *GetApprovalProcessResponseBodyProcess {
	s.AppUninstallPolicies = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetApprovalType(v int32) *GetApprovalProcessResponseBodyProcess {
	s.ApprovalType = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetCreateTime(v string) *GetApprovalProcessResponseBodyProcess {
	s.CreateTime = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetDescription(v string) *GetApprovalProcessResponseBodyProcess {
	s.Description = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetDeviceRegistrationPolicies(v *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) *GetApprovalProcessResponseBodyProcess {
	s.DeviceRegistrationPolicies = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetDlpSendPolicies(v *GetApprovalProcessResponseBodyProcessDlpSendPolicies) *GetApprovalProcessResponseBodyProcess {
	s.DlpSendPolicies = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetDomainBlacklistPolicies(v *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) *GetApprovalProcessResponseBodyProcess {
	s.DomainBlacklistPolicies = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetDomainWhitelistPolicies(v *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) *GetApprovalProcessResponseBodyProcess {
	s.DomainWhitelistPolicies = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetEndpointHardeningPolicies(v *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) *GetApprovalProcessResponseBodyProcess {
	s.EndpointHardeningPolicies = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetEventLabel(v string) *GetApprovalProcessResponseBodyProcess {
	s.EventLabel = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetExternalConfig(v string) *GetApprovalProcessResponseBodyProcess {
	s.ExternalConfig = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetPeripheralBlockPolicies(v *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) *GetApprovalProcessResponseBodyProcess {
	s.PeripheralBlockPolicies = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetPrivateAccessBlockPolicies(v *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies) *GetApprovalProcessResponseBodyProcess {
	s.PrivateAccessBlockPolicies = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetProcessId(v string) *GetApprovalProcessResponseBodyProcess {
	s.ProcessId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetProcessName(v string) *GetApprovalProcessResponseBodyProcess {
	s.ProcessName = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetProcessNodes(v [][]*GetApprovalProcessResponseBodyProcessProcessNodes) *GetApprovalProcessResponseBodyProcess {
	s.ProcessNodes = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetSoftwareBlockPolicies(v *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) *GetApprovalProcessResponseBodyProcess {
	s.SoftwareBlockPolicies = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) SetSoftwareHardeningPolicies(v *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) *GetApprovalProcessResponseBodyProcess {
	s.SoftwareHardeningPolicies = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcess) Validate() error {
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
	if s.PeripheralBlockPolicies != nil {
		if err := s.PeripheralBlockPolicies.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateAccessBlockPolicies != nil {
		if err := s.PrivateAccessBlockPolicies.Validate(); err != nil {
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

type GetApprovalProcessResponseBodyProcessAppUninstallPolicies struct {
	// The external flow ID.
	ExternalProcessId *string `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	// The field mapping.
	FieldMap []*GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	// The list of endpoint uninstall policy IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The approval template ID.
	//
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessAppUninstallPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessAppUninstallPolicies) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPolicies) GetFieldMap() []*GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap {
	return s.FieldMap
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPolicies) SetExternalProcessId(v string) *GetApprovalProcessResponseBodyProcessAppUninstallPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPolicies) SetFieldMap(v []*GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) *GetApprovalProcessResponseBodyProcessAppUninstallPolicies {
	s.FieldMap = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPolicies) SetPolicyIds(v []*string) *GetApprovalProcessResponseBodyProcessAppUninstallPolicies {
	s.PolicyIds = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPolicies) SetSchemaId(v string) *GetApprovalProcessResponseBodyProcessAppUninstallPolicies {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPolicies) Validate() error {
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

type GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap struct {
	// The display field.
	//
	// example:
	//
	// 名称
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	// The system field.
	//
	// example:
	//
	// 名称
	SystemField *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) SetDisplayField(v string) *GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) SetDisplayFieldValue(v string) *GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) SetSystemField(v string) *GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies struct {
	// The external flow ID.
	ExternalProcessId *string `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	// The field mapping.
	FieldMap []*GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	// The list of device registration policy IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The approval template ID.
	//
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetFieldMap() []*GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap {
	return s.FieldMap
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) SetExternalProcessId(v string) *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) SetFieldMap(v []*GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	s.FieldMap = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) SetPolicyIds(v []*string) *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	s.PolicyIds = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) SetSchemaId(v string) *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) Validate() error {
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

type GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap struct {
	// The display field.
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	// The system field.
	SystemField *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) SetDisplayField(v string) *GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) SetDisplayFieldValue(v string) *GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) SetSystemField(v string) *GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessDlpSendPolicies struct {
	// The external flow ID.
	ExternalProcessId *string `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	// The field mapping.
	FieldMap []*GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	// The list of file outgoing policy IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The approval template ID.
	//
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessDlpSendPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessDlpSendPolicies) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPolicies) GetFieldMap() []*GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap {
	return s.FieldMap
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPolicies) SetExternalProcessId(v string) *GetApprovalProcessResponseBodyProcessDlpSendPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPolicies) SetFieldMap(v []*GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) *GetApprovalProcessResponseBodyProcessDlpSendPolicies {
	s.FieldMap = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPolicies) SetPolicyIds(v []*string) *GetApprovalProcessResponseBodyProcessDlpSendPolicies {
	s.PolicyIds = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPolicies) SetSchemaId(v string) *GetApprovalProcessResponseBodyProcessDlpSendPolicies {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPolicies) Validate() error {
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

type GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap struct {
	// The display field.
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	// The system field.
	SystemField *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) SetDisplayField(v string) *GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) SetDisplayFieldValue(v string) *GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) SetSystemField(v string) *GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies struct {
	// The external flow ID.
	ExternalProcessId *string `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	// The field mapping.
	FieldMap []*GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	// The list of domain name blacklist policy IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The approval template ID.
	//
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetFieldMap() []*GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap {
	return s.FieldMap
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) SetExternalProcessId(v string) *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) SetFieldMap(v []*GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	s.FieldMap = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) SetPolicyIds(v []*string) *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	s.PolicyIds = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) SetSchemaId(v string) *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) Validate() error {
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

type GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap struct {
	// The display field.
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	// The system field.
	SystemField *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) SetDisplayField(v string) *GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) SetDisplayFieldValue(v string) *GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) SetSystemField(v string) *GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies struct {
	// The external flow ID.
	ExternalProcessId *string `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	// The field mapping table.
	FieldMap []*GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	// The list of domain name whitelist policy IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The approval template ID.
	//
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetFieldMap() []*GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap {
	return s.FieldMap
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) SetExternalProcessId(v string) *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) SetFieldMap(v []*GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	s.FieldMap = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) SetPolicyIds(v []*string) *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	s.PolicyIds = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) SetSchemaId(v string) *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) Validate() error {
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

type GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap struct {
	// The display field.
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	// The system field.
	SystemField *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) SetDisplayField(v string) *GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) SetDisplayFieldValue(v string) *GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) SetSystemField(v string) *GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies struct {
	// The external flow ID.
	ExternalProcessId *string `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	// The field mapping.
	FieldMap []*GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	// The policy IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The template ID.
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetFieldMap() []*GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap {
	return s.FieldMap
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) SetExternalProcessId(v string) *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) SetFieldMap(v []*GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	s.FieldMap = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) SetPolicyIds(v []*string) *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	s.PolicyIds = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) SetSchemaId(v string) *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) Validate() error {
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

type GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap struct {
	// The display field.
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	// The system field.
	SystemField *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) SetDisplayField(v string) *GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) SetDisplayFieldValue(v string) *GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) SetSystemField(v string) *GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies struct {
	// The external flow ID.
	ExternalProcessId *string `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	// The field mapping.
	FieldMap []*GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	// The list of peripheral control policy IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The approval template ID.
	//
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) GetFieldMap() []*GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap {
	return s.FieldMap
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) SetExternalProcessId(v string) *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) SetFieldMap(v []*GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap) *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies {
	s.FieldMap = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) SetPolicyIds(v []*string) *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies {
	s.PolicyIds = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) SetSchemaId(v string) *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) Validate() error {
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

type GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap struct {
	// The display field.
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	// The system field.
	SystemField *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap) SetDisplayField(v string) *GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap) SetDisplayFieldValue(v string) *GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap) SetSystemField(v string) *GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies struct {
	ExternalProcessId *string                                                                    `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	FieldMap          []*GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	PolicyIds         []*string                                                                  `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	SchemaId          *string                                                                    `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies) GetFieldMap() []*GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap {
	return s.FieldMap
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies) SetExternalProcessId(v string) *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies) SetFieldMap(v []*GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap) *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies {
	s.FieldMap = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies) SetPolicyIds(v []*string) *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies {
	s.PolicyIds = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies) SetSchemaId(v string) *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPolicies) Validate() error {
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

type GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap struct {
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	SystemField       *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap) SetDisplayField(v string) *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap) SetDisplayFieldValue(v string) *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap) SetSystemField(v string) *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessPrivateAccessBlockPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessProcessNodes struct {
	// The approver ID.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// The username of the approver.
	//
	// example:
	//
	// 王先生
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessProcessNodes) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessProcessNodes) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessProcessNodes) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *GetApprovalProcessResponseBodyProcessProcessNodes) GetUsername() *string {
	return s.Username
}

func (s *GetApprovalProcessResponseBodyProcessProcessNodes) SetSaseUserId(v string) *GetApprovalProcessResponseBodyProcessProcessNodes {
	s.SaseUserId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessProcessNodes) SetUsername(v string) *GetApprovalProcessResponseBodyProcessProcessNodes {
	s.Username = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessProcessNodes) Validate() error {
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies struct {
	// The external flow ID.
	ExternalProcessId *string `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	// The field mapping.
	FieldMap []*GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	// The list of software blocking policy IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The approval template ID.
	//
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetFieldMap() []*GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap {
	return s.FieldMap
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) SetExternalProcessId(v string) *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) SetFieldMap(v []*GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	s.FieldMap = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) SetPolicyIds(v []*string) *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	s.PolicyIds = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) SetSchemaId(v string) *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) Validate() error {
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

type GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap struct {
	// The display field.
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	// The system field.
	SystemField *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) SetDisplayField(v string) *GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) SetDisplayFieldValue(v string) *GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) SetSystemField(v string) *GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies struct {
	// The external flow ID.
	ExternalProcessId *string `json:"ExternalProcessId,omitempty" xml:"ExternalProcessId,omitempty"`
	// The field mapping.
	FieldMap []*GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap `json:"FieldMap,omitempty" xml:"FieldMap,omitempty" type:"Repeated"`
	// The policy IDs.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The template ID.
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetExternalProcessId() *string {
	return s.ExternalProcessId
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetFieldMap() []*GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap {
	return s.FieldMap
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) SetExternalProcessId(v string) *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	s.ExternalProcessId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) SetFieldMap(v []*GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	s.FieldMap = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) SetPolicyIds(v []*string) *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	s.PolicyIds = v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) SetSchemaId(v string) *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) Validate() error {
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

type GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap struct {
	// The display field.
	DisplayField      *string `json:"DisplayField,omitempty" xml:"DisplayField,omitempty"`
	DisplayFieldValue *string `json:"DisplayFieldValue,omitempty" xml:"DisplayFieldValue,omitempty"`
	// The system field.
	SystemField *string `json:"SystemField,omitempty" xml:"SystemField,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) GetDisplayField() *string {
	return s.DisplayField
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) GetDisplayFieldValue() *string {
	return s.DisplayFieldValue
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) GetSystemField() *string {
	return s.SystemField
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) SetDisplayField(v string) *GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap {
	s.DisplayField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) SetDisplayFieldValue(v string) *GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap {
	s.DisplayFieldValue = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) SetSystemField(v string) *GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap {
	s.SystemField = &v
	return s
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPoliciesFieldMap) Validate() error {
	return dara.Validate(s)
}
