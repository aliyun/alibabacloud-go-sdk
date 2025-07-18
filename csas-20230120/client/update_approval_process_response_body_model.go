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
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcess struct {
	AppUninstallPolicies *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies `json:"AppUninstallPolicies,omitempty" xml:"AppUninstallPolicies,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessAppUninstallPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessDeviceRegistrationPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessDlpSendPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessDlpSendPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainBlacklistPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessDomainWhitelistPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	SchemaId  *string   `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessEndpointHardeningPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessPeripheraBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
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

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareBlockPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}

type UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies struct {
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	SchemaId  *string   `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateApprovalProcessResponseBodyProcessSoftwareHardeningPolicies) GetSchemaId() *string {
	return s.SchemaId
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
	return dara.Validate(s)
}
