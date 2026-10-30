// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPenetrationTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetList(v string) *SubmitPenetrationTestRequest
	GetAssetList() *string
	SetIsThirdpartyTest(v bool) *SubmitPenetrationTestRequest
	GetIsThirdpartyTest() *bool
	SetTestContent(v string) *SubmitPenetrationTestRequest
	GetTestContent() *string
	SetTestEndDate(v string) *SubmitPenetrationTestRequest
	GetTestEndDate() *string
	SetTestSourceIps(v string) *SubmitPenetrationTestRequest
	GetTestSourceIps() *string
	SetTestStartDate(v string) *SubmitPenetrationTestRequest
	GetTestStartDate() *string
	SetTestSteps(v string) *SubmitPenetrationTestRequest
	GetTestSteps() *string
	SetThirdpartyEmail(v string) *SubmitPenetrationTestRequest
	GetThirdpartyEmail() *string
	SetThirdpartyMobile(v string) *SubmitPenetrationTestRequest
	GetThirdpartyMobile() *string
	SetThirdpartyTester(v string) *SubmitPenetrationTestRequest
	GetThirdpartyTester() *string
}

type SubmitPenetrationTestRequest struct {
	// The asset list.
	//
	// > The array supports up to 50 IP addresses.
	//
	// example:
	//
	// [{"AssetType":"SLB","Ip":"47.114.52.93"}]
	AssetList *string `json:"AssetList,omitempty" xml:"AssetList,omitempty"`
	// Specifies whether the test is performed by a third-party testing service provider. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// false
	IsThirdpartyTest *bool `json:"IsThirdpartyTest,omitempty" xml:"IsThirdpartyTest,omitempty"`
	// The test content.
	//
	// > The content must be up to 500 characters in length. Provide detailed information to facilitate the approval process.
	//
	// example:
	//
	// security test
	TestContent *string `json:"TestContent,omitempty" xml:"TestContent,omitempty"`
	// The end time of the test.
	//
	// example:
	//
	// 2026-10-30
	TestEndDate *string `json:"TestEndDate,omitempty" xml:"TestEndDate,omitempty"`
	// The source IP addresses used for the test.
	//
	// > Up to 50 IP addresses are supported. Separate multiple IP addresses with line breaks or commas (,). Enter 0.0.0.0 to allow all IP addresses. CIDR blocks such as 140.168.0.0/16 are not supported.
	//
	// example:
	//
	// 0.0.0.0
	TestSourceIps *string `json:"TestSourceIps,omitempty" xml:"TestSourceIps,omitempty"`
	// The start date of the test.
	//
	// example:
	//
	// 2026-09-30
	TestStartDate *string `json:"TestStartDate,omitempty" xml:"TestStartDate,omitempty"`
	// The test steps.
	//
	// > The steps must be up to 500 characters in length. Provide detailed information to facilitate the approval process.
	//
	// example:
	//
	// none
	TestSteps *string `json:"TestSteps,omitempty" xml:"TestSteps,omitempty"`
	// The email address of the third-party tester.
	//
	// example:
	//
	// 123xxxxx@xx.com
	ThirdpartyEmail *string `json:"ThirdpartyEmail,omitempty" xml:"ThirdpartyEmail,omitempty"`
	// The mobile phone number of the third-party tester.
	//
	// example:
	//
	// 177xxxxxx21
	ThirdpartyMobile *string `json:"ThirdpartyMobile,omitempty" xml:"ThirdpartyMobile,omitempty"`
	// The name of the third-party tester.
	//
	// example:
	//
	// zhangsan
	ThirdpartyTester *string `json:"ThirdpartyTester,omitempty" xml:"ThirdpartyTester,omitempty"`
}

func (s SubmitPenetrationTestRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitPenetrationTestRequest) GoString() string {
	return s.String()
}

func (s *SubmitPenetrationTestRequest) GetAssetList() *string {
	return s.AssetList
}

func (s *SubmitPenetrationTestRequest) GetIsThirdpartyTest() *bool {
	return s.IsThirdpartyTest
}

func (s *SubmitPenetrationTestRequest) GetTestContent() *string {
	return s.TestContent
}

func (s *SubmitPenetrationTestRequest) GetTestEndDate() *string {
	return s.TestEndDate
}

func (s *SubmitPenetrationTestRequest) GetTestSourceIps() *string {
	return s.TestSourceIps
}

func (s *SubmitPenetrationTestRequest) GetTestStartDate() *string {
	return s.TestStartDate
}

func (s *SubmitPenetrationTestRequest) GetTestSteps() *string {
	return s.TestSteps
}

func (s *SubmitPenetrationTestRequest) GetThirdpartyEmail() *string {
	return s.ThirdpartyEmail
}

func (s *SubmitPenetrationTestRequest) GetThirdpartyMobile() *string {
	return s.ThirdpartyMobile
}

func (s *SubmitPenetrationTestRequest) GetThirdpartyTester() *string {
	return s.ThirdpartyTester
}

func (s *SubmitPenetrationTestRequest) SetAssetList(v string) *SubmitPenetrationTestRequest {
	s.AssetList = &v
	return s
}

func (s *SubmitPenetrationTestRequest) SetIsThirdpartyTest(v bool) *SubmitPenetrationTestRequest {
	s.IsThirdpartyTest = &v
	return s
}

func (s *SubmitPenetrationTestRequest) SetTestContent(v string) *SubmitPenetrationTestRequest {
	s.TestContent = &v
	return s
}

func (s *SubmitPenetrationTestRequest) SetTestEndDate(v string) *SubmitPenetrationTestRequest {
	s.TestEndDate = &v
	return s
}

func (s *SubmitPenetrationTestRequest) SetTestSourceIps(v string) *SubmitPenetrationTestRequest {
	s.TestSourceIps = &v
	return s
}

func (s *SubmitPenetrationTestRequest) SetTestStartDate(v string) *SubmitPenetrationTestRequest {
	s.TestStartDate = &v
	return s
}

func (s *SubmitPenetrationTestRequest) SetTestSteps(v string) *SubmitPenetrationTestRequest {
	s.TestSteps = &v
	return s
}

func (s *SubmitPenetrationTestRequest) SetThirdpartyEmail(v string) *SubmitPenetrationTestRequest {
	s.ThirdpartyEmail = &v
	return s
}

func (s *SubmitPenetrationTestRequest) SetThirdpartyMobile(v string) *SubmitPenetrationTestRequest {
	s.ThirdpartyMobile = &v
	return s
}

func (s *SubmitPenetrationTestRequest) SetThirdpartyTester(v string) *SubmitPenetrationTestRequest {
	s.ThirdpartyTester = &v
	return s
}

func (s *SubmitPenetrationTestRequest) Validate() error {
	return dara.Validate(s)
}
