// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocParserJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *CreateDocParserJobRequest
	GetAgentName() *string
	SetAsrLanguage(v string) *CreateDocParserJobRequest
	GetAsrLanguage() *string
	SetAudioClipOutput(v bool) *CreateDocParserJobRequest
	GetAudioClipOutput() *bool
	SetAudioWindowSeconds(v int32) *CreateDocParserJobRequest
	GetAudioWindowSeconds() *int32
	SetChunkSummary(v bool) *CreateDocParserJobRequest
	GetChunkSummary() *bool
	SetFileFormat(v string) *CreateDocParserJobRequest
	GetFileFormat() *string
	SetFileName(v string) *CreateDocParserJobRequest
	GetFileName() *string
	SetFileUrl(v string) *CreateDocParserJobRequest
	GetFileUrl() *string
	SetFrameOutput(v bool) *CreateDocParserJobRequest
	GetFrameOutput() *bool
	SetGlobalSummary(v bool) *CreateDocParserJobRequest
	GetGlobalSummary() *bool
	SetImageMode(v string) *CreateDocParserJobRequest
	GetImageMode() *string
	SetImageUnderstanding(v string) *CreateDocParserJobRequest
	GetImageUnderstanding() *string
	SetMediaChunkIntervalSeconds(v int32) *CreateDocParserJobRequest
	GetMediaChunkIntervalSeconds() *int32
	SetMediaChunkStrategy(v string) *CreateDocParserJobRequest
	GetMediaChunkStrategy() *string
	SetMediaFramesPerMinute(v float64) *CreateDocParserJobRequest
	GetMediaFramesPerMinute() *float64
	SetMediaMaxFrameBudget(v int32) *CreateDocParserJobRequest
	GetMediaMaxFrameBudget() *int32
	SetMediaMinFrameBudget(v int32) *CreateDocParserJobRequest
	GetMediaMinFrameBudget() *int32
	SetOssFileUrl(v string) *CreateDocParserJobRequest
	GetOssFileUrl() *string
	SetOutputFormat(v string) *CreateDocParserJobRequest
	GetOutputFormat() *string
	SetParseScene(v string) *CreateDocParserJobRequest
	GetParseScene() *string
	SetRegionId(v string) *CreateDocParserJobRequest
	GetRegionId() *string
	SetResponseMode(v string) *CreateDocParserJobRequest
	GetResponseMode() *string
	SetResultType(v string) *CreateDocParserJobRequest
	GetResultType() *string
	SetTableFormat(v string) *CreateDocParserJobRequest
	GetTableFormat() *string
}

