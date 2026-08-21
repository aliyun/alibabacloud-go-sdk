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
	// The validity period of the auxiliary media asset URL. Unit: seconds.
	//
	//  - If OutputType is set to **cdn**:
	//
	//     - The URL expires only if URL authentication is enabled. Otherwise, the URL is permanently valid.
	//
	//     - Minimum value: **1**.
	//
	//     - Maximum value: unlimited.
	//
	//     - Default value: If you do not specify this parameter, the default validity period specified in URL authentication is used.
	//
	// - If OutputType is set to **oss**:
	//
	//     - The URL expires only if the storage permission is set to private. Otherwise, the URL is permanently valid.
	//
	//     - Minimum value: **1**.
	//
	//     - Maximum value: To reduce security risks to the origin server, the maximum value is **2592000*	- (30 days) if the auxiliary media asset is stored in a bucket managed by ApsaraVideo VOD, and **129600*	- (36 hours) if the auxiliary media asset is stored in your own OSS bucket.
	//
	//     - Default value: If you do not specify this parameter, the value is **3600**.
	//
	// example:
	//
	// 3600
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The auxiliary media asset IDs.
	//
	// - Separate multiple IDs with commas (,). You can specify up to 20 IDs.
	//
	// - The IDs are returned after you call the [CreateUploadAttachedMedia](~~CreateUploadAttachedMedia~~) operation to obtain the upload URL and credential for the auxiliary media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// eb1861d2c9a842340e989dd56****,0222e203cf80f9c22870a4d2c****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The type of the output URL. Valid values:
	//
	// - **oss**: the back-to-origin URL.
	//
	// - **cdn*	- (default): the CDN-accelerated URL.
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
