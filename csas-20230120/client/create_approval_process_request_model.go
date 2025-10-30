// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApprovalProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateApprovalProcessRequest
	GetDescription() *string
	SetMatchSchemas(v *CreateApprovalProcessRequestMatchSchemas) *CreateApprovalProcessRequest
	GetMatchSchemas() *CreateApprovalProcessRequestMatchSchemas
	SetProcessName(v string) *CreateApprovalProcessRequest
	GetProcessName() *string
	SetProcessNodes(v [][]*string) *CreateApprovalProcessRequest
	GetProcessNodes() [][]*string
}

type CreateApprovalProcessRequest struct {
	Description  *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	MatchSchemas *CreateApprovalProcessRequestMatchSchemas `json:"MatchSchemas,omitempty" xml:"MatchSchemas,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// test_process
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// This parameter is required.
	ProcessNodes [][]*string `json:"ProcessNodes,omitempty" xml:"ProcessNodes,omitempty" type:"Repeated"`
}

func (s CreateApprovalProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessRequest) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApprovalProcessRequest) GetMatchSchemas() *CreateApprovalProcessRequestMatchSchemas {
	return s.MatchSchemas
}

func (s *CreateApprovalProcessRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *CreateApprovalProcessRequest) GetProcessNodes() [][]*string {
	return s.ProcessNodes
}

func (s *CreateApprovalProcessRequest) SetDescription(v string) *CreateApprovalProcessRequest {
	s.Description = &v
	return s
}

func (s *CreateApprovalProcessRequest) SetMatchSchemas(v *CreateApprovalProcessRequestMatchSchemas) *CreateApprovalProcessRequest {
	s.MatchSchemas = v
	return s
}

func (s *CreateApprovalProcessRequest) SetProcessName(v string) *CreateApprovalProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *CreateApprovalProcessRequest) SetProcessNodes(v [][]*string) *CreateApprovalProcessRequest {
	s.ProcessNodes = v
	return s
}

func (s *CreateApprovalProcessRequest) Validate() error {
	if s.MatchSchemas != nil {
		if err := s.MatchSchemas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApprovalProcessRequestMatchSchemas struct {
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

func (s CreateApprovalProcessRequestMatchSchemas) String() string {
	return dara.Prettify(s)
}

func (s CreateApprovalProcessRequestMatchSchemas) GoString() string {
	return s.String()
}

func (s *CreateApprovalProcessRequestMatchSchemas) GetAppUninstallSchemaId() *string {
	return s.AppUninstallSchemaId
}

func (s *CreateApprovalProcessRequestMatchSchemas) GetDeviceRegistrationSchemaId() *string {
	return s.DeviceRegistrationSchemaId
}

func (s *CreateApprovalProcessRequestMatchSchemas) GetDlpSendSchemaId() *string {
	return s.DlpSendSchemaId
}

func (s *CreateApprovalProcessRequestMatchSchemas) GetDomainBlacklistSchemaId() *string {
	return s.DomainBlacklistSchemaId
}

func (s *CreateApprovalProcessRequestMatchSchemas) GetDomainWhitelistSchemaId() *string {
	return s.DomainWhitelistSchemaId
}

func (s *CreateApprovalProcessRequestMatchSchemas) GetEndpointHardeningSchemaId() *string {
	return s.EndpointHardeningSchemaId
}

func (s *CreateApprovalProcessRequestMatchSchemas) GetPeripheralBlockSchemaId() *string {
	return s.PeripheralBlockSchemaId
}

func (s *CreateApprovalProcessRequestMatchSchemas) GetSoftwareBlockSchemaId() *string {
	return s.SoftwareBlockSchemaId
}

func (s *CreateApprovalProcessRequestMatchSchemas) GetSoftwareHardeningSchemaId() *string {
	return s.SoftwareHardeningSchemaId
}

func (s *CreateApprovalProcessRequestMatchSchemas) SetAppUninstallSchemaId(v string) *CreateApprovalProcessRequestMatchSchemas {
	s.AppUninstallSchemaId = &v
	return s
}

func (s *CreateApprovalProcessRequestMatchSchemas) SetDeviceRegistrationSchemaId(v string) *CreateApprovalProcessRequestMatchSchemas {
	s.DeviceRegistrationSchemaId = &v
	return s
}

func (s *CreateApprovalProcessRequestMatchSchemas) SetDlpSendSchemaId(v string) *CreateApprovalProcessRequestMatchSchemas {
	s.DlpSendSchemaId = &v
	return s
}

func (s *CreateApprovalProcessRequestMatchSchemas) SetDomainBlacklistSchemaId(v string) *CreateApprovalProcessRequestMatchSchemas {
	s.DomainBlacklistSchemaId = &v
	return s
}

func (s *CreateApprovalProcessRequestMatchSchemas) SetDomainWhitelistSchemaId(v string) *CreateApprovalProcessRequestMatchSchemas {
	s.DomainWhitelistSchemaId = &v
	return s
}

func (s *CreateApprovalProcessRequestMatchSchemas) SetEndpointHardeningSchemaId(v string) *CreateApprovalProcessRequestMatchSchemas {
	s.EndpointHardeningSchemaId = &v
	return s
}

func (s *CreateApprovalProcessRequestMatchSchemas) SetPeripheralBlockSchemaId(v string) *CreateApprovalProcessRequestMatchSchemas {
	s.PeripheralBlockSchemaId = &v
	return s
}

func (s *CreateApprovalProcessRequestMatchSchemas) SetSoftwareBlockSchemaId(v string) *CreateApprovalProcessRequestMatchSchemas {
	s.SoftwareBlockSchemaId = &v
	return s
}

func (s *CreateApprovalProcessRequestMatchSchemas) SetSoftwareHardeningSchemaId(v string) *CreateApprovalProcessRequestMatchSchemas {
	s.SoftwareHardeningSchemaId = &v
	return s
}

func (s *CreateApprovalProcessRequestMatchSchemas) Validate() error {
	return dara.Validate(s)
}
