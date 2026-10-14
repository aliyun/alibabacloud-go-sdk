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
	// The business error code. The value "OK" is returned when the request succeeds.
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
	// This field is returned only when Status is succeeded.
	//
	// example:
	//
	// tamper
	DetectMode *string `json:"DetectMode,omitempty" xml:"DetectMode,omitempty"`
	// The HTTP status code. The value 200 is returned when the request succeeds.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The list of detection result labels. This field is returned only when Status is succeeded and AIGC detection is included. The label values and confidence ranges vary depending on DetectMode. In aigc mode, AIGC generation identification labels are returned (see the table below). In tamper mode, aigc/aips labels are returned (confidence range: 0 to 1).
	Labels []*GetImageDetectionTaskResultResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The additional information. The value "success" is returned when the request succeeds. The error message is returned when the task fails. This field is returned only when Status is failed.
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
	// - pending: waiting.
	//
	// - running: in progress.
	//
	// - succeeded: completed.
	//
	// - failed: failed.
	//
	// example:
	//
	// succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The tamper detection result.
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
	// The detection label. When DetectMode is aigc, valid values are: risk_aigc (the image is suspected to be AIGC-generated), risk_fake (the image is suspected to be a composite image), risk_edit (the image is suspected to be AI-edited). When DetectMode is tamper, valid values are: aigc (AIGC detection score), aips (AI post-processing detection score).
	//
	// example:
	//
	// risk_aigc
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
	// The AI visible watermark. This field is returned only in large model mode. It is empty in small model mode.
	//
	// example:
	//
	// 是
	AiWatermark *string `json:"AiWatermark,omitempty" xml:"AiWatermark,omitempty"`
	// The basic algorithm detection results, including multi-dimensional detection scores and risk labels.
	BaseResults *GetImageDetectionTaskResultResponseBodyTamperBaseResults `json:"BaseResults,omitempty" xml:"BaseResults,omitempty" type:"Struct"`
	// The final conclusion.
	//
	// example:
	//
	// Metadata data analytics result verified as high risk
	FinalConclusion *string `json:"FinalConclusion,omitempty" xml:"FinalConclusion,omitempty"`
	// The comprehensive analysis result from the large language model.
	//
	// example:
	//
	// The image has been detected with traces of AI generation. The AIGC detection score is 0.95. We recommend further verifying the image source
	LlmResult *string `json:"LlmResult,omitempty" xml:"LlmResult,omitempty"`
	// The metadata AI flag. This field is always returned.
	//
	// example:
	//
	// 高风险
	MetadataAiFlag *string `json:"MetadataAiFlag,omitempty" xml:"MetadataAiFlag,omitempty"`
	// The source of the metadata AI flag. This field is returned only in large model mode. It is empty in small model mode.
	//
	// example:
	//
	// 通义
	MetadataAiSource *string `json:"MetadataAiSource,omitempty" xml:"MetadataAiSource,omitempty"`
	// The determination result. This field is always returned.
	//
	// example:
	//
	// 高风险
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetImageDetectionTaskResultResponseBodyTamper) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyTamper) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) GetAiWatermark() *string {
	return s.AiWatermark
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) GetBaseResults() *GetImageDetectionTaskResultResponseBodyTamperBaseResults {
	return s.BaseResults
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) GetFinalConclusion() *string {
	return s.FinalConclusion
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) GetLlmResult() *string {
	return s.LlmResult
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) GetMetadataAiFlag() *string {
	return s.MetadataAiFlag
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) GetMetadataAiSource() *string {
	return s.MetadataAiSource
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) SetAiWatermark(v string) *GetImageDetectionTaskResultResponseBodyTamper {
	s.AiWatermark = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) SetBaseResults(v *GetImageDetectionTaskResultResponseBodyTamperBaseResults) *GetImageDetectionTaskResultResponseBodyTamper {
	s.BaseResults = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) SetFinalConclusion(v string) *GetImageDetectionTaskResultResponseBodyTamper {
	s.FinalConclusion = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) SetLlmResult(v string) *GetImageDetectionTaskResultResponseBodyTamper {
	s.LlmResult = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) SetMetadataAiFlag(v string) *GetImageDetectionTaskResultResponseBodyTamper {
	s.MetadataAiFlag = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) SetMetadataAiSource(v string) *GetImageDetectionTaskResultResponseBodyTamper {
	s.MetadataAiSource = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamper) SetRiskLevel(v string) *GetImageDetectionTaskResultResponseBodyTamper {
	s.RiskLevel = &v
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
	// The AIGC detection result.
	Aigc *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc `json:"Aigc,omitempty" xml:"Aigc,omitempty" type:"Struct"`
	// The AI post-processing detection score.
	Aips *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips `json:"Aips,omitempty" xml:"Aips,omitempty" type:"Struct"`
	// The image type.
	//
	// example:
	//
	// natural
	ImgType *string `json:"ImgType,omitempty" xml:"ImgType,omitempty"`
	// The PS tamper localization result.
	PsLoc *GetImageDetectionTaskResultResponseBodyTamperBaseResultsPsLoc `json:"PsLoc,omitempty" xml:"PsLoc,omitempty" type:"Struct"`
	// The recapture/copy detection result.
	Recap *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap `json:"Recap,omitempty" xml:"Recap,omitempty" type:"Struct"`
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

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResults) GetRecap() *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap {
	return s.Recap
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

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResults) SetRecap(v *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap) *GetImageDetectionTaskResultResponseBodyTamperBaseResults {
	s.Recap = v
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
	if s.Recap != nil {
		if err := s.Recap.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc struct {
	// The AIGC detection result description.
	//
	// example:
	//
	// Suspected AI-generated image
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The AIGC dimension risk label.
	//
	// example:
	//
	// 高风险
	Risk *string `json:"Risk,omitempty" xml:"Risk,omitempty"`
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

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc) GetRisk() *string {
	return s.Risk
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc) GetScore() *float32 {
	return s.Score
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc) SetDesc(v string) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc {
	s.Desc = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc) SetRisk(v string) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAigc {
	s.Risk = &v
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
	// The AI post-processing detection result description.
	//
	// example:
	//
	// No AI post-processing traces detected
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The local PS/AI editing dimension risk label.
	//
	// example:
	//
	// 高风险
	Risk *string `json:"Risk,omitempty" xml:"Risk,omitempty"`
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

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips) GetRisk() *string {
	return s.Risk
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips) GetScore() *float32 {
	return s.Score
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips) SetDesc(v string) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips {
	s.Desc = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips) SetRisk(v string) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsAips {
	s.Risk = &v
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
	// The PS tamper localization result description.
	//
	// example:
	//
	// 1 suspected PS-modified region detected
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
	// The bounding box coordinates of the tampered region, in the format `[x1, y1, x2, y2]`.
	Bbox []*float32 `json:"Bbox,omitempty" xml:"Bbox,omitempty" type:"Repeated"`
	// The tamper confidence of the region. Value range: 0 to 1.
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

type GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap struct {
	// The recapture/copy detection result description.
	//
	// example:
	//
	// Recapture detection score
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The recapture/copy dimension risk label.
	//
	// example:
	//
	// 高风险
	Risk *string `json:"Risk,omitempty" xml:"Risk,omitempty"`
	// The recapture/copy detection score. Value range: 0 to 1. A higher value indicates a greater likelihood of recapture or copying.
	//
	// example:
	//
	// 0.9998
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap) GetDesc() *string {
	return s.Desc
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap) GetRisk() *string {
	return s.Risk
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap) GetScore() *float32 {
	return s.Score
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap) SetDesc(v string) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap {
	s.Desc = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap) SetRisk(v string) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap {
	s.Risk = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap) SetScore(v float32) *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap {
	s.Score = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyTamperBaseResultsRecap) Validate() error {
	return dara.Validate(s)
}
