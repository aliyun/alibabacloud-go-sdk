// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDeviceCertificateResponseBodyData) *ListDeviceCertificateResponseBody
	GetData() *ListDeviceCertificateResponseBodyData
	SetRequestId(v string) *ListDeviceCertificateResponseBody
	GetRequestId() *string
}

type ListDeviceCertificateResponseBody struct {
	// Query result.
	Data *ListDeviceCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Public parameters, each request ID is unique and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeviceCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateResponseBody) GetData() *ListDeviceCertificateResponseBodyData {
	return s.Data
}

func (s *ListDeviceCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceCertificateResponseBody) SetData(v *ListDeviceCertificateResponseBodyData) *ListDeviceCertificateResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceCertificateResponseBody) SetRequestId(v string) *ListDeviceCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceCertificateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDeviceCertificateResponseBodyData struct {
	// Details of the device certificate.
	DeviceCertificateVOS []*ListDeviceCertificateResponseBodyDataDeviceCertificateVOS `json:"DeviceCertificateVOS,omitempty" xml:"DeviceCertificateVOS,omitempty" type:"Repeated"`
	// The current page number of the returned query records.
	//
	// example:
	//
	// 2
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The maximum number of results to display per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Maximum number of pages in the query result.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDeviceCertificateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCertificateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateResponseBodyData) GetDeviceCertificateVOS() []*ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	return s.DeviceCertificateVOS
}

func (s *ListDeviceCertificateResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListDeviceCertificateResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeviceCertificateResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListDeviceCertificateResponseBodyData) SetDeviceCertificateVOS(v []*ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) *ListDeviceCertificateResponseBodyData {
	s.DeviceCertificateVOS = v
	return s
}

func (s *ListDeviceCertificateResponseBodyData) SetPageNo(v int32) *ListDeviceCertificateResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyData) SetPageSize(v int32) *ListDeviceCertificateResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyData) SetTotal(v int32) *ListDeviceCertificateResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyData) Validate() error {
	if s.DeviceCertificateVOS != nil {
		for _, item := range s.DeviceCertificateVOS {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeviceCertificateResponseBodyDataDeviceCertificateVOS struct {
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
	// - 0: indicates that the certificate is in an inactive state. - 1: indicates that the certificate is in an active state.
	//
	// After the device certificate is registered, it defaults to the active state.
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
	// The end time when the device certificate becomes effective. Formatted as a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1969497303000
	ValidEnd *string `json:"ValidEnd,omitempty" xml:"ValidEnd,omitempty"`
}

func (s ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) GetCaSn() *string {
	return s.CaSn
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) GetDeviceContent() *string {
	return s.DeviceContent
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) GetDeviceName() *string {
	return s.DeviceName
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) GetDeviceSn() *string {
	return s.DeviceSn
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) GetStatus() *string {
	return s.Status
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) GetValidBegin() *string {
	return s.ValidBegin
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) GetValidEnd() *string {
	return s.ValidEnd
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetCaSn(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.CaSn = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetDeviceContent(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.DeviceContent = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetDeviceName(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.DeviceName = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetDeviceSn(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.DeviceSn = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetStatus(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.Status = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetValidBegin(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.ValidBegin = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetValidEnd(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.ValidEnd = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) Validate() error {
	return dara.Validate(s)
}
