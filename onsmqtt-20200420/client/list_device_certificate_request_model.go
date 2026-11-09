// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMqttInstanceId(v string) *ListDeviceCertificateRequest
	GetMqttInstanceId() *string
	SetPageNo(v string) *ListDeviceCertificateRequest
	GetPageNo() *string
	SetPageSize(v string) *ListDeviceCertificateRequest
	GetPageSize() *string
}

type ListDeviceCertificateRequest struct {
	// The instance ID of the Cloud Message Queue MQTT version, indicating which instance\\"s device certificates need to be viewed.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-7mz2d******
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	// Indicates which page of the results to return. The starting page is counted from 1.
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

func (s ListDeviceCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCertificateRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *ListDeviceCertificateRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *ListDeviceCertificateRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDeviceCertificateRequest) SetMqttInstanceId(v string) *ListDeviceCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *ListDeviceCertificateRequest) SetPageNo(v string) *ListDeviceCertificateRequest {
	s.PageNo = &v
	return s
}

func (s *ListDeviceCertificateRequest) SetPageSize(v string) *ListDeviceCertificateRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeviceCertificateRequest) Validate() error {
	return dara.Validate(s)
}
