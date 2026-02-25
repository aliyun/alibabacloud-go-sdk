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
	MediaCensorJobList *QueryMediaCensorJobListResponseBodyMediaCensorJobList `json:"MediaCensorJobList,omitempty" xml:"MediaCensorJobList,omitempty" type:"Struct"`
	// The token that is used to retrieve the next page of the query results. The value is a UUID that contains 32 characters. If the returned query results cannot be displayed within one page, this parameter is returned. The value of this parameter is updated for each query.
	//
	// example:
	//
	// 9b1a42bc6e8d46e6a1383b7e7f01****
	NextPageToken *string                                         `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	NonExistIds   *QueryMediaCensorJobListResponseBodyNonExistIds `json:"NonExistIds,omitempty" xml:"NonExistIds,omitempty" type:"Struct"`
	// The ID of the request.
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
	AudioCensorResult       *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult       `json:"AudioCensorResult,omitempty" xml:"AudioCensorResult,omitempty" type:"Struct"`
	BarrageCensorResult     *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult     `json:"BarrageCensorResult,omitempty" xml:"BarrageCensorResult,omitempty" type:"Struct"`
	Code                    *string                                                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	CoverImageCensorResults *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResults `json:"CoverImageCensorResults,omitempty" xml:"CoverImageCensorResults,omitempty" type:"Struct"`
	CreationTime            *string                                                                                     `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DescCensorResult        *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobDescCensorResult        `json:"DescCensorResult,omitempty" xml:"DescCensorResult,omitempty" type:"Struct"`
	FinishTime              *string                                                                                     `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Input                   *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobInput                   `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	JobId                   *string                                                                                     `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message                 *string                                                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PipelineId              *string                                                                                     `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	State                   *string                                                                                     `json:"State,omitempty" xml:"State,omitempty"`
	Suggestion              *string                                                                                     `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TitleCensorResult       *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobTitleCensorResult       `json:"TitleCensorResult,omitempty" xml:"TitleCensorResult,omitempty" type:"Struct"`
	UserData                *string                                                                                     `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VensorCensorResult      *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResult      `json:"VensorCensorResult,omitempty" xml:"VensorCensorResult,omitempty" type:"Struct"`
	VideoCensorConfig       *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfig       `json:"VideoCensorConfig,omitempty" xml:"VideoCensorConfig,omitempty" type:"Struct"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) GetAudioCensorResult() *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult {
	return s.AudioCensorResult
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

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob) SetAudioCensorResult(v *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJob {
	s.AudioCensorResult = v
	return s
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
	if s.AudioCensorResult != nil {
		if err := s.AudioCensorResult.Validate(); err != nil {
			return err
		}
	}
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

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult struct {
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult) GetLabel() *string {
	return s.Label
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult) GetSuggestion() *string {
	return s.Suggestion
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult) SetLabel(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult {
	s.Label = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult) SetSuggestion(v string) *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult {
	s.Suggestion = &v
	return s
}

func (s *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobAudioCensorResult) Validate() error {
	return dara.Validate(s)
}

type QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobBarrageCensorResult struct {
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Rate       *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
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
	Bucket   *string                                                                                                                  `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string                                                                                                                  `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string                                                                                                                  `json:"Object,omitempty" xml:"Object,omitempty"`
	Results  *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobCoverImageCensorResultsCoverImageCensorResultResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
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
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Rate       *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
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
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Rate       *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
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
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Rate       *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
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
	CensorResults  *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultCensorResults  `json:"CensorResults,omitempty" xml:"CensorResults,omitempty" type:"Struct"`
	NextPageToken  *string                                                                                              `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
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
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Rate       *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
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
	CensorResults *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVensorCensorResultVideoTimelinesVideoTimelineCensorResults `json:"CensorResults,omitempty" xml:"CensorResults,omitempty" type:"Struct"`
	Object        *string                                                                                                                        `json:"Object,omitempty" xml:"Object,omitempty"`
	Timestamp     *string                                                                                                                        `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
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
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Rate       *string `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
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
	BizType     *string                                                                                         `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OutputFile  *QueryMediaCensorJobListResponseBodyMediaCensorJobListMediaCensorJobVideoCensorConfigOutputFile `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
	VideoCensor *string                                                                                         `json:"VideoCensor,omitempty" xml:"VideoCensor,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
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
