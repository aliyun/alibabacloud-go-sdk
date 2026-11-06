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
	SetDocumentStructure(v bool) *CreateDocParserJobRequest
	GetDocumentStructure() *bool
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
	// The name of the agent.
	//
	// example:
	//
	// codex
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The Automatic Speech Recognition (ASR) language hint. Valid values: auto, zh, and en. Default value: auto. If set to auto, the language field in the result is empty.
	//
	// example:
	//
	// zh
	AsrLanguage *string `json:"AsrLanguage,omitempty" xml:"AsrLanguage,omitempty"`
	// Specifies whether to output segmented audio clips to OSS. Default value: true.
	//
	// example:
	//
	// true
	AudioClipOutput *bool `json:"AudioClipOutput,omitempty" xml:"AudioClipOutput,omitempty"`
	// The audio chunk window duration in seconds. This parameter takes effect only for pure audio files. Valid values: 10 to 600. If this parameter is passed in a video request, an InvalidParameter error is returned.
	//
	// example:
	//
	// 200
	AudioWindowSeconds *int32 `json:"AudioWindowSeconds,omitempty" xml:"AudioWindowSeconds,omitempty"`
	// Specifies whether to generate a one-sentence summary for each segment. Default value: true.
	//
	// example:
	//
	// true
	ChunkSummary *bool `json:"ChunkSummary,omitempty" xml:"ChunkSummary,omitempty"`
	// Specifies whether to extract the directory and section structure of the document.
	DocumentStructure *bool `json:"DocumentStructure,omitempty" xml:"DocumentStructure,omitempty"`
	// The format of the input file. Valid values:
	//
	// - **pdf**: PDF file.
	//
	// - **docx**: Word file in DOCX format.
	//
	// - **doc**: Word file in DOC format.
	//
	// - **pptx**: PowerPoint file in PPTX format.
	//
	// - **ppt**: PowerPoint file in PPT format.
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
	// The name of the file, which must include the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// document.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The HTTP or HTTPS URL of the file to be parsed.
	//
	// > SDKs for various programming languages also provide the `CreateDocParserJobAdvance` method, which allows you to pass a local file stream (such as InputStream in Java) directly without uploading the file to OSS and constructing a FileUrl in advance. When you use the Advance method, replace the `FileUrl` parameter (URL string) with the `FileUrlObject` parameter (file stream) and keep other request parameters unchanged. The SDK automatically performs the following operations:
	//
	// > 1. Obtains temporary OSS upload credentials.
	//
	// > 2. Uploads the file stream directly to OSS.
	//
	// > 3. Calls the CreateDocParserJob operation by using the generated OSS URL.
	//
	// example:
	//
	// https://xxx.oss-cn-beijing.aliyuncs.com/document.pdf?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Specifies whether to output the original images of keyframes to OSS. This parameter takes effect only for videos. Default value: true. The media references in the result are oss:// protocol addresses, and the caller must have OSS permissions to read them.
	//
	// example:
	//
	// true
	FrameOutput *bool `json:"FrameOutput,omitempty" xml:"FrameOutput,omitempty"`
	// Specifies whether to generate a global summary for the entire media. Default value: true.
	//
	// example:
	//
	// true
	GlobalSummary *bool `json:"GlobalSummary,omitempty" xml:"GlobalSummary,omitempty"`
	// The image handling mode. Valid values: ignore, base64, placeholder, and oss. Default value: oss. If set to oss, image links are returned in the body text.
	//
	// example:
	//
	// ignore
	ImageMode *string `json:"ImageMode,omitempty" xml:"ImageMode,omitempty"`
	// Specifies whether to generate image understanding text. Valid values: full and none. Default value: full.
	//
	// example:
	//
	// full
	ImageUnderstanding *string `json:"ImageUnderstanding,omitempty" xml:"ImageUnderstanding,omitempty"`
	// The media chunking interval in seconds. This parameter takes effect only for videos and when the chunking strategy is set to interval. Valid values: 5 to 600.
	//
	// example:
	//
	// 200
	MediaChunkIntervalSeconds *int32 `json:"MediaChunkIntervalSeconds,omitempty" xml:"MediaChunkIntervalSeconds,omitempty"`
	// The chunking strategy. This parameter takes effect only for videos. Valid values: scene and interval. Default value: scene, which chunks the video by scene. If set to interval, the video is chunked by fixed intervals.
	//
	// example:
	//
	// scene
	MediaChunkStrategy *string `json:"MediaChunkStrategy,omitempty" xml:"MediaChunkStrategy,omitempty"`
	// The number of media frames extracted per minute. This parameter takes effect only for videos. Valid values: 0 to 30. A value of 0 indicates that the dynamic frame budget is disabled.
	//
	// example:
	//
	// 200
	MediaFramesPerMinute *float64 `json:"MediaFramesPerMinute,omitempty" xml:"MediaFramesPerMinute,omitempty"`
	// The maximum media frame budget. This parameter takes effect only for videos. It specifies the upper limit of the frame budget. Valid values: 1 to 600. If the lower limit is also set, the lower limit must not be greater than the upper limit.
	//
	// example:
	//
	// 200
	MediaMaxFrameBudget *int32 `json:"MediaMaxFrameBudget,omitempty" xml:"MediaMaxFrameBudget,omitempty"`
	// The minimum media frame budget. This parameter takes effect only for videos. It specifies the lower limit of the frame budget. Valid values: 1 to 600. If the upper limit is also set, the lower limit must not be greater than the upper limit.
	//
	// example:
	//
	// 100
	MediaMinFrameBudget *int32 `json:"MediaMinFrameBudget,omitempty" xml:"MediaMinFrameBudget,omitempty"`
	// The OSS file URL.
	//
	// example:
	//
	// https://pop-file-upload-cn-beijing.oss-cn-beijing.aliyuncs.com/121852226****
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
	// The parsing scene. This parameter takes effect only for document and image parsing. Valid values: general and formula_ocr. Default value: general. If set to formula_ocr, full-page OCR is performed and formulas are output in LaTeX format.
	//
	// example:
	//
	// formula_ocr
	ParseScene *string `json:"ParseScene,omitempty" xml:"ParseScene,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The response mode. Valid values: async and sync. Default value: async, which returns only the JobId and requires the caller to poll for the result. If set to sync, the parsing result is returned synchronously within the same request. For more information, see the response description below.
	//
	// example:
	//
	// sync
	ResponseMode *string `json:"ResponseMode,omitempty" xml:"ResponseMode,omitempty"`
	// The result return method. Valid values: content and url. Default value: content, which returns the Markdown body text. If set to url, a temporary OSS download URL is returned.
	//
	// example:
	//
	// content
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// The table format. Valid values: html and markdown. Default value: markdown.
	//
	// example:
	//
	// html
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

func (s *CreateDocParserJobRequest) GetDocumentStructure() *bool {
	return s.DocumentStructure
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

func (s *CreateDocParserJobRequest) SetDocumentStructure(v bool) *CreateDocParserJobRequest {
	s.DocumentStructure = &v
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
