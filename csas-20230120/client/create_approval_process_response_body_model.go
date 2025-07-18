// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApprovalProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProcess(v *CreateApprovalProcessResponseBodyProcess) *CreateApprovalProcessResponseBody
	GetProcess() *CreateApprovalProcessResponseBodyProcess
	SetRequestId(v string) *CreateApprovalProcessResponseBody
	GetRequestId() *string
}

type CreateApprovalProcessResponseBody struct {
	Process *CreateApprovalProcessResponseBodyProcess `json:"Process,omitempty" xml:"Process,omitempty" type:"Struct"`
	// example:
	//
	// 2CABFEBB-0CE7-575E-833A-266F75D46713
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApprovalProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponseBody) GetProcess() *CreateApprovalProcessResponseBodyProcess {
	return s.Process
}

func (s *CreateApprovalProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApprovalProcessResponseBody) SetProcess(v *CreateApprovalProcessResponseBodyProcess) *CreateApprovalProcessResponseBody {
	s.Process = v
	return s
}

func (s *CreateApprovalProcessResponseBody) SetRequestId(v string) *CreateApprovalProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApprovalProcessResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateApprovalProcessResponseBodyProcess struct {
	AppUninstallPolicies *CreateApprovalProcessResponseBodyProcessAppUninstallPolicies `json:"AppUninstallPolicies,omitempty" xml:"AppUninstallPolicies,omitempty" type:"Struct"`
	// example:
	//
	// 2022-10-25 10:44:09
	CreateTime                 *string                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description                *string                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceRegistrationPolicies *CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies `json:"DeviceRegistrationPolicies,omitempty" xml:"DeviceRegistrationPolicies,omitempty" type:"Struct"`
	DlpSendPolicies            *CreateApprovalProcessResponseBodyProcessDlpSendPolicies            `json:"DlpSendPolicies,omitempty" xml:"DlpSendPolicies,omitempty" type:"Struct"`
	DomainBlacklistPolicies    *CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies    `json:"DomainBlacklistPolicies,omitempty" xml:"DomainBlacklistPolicies,omitempty" type:"Struct"`
	DomainWhitelistPolicies    *CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies    `json:"DomainWhitelistPolicies,omitempty" xml:"DomainWhitelistPolicies,omitempty" type:"Struct"`
	EndpointHardeningPolicies  *CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies  `json:"EndpointHardeningPolicies,omitempty" xml:"EndpointHardeningPolicies,omitempty" type:"Struct"`
	PeripheralBlockPolicies    *CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies    `json:"PeripheralBlockPolicies,omitempty" xml:"PeripheralBlockPolicies,omitempty" type:"Struct"`
	// example:
	//
	// approval-process-dc61e92ba5c5****
	ProcessId                 *string                                                            `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ProcessName               *string                                                            `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessNodes              [][]*CreateApprovalProcessResponseBodyProcessProcessNodes          `json:"ProcessNodes,omitempty" xml:"ProcessNodes,omitempty" type:"Repeated"`
	SoftwareBlockPolicies     *CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies     `json:"SoftwareBlockPolicies,omitempty" xml:"SoftwareBlockPolicies,omitempty" type:"Struct"`
	SoftwareHardeningPolicies *CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies `json:"SoftwareHardeningPolicies,omitempty" xml:"SoftwareHardeningPolicies,omitempty" type:"Struct"`
}

func (s CreateApprovalProcessResponseBodyProcess) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponseBodyProcess) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponseBodyProcess) GetAppUninstallPolicies() *CreateApprovalProcessResponseBodyProcessAppUninstallPolicies {
	return s.AppUninstallPolicies
}

func (s *CreateApprovalProcessResponseBodyProcess) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateApprovalProcessResponseBodyProcess) GetDescription() *string {
	return s.Description
}

func (s *CreateApprovalProcessResponseBodyProcess) GetDeviceRegistrationPolicies() *CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	return s.DeviceRegistrationPolicies
}

func (s *CreateApprovalProcessResponseBodyProcess) GetDlpSendPolicies() *CreateApprovalProcessResponseBodyProcessDlpSendPolicies {
	return s.DlpSendPolicies
}

func (s *CreateApprovalProcessResponseBodyProcess) GetDomainBlacklistPolicies() *CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	return s.DomainBlacklistPolicies
}

func (s *CreateApprovalProcessResponseBodyProcess) GetDomainWhitelistPolicies() *CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	return s.DomainWhitelistPolicies
}

func (s *CreateApprovalProcessResponseBodyProcess) GetEndpointHardeningPolicies() *CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	return s.EndpointHardeningPolicies
}

func (s *CreateApprovalProcessResponseBodyProcess) GetPeripheralBlockPolicies() *CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies {
	return s.PeripheralBlockPolicies
}

func (s *CreateApprovalProcessResponseBodyProcess) GetProcessId() *string {
	return s.ProcessId
}

func (s *CreateApprovalProcessResponseBodyProcess) GetProcessName() *string {
	return s.ProcessName
}

func (s *CreateApprovalProcessResponseBodyProcess) GetProcessNodes() [][]*CreateApprovalProcessResponseBodyProcessProcessNodes {
	return s.ProcessNodes
}

func (s *CreateApprovalProcessResponseBodyProcess) GetSoftwareBlockPolicies() *CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	return s.SoftwareBlockPolicies
}

func (s *CreateApprovalProcessResponseBodyProcess) GetSoftwareHardeningPolicies() *CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	return s.SoftwareHardeningPolicies
}