type CreateDocParserJobRequest struct {
	// The agent name.
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The language type for speech recognition.
	AsrLanguage *string `json:"AsrLanguage,omitempty" xml:"AsrLanguage,omitempty"`
	// The audio clip output.
	AudioClipOutput *bool `json:"AudioClipOutput,omitempty" xml:"AudioClipOutput,omitempty"`
	// The audio window duration in seconds.
	AudioWindowSeconds *int32 `json:"AudioWindowSeconds,omitempty" xml:"AudioWindowSeconds,omitempty"`
	// The chunk summary information.
	ChunkSummary *bool `json:"ChunkSummary,omitempty" xml:"ChunkSummary,omitempty"`
	// The format of the input file. Valid values:
	//
	// - **pdf**: PDF file.
	//
	// - **docx**: Word file in docx format.
	//
	// - **doc**: Word file in doc format.
	//
	// - **pptx**: PPT file in pptx format.
	//
	// - **ppt**: PPT file in ppt format.
	//
	// - **txt**: Plain text file.
	//
	// - **md**: Markdown file.
	//
	// - **png**: PNG image.
	//
	// - **jpg**: JPG image.
	//
	// - **jpeg**: JPEG image.
	//
	// This parameter is required.
	//
	// example:
	//
	// pdf
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	// The file name, which must include the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// document.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The HTTP or HTTPS URL of the file to be parsed.
	//
	// >SDKs for various programming languages additionally provide a `CreateDocParserJobAdvance` method that supports passing a local file stream directly (such as Java InputStream), without the need to upload the file to OSS and construct a FileUrl in advance. When using the Advance method, replace the `FileUrl` parameter (URL string) with the `FileUrlObject` parameter (file stream). All other request parameters remain unchanged. The SDK automatically performs the following operations:
	//
	// >1. Obtains temporary OSS upload credentials.
	//
	// >2. Uploads the file stream directly to OSS.
	//
	// >3. Calls the CreateDocParserJob operation using the generated OSS URL.
	//
	// example:
	//
	// https://xxx.oss-cn-beijing.aliyuncs.com/document.pdf?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The frame output result.
	FrameOutput *bool `json:"FrameOutput,omitempty" xml:"FrameOutput,omitempty"`
	// The global summary information.
	GlobalSummary *bool `json:"GlobalSummary,omitempty" xml:"GlobalSummary,omitempty"`
	// The image processing format.
	ImageMode *string `json:"ImageMode,omitempty" xml:"ImageMode,omitempty"`
	// The image understanding and analysis setting.
	ImageUnderstanding *string `json:"ImageUnderstanding,omitempty" xml:"ImageUnderstanding,omitempty"`
	// The media chunk interval in seconds.
	MediaChunkIntervalSeconds *int32 `json:"MediaChunkIntervalSeconds,omitempty" xml:"MediaChunkIntervalSeconds,omitempty"`
	// The media chunk strategy.
	MediaChunkStrategy *string `json:"MediaChunkStrategy,omitempty" xml:"MediaChunkStrategy,omitempty"`
	// The number of media frames per minute.
	MediaFramesPerMinute *float64 `json:"MediaFramesPerMinute,omitempty" xml:"MediaFramesPerMinute,omitempty"`
	// The maximum frame budget for media.
	MediaMaxFrameBudget *int32 `json:"MediaMaxFrameBudget,omitempty" xml:"MediaMaxFrameBudget,omitempty"`
	// The minimum frame budget for media.
	MediaMinFrameBudget *int32 `json:"MediaMinFrameBudget,omitempty" xml:"MediaMinFrameBudget,omitempty"`
	// The OSS file URL.
	OssFileUrl *string `json:"OssFileUrl,omitempty" xml:"OssFileUrl,omitempty"`
	// The output format of the parsing result. Valid values:
	//
	// - **markdown**: Markdown format.
	//
	// This parameter is required.
	//
	// example:
	//
	// markdown
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// The parsing scene.
	ParseScene *string `json:"ParseScene,omitempty" xml:"ParseScene,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The response mode.
	ResponseMode *string `json:"ResponseMode,omitempty" xml:"ResponseMode,omitempty"`
	// The result type.
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// The table processing format.
	TableFormat *string `json:"TableFormat,omitempty" xml:"TableFormat,omitempty"`
}

func (s CreateDocParserJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocParserJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDocParserJobRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *CreateDocParserJobRequest) GetAsrLanguage() *string {
	return s.AsrLanguage
}

func (s *CreateDocParserJobRequest) GetAudioClipOutput() *bool {
	return s.AudioClipOutput
}

func (s *CreateDocParserJobRequest) GetAudioWindowSeconds() *int32 {
	return s.AudioWindowSeconds
}

func (s *CreateDocParserJobRequest) GetChunkSummary() *bool {
	return s.ChunkSummary
}

func (s *CreateDocParserJobRequest) GetFileFormat() *string {
	return s.FileFormat
}

func (s *CreateDocParserJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateDocParserJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateDocParserJobRequest) GetFrameOutput() *bool {
	return s.FrameOutput
}

func (s *CreateDocParserJobRequest) GetGlobalSummary() *bool {
	return s.GlobalSummary
}

func (s *CreateDocParserJobRequest) GetImageMode() *string {
	return s.ImageMode
}

func (s *CreateDocParserJobRequest) GetImageUnderstanding() *string {
	return s.ImageUnderstanding
}

func (s *CreateDocParserJobRequest) GetMediaChunkIntervalSeconds() *int32 {
	return s.MediaChunkIntervalSeconds
}

func (s *CreateDocParserJobRequest) GetMediaChunkStrategy() *string {
	return s.MediaChunkStrategy
}

func (s *CreateDocParserJobRequest) GetMediaFramesPerMinute() *float64 {
	return s.MediaFramesPerMinute
}

func (s *CreateDocParserJobRequest) GetMediaMaxFrameBudget() *int32 {
	return s.MediaMaxFrameBudget
}

func (s *CreateDocParserJobRequest) GetMediaMinFrameBudget() *int32 {
	return s.MediaMinFrameBudget
}

func (s *CreateDocParserJobRequest) GetOssFileUrl() *string {
	return s.OssFileUrl
}

func (s *CreateDocParserJobRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *CreateDocParserJobRequest) GetParseScene() *string {
	return s.ParseScene
}

func (s *CreateDocParserJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDocParserJobRequest) GetResponseMode() *string {
	return s.ResponseMode
}

func (s *CreateDocParserJobRequest) GetResultType() *string {
	return s.ResultType
}

func (s *CreateDocParserJobRequest) GetTableFormat() *string {
	return s.TableFormat
}

func (s *CreateDocParserJobRequest) SetAgentName(v string) *CreateDocParserJobRequest {
	s.AgentName = &v
	return s
}

func (s *CreateDocParserJobRequest) SetAsrLanguage(v string) *CreateDocParserJobRequest {
	s.AsrLanguage = &v
	return s
}

func (s *CreateDocParserJobRequest) SetAudioClipOutput(v bool) *CreateDocParserJobRequest {
	s.AudioClipOutput = &v
	return s
}

func (s *CreateDocParserJobRequest) SetAudioWindowSeconds(v int32) *CreateDocParserJobRequest {
	s.AudioWindowSeconds = &v
	return s
}

func (s *CreateDocParserJobRequest) SetChunkSummary(v bool) *CreateDocParserJobRequest {
	s.ChunkSummary = &v
	return s
}

func (s *CreateDocParserJobRequest) SetFileFormat(v string) *CreateDocParserJobRequest {
	s.FileFormat = &v
	return s
}

func (s *CreateDocParserJobRequest) SetFileName(v string) *CreateDocParserJobRequest {
	s.FileName = &v
	return s
}

func (s *CreateDocParserJobRequest) SetFileUrl(v string) *CreateDocParserJobRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateDocParserJobRequest) SetFrameOutput(v bool) *CreateDocParserJobRequest {
	s.FrameOutput = &v
	return s
}

func (s *CreateDocParserJobRequest) SetGlobalSummary(v bool) *CreateDocParserJobRequest {
	s.GlobalSummary = &v
	return s
}

func (s *CreateDocParserJobRequest) SetImageMode(v string) *CreateDocParserJobRequest {
	s.ImageMode = &v
	return s
}

func (s *CreateDocParserJobRequest) SetImageUnderstanding(v string) *CreateDocParserJobRequest {
	s.ImageUnderstanding = &v
	return s
}

func (s *CreateDocParserJobRequest) SetMediaChunkIntervalSeconds(v int32) *CreateDocParserJobRequest {
	s.MediaChunkIntervalSeconds = &v
	return s
}

func (s *CreateDocParserJobRequest) SetMediaChunkStrategy(v string) *CreateDocParserJobRequest {
	s.MediaChunkStrategy = &v
	return s
}

func (s *CreateDocParserJobRequest) SetMediaFramesPerMinute(v float64) *CreateDocParserJobRequest {
	s.MediaFramesPerMinute = &v
	return s
}

func (s *CreateDocParserJobRequest) SetMediaMaxFrameBudget(v int32) *CreateDocParserJobRequest {
	s.MediaMaxFrameBudget = &v
	return s
}

func (s *CreateDocParserJobRequest) SetMediaMinFrameBudget(v int32) *CreateDocParserJobRequest {
	s.MediaMinFrameBudget = &v
	return s
}

func (s *CreateDocParserJobRequest) SetOssFileUrl(v string) *CreateDocParserJobRequest {
	s.OssFileUrl = &v
	return s
}

func (s *CreateDocParserJobRequest) SetOutputFormat(v string) *CreateDocParserJobRequest {
	s.OutputFormat = &v
	return s
}

func (s *CreateDocParserJobRequest) SetParseScene(v string) *CreateDocParserJobRequest {
	s.ParseScene = &v
	return s
}

func (s *CreateDocParserJobRequest) SetRegionId(v string) *CreateDocParserJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDocParserJobRequest) SetResponseMode(v string) *CreateDocParserJobRequest {
	s.ResponseMode = &v
	return s
}

func (s *CreateDocParserJobRequest) SetResultType(v string) *CreateDocParserJobRequest {
	s.ResultType = &v
	return s
}

func (s *CreateDocParserJobRequest) SetTableFormat(v string) *CreateDocParserJobRequest {
	s.TableFormat = &v
	return s
}

func (s *CreateDocParserJobRequest) Validate() error {
	return dara.Validate(s)
}
