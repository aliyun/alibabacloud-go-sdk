// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaCensorJobDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaCensorJobDetail(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) *QueryMediaCensorJobDetailResponseBody
	GetMediaCensorJobDetail() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail
	SetRequestId(v string) *QueryMediaCensorJobDetailResponseBody
	GetRequestId() *string
}

type QueryMediaCensorJobDetailResponseBody struct {
	// The results of the content moderation job.
	MediaCensorJobDetail *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail `json:"MediaCensorJobDetail,omitempty" xml:"MediaCensorJobDetail,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B42299E6-F71F-465F-8FE9-4FC2E3D3C2CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMediaCensorJobDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBody) GetMediaCensorJobDetail() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	return s.MediaCensorJobDetail
}

func (s *QueryMediaCensorJobDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMediaCensorJobDetailResponseBody) SetMediaCensorJobDetail(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) *QueryMediaCensorJobDetailResponseBody {
	s.MediaCensorJobDetail = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBody) SetRequestId(v string) *QueryMediaCensorJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail struct {
	// The moderation results of live comments.
	BarrageCensorResult *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult `json:"BarrageCensorResult,omitempty" xml:"BarrageCensorResult,omitempty" type:"Struct"`
	// The error code returned if the job failed. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// InvalidParameter.ResourceNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The moderation results of thumbnails.
	CoverImageCensorResults *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResults `json:"CoverImageCensorResults,omitempty" xml:"CoverImageCensorResults,omitempty" type:"Struct"`
	// The time when the content moderation job was created.
	//
	// example:
	//
	// 2018-09-13T16:32:24Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The moderation results of descriptions.
	DescCensorResult *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult `json:"DescCensorResult,omitempty" xml:"DescCensorResult,omitempty" type:"Struct"`
	// The time when the content moderation job was complete.
	//
	// example:
	//
	// 2018-09-13T16:38:24Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The information about the job input.
	Input *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The ID of the content moderation job.
	//
	// example:
	//
	// f8f166eea7a44e9bb0a4aecf9543****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The error message returned if the job failed. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// The resource operated cannot be found
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the ApsaraVideo Media Processing (MPS) queue to which the job was submitted.
	//
	// example:
	//
	// c5b30b7c0d0e4a0abde1d5f9e751****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The job state.
	//
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The overall result of the content moderation job. Valid values:
	//
	// 	- **pass**: The content passes the moderation.
	//
	// 	- **review**: The content needs to be manually reviewed.
	//
	// 	- **block**: The content needs to be blocked.
	//
	// >  If the moderation result of any type of content is review, the overall result is review. If the moderation result of any type of content is block, the overall result is block.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The moderation results of titles.
	TitleCensorResult *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult `json:"TitleCensorResult,omitempty" xml:"TitleCensorResult,omitempty" type:"Struct"`
	// The user-defined data.
	//
	// example:
	//
	// example userdata ****
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The moderation results of videos.
	VensorCensorResult *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult `json:"VensorCensorResult,omitempty" xml:"VensorCensorResult,omitempty" type:"Struct"`
	// The video moderation configurations.
	VideoCensorConfig *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig `json:"VideoCensorConfig,omitempty" xml:"VideoCensorConfig,omitempty" type:"Struct"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetBarrageCensorResult() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult {
	return s.BarrageCensorResult
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetCode() *string {
	return s.Code
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetCoverImageCensorResults() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResults {
	return s.CoverImageCensorResults
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetDescCensorResult() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult {
	return s.DescCensorResult
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetFinishTime() *string {
	return s.FinishTime
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetInput() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput {
	return s.Input
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetJobId() *string {
	return s.JobId
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetMessage() *string {
	return s.Message
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetState() *string {
	return s.State
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetTitleCensorResult() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult {
	return s.TitleCensorResult
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetUserData() *string {
	return s.UserData
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetVensorCensorResult() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult {
	return s.VensorCensorResult
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) GetVideoCensorConfig() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig {
	return s.VideoCensorConfig
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetBarrageCensorResult(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.BarrageCensorResult = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetCode(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.Code = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetCoverImageCensorResults(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResults) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.CoverImageCensorResults = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetCreationTime(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.CreationTime = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetDescCensorResult(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.DescCensorResult = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetFinishTime(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.FinishTime = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetInput(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.Input = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetJobId(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.JobId = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetMessage(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.Message = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetPipelineId(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.PipelineId = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetState(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.State = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetSuggestion(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetTitleCensorResult(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.TitleCensorResult = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetUserData(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.UserData = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetVensorCensorResult(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.VensorCensorResult = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) SetVideoCensorConfig(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail {
	s.VideoCensorConfig = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetail) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult struct {
	// The label of the moderation result. Valid values:
	//
	// 	- **normal**: normal content.
	//
	// 	- **spam**: spam.
	//
	// 	- **ad**: ads.
	//
	// 	- **abuse**: abuse content.
	//
	// 	- **flood**: excessive junk content.
	//
	// 	- **contraband**: prohibited content.
	//
	// 	- **meaningless**: meaningless content.
	//
	// example:
	//
	// normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score.
	//
	// example:
	//
	// 99.91
	Rate *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The moderation scenario. The value is **antispam**.
	//
	// example:
	//
	// antispam
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The recommended subsequent operation. Valid values:
	//
	// 	- **pass**: The content passes the moderation.
	//
	// 	- **review**: The content needs to be manually reviewed.
	//
	// 	- **block**: The content needs to be blocked.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult) GetRate() *string {
	return s.Rate
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult) GetScene() *string {
	return s.Scene
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult) SetLabel(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult) SetRate(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult {
	s.Rate = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult) SetScene(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult {
	s.Scene = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult) SetSuggestion(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailBarrageCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResults struct {
	CoverImageCensorResult []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult `json:"CoverImageCensorResult,omitempty" xml:"CoverImageCensorResult,omitempty" type:"Repeated"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResults) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResults) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResults) GetCoverImageCensorResult() []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult {
	return s.CoverImageCensorResult
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResults) SetCoverImageCensorResult(v []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResults {
	s.CoverImageCensorResult = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResults) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult struct {
	// The OSS bucket in which the thumbnail is stored.
	//
	// example:
	//
	// bucket-out-test-****
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS region in which the thumbnail resides.
	//
	// example:
	//
	// oss-cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The Object Storage Service (OSS) object that is used as the thumbnail.
	//
	// example:
	//
	// test/ai/censor/v2/vme-****.jpg
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The moderation results.
	Results *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult) GetBucket() *string {
	return s.Bucket
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult) GetLocation() *string {
	return s.Location
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult) GetObject() *string {
	return s.Object
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult) GetResults() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResults {
	return s.Results
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult) SetBucket(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult {
	s.Bucket = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult) SetLocation(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult {
	s.Location = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult) SetObject(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult {
	s.Object = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult) SetResults(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResults) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult {
	s.Results = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResults struct {
	Result []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResults) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResults) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResults) GetResult() []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult {
	return s.Result
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResults) SetResult(v []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResults {
	s.Result = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResults) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult struct {
	// The label of the moderation result.
	//
	// 	- Valid values in the pornographic content moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **sexy**: sexy content.
	//
	//     	- **porn**: pornographic content.
	//
	// 	- Valid values in the terrorist content moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **bloody**: bloody content.
	//
	//     	- **explosion**: explosion and smoke.
	//
	//     	- **outfit**: special costume.
	//
	//     	- **logo**: special logo.
	//
	//     	- **weapon**: weapon.
	//
	//     	- **politics**: political content.
	//
	//     	- **violence**: violence.
	//
	//     	- **crowd**: crowd.
	//
	//     	- **parade**: parade.
	//
	//     	- **carcrash**: car accident.
	//
	//     	- **flag**: flag.
	//
	//     	- **location**: landmark.
	//
	//     	- **others**: other content.
	//
	// 	- Valid values in the ad moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **ad**: other ads.
	//
	//     	- **politics**: political content in text.
	//
	//     	- **porn**: pornographic content in text.
	//
	//     	- **abuse**: abuse in text.
	//
	//     	- **terrorism**: terrorist content in text.
	//
	//     	- **contraband**: prohibited content in text.
	//
	//     	- **spam**: spam in text.
	//
	//     	- **npx**: illegal ad.
	//
	//     	- **qrcode**: QR code.
	//
	//     	- **programCode**: mini program code.
	//
	// 	- Valid values in the undesirable scene moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **meaningless**: meaningless content, such as a black or white screen.
	//
	//     	- **PIP**: picture-in-picture.
	//
	//     	- **smoking**: smoking.
	//
	//     	- **drivelive**: live broadcasting in a running vehicle.
	//
	// 	- Valid values in the logo moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **TV**: controlled logo.
	//
	//     	- **trademark**: trademark.
	//
	// example:
	//
	// Normal
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score. Valid values: 0 to 100.
	//
	// example:
	//
	// 100
	Rate *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The moderation scenario. Valid values:
	//
	// 	- **porn**: pornographic content moderation.
	//
	// 	- **terrorism**: terrorist content moderation.
	//
	// 	- **ad**: ad moderation.
	//
	// 	- **live**: undesirable scene moderation.
	//
	// 	- **logo**: logo moderation.
	//
	// example:
	//
	// Antispam
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The recommended subsequent operation. Valid values:
	//
	// 	- **pass**: The content passes the moderation.
	//
	// 	- **review**: The content needs to be manually reviewed.
	//
	// 	- **block**: The content needs to be blocked.
	//
	// example:
	//
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult) GetRate() *string {
	return s.Rate
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult) GetScene() *string {
	return s.Scene
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult) SetLabel(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult) SetRate(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult {
	s.Rate = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult) SetScene(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult {
	s.Scene = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult) SetSuggestion(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailCoverImageCensorResultsCoverImageCensorResultResultsResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult struct {
	// The label of the moderation result. Valid values:
	//
	// 	- **normal**: normal content.
	//
	// 	- **spam**: spam.
	//
	// 	- **ad**: ads.
	//
	// 	- **abuse**: abuse content.
	//
	// 	- **flood**: excessive junk content.
	//
	// 	- **contraband**: prohibited content.
	//
	// 	- **meaningless**: meaningless content.
	//
	// example:
	//
	// terrorism
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score.
	//
	// example:
	//
	// 100
	Rate *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The moderation scenario. The value is **antispam**.
	//
	// example:
	//
	// antispam
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The recommended subsequent operation. Valid values:
	//
	// 	- **pass**: The content passes the moderation.
	//
	// 	- **review**: The content needs to be manually reviewed.
	//
	// 	- **block**: The content needs to be blocked.
	//
	// example:
	//
	// review
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult) GetRate() *string {
	return s.Rate
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult) GetScene() *string {
	return s.Scene
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult) SetLabel(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult) SetRate(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult {
	s.Rate = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult) SetScene(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult {
	s.Scene = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult) SetSuggestion(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailDescCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput struct {
	// The name of the OSS bucket in which the input file is stored.
	//
	// example:
	//
	// bucket-test-in-****
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS region in which the input file resides.
	//
	// example:
	//
	// oss-cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the OSS object that is used as the input file.
	//
	// example:
	//
	// test/ai/censor/test-****.mp4
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput) GetBucket() *string {
	return s.Bucket
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput) GetLocation() *string {
	return s.Location
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput) GetObject() *string {
	return s.Object
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput) SetBucket(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput {
	s.Bucket = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput) SetLocation(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput {
	s.Location = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput) SetObject(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput {
	s.Object = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailInput) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult struct {
	// The label of the moderation result. Valid values:
	//
	// 	- **normal**: normal content.
	//
	// 	- **spam**: spam.
	//
	// 	- **ad**: ads.
	//
	// 	- **abuse**: abuse content.
	//
	// 	- **flood**: excessive junk content.
	//
	// 	- **contraband**: prohibited content.
	//
	// 	- **meaningless**: meaningless content.
	//
	// example:
	//
	// meaningless
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score.
	//
	// example:
	//
	// 99.91
	Rate *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The moderation scenario. The value is **antispam**.
	//
	// example:
	//
	// antispam
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The recommended subsequent operation. Valid values:
	//
	// 	- **pass**: The content passes the moderation.
	//
	// 	- **review**: The content needs to be manually reviewed.
	//
	// 	- **block**: The content needs to be blocked.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult) GetRate() *string {
	return s.Rate
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult) GetScene() *string {
	return s.Scene
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult) SetLabel(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult) SetRate(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult {
	s.Rate = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult) SetScene(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult {
	s.Scene = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult) SetSuggestion(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailTitleCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult struct {
	// A collection of moderation results. The information includes the summary about various scenarios such as pornographic content moderation and terrorist content moderation.
	CensorResults *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResults `json:"CensorResults,omitempty" xml:"CensorResults,omitempty" type:"Struct"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ea04afcca7cd4e80b9ece8fbb251****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The moderation results that are sorted in ascending order by time.
	VideoTimelines *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelines `json:"VideoTimelines,omitempty" xml:"VideoTimelines,omitempty" type:"Struct"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult) GetCensorResults() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResults {
	return s.CensorResults
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult) GetVideoTimelines() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelines {
	return s.VideoTimelines
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult) SetCensorResults(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResults) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult {
	s.CensorResults = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult) SetNextPageToken(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult {
	s.NextPageToken = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult) SetVideoTimelines(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelines) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult {
	s.VideoTimelines = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResults struct {
	CensorResult []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult `json:"CensorResult,omitempty" xml:"CensorResult,omitempty" type:"Repeated"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResults) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResults) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResults) GetCensorResult() []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult {
	return s.CensorResult
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResults) SetCensorResult(v []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResults {
	s.CensorResult = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResults) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult struct {
	// The label of the moderation result.
	//
	// 	- Valid values in the pornographic content moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **sexy**: sexy content.
	//
	//     	- **porn**: pornographic content.
	//
	// 	- Valid values in the terrorist content moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **bloody**: bloody content.
	//
	//     	- **explosion**: explosion and smoke.
	//
	//     	- **outfit**: special costume.
	//
	//     	- **logo**: special logo.
	//
	//     	- **weapon**: weapon.
	//
	//     	- **politics**: political content.
	//
	//     	- **violence**: violence.
	//
	//     	- **crowd**: crowd.
	//
	//     	- **parade**: parade.
	//
	//     	- **carcrash**: car accident.
	//
	//     	- **flag**: flag.
	//
	//     	- **location**: landmark.
	//
	//     	- **others**: other content.
	//
	// 	- Valid values in the ad moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **ad**: other ads.
	//
	//     	- **politics**: political content in text.
	//
	//     	- **porn**: pornographic content in text.
	//
	//     	- **abuse**: abuse in text.
	//
	//     	- **terrorism**: terrorist content in text.
	//
	//     	- **contraband**: prohibited content in text.
	//
	//     	- **spam**: spam in text.
	//
	//     	- **npx**: illegal ad.
	//
	//     	- **qrcode**: QR code.
	//
	//     	- **programCode**: mini program code.
	//
	// 	- Valid values in the undesirable scene moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **meaningless**: meaningless content, such as a black or white screen.
	//
	//     	- **PIP**: picture-in-picture.
	//
	//     	- **smoking**: smoking.
	//
	//     	- **drivelive**: live broadcasting in a running vehicle.
	//
	// 	- Valid values in the logo moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **TV**: controlled logo.
	//
	//     	- **trademark**: trademark.
	//
	// example:
	//
	// meaningless
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score.
	//
	// example:
	//
	// 100
	Rate *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The moderation scenario. Valid values:
	//
	// 	- **porn**: pornographic content moderation.
	//
	// 	- **terrorism**: terrorist content moderation.
	//
	// 	- **ad**: ad moderation.
	//
	// 	- **live**: undesirable scene moderation.
	//
	// 	- **logo**: logo moderation.
	//
	// example:
	//
	// terrorism
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The recommended subsequent operation. Valid values:
	//
	// 	- **pass**: The content passes the moderation.
	//
	// 	- **review**: The content needs to be manually reviewed.
	//
	// 	- **block**: The content needs to be blocked.
	//
	// example:
	//
	// review
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult) GetRate() *string {
	return s.Rate
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult) GetScene() *string {
	return s.Scene
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult) SetLabel(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult) SetRate(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult {
	s.Rate = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult) SetScene(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult {
	s.Scene = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult) SetSuggestion(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultCensorResultsCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelines struct {
	VideoTimeline []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline `json:"VideoTimeline,omitempty" xml:"VideoTimeline,omitempty" type:"Repeated"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelines) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelines) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelines) GetVideoTimeline() []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline {
	return s.VideoTimeline
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelines) SetVideoTimeline(v []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelines {
	s.VideoTimeline = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelines) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline struct {
	// The moderation results that include information such as labels and scores.
	CensorResults *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResults `json:"CensorResults,omitempty" xml:"CensorResults,omitempty" type:"Struct"`
	// The OSS object that is generated as the output snapshot.
	//
	// >  In the example, {Count} is a placeholder. The OSS objects that are generated as output snapshots are named `output00001-****.jpg`, `output00002-****.jpg`, and so on.
	//
	// example:
	//
	// output{Count}.jpg
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The position in the video. Format: `hh:mm:ss[.SSS]`.
	//
	// example:
	//
	// 00:02:59.999
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline) GetCensorResults() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResults {
	return s.CensorResults
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline) GetObject() *string {
	return s.Object
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline) GetTimestamp() *string {
	return s.Timestamp
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline) SetCensorResults(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResults) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline {
	s.CensorResults = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline) SetObject(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline {
	s.Object = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline) SetTimestamp(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline {
	s.Timestamp = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimeline) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResults struct {
	CensorResult []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult `json:"CensorResult,omitempty" xml:"CensorResult,omitempty" type:"Repeated"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResults) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResults) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResults) GetCensorResult() []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult {
	return s.CensorResult
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResults) SetCensorResult(v []*QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResults {
	s.CensorResult = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResults) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult struct {
	// The label of the moderation result.
	//
	// 	- Valid values in the pornographic content moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **sexy**: sexy content.
	//
	//     	- **porn**: pornographic content.
	//
	// 	- Valid values in the terrorist content moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **bloody**: bloody content.
	//
	//     	- **explosion**: explosion and smoke.
	//
	//     	- **outfit**: special costume.
	//
	//     	- **logo**: special logo.
	//
	//     	- **weapon**: weapon.
	//
	//     	- **politics**: political content.
	//
	//     	- **violence**: violence.
	//
	//     	- **crowd**: crowd.
	//
	//     	- **parade**: parade.
	//
	//     	- **carcrash**: car accident.
	//
	//     	- **flag**: flag.
	//
	//     	- **location**: landmark.
	//
	//     	- **others**: other content.
	//
	// 	- Valid values in the ad moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **ad**: other ads.
	//
	//     	- **politics**: political content in text.
	//
	//     	- **porn**: pornographic content in text.
	//
	//     	- **abuse**: abuse in text.
	//
	//     	- **terrorism**: terrorist content in text.
	//
	//     	- **contraband**: prohibited content in text.
	//
	//     	- **spam**: spam in text.
	//
	//     	- **npx**: illegal ad.
	//
	//     	- **qrcode**: QR code.
	//
	//     	- **programCode**: mini program code.
	//
	// 	- Valid values in the undesirable scene moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **meaningless**: meaningless content, such as a black or white screen.
	//
	//     	- **PIP**: picture-in-picture.
	//
	//     	- **smoking**: smoking.
	//
	//     	- **drivelive**: live broadcasting in a running vehicle.
	//
	// 	- Valid values in the logo moderation scenario:
	//
	//     	- **normal**: normal content.
	//
	//     	- **TV**: controlled logo.
	//
	//     	- **trademark**: trademark.
	//
	// example:
	//
	// flood
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The score.
	//
	// example:
	//
	// 99.99
	Rate *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The moderation scenario. Valid values:
	//
	// 	- **porn**: pornographic content moderation.
	//
	// 	- **terrorism**: terrorist content moderation.
	//
	// 	- **ad**: ad moderation.
	//
	// 	- **live**: undesirable scene moderation.
	//
	// 	- **logo**: logo moderation.
	//
	// example:
	//
	// porn
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The recommended subsequent operation. Valid values:
	//
	// 	- **pass**: The content passes the moderation.
	//
	// 	- **review**: The content needs to be manually reviewed.
	//
	// 	- **block**: The content needs to be blocked.
	//
	// example:
	//
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) GetRate() *string {
	return s.Rate
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) GetScene() *string {
	return s.Scene
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) SetLabel(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) SetRate(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult {
	s.Rate = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) SetScene(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult {
	s.Scene = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) SetSuggestion(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig struct {
	// The custom business type. Default value: common.
	//
	// example:
	//
	// common
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The information about output snapshots.
	OutputFile *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
	// Indicates whether the video content needs to be moderated. Default value: **true**. Valid values:
	//
	// 	- **true**: The video content needs to be moderated.
	//
	// 	- **false**: The video content does not need to be moderated.
	//
	// example:
	//
	// true
	VideoCensor *string `json:"VideoCensor,omitempty" xml:"VideoCensor,omitempty"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig) GetBizType() *string {
	return s.BizType
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig) GetOutputFile() *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile {
	return s.OutputFile
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig) GetVideoCensor() *string {
	return s.VideoCensor
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig) SetBizType(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig {
	s.BizType = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig) SetOutputFile(v *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig {
	s.OutputFile = v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig) SetVideoCensor(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig {
	s.VideoCensor = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfig) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile struct {
	// The OSS bucket in which the output snapshot is stored.
	//
	// example:
	//
	// test-bucket-****
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS region in which the output snapshot resides.
	//
	// example:
	//
	// oss-cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The OSS object that is generated as the output snapshot.
	//
	// >  In the example, {Count} is a placeholder. The OSS objects that are generated as output snapshots are named `output00001-****.jpg`, `output00002-****.jpg`, and so on.
	//
	// example:
	//
	// output{Count}.jpg
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile) GetBucket() *string {
	return s.Bucket
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile) GetLocation() *string {
	return s.Location
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile) GetObject() *string {
	return s.Object
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile) SetBucket(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile {
	s.Bucket = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile) SetLocation(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile {
	s.Location = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile) SetObject(v string) *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile {
	s.Object = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponseBodyMediaCensorJobDetailVideoCensorConfigOutputFile) Validate() error {
	return dara.Validate(s)
}
