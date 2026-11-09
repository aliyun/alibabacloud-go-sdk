// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceCertificateByCaSnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaSn(v string) *ListDeviceCertificateByCaSnRequest
	GetCaSn() *string
	SetMqttInstanceId(v string) *ListDeviceCertificateByCaSnRequest
	GetMqttInstanceId() *string
	SetPageNo(v string) *ListDeviceCertificateByCaSnRequest
	GetPageNo() *string
	SetPageSize(v string) *ListDeviceCertificateByCaSnRequest
	GetPageSize() *string
}

type ListDeviceCertificateByCaSnRequest struct {
	// The SN serial number of the CA certificate to be queried, indicating which CA certificate\\"s registered device certificates are to be retrieved.
	//
	// This parameter is required.
	//
	// example:
	//
	// 007269004887******
	CaSn *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	// The instance ID bound to the CA certificate, which is the instance ID of the MQTT version of the cloud message queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-7mz2d******
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	// Indicates the page number of the returned results. The starting page is counted from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The maximum number of query records to display per page. Value range: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDeviceCertificateByCaSnRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCertificateByCaSnRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateByCaSnRequest) GetCaSn() *string {
	return s.CaSn
}

func (s *ListDeviceCertificateByCaSnRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *ListDeviceCertificateByCaSnRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *ListDeviceCertificateByCaSnRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDeviceCertificateByCaSnRequest) SetCaSn(v string) *ListDeviceCertificateByCaSnRequest {
	s.CaSn = &v
	return s
}

func (s *ListDeviceCertificateByCaSnRequest) SetMqttInstanceId(v string) *ListDeviceCertificateByCaSnRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *ListDeviceCertificateByCaSnRequest) SetPageNo(v string) *ListDeviceCertificateByCaSnRequest {
	s.PageNo = &v
	return s
}

func (s *ListDeviceCertificateByCaSnRequest) SetPageSize(v string) *ListDeviceCertificateByCaSnRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeviceCertificateByCaSnRequest) Validate() error {
	return dara.Validate(s)
}
