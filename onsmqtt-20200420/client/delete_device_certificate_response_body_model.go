// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeviceCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceSn(v string) *DeleteDeviceCertificateResponseBody
	GetDeviceSn() *string
	SetRequestId(v string) *DeleteDeviceCertificateResponseBody
	GetRequestId() *string
}

type DeleteDeviceCertificateResponseBody struct {
	// The SN of the deleted device certificate. The SN uniquely identifies a device certificate.
	//
	// example:
	//
	// 356217374433******
	DeviceSn *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	// The common parameter. Each request has a unique ID that can be used for troubleshooting.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeviceCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeviceCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceCertificateResponseBody) GetDeviceSn() *string {
	return s.DeviceSn
}

func (s *DeleteDeviceCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeviceCertificateResponseBody) SetDeviceSn(v string) *DeleteDeviceCertificateResponseBody {
	s.DeviceSn = &v
	return s
}

func (s *DeleteDeviceCertificateResponseBody) SetRequestId(v string) *DeleteDeviceCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeviceCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
