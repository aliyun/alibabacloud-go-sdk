// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceCertificateByCaSnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDeviceCertificateByCaSnResponseBodyData) *ListDeviceCertificateByCaSnResponseBody
	GetData() *ListDeviceCertificateByCaSnResponseBodyData
	SetRequestId(v string) *ListDeviceCertificateByCaSnResponseBody
	GetRequestId() *string
}

type ListDeviceCertificateByCaSnResponseBody struct {
	// Query result.
	Data *ListDeviceCertificateByCaSnResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Public parameters, each request ID is unique and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeviceCertificateByCaSnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCertificateByCaSnResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateByCaSnResponseBody) GetData() *ListDeviceCertificateByCaSnResponseBodyData {
	return s.Data
}

func (s *ListDeviceCertificateByCaSnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceCertificateByCaSnResponseBody) SetData(v *ListDeviceCertificateByCaSnResponseBodyData) *ListDeviceCertificateByCaSnResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBody) SetRequestId(v string) *ListDeviceCertificateByCaSnResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDeviceCertificateByCaSnResponseBodyData struct {
	// Details of the device certificate.
	DeviceCertificateVOS []*ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS `json:"DeviceCertificateVOS,omitempty" xml:"DeviceCertificateVOS,omitempty" type:"Repeated"`
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
	// Total number of query results.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDeviceCertificateByCaSnResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCertificateByCaSnResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) GetDeviceCertificateVOS() []*ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	return s.DeviceCertificateVOS
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) SetDeviceCertificateVOS(v []*ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) *ListDeviceCertificateByCaSnResponseBodyData {
	s.DeviceCertificateVOS = v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) SetPageNo(v int32) *ListDeviceCertificateByCaSnResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) SetPageSize(v int32) *ListDeviceCertificateByCaSnResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) SetTotal(v int32) *ListDeviceCertificateByCaSnResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) Validate() error {
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

type ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS struct {
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
	// - 0: indicates that the certificate is in an inactive state.
	//
	// - 1: indicates that the certificate is in an active state.
	//
	// After the device certificate is registered, it is in an active state by default.
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

func (s ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) GetCaSn() *string {
	return s.CaSn
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) GetDeviceContent() *string {
	return s.DeviceContent
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) GetDeviceName() *string {
	return s.DeviceName
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) GetDeviceSn() *string {
	return s.DeviceSn
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) GetStatus() *string {
	return s.Status
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) GetValidBegin() *string {
	return s.ValidBegin
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) GetValidEnd() *string {
	return s.ValidEnd
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetCaSn(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.CaSn = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetDeviceContent(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.DeviceContent = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetDeviceName(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.DeviceName = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetDeviceSn(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.DeviceSn = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetStatus(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.Status = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetValidBegin(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.ValidBegin = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetValidEnd(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.ValidEnd = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) Validate() error {
	return dara.Validate(s)
}
