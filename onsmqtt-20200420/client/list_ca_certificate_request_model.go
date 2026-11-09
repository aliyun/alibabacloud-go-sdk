// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMqttInstanceId(v string) *ListCaCertificateRequest
	GetMqttInstanceId() *string
	SetPageNo(v string) *ListCaCertificateRequest
	GetPageNo() *string
	SetPageSize(v string) *ListCaCertificateRequest
	GetPageSize() *string
}

type ListCaCertificateRequest struct {
	// The instance ID of the Cloud Message Queue MQTT version, indicating which instance\\"s CA certificates need to be viewed.
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

func (s ListCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *ListCaCertificateRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *ListCaCertificateRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *ListCaCertificateRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListCaCertificateRequest) SetMqttInstanceId(v string) *ListCaCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *ListCaCertificateRequest) SetPageNo(v string) *ListCaCertificateRequest {
	s.PageNo = &v
	return s
}

func (s *ListCaCertificateRequest) SetPageSize(v string) *ListCaCertificateRequest {
	s.PageSize = &v
	return s
}

func (s *ListCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
