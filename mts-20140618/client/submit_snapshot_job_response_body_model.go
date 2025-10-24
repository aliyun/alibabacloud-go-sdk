// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSnapshotJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitSnapshotJobResponseBody
	GetRequestId() *string
	SetSnapshotJob(v *SubmitSnapshotJobResponseBodySnapshotJob) *SubmitSnapshotJobResponseBody
	GetSnapshotJob() *SubmitSnapshotJobResponseBodySnapshotJob
}

type SubmitSnapshotJobResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 19B6D8C5-A5DD-467A-B435-29D393C71E2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the snapshot job.
	SnapshotJob *SubmitSnapshotJobResponseBodySnapshotJob `json:"SnapshotJob,omitempty" xml:"SnapshotJob,omitempty" type:"Struct"`
}

func (s SubmitSnapshotJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSnapshotJobResponseBody) GetSnapshotJob() *SubmitSnapshotJobResponseBodySnapshotJob {
	return s.SnapshotJob
}

func (s *SubmitSnapshotJobResponseBody) SetRequestId(v string) *SubmitSnapshotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSnapshotJobResponseBody) SetSnapshotJob(v *SubmitSnapshotJobResponseBodySnapshotJob) *SubmitSnapshotJobResponseBody {
	s.SnapshotJob = v
	return s
}

