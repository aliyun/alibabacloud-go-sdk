// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySnapshotJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *QuerySnapshotJobListResponseBody
	GetNextPageToken() *string
	SetNonExistSnapshotJobIds(v *QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds) *QuerySnapshotJobListResponseBody
	GetNonExistSnapshotJobIds() *QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds
	SetRequestId(v string) *QuerySnapshotJobListResponseBody
	GetRequestId() *string
	SetSnapshotJobList(v *QuerySnapshotJobListResponseBodySnapshotJobList) *QuerySnapshotJobListResponseBody
	GetSnapshotJobList() *QuerySnapshotJobListResponseBodySnapshotJobList
}

type QuerySnapshotJobListResponseBody struct {
	// The OSS object that is used as the input file.
	//
	// example:
	//
	// b11c171cced04565b1f38f1ecc39****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The OSS object that is generated as the output file of the tiling job.
	NonExistSnapshotJobIds *QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds `json:"NonExistSnapshotJobIds,omitempty" xml:"NonExistSnapshotJobIds,omitempty" type:"Struct"`
	// The ID of the snapshot job.
	//
	// example:
	//
	// 34BCAB31-2833-43A7-9FBD-B34302AB23EQ
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The distance between images.
	//
	// 	- Default value: **0**.
	//
	// 	- Unit: pixel.
	SnapshotJobList *QuerySnapshotJobListResponseBodySnapshotJobList `json:"SnapshotJobList,omitempty" xml:"SnapshotJobList,omitempty" type:"Struct"`
}

func (s QuerySnapshotJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *QuerySnapshotJobListResponseBody) GetNonExistSnapshotJobIds() *QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds {
	return s.NonExistSnapshotJobIds
}

func (s *QuerySnapshotJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySnapshotJobListResponseBody) GetSnapshotJobList() *QuerySnapshotJobListResponseBodySnapshotJobList {
	return s.SnapshotJobList
}

func (s *QuerySnapshotJobListResponseBody) SetNextPageToken(v string) *QuerySnapshotJobListResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *QuerySnapshotJobListResponseBody) SetNonExistSnapshotJobIds(v *QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds) *QuerySnapshotJobListResponseBody {
	s.NonExistSnapshotJobIds = v
	return s
}

func (s *QuerySnapshotJobListResponseBody) SetRequestId(v string) *QuerySnapshotJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySnapshotJobListResponseBody) SetSnapshotJobList(v *QuerySnapshotJobListResponseBodySnapshotJobList) *QuerySnapshotJobListResponseBody {
	s.SnapshotJobList = v
	return s
}

