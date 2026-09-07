// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackendReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v int64) *CreateBackendReportRequest
	GetEndTimestamp() *int64
	SetPolicyType(v string) *CreateBackendReportRequest
	GetPolicyType() *string
	SetReason(v string) *CreateBackendReportRequest
	GetReason() *string
	SetReportObjects(v []*CreateBackendReportRequestReportObjects) *CreateBackendReportRequest
	GetReportObjects() []*CreateBackendReportRequestReportObjects
	SetTargets(v []*CreateBackendReportRequestTargets) *CreateBackendReportRequest
	GetTargets() []*CreateBackendReportRequestTargets
	SetValidityType(v string) *CreateBackendReportRequest
	GetValidityType() *string
}

type CreateBackendReportRequest struct {
	// The filing expiration time as a UNIX timestamp in seconds. This parameter is required when ValidityType is set to FixedTime or ValidityType is not specified, and the value must be later than the current time. When ValidityType is set to Permanent, do not specify this parameter or set it to 0.
	//
	// example:
	//
	// 1788192000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The filing policy type. Valid values:
	//
	// 	- PrivateAccessBlock: private access.
	//
	// 	- DomainWhitelist: domain name whitelist.
	//
	// 	- DomainBlacklist: domain name blacklist.
	//
	// 	- SoftwareBlock: software blocking.
	//
	// 	- DlpSend: file outbound transfer.
	//
	// 	- PeripheralBlock: peripheral control.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrivateAccessBlock
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The filing reason. The value must be 1 to 1024 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Temporary project access
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The list of filing objects, serialized in Flat format. You can specify 1 to 100 filing objects of the same policy type. The object fields must match the PolicyType value.
	//
	// This parameter is required.
	ReportObjects []*CreateBackendReportRequestReportObjects `json:"ReportObjects,omitempty" xml:"ReportObjects,omitempty" type:"Repeated"`
	// The list of filing users, serialized in Flat format. You can specify 1 to 100 users. Only specific SASE users under the current Alibaba Cloud account are supported. The product of the number of deduplicated users and the number of filing objects cannot exceed 100.
	//
	// This parameter is required.
	Targets []*CreateBackendReportRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// The validity duration type. Default value: FixedTime. Valid values:
	//
	// 	- FixedTime: Expires at the specified time.
	//
	// 	- Permanent: Permanently valid.
	//
	// example:
	//
	// FixedTime
	ValidityType *string `json:"ValidityType,omitempty" xml:"ValidityType,omitempty"`
}

func (s CreateBackendReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendReportRequest) GoString() string {
	return s.String()
}

func (s *CreateBackendReportRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *CreateBackendReportRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateBackendReportRequest) GetReason() *string {
	return s.Reason
}

func (s *CreateBackendReportRequest) GetReportObjects() []*CreateBackendReportRequestReportObjects {
	return s.ReportObjects
}

func (s *CreateBackendReportRequest) GetTargets() []*CreateBackendReportRequestTargets {
	return s.Targets
}

func (s *CreateBackendReportRequest) GetValidityType() *string {
	return s.ValidityType
}

func (s *CreateBackendReportRequest) SetEndTimestamp(v int64) *CreateBackendReportRequest {
	s.EndTimestamp = &v
	return s
}

func (s *CreateBackendReportRequest) SetPolicyType(v string) *CreateBackendReportRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateBackendReportRequest) SetReason(v string) *CreateBackendReportRequest {
	s.Reason = &v
	return s
}

func (s *CreateBackendReportRequest) SetReportObjects(v []*CreateBackendReportRequestReportObjects) *CreateBackendReportRequest {
	s.ReportObjects = v
	return s
}

func (s *CreateBackendReportRequest) SetTargets(v []*CreateBackendReportRequestTargets) *CreateBackendReportRequest {
	s.Targets = v
	return s
}

func (s *CreateBackendReportRequest) SetValidityType(v string) *CreateBackendReportRequest {
	s.ValidityType = &v
	return s
}