func (s *SubmitSnapshotJobResponseBody) Validate() error {
	if s.SnapshotJob != nil {
		if err := s.SnapshotJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitSnapshotJobResponseBodySnapshotJob struct {
	// The error code returned if the job fails. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// ResourceContentBad
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of snapshots that are captured.
	//
	// example:
	//
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2021-05-19T03:11:48Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the snapshot job.
	//
	// example:
	//
	// f4e3b9ba9f3840c39d6e288056f0****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the job input.
	Input *SubmitSnapshotJobResponseBodySnapshotJobInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The message sent by MNS to notify the user of the job result.
	MNSMessageResult *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	// The error message returned if the job fails. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// The resource operated InputFile is bad
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the MPS queue to which the snapshot job is submitted.
	//
	// example:
	//
	// dd3dae411e704030b921e52698e5****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The snapshot configurations.
	SnapshotConfig *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig `json:"SnapshotConfig,omitempty" xml:"SnapshotConfig,omitempty" type:"Struct"`
	// The status of the snapshot job. Valid values:
	//
	// 	- **Submitted**: The job is submitted.
	//
	// 	- **Snapshoting**: The job is being processed.
	//
	// 	- **Success**: The job is successful.
	//
	// 	- **Fail**: The job fails.
	//
	// example:
	//
	// Snapshoting
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The number of single images that are contained in the tiled image.
	//
	// example:
	//
	// 5
	TileCount *string `json:"TileCount,omitempty" xml:"TileCount,omitempty"`
	// The custom data.
	//
	// example:
	//
	// testid-001
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSnapshotJobResponseBodySnapshotJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponseBodySnapshotJob) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetCode() *string {
	return s.Code
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetCount() *string {
	return s.Count
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetId() *string {
	return s.Id
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetInput() *SubmitSnapshotJobResponseBodySnapshotJobInput {
	return s.Input
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetMNSMessageResult() *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult {
	return s.MNSMessageResult
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetMessage() *string {
	return s.Message
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetSnapshotConfig() *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig {
	return s.SnapshotConfig
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetState() *string {
	return s.State
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetTileCount() *string {
	return s.TileCount
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetCode(v string) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.Code = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetCount(v string) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.Count = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetCreationTime(v string) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.CreationTime = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetId(v string) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.Id = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetInput(v *SubmitSnapshotJobResponseBodySnapshotJobInput) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.Input = v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetMNSMessageResult(v *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.MNSMessageResult = v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetMessage(v string) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.Message = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetPipelineId(v string) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.PipelineId = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetSnapshotConfig(v *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.SnapshotConfig = v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetState(v string) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.State = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetTileCount(v string) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.TileCount = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetUserData(v string) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.UserData = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) Validate() error {
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

type SubmitSnapshotJobResponseBodySnapshotJobInput struct {
	// The OSS bucket that stores the object.
	//
	// example:
	//
	// example
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The region in which the OSS bucket resides.
	//
	// example:
	//
	// example-location\\"
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The OSS object that is used as the input file.
	//
	// example:
	//
	// example.flv
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The ARN of the specified RAM role. Format: acs:ram::$accountID:role/$roleName.
	//
	// example:
	//
	// acs:ram::1:role/testrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s SubmitSnapshotJobResponseBodySnapshotJobInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponseBodySnapshotJobInput) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobInput) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobInput) GetLocation() *string {
	return s.Location
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobInput) GetObject() *string {
	return s.Object
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobInput) GetRoleArn() *string {
	return s.RoleArn
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobInput) SetBucket(v string) *SubmitSnapshotJobResponseBodySnapshotJobInput {
	s.Bucket = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobInput) SetLocation(v string) *SubmitSnapshotJobResponseBodySnapshotJobInput {
	s.Location = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobInput) SetObject(v string) *SubmitSnapshotJobResponseBodySnapshotJobInput {
	s.Object = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobInput) SetRoleArn(v string) *SubmitSnapshotJobResponseBodySnapshotJobInput {
	s.RoleArn = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobInput) Validate() error {
	return dara.Validate(s)
}

type SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult struct {
	// The error code returned if the job fails. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the job fails. This parameter is not returned if the job is successful.
	//
	// example:
	//
	// The resource operated InputFile is bad
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the message. This parameter is not returned if the job fails.
	//
	// example:
	//
	// 799454621135656C7F815F198A76****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult) GetMessageId() *string {
	return s.MessageId
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult) SetErrorCode(v string) *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult {
	s.ErrorCode = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult) SetErrorMessage(v string) *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult) SetMessageId(v string) *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult {
	s.MessageId = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobMNSMessageResult) Validate() error {
	return dara.Validate(s)
}

type SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig struct {
	// The type of the snapshot. Default value: **Normal**. Valid values:
	//
	// 	- **normal**: normal frames.
	//
	// 	- **intra**: I-frames (keyframes).
	//
	// > If the FrameType parameter is set to intra in the request, only keyframes are captured. If no keyframe is found at the specified point in time, the keyframe closest to the specified point in time is captured. Keyframes are captured faster than normal frames if the same snapshot rules are applied.
	//
	// example:
	//
	// intra
	FrameType *string `json:"FrameType,omitempty" xml:"FrameType,omitempty"`
	// The height of the output snapshot.
	//
	// example:
	//
	// 8
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The interval for capturing snapshots.
	//
	// 	- If this parameter is specified in the request, snapshots are captured at intervals. The value must be greater than 0 in the request.
	//
	// 	- Unit: seconds.
	//
	// 	- Default value: **10**.
	//
	// example:
	//
	// 20
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The number of snapshots. If the Num parameter is set in the request, snapshots are captured at intervals.
	//
	// example:
	//
	// 10
	Num *string `json:"Num,omitempty" xml:"Num,omitempty"`
	// The information about the output file of the snapshot job.
	OutputFile *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
	// The tiling configurations.
	TileOut *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut `json:"TileOut,omitempty" xml:"TileOut,omitempty" type:"Struct"`
	// The information about the output file of the tiling job.
	TileOutputFile *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile `json:"TileOutputFile,omitempty" xml:"TileOutputFile,omitempty" type:"Struct"`
	// The start time for capturing snapshots. Unit: milliseconds.
	//
	// example:
	//
	// 5
	Time      *string                                                          `json:"Time,omitempty" xml:"Time,omitempty"`
	TimeArray *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTimeArray `json:"TimeArray,omitempty" xml:"TimeArray,omitempty" type:"Struct"`
	// The width of the output snapshot.
	//
	// example:
	//
	// 8
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) GetFrameType() *string {
	return s.FrameType
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) GetHeight() *string {
	return s.Height
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) GetInterval() *string {
	return s.Interval
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) GetNum() *string {
	return s.Num
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) GetOutputFile() *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile {
	return s.OutputFile
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) GetTileOut() *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut {
	return s.TileOut
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) GetTileOutputFile() *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile {
	return s.TileOutputFile
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) GetTime() *string {
	return s.Time
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) GetTimeArray() *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTimeArray {
	return s.TimeArray
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) GetWidth() *string {
	return s.Width
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) SetFrameType(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig {
	s.FrameType = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) SetHeight(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig {
	s.Height = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) SetInterval(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig {
	s.Interval = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) SetNum(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig {
	s.Num = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) SetOutputFile(v *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig {
	s.OutputFile = v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) SetTileOut(v *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig {
	s.TileOut = v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) SetTileOutputFile(v *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig {
	s.TileOutputFile = v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) SetTime(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig {
	s.Time = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) SetTimeArray(v *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTimeArray) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig {
	s.TimeArray = v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) SetWidth(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig {
	s.Width = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfig) Validate() error {
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

type SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile struct {
	// The OSS bucket that stores the output snapshot.
	//
	// example:
	//
	// example
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS region in which the OSS bucket for storing the output snapshot resides.
	//
	// example:
	//
	// example-location
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The OSS object that is generated as the output file of the snapshot job.
	//
	// example:
	//
	// test.png
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the specified RAM role. Format: acs:ram::$accountID:role/$roleName.
	//
	// example:
	//
	// acs:ram::1:role/testrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile) GetLocation() *string {
	return s.Location
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile) GetObject() *string {
	return s.Object
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile) GetRoleArn() *string {
	return s.RoleArn
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile) SetBucket(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile {
	s.Bucket = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile) SetLocation(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile {
	s.Location = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile) SetObject(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile {
	s.Object = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile) SetRoleArn(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile {
	s.RoleArn = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigOutputFile) Validate() error {
	return dara.Validate(s)
}

type SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut struct {
	// The height of a single image. The default value is the height of the output snapshot.
	//
	// example:
	//
	// 100
	CellHeight *string `json:"CellHeight,omitempty" xml:"CellHeight,omitempty"`
	// The step for selecting a single image.
	//
	// example:
	//
	// 3
	CellSelStep *string `json:"CellSelStep,omitempty" xml:"CellSelStep,omitempty"`
	// The width of a single image. The default value is the width of the output snapshot.
	//
	// example:
	//
	// 100
	CellWidth *string `json:"CellWidth,omitempty" xml:"CellWidth,omitempty"`
	// The background color.
	//
	// 	- Default value: **black**.
	//
	// 	- You can set the Color parameter to a **color keyword*	- or **random*	- in the request.
	//
	// > If you want to set the background color to black, you can specify the color keyword in one of the following three formats: Black, black, and #000000.
	//
	// example:
	//
	// black
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The number of columns that the tiled image contains. Default value: **10**.
	//
	// example:
	//
	// 10
	Columns *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// Indicates whether the single images are retained. Valid values:
	//
	// 	- **true**: The single images are retained.
	//
	// 	- **false**: The single images are not retained.
	//
	// 	- Default value: **true**.
	//
	// example:
	//
	// false
	IsKeepCellPic *string `json:"IsKeepCellPic,omitempty" xml:"IsKeepCellPic,omitempty"`
	// The number of rows that the tiled image contains. Default value: **10**.
	//
	// example:
	//
	// 10
	Lines *string `json:"Lines,omitempty" xml:"Lines,omitempty"`
	// The margin width of the tiled image.
	//
	// 	- Default value: **0**.
	//
	// 	- Unit: pixel.
	//
	// example:
	//
	// 5
	Margin *string `json:"Margin,omitempty" xml:"Margin,omitempty"`
	// The distance between two consecutive single images in the tiled image.
	//
	// 	- Default value: **0**.
	//
	// 	- Unit: pixel.
	//
	// example:
	//
	// 0
	Padding *string `json:"Padding,omitempty" xml:"Padding,omitempty"`
}

func (s SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) GetCellHeight() *string {
	return s.CellHeight
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) GetCellSelStep() *string {
	return s.CellSelStep
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) GetCellWidth() *string {
	return s.CellWidth
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) GetColor() *string {
	return s.Color
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) GetColumns() *string {
	return s.Columns
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) GetIsKeepCellPic() *string {
	return s.IsKeepCellPic
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) GetLines() *string {
	return s.Lines
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) GetMargin() *string {
	return s.Margin
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) GetPadding() *string {
	return s.Padding
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) SetCellHeight(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut {
	s.CellHeight = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) SetCellSelStep(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut {
	s.CellSelStep = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) SetCellWidth(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut {
	s.CellWidth = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) SetColor(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut {
	s.Color = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) SetColumns(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut {
	s.Columns = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) SetIsKeepCellPic(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut {
	s.IsKeepCellPic = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) SetLines(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut {
	s.Lines = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) SetMargin(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut {
	s.Margin = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) SetPadding(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut {
	s.Padding = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOut) Validate() error {
	return dara.Validate(s)
}

type SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile struct {
	// The OSS bucket that stores the object.
	//
	// example:
	//
	// example
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The ID of the region in which the OSS bucket that stores the object is located.
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

func (s SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile) GetBucket() *string {
	return s.Bucket
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile) GetLocation() *string {
	return s.Location
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile) GetObject() *string {
	return s.Object
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile) GetRoleArn() *string {
	return s.RoleArn
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile) SetBucket(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile {
	s.Bucket = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile) SetLocation(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile {
	s.Location = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile) SetObject(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile {
	s.Object = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile) SetRoleArn(v string) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile {
	s.RoleArn = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTileOutputFile) Validate() error {
	return dara.Validate(s)
}

type SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTimeArray struct {
	TimePointList []*int64 `json:"TimePointList,omitempty" xml:"TimePointList,omitempty" type:"Repeated"`
}

func (s SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTimeArray) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTimeArray) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTimeArray) GetTimePointList() []*int64 {
	return s.TimePointList
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTimeArray) SetTimePointList(v []*int64) *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTimeArray {
	s.TimePointList = v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJobSnapshotConfigTimeArray) Validate() error {
	return dara.Validate(s)
}
