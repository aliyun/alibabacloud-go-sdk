// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateEdgeMobileAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceClass(v string) *ActivateEdgeMobileAgentRequest
	GetDeviceClass() *string
	SetDeviceId(v string) *ActivateEdgeMobileAgentRequest
	GetDeviceId() *string
	SetDeviceMeta(v string) *ActivateEdgeMobileAgentRequest
	GetDeviceMeta() *string
	SetLicenseKey(v string) *ActivateEdgeMobileAgentRequest
	GetLicenseKey() *string
}

type ActivateEdgeMobileAgentRequest struct {
	// The device form factor. Valid values:
	//
	// - BOX
	//
	// - PHONE
	//
	// - PAD
	//
	// - OTHER
	//
	// example:
	//
	// BOX
	DeviceClass *string `json:"DeviceClass,omitempty" xml:"DeviceClass,omitempty"`
	// The unique identifier of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sn-0001eevqa6jeapl*****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The extended device metadata in JSON format. The string contains information such as fingerprint, deviceModel, and firmwareVersion.
	//
	// example:
	//
	// {"frmwareVersion": "1.0.0"}
	DeviceMeta *string `json:"DeviceMeta,omitempty" xml:"DeviceMeta,omitempty"`
	// The license key.
	//
	// This parameter is required.
	//
	// example:
	//
	// lic-ez197xvdf0j5eo0*****
	LicenseKey *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
}

func (s ActivateEdgeMobileAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateEdgeMobileAgentRequest) GoString() string {
	return s.String()
}

func (s *ActivateEdgeMobileAgentRequest) GetDeviceClass() *string {
	return s.DeviceClass
}

func (s *ActivateEdgeMobileAgentRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ActivateEdgeMobileAgentRequest) GetDeviceMeta() *string {
	return s.DeviceMeta
}

func (s *ActivateEdgeMobileAgentRequest) GetLicenseKey() *string {
	return s.LicenseKey
}

func (s *ActivateEdgeMobileAgentRequest) SetDeviceClass(v string) *ActivateEdgeMobileAgentRequest {
	s.DeviceClass = &v
	return s
}

func (s *ActivateEdgeMobileAgentRequest) SetDeviceId(v string) *ActivateEdgeMobileAgentRequest {
	s.DeviceId = &v
	return s
}

func (s *ActivateEdgeMobileAgentRequest) SetDeviceMeta(v string) *ActivateEdgeMobileAgentRequest {
	s.DeviceMeta = &v
	return s
}

func (s *ActivateEdgeMobileAgentRequest) SetLicenseKey(v string) *ActivateEdgeMobileAgentRequest {
	s.LicenseKey = &v
	return s
}

func (s *ActivateEdgeMobileAgentRequest) Validate() error {
	return dara.Validate(s)
}
