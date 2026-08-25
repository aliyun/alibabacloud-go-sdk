// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTestResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetImageTestResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetImageTestResultResponseBody
	GetSuccess() *bool
	SetTestResult(v *GetImageTestResultResponseBodyTestResult) *GetImageTestResultResponseBody
	GetTestResult() *GetImageTestResultResponseBodyTestResult
}

type GetImageTestResultResponseBody struct {
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of the image test result.
	TestResult *GetImageTestResultResponseBodyTestResult `json:"TestResult,omitempty" xml:"TestResult,omitempty" type:"Struct"`
}

func (s GetImageTestResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageTestResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageTestResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageTestResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageTestResultResponseBody) GetTestResult() *GetImageTestResultResponseBodyTestResult {
	return s.TestResult
}

func (s *GetImageTestResultResponseBody) SetRequestId(v string) *GetImageTestResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageTestResultResponseBody) SetSuccess(v bool) *GetImageTestResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageTestResultResponseBody) SetTestResult(v *GetImageTestResultResponseBodyTestResult) *GetImageTestResultResponseBody {
	s.TestResult = v
	return s
}

func (s *GetImageTestResultResponseBody) Validate() error {
	if s.TestResult != nil {
		if err := s.TestResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageTestResultResponseBodyTestResult struct {
	// The image ID.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The test result message.
	//
	// example:
	//
	// test finished
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The operation time, represented as a 64-bit timestamp.
	//
	// example:
	//
	// 1727055811000
	OperateTime *int64 `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	// The process ID.
	//
	// example:
	//
	// 582d4896-d224-413b-b883-239eeebe0bc5
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The image publish status. Valid values:
	//
	// - Untest: Not tested.
	//
	// - Testing: Testing in progress.
	//
	// - TestFailed: Test failed.
	//
	// - Unpublished: Not published.
	//
	// - Publishing: Publishing in progress.
	//
	// - Published: Published.
	//
	// - PublishFailed: Publish failed.
	//
	// - Building: Building in progress.
	//
	// - BuildSuccess: Build succeeded.
	//
	// - BuildFailed: Build failed.
	//
	// - Accelerating: Acceleration in progress.
	//
	// - AccelerateSuccess: Acceleration succeeded.
	//
	// - AccelerateFailed: Acceleration failed.
	//
	// example:
	//
	// Unpublished
	PublishStage *string `json:"PublishStage,omitempty" xml:"PublishStage,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// Serverless_res_group_****
	ResourceGroupId *int64 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The test process status. Valid values:
	//
	// - running: Running.
	//
	// - completed: Completed.
	//
	// - failed: Failed.
	//
	// - cancelled: Cancelled.
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetImageTestResultResponseBodyTestResult) String() string {
	return dara.Prettify(s)
}

func (s GetImageTestResultResponseBodyTestResult) GoString() string {
	return s.String()
}

func (s *GetImageTestResultResponseBodyTestResult) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageTestResultResponseBodyTestResult) GetMessage() *string {
	return s.Message
}

func (s *GetImageTestResultResponseBodyTestResult) GetOperateTime() *int64 {
	return s.OperateTime
}

func (s *GetImageTestResultResponseBodyTestResult) GetProcessId() *string {
	return s.ProcessId
}

func (s *GetImageTestResultResponseBodyTestResult) GetPublishStage() *string {
	return s.PublishStage
}

func (s *GetImageTestResultResponseBodyTestResult) GetResourceGroupId() *int64 {
	return s.ResourceGroupId
}

func (s *GetImageTestResultResponseBodyTestResult) GetStatus() *string {
	return s.Status
}

func (s *GetImageTestResultResponseBodyTestResult) SetImageId(v string) *GetImageTestResultResponseBodyTestResult {
	s.ImageId = &v
	return s
}

func (s *GetImageTestResultResponseBodyTestResult) SetMessage(v string) *GetImageTestResultResponseBodyTestResult {
	s.Message = &v
	return s
}

func (s *GetImageTestResultResponseBodyTestResult) SetOperateTime(v int64) *GetImageTestResultResponseBodyTestResult {
	s.OperateTime = &v
	return s
}

func (s *GetImageTestResultResponseBodyTestResult) SetProcessId(v string) *GetImageTestResultResponseBodyTestResult {
	s.ProcessId = &v
	return s
}

func (s *GetImageTestResultResponseBodyTestResult) SetPublishStage(v string) *GetImageTestResultResponseBodyTestResult {
	s.PublishStage = &v
	return s
}

func (s *GetImageTestResultResponseBodyTestResult) SetResourceGroupId(v int64) *GetImageTestResultResponseBodyTestResult {
	s.ResourceGroupId = &v
	return s
}

func (s *GetImageTestResultResponseBodyTestResult) SetStatus(v string) *GetImageTestResultResponseBodyTestResult {
	s.Status = &v
	return s
}

func (s *GetImageTestResultResponseBodyTestResult) Validate() error {
	return dara.Validate(s)
}
