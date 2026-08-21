// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAiAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthTimeout(v string) *GetMediaAiAnalysisRequest
	GetAuthTimeout() *string
	SetMediaId(v string) *GetMediaAiAnalysisRequest
	GetMediaId() *string
	SetOutputType(v string) *GetMediaAiAnalysisRequest
	GetOutputType() *string
	SetResultTypes(v string) *GetMediaAiAnalysisRequest
	GetResultTypes() *string
}

type GetMediaAiAnalysisRequest struct {
	// The expiration time of the image access URL. Unit: seconds.
	//
	// - If OutputType is set to cdn:
	//
	//     - Only image URLs with URL authentication enabled expire. Otherwise, the URLs are permanently valid.
	//
	//     - Minimum value: 1.
	//
	//     - Maximum value: unlimited.
	//
	//     - Default value: If this parameter is not specified, the default validity period specified in URL authentication settings is used.
	//
	// - If OutputType is set to oss:
	//
	//     - Only image URLs with private storage permissions expire. Otherwise, the URLs are permanently valid.
	//
	//     - Minimum value: 1.
	//
	//     - Maximum value: To reduce security risks to the origin server, the maximum value is **2592000*	- (30 days) if images are stored in a bucket managed by ApsaraVideo VOD, and **129600*	- (36 hours) if images are stored in your own OSS bucket.
	//
	//     - Default value: If this parameter is not specified, the value is 3600.
	//
	// example:
	//
	// 3600
	AuthTimeout *string `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The audio ID. You can query the audio ID in the ApsaraVideo VOD console or obtain it from the response of the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation.
	//
	// example:
	//
	// 006204a11bb386bb25491f95f****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The type of the output URL. Valid values:
	//
	// - **oss**: back-to-origin URL.
	//
	// - **cdn*	- (default): accelerated URL.
	//
	// example:
	//
	// cdn
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	// The type of analysis results. Separate multiple types with commas (,).
	//
	// example:
	//
	// Chapter
	ResultTypes *string `json:"ResultTypes,omitempty" xml:"ResultTypes,omitempty"`
}

func (s GetMediaAiAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAiAnalysisRequest) GoString() string {
	return s.String()
}

func (s *GetMediaAiAnalysisRequest) GetAuthTimeout() *string {
	return s.AuthTimeout
}

func (s *GetMediaAiAnalysisRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaAiAnalysisRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *GetMediaAiAnalysisRequest) GetResultTypes() *string {
	return s.ResultTypes
}

func (s *GetMediaAiAnalysisRequest) SetAuthTimeout(v string) *GetMediaAiAnalysisRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetMediaAiAnalysisRequest) SetMediaId(v string) *GetMediaAiAnalysisRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaAiAnalysisRequest) SetOutputType(v string) *GetMediaAiAnalysisRequest {
	s.OutputType = &v
	return s
}

func (s *GetMediaAiAnalysisRequest) SetResultTypes(v string) *GetMediaAiAnalysisRequest {
	s.ResultTypes = &v
	return s
}

func (s *GetMediaAiAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
