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
	// The information about the intelligent review job.
	MediaAuditJob *GetAIMediaAuditJobResponseBodyMediaAuditJob `json:"MediaAuditJob,omitempty" xml:"MediaAuditJob,omitempty" type:"Struct"`
	// The ID of the request.
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
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJob struct {
	// The error code. This parameter is returned if the value of Status is fail.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the job is complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T13:00:00Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the job started to run. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The job result.
	Data *GetAIMediaAuditJobResponseBodyMediaAuditJobData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the job.
	//
	// example:
	//
	// bdbc266af6894*****943a70176d92e9
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the video.
	//
	// example:
	//
	// fe028d09441d*****d1afffb138cd7e
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The error message. This parameter is returned if the value of Status is fail.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- **success**: The job is successful.
	//
	// 	- **fail**: The job failed.
	//
	// 	- **init**: The job is being initialized.
	//
	// 	- **Processing**: The job is in progress.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the job. The value is AIMediaAudit.
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
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobData struct {
	// The content that violates the regulations. Separate multiple values with commas (,). Valid values:
	//
	// 	- **video**: the video.
	//
	// 	- **image-cover**: the cover.
	//
	// 	- **text-title**: the title.
	//
	// example:
	//
	// video
	AbnormalModules *string `json:"AbnormalModules,omitempty" xml:"AbnormalModules,omitempty"`
	// The results of audio review.
	AudioResult []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Repeated"`
	// The results of image review.
	ImageResult []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult `json:"ImageResult,omitempty" xml:"ImageResult,omitempty" type:"Repeated"`
	// The category of the review result. Multiple values are separated by commas (,). Valid values:
	//
	// 	- **porn**: pornographic content
	//
	// 	- **terrorism**: terrorist or politically sensitive content
	//
	// 	- **ad**: ad violation
	//
	// 	- **live**: undesirable scene
	//
	// 	- **logo**: logo
	//
	// 	- **audio**: audio anti-spam
	//
	// 	- **normal**: normal content
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// 	- **block**: The content violates the regulations.
	//
	// 	- **review**: The content may violate the regulations.
	//
	// 	- **pass**: The content passes the review.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The text moderation results.
	TextResult []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult `json:"TextResult,omitempty" xml:"TextResult,omitempty" type:"Repeated"`
	// The results of video review.
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
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult struct {
	// The category of the review result.
	//
	// 	- **normal**: normal content
	//
	// 	- **spam**: spam
	//
	// 	- **ad**: ads
	//
	// 	- **politics**: political content
	//
	// 	- **terrorism**: terrorist content
	//
	// 	- **abuse**: abuse
	//
	// 	- **porn**: pornographic content.
	//
	// 	- **flood**: excessive junk content
	//
	// 	- **contraband**: prohibited content
	//
	// 	- **meaningless**: meaningless content
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The review scenario. The value is **antispam**.
	//
	// example:
	//
	// antispam
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The score.
	//
	// example:
	//
	// 99.91
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// 	- **block**
	//
	// 	- **review**
	//
	// 	- **pass**
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
	// The categories of the image review results. Multiple values are separated by commas (,). Valid values:
	//
	// 	- **porn**: pornographic content
	//
	// 	- **terrorism**: terrorist or politically sensitive content
	//
	// 	- **ad**: ad violation
	//
	// 	- **live**: undesirable scene
	//
	// 	- **logo**: logo
	//
	// 	- **normal**: normal content
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Details of image review results.
	Result []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The recommendation for review results. Valid values:
	//
	// 	- **block**
	//
	// 	- **review**
	//
	// 	- **pass**
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The type of the image. Valid value: **cover**.
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
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult struct {
	// The category of the review result.
	//
	// Valid values if scene is **porn**:
	//
	// 	- **porn**
	//
	// 	- **sexy**
	//
	// 	- **normal**
	//
	// Valid values if scene is **terrorism**:
	//
	// 	- **normal**
	//
	// 	- **bloody**
	//
	// 	- **explosion**
	//
	// 	- **outfit**
	//
	// 	- **logo**
	//
	// 	- **weapon**
	//
	// 	- **politics**
	//
	// 	- **violence**
	//
	// 	- **crowd**
	//
	// 	- **parade**
	//
	// 	- **carcrash**
	//
	// 	- **flag**
	//
	// 	- **location**
	//
	// 	- **others**
	//
	// Valid values if scene is **ad**:
	//
	// 	- **normal**: normal content
	//
	// 	- **ad**: ads
	//
	// 	- **politics**: political content
	//
	// 	- **porn**: pornographic content
	//
	// 	- **abuse**: verbal abuse
	//
	// 	- **terrorism**: terrorist content
	//
	// 	- **contraband**: prohibited content
	//
	// 	- **spam**: spam content
	//
	// 	- **npx**: illegal ad
	//
	// 	- **qrcode**: QR code
	//
	// 	- **programCode**: mini program code
	//
	// Valid values if scene is **live**:
	//
	// 	- **normal**: normal content
	//
	// 	- **meaningless**: meaningless content, such as a black or white screen.
	//
	// 	- **PIP**: picture-in-picture
	//
	// 	- **smoking**: smoking
	//
	// 	- **drivelive**: live broadcasting in a running vehicle
	//
	// Valid values if scene is **logo**:
	//
	// 	- **normal**: normal content
	//
	// 	- **TV**: controlled TV station logo
	//
	// 	- **trademark**: trademark
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The review scenario. Valid values:
	//
	// 	- **porn**: pornographic content
	//
	// 	- **terrorism**: terrorist or politically sensitive content
	//
	// 	- **ad**: ad violation
	//
	// 	- **live**: undesirable scene
	//
	// 	- **logo**: logo
	//
	// example:
	//
	// porn
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The score of the image of the category that is indicated by Label. Valid values: `[0, 100]`. The score is representative of the confidence.
	//
	// example:
	//
	// 0
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// 	- **block**
	//
	// 	- **review**
	//
	// 	- **pass**
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
	// The category of the review result. Valid values:
	//
	// 	- **spam**: spam content
	//
	// 	- **ad**: ads
	//
	// 	- **abuse**: abuse
	//
	// 	- **flood**: excessive junk content
	//
	// 	- **contraband**: prohibited content
	//
	// 	- **meaningless**: meaningless content
	//
	// 	- **normal**: normal content
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The review scenario. The value is **antispam**.
	//
	// example:
	//
	// antispam
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The score of the image of the category that is indicated by Label. Valid values: `[0, 100]`. The score is representative of the confidence.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// 	- **block**
	//
	// 	- **review**
	//
	// 	- **pass**
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The type of the text. The value is **title**.
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
	// The results of ad review.
	AdResult *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult `json:"AdResult,omitempty" xml:"AdResult,omitempty" type:"Struct"`
	// The category of the review result. Valid values:
	//
	// 	- **porn**: pornographic content
	//
	// 	- **terrorism**: terrorist or politically sensitive content
	//
	// 	- **ad**: ad violation
	//
	// 	- **live**: undesirable scene
	//
	// 	- **logo**: logo
	//
	// 	- **normal**
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The results of undesired content review.
	LiveResult *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult `json:"LiveResult,omitempty" xml:"LiveResult,omitempty" type:"Struct"`
	// The results of logo review.
	LogoResult *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult `json:"LogoResult,omitempty" xml:"LogoResult,omitempty" type:"Struct"`
	// The results of pornographic content review.
	PornResult *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult `json:"PornResult,omitempty" xml:"PornResult,omitempty" type:"Struct"`
	// The recommendation for video review results. Valid values:
	//
	// 	- **block**
	//
	// 	- **review**
	//
	// 	- **pass**
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The results of terrorist content review.
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
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult struct {
	// The average score of the review results.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The number of snapshots of each category in the review result.
	CounterList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The categories of the ad review results. Valid values:
	//
	// 	- **normal**: normal content
	//
	// 	- **ad**: other ads
	//
	// 	- **politics**: political content
	//
	// 	- **porn**: pornographic content
	//
	// 	- **abuse**: abuse
	//
	// 	- **terrorism**: terrorist content
	//
	// 	- **contraband**: prohibited content
	//
	// 	- **spam**: spam content
	//
	// 	- **npx**: illegal ad
	//
	// 	- **qrcode**: QR code
	//
	// 	- **programCode**: mini program code
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
	// The recommendation for review results. Valid values:
	//
	// 	- **block**: The content violates the regulations.
	//
	// 	- **review**: The content may violate the regulations.
	//
	// 	- **pass**: The content passes the review.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the snapshot that has the highest score in the category.
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
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The categories of the ad review results. Valid values:
	//
	// 	- **normal**: normal content
	//
	// 	- **ad**: other ads
	//
	// 	- **politics**: political content
	//
	// 	- **porn**: pornographic content
	//
	// 	- **abuse**: abuse
	//
	// 	- **terrorism**: terrorist content
	//
	// 	- **contraband**: prohibited content
	//
	// 	- **spam**: spam content
	//
	// 	- **npx**: illegal ad
	//
	// 	- **qrcode**: QR code
	//
	// 	- **programCode**: mini program code
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
	// The categories of the ad review results. Valid values:
	//
	// 	- **normal**: normal content
	//
	// 	- **ad**: other ads
	//
	// 	- **politics**: political content
	//
	// 	- **porn**: pornographic content
	//
	// 	- **abuse**: abuse
	//
	// 	- **terrorism**: terrorist content
	//
	// 	- **contraband**: prohibited content
	//
	// 	- **spam**: spam content
	//
	// 	- **npx**: illegal ad
	//
	// 	- **qrcode**: QR code
	//
	// 	- **programCode**: mini program code
	//
	// example:
	//
	// ad
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the snapshot in the category that is indicated by Label.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The timestamp of the snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 500
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	//
	// example:
	//
	// http://temp-****bucket.oss-cn-shanghai.aliyuncs.com/aivideocensor/****.jpg
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

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult struct {
	// The average score of the review results.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The categories of the review results and the number of video snapshots in each category.
	CounterList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The category of the review result. Valid values:
	//
	// 	- **normal**: normal content
	//
	// 	- **meaningless**: meaningless content, such as a black or white screen.
	//
	// 	- **PIP**: picture-in-picture
	//
	// 	- **smoking**: smoking
	//
	// 	- **drivelive**: live broadcasting in a running vehicle
	//
	// example:
	//
	// live
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest review score.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// 	- **block**: The content violates the regulations.
	//
	// 	- **review**: The content may violate the regulations.
	//
	// 	- **pass**: The content passes the review.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the snapshot that has the highest score in the category.
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
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 0
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The category of the review result. Valid values:
	//
	// 	- **normal**: normal content
	//
	// 	- **meaningless**: meaningless content, such as a black or white screen.
	//
	// 	- **PIP**: picture-in-picture
	//
	// 	- **smoking**: smoking
	//
	// 	- **drivelive**: live broadcasting in a running vehicle
	//
	// example:
	//
	// live
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
	// The category of the review result. Valid values:
	//
	// 	- **normal**: normal content
	//
	// 	- **meaningless**: meaningless content, such as a black or white screen.
	//
	// 	- **PIP**: picture-in-picture
	//
	// 	- **smoking**: smoking
	//
	// 	- **drivelive**: live broadcasting in a running vehicle
	//
	// example:
	//
	// live
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the snapshot in the category that is indicated by Label.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The timestamp of the snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 500
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	//
	// example:
	//
	// http://temp-****bucket.oss-cn-shanghai.aliyuncs.com/aivideocensor/****.jpg
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
	// The average score of the snapshots in the category indicated by Label.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The categories of the review results and the number of video snapshots in each category.
	CounterList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The category of the review result. Valid values:
	//
	// 	- **normal**: normal content
	//
	// 	- **TV**: controlled TV station logo
	//
	// 	- **trademark**: trademark
	//
	// example:
	//
	// logo
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest score of the snapshot of the category that is indicated by Label.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// 	- **block**: The content violates the regulations.
	//
	// 	- **review**: The content may violate the regulations.
	//
	// 	- **pass**: The content passes the review.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the snapshot that has the highest score in the category.
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
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The category of the review result. Valid values:
	//
	// 	- **normal**: normal content
	//
	// 	- **TV**: controlled TV station logo
	//
	// 	- **trademark**: trademark
	//
	// example:
	//
	// logo
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
	// The category of the review result. Valid values:
	//
	// 	- **normal**: normal content
	//
	// 	- **TV**: controlled TV station logo
	//
	// 	- **trademark**: trademark
	//
	// example:
	//
	// logo
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the snapshot in the category that is indicated by Label.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The timestamp of the snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 5000
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	//
	// example:
	//
	// http://temp-****bucket.oss-cn-shanghai.aliyuncs.com/aivideocensor/****.jpg
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
	// The average score of the snapshots of the category that is indicated by Label. Valid values: `[0, 100]`. The value is accurate to 10 decimal places. The score is representative of the confidence.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The number of snapshots of each category in the review result.
	CounterList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The category of the review result. Valid values:
	//
	// 	- **porn**
	//
	// 	- **sexy**
	//
	// 	- **normal**
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest score of the snapshot of the category that is indicated by Label. Valid values: `[0, 100]`. The value is accurate to 10 decimal places. The score is representative of the confidence.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The recommendation for review results. Valid values:
	//
	// 	- **block**: The content violates the regulations.
	//
	// 	- **review**: The content may violate the regulations.
	//
	// 	- **pass**: The content passes the review.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the snapshot that has the highest score in the category.
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
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 0
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The results of pornographic content review. Valid values:
	//
	// 	- **porn**
	//
	// 	- **sexy**
	//
	// 	- **normal**
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
	// The results of pornographic content review. Valid values:
	//
	// 	- **porn**
	//
	// 	- **sexy**
	//
	// 	- **normal**
	//
	// example:
	//
	// porn
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the snapshot in the category that is indicated by Label. Valid values: `[0, 100]`. The value is accurate to 10 decimal places. The score is representative of the confidence.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The timestamp of the snapshot in the video. Unit: milliseconds.
	//
	// example:
	//
	// 3005
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The URL of the video snapshot.
	//
	// example:
	//
	// http://temp-****bucket.oss-cn-shanghai.aliyuncs.com/aivideocensor/****.jpg
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
	// The average score of the snapshots of the category that is indicated by Label. Valid values: `[0, 100]`. The value is accurate to 10 decimal places. The score is representative of the confidence.
	//
	// example:
	//
	// 100
	AverageScore *string `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	// The categories of terrorist content review results and the number of video snapshots in each category.
	CounterList []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	// The category of the review result. Valid values:
	//
	// 	- **normal**
	//
	// 	- **bloody**
	//
	// 	- **explosion**
	//
	// 	- **outfit**
	//
	// 	- **logo**
	//
	// 	- **weapon**
	//
	// 	- **politics**
	//
	// 	- **violence**
	//
	// 	- **crowd**
	//
	// 	- **parade**
	//
	// 	- **carcrash**
	//
	// 	- **flag**
	//
	// 	- **location**
	//
	// 	- **others**
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The highest score of the snapshot of the category that is indicated by Label. Valid values: `[0, 100]`. The value is accurate to 10 decimal places. The score is representative of the confidence.
	//
	// example:
	//
	// 100
	MaxScore *string `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	// The recommendation for terrorist content review results. Valid values:
	//
	// 	- **block**
	//
	// 	- **review**
	//
	// 	- **pass**
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The information about the snapshot that has the highest score in the category.
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
	return dara.Validate(s)
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList struct {
	// The number of video snapshots.
	//
	// example:
	//
	// 0
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The results of terrorist content review. Valid values:
	//
	// 	- **normal**
	//
	// 	- **bloody**
	//
	// 	- **explosion**
	//
	// 	- **outfit**
	//
	// 	- **logo**
	//
	// 	- **weapon**
	//
	// 	- **politics**
	//
	// 	- **violence**
	//
	// 	- **crowd**
	//
	// 	- **parade**
	//
	// 	- **carcrash**
	//
	// 	- **flag**
	//
	// 	- **location**
	//
	// 	- **others**
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
	// The results of terrorist content review. Valid values:
	//
	// 	- **normal**
	//
	// 	- **bloody**
	//
	// 	- **explosion**
	//
	// 	- **outfit**
	//
	// 	- **logo**
	//
	// 	- **weapon**
	//
	// 	- **politics**
	//
	// 	- **violence**
	//
	// 	- **crowd**
	//
	// 	- **parade**
	//
	// 	- **carcrash**
	//
	// 	- **flag**
	//
	// 	- **location**
	//
	// 	- **others**
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score of the snapshot in the category that is indicated by Label. Valid values: `[0, 100]`. The value is accurate to 10 decimal places. The score is representative of the confidence.
	//
	// example:
	//
	// 100
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// The timestamp of the snapshot in the video. Unit: milliseconds.
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
