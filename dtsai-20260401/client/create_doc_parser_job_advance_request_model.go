// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCreateDocParserJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *CreateDocParserJobAdvanceRequest
	GetAgentName() *string
	SetAsrLanguage(v string) *CreateDocParserJobAdvanceRequest
	GetAsrLanguage() *string
	SetAudioClipOutput(v bool) *CreateDocParserJobAdvanceRequest
	GetAudioClipOutput() *bool
	SetAudioWindowSeconds(v int32) *CreateDocParserJobAdvanceRequest
	GetAudioWindowSeconds() *int32
	SetChunkSummary(v bool) *CreateDocParserJobAdvanceRequest
	GetChunkSummary() *bool
	SetFileFormat(v string) *CreateDocParserJobAdvanceRequest
	GetFileFormat() *string
	SetFileName(v string) *CreateDocParserJobAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *CreateDocParserJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetFrameOutput(v bool) *CreateDocParserJobAdvanceRequest
	GetFrameOutput() *bool
	SetGlobalSummary(v bool) *CreateDocParserJobAdvanceRequest
	GetGlobalSummary() *bool
	SetImageMode(v string) *CreateDocParserJobAdvanceRequest
	GetImageMode() *string
	SetImageUnderstanding(v string) *CreateDocParserJobAdvanceRequest
	GetImageUnderstanding() *string
	SetMediaChunkIntervalSeconds(v int32) *CreateDocParserJobAdvanceRequest
	GetMediaChunkIntervalSeconds() *int32
	SetMediaChunkStrategy(v string) *CreateDocParserJobAdvanceRequest
	GetMediaChunkStrategy() *string
	SetMediaFramesPerMinute(v float64) *CreateDocParserJobAdvanceRequest
	GetMediaFramesPerMinute() *float64
	SetMediaMaxFrameBudget(v int32) *CreateDocParserJobAdvanceRequest
	GetMediaMaxFrameBudget() *int32
	SetMediaMinFrameBudget(v int32) *CreateDocParserJobAdvanceRequest
	GetMediaMinFrameBudget() *int32
	SetOssFileUrl(v string) *CreateDocParserJobAdvanceRequest
	GetOssFileUrl() *string
	SetOutputFormat(v string) *CreateDocParserJobAdvanceRequest
	GetOutputFormat() *string
	SetParseScene(v string) *CreateDocParserJobAdvanceRequest
	GetParseScene() *string
	SetRegionId(v string) *CreateDocParserJobAdvanceRequest
	GetRegionId() *string
	SetResponseMode(v string) *CreateDocParserJobAdvanceRequest
	GetResponseMode() *string
	SetResultType(v string) *CreateDocParserJobAdvanceRequest
	GetResultType() *string
	SetTableFormat(v string) *CreateDocParserJobAdvanceRequest
	GetTableFormat() *string
}

