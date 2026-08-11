// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthTimeout(v int64) *GetMediaInfoRequest
	GetAuthTimeout() *int64
	SetInputURL(v string) *GetMediaInfoRequest
	GetInputURL() *string
	SetMediaId(v string) *GetMediaInfoRequest
	GetMediaId() *string
	SetOutputType(v string) *GetMediaInfoRequest
	GetOutputType() *string
	SetReturnDetailedInfo(v string) *GetMediaInfoRequest
	GetReturnDetailedInfo() *string
}

type GetMediaInfoRequest struct {
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The address of the media asset in the corresponding system. Before use, the media asset must be registered in the IMS content library and bound to an IMS mediaId.
	//
	// - OSS address. Two formats are supported:
	//
	// http(s)://example-bucket.oss-ap-southeast-1.aliyuncs.com/example.mp4 or
	//
	// oss://example-bucket/example.mp4. This format assumes the OSS region is the same as the service access region by default.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The IMS media asset ID. If this parameter is empty, InputURL is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The type of the media file address in the response:
	//
	// example:
	//
	// cdn
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	// Specifies whether to return detailed information for the corresponding media asset fields. The following fields are supported:
	//
	// AiRoughData.StandardSmartTagJob: Specifies whether to return detailed tagging results if the media asset has been submitted for tag analysis.
	//
	// - Default value: false. The task result is returned as a URL.
	//
	// - true: The task result is returned as text.
	//
	// example:
	//
	// {"AiRoughData.StandardSmartTagJob": false}
	ReturnDetailedInfo *string `json:"ReturnDetailedInfo,omitempty" xml:"ReturnDetailedInfo,omitempty"`
}

func (s GetMediaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMediaInfoRequest) GetAuthTimeout() *int64 {
	return s.AuthTimeout
}

func (s *GetMediaInfoRequest) GetInputURL() *string {
	return s.InputURL
}

func (s *GetMediaInfoRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaInfoRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *GetMediaInfoRequest) GetReturnDetailedInfo() *string {
	return s.ReturnDetailedInfo
}

func (s *GetMediaInfoRequest) SetAuthTimeout(v int64) *GetMediaInfoRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetMediaInfoRequest) SetInputURL(v string) *GetMediaInfoRequest {
	s.InputURL = &v
	return s
}

func (s *GetMediaInfoRequest) SetMediaId(v string) *GetMediaInfoRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaInfoRequest) SetOutputType(v string) *GetMediaInfoRequest {
	s.OutputType = &v
	return s
}

func (s *GetMediaInfoRequest) SetReturnDetailedInfo(v string) *GetMediaInfoRequest {
	s.ReturnDetailedInfo = &v
	return s
}

func (s *GetMediaInfoRequest) Validate() error {
	return dara.Validate(s)
}
