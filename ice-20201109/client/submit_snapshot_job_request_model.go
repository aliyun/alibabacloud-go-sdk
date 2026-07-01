// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSnapshotJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *SubmitSnapshotJobRequestInput) *SubmitSnapshotJobRequest
	GetInput() *SubmitSnapshotJobRequestInput
	SetName(v string) *SubmitSnapshotJobRequest
	GetName() *string
	SetOutput(v *SubmitSnapshotJobRequestOutput) *SubmitSnapshotJobRequest
	GetOutput() *SubmitSnapshotJobRequestOutput
	SetScheduleConfig(v *SubmitSnapshotJobRequestScheduleConfig) *SubmitSnapshotJobRequest
	GetScheduleConfig() *SubmitSnapshotJobRequestScheduleConfig
	SetTemplateConfig(v *SubmitSnapshotJobRequestTemplateConfig) *SubmitSnapshotJobRequest
	GetTemplateConfig() *SubmitSnapshotJobRequestTemplateConfig
	SetUserData(v string) *SubmitSnapshotJobRequest
	GetUserData() *string
}

type SubmitSnapshotJobRequest struct {
	// The input for the snapshot job.
	//
	// This parameter is required.
	Input *SubmitSnapshotJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The name of the snapshot job.
	//
	// example:
	//
	// SampleJob
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output destination for the snapshot job.
	//
	// This parameter is required.
	Output *SubmitSnapshotJobRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The scheduling configuration.
	ScheduleConfig *SubmitSnapshotJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The snapshot template configuration.
	//
	// This parameter is required.
	TemplateConfig *SubmitSnapshotJobRequestTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty" type:"Struct"`
	// Custom user data, passed as a JSON-formatted string.
	//
	// example:
	//
	// {"test parameter": "test value"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSnapshotJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobRequest) GetInput() *SubmitSnapshotJobRequestInput {
	return s.Input
}

func (s *SubmitSnapshotJobRequest) GetName() *string {
	return s.Name
}

func (s *SubmitSnapshotJobRequest) GetOutput() *SubmitSnapshotJobRequestOutput {
	return s.Output
}

func (s *SubmitSnapshotJobRequest) GetScheduleConfig() *SubmitSnapshotJobRequestScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitSnapshotJobRequest) GetTemplateConfig() *SubmitSnapshotJobRequestTemplateConfig {
	return s.TemplateConfig
}

func (s *SubmitSnapshotJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSnapshotJobRequest) SetInput(v *SubmitSnapshotJobRequestInput) *SubmitSnapshotJobRequest {
	s.Input = v
	return s
}

func (s *SubmitSnapshotJobRequest) SetName(v string) *SubmitSnapshotJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetOutput(v *SubmitSnapshotJobRequestOutput) *SubmitSnapshotJobRequest {
	s.Output = v
	return s
}

func (s *SubmitSnapshotJobRequest) SetScheduleConfig(v *SubmitSnapshotJobRequestScheduleConfig) *SubmitSnapshotJobRequest {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitSnapshotJobRequest) SetTemplateConfig(v *SubmitSnapshotJobRequestTemplateConfig) *SubmitSnapshotJobRequest {
	s.TemplateConfig = v
	return s
}

