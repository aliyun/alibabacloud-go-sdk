// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPenetrationTestDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPenetrationTestDetailResponseBody
	GetCode() *string
	SetData(v *QueryPenetrationTestDetailResponseBodyData) *QueryPenetrationTestDetailResponseBody
	GetData() *QueryPenetrationTestDetailResponseBodyData
	SetMessage(v string) *QueryPenetrationTestDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryPenetrationTestDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPenetrationTestDetailResponseBody
	GetSuccess() *bool
}

type QueryPenetrationTestDetailResponseBody struct {
	// The status code.
	//
	// - **200**: Success.
	//
	// - **Other values (400 and 500)**: Failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response content.
	Data *QueryPenetrationTestDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 855FCC89-0B13-5FC0-AAD2-120878081C1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - **true**: The call is successful.
	//
	// - **false**: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPenetrationTestDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPenetrationTestDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPenetrationTestDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPenetrationTestDetailResponseBody) GetData() *QueryPenetrationTestDetailResponseBodyData {
	return s.Data
}

func (s *QueryPenetrationTestDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPenetrationTestDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPenetrationTestDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPenetrationTestDetailResponseBody) SetCode(v string) *QueryPenetrationTestDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBody) SetData(v *QueryPenetrationTestDetailResponseBodyData) *QueryPenetrationTestDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryPenetrationTestDetailResponseBody) SetMessage(v string) *QueryPenetrationTestDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBody) SetRequestId(v string) *QueryPenetrationTestDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBody) SetSuccess(v bool) *QueryPenetrationTestDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPenetrationTestDetailResponseBodyData struct {
	// The application time.
	//
	// example:
	//
	// 1788873796000
	ApplyTime *int64 `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The type of the asset. Valid values:
	//
	// - **EIP**: public IP address in a VPC.
	//
	// - **SLB**: load balancing.
	//
	// - **ECS**: public IP address in the classic network.
	//
	// example:
	//
	// EIP
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The unique identity of the application record.
	//
	// > This field is of the Long type. Precision loss may occur during the procedure to serialize or deserialize the numeric sequence. Ensure that the value does not exceed 9007199254740991.
	//
	// example:
	//
	// 2387
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// IP
	//
	// example:
	//
	// 47.113.186.91
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// Indicates whether the test is performed by a third-party service provider. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// false
	IsThirdpartyTest *bool `json:"IsThirdpartyTest,omitempty" xml:"IsThirdpartyTest,omitempty"`
	// The task status. Valid values:
	//
	// - **early_end**: The user applied for early termination.
	//
	// - **expired**: The authorization has expired.
	//
	// - **auditing**: Under review.
	//
	// - **approved**: Approved.
	//
	// - **rejected**: Rejected.
	//
	// example:
	//
	// early_end
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The test content.
	//
	// > - Maximum 500 characters. Provide detailed information to increase the chance of passing the review.
	//
	// example:
	//
	// security test
	TestContent *string `json:"TestContent,omitempty" xml:"TestContent,omitempty"`
	// The test end time.
	//
	// example:
	//
	// 1788873796000
	TestEndDate *int64 `json:"TestEndDate,omitempty" xml:"TestEndDate,omitempty"`
	// The source IP addresses used for the test.
	//
	// > A maximum of 50 IP addresses are supported. Separate multiple IP addresses with line breaks or commas (,). Specify 0.0.0.0 to allow all IP addresses. CIDR blocks are not supported, such as 140.168.0.0/16.
	//
	// example:
	//
	// 1.1.1.1
	TestSourceIps *string `json:"TestSourceIps,omitempty" xml:"TestSourceIps,omitempty"`
	// The test start date.
	//
	// example:
	//
	// 1788873796000
	TestStartDate *int64 `json:"TestStartDate,omitempty" xml:"TestStartDate,omitempty"`
	// The test steps.
	//
	// > Maximum 500 characters. Provide detailed information to increase the chance of passing the review.
	//
	// example:
	//
	// none
	TestSteps *string `json:"TestSteps,omitempty" xml:"TestSteps,omitempty"`
	// The email address of the third-party tester.
	//
	// example:
	//
	// 12312@xxx.com
	ThirdpartyEmail *string `json:"ThirdpartyEmail,omitempty" xml:"ThirdpartyEmail,omitempty"`
	// The mobile phone number of the third-party tester.
	//
	// example:
	//
	// 177xxxxxx21
	ThirdpartyMobile *string `json:"ThirdpartyMobile,omitempty" xml:"ThirdpartyMobile,omitempty"`
	// The third-party tester.
	//
	// example:
	//
	// zhangsan
	ThirdpartyTester *string `json:"ThirdpartyTester,omitempty" xml:"ThirdpartyTester,omitempty"`
}

func (s QueryPenetrationTestDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPenetrationTestDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetApplyTime() *int64 {
	return s.ApplyTime
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetAssetType() *string {
	return s.AssetType
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetIp() *string {
	return s.Ip
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetIsThirdpartyTest() *bool {
	return s.IsThirdpartyTest
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetTestContent() *string {
	return s.TestContent
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetTestEndDate() *int64 {
	return s.TestEndDate
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetTestSourceIps() *string {
	return s.TestSourceIps
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetTestStartDate() *int64 {
	return s.TestStartDate
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetTestSteps() *string {
	return s.TestSteps
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetThirdpartyEmail() *string {
	return s.ThirdpartyEmail
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetThirdpartyMobile() *string {
	return s.ThirdpartyMobile
}

func (s *QueryPenetrationTestDetailResponseBodyData) GetThirdpartyTester() *string {
	return s.ThirdpartyTester
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetApplyTime(v int64) *QueryPenetrationTestDetailResponseBodyData {
	s.ApplyTime = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetAssetType(v string) *QueryPenetrationTestDetailResponseBodyData {
	s.AssetType = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetId(v int64) *QueryPenetrationTestDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetIp(v string) *QueryPenetrationTestDetailResponseBodyData {
	s.Ip = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetIsThirdpartyTest(v bool) *QueryPenetrationTestDetailResponseBodyData {
	s.IsThirdpartyTest = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetStatus(v string) *QueryPenetrationTestDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetTestContent(v string) *QueryPenetrationTestDetailResponseBodyData {
	s.TestContent = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetTestEndDate(v int64) *QueryPenetrationTestDetailResponseBodyData {
	s.TestEndDate = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetTestSourceIps(v string) *QueryPenetrationTestDetailResponseBodyData {
	s.TestSourceIps = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetTestStartDate(v int64) *QueryPenetrationTestDetailResponseBodyData {
	s.TestStartDate = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetTestSteps(v string) *QueryPenetrationTestDetailResponseBodyData {
	s.TestSteps = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetThirdpartyEmail(v string) *QueryPenetrationTestDetailResponseBodyData {
	s.ThirdpartyEmail = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetThirdpartyMobile(v string) *QueryPenetrationTestDetailResponseBodyData {
	s.ThirdpartyMobile = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) SetThirdpartyTester(v string) *QueryPenetrationTestDetailResponseBodyData {
	s.ThirdpartyTester = &v
	return s
}

func (s *QueryPenetrationTestDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
