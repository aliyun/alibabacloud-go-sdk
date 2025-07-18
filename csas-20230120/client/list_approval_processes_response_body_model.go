// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalProcessesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProcesses(v []*ListApprovalProcessesResponseBodyProcesses) *ListApprovalProcessesResponseBody
	GetProcesses() []*ListApprovalProcessesResponseBodyProcesses
	SetRequestId(v string) *ListApprovalProcessesResponseBody
	GetRequestId() *string
	SetTotalNum(v string) *ListApprovalProcessesResponseBody
	GetTotalNum() *string
}

type ListApprovalProcessesResponseBody struct {
	Processes []*ListApprovalProcessesResponseBodyProcesses `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Repeated"`
	// example:
	//
	// 7E39C33B-F565-55C6-ACC2-953FCE7DA7D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *string `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListApprovalProcessesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponseBody) GetProcesses() []*ListApprovalProcessesResponseBodyProcesses {
	return s.Processes
}

func (s *ListApprovalProcessesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApprovalProcessesResponseBody) GetTotalNum() *string {
	return s.TotalNum
}

func (s *ListApprovalProcessesResponseBody) SetProcesses(v []*ListApprovalProcessesResponseBodyProcesses) *ListApprovalProcessesResponseBody {
	s.Processes = v
	return s
}

func (s *ListApprovalProcessesResponseBody) SetRequestId(v string) *ListApprovalProcessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApprovalProcessesResponseBody) SetTotalNum(v string) *ListApprovalProcessesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListApprovalProcessesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesResponseBodyProcesses struct {
	AppUninstallPolicies *ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies `json:"AppUninstallPolicies,omitempty" xml:"AppUninstallPolicies,omitempty" type:"Struct"`
	// example:
	//
	// 2024-02-27 14:04:27
	CreateTime                 *string                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description                *string                                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceRegistrationPolicies *ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies `json:"DeviceRegistrationPolicies,omitempty" xml:"DeviceRegistrationPolicies,omitempty" type:"Struct"`
	DlpSendPolicies            *ListApprovalProcessesResponseBodyProcessesDlpSendPolicies            `json:"DlpSendPolicies,omitempty" xml:"DlpSendPolicies,omitempty" type:"Struct"`
	DomainBlacklistPolicies    *ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies    `json:"DomainBlacklistPolicies,omitempty" xml:"DomainBlacklistPolicies,omitempty" type:"Struct"`
	DomainWhitelistPolicies    *ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies    `json:"DomainWhitelistPolicies,omitempty" xml:"DomainWhitelistPolicies,omitempty" type:"Struct"`
	EndpointHardeningPolicies  *ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies  `json:"EndpointHardeningPolicies,omitempty" xml:"EndpointHardeningPolicies,omitempty" type:"Struct"`
	PeripheralBlockPolicies    *ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies    `json:"PeripheralBlockPolicies,omitempty" xml:"PeripheralBlockPolicies,omitempty" type:"Struct"`
	// example:
	//
	// approval-process-35ee09077ee9****
	ProcessId                 *string                                                              `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	ProcessName               *string                                                              `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessNodes              [][]*ListApprovalProcessesResponseBodyProcessesProcessNodes          `json:"ProcessNodes,omitempty" xml:"ProcessNodes,omitempty" type:"Repeated"`
	SoftwareBlockPolicies     *ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies     `json:"SoftwareBlockPolicies,omitempty" xml:"SoftwareBlockPolicies,omitempty" type:"Struct"`
	SoftwareHardeningPolicies *ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies `json:"SoftwareHardeningPolicies,omitempty" xml:"SoftwareHardeningPolicies,omitempty" type:"Struct"`
}