func (s *SubmitSnapshotJobRequest) SetUserData(v string) *SubmitSnapshotJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSnapshotJobRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateConfig != nil {
		if err := s.TemplateConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitSnapshotJobRequestInput struct {
	// The input media asset.
	//
	// - If `Type` is `OSS`, specify the OSS URL of the input file.
	//
	// - If `Type` is `Media`, specify the ID of the media asset.
	//
	// The OSS URL must be in one of the following formats:
	//
	// 1. `oss://bucket/object`
	//
	// 2. `http(s)://bucket.oss-[RegionId].aliyuncs.com/object`
	//
	//    <br>In these formats, `bucket` is the name of an OSS bucket located in the same region as the current project, and `object` is the file path.<br>
	//
	// > The OSS bucket specified in the URL must be added to IMS [storage management](https://help.aliyun.com/document_detail/609918.html) before use.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket/object.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input. Valid values:
	//
	// - `OSS`: an OSS file URL.
	//
	// - `Media`: a media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitSnapshotJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitSnapshotJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitSnapshotJobRequestInput) SetMedia(v string) *SubmitSnapshotJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitSnapshotJobRequestInput) SetType(v string) *SubmitSnapshotJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitSnapshotJobRequestInput) Validate() error {
	return dara.Validate(s)
}

type SubmitSnapshotJobRequestOutput struct {
	// The output media asset.
	//
	// - If `Type` is `OSS`, specify the OSS URL for the output file.
	//
	// - If `Type` is `Media`, specify the ID of the output media asset.
	//
	// The OSS URL must be in one of the following formats:
	//
	// 1. `oss://bucket/object`
	//
	// 2. `http(s)://bucket.oss-[RegionId].aliyuncs.com/object`
	//
	// In these formats, `bucket` is the name of an OSS bucket located in the same region as the current project, and `object` is the file path.
	//
	// - When capturing multiple static snapshots, the `object` must contain the `{Count}` placeholder.
	//
	// - When capturing a sprite, the `object` must contain the `{TileCount}` placeholder.
	//
	// - For WebVTT snapshots, the filename in the `object` path must end with `.vtt`.
	//
	// > The OSS bucket specified in the URL must be added to IMS [storage management](https://help.aliyun.com/document_detail/609918.html) before use.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/output-{Count}.jpg
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the output. Valid values:
	//
	// - `OSS`: an OSS file URL.
	//
	// - `Media`: a media asset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitSnapshotJobRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobRequestOutput) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobRequestOutput) GetMedia() *string {
	return s.Media
}

func (s *SubmitSnapshotJobRequestOutput) GetType() *string {
	return s.Type
}

func (s *SubmitSnapshotJobRequestOutput) SetMedia(v string) *SubmitSnapshotJobRequestOutput {
	s.Media = &v
	return s
}

func (s *SubmitSnapshotJobRequestOutput) SetType(v string) *SubmitSnapshotJobRequestOutput {
	s.Type = &v
	return s
}

func (s *SubmitSnapshotJobRequestOutput) Validate() error {
	return dara.Validate(s)
}

type SubmitSnapshotJobRequestScheduleConfig struct {
	// The pipeline ID.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s SubmitSnapshotJobRequestScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobRequestScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobRequestScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitSnapshotJobRequestScheduleConfig) SetPipelineId(v string) *SubmitSnapshotJobRequestScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitSnapshotJobRequestScheduleConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitSnapshotJobRequestTemplateConfig struct {
	// Parameters used to override settings in the specified template.
	OverwriteParams *SubmitSnapshotJobRequestTemplateConfigOverwriteParams `json:"OverwriteParams,omitempty" xml:"OverwriteParams,omitempty" type:"Struct"`
	// The snapshot template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitSnapshotJobRequestTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobRequestTemplateConfig) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobRequestTemplateConfig) GetOverwriteParams() *SubmitSnapshotJobRequestTemplateConfigOverwriteParams {
	return s.OverwriteParams
}

func (s *SubmitSnapshotJobRequestTemplateConfig) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitSnapshotJobRequestTemplateConfig) SetOverwriteParams(v *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) *SubmitSnapshotJobRequestTemplateConfig {
	s.OverwriteParams = v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfig) SetTemplateId(v string) *SubmitSnapshotJobRequestTemplateConfig {
	s.TemplateId = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfig) Validate() error {
	if s.OverwriteParams != nil {
		if err := s.OverwriteParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitSnapshotJobRequestTemplateConfigOverwriteParams struct {
	// The threshold for detecting and filtering black content in the first frame. This applies only to multi-frame snapshots.
	//
	// example:
	//
	// 30
	BlackLevel *int32 `json:"BlackLevel,omitempty" xml:"BlackLevel,omitempty"`
	// The number of snapshots to capture.
	//
	// example:
	//
	// 5
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The frame type.
	//
	// example:
	//
	// intra
	FrameType *string `json:"FrameType,omitempty" xml:"FrameType,omitempty"`
	// The output image height.
	//
	// example:
	//
	// 480
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The interval between snapshots.
	//
	// example:
	//
	// 10
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Specifies whether to stitch snapshots into a single sprite. This applies only to WebVTT snapshots.
	//
	// example:
	//
	// true
	IsSptFrag *bool `json:"IsSptFrag,omitempty" xml:"IsSptFrag,omitempty"`
	// The threshold for determining whether a pixel is black.
	//
	// example:
	//
	// 70
	PixelBlackThreshold *int32 `json:"PixelBlackThreshold,omitempty" xml:"PixelBlackThreshold,omitempty"`
	// The sprite configuration.
	SpriteSnapshotConfig *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig `json:"SpriteSnapshotConfig,omitempty" xml:"SpriteSnapshotConfig,omitempty" type:"Struct"`
	// The start time for capturing snapshots.
	//
	// example:
	//
	// 1000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The snapshot type.
	//
	// example:
	//
	// Sprite
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The output image width.
	//
	// example:
	//
	// 720
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitSnapshotJobRequestTemplateConfigOverwriteParams) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobRequestTemplateConfigOverwriteParams) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) GetBlackLevel() *int32 {
	return s.BlackLevel
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) GetCount() *int64 {
	return s.Count
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) GetFrameType() *string {
	return s.FrameType
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) GetHeight() *int32 {
	return s.Height
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) GetInterval() *int64 {
	return s.Interval
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) GetIsSptFrag() *bool {
	return s.IsSptFrag
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) GetPixelBlackThreshold() *int32 {
	return s.PixelBlackThreshold
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) GetSpriteSnapshotConfig() *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig {
	return s.SpriteSnapshotConfig
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) GetTime() *int64 {
	return s.Time
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) GetType() *string {
	return s.Type
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) GetWidth() *int32 {
	return s.Width
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) SetBlackLevel(v int32) *SubmitSnapshotJobRequestTemplateConfigOverwriteParams {
	s.BlackLevel = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) SetCount(v int64) *SubmitSnapshotJobRequestTemplateConfigOverwriteParams {
	s.Count = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) SetFrameType(v string) *SubmitSnapshotJobRequestTemplateConfigOverwriteParams {
	s.FrameType = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) SetHeight(v int32) *SubmitSnapshotJobRequestTemplateConfigOverwriteParams {
	s.Height = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) SetInterval(v int64) *SubmitSnapshotJobRequestTemplateConfigOverwriteParams {
	s.Interval = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) SetIsSptFrag(v bool) *SubmitSnapshotJobRequestTemplateConfigOverwriteParams {
	s.IsSptFrag = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) SetPixelBlackThreshold(v int32) *SubmitSnapshotJobRequestTemplateConfigOverwriteParams {
	s.PixelBlackThreshold = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) SetSpriteSnapshotConfig(v *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) *SubmitSnapshotJobRequestTemplateConfigOverwriteParams {
	s.SpriteSnapshotConfig = v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) SetTime(v int64) *SubmitSnapshotJobRequestTemplateConfigOverwriteParams {
	s.Time = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) SetType(v string) *SubmitSnapshotJobRequestTemplateConfigOverwriteParams {
	s.Type = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) SetWidth(v int32) *SubmitSnapshotJobRequestTemplateConfigOverwriteParams {
	s.Width = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParams) Validate() error {
	if s.SpriteSnapshotConfig != nil {
		if err := s.SpriteSnapshotConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig struct {
	// The height of each tile. Default: the height of the output snapshot.
	//
	// example:
	//
	// 480
	CellHeight *int32 `json:"CellHeight,omitempty" xml:"CellHeight,omitempty"`
	// The width of each tile. Default: the width of the output snapshot.
	//
	// example:
	//
	// 720
	CellWidth *int32 `json:"CellWidth,omitempty" xml:"CellWidth,omitempty"`
	// The background color.
	//
	// example:
	//
	// #000000
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The number of columns in the sprite grid.
	//
	// example:
	//
	// 20
	Columns *int32 `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// The number of rows in the sprite grid.
	//
	// example:
	//
	// 20
	Lines *int32 `json:"Lines,omitempty" xml:"Lines,omitempty"`
	// The margin around the sprite, in pixels. Default value: 0.
	//
	// example:
	//
	// 20
	Margin *int32 `json:"Margin,omitempty" xml:"Margin,omitempty"`
	// The padding between tiles, in pixels. Default value: 0.
	//
	// example:
	//
	// 20
	Padding *int32 `json:"Padding,omitempty" xml:"Padding,omitempty"`
}

func (s SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) GetCellHeight() *int32 {
	return s.CellHeight
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) GetCellWidth() *int32 {
	return s.CellWidth
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) GetColor() *string {
	return s.Color
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) GetColumns() *int32 {
	return s.Columns
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) GetLines() *int32 {
	return s.Lines
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) GetMargin() *int32 {
	return s.Margin
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) GetPadding() *int32 {
	return s.Padding
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) SetCellHeight(v int32) *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig {
	s.CellHeight = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) SetCellWidth(v int32) *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig {
	s.CellWidth = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) SetColor(v string) *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig {
	s.Color = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) SetColumns(v int32) *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig {
	s.Columns = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) SetLines(v int32) *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig {
	s.Lines = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) SetMargin(v int32) *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig {
	s.Margin = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) SetPadding(v int32) *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig {
	s.Padding = &v
	return s
}

func (s *SubmitSnapshotJobRequestTemplateConfigOverwriteParamsSpriteSnapshotConfig) Validate() error {
	return dara.Validate(s)
}
