// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDetectionTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetImageDetectionTaskResultResponseBody
	GetCode() *string
	SetDetectMode(v string) *GetImageDetectionTaskResultResponseBody
	GetDetectMode() *string
	SetHttpStatusCode(v int32) *GetImageDetectionTaskResultResponseBody
	GetHttpStatusCode() *int32
	SetLabels(v []*GetImageDetectionTaskResultResponseBodyLabels) *GetImageDetectionTaskResultResponseBody
	GetLabels() []*GetImageDetectionTaskResultResponseBodyLabels
	SetMessage(v string) *GetImageDetectionTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageDetectionTaskResultResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetImageDetectionTaskResultResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetImageDetectionTaskResultResponseBody
	GetSuccess() *bool
	SetTamper(v *GetImageDetectionTaskResultResponseBodyTamper) *GetImageDetectionTaskResultResponseBody
	GetTamper() *GetImageDetectionTaskResultResponseBodyTamper
	SetTaskId(v string) *GetImageDetectionTaskResultResponseBody
	GetTaskId() *string
}

type GetImageDetectionTaskResultResponseBody struct {
	// The business error code. The value `"OK"` is returned when the request succeeds.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detection mode that was actually executed. Valid values:
	//
	// - aigc: AIGC only.
	//
	// - tamper: tamper detection.
	//
	// This parameter is returned only when `Status` is `succeeded`.
	//
	// example:
	//
	// tamper
	DetectMode *string `json:"DetectMode,omitempty" xml:"DetectMode,omitempty"`
	// The HTTP status code. The value `200` is returned when the request succeeds.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The list of AIGC detection result labels. This parameter is returned only when `Status` is `succeeded` and the task includes AIGC detection.
	Labels []*GetImageDetectionTaskResultResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The additional information. The value `"success"` is returned when the request succeeds. An error message is returned when the task fails. This parameter is returned only when `Status` is `failed`.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task status. Valid values:
	//
	// - `pending`: waiting.
	//
	// - `running`: in progress.
	//
	// - `succeeded`: completed.
	//
	// - `failed`: failed.
	//
	// example:
	//
	// succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The tamper detection results. This parameter is returned only when `DetectType` is `tamper` or `auto` (and the image is identified as a credential-type image).
	Tamper *GetImageDetectionTaskResultResponseBodyTamper `json:"Tamper,omitempty" xml:"Tamper,omitempty" type:"Struct"`
	// The task ID.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetImageDetectionTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetImageDetectionTaskResultResponseBody) GetDetectMode() *string {
	return s.DetectMode
}

func (s *GetImageDetectionTaskResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetImageDetectionTaskResultResponseBody) GetLabels() []*GetImageDetectionTaskResultResponseBodyLabels {
	return s.Labels
}

func (s *GetImageDetectionTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageDetectionTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageDetectionTaskResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetImageDetectionTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageDetectionTaskResultResponseBody) GetTamper() *GetImageDetectionTaskResultResponseBodyTamper {
	return s.Tamper
}

