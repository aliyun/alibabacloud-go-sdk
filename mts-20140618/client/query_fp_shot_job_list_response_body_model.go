// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFpShotJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFpShotJobList(v *QueryFpShotJobListResponseBodyFpShotJobList) *QueryFpShotJobListResponseBody
	GetFpShotJobList() *QueryFpShotJobListResponseBodyFpShotJobList
	SetNextPageToken(v string) *QueryFpShotJobListResponseBody
	GetNextPageToken() *string
	SetNonExistIds(v *QueryFpShotJobListResponseBodyNonExistIds) *QueryFpShotJobListResponseBody
	GetNonExistIds() *QueryFpShotJobListResponseBodyNonExistIds
	SetRequestId(v string) *QueryFpShotJobListResponseBody
	GetRequestId() *string
}

type QueryFpShotJobListResponseBody struct {
	// The information about media fingerprint analysis jobs.
	FpShotJobList *QueryFpShotJobListResponseBodyFpShotJobList `json:"FpShotJobList,omitempty" xml:"FpShotJobList,omitempty" type:"Struct"`
	// The token that is used to retrieve the next page of the query results. The value is a 32-bit UUID. If the returned query results cannot be displayed within one page, this parameter is returned. The value of this parameter is updated for each query.
	//
	// example:
	//
	// b11c171cced04565b1f38f1ecc39****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The IDs of the jobs that do not exist.
	NonExistIds *QueryFpShotJobListResponseBodyNonExistIds `json:"NonExistIds,omitempty" xml:"NonExistIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFpShotJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBody) GetFpShotJobList() *QueryFpShotJobListResponseBodyFpShotJobList {
	return s.FpShotJobList
}

func (s *QueryFpShotJobListResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *QueryFpShotJobListResponseBody) GetNonExistIds() *QueryFpShotJobListResponseBodyNonExistIds {
	return s.NonExistIds
}

func (s *QueryFpShotJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFpShotJobListResponseBody) SetFpShotJobList(v *QueryFpShotJobListResponseBodyFpShotJobList) *QueryFpShotJobListResponseBody {
	s.FpShotJobList = v
	return s
}

func (s *QueryFpShotJobListResponseBody) SetNextPageToken(v string) *QueryFpShotJobListResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *QueryFpShotJobListResponseBody) SetNonExistIds(v *QueryFpShotJobListResponseBodyNonExistIds) *QueryFpShotJobListResponseBody {
	s.NonExistIds = v
	return s
}

func (s *QueryFpShotJobListResponseBody) SetRequestId(v string) *QueryFpShotJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFpShotJobListResponseBody) Validate() error {
	if s.FpShotJobList != nil {
		if err := s.FpShotJobList.Validate(); err != nil {
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

type QueryFpShotJobListResponseBodyFpShotJobList struct {
	FpShotJob []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJob `json:"FpShotJob,omitempty" xml:"FpShotJob,omitempty" type:"Repeated"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobList) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobList) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobList) GetFpShotJob() []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	return s.FpShotJob
}

