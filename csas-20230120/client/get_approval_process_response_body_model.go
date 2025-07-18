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
	Process *GetApprovalProcessResponseBodyProcess `json:"Process,omitempty" xml:"Process,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcess struct {
	AppUninstallPolicies *GetApprovalProcessResponseBodyProcessAppUninstallPolicies `json:"AppUninstallPolicies,omitempty" xml:"AppUninstallPolicies,omitempty" type:"Struct"`
	// example:
	//
	// 2022-10-25 10:44:09
	CreateTime                 *string                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description                *string                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceRegistrationPolicies *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies `json:"DeviceRegistrationPolicies,omitempty" xml:"DeviceRegistrationPolicies,omitempty" type:"Struct"`
	DlpSendPolicies            *GetApprovalProcessResponseBodyProcessDlpSendPolicies            `json:"DlpSendPolicies,omitempty" xml:"DlpSendPolicies,omitempty" type:"Struct"`
	DomainBlacklistPolicies    *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies    `json:"DomainBlacklistPolicies,omitempty" xml:"DomainBlacklistPolicies,omitempty" type:"Struct"`
	DomainWhitelistPolicies    *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies    `json:"DomainWhitelistPolicies,omitempty" xml:"DomainWhitelistPolicies,omitempty" type:"Struct"`
	EndpointHardeningPolicies  *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies  `json:"EndpointHardeningPolicies,omitempty" xml:"EndpointHardeningPolicies,omitempty" type:"Struct"`
	PeripheralBlockPolicies    *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies    `json:"PeripheralBlockPolicies,omitempty" xml:"PeripheralBlockPolicies,omitempty" type:"Struct"`
	// example:
	//
	// approval-process-35ee09077ee9****
	ProcessId                 *string                                                         `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ProcessName               *string                                                         `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessNodes              [][]*GetApprovalProcessResponseBodyProcessProcessNodes          `json:"ProcessNodes,omitempty" xml:"ProcessNodes,omitempty" type:"Repeated"`
	SoftwareBlockPolicies     *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies     `json:"SoftwareBlockPolicies,omitempty" xml:"SoftwareBlockPolicies,omitempty" type:"Struct"`
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

func (s *GetApprovalProcessResponseBodyProcess) GetPeripheralBlockPolicies() *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies {
	return s.PeripheralBlockPolicies
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

func (s *GetApprovalProcessResponseBodyProcess) SetPeripheralBlockPolicies(v *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) *GetApprovalProcessResponseBodyProcess {
	s.PeripheralBlockPolicies = v
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
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessAppUninstallPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessAppUninstallPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessDlpSendPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *GetApprovalProcessResponseBodyProcessDlpSendPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessDlpSendPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	SchemaId  *string   `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessPeripheralBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessProcessNodes struct {
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	SchemaId  *string   `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GoString() string {
	return s.String()
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *GetApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}
