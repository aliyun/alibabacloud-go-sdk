// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMezzanineInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionType(v string) *GetMezzanineInfoRequest
	GetAdditionType() *string
	SetAuthTimeout(v int64) *GetMezzanineInfoRequest
	GetAuthTimeout() *int64
	SetOutputType(v string) *GetMezzanineInfoRequest
	GetOutputType() *string
	SetReferenceId(v string) *GetMezzanineInfoRequest
	GetReferenceId() *string
	SetVideoId(v string) *GetMezzanineInfoRequest
	GetVideoId() *string
}

type GetMezzanineInfoRequest struct {
	// The type of additional information. Separate multiple values with commas (,). By default, only the basic information is returned. Valid values:
	//
	// 	- **video**: video stream information
	//
	// 	- **audio**: audio stream information
	//
	// example:
	//
	// video
	AdditionType *string `json:"AdditionType,omitempty" xml:"AdditionType,omitempty"`
	// The validity period of the mezzanine file URL. Unit: seconds. Default value: **1800**. Minimum value: **1**.
	//
	// 	- If the OutputType parameter is set to **cdn**:
	//
	//     	- The mezzanine file URL has a validity period only if URL signing is enabled. Otherwise, the mezzanine file URL is permanently valid.
	//
	//     	- Minimum value: **1**.
	//
	//     	- Maximum Value: unlimited.
	//
	//     	- Default value: If you do not set this parameter, the default validity period that is specified in URL signing is used.
	//
	// <!---->
	//
	// 	- If the OutputType parameter is set to **oss**:
	//
	//     	- The mezzanine file URL has a validity period only if the permissions on the Object Storage Service (OSS) bucket are private. Otherwise, the mezzanine file URL is permanently valid.
	//
	//     	- Minimum value: **1**.
	//
	//     	- Maximum value: **2592000*	- (30 days). The maximum value is limited to reduce security risks of the origin.
	//
	//     	- Default value: If you do not set this parameter, the default value is **3600**.
	//
	// example:
	//
	// 3600
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The type of the mezzanine file URL. Valid values:
	//
	// - **oss**: OSS URL
	//
	// - **cdn*	- (default): Content Delivery Network (CDN) URL
	//
	// > If the mezzanine file is stored in a bucket of the in type, only an OSS URL is returned.
	//
	// example:
	//
	// oss
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	// example:
	//
	// 123-123
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The ID of the video.
	//
	// example:
	//
	// 1f1a6fc03ca04814031b8a6559e****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetMezzanineInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMezzanineInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoRequest) GetAdditionType() *string {
	return s.AdditionType
}

func (s *GetMezzanineInfoRequest) GetAuthTimeout() *int64 {
	return s.AuthTimeout
}

func (s *GetMezzanineInfoRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *GetMezzanineInfoRequest) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *GetMezzanineInfoRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *GetMezzanineInfoRequest) SetAdditionType(v string) *GetMezzanineInfoRequest {
	s.AdditionType = &v
	return s
}

func (s *GetMezzanineInfoRequest) SetAuthTimeout(v int64) *GetMezzanineInfoRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetMezzanineInfoRequest) SetOutputType(v string) *GetMezzanineInfoRequest {
	s.OutputType = &v
	return s
}

func (s *GetMezzanineInfoRequest) SetReferenceId(v string) *GetMezzanineInfoRequest {
	s.ReferenceId = &v
	return s
}

func (s *GetMezzanineInfoRequest) SetVideoId(v string) *GetMezzanineInfoRequest {
	s.VideoId = &v
	return s
}

func (s *GetMezzanineInfoRequest) Validate() error {
	return dara.Validate(s)
}
