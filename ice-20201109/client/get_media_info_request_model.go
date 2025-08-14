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
	// The input URL of the media asset in another service. The URL must be registered in the IMS content library and bound to the ID of the media asset in IMS.
	//
	// 	- For a media asset from Object Storage Service (OSS), the URL may have one of the following formats:
	//
	// http(s)://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4 or
	//
	// oss://example-bucket/example.mp4. The second format indicates that the region in which the OSS bucket of the media asset resides is the same as the region in which OSS is activated.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The ID of the media asset in IMS. If this parameter is left empty, the InputURL parameter must be specified.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The type of the URL of the media asset to return in the response. Valid values:
	//
	// 	- oss (default): an OSS URL.
	//
	// 	- cdn: a CDN URL. A CDN URL is returned only if the media asset is imported from ApsaraVideo VOD and the relevant domain name is an accelerated domain name in ApsaraVideo VOD.
	//
	// example:
	//
	// cdn
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	// Specifies whether to return detailed information for specific media asset attributes. Supported attributes: AiRoughData.StandardSmartTagJob, which specifies whether to return detailed tag information if a tagging job has been submitted for the media asset. Valid values for the attribute:
	//
	// 	- false (default): The job result is returned as a URL.
	//
	// 	- true: The job result is returned as text.
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