func (s *CreateApprovalProcessResponseBodyProcess) SetAppUninstallPolicies(v *CreateApprovalProcessResponseBodyProcessAppUninstallPolicies) *CreateApprovalProcessResponseBodyProcess {
	s.AppUninstallPolicies = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetCreateTime(v string) *CreateApprovalProcessResponseBodyProcess {
	s.CreateTime = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetDescription(v string) *CreateApprovalProcessResponseBodyProcess {
	s.Description = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetDeviceRegistrationPolicies(v *CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) *CreateApprovalProcessResponseBodyProcess {
	s.DeviceRegistrationPolicies = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetDlpSendPolicies(v *CreateApprovalProcessResponseBodyProcessDlpSendPolicies) *CreateApprovalProcessResponseBodyProcess {
	s.DlpSendPolicies = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetDomainBlacklistPolicies(v *CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) *CreateApprovalProcessResponseBodyProcess {
	s.DomainBlacklistPolicies = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetDomainWhitelistPolicies(v *CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) *CreateApprovalProcessResponseBodyProcess {
	s.DomainWhitelistPolicies = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetEndpointHardeningPolicies(v *CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) *CreateApprovalProcessResponseBodyProcess {
	s.EndpointHardeningPolicies = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetPeripheralBlockPolicies(v *CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies) *CreateApprovalProcessResponseBodyProcess {
	s.PeripheralBlockPolicies = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetProcessId(v string) *CreateApprovalProcessResponseBodyProcess {
	s.ProcessId = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetProcessName(v string) *CreateApprovalProcessResponseBodyProcess {
	s.ProcessName = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetProcessNodes(v [][]*CreateApprovalProcessResponseBodyProcessProcessNodes) *CreateApprovalProcessResponseBodyProcess {
	s.ProcessNodes = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetSoftwareBlockPolicies(v *CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) *CreateApprovalProcessResponseBodyProcess {
	s.SoftwareBlockPolicies = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) SetSoftwareHardeningPolicies(v *CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) *CreateApprovalProcessResponseBodyProcess {
	s.SoftwareHardeningPolicies = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcess) Validate() error {
	return dara.Validate(s)
}

type CreateApprovalProcessResponseBodyProcessAppUninstallPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s CreateApprovalProcessResponseBodyProcessAppUninstallPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponseBodyProcessAppUninstallPolicies) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponseBodyProcessAppUninstallPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *CreateApprovalProcessResponseBodyProcessAppUninstallPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *CreateApprovalProcessResponseBodyProcessAppUninstallPolicies) SetPolicyIds(v []*string) *CreateApprovalProcessResponseBodyProcessAppUninstallPolicies {
	s.PolicyIds = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessAppUninstallPolicies) SetSchemaId(v string) *CreateApprovalProcessResponseBodyProcessAppUninstallPolicies {
	s.SchemaId = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessAppUninstallPolicies) Validate() error {
	return dara.Validate(s)
}

type CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) SetPolicyIds(v []*string) *CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	s.PolicyIds = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) SetSchemaId(v string) *CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies {
	s.SchemaId = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) Validate() error {
	return dara.Validate(s)
}

type CreateApprovalProcessResponseBodyProcessDlpSendPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s CreateApprovalProcessResponseBodyProcessDlpSendPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponseBodyProcessDlpSendPolicies) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponseBodyProcessDlpSendPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *CreateApprovalProcessResponseBodyProcessDlpSendPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *CreateApprovalProcessResponseBodyProcessDlpSendPolicies) SetPolicyIds(v []*string) *CreateApprovalProcessResponseBodyProcessDlpSendPolicies {
	s.PolicyIds = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessDlpSendPolicies) SetSchemaId(v string) *CreateApprovalProcessResponseBodyProcessDlpSendPolicies {
	s.SchemaId = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessDlpSendPolicies) Validate() error {
	return dara.Validate(s)
}

type CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) SetPolicyIds(v []*string) *CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	s.PolicyIds = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) SetSchemaId(v string) *CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies {
	s.SchemaId = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) Validate() error {
	return dara.Validate(s)
}

type CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) SetPolicyIds(v []*string) *CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	s.PolicyIds = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) SetSchemaId(v string) *CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies {
	s.SchemaId = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) Validate() error {
	return dara.Validate(s)
}

type CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	SchemaId  *string   `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) SetPolicyIds(v []*string) *CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	s.PolicyIds = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) SetSchemaId(v string) *CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies {
	s.SchemaId = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) Validate() error {
	return dara.Validate(s)
}

type CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies) SetPolicyIds(v []*string) *CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies {
	s.PolicyIds = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies) SetSchemaId(v string) *CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies {
	s.SchemaId = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessPeripheralBlockPolicies) Validate() error {
	return dara.Validate(s)
}

type CreateApprovalProcessResponseBodyProcessProcessNodes struct {
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateApprovalProcessResponseBodyProcessProcessNodes) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponseBodyProcessProcessNodes) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponseBodyProcessProcessNodes) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *CreateApprovalProcessResponseBodyProcessProcessNodes) GetUsername() *string {
	return s.Username
}

func (s *CreateApprovalProcessResponseBodyProcessProcessNodes) SetSaseUserId(v string) *CreateApprovalProcessResponseBodyProcessProcessNodes {
	s.SaseUserId = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessProcessNodes) SetUsername(v string) *CreateApprovalProcessResponseBodyProcessProcessNodes {
	s.Username = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessProcessNodes) Validate() error {
	return dara.Validate(s)
}

type CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) SetPolicyIds(v []*string) *CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	s.PolicyIds = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) SetSchemaId(v string) *CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies {
	s.SchemaId = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) Validate() error {
	return dara.Validate(s)
}

type CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	SchemaId  *string   `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) SetPolicyIds(v []*string) *CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	s.PolicyIds = v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) SetSchemaId(v string) *CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies {
	s.SchemaId = &v
	return s
}

func (s *CreateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) Validate() error {
	return dara.Validate(s)
}
