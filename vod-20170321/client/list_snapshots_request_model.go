// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthTimeout(v string) *ListSnapshotsRequest
	GetAuthTimeout() *string
	SetPageNo(v string) *ListSnapshotsRequest
	GetPageNo() *string
	SetPageSize(v string) *ListSnapshotsRequest
	GetPageSize() *string
	SetSnapshotType(v string) *ListSnapshotsRequest
	GetSnapshotType() *string
	SetVideoId(v string) *ListSnapshotsRequest
	GetVideoId() *string
}

type ListSnapshotsRequest struct {
	// The validity period of the snapshot URL. Default value: **3600**. Minimum value: **3600**. Unit: seconds.
	//
	// 	- This parameter takes effect only when you enable URL signing. For more information, see [Configure URL signing](https://help.aliyun.com/document_detail/57007.html).
	//
	// 	- If you specify a value smaller than **3,600 seconds**, **3600*	- is used by default.
	//
	// 	- If the snapshot URL is an Object Storage Service (OSS) URL, the maximum value for this parameter is **2592000*	- (30 days). This reduces risks on the origin.
	//
	// example:
	//
	// 3600
	AuthTimeout *string `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: **20**. Maximum value: **100**.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of snapshots to return. Valid values:
	//
	// 	- **CoverSnapshot**: thumbnail snapshot
	//
	// 	- **NormalSnapshot**: regular snapshot
	//
	// 	- **SpriteSnapshot**: sprite snapshot
	//
	// 	- **SpriteOriginSnapshot**: sprite source snapshot
	//
	// 	- **WebVttSnapshot**: WebVTT snapshot
	//
	// example:
	//
	// CoverSnapshot
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	// The ID of the video. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Media Files*	- > **Audio/Video*	- to view the video ID.
	//
	// 	- Obtain the video ID from the response to the [CreateUploadVideo](~~CreateUploadVideo~~) operation that you call to obtain the upload URL and credential.
	//
	// 	- Obtain the video ID from the response to the [SearchMedia](~~SearchMedia~~) operation that you call to query videos.
	//
	// This parameter is required.
	//
	// example:
	//
	// d3e680e618708fbf2cae7cc931****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequest) GetAuthTimeout() *string {
	return s.AuthTimeout
}

func (s *ListSnapshotsRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *ListSnapshotsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListSnapshotsRequest) GetSnapshotType() *string {
	return s.SnapshotType
}

func (s *ListSnapshotsRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *ListSnapshotsRequest) SetAuthTimeout(v string) *ListSnapshotsRequest {
	s.AuthTimeout = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageNo(v string) *ListSnapshotsRequest {
	s.PageNo = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageSize(v string) *ListSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotsRequest) SetSnapshotType(v string) *ListSnapshotsRequest {
	s.SnapshotType = &v
	return s
}

func (s *ListSnapshotsRequest) SetVideoId(v string) *ListSnapshotsRequest {
	s.VideoId = &v
	return s
}

func (s *ListSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