func (s *QueryFpShotJobListResponseBodyFpShotJobList) SetFpShotJob(v []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) *QueryFpShotJobListResponseBodyFpShotJobList {
	s.FpShotJob = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobList) Validate() error {
	if s.FpShotJob != nil {
		for _, item := range s.FpShotJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJob struct {
	// The error code returned if the job fails.
	//
	// example:
	//
	// InvalidParameter.UUIDFormatInvalid
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2017-01-10T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The length of the input file.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 5
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the uploaded file.
	//
	// example:
	//
	// ebb51ee30f0b49aba959823fa991****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The time when the job was complete.
	//
	// example:
	//
	// 0
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The configurations of the job.
	FpShotConfig *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig `json:"FpShotConfig,omitempty" xml:"FpShotConfig,omitempty" type:"Struct"`
	// The results of the media fingerprint analysis job.
	FpShotResult *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult `json:"FpShotResult,omitempty" xml:"FpShotResult,omitempty" type:"Struct"`
	// The ID of the job.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the job input.
	//
	// example:
	//
	// {"Bucket":"oss-test","Location":"oss-cn-beijing","Object":"test.mp4"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The information about the job input.
	InputFile *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile `json:"InputFile,omitempty" xml:"InputFile,omitempty" type:"Struct"`
	// The error message returned if the job fails. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// The parameter \\"Id\\" is invalid.A uuid must:1)be comprised of chracters[a-f],numbers[0-9];2)be 32 characters long
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the MPS queue to which the analysis job is submitted.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- **Queuing**: The job is waiting in the queue.
	//
	// 	- **Analysing**: The job is in progress.
	//
	// 	- **Success**: The job is successful.
	//
	// 	- **Fail**: The job fails.
	//
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The custom data.
	//
	// example:
	//
	// testid-001
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetCode() *string {
	return s.Code
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetFileId() *string {
	return s.FileId
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetFpShotConfig() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig {
	return s.FpShotConfig
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetFpShotResult() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult {
	return s.FpShotResult
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetId() *string {
	return s.Id
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetInput() *string {
	return s.Input
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetInputFile() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile {
	return s.InputFile
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetMessage() *string {
	return s.Message
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetState() *string {
	return s.State
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) GetUserData() *string {
	return s.UserData
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetCode(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.Code = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetCreationTime(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.CreationTime = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetDuration(v int32) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.Duration = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetFileId(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.FileId = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetFinishTime(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.FinishTime = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetFpShotConfig(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.FpShotConfig = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetFpShotResult(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.FpShotResult = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetId(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.Id = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetInput(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.Input = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetInputFile(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.InputFile = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetMessage(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.Message = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetPipelineId(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.PipelineId = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetState(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.State = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) SetUserData(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob {
	s.UserData = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJob) Validate() error {
	if s.FpShotConfig != nil {
		if err := s.FpShotConfig.Validate(); err != nil {
			return err
		}
	}
	if s.FpShotResult != nil {
		if err := s.FpShotResult.Validate(); err != nil {
			return err
		}
	}
	if s.InputFile != nil {
		if err := s.InputFile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig struct {
	// The ID of the media fingerprint library.
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****
	FpDBId *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	// The unique primary key of the video.
	//
	// example:
	//
	// 3ca84a39a9024f19853b21be9cf9****
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The storage type. Valid values:
	//
	// 	- **nosave**: The fingerprints of the job input are not saved to the media fingerprint library.
	//
	// 	- **save**: The fingerprints of the job input are saved to the media fingerprint library only if the job input is not duplicated with media content in the media fingerprint library.
	//
	// 	- **forcesave**: The fingerprints of the job input are forcibly saved to the media fingerprint library.
	//
	// example:
	//
	// save
	SaveType *string `json:"SaveType,omitempty" xml:"SaveType,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig) GetFpDBId() *string {
	return s.FpDBId
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig) GetSaveType() *string {
	return s.SaveType
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig) SetFpDBId(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig {
	s.FpDBId = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig) SetPrimaryKey(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig {
	s.PrimaryKey = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig) SetSaveType(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig {
	s.SaveType = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotConfig) Validate() error {
	return dara.Validate(s)
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult struct {
	// The audio fingerprint analysis results.
	AudioFpShots *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShots `json:"AudioFpShots,omitempty" xml:"AudioFpShots,omitempty" type:"Struct"`
	// The video fingerprint analysis results.
	FpShots *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShots `json:"FpShots,omitempty" xml:"FpShots,omitempty" type:"Struct"`
	// The text fingerprint analysis results.
	TextFpShots *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShots `json:"TextFpShots,omitempty" xml:"TextFpShots,omitempty" type:"Struct"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult) GetAudioFpShots() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShots {
	return s.AudioFpShots
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult) GetFpShots() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShots {
	return s.FpShots
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult) GetTextFpShots() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShots {
	return s.TextFpShots
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult) SetAudioFpShots(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShots) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult {
	s.AudioFpShots = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult) SetFpShots(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShots) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult {
	s.FpShots = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult) SetTextFpShots(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShots) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult {
	s.TextFpShots = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResult) Validate() error {
	if s.AudioFpShots != nil {
		if err := s.AudioFpShots.Validate(); err != nil {
			return err
		}
	}
	if s.FpShots != nil {
		if err := s.FpShots.Validate(); err != nil {
			return err
		}
	}
	if s.TextFpShots != nil {
		if err := s.TextFpShots.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShots struct {
	FpShot []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot `json:"FpShot,omitempty" xml:"FpShot,omitempty" type:"Repeated"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShots) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShots) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShots) GetFpShot() []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot {
	return s.FpShot
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShots) SetFpShot(v []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShots {
	s.FpShot = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShots) Validate() error {
	if s.FpShot != nil {
		for _, item := range s.FpShot {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot struct {
	// The audio files that have similar fingerprints to the input audio in the audio fingerprint library.
	FpShotSlices *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlices `json:"FpShotSlices,omitempty" xml:"FpShotSlices,omitempty" type:"Struct"`
	// The unique primary key of the input audio.
	//
	// example:
	//
	// 498ac941373341599c4777c8d884****
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The overall similarity of the input audio against audio files that have similar fingerprints to the input audio in the audio fingerprint library.
	//
	// example:
	//
	// 0
	Similarity *string `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot) GetFpShotSlices() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlices {
	return s.FpShotSlices
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot) GetSimilarity() *string {
	return s.Similarity
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot) SetFpShotSlices(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlices) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot {
	s.FpShotSlices = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot) SetPrimaryKey(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot {
	s.PrimaryKey = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot) SetSimilarity(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot {
	s.Similarity = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShot) Validate() error {
	if s.FpShotSlices != nil {
		if err := s.FpShotSlices.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlices struct {
	FpShotSlice []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice `json:"FpShotSlice,omitempty" xml:"FpShotSlice,omitempty" type:"Repeated"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlices) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlices) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlices) GetFpShotSlice() []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice {
	return s.FpShotSlice
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlices) SetFpShotSlice(v []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlices {
	s.FpShotSlice = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlices) Validate() error {
	if s.FpShotSlice != nil {
		for _, item := range s.FpShotSlice {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice struct {
	// The start point in time and duration of the similar audio clip in the audio file that has similar fingerprints to the input audio in the audio fingerprint library.
	Duplication *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication `json:"Duplication,omitempty" xml:"Duplication,omitempty" type:"Struct"`
	// The start point in time and duration of the similar audio clip in the input audio.
	Input *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The similarity of the input audio against the audio file that has similar fingerprints to the input audio in the audio fingerprint library.
	//
	// example:
	//
	// 0
	Similarity *string `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice) GetDuplication() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication {
	return s.Duplication
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice) GetInput() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput {
	return s.Input
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice) GetSimilarity() *string {
	return s.Similarity
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice) SetDuplication(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice {
	s.Duplication = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice) SetInput(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice {
	s.Input = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice) SetSimilarity(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice {
	s.Similarity = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSlice) Validate() error {
	if s.Duplication != nil {
		if err := s.Duplication.Validate(); err != nil {
			return err
		}
	}
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication struct {
	// The duration of the similar audio clip in the audio file that has similar fingerprints to the input audio in the audio fingerprint library.
	//
	// example:
	//
	// 3
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start point in time of the similar audio clip in the audio file that has similar fingerprints to the input audio in the audio fingerprint library.
	//
	// example:
	//
	// 0
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication) GetDuration() *string {
	return s.Duration
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication) GetStart() *string {
	return s.Start
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication) SetDuration(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication {
	s.Duration = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication) SetStart(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication {
	s.Start = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceDuplication) Validate() error {
	return dara.Validate(s)
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput struct {
	// The duration of the similar audio clip in the input audio.
	//
	// example:
	//
	// 5
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start point in time of the similar audio clip in the input audio.
	//
	// example:
	//
	// 0
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput) GetDuration() *string {
	return s.Duration
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput) GetStart() *string {
	return s.Start
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput) SetDuration(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput {
	s.Duration = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput) SetStart(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput {
	s.Start = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultAudioFpShotsFpShotFpShotSlicesFpShotSliceInput) Validate() error {
	return dara.Validate(s)
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShots struct {
	FpShot []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot `json:"FpShot,omitempty" xml:"FpShot,omitempty" type:"Repeated"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShots) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShots) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShots) GetFpShot() []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot {
	return s.FpShot
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShots) SetFpShot(v []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShots {
	s.FpShot = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShots) Validate() error {
	if s.FpShot != nil {
		for _, item := range s.FpShot {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot struct {
	// The video files that have similar fingerprints to the input video in the video fingerprint library.
	FpShotSlices *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlices `json:"FpShotSlices,omitempty" xml:"FpShotSlices,omitempty" type:"Struct"`
	// The unique primary key of the input video.
	//
	// example:
	//
	// 498ac941373341599c4777c8d884****
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The overall similarity of the input video against video files that have similar fingerprints to the input video in the video fingerprint library.
	//
	// >  The overall similarity is the average value of the similarities of the input video clips with the clips of the video that has a similar fingerprint. If multiple video files that have similar fingerprints to the input video exist in the video fingerprint library, the similarities of the input video against multiple similar video clips are returned.
	//
	// example:
	//
	// 0.8914769887924194
	Similarity *string `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot) GetFpShotSlices() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlices {
	return s.FpShotSlices
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot) GetSimilarity() *string {
	return s.Similarity
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot) SetFpShotSlices(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlices) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot {
	s.FpShotSlices = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot) SetPrimaryKey(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot {
	s.PrimaryKey = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot) SetSimilarity(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot {
	s.Similarity = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShot) Validate() error {
	if s.FpShotSlices != nil {
		if err := s.FpShotSlices.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlices struct {
	FpShotSlice []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice `json:"FpShotSlice,omitempty" xml:"FpShotSlice,omitempty" type:"Repeated"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlices) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlices) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlices) GetFpShotSlice() []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice {
	return s.FpShotSlice
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlices) SetFpShotSlice(v []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlices {
	s.FpShotSlice = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlices) Validate() error {
	if s.FpShotSlice != nil {
		for _, item := range s.FpShotSlice {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice struct {
	// The start point in time and duration of the similar video clip in the video file that has similar fingerprints to the input video in the video fingerprint library.
	Duplication *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication `json:"Duplication,omitempty" xml:"Duplication,omitempty" type:"Struct"`
	// The start time and duration of the similar video clip in the input video.
	Input *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The similarity of the input video clip against the video file that has similar fingerprints to the input video in the video fingerprint library.
	//
	// example:
	//
	// 0
	Similarity *string `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice) GetDuplication() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication {
	return s.Duplication
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice) GetInput() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput {
	return s.Input
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice) GetSimilarity() *string {
	return s.Similarity
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice) SetDuplication(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice {
	s.Duplication = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice) SetInput(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice {
	s.Input = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice) SetSimilarity(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice {
	s.Similarity = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSlice) Validate() error {
	if s.Duplication != nil {
		if err := s.Duplication.Validate(); err != nil {
			return err
		}
	}
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication struct {
	// The duration of the similar video clip in the video file that has similar fingerprints to the input video in the video fingerprint library.
	//
	// example:
	//
	// 48
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start point in time of the similar video clip in the video file that has similar fingerprints to the input video in the video fingerprint library.
	//
	// example:
	//
	// 1260
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication) GetDuration() *string {
	return s.Duration
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication) GetStart() *string {
	return s.Start
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication) SetDuration(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication {
	s.Duration = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication) SetStart(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication {
	s.Start = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceDuplication) Validate() error {
	return dara.Validate(s)
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput struct {
	// The duration of the similar video clip in the input video.
	//
	// example:
	//
	// 48
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start point in time of the similar video clip in the input video.
	//
	// example:
	//
	// 46
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput) GetDuration() *string {
	return s.Duration
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput) GetStart() *string {
	return s.Start
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput) SetDuration(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput {
	s.Duration = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput) SetStart(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput {
	s.Start = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultFpShotsFpShotFpShotSlicesFpShotSliceInput) Validate() error {
	return dara.Validate(s)
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShots struct {
	TextFpShot []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot `json:"TextFpShot,omitempty" xml:"TextFpShot,omitempty" type:"Repeated"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShots) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShots) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShots) GetTextFpShot() []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot {
	return s.TextFpShot
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShots) SetTextFpShot(v []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShots {
	s.TextFpShot = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShots) Validate() error {
	if s.TextFpShot != nil {
		for _, item := range s.TextFpShot {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot struct {
	// The unique primary key of the input text.
	//
	// example:
	//
	// 3e34ac649945b53a1b0f863ce030****
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The similarity of the input text against text snippets that have similar fingerprints to the input text in the text fingerprint library.
	//
	// example:
	//
	// 1.0
	Similarity *string `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
	// The text snippets that have similar fingerprints to the input text in the text fingerprint library.
	TextFpShotSlices *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlices `json:"TextFpShotSlices,omitempty" xml:"TextFpShotSlices,omitempty" type:"Struct"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot) GetSimilarity() *string {
	return s.Similarity
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot) GetTextFpShotSlices() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlices {
	return s.TextFpShotSlices
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot) SetPrimaryKey(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot {
	s.PrimaryKey = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot) SetSimilarity(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot {
	s.Similarity = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot) SetTextFpShotSlices(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlices) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot {
	s.TextFpShotSlices = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShot) Validate() error {
	if s.TextFpShotSlices != nil {
		if err := s.TextFpShotSlices.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlices struct {
	TextFpShotSlice []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice `json:"TextFpShotSlice,omitempty" xml:"TextFpShotSlice,omitempty" type:"Repeated"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlices) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlices) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlices) GetTextFpShotSlice() []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice {
	return s.TextFpShotSlice
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlices) SetTextFpShotSlice(v []*QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlices {
	s.TextFpShotSlice = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlices) Validate() error {
	if s.TextFpShotSlice != nil {
		for _, item := range s.TextFpShotSlice {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice struct {
	// The text snippet that has similar fingerprints to the input text in the text fingerprint library.
	//
	// example:
	//
	// It\\"s snowy today.
	DuplicationText *string `json:"DuplicationText,omitempty" xml:"DuplicationText,omitempty"`
	// The start point in time and duration of the similar text snippet in the input text.
	InputFragment *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment `json:"InputFragment,omitempty" xml:"InputFragment,omitempty" type:"Struct"`
	// The input text for text fingerprint analysis.
	//
	// example:
	//
	// It\\"s snowy today.
	InputText *string `json:"InputText,omitempty" xml:"InputText,omitempty"`
	// The similarity of the input text against the text snippet that has similar fingerprints to the input text in the text fingerprint library.
	//
	// example:
	//
	// 1.0
	Similarity *string `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice) GetDuplicationText() *string {
	return s.DuplicationText
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice) GetInputFragment() *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment {
	return s.InputFragment
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice) GetInputText() *string {
	return s.InputText
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice) GetSimilarity() *string {
	return s.Similarity
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice) SetDuplicationText(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice {
	s.DuplicationText = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice) SetInputFragment(v *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice {
	s.InputFragment = v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice) SetInputText(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice {
	s.InputText = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice) SetSimilarity(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice {
	s.Similarity = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSlice) Validate() error {
	if s.InputFragment != nil {
		if err := s.InputFragment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment struct {
	// The duration of the similar text snippet in the input text.
	//
	// example:
	//
	// 3
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start time of the similar text snippet in the input text.
	//
	// example:
	//
	// 0
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment) GetDuration() *string {
	return s.Duration
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment) GetStart() *string {
	return s.Start
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment) SetDuration(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment {
	s.Duration = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment) SetStart(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment {
	s.Start = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobFpShotResultTextFpShotsTextFpShotTextFpShotSlicesTextFpShotSliceInputFragment) Validate() error {
	return dara.Validate(s)
}

type QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile struct {
	// The OSS bucket in which the job input resides.
	//
	// example:
	//
	// oss-test
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS region in which the job input resides.
	//
	// example:
	//
	// oss-cn-beijing
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The Object Storage Service (OSS) object that is used as the job input.
	//
	// example:
	//
	// test.mp4
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile) GetBucket() *string {
	return s.Bucket
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile) GetLocation() *string {
	return s.Location
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile) GetObject() *string {
	return s.Object
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile) SetBucket(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile {
	s.Bucket = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile) SetLocation(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile {
	s.Location = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile) SetObject(v string) *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile {
	s.Object = &v
	return s
}

func (s *QueryFpShotJobListResponseBodyFpShotJobListFpShotJobInputFile) Validate() error {
	return dara.Validate(s)
}

type QueryFpShotJobListResponseBodyNonExistIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s QueryFpShotJobListResponseBodyNonExistIds) String() string {
	return dara.Prettify(s)
}

func (s QueryFpShotJobListResponseBodyNonExistIds) GoString() string {
	return s.String()
}

func (s *QueryFpShotJobListResponseBodyNonExistIds) GetString_() []*string {
	return s.String_
}

func (s *QueryFpShotJobListResponseBodyNonExistIds) SetString_(v []*string) *QueryFpShotJobListResponseBodyNonExistIds {
	s.String_ = v
	return s
}

func (s *QueryFpShotJobListResponseBodyNonExistIds) Validate() error {
	return dara.Validate(s)
}
