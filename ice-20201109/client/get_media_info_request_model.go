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
	// The validity period of the signed URL, in seconds.
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The address of the media asset to query. You must first register the media asset in the IMS media library and bind it to a `mediaId`.
	//
	// - Object Storage Service (OSS) URL. Two formats are supported:
	//
	// `http(s)://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4`
	//
	// `oss://example-bucket/example.mp4`. When you use this format, the OSS region defaults to the service endpoint region.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The ID of the media asset in Intelligent Media Services (IMS). If you omit this parameter, you must specify `InputURL`.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The type of URL to return for the media asset file.
	//
	// - `oss`: Returns the OSS URL. This is the default value.
	//
	// - `cdn`: Returns the Content Delivery Network (CDN) URL. A CDN URL is returned only if the media asset was imported from Video on Demand (VOD) and has a CDN domain name configured in VOD.
	//
	// example:
	//
	// cdn
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	// Whether to return detailed information for specific media asset fields. The only supported field is `AiRoughData.StandardSmartTagJob`, which specifies how the result of a tag analysis task is returned.
	//
	// - `false`: The task result is returned as a URL. This is the default value.
	//
	// - `true`: The task result is returned as a string.
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
