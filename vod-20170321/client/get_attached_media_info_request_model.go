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
	// The expiration time of the auxiliary media asset URLs. Unit: seconds.
	//
	//  - If OutputType is set to **cdn**:
	//
	//     - The URLs expire only if URL authentication is enabled. Otherwise, the URLs are permanently valid.
	//
	//     - Minimum value: **1**.
	//
	//     - Maximum value: unlimited.
	//
	//     - Default value: If this parameter is not specified, the default validity period configured for URL authentication is used.
	//
	// - If OutputType is set to **oss**:
	//
	//     - The URLs expire only if the storage permission is set to private. Otherwise, the URLs are permanently valid.
	//
	//     - Minimum value: **1**.
	//
	//     - Maximum value: To reduce security risks to the origin server, the maximum value is **2592000*	- (30 days) if the auxiliary media assets are stored in an ApsaraVideo VOD bucket, and **129600*	- (36 hours) if the auxiliary media assets are stored in a user-owned OSS bucket.
	//
	//     - Default value: If this parameter is not specified, the value is **3600**.
	//
	// example:
	//
	// 3600
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The IDs of the auxiliary media assets.
	//
	// - Separate multiple IDs with commas (,). You can specify up to 20 IDs.
	//
	// - The IDs are returned after you call the [CreateUploadAttachedMedia](~~CreateUploadAttachedMedia~~) operation to obtain the upload URLs and credentials for the auxiliary media assets.
	//
	// This parameter is required.
	//
	// example:
	//
	// eb1861d2c9a842340e989dd56****,0222e203cf80f9c22870a4d2c****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The type of the output URL. Valid values:
	//
	// - **oss**: back-to-origin URL.
	//
	// - **cdn*	- (default): accelerated URL.
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