func (s *QuerySnapshotJobListResponseBody) Validate() error {
	if s.NonExistSnapshotJobIds != nil {
		if err := s.NonExistSnapshotJobIds.Validate(); err != nil {
			return err
		}
	}
	if s.SnapshotJobList != nil {
		if err := s.SnapshotJobList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds) GetString_() []*string {
	return s.String_
}

func (s *QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds) SetString_(v []*string) *QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds {
	s.String_ = v
	return s
}

func (s *QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds) Validate() error {
	return dara.Validate(s)
}

type QuerySnapshotJobListResponseBodySnapshotJobList struct {
	SnapshotJob []*QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob `json:"SnapshotJob,omitempty" xml:"SnapshotJob,omitempty" type:"Repeated"`
}

func (s QuerySnapshotJobListResponseBodySnapshotJobList) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListResponseBodySnapshotJobList) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobList) GetSnapshotJob() []*QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	return s.SnapshotJob
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobList) SetSnapshotJob(v []*QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) *QuerySnapshotJobListResponseBodySnapshotJobList {
	s.SnapshotJob = v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobList) Validate() error {
	if s.SnapshotJob != nil {
		for _, item := range s.SnapshotJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob struct {
	// You can call this operation to query up to 10 snapshot jobs at a time.
	//
	//
	// ## Limits on QPS
	//
	// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation. For more information, see [QPS limit](https://www.alibabacloud.com/help/en/apsaravideo-for-media-processing/latest/qps-limit).
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The stride of a single image.
	//
	// example:
	//
	// 2021-06-30T12:34:29Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The OSS output file of the tiling job.
	//
	// example:
	//
	// cc6cbef8e8d5481ca536f5d2a466****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of snapshots that are contained in the tiled image.
	Input *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The OSS object that is used as the input file.
	MNSMessageResult *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	// The ARN of the specified RAM role. Format: acs:ram::$accountID:role/$roleName.
	//
	// example:
	//
	// The resource operated InputFile is bad
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The start time for taking snapshots. Unit: milliseconds.
	//
	// example:
	//
	// b11c171cced04565b1f38f1ecc39****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The height of a single image. The default value is the height of the output snapshot.
	SnapshotConfig *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig `json:"SnapshotConfig,omitempty" xml:"SnapshotConfig,omitempty" type:"Struct"`
	// The information about the job input.
	//
	// example:
	//
	// Snapshoting
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The snapshot job IDs that do not exist. This parameter is not returned if all specified snapshot jobs are found.
	//
	// example:
	//
	// 7
	TileCount *string `json:"TileCount,omitempty" xml:"TileCount,omitempty"`
	// The token that is used to retrieve the next page of the query results. The value is a 32-bit UUID. If the returned query results cannot be displayed within one page, this parameter is returned. The value of this parameter is updated for each query.
	//
	// example:
	//
	// testid-001
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GetCode() *string {
	return s.Code
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GetCount() *string {
	return s.Count
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GetId() *string {
	return s.Id
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GetInput() *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput {
	return s.Input
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GetMNSMessageResult() *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult {
	return s.MNSMessageResult
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GetMessage() *string {
	return s.Message
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GetSnapshotConfig() *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig {
	return s.SnapshotConfig
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GetState() *string {
	return s.State
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GetTileCount() *string {
	return s.TileCount
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) GetUserData() *string {
	return s.UserData
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) SetCode(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	s.Code = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) SetCount(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	s.Count = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) SetCreationTime(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	s.CreationTime = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) SetId(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	s.Id = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) SetInput(v *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	s.Input = v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) SetMNSMessageResult(v *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	s.MNSMessageResult = v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) SetMessage(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	s.Message = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) SetPipelineId(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	s.PipelineId = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) SetSnapshotConfig(v *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	s.SnapshotConfig = v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) SetState(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	s.State = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) SetTileCount(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	s.TileCount = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) SetUserData(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob {
	s.UserData = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJob) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.MNSMessageResult != nil {
		if err := s.MNSMessageResult.Validate(); err != nil {
			return err
		}
	}
	if s.SnapshotConfig != nil {
		if err := s.SnapshotConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput struct {
	// The ID of the snapshot job.
	//
	// example:
	//
	// example
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The ID of the MPS queue to which the snapshot job was submitted.
	//
	// example:
	//
	// example-location
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The error code returned when the job fails. This parameter is not returned if the job is successfully processed.
	//
	// example:
	//
	// example.flv
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The custom data.
	//
	// example:
	//
	// acs:ram::1:role/testrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput) GetBucket() *string {
	return s.Bucket
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput) GetLocation() *string {
	return s.Location
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput) GetObject() *string {
	return s.Object
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput) GetRoleArn() *string {
	return s.RoleArn
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput) SetBucket(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput {
	s.Bucket = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput) SetLocation(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput {
	s.Location = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput) SetObject(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput {
	s.Object = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput) SetRoleArn(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput {
	s.RoleArn = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput) Validate() error {
	return dara.Validate(s)
}

type QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult struct {
	// The number of snapshots that were taken.
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The OSS bucket that stores the input file.
	//
	// example:
	//
	// The resource operated InputFile is bad
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the region in which the input OSS bucket is located.
	//
	// example:
	//
	// 799454621135656C7F815F198A76****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult) GetMessageId() *string {
	return s.MessageId
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult) SetErrorCode(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult {
	s.ErrorCode = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult) SetErrorMessage(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult {
	s.ErrorMessage = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult) SetMessageId(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult {
	s.MessageId = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult) Validate() error {
	return dara.Validate(s)
}

type QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig struct {
	// The ID of the region in which the output OSS bucket is located.
	//
	// example:
	//
	// intra
	FrameType *string `json:"FrameType,omitempty" xml:"FrameType,omitempty"`
	// The number of snapshots to take. If the Num parameter is set in the request, snapshots are taken at intervals.
	//
	// example:
	//
	// 8
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The OSS object that is generated as the output file of the snapshot job.
	//
	// example:
	//
	// 10
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The status of the snapshot job.
	//
	// - **Submitted**: The job was submitted.
	//
	// - **Snapshoting**: The job is being processed.
	//
	// - **Success**: The job was successfully processed.
	//
	// - **Fail**: The job failed.
	//
	// example:
	//
	// 10
	Num *string `json:"Num,omitempty" xml:"Num,omitempty"`
	// The OSS output file of the tiling job.
	OutputFile *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
	// The margin width of the tiled image.
	//
	// 	- Default value: **0**.
	//
	// 	- Unit: pixel.
	TileOut *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut `json:"TileOut,omitempty" xml:"TileOut,omitempty" type:"Struct"`
	// The error message returned when the job fails. This parameter is not returned if the job is successfully processed.
	TileOutputFile *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile `json:"TileOutputFile,omitempty" xml:"TileOutputFile,omitempty" type:"Struct"`
	// The width of a single image. The default value is the width of the output snapshot.
	//
	// example:
	//
	// 4
	Time      *string                                                                            `json:"Time,omitempty" xml:"Time,omitempty"`
	TimeArray *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTimeArray `json:"TimeArray,omitempty" xml:"TimeArray,omitempty" type:"Struct"`
	// The OSS bucket that stores the output file.
	//
	// example:
	//
	// 8
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) GetFrameType() *string {
	return s.FrameType
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) GetHeight() *string {
	return s.Height
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) GetInterval() *string {
	return s.Interval
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) GetNum() *string {
	return s.Num
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) GetOutputFile() *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile {
	return s.OutputFile
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) GetTileOut() *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut {
	return s.TileOut
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) GetTileOutputFile() *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile {
	return s.TileOutputFile
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) GetTime() *string {
	return s.Time
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) GetTimeArray() *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTimeArray {
	return s.TimeArray
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) GetWidth() *string {
	return s.Width
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) SetFrameType(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig {
	s.FrameType = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) SetHeight(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig {
	s.Height = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) SetInterval(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig {
	s.Interval = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) SetNum(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig {
	s.Num = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) SetOutputFile(v *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig {
	s.OutputFile = v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) SetTileOut(v *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig {
	s.TileOut = v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) SetTileOutputFile(v *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig {
	s.TileOutputFile = v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) SetTime(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig {
	s.Time = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) SetTimeArray(v *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTimeArray) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig {
	s.TimeArray = v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) SetWidth(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig {
	s.Width = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig) Validate() error {
	if s.OutputFile != nil {
		if err := s.OutputFile.Validate(); err != nil {
			return err
		}
	}
	if s.TileOut != nil {
		if err := s.TileOut.Validate(); err != nil {
			return err
		}
	}
	if s.TileOutputFile != nil {
		if err := s.TileOutputFile.Validate(); err != nil {
			return err
		}
	}
	if s.TimeArray != nil {
		if err := s.TimeArray.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile struct {
	// The OSS bucket that stores the output file.
	//
	// example:
	//
	// example
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The ID of the region in which the output OSS bucket is located.
	//
	// example:
	//
	// example-location
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The OSS object that is generated as the output file of the tiling job.
	//
	// example:
	//
	// example.png
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The ARN of the specified RAM role. Format: acs:ram::$accountID:role/$roleName.
	//
	// example:
	//
	// acs:ram::1:role/testrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile) GetBucket() *string {
	return s.Bucket
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile) GetLocation() *string {
	return s.Location
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile) GetObject() *string {
	return s.Object
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile) GetRoleArn() *string {
	return s.RoleArn
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile) SetBucket(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile {
	s.Bucket = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile) SetLocation(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile {
	s.Location = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile) SetObject(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile {
	s.Object = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile) SetRoleArn(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile {
	s.RoleArn = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile) Validate() error {
	return dara.Validate(s)
}

type QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut struct {
	// The interval for taking snapshots.
	//
	// 	- If this Interval parameter is specified in the request, snapshots are taken at intervals. The value must be greater than 0.
	//
	// 	- Unit: seconds.
	//
	// 	- Default value: **10**.
	//
	// example:
	//
	// 8
	CellHeight *string `json:"CellHeight,omitempty" xml:"CellHeight,omitempty"`
	// The number of rows that the tiled image can contain. Default value: **10**.
	//
	// example:
	//
	// 3
	CellSelStep *string `json:"CellSelStep,omitempty" xml:"CellSelStep,omitempty"`
	// The type of the snapshot. Valid values:
	//
	// 	- **normal**: normal frames.
	//
	// 	- **intra**: I-frames.
	//
	// 	- Default value: **intra**.
	//
	// example:
	//
	// 8
	CellWidth *string `json:"CellWidth,omitempty" xml:"CellWidth,omitempty"`
	// Indicates whether the single images are retained. Default value: **true**.
	//
	// example:
	//
	// black
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The height of the output snapshot.
	//
	// example:
	//
	// 10
	Columns *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// The Object Storage Service (OSS) output file of the snapshot job.
	//
	// example:
	//
	// false
	IsKeepCellPic *string `json:"IsKeepCellPic,omitempty" xml:"IsKeepCellPic,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the specified RAM role. Format: acs:ram::$accountID:role/$roleName.
	//
	// example:
	//
	// 10
	Lines *string `json:"Lines,omitempty" xml:"Lines,omitempty"`
	// The width of the output snapshot.
	//
	// example:
	//
	// 0
	Margin *string `json:"Margin,omitempty" xml:"Margin,omitempty"`
	// The number of columns that the tiled image can contain. Default value: **10**.
	//
	// example:
	//
	// 0
	Padding *string `json:"Padding,omitempty" xml:"Padding,omitempty"`
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) GetCellHeight() *string {
	return s.CellHeight
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) GetCellSelStep() *string {
	return s.CellSelStep
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) GetCellWidth() *string {
	return s.CellWidth
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) GetColor() *string {
	return s.Color
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) GetColumns() *string {
	return s.Columns
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) GetIsKeepCellPic() *string {
	return s.IsKeepCellPic
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) GetLines() *string {
	return s.Lines
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) GetMargin() *string {
	return s.Margin
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) GetPadding() *string {
	return s.Padding
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) SetCellHeight(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut {
	s.CellHeight = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) SetCellSelStep(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut {
	s.CellSelStep = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) SetCellWidth(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut {
	s.CellWidth = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) SetColor(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut {
	s.Color = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) SetColumns(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut {
	s.Columns = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) SetIsKeepCellPic(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut {
	s.IsKeepCellPic = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) SetLines(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut {
	s.Lines = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) SetMargin(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut {
	s.Margin = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) SetPadding(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut {
	s.Padding = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut) Validate() error {
	return dara.Validate(s)
}

type QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile struct {
	// The error code returned when the job fails. This parameter is not returned if the job is successfully processed.
	//
	// example:
	//
	// example
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The error message returned when the job fails. This parameter is not returned if the job is successfully processed.
	//
	// example:
	//
	// example-location
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The ID of the message. This parameter is not returned if the job fails.
	//
	// example:
	//
	// example.png
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The message sent by MNS to notify the user of the job result.
	//
	// example:
	//
	// acs:ram::1:role/testrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile) GetBucket() *string {
	return s.Bucket
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile) GetLocation() *string {
	return s.Location
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile) GetObject() *string {
	return s.Object
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile) GetRoleArn() *string {
	return s.RoleArn
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile) SetBucket(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile {
	s.Bucket = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile) SetLocation(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile {
	s.Location = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile) SetObject(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile {
	s.Object = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile) SetRoleArn(v string) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile {
	s.RoleArn = &v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile) Validate() error {
	return dara.Validate(s)
}

type QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTimeArray struct {
	TimePointList []*int64 `json:"TimePointList,omitempty" xml:"TimePointList,omitempty" type:"Repeated"`
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTimeArray) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTimeArray) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTimeArray) GetTimePointList() []*int64 {
	return s.TimePointList
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTimeArray) SetTimePointList(v []*int64) *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTimeArray {
	s.TimePointList = v
	return s
}

func (s *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTimeArray) Validate() error {
	return dara.Validate(s)
}
