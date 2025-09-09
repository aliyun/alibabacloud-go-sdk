// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDefaultServiceTestConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateDefaultServiceTestConfigResponseBody
	GetRequestId() *string
	SetTestConfig(v string) *GenerateDefaultServiceTestConfigResponseBody
	GetTestConfig() *string
}

type GenerateDefaultServiceTestConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service test config
	//
	// example:
	//
	// ---
	//
	// parameters:
	//
	//   PayType: "PostPaid"
	//
	//   EcsInstanceType: "$[iact3-auto]"
	//
	//   InstancePassword: "$[iact3-auto]"
	TestConfig *string `json:"TestConfig,omitempty" xml:"TestConfig,omitempty"`
}

func (s GenerateDefaultServiceTestConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateDefaultServiceTestConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDefaultServiceTestConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateDefaultServiceTestConfigResponseBody) GetTestConfig() *string {
	return s.TestConfig
}

func (s *GenerateDefaultServiceTestConfigResponseBody) SetRequestId(v string) *GenerateDefaultServiceTestConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDefaultServiceTestConfigResponseBody) SetTestConfig(v string) *GenerateDefaultServiceTestConfigResponseBody {
	s.TestConfig = &v
	return s
}

func (s *GenerateDefaultServiceTestConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
