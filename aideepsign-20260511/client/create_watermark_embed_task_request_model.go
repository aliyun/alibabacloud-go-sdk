// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWatermarkEmbedTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateWatermarkEmbedTaskRequest
	GetClientToken() *string
	SetFileType(v string) *CreateWatermarkEmbedTaskRequest
	GetFileType() *string
	SetFileUrl(v string) *CreateWatermarkEmbedTaskRequest
	GetFileUrl() *string
	SetLevel(v string) *CreateWatermarkEmbedTaskRequest
	GetLevel() *string
	SetReserved(v string) *CreateWatermarkEmbedTaskRequest
	GetReserved() *string
	SetWatermarkContent(v string) *CreateWatermarkEmbedTaskRequest
	GetWatermarkContent() *string
}

type CreateWatermarkEmbedTaskRequest struct {
	// A reserved idempotency field. This field is not supported. Passing a value does not cause an error, but it does not produce an idempotency effect.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The file type. Valid values:
	//
	// - image: image.
	//
	// - video: video.
	//
	// The value must match the actual file type.
	//
	// This parameter is required.
	//
	// example:
	//
	// video
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The URL of the file to be processed. Images and videos are supported. Only the HTTP and HTTPS protocols are supported. Videos must be no longer than 10 minutes and no larger than 100 MB.
	//
	// example:
	//
	// https://example.com/demo.mp4
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The embedding strength. Valid values: 0 to 5. A higher strength provides better robustness but lower visual quality.
	//
	// This parameter is required.
	//
	// example:
	//
	// "2"
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// A reserved field. No valid values are available. Do not specify this parameter.
	//
	// example:
	//
	// ""
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The watermark content. The embedded content can be restored to the original text by calling the extraction operation.
	//
	// example:
	//
	// wm-e2e-85970647
	WatermarkContent *string `json:"WatermarkContent,omitempty" xml:"WatermarkContent,omitempty"`
}

func (s CreateWatermarkEmbedTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWatermarkEmbedTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateWatermarkEmbedTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateWatermarkEmbedTaskRequest) GetFileType() *string {
	return s.FileType
}

func (s *CreateWatermarkEmbedTaskRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateWatermarkEmbedTaskRequest) GetLevel() *string {
	return s.Level
}

func (s *CreateWatermarkEmbedTaskRequest) GetReserved() *string {
	return s.Reserved
}

func (s *CreateWatermarkEmbedTaskRequest) GetWatermarkContent() *string {
	return s.WatermarkContent
}

func (s *CreateWatermarkEmbedTaskRequest) SetClientToken(v string) *CreateWatermarkEmbedTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWatermarkEmbedTaskRequest) SetFileType(v string) *CreateWatermarkEmbedTaskRequest {
	s.FileType = &v
	return s
}

func (s *CreateWatermarkEmbedTaskRequest) SetFileUrl(v string) *CreateWatermarkEmbedTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateWatermarkEmbedTaskRequest) SetLevel(v string) *CreateWatermarkEmbedTaskRequest {
	s.Level = &v
	return s
}

func (s *CreateWatermarkEmbedTaskRequest) SetReserved(v string) *CreateWatermarkEmbedTaskRequest {
	s.Reserved = &v
	return s
}

func (s *CreateWatermarkEmbedTaskRequest) SetWatermarkContent(v string) *CreateWatermarkEmbedTaskRequest {
	s.WatermarkContent = &v
	return s
}

func (s *CreateWatermarkEmbedTaskRequest) Validate() error {
	return dara.Validate(s)
}