func (s *GetImageDetectionTaskResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageDetectionTaskResultResponseBody) SetCode(v string) *GetImageDetectionTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetDetectMode(v string) *GetImageDetectionTaskResultResponseBody {
	s.DetectMode = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetHttpStatusCode(v int32) *GetImageDetectionTaskResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetLabels(v []*GetImageDetectionTaskResultResponseBodyLabels) *GetImageDetectionTaskResultResponseBody {
	s.Labels = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetMessage(v string) *GetImageDetectionTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetRequestId(v string) *GetImageDetectionTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetStatus(v string) *GetImageDetectionTaskResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetSuccess(v bool) *GetImageDetectionTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetTamper(v *GetImageDetectionTaskResultResponseBodyTamper) *GetImageDetectionTaskResultResponseBody {
	s.Tamper = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetTaskId(v string) *GetImageDetectionTaskResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tamper != nil {
		if err := s.Tamper.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageDetectionTaskResultResponseBodyLabels struct {
	// The confidence level. Value range: 0 to 1.
	//
	// example:
	//
	// 0.9562
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The detection label. Valid values:
	//
	// - ai_generated: AI-generated.
	//
	// - non_ai_generated: not AI-generated.
	//
	// example:
	//
	// non_ai_generated
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetImageDetectionTaskResultResponseBodyLabels) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyLabels) GetConfidence() *float32 {
	return s.Confidence
}

func (s *GetImageDetectionTaskResultResponseBodyLabels) GetLabel() *string {
	return s.Label
}

func (s *GetImageDetectionTaskResultResponseBodyLabels) SetConfidence(v float32) *GetImageDetectionTaskResultResponseBodyLabels {
	s.Confidence = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyLabels) SetLabel(v string) *GetImageDetectionTaskResultResponseBodyLabels {
	s.Label = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyLabels) Validate() error {
	return dara.Validate(s)
}

type GetImageDetectionTaskResultResponseBodyTamper struct {
	// The basic detection results, including detection scores across multiple dimensions.
	BaseResults *GetImageDetectionTaskResultResponseBodyTamperBaseResults `json:"BaseResults,omitempty" xml:"BaseResults,omitempty" type:"Struct"`
	// The comprehensive analysis result from the large language model.
	//
	// example:
	//
	// 该图片经检测存在AI生成痕迹，AIGC检测得分0.95，建议进一步核实图片来源。
	LlmResult *string `json:"LlmResult,omitempty" xml:"LlmResult,omitempty"`
	// The risk code. A value of `"0"` indicates no risk.
	//
	// example:
	//
	// 0
	RiskCode *string `json:"RiskCode,omitempty" xml:"RiskCode,omitempty"`
	// The list of risk reasons. An empty array is returned when no risk is detected.
	RiskReasons []*string `json:"RiskReasons,omitempty" xml:"RiskReasons,omitempty" type:"Repeated"`
}

func (s GetImageDetectionTaskResultResponseBodyTamper) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyTamper) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) GetBaseResults() *GetImageDetectionTaskResultResponseBodyTamperBaseResults {
	return s.BaseResults
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) GetLlmResult() *string {
	return s.LlmResult
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) GetRiskCode() *string {
	return s.RiskCode
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) GetRiskReasons() []*string {
	return s.RiskReasons
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) SetBaseResults(v *GetImageDetectionTaskResultResponseBodyTamperBaseResults) *GetImageDetectionTaskResultResponseBodyTamper {
	s.BaseResults = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) SetLlmResult(v string) *GetImageDetectionTaskResultResponseBodyTamper {
	s.LlmResult = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) SetRiskCode(v string) *GetImageDetectionTaskResultResponseBodyTamper {
	s.RiskCode = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) SetRiskReasons(v []*string) *GetImageDetectionTaskResultResponseBodyTamper {
	s.RiskReasons = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) Validate() error {
	if s.BaseResults != nil {
		if err := s.BaseResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageDetectionTaskResultResponseBodyTamperBaseResults struct {
	// The AIGC detection score.
	Aigc *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc `json:"Aigc,omitempty" xml:"Aigc,omitempty" type:"Struct"`
	// The AI post-processing detection score.
	Aips *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips `json:"Aips,omitempty" xml:"Aips,omitempty" type:"Struct"`
	// The image type.
	//
	// example:
	//
	// natural
	ImgType *string `json:"ImgType,omitempty" xml:"ImgType,omitempty"`
	// The PS tamper localization results.
	PsLoc *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc `json:"PsLoc,omitempty" xml:"PsLoc,omitempty" type:"Struct"`
}

func (s GetImageDetectionTaskResultResponseBodyTamperBaseResults) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyTamperBaseResults) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResults) GetAigc() *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc {
	return s.Aigc
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResults) GetAips() *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips {
	return s.Aips
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResults) GetImgType() *string {
	return s.ImgType
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResults) GetPsLoc() *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc {
	return s.PsLoc
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResults) SetAigc(v *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc) *GetImageDetectionTaskResultResponseBodyTamperBaseResults {
	s.Aigc = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResults) SetAips(v *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips) *GetImageDetectionTaskResultResponseBodyTamperBaseResults {
	s.Aips = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResults) SetImgType(v string) *GetImageDetectionTaskResultResponseBodyTamperBaseResults {
	s.ImgType = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResults) SetPsLoc(v *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc) *GetImageDetectionTaskResultResponseBodyTamperBaseResults {
	s.PsLoc = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResults) Validate() error {
	if s.Aigc != nil {
		if err := s.Aigc.Validate(); err != nil {
			return err
		}
	}
	if s.Aips != nil {
		if err := s.Aips.Validate(); err != nil {
			return err
		}
	}
	if s.PsLoc != nil {
		if err := s.PsLoc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc struct {
	// The description of the AIGC detection result.
	//
	// example:
	//
	// 疑似AI生成图片
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The AIGC detection score. Value range: 0 to 1. A higher value indicates a greater likelihood of AI generation.
	//
	// example:
	//
	// 0.47
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc) GetDesc() *string {
	return s.Desc
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc) GetScore() *float32 {
	return s.Score
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc) SetDesc(v string) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc {
	s.Desc = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc) SetScore(v float32) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc {
	s.Score = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc) Validate() error {
	return dara.Validate(s)
}

type GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips struct {
	// The description of the AI post-processing detection result.
	//
	// example:
	//
	// 未发现AI后处理痕迹
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The AI post-processing detection score. Value range: 0 to 1. A higher value indicates a greater likelihood of AI post-processing.
	//
	// example:
	//
	// 0.12
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips) GetDesc() *string {
	return s.Desc
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips) GetScore() *float32 {
	return s.Score
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips) SetDesc(v string) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips {
	s.Desc = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips) SetScore(v float32) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips {
	s.Score = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips) Validate() error {
	return dara.Validate(s)
}

type GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc struct {
	// The description of the PS tamper localization result.
	//
	// example:
	//
	// 发现1处疑似PS修改区域
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The list of tampered regions.
	Items []*GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc) GetDesc() *string {
	return s.Desc
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc) GetItems() []*GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems {
	return s.Items
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc) SetDesc(v string) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc {
	s.Desc = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc) SetItems(v []*GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc {
	s.Items = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems struct {
	// The bounding box coordinates of the tampered region in the format `[x1, y1, x2, y2]`.
	Bbox []*float32 `json:"Bbox,omitempty" xml:"Bbox,omitempty" type:"Repeated"`
	// The tamper confidence level for the region. Value range: `0 to 1`.
	//
	// example:
	//
	// 0.67
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems) GetBbox() []*float32 {
	return s.Bbox
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems) GetScore() *float32 {
	return s.Score
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems) SetBbox(v []*float32) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems {
	s.Bbox = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems) SetScore(v float32) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems {
	s.Score = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLocItems) Validate() error {
	return dara.Validate(s)
}