func (s ListApprovalProcessesResponseBodyProcesses) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponseBodyProcesses) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetAppUninstallPolicies() *ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies {
	return s.AppUninstallPolicies
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetDescription() *string {
	return s.Description
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetDeviceRegistrationPolicies() *ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies {
	return s.DeviceRegistrationPolicies
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetDlpSendPolicies() *ListApprovalProcessesResponseBodyProcessesDlpSendPolicies {
	return s.DlpSendPolicies
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetDomainBlacklistPolicies() *ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies {
	return s.DomainBlacklistPolicies
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetDomainWhitelistPolicies() *ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies {
	return s.DomainWhitelistPolicies
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetEndpointHardeningPolicies() *ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies {
	return s.EndpointHardeningPolicies
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetPeripheralBlockPolicies() *ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies {
	return s.PeripheralBlockPolicies
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetProcessName() *string {
	return s.ProcessName
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetProcessNodes() [][]*ListApprovalProcessesResponseBodyProcessesProcessNodes {
	return s.ProcessNodes
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetSoftwareBlockPolicies() *ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies {
	return s.SoftwareBlockPolicies
}

func (s *ListApprovalProcessesResponseBodyProcesses) GetSoftwareHardeningPolicies() *ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies {
	return s.SoftwareHardeningPolicies
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetAppUninstallPolicies(v *ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies) *ListApprovalProcessesResponseBodyProcesses {
	s.AppUninstallPolicies = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetCreateTime(v string) *ListApprovalProcessesResponseBodyProcesses {
	s.CreateTime = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetDescription(v string) *ListApprovalProcessesResponseBodyProcesses {
	s.Description = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetDeviceRegistrationPolicies(v *ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies) *ListApprovalProcessesResponseBodyProcesses {
	s.DeviceRegistrationPolicies = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetDlpSendPolicies(v *ListApprovalProcessesResponseBodyProcessesDlpSendPolicies) *ListApprovalProcessesResponseBodyProcesses {
	s.DlpSendPolicies = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetDomainBlacklistPolicies(v *ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies) *ListApprovalProcessesResponseBodyProcesses {
	s.DomainBlacklistPolicies = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetDomainWhitelistPolicies(v *ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies) *ListApprovalProcessesResponseBodyProcesses {
	s.DomainWhitelistPolicies = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetEndpointHardeningPolicies(v *ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies) *ListApprovalProcessesResponseBodyProcesses {
	s.EndpointHardeningPolicies = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetPeripheralBlockPolicies(v *ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies) *ListApprovalProcessesResponseBodyProcesses {
	s.PeripheralBlockPolicies = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetProcessId(v string) *ListApprovalProcessesResponseBodyProcesses {
	s.ProcessId = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetProcessName(v string) *ListApprovalProcessesResponseBodyProcesses {
	s.ProcessName = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetProcessNodes(v [][]*ListApprovalProcessesResponseBodyProcessesProcessNodes) *ListApprovalProcessesResponseBodyProcesses {
	s.ProcessNodes = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetSoftwareBlockPolicies(v *ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies) *ListApprovalProcessesResponseBodyProcesses {
	s.SoftwareBlockPolicies = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) SetSoftwareHardeningPolicies(v *ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies) *ListApprovalProcessesResponseBodyProcesses {
	s.SoftwareHardeningPolicies = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcesses) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies) SetPolicyIds(v []*string) *ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies {
	s.PolicyIds = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies) SetSchemaId(v string) *ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesAppUninstallPolicies) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies) SetPolicyIds(v []*string) *ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies {
	s.PolicyIds = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies) SetSchemaId(v string) *ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesDeviceRegistrationPolicies) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesResponseBodyProcessesDlpSendPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s ListApprovalProcessesResponseBodyProcessesDlpSendPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponseBodyProcessesDlpSendPolicies) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponseBodyProcessesDlpSendPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListApprovalProcessesResponseBodyProcessesDlpSendPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalProcessesResponseBodyProcessesDlpSendPolicies) SetPolicyIds(v []*string) *ListApprovalProcessesResponseBodyProcessesDlpSendPolicies {
	s.PolicyIds = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesDlpSendPolicies) SetSchemaId(v string) *ListApprovalProcessesResponseBodyProcessesDlpSendPolicies {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesDlpSendPolicies) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies) SetPolicyIds(v []*string) *ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies {
	s.PolicyIds = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies) SetSchemaId(v string) *ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesDomainBlacklistPolicies) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies) SetPolicyIds(v []*string) *ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies {
	s.PolicyIds = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies) SetSchemaId(v string) *ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesDomainWhitelistPolicies) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	SchemaId  *string   `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies) SetPolicyIds(v []*string) *ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies {
	s.PolicyIds = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies) SetSchemaId(v string) *ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesEndpointHardeningPolicies) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies) SetPolicyIds(v []*string) *ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies {
	s.PolicyIds = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies) SetSchemaId(v string) *ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesPeripheralBlockPolicies) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesResponseBodyProcessesProcessNodes struct {
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListApprovalProcessesResponseBodyProcessesProcessNodes) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponseBodyProcessesProcessNodes) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponseBodyProcessesProcessNodes) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListApprovalProcessesResponseBodyProcessesProcessNodes) GetUsername() *string {
	return s.Username
}

func (s *ListApprovalProcessesResponseBodyProcessesProcessNodes) SetSaseUserId(v string) *ListApprovalProcessesResponseBodyProcessesProcessNodes {
	s.SaseUserId = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesProcessNodes) SetUsername(v string) *ListApprovalProcessesResponseBodyProcessesProcessNodes {
	s.Username = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesProcessNodes) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies) SetPolicyIds(v []*string) *ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies {
	s.PolicyIds = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies) SetSchemaId(v string) *ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesSoftwareBlockPolicies) Validate() error {
	return dara.Validate(s)
}

type ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	SchemaId  *string   `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies) SetPolicyIds(v []*string) *ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies {
	s.PolicyIds = v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies) SetSchemaId(v string) *ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalProcessesResponseBodyProcessesSoftwareHardeningPolicies) Validate() error {
	return dara.Validate(s)
}
