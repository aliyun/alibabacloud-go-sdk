// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDefectDetectionTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAbnormalityReason(v string) *GetImageDefectDetectionTaskResultResponseBody
	GetAbnormalityReason() *string
	SetAbnormalityScore(v float32) *GetImageDefectDetectionTaskResultResponseBody
	GetAbnormalityScore() *float32
	SetCode(v string) *GetImageDefectDetectionTaskResultResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetImageDefectDetectionTaskResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetImageDefectDetectionTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageDefectDetectionTaskResultResponseBody
	GetRequestId() *string
	SetResult(v string) *GetImageDefectDetectionTaskResultResponseBody
	GetResult() *string
	SetRiskCode(v string) *GetImageDefectDetectionTaskResultResponseBody
	GetRiskCode() *string
	SetStatus(v string) *GetImageDefectDetectionTaskResultResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetImageDefectDetectionTaskResultResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *GetImageDefectDetectionTaskResultResponseBody
	GetTaskId() *string
}

type GetImageDefectDetectionTaskResultResponseBody struct {
	// **The image abnormality reason.*	- This field is returned only when `Status` is `succeeded`.
	//
	// example:
	//
	// The main subject of the image is blurry
	AbnormalityReason *string `json:"AbnormalityReason,omitempty" xml:"AbnormalityReason,omitempty"`
	// **The image abnormality score.*	- The value ranges from 0 to 1. A higher value indicates a more severe image abnormality. This field is returned only when `Status` is `succeeded`.
	//
	// example:
	//
	// 0.9667
	AbnormalityScore *float32 `json:"AbnormalityScore,omitempty" xml:"AbnormalityScore,omitempty"`
	// **The business error code.*	- The value "OK" is returned for successful requests.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// **The HTTP status code.*	- The value 200 is returned for successful requests.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The failure reason. This field is returned only when `Status` is `failed`.
	//
	// example:
	//
	// poll timeout: no result within 30m0s
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The raw downstream detection result (JSON string). The `AbnormalityScore` and `AbnormalityReason` fields are parsed from this raw result. The raw result is retained for debugging and passing through newly added upstream fields.
	//
	// example:
	//
	// {"image_abnormality_score": 0.75, "image_abnormality_reason": "The main subject of the image is blurry"}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The risk code. **This field is deprecated and has no meaning.*	- Determine whether an abnormality exists based on the AbnormalityScore, AbnormalityReason, and Result fields.
	//
	// example:
	//
	// "0"
	RiskCode *string `json:"RiskCode,omitempty" xml:"RiskCode,omitempty"`
	// **The task status.*	- Valid values:
	//
	// - `running`: The task is running.
	//
	// - `succeeded`: The task is completed.
	//
	// - `failed`: The task failed.
	//
	// example:
	//
	// succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// **The task ID.**
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetImageDefectDetectionTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageDefectDetectionTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageDefectDetectionTaskResultResponseBody) GetAbnormalityReason() *string {
	return s.AbnormalityReason
}

func (s *GetImageDefectDetectionTaskResultResponseBody) GetAbnormalityScore() *float32 {
	return s.AbnormalityScore
}

func (s *GetImageDefectDetectionTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetImageDefectDetectionTaskResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetImageDefectDetectionTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageDefectDetectionTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageDefectDetectionTaskResultResponseBody) GetResult() *string {
	return s.Result
}

func (s *GetImageDefectDetectionTaskResultResponseBody) GetRiskCode() *string {
	return s.RiskCode
}

func (s *GetImageDefectDetectionTaskResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetImageDefectDetectionTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageDefectDetectionTaskResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageDefectDetectionTaskResultResponseBody) SetAbnormalityReason(v string) *GetImageDefectDetectionTaskResultResponseBody {
	s.AbnormalityReason = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponseBody) SetAbnormalityScore(v float32) *GetImageDefectDetectionTaskResultResponseBody {
	s.AbnormalityScore = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponseBody) SetCode(v string) *GetImageDefectDetectionTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponseBody) SetHttpStatusCode(v int32) *GetImageDefectDetectionTaskResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponseBody) SetMessage(v string) *GetImageDefectDetectionTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponseBody) SetRequestId(v string) *GetImageDefectDetectionTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponseBody) SetResult(v string) *GetImageDefectDetectionTaskResultResponseBody {
	s.Result = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponseBody) SetRiskCode(v string) *GetImageDefectDetectionTaskResultResponseBody {
	s.RiskCode = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponseBody) SetStatus(v string) *GetImageDefectDetectionTaskResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponseBody) SetSuccess(v bool) *GetImageDefectDetectionTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponseBody) SetTaskId(v string) *GetImageDefectDetectionTaskResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetImageDefectDetectionTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}
