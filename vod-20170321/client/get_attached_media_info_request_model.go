// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttachedMediaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthTimeout(v int64) *GetAttachedMediaInfoRequest
	GetAuthTimeout() *int64
	SetMediaIds(v string) *GetAttachedMediaInfoRequest
	GetMediaIds() *string
	SetOutputType(v string) *GetAttachedMediaInfoRequest
	GetOutputType() *string
}

type GetAttachedMediaInfoRequest struct {
	// The validity period of the URL. Unit: seconds.
	//
	// 	- If you set the OutputType parameter to **cdn**:
	//
	//     	- The URL of the auxiliary media asset has a validity period only if URL signing is enabled. Otherwise, the URL of the auxiliary media asset is permanently valid.
	//
	//     	- Minimum value: **1**.
	//
	//     	- Maximum value: unlimited.
	//
	//     	- Default value: If you do not set this parameter, the default validity period that is specified in URL signing is used.
	//
	// 	- If you set the OutputType parameter to **oss**:
	//
	//     	- The URL of the auxiliary media asset has a validity period only if the permissions on the Object Storage Service (OSS) bucket are private. Otherwise, the URL of the auxiliary media asset is permanently valid.
	//
	//     	- Minimum value: **1**.
	//
	//     	- The maximum value for a media asset stored in the VOD bucket is **2592000*	- (30 days) and the maximum value for a media asset stored in an OSS bucket is **129600*	- (36 hours). The maximum value is limited to reduce security risks of the origin.
	//
	//     	- Default value: If you do not set this parameter, the default value **3600*	- is used.
	//
	// example:
	//
	// 3600
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The ID of the auxiliary media asset.
	//
	// 	- Separate multiple IDs with commas (,). You can specify up to 20 IDs.
	//
	// 	- You can obtain the ID from the response to the [CreateUploadAttachedMedia](~~CreateUploadAttachedMedia~~) operation that you call to obtain the upload URL and credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// eb1861d2c9a842340e989dd56****,0222e203cf80f9c22870a4d2c****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The type of the media asset URL. Valid values:
	//
	// 	- **oss**
	//
	// 	- **cdn*	- (default)
	//
	// example:
	//
	// oss
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
}

func (s GetAttachedMediaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAttachedMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAttachedMediaInfoRequest) GetAuthTimeout() *int64 {
	return s.AuthTimeout
}

func (s *GetAttachedMediaInfoRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *GetAttachedMediaInfoRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *GetAttachedMediaInfoRequest) SetAuthTimeout(v int64) *GetAttachedMediaInfoRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetAttachedMediaInfoRequest) SetMediaIds(v string) *GetAttachedMediaInfoRequest {
	s.MediaIds = &v
	return s
}

func (s *GetAttachedMediaInfoRequest) SetOutputType(v string) *GetAttachedMediaInfoRequest {
	s.OutputType = &v
	return s
}

func (s *GetAttachedMediaInfoRequest) Validate() error {
	return dara.Validate(s)
}
