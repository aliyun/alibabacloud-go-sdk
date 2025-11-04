// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaCensorJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaCensorJobList(v *QueryMediaCensorJobListResponseBodyMediaCensorJobList) *QueryMediaCensorJobListResponseBody
	GetMediaCensorJobList() *QueryMediaCensorJobListResponseBodyMediaCensorJobList
	SetNextPageToken(v string) *QueryMediaCensorJobListResponseBody
	GetNextPageToken() *string
	SetNonExistIds(v *QueryMediaCensorJobListResponseBodyNonExistIds) *QueryMediaCensorJobListResponseBody
	GetNonExistIds() *QueryMediaCensorJobListResponseBodyNonExistIds
	SetRequestId(v string) *QueryMediaCensorJobListResponseBody
	GetRequestId() *string
}

type QueryMediaCensorJobListResponseBody struct {
	// The queried content moderation jobs.
	MediaCensorJobList *QueryMediaCensorJobListResponseBodyMediaCensorJobList `json:"MediaCensorJobList,omitempty" xml:"MediaCensorJobList,omitempty" type:"Struct"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. The value is 32-character UUID. If the returned query results cannot be displayed within one page, this parameter is returned. The value of this parameter is updated for each query.
	//
	// example:
	//
	// 9b1a42bc6e8d46e6a1383b7e7f01****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The IDs of the jobs that do not exist. This parameter is not returned if all the specified jobs are found.
	NonExistIds *QueryMediaCensorJobListResponseBodyNonExistIds `json:"NonExistIds,omitempty" xml:"NonExistIds,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D1D5C080-8E2F-5030-8AB4-13092F17631B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMediaCensorJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBody) GetMediaCensorJobList() *QueryMediaCensorJobListResponseBodyMediaCensorJobList {
	return s.MediaCensorJobList
}

func (s *QueryMediaCensorJobListResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *QueryMediaCensorJobListResponseBody) GetNonExistIds() *QueryMediaCensorJobListResponseBodyNonExistIds {
	return s.NonExistIds
}

func (s *QueryMediaCensorJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMediaCensorJobListResponseBody) SetMediaCensorJobList(v *QueryMediaCensorJobListResponseBodyMediaCensorJobList) *QueryMediaCensorJobListResponseBody {
	s.MediaCensorJobList = v
	return s
}

func (s *QueryMediaCensorJobListResponseBody) SetNextPageToken(v string) *QueryMediaCensorJobListResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBody) SetNonExistIds(v *QueryMediaCensorJobListResponseBodyNonExistIds) *QueryMediaCensorJobListResponseBody {
	s.NonExistIds = v
	return s
}

