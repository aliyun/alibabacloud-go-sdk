// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterServiceInput interface {
	dara.Model
	String() string
	GoString() string
	SetServiceName(v string) *DeregisterServiceInput
	GetServiceName() *string
}

type DeregisterServiceInput struct {
	// 要注销的服务名称（UUID格式）
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789abc
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s DeregisterServiceInput) String() string {
	return dara.Prettify(s)
}

func (s DeregisterServiceInput) GoString() string {
	return s.String()
}

func (s *DeregisterServiceInput) GetServiceName() *string {
	return s.ServiceName
}

func (s *DeregisterServiceInput) SetServiceName(v string) *DeregisterServiceInput {
	s.ServiceName = &v
	return s
}

func (s *DeregisterServiceInput) Validate() error {
	return dara.Validate(s)
}