func (s *CreateBackendReportRequest) Validate() error {
	if s.ReportObjects != nil {
		for _, item := range s.ReportObjects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Targets != nil {
		for _, item := range s.Targets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateBackendReportRequestReportObjects struct {
	// The private access application ID. This parameter is required when PolicyType is set to PrivateAccessBlock. You can call ListPrivateAccessApplications to query the ID.
	//
	// example:
	//
	// pa-app-****************1234
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The endpoint operating system. This parameter is required when PolicyType is set to PeripheralBlock. Valid values:
	//
	// 	- windows: Windows.
	//
	// 	- macOS: macOS.
	//
	// example:
	//
	// windows
	DevType *string `json:"DevType,omitempty" xml:"DevType,omitempty"`
	// The peripheral channel. This parameter is required when PolicyType is set to PeripheralBlock. Windows supports usbStorage, printer, mobile, cardReader, cdrom, and bluetooth. macOS supports usbStorage, airDrop, mobile, and bluetooth.
	//
	// example:
	//
	// usbStorage
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The file MD5 hash. This parameter is required when PolicyType is set to DlpSend. The value must be a 32-character hexadecimal string and is case-insensitive.
	//
	// example:
	//
	// c936226c4745125b5786527d205a****
	FileMd5 *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// The filing domain name. This parameter is required when PolicyType is set to DomainWhitelist or DomainBlacklist. Regular domain names and wildcard domain names that start with *. are supported. Protocols, ports, and paths are not supported.
	//
	// example:
	//
	// *.example.com
	ReportDomain *string `json:"ReportDomain,omitempty" xml:"ReportDomain,omitempty"`
	// The peripheral filing granularity. This parameter is required when PolicyType is set to PeripheralBlock. Currently, only Channel is supported, which indicates filing by peripheral channel.
	//
	// example:
	//
	// Channel
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The blocked software ID. This parameter is required when PolicyType is set to SoftwareBlock.
	//
	// example:
	//
	// swb-c717ee516145****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
}

func (s CreateBackendReportRequestReportObjects) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendReportRequestReportObjects) GoString() string {
	return s.String()
}

func (s *CreateBackendReportRequestReportObjects) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateBackendReportRequestReportObjects) GetDevType() *string {
	return s.DevType
}

func (s *CreateBackendReportRequestReportObjects) GetDeviceType() *string {
	return s.DeviceType
}

func (s *CreateBackendReportRequestReportObjects) GetFileMd5() *string {
	return s.FileMd5
}

func (s *CreateBackendReportRequestReportObjects) GetReportDomain() *string {
	return s.ReportDomain
}

func (s *CreateBackendReportRequestReportObjects) GetScope() *string {
	return s.Scope
}

func (s *CreateBackendReportRequestReportObjects) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *CreateBackendReportRequestReportObjects) SetApplicationId(v string) *CreateBackendReportRequestReportObjects {
	s.ApplicationId = &v
	return s
}

func (s *CreateBackendReportRequestReportObjects) SetDevType(v string) *CreateBackendReportRequestReportObjects {
	s.DevType = &v
	return s
}

func (s *CreateBackendReportRequestReportObjects) SetDeviceType(v string) *CreateBackendReportRequestReportObjects {
	s.DeviceType = &v
	return s
}

func (s *CreateBackendReportRequestReportObjects) SetFileMd5(v string) *CreateBackendReportRequestReportObjects {
	s.FileMd5 = &v
	return s
}

func (s *CreateBackendReportRequestReportObjects) SetReportDomain(v string) *CreateBackendReportRequestReportObjects {
	s.ReportDomain = &v
	return s
}

func (s *CreateBackendReportRequestReportObjects) SetScope(v string) *CreateBackendReportRequestReportObjects {
	s.Scope = &v
	return s
}

func (s *CreateBackendReportRequestReportObjects) SetSoftwareId(v string) *CreateBackendReportRequestReportObjects {
	s.SoftwareId = &v
	return s
}

func (s *CreateBackendReportRequestReportObjects) Validate() error {
	return dara.Validate(s)
}

type CreateBackendReportRequestTargets struct {
	// The SASE user ID. You can call ListUsers to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// su_8548af20c3b30e931e75cd847a4c****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateBackendReportRequestTargets) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendReportRequestTargets) GoString() string {
	return s.String()
}

func (s *CreateBackendReportRequestTargets) GetUserId() *string {
	return s.UserId
}

func (s *CreateBackendReportRequestTargets) SetUserId(v string) *CreateBackendReportRequestTargets {
	s.UserId = &v
	return s
}

func (s *CreateBackendReportRequestTargets) Validate() error {
	return dara.Validate(s)
}
