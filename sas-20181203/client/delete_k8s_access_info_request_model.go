// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sAccessInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunYundunGatewayApiName(v string) *DeleteK8sAccessInfoRequest
	GetAliyunYundunGatewayApiName() *string
	SetAliyunYundunGatewayPopName(v string) *DeleteK8sAccessInfoRequest
	GetAliyunYundunGatewayPopName() *string
	SetAliyunYundunGatewayProjectName(v string) *DeleteK8sAccessInfoRequest
	GetAliyunYundunGatewayProjectName() *string
	SetId(v int64) *DeleteK8sAccessInfoRequest
	GetId() *int64
}

type DeleteK8sAccessInfoRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// None
	AliyunYundunGatewayApiName *string `json:"AliyunYundunGatewayApiName,omitempty" xml:"AliyunYundunGatewayApiName,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// None
	AliyunYundunGatewayPopName *string `json:"AliyunYundunGatewayPopName,omitempty" xml:"AliyunYundunGatewayPopName,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// None
	AliyunYundunGatewayProjectName *string `json:"AliyunYundunGatewayProjectName,omitempty" xml:"AliyunYundunGatewayProjectName,omitempty"`
	// The ID generated when Kubernetes is connected. You can call the GenerateK8sAccessInfo operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteK8sAccessInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sAccessInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteK8sAccessInfoRequest) GetAliyunYundunGatewayApiName() *string {
	return s.AliyunYundunGatewayApiName
}

func (s *DeleteK8sAccessInfoRequest) GetAliyunYundunGatewayPopName() *string {
	return s.AliyunYundunGatewayPopName
}

func (s *DeleteK8sAccessInfoRequest) GetAliyunYundunGatewayProjectName() *string {
	return s.AliyunYundunGatewayProjectName
}

func (s *DeleteK8sAccessInfoRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteK8sAccessInfoRequest) SetAliyunYundunGatewayApiName(v string) *DeleteK8sAccessInfoRequest {
	s.AliyunYundunGatewayApiName = &v
	return s
}

func (s *DeleteK8sAccessInfoRequest) SetAliyunYundunGatewayPopName(v string) *DeleteK8sAccessInfoRequest {
	s.AliyunYundunGatewayPopName = &v
	return s
}

func (s *DeleteK8sAccessInfoRequest) SetAliyunYundunGatewayProjectName(v string) *DeleteK8sAccessInfoRequest {
	s.AliyunYundunGatewayProjectName = &v
	return s
}

func (s *DeleteK8sAccessInfoRequest) SetId(v int64) *DeleteK8sAccessInfoRequest {
	s.Id = &v
	return s
}

func (s *DeleteK8sAccessInfoRequest) Validate() error {
	return dara.Validate(s)
}
