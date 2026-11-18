// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIMediaAuditJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaAuditJob(v *GetAIMediaAuditJobResponseBodyMediaAuditJob) *GetAIMediaAuditJobResponseBody
	GetMediaAuditJob() *GetAIMediaAuditJobResponseBodyMediaAuditJob
	SetRequestId(v string) *GetAIMediaAuditJobResponseBody
	GetRequestId() *string
}

type GetAIMediaAuditJobResponseBody struct {
	// The automated review job information.
	MediaAuditJob *GetAIMediaAuditJobResponseBodyMediaAuditJob `json:"MediaAuditJob,omitempty" xml:"MediaAuditJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EAA3E96A-02E2-41*****85-08E1D568ED3A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAIMediaAuditJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBody) GetMediaAuditJob() *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	return s.MediaAuditJob
}

func (s *GetAIMediaAuditJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAIMediaAuditJobResponseBody) SetMediaAuditJob(v *GetAIMediaAuditJobResponseBodyMediaAuditJob) *GetAIMediaAuditJobResponseBody {
	s.MediaAuditJob = v
	return s
}

func (s *GetAIMediaAuditJobResponseBody) SetRequestId(v string) *GetAIMediaAuditJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBody) Validate() error {
	if s.MediaAuditJob != nil {
		if err := s.MediaAuditJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAIMediaAuditJobResponseBodyMediaAuditJob struct {
	// The job error code. Check this field when the job status is fail.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The end time of the job. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// example:
	//
	// 2017-01-11T13:00:00Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The start time of the job. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The job result data.
	Data *GetAIMediaAuditJobResponseBodyMediaAuditJobData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The job ID.
	//
	// example:
	//
	// bdbc266af6894*****943a70176d92e9
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The video ID.
	//
	// example:
	//
	// fe028d09441d*****d1afffb138cd7e
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The job error message. Check this field when the job status is fail.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The job status. Valid values:
	//
	// - **success**: Succeeded.
	//
	// - **fail**: Failed.
	//
	// - **init**: Initialized.
	//
	// - **processing**: Processing.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The job type. Currently, only automated review is supported.
	//
	// example:
	//
	// AIMediaAudit
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJob) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJob) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) GetCode() *string {
	return s.Code
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) GetData() *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	return s.Data
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) GetJobId() *string {
	return s.JobId
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) GetMediaId() *string {
	return s.MediaId
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) GetMessage() *string {
	return s.Message
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) GetStatus() *string {
	return s.Status
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) GetType() *string {
	return s.Type
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetCode(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.Code = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetCompleteTime(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.CompleteTime = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetCreationTime(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.CreationTime = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetData(v *GetAIMediaAuditJobResponseBodyMediaAuditJobData) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.Data = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetJobId(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.JobId = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetMediaId(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.MediaId = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetMessage(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.Message = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetStatus(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.Status = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetType(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.Type = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobData struct {
	// The reviewed content that contains violations. Multiple values are separated by commas (,). Valid values:
	//
	// - **video**: Video.
	//
	// - **image-cover**: Thumbnail.
	//
	// - **text-title**: Title.
	//
	// example:
	//
	// video
	AbnormalModules *string `json:"AbnormalModules,omitempty" xml:"AbnormalModules,omitempty"`
	// The audio review result.
	AudioResult []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Repeated"`
	// The image review result.
	ImageResult []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult `json:"ImageResult,omitempty" xml:"ImageResult,omitempty" type:"Repeated"`
	// The review result category. Multiple values are separated by commas (,). Valid values:
	//
	// - **porn**: Pornographic content.
	//
	// - **terrorism**: Terrorist and political content.
	//
	// - **ad**: Image and text violation.
	//
	// - **live**: Undesirable scene.
	//
	// - **logo**: Image logo.
	//
	// - **audio**: Audio anti-spam.
	//
	// - **normal**: Normal.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The review suggestion. Valid values:
	//
	// - **block**: Violation.
	//
	// - **review**: Suspected violation.
	//
	// - **pass**: Passed.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The text review result.
	TextResult []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult `json:"TextResult,omitempty" xml:"TextResult,omitempty" type:"Repeated"`
	// The video review result.
	VideoResult *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult `json:"VideoResult,omitempty" xml:"VideoResult,omitempty" type:"Struct"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobData) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobData) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) GetAbnormalModules() *string {
	return s.AbnormalModules
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) GetAudioResult() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult {
	return s.AudioResult
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) GetImageResult() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult {
	return s.ImageResult
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) GetTextResult() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	return s.TextResult
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) GetVideoResult() *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	return s.VideoResult
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetAbnormalModules(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.AbnormalModules = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetAudioResult(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.AudioResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetImageResult(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.ImageResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetTextResult(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.TextResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetVideoResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.VideoResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) Validate() error {
	if s.AudioResult != nil {
		for _, item := range s.AudioResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageResult != nil {
		for _, item := range s.ImageResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TextResult != nil {
		for _, item := range s.TextResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoResult != nil {
		if err := s.VideoResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult struct {
	// The review result category. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **spam**: Spam content.
	//
	// - **ad**: Ad.
	//
	// - **politics**: Political content.
	//
	// - **terrorism**: Terrorist content.
	//
	// - **abuse**: Abuse.
	//
	// - **porn**: Pornographic content.
	//
	// - **flood**: Flooding.
	//
	// - **contraband**: Contraband.
	//
	// - **meaningless**: Meaningless.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The review scene. Fixed value: **antispam**.
	//
	// example:
	//
	// antispam
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The result score.
	//
	// example:
	//
	// 99.91
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The review suggestion. Valid values:
	//
	// - **block**: Violation.
	//
	// - **review**: Suspected violation.
	//
	// - **pass**: Passed.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) GetScene() *string {
	return s.Scene
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) GetScore() *string {
	return s.Score
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) SetScene(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult {
	s.Scene = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult struct {
	// The image review result category. Multiple values are separated by commas (,). Valid values:
	//
	// - **porn**: Pornographic content.
	//
	// - **terrorism**: Terrorist and political content.
	//
	// - **ad**: Image and text violation.
	//
	// - **live**: Undesirable scene.
	//
	// - **logo**: Image logo.
	//
	// - **normal**: Normal.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The image review result details.
	Result []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The review suggestion. Valid values:
	//
	// - **block**: Violation.
	//
	// - **review**: Suspected violation.
	//
	// - **pass**: Passed.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The image category. Valid value: **cover*	- (thumbnail).
	//
	// example:
	//
	// cover
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the image.
	//
	// example:
	//
	// http://www.test.com/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) GetResult() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult {
	return s.Result
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) GetType() *string {
	return s.Type
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) GetUrl() *string {
	return s.Url
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) SetResult(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult {
	s.Result = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) SetType(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult {
	s.Type = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult {
	s.Url = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult struct {
	// The review result category.
	//
	// When the scene is **porn**, valid values:
	//
	// - **porn**: Pornographic.
	//
	// - **sexy**: Sexy.
	//
	// - **normal**: Normal.
	//
	// When the scene is **terrorism**, valid values:
	//
	// - **normal**: Normal.
	//
	// - **bloody**: Bloody.
	//
	// - **explosion**: Explosion and smoke.
	//
	// - **outfit**: Special outfit.
	//
	// - **logo**: Special logo.
	//
	// - **weapon**: Weapon.
	//
	// - **politics**: Political content.
	//
	// - **violence**: Violence.
	//
	// - **crowd**: Crowd.
	//
	// - **parade**: Parade.
	//
	// - **carcrash**: Car crash.
	//
	// - **flag**: Flag.
	//
	// - **location**: Landmark.
	//
	// - **others**: Others.
	//
	// When the scene is **ad**, valid values:
	//
	// - **normal**: Normal.
	//
	// - **ad**: Other ads.
	//
	// - **politics**: Text containing political content.
	//
	// - **porn**: Text containing pornographic content.
	//
	// - **abuse**: Text containing abusive content.
	//
	// - **terrorism**: Text containing terrorist content.
	//
	// - **contraband**: Text containing contraband content.
	//
	// - **spam**: Text containing other spam content.
	//
	// - **npx**: Nuisance ad.
	//
	// - **qrcode**: QR code included.
	//
	// - **programCode**: Mini program code included.
	//
	// When the scene is **live**, valid values:
	//
	// - **normal**: Normal.
	//
	// - **meaningless**: No content in the image, such as a black or white screen.
	//
	// - **PIP**: Picture-in-Picture (PiP).
	//
	// - **smoking**: Smoking.
	//
	// - **drivelive**: In-car live streaming.
	//
	// When the scene is **logo**, valid values:
	//
	// - **normal**: Normal.
	//
	// - **TV**: Regulated logo included.
	//
	// - **trademark**: Trademark included.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The review scene. Valid values:
	//
	// - **porn**: Pornography detection.
	//
	// - **terrorism**: Terrorist and political content detection.
	//
	// - **ad**: Image and text violation detection.
	//
	// - **live**: Undesirable scene detection.
	//
	// - **logo**: Image logo detection.
	//
	// example:
	//
	// porn
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The score of the image that hits the label. Valid values: `[0, 100]`. The score indicates the probability that the image belongs to the corresponding label category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 0
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The review suggestion. Valid values:
	//
	// - **block**: Violation.
	//
	// - **review**: Suspected violation.
	//
	// - **pass**: Passed.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) GetScene() *string {
	return s.Scene
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) GetScore() *string {
	return s.Score
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) SetScene(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult {
	s.Scene = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult struct {
	// The text content.
	//
	// example:
	//
	// Test
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The review result category. Valid values:
	//
	// - **spam**: Spam content.
	//
	// - **ad**: Ad.
	//
	// - **abuse**: Abuse.
	//
	// - **flood**: Flooding.
	//
	// - **contraband**: Contraband.
	//
	// - **meaningless**: Meaningless.
	//
	// - **normal**: Normal.
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The review scene. Fixed value: **antispam**.
	//
	// example:
	//
	// antispam
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The score of the image that hits the label. Valid values: `[0, 100]`. The score indicates the probability that the image belongs to the corresponding label category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The review suggestion. Valid values:
	//
	// - **block**: Violation.
	//
	// - **review**: Suspected violation.
	//
	// - **pass**: Passed.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The text categorization. Valid value: **title*	- (title).
	//
	// example:
	//
	// title
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) GetContent() *string {
	return s.Content
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) GetScene() *string {
	return s.Scene
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) GetScore() *string {
	return s.Score
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) GetType() *string {
	return s.Type
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) SetContent(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	s.Content = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) SetScene(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	s.Scene = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) SetType(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	s.Type = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult struct {
	// The ad review result.
	AdResult *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult `json:"AdResult,omitempty" xml:"AdResult,omitempty" type:"Struct"`
	// The GreenEnhanced moderation result.
	GreenEnhancedResult *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult `json:"GreenEnhancedResult,omitempty" xml:"GreenEnhancedResult,omitempty" type:"Struct"`
	// The review result category. Valid values:
	//
	// - **porn**: Pornographic content.
	//
	// - **terrorism**: Terrorist and political content.
	//
	// - **ad**: Image and text violation.
	//
	// - **live**: Undesirable scene.
	//
	// - **logo**: Image logo.
	//
	// - **normal**: Normal.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The harmful content moderation result.
	LiveResult *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult `json:"LiveResult,omitempty" xml:"LiveResult,omitempty" type:"Struct"`
	// The logo moderation result.
	LogoResult *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult `json:"LogoResult,omitempty" xml:"LogoResult,omitempty" type:"Struct"`
	// The pornography detection result.
	PornResult *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult `json:"PornResult,omitempty" xml:"PornResult,omitempty" type:"Struct"`
	// The video review suggestion. Valid values:
	//
	// - **block**: Violation.
	//
	// - **review**: Suspected violation.
	//
	// - **pass**: Passed.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The terrorist and political content review result.
	TerrorismResult *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult `json:"TerrorismResult,omitempty" xml:"TerrorismResult,omitempty" type:"Struct"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) GetAdResult() *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	return s.AdResult
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) GetGreenEnhancedResult() *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult {
	return s.GreenEnhancedResult
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) GetLiveResult() *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	return s.LiveResult
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) GetLogoResult() *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	return s.LogoResult
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) GetPornResult() *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	return s.PornResult
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) GetTerrorismResult() *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	return s.TerrorismResult
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetAdResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.AdResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetGreenEnhancedResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.GreenEnhancedResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetLiveResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.LiveResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetLogoResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.LogoResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetPornResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.PornResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetTerrorismResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.TerrorismResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) Validate() error {
	if s.AdResult != nil {
		if err := s.AdResult.Validate(); err != nil {
			return err
		}
	}
	if s.GreenEnhancedResult != nil {
		if err := s.GreenEnhancedResult.Validate(); err != nil {
			return err
		}
	}
	if s.LiveResult != nil {
		if err := s.LiveResult.Validate(); err != nil {
			return err
		}
	}
	if s.LogoResult != nil {
		if err := s.LogoResult.Validate(); err != nil {
			return err
		}
	}
	if s.PornResult != nil {
		if err := s.PornResult.Validate(); err != nil {
			return err
		}
	}
	if s.TerrorismResult != nil {
		if err := s.TerrorismResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult struct {
	// The average score of the ad review result.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The categories of the review results and the number of video snapshots for each category.
	CounterList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The ad review result category. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **ad**: Other ads.
	//
	// - **politics**: Text containing political content.
	//
	// - **porn**: Text containing pornographic content.
	//
	// - **abuse**: Text containing abusive content.
	//
	// - **terrorism**: Text containing terrorist content.
	//
	// - **contraband**: Text containing contraband content.
	//
	// - **spam**: Text containing other spam content.
	//
	// - **npx**: Nuisance ad.
	//
	// - **qrcode**: QR code included.
	//
	// - **programCode**: Mini program code included.
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest review score.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The review suggestion. Valid values:
	//
	// - **block**: Violation.
	//
	// - **review**: Suspected violation.
	//
	// - **pass**: Passed.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the video snapshots with the highest scores for the hit label.
	TopList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) GetAverageScore() *string {
	return s.AverageScore
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) GetCounterList() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList {
	return s.CounterList
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) GetMaxScore() *string {
	return s.MaxScore
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) GetTopList() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList {
	return s.TopList
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) SetAverageScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	s.AverageScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) SetCounterList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	s.CounterList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) SetMaxScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	s.MaxScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) SetTopList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	s.TopList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) Validate() error {
	if s.CounterList != nil {
		for _, item := range s.CounterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopList != nil {
		for _, item := range s.TopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ad review result category. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **ad**: Other ads.
	//
	// - **politics**: Text containing political content.
	//
	// - **porn**: Text containing pornographic content.
	//
	// - **abuse**: Text containing abusive content.
	//
	// - **terrorism**: Text containing terrorist content.
	//
	// - **contraband**: Text containing contraband content.
	//
	// - **spam**: Text containing other spam content.
	//
	// - **npx**: Nuisance ad.
	//
	// - **qrcode**: QR code included.
	//
	// - **programCode**: Mini program code included.
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) GetCount() *int32 {
	return s.Count
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) SetCount(v int32) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList {
	s.Count = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList struct {
	// The ad review result category. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **ad**: Other ads.
	//
	// - **politics**: Text containing political content.
	//
	// - **porn**: Text containing pornographic content.
	//
	// - **abuse**: Text containing abusive content.
	//
	// - **terrorism**: Text containing terrorist content.
	//
	// - **contraband**: Text containing contraband content.
	//
	// - **spam**: Text containing other spam content.
	//
	// - **npx**: Nuisance ad.
	//
	// - **qrcode**: QR code included.
	//
	// - **programCode**: Mini program code included.
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the video snapshot that hits the label.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 500
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	//
	// example:
	//
	// http://temp-testbucket.oss-cn-shanghai.aliyuncs.com/aivideocensor/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) GetScore() *string {
	return s.Score
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) GetUrl() *string {
	return s.Url
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) SetTimestamp(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList {
	s.Url = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult struct {
	// The average confidence score of the hit frames. This parameter is not returned if no frames are hit.
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The aggregated count of violation labels, including Label (the Green label) and Count (the number of frames that hit the label).
	CounterList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The union of hit Green native labels, separated by commas, such as pornographic_adultContent_tii. If no labels are hit, the value is normal.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest confidence score of the hit frames. This parameter is not returned if no frames are hit.
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The image moderation conclusion, which is mapped from frameResult.riskLevel: high to block, medium or low to review, and none to pass.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The details of the hit frames, sorted in descending order of confidence score.
	TopList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) GetAverageScore() *string {
	return s.AverageScore
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) GetCounterList() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList {
	return s.CounterList
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) GetMaxScore() *string {
	return s.MaxScore
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) GetTopList() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList {
	return s.TopList
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) SetAverageScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult {
	s.AverageScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) SetCounterList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult {
	s.CounterList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) SetMaxScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult {
	s.MaxScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) SetTopList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult {
	s.TopList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResult) Validate() error {
	if s.CounterList != nil {
		for _, item := range s.CounterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopList != nil {
		for _, item := range s.TopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList struct {
	// The number of video frames that correspond to the label.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The moderation result label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList) GetCount() *int32 {
	return s.Count
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList) SetCount(v int32) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList {
	s.Count = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultCounterList) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList struct {
	// The moderation category.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The confidence score of the video snapshot that hits the label.
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video snapshot in the video.
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList) GetScore() *string {
	return s.Score
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList) GetUrl() *string {
	return s.Url
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList) SetTimestamp(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList {
	s.Url = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultGreenEnhancedResultTopList) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult struct {
	// The average score of the moderation result.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The moderation categories and the number of video snapshots for each category.
	CounterList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The moderation category. Valid values:
	//
	// - **normal**: normal
	//
	// - **meaningless**: the image contains no content, such as a black or white screen
	//
	// - **PIP**: Picture-in-Picture (PiP)
	//
	// - **smoking**: smoking
	//
	// - **drivelive**: live streaming inside a vehicle
	//
	// example:
	//
	// smoking
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest review score.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The moderation suggestion. Valid values:
	//
	// - **block**: The content violates regulations.
	//
	// - **review**: The content is suspected of violating regulations.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the video snapshots with the highest scores for the hit label.
	TopList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) GetAverageScore() *string {
	return s.AverageScore
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) GetCounterList() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList {
	return s.CounterList
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) GetMaxScore() *string {
	return s.MaxScore
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) GetTopList() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList {
	return s.TopList
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) SetAverageScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	s.AverageScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) SetCounterList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	s.CounterList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) SetMaxScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	s.MaxScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) SetTopList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	s.TopList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) Validate() error {
	if s.CounterList != nil {
		for _, item := range s.CounterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopList != nil {
		for _, item := range s.TopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The moderation category. Valid values:
	//
	// - **normal**: normal
	//
	// - **meaningless**: the image contains no content, such as a black or white screen
	//
	// - **PIP**: Picture-in-Picture (PiP)
	//
	// - **smoking**: smoking
	//
	// - **drivelive**: live streaming inside a vehicle
	//
	// example:
	//
	// smoking
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) GetCount() *int32 {
	return s.Count
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) SetCount(v int32) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList {
	s.Count = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList struct {
	// The moderation category. Valid values:
	//
	// - **normal**: normal
	//
	// - **meaningless**: the image contains no content, such as a black or white screen
	//
	// - **PIP**: Picture-in-Picture (PiP)
	//
	// - **smoking**: smoking
	//
	// - **drivelive**: live streaming inside a vehicle
	//
	// example:
	//
	// smoking
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the video snapshot that hits the label.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 500
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	//
	// example:
	//
	// http://temp-testbucket.oss-cn-shanghai.aliyuncs.com/aivideocensor/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) GetScore() *string {
	return s.Score
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) GetUrl() *string {
	return s.Url
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) SetTimestamp(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList {
	s.Url = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult struct {
	// The average score of the video snapshots that hit the label.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The moderation categories and the number of video snapshots for each category.
	CounterList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The logo moderation category. Valid values:
	//
	// - **normal**: normal
	//
	// - **TV**: the image contains a regulated logo
	//
	// - **trademark**: the image contains a trademark
	//
	// example:
	//
	// TV
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest score of the video snapshots that hit the label.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The logo moderation suggestion. Valid values:
	//
	// - **block**: The content violates regulations.
	//
	// - **review**: The content is suspected of violating regulations.
	//
	// - **pass**: The content passes the moderation.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the video snapshots with the highest scores for the hit label.
	TopList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) GetAverageScore() *string {
	return s.AverageScore
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) GetCounterList() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList {
	return s.CounterList
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) GetMaxScore() *string {
	return s.MaxScore
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) GetTopList() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList {
	return s.TopList
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) SetAverageScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	s.AverageScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) SetCounterList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	s.CounterList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) SetMaxScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	s.MaxScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) SetTopList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	s.TopList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) Validate() error {
	if s.CounterList != nil {
		for _, item := range s.CounterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopList != nil {
		for _, item := range s.TopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The logo moderation category. Valid values:
	//
	// - **normal**: normal
	//
	// - **TV**: the image contains a regulated logo
	//
	// - **trademark**: the image contains a trademark
	//
	// example:
	//
	// TV
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) GetCount() *int32 {
	return s.Count
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) SetCount(v int32) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList {
	s.Count = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList struct {
	// The logo moderation category. Valid values:
	//
	// - **normal**: normal
	//
	// - **TV**: the image contains a regulated logo
	//
	// - **trademark**: the image contains a trademark
	//
	// example:
	//
	// TV
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the video snapshot that hits the label.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 5000
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	//
	// example:
	//
	// http://temp-testbucket.oss-cn-shanghai.aliyuncs.com/aivideocensor/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) GetScore() *string {
	return s.Score
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) GetUrl() *string {
	return s.Url
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) SetTimestamp(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList {
	s.Url = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult struct {
	// The average score of the video snapshots that hit the label. Valid values: `[0, 100]`. The precision is 10 decimal places. The score indicates the probability that the snapshot belongs to the corresponding label category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The categories of the review results and the number of video snapshots for each category.
	CounterList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The pornography detection result. Valid values:
	//
	// - **porn**: Pornographic.
	//
	// - **sexy**: Sexy.
	//
	// - **normal**: Normal.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest score of the video snapshots that hit the label. Valid values: `[0, 100]`. The precision is 10 decimal places. The score indicates the probability that the snapshot belongs to the corresponding label category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The pornography detection suggestion. Valid values:
	//
	// - **block**: Violation.
	//
	// - **review**: Suspected violation.
	//
	// - **pass**: Passed.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the video snapshots with the highest scores for the hit label.
	TopList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) GetAverageScore() *string {
	return s.AverageScore
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) GetCounterList() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList {
	return s.CounterList
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) GetMaxScore() *string {
	return s.MaxScore
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) GetTopList() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList {
	return s.TopList
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) SetAverageScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	s.AverageScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) SetCounterList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	s.CounterList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) SetMaxScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	s.MaxScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) SetTopList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	s.TopList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) Validate() error {
	if s.CounterList != nil {
		for _, item := range s.CounterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopList != nil {
		for _, item := range s.TopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 0
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The pornography detection result. Valid values:
	//
	// - **porn**: Pornographic.
	//
	// - **sexy**: Sexy.
	//
	// - **normal**: Normal.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) GetCount() *int32 {
	return s.Count
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) SetCount(v int32) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList {
	s.Count = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList struct {
	// The pornography detection result. Valid values:
	//
	// - **porn**: Pornographic.
	//
	// - **sexy**: Sexy.
	//
	// - **normal**: Normal.
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the video snapshot that hits the label. Valid values: `[0, 100]`. The precision is 10 decimal places. The score indicates the probability that the snapshot belongs to the corresponding label category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 3005
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	//
	// example:
	//
	// http://temp-testbucket.oss-cn-shanghai.aliyuncs.com/aivideocensor/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) GetScore() *string {
	return s.Score
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) GetUrl() *string {
	return s.Url
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) SetTimestamp(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList {
	s.Url = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult struct {
	// The average score of the video snapshots that hit the label. Valid values: `[0, 100]`. The precision is 10 decimal places. The score indicates the probability that the snapshot belongs to the corresponding label category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The categories of the terrorist and political content review results and the number of video snapshots for each category.
	CounterList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The terrorist and political content review result. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **bloody**: Bloody.
	//
	// - **explosion**: Explosion and smoke.
	//
	// - **outfit**: Special outfit.
	//
	// - **logo**: Special logo.
	//
	// - **weapon**: Weapon.
	//
	// - **politics**: Political content.
	//
	// - **violence**: Violence.
	//
	// - **crowd**: Crowd.
	//
	// - **parade**: Parade.
	//
	// - **carcrash**: Car crash.
	//
	// - **flag**: Flag.
	//
	// - **location**: Landmark.
	//
	// - **others**: Others.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest score of the video snapshots that hit the label. Valid values: `[0, 100]`. The precision is 10 decimal places. The score indicates the probability that the snapshot belongs to the corresponding label category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The terrorist and political content review suggestion. Valid values:
	//
	// - **block**: Violation.
	//
	// - **review**: Suspected violation.
	//
	// - **pass**: Passed.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the video snapshots with the highest scores for the hit label.
	TopList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) GetAverageScore() *string {
	return s.AverageScore
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) GetCounterList() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList {
	return s.CounterList
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) GetMaxScore() *string {
	return s.MaxScore
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) GetTopList() []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList {
	return s.TopList
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) SetAverageScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	s.AverageScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) SetCounterList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	s.CounterList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) SetMaxScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	s.MaxScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) SetTopList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	s.TopList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) Validate() error {
	if s.CounterList != nil {
		for _, item := range s.CounterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopList != nil {
		for _, item := range s.TopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 0
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The terrorist and political content review result. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **bloody**: Bloody.
	//
	// - **explosion**: Explosion and smoke.
	//
	// - **outfit**: Special outfit.
	//
	// - **logo**: Special logo.
	//
	// - **weapon**: Weapon.
	//
	// - **politics**: Political content.
	//
	// - **violence**: Violence.
	//
	// - **crowd**: Crowd.
	//
	// - **parade**: Parade.
	//
	// - **carcrash**: Car crash.
	//
	// - **flag**: Flag.
	//
	// - **location**: Landmark.
	//
	// - **others**: Others.
	//
	// example:
	//
	// terrorism
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) GetCount() *int32 {
	return s.Count
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) SetCount(v int32) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList {
	s.Count = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) Validate() error {
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList struct {
	// The terrorist and political content review result. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **bloody**: Bloody.
	//
	// - **explosion**: Explosion and smoke.
	//
	// - **outfit**: Special outfit.
	//
	// - **logo**: Special logo.
	//
	// - **weapon**: Weapon.
	//
	// - **politics**: Political content.
	//
	// - **violence**: Violence.
	//
	// - **crowd**: Crowd.
	//
	// - **parade**: Parade.
	//
	// - **carcrash**: Car crash.
	//
	// - **flag**: Flag.
	//
	// - **location**: Landmark.
	//
	// - **others**: Others.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the video snapshot that hits the label. Valid values: `[0, 100]`. The precision is 10 decimal places. The score indicates the probability that the snapshot belongs to the corresponding label category. A higher score indicates higher accuracy.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The position of the video snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 5
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	//
	// example:
	//
	// http://ali*****.com/aivideocensor/yytysursrutyrxuq/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) String() string {
	return dara.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) GetLabel() *string {
	return s.Label
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) GetScore() *string {
	return s.Score
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) GetUrl() *string {
	return s.Url
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) SetTimestamp(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList {
	s.Url = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) Validate() error {
	return dara.Validate(s)
}
