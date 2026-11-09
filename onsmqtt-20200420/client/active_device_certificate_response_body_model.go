// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveDeviceCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceSn(v string) *ActiveDeviceCertificateResponseBody
	GetDeviceSn() *string
	SetRequestId(v string) *ActiveDeviceCertificateResponseBody
	GetRequestId() *string
}

type ActiveDeviceCertificateResponseBody struct {
	// The SN of the activated device certificate. The SN uniquely identifies a device certificate.
	//
	// example:
	//
	// 356217374433******
	DeviceSn *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	// The common parameter. Each request has a unique ID, which can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActiveDeviceCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActiveDeviceCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveDeviceCertificateResponseBody) GetDeviceSn() *string {
	return s.DeviceSn
}

func (s *ActiveDeviceCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActiveDeviceCertificateResponseBody) SetDeviceSn(v string) *ActiveDeviceCertificateResponseBody {
	s.DeviceSn = &v
	return s
}

func (s *ActiveDeviceCertificateResponseBody) SetRequestId(v string) *ActiveDeviceCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveDeviceCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
