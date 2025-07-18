// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApprovalProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateApprovalProcessRequest
	GetDescription() *string
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
	Description  *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	MatchSchemas *UpdateApprovalProcessRequestMatchSchemas `json:"MatchSchemas,omitempty" xml:"MatchSchemas,omitempty" type:"Struct"`
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

func (s *UpdateApprovalProcessRequest) GetDescription() *string {
	return s.Description
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

func (s *UpdateApprovalProcessRequest) SetDescription(v string) *UpdateApprovalProcessRequest {
	s.Description = &v
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
