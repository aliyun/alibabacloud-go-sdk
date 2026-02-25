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
	NextPageToken          *string                                                 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	NonExistSnapshotJobIds *QuerySnapshotJobListResponseBodyNonExistSnapshotJobIds `json:"NonExistSnapshotJobIds,omitempty" xml:"NonExistSnapshotJobIds,omitempty" type:"Struct"`
	// The ID of the snapshot job.
	//
	// example:
	//
	// 34BCAB31-2833-43A7-9FBD-B34302AB23EQ
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Code             *string                                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Count            *string                                                                     `json:"Count,omitempty" xml:"Count,omitempty"`
	CreationTime     *string                                                                     `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Id               *string                                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	Input            *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobInput            `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	MNSMessageResult *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobMNSMessageResult `json:"MNSMessageResult,omitempty" xml:"MNSMessageResult,omitempty" type:"Struct"`
	Message          *string                                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PipelineId       *string                                                                     `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	SnapshotConfig   *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfig   `json:"SnapshotConfig,omitempty" xml:"SnapshotConfig,omitempty" type:"Struct"`
	State            *string                                                                     `json:"State,omitempty" xml:"State,omitempty"`
	TileCount        *string                                                                     `json:"TileCount,omitempty" xml:"TileCount,omitempty"`
	UserData         *string                                                                     `json:"UserData,omitempty" xml:"UserData,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
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
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	MessageId    *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
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
	FrameType      *string                                                                                 `json:"FrameType,omitempty" xml:"FrameType,omitempty"`
	Height         *string                                                                                 `json:"Height,omitempty" xml:"Height,omitempty"`
	Interval       *string                                                                                 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Num            *string                                                                                 `json:"Num,omitempty" xml:"Num,omitempty"`
	OutputFile     *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigOutputFile     `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
	TileOut        *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOut        `json:"TileOut,omitempty" xml:"TileOut,omitempty" type:"Struct"`
	TileOutputFile *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTileOutputFile `json:"TileOutputFile,omitempty" xml:"TileOutputFile,omitempty" type:"Struct"`
	Time           *string                                                                                 `json:"Time,omitempty" xml:"Time,omitempty"`
	TimeArray      *QuerySnapshotJobListResponseBodySnapshotJobListSnapshotJobSnapshotConfigTimeArray      `json:"TimeArray,omitempty" xml:"TimeArray,omitempty" type:"Struct"`
	Width          *string                                                                                 `json:"Width,omitempty" xml:"Width,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
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
	CellHeight    *string `json:"CellHeight,omitempty" xml:"CellHeight,omitempty"`
	CellSelStep   *string `json:"CellSelStep,omitempty" xml:"CellSelStep,omitempty"`
	CellWidth     *string `json:"CellWidth,omitempty" xml:"CellWidth,omitempty"`
	Color         *string `json:"Color,omitempty" xml:"Color,omitempty"`
	Columns       *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	IsKeepCellPic *string `json:"IsKeepCellPic,omitempty" xml:"IsKeepCellPic,omitempty"`
	Lines         *string `json:"Lines,omitempty" xml:"Lines,omitempty"`
	Margin        *string `json:"Margin,omitempty" xml:"Margin,omitempty"`
	Padding       *string `json:"Padding,omitempty" xml:"Padding,omitempty"`
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
	Bucket   *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Object   *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RoleArn  *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
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
