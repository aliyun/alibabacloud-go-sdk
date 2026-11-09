// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDeviceCertificateResponseBodyData) *GetDeviceCertificateResponseBody
	GetData() *GetDeviceCertificateResponseBodyData
	SetRequestId(v string) *GetDeviceCertificateResponseBody
	GetRequestId() *string
}

type GetDeviceCertificateResponseBody struct {
	// Certificate details.
	Data *GetDeviceCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Public parameters, each request ID is unique and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceCertificateResponseBody) GetData() *GetDeviceCertificateResponseBodyData {
	return s.Data
}

func (s *GetDeviceCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceCertificateResponseBody) SetData(v *GetDeviceCertificateResponseBodyData) *GetDeviceCertificateResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceCertificateResponseBody) SetRequestId(v string) *GetDeviceCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceCertificateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeviceCertificateResponseBodyData struct {
	// The SN serial number of the CA certificate to which the device certificate belongs, used to uniquely identify a CA certificate.
	//
	// example:
	//
	// 00f26900ba87******
	CaSn *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	// Content of the device certificate.
	//
	//  represents a new line.
	//
	// example:
	//
	// -----BEGIN DEVICECERTIFICATE-----\\nMIIDuzCCAqdGVzdC5jbi1xaW5n******\\n-----END DEVICECERTIFICATE-----
	DeviceContent *string `json:"DeviceContent,omitempty" xml:"DeviceContent,omitempty"`
	// Name of the device certificate.
	//
	// example:
	//
	// mqtt_device
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The SN serial number of the device certificate, used to uniquely identify a device certificate.
	//
	// example:
	//
	// 356217374433******
	DeviceSn *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	// The status of the device certificate. The values are as follows:
	//
	// - **0**: Indicates that the certificate is in an inactive state. - **1**: Indicates that the certificate is in an active state.
	//
	// > After the device certificate is registered, it is in an active state by default.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The start time when the device certificate becomes effective. The format is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1654137303000
	ValidBegin *string `json:"ValidBegin,omitempty" xml:"ValidBegin,omitempty"`
	// The end time when the device certificate becomes effective. The format is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1969497303000
	ValidEnd *string `json:"ValidEnd,omitempty" xml:"ValidEnd,omitempty"`
}

func (s GetDeviceCertificateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceCertificateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceCertificateResponseBodyData) GetCaSn() *string {
	return s.CaSn
}

func (s *GetDeviceCertificateResponseBodyData) GetDeviceContent() *string {
	return s.DeviceContent
}

func (s *GetDeviceCertificateResponseBodyData) GetDeviceName() *string {
	return s.DeviceName
}

func (s *GetDeviceCertificateResponseBodyData) GetDeviceSn() *string {
	return s.DeviceSn
}

func (s *GetDeviceCertificateResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDeviceCertificateResponseBodyData) GetValidBegin() *string {
	return s.ValidBegin
}

func (s *GetDeviceCertificateResponseBodyData) GetValidEnd() *string {
	return s.ValidEnd
}

func (s *GetDeviceCertificateResponseBodyData) SetCaSn(v string) *GetDeviceCertificateResponseBodyData {
	s.CaSn = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) SetDeviceContent(v string) *GetDeviceCertificateResponseBodyData {
	s.DeviceContent = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) SetDeviceName(v string) *GetDeviceCertificateResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) SetDeviceSn(v string) *GetDeviceCertificateResponseBodyData {
	s.DeviceSn = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) SetStatus(v string) *GetDeviceCertificateResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) SetValidBegin(v string) *GetDeviceCertificateResponseBodyData {
	s.ValidBegin = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) SetValidEnd(v string) *GetDeviceCertificateResponseBodyData {
	s.ValidEnd = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