type CreateDocParserJobAdvanceRequest struct {
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
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
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

func (s CreateDocParserJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocParserJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDocParserJobAdvanceRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *CreateDocParserJobAdvanceRequest) GetAsrLanguage() *string {
	return s.AsrLanguage
}

func (s *CreateDocParserJobAdvanceRequest) GetAudioClipOutput() *bool {
	return s.AudioClipOutput
}

func (s *CreateDocParserJobAdvanceRequest) GetAudioWindowSeconds() *int32 {
	return s.AudioWindowSeconds
}

func (s *CreateDocParserJobAdvanceRequest) GetChunkSummary() *bool {
	return s.ChunkSummary
}

func (s *CreateDocParserJobAdvanceRequest) GetFileFormat() *string {
	return s.FileFormat
}

func (s *CreateDocParserJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateDocParserJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *CreateDocParserJobAdvanceRequest) GetFrameOutput() *bool {
	return s.FrameOutput
}

func (s *CreateDocParserJobAdvanceRequest) GetGlobalSummary() *bool {
	return s.GlobalSummary
}

func (s *CreateDocParserJobAdvanceRequest) GetImageMode() *string {
	return s.ImageMode
}

func (s *CreateDocParserJobAdvanceRequest) GetImageUnderstanding() *string {
	return s.ImageUnderstanding
}

func (s *CreateDocParserJobAdvanceRequest) GetMediaChunkIntervalSeconds() *int32 {
	return s.MediaChunkIntervalSeconds
}

func (s *CreateDocParserJobAdvanceRequest) GetMediaChunkStrategy() *string {
	return s.MediaChunkStrategy
}

func (s *CreateDocParserJobAdvanceRequest) GetMediaFramesPerMinute() *float64 {
	return s.MediaFramesPerMinute
}

func (s *CreateDocParserJobAdvanceRequest) GetMediaMaxFrameBudget() *int32 {
	return s.MediaMaxFrameBudget
}

func (s *CreateDocParserJobAdvanceRequest) GetMediaMinFrameBudget() *int32 {
	return s.MediaMinFrameBudget
}

func (s *CreateDocParserJobAdvanceRequest) GetOssFileUrl() *string {
	return s.OssFileUrl
}

func (s *CreateDocParserJobAdvanceRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *CreateDocParserJobAdvanceRequest) GetParseScene() *string {
	return s.ParseScene
}

func (s *CreateDocParserJobAdvanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDocParserJobAdvanceRequest) GetResponseMode() *string {
	return s.ResponseMode
}

func (s *CreateDocParserJobAdvanceRequest) GetResultType() *string {
	return s.ResultType
}

func (s *CreateDocParserJobAdvanceRequest) GetTableFormat() *string {
	return s.TableFormat
}

func (s *CreateDocParserJobAdvanceRequest) SetAgentName(v string) *CreateDocParserJobAdvanceRequest {
	s.AgentName = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetAsrLanguage(v string) *CreateDocParserJobAdvanceRequest {
	s.AsrLanguage = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetAudioClipOutput(v bool) *CreateDocParserJobAdvanceRequest {
	s.AudioClipOutput = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetAudioWindowSeconds(v int32) *CreateDocParserJobAdvanceRequest {
	s.AudioWindowSeconds = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetChunkSummary(v bool) *CreateDocParserJobAdvanceRequest {
	s.ChunkSummary = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetFileFormat(v string) *CreateDocParserJobAdvanceRequest {
	s.FileFormat = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetFileName(v string) *CreateDocParserJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetFileUrlObject(v io.Reader) *CreateDocParserJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetFrameOutput(v bool) *CreateDocParserJobAdvanceRequest {
	s.FrameOutput = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetGlobalSummary(v bool) *CreateDocParserJobAdvanceRequest {
	s.GlobalSummary = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetImageMode(v string) *CreateDocParserJobAdvanceRequest {
	s.ImageMode = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetImageUnderstanding(v string) *CreateDocParserJobAdvanceRequest {
	s.ImageUnderstanding = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetMediaChunkIntervalSeconds(v int32) *CreateDocParserJobAdvanceRequest {
	s.MediaChunkIntervalSeconds = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetMediaChunkStrategy(v string) *CreateDocParserJobAdvanceRequest {
	s.MediaChunkStrategy = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetMediaFramesPerMinute(v float64) *CreateDocParserJobAdvanceRequest {
	s.MediaFramesPerMinute = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetMediaMaxFrameBudget(v int32) *CreateDocParserJobAdvanceRequest {
	s.MediaMaxFrameBudget = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetMediaMinFrameBudget(v int32) *CreateDocParserJobAdvanceRequest {
	s.MediaMinFrameBudget = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetOssFileUrl(v string) *CreateDocParserJobAdvanceRequest {
	s.OssFileUrl = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetOutputFormat(v string) *CreateDocParserJobAdvanceRequest {
	s.OutputFormat = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetParseScene(v string) *CreateDocParserJobAdvanceRequest {
	s.ParseScene = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetRegionId(v string) *CreateDocParserJobAdvanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetResponseMode(v string) *CreateDocParserJobAdvanceRequest {
	s.ResponseMode = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetResultType(v string) *CreateDocParserJobAdvanceRequest {
	s.ResultType = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetTableFormat(v string) *CreateDocParserJobAdvanceRequest {
	s.TableFormat = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