func (s *QueryMediaCensorJobListResponseBody) SetRequestId(v string) *QueryMediaCensorJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBody) Validate() error {
	if s.MediaCensorJobList != nil {
		if err := s.MediaCensorJobList.Validate(); err != nil {
			return err
		}
	}
	if s.NonExistIds != nil {
		if err := s.NonExistIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobList struct {
	MediaCensorJob []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob `json:"MediaCensorJob,omitempty" xml:"MediaCensorJob,omitempty" type:"Repeated"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobList) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobList) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobList) GetMediaCensorJob() []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	return s.MediaCensorJob
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobList) SetMediaCensorJob(v []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) *QueryMediaCensorJobListResponseBodyMediaCensorJobList {
	s.MediaCensorJob = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobList) Validate() error {
	if s.MediaCensorJob != nil {
		for _, item := range s.MediaCensorJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob struct {
	// The moderation results of live comments.
	BarrageCensorResult *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult `json:"BarrageCensorResult,omitempty" xml:"BarrageCensorResult,omitempty" type:"Struct"`
	// The error code returned if the job failed. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// InvalidParameter.ResourceNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The moderation results of thumbnails.
	CoverImageCensorResults *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResults `json:"CoverImageCensorResults,omitempty" xml:"CoverImageCensorResults,omitempty" type:"Struct"`
	// The time when the content moderation job was created.
	//
	// example:
	//
	// 2021-11-04T07:25:48Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The moderation results of descriptions.
	DescCensorResult *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult `json:"DescCensorResult,omitempty" xml:"DescCensorResult,omitempty" type:"Struct"`
	// The time when the content moderation job was complete.
	//
	// example:
	//
	// 2021-11-04T07:25:50Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The information about the job input.
	Input *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The ID of the content moderation job.
	//
	// example:
	//
	// f8f166eea7a44e9bb0a4aecf9543
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The error message returned if the job failed. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// The resource operated cannot be found
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the MPS queue to which the job was submitted.
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
	// The moderation results of titles.
	TitleCensorResult *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult `json:"TitleCensorResult,omitempty" xml:"TitleCensorResult,omitempty" type:"Struct"`
	// The user-defined data.
	//
	// example:
	//
	// example userdata ****
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The moderation results of videos.
	VensorCensorResult *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult `json:"VensorCensorResult,omitempty" xml:"VensorCensorResult,omitempty" type:"Struct"`
	// The video moderation configurations.
	VideoCensorConfig *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig `json:"VideoCensorConfig,omitempty" xml:"VideoCensorConfig,omitempty" type:"Struct"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetBarrageCensorResult() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult {
	return s.BarrageCensorResult
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetCode() *string {
	return s.Code
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetCoverImageCensorResults() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResults {
	return s.CoverImageCensorResults
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetDescCensorResult() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult {
	return s.DescCensorResult
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetInput() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput {
	return s.Input
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetJobId() *string {
	return s.JobId
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetMessage() *string {
	return s.Message
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetState() *string {
	return s.State
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetTitleCensorResult() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult {
	return s.TitleCensorResult
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetUserData() *string {
	return s.UserData
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetVensorCensorResult() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult {
	return s.VensorCensorResult
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetVideoCensorConfig() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig {
	return s.VideoCensorConfig
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetBarrageCensorResult(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.BarrageCensorResult = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetCode(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.Code = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetCoverImageCensorResults(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResults) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.CoverImageCensorResults = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetCreationTime(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.CreationTime = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetDescCensorResult(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.DescCensorResult = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetFinishTime(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.FinishTime = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetInput(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.Input = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetJobId(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.JobId = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetMessage(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.Message = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetPipelineId(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.PipelineId = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetState(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.State = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetSuggestion(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetTitleCensorResult(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.TitleCensorResult = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetUserData(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.UserData = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetVensorCensorResult(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.VensorCensorResult = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetVideoCensorConfig(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.VideoCensorConfig = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) Validate() error {
	if s.BarrageCensorResult != nil {
		if err := s.BarrageCensorResult.Validate(); err != nil {
			return err
		}
	}
	if s.CoverImageCensorResults != nil {
		if err := s.CoverImageCensorResults.Validate(); err != nil {
			return err
		}
	}
	if s.DescCensorResult != nil {
		if err := s.DescCensorResult.Validate(); err != nil {
			return err
		}
	}
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.TitleCensorResult != nil {
		if err := s.TitleCensorResult.Validate(); err != nil {
			return err
		}
	}
	if s.VensorCensorResult != nil {
		if err := s.VensorCensorResult.Validate(); err != nil {
			return err
		}
	}
	if s.VideoCensorConfig != nil {
		if err := s.VideoCensorConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult struct {
	// The label of the moderation result. Separate multiple labels with commas (,). Valid values:
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
	// The score. Valid values: 0 to 100.
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

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult) GetRate() *string {
	return s.Rate
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult) GetScene() *string {
	return s.Scene
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult) SetLabel(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult) SetRate(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult {
	s.Rate = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult) SetScene(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult {
	s.Scene = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult) SetSuggestion(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResults struct {
	CoverImageCensorResult []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult `json:"CoverImageCensorResult,omitempty" xml:"CoverImageCensorResult,omitempty" type:"Repeated"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResults) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResults) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResults) GetCoverImageCensorResult() []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult {
	return s.CoverImageCensorResult
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResults) SetCoverImageCensorResult(v []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResults {
	s.CoverImageCensorResult = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResults) Validate() error {
	if s.CoverImageCensorResult != nil {
		for _, item := range s.CoverImageCensorResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult struct {
	// The OSS bucket in which the thumbnail is stored.
	//
	// example:
	//
	// example-Bucket-****
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
	Results *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult) GetBucket() *string {
	return s.Bucket
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult) GetLocation() *string {
	return s.Location
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult) GetObject() *string {
	return s.Object
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult) GetResults() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResults {
	return s.Results
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult) SetBucket(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult {
	s.Bucket = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult) SetLocation(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult {
	s.Location = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult) SetObject(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult {
	s.Object = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult) SetResults(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResults) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult {
	s.Results = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResult) Validate() error {
	if s.Results != nil {
		if err := s.Results.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResults struct {
	Result []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResults) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResults) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResults) GetResult() []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult {
	return s.Result
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResults) SetResult(v []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResults {
	s.Result = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResults) Validate() error {
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

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult struct {
	// The label of the moderation result. Separate multiple labels with commas (,).
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
	//     	- **drivelive**: live streaming in a running vehicle.
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
	// normal
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
	// live
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The overall result of the moderation job. Valid values:
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
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult) GetRate() *string {
	return s.Rate
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult) GetScene() *string {
	return s.Scene
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult) SetLabel(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult) SetRate(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult {
	s.Rate = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult) SetScene(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult {
	s.Scene = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult) SetSuggestion(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResultsResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult struct {
	// The label of the moderation result. Separate multiple labels with commas (,). Valid values:
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
	// The score. Valid values: 0 to 100.
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
	// pass
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult) GetRate() *string {
	return s.Rate
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult) GetScene() *string {
	return s.Scene
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult) SetLabel(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult) SetRate(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult {
	s.Rate = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult) SetScene(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult {
	s.Scene = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult) SetSuggestion(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput struct {
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

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput) GetBucket() *string {
	return s.Bucket
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput) GetLocation() *string {
	return s.Location
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput) GetObject() *string {
	return s.Object
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput) SetBucket(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput {
	s.Bucket = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput) SetLocation(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput {
	s.Location = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput) SetObject(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput {
	s.Object = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult struct {
	// The label of the moderation result. Separate multiple labels with commas (,). Valid values:
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
	// The score. Valid values: 0 to 100.
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
	// block
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult) GetRate() *string {
	return s.Rate
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult) GetScene() *string {
	return s.Scene
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult) SetLabel(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult) SetRate(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult {
	s.Rate = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult) SetScene(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult {
	s.Scene = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult) SetSuggestion(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult struct {
	// A collection of moderation results. The information includes the summary about various scenarios such as pornographic content moderation and terrorist content moderation.
	CensorResults *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResults `json:"CensorResults,omitempty" xml:"CensorResults,omitempty" type:"Struct"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ea04afcca7cd4e80b9ece8fbb251
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The moderation results that are sorted in ascending order by time.
	VideoTimelines *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelines `json:"VideoTimelines,omitempty" xml:"VideoTimelines,omitempty" type:"Struct"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult) GetCensorResults() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResults {
	return s.CensorResults
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult) GetVideoTimelines() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelines {
	return s.VideoTimelines
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult) SetCensorResults(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResults) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult {
	s.CensorResults = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult) SetNextPageToken(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult {
	s.NextPageToken = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult) SetVideoTimelines(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelines) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult {
	s.VideoTimelines = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult) Validate() error {
	if s.CensorResults != nil {
		if err := s.CensorResults.Validate(); err != nil {
			return err
		}
	}
	if s.VideoTimelines != nil {
		if err := s.VideoTimelines.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResults struct {
	CensorResult []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult `json:"CensorResult,omitempty" xml:"CensorResult,omitempty" type:"Repeated"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResults) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResults) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResults) GetCensorResult() []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult {
	return s.CensorResult
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResults) SetCensorResult(v []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResults {
	s.CensorResult = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResults) Validate() error {
	if s.CensorResult != nil {
		for _, item := range s.CensorResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult struct {
	// The label of the moderation result. Separate multiple labels with commas (,).
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
	//     	- **drivelive**: live streaming in a running vehicle.
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
	// ad
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

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult) GetRate() *string {
	return s.Rate
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult) GetScene() *string {
	return s.Scene
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult) SetLabel(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult) SetRate(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult {
	s.Rate = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult) SetScene(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult {
	s.Scene = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult) SetSuggestion(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResultsCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelines struct {
	VideoTimeline []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline `json:"VideoTimeline,omitempty" xml:"VideoTimeline,omitempty" type:"Repeated"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelines) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelines) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelines) GetVideoTimeline() []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline {
	return s.VideoTimeline
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelines) SetVideoTimeline(v []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelines {
	s.VideoTimeline = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelines) Validate() error {
	if s.VideoTimeline != nil {
		for _, item := range s.VideoTimeline {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline struct {
	// The moderation results that include information such as labels and scores.
	CensorResults *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResults `json:"CensorResults,omitempty" xml:"CensorResults,omitempty" type:"Struct"`
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

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline) GetCensorResults() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResults {
	return s.CensorResults
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline) GetObject() *string {
	return s.Object
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline) GetTimestamp() *string {
	return s.Timestamp
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline) SetCensorResults(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResults) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline {
	s.CensorResults = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline) SetObject(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline {
	s.Object = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline) SetTimestamp(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline {
	s.Timestamp = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimeline) Validate() error {
	if s.CensorResults != nil {
		if err := s.CensorResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResults struct {
	CensorResult []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult `json:"CensorResult,omitempty" xml:"CensorResult,omitempty" type:"Repeated"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResults) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResults) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResults) GetCensorResult() []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult {
	return s.CensorResult
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResults) SetCensorResult(v []*QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResults {
	s.CensorResult = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResults) Validate() error {
	if s.CensorResult != nil {
		for _, item := range s.CensorResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult struct {
	// The label of the moderation result. Separate multiple labels with commas (,).
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
	//     	- **drivelive**: live streaming in a running vehicle.
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
	// normal
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

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) GetRate() *string {
	return s.Rate
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) GetScene() *string {
	return s.Scene
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) SetLabel(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) SetRate(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult {
	s.Rate = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) SetScene(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult {
	s.Scene = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) SetSuggestion(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResultsCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig struct {
	// The moderation template. Default value: common. The default value indicates that the default template is used.
	//
	// >  If the moderation template is not specified, the default value common is returned. If a custom moderation template that is created by submitting a ticket is specified, the UID of the template is returned.
	//
	// example:
	//
	// common
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The information about output snapshots.
	OutputFile *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
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

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig) GetBizType() *string {
	return s.BizType
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig) GetOutputFile() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile {
	return s.OutputFile
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig) GetVideoCensor() *string {
	return s.VideoCensor
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig) SetBizType(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig {
	s.BizType = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig) SetOutputFile(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig {
	s.OutputFile = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig) SetVideoCensor(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig {
	s.VideoCensor = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig) Validate() error {
	if s.OutputFile != nil {
		if err := s.OutputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile struct {
	// The OSS bucket in which the output snapshot is stored.
	//
	// example:
	//
	// test-bucket-****
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The ID of the region in which the output snapshot resides.
	//
	// example:
	//
	// oss-cn-shanghai
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The OSS object that is generated as the output snapshot.
	//
	// >  In the example, {Count} is a placeholder. The OSS objects that are generated as output snapshots are named `output00001-****.jpg, output00002-****.jpg`, and so on.
	//
	// example:
	//
	// output{Count}.jpg
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile) GetBucket() *string {
	return s.Bucket
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile) GetLocation() *string {
	return s.Location
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile) GetObject() *string {
	return s.Object
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile) SetBucket(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile {
	s.Bucket = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile) SetLocation(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile {
	s.Location = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile) SetObject(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile {
	s.Object = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobListResponseBodyNonExistIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s QueryMediaCensorJobListResponseBodyNonExistIds) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyNonExistIds) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyNonExistIds) GetString_() []*string {
	return s.String_
}

func (s *QueryMediaCensorJobListResponseBodyNonExistIds) SetString_(v []*string) *QueryMediaCensorJobListResponseBodyNonExistIds {
	s.String_ = v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyNonExistIds) Validate() error {
	return dara.Validate(s)
}
