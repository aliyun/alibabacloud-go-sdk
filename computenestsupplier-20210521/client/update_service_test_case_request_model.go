// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceTestCaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdateServiceTestCaseRequest
	GetRegionId() *string
	SetTestCaseId(v string) *UpdateServiceTestCaseRequest
	GetTestCaseId() *string
	SetTestCaseName(v string) *UpdateServiceTestCaseRequest
	GetTestCaseName() *string
	SetTestConfig(v string) *UpdateServiceTestCaseRequest
	GetTestConfig() *string
}

type UpdateServiceTestCaseRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service test case ID
	//
	// This parameter is required.
	//
	// example:
	//
	// stc-2deec15c20b24aaf9f16
	TestCaseId *string `json:"TestCaseId,omitempty" xml:"TestCaseId,omitempty"`
	// Test case name
	//
	// This parameter is required.
	//
	// example:
	//
	// case1
	TestCaseName *string `json:"TestCaseName,omitempty" xml:"TestCaseName,omitempty"`
	// Test configuration
	//
	// This parameter is required.
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

func (s UpdateServiceTestCaseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceTestCaseRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceTestCaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateServiceTestCaseRequest) GetTestCaseId() *string {
	return s.TestCaseId
}

func (s *UpdateServiceTestCaseRequest) GetTestCaseName() *string {
	return s.TestCaseName
}

func (s *UpdateServiceTestCaseRequest) GetTestConfig() *string {
	return s.TestConfig
}

func (s *UpdateServiceTestCaseRequest) SetRegionId(v string) *UpdateServiceTestCaseRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceTestCaseRequest) SetTestCaseId(v string) *UpdateServiceTestCaseRequest {
	s.TestCaseId = &v
	return s
}

func (s *UpdateServiceTestCaseRequest) SetTestCaseName(v string) *UpdateServiceTestCaseRequest {
	s.TestCaseName = &v
	return s
}

func (s *UpdateServiceTestCaseRequest) SetTestConfig(v string) *UpdateServiceTestCaseRequest {
	s.TestConfig = &v
	return s
}

func (s *UpdateServiceTestCaseRequest) Validate() error {
	return dara.Validate(s)
}
