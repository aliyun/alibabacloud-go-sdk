// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWatermarkExtractTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateWatermarkExtractTaskRequest
	GetClientToken() *string
	SetFileType(v string) *CreateWatermarkExtractTaskRequest
	GetFileType() *string
	SetFileUrl(v string) *CreateWatermarkExtractTaskRequest
	GetFileUrl() *string
	SetReserved(v string) *CreateWatermarkExtractTaskRequest
	GetReserved() *string
}

type CreateWatermarkExtractTaskRequest struct {
	// A reserved idempotency field. Idempotency is not supported. Passing in a value does not cause an error, but it does not produce an idempotency effect.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The file type. Valid values: `image` and `video`. The value must be the same as the `FileType` used when the watermark was embedded.
	//
	// This parameter is required.
	//
	// example:
	//
	// video
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The URL of the file from which the watermark is to be extracted. Images and videos are supported. Only the `http/https` protocol is supported. Videos must be no longer than 10 minutes and no larger than 100 MB.
	//
	// example:
	//
	// https://example.com/demo.mp4
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// A reserved field. No valid values are available. Do not pass in this parameter.
	//
	// example:
	//
	// ""
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
}

func (s CreateWatermarkExtractTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWatermarkExtractTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateWatermarkExtractTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateWatermarkExtractTaskRequest) GetFileType() *string {
	return s.FileType
}

func (s *CreateWatermarkExtractTaskRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateWatermarkExtractTaskRequest) GetReserved() *string {
	return s.Reserved
}

func (s *CreateWatermarkExtractTaskRequest) SetClientToken(v string) *CreateWatermarkExtractTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWatermarkExtractTaskRequest) SetFileType(v string) *CreateWatermarkExtractTaskRequest {
	s.FileType = &v
	return s
}

func (s *CreateWatermarkExtractTaskRequest) SetFileUrl(v string) *CreateWatermarkExtractTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateWatermarkExtractTaskRequest) SetReserved(v string) *CreateWatermarkExtractTaskRequest {
	s.Reserved = &v
	return s
}

func (s *CreateWatermarkExtractTaskRequest) Validate() error {
	return dara.Validate(s)
}
