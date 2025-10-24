// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeMediaInfo(v bool) *QueryMediaListRequest
	GetIncludeMediaInfo() *bool
	SetIncludePlayList(v bool) *QueryMediaListRequest
	GetIncludePlayList() *bool
	SetIncludeSnapshotList(v bool) *QueryMediaListRequest
	GetIncludeSnapshotList() *bool
	SetIncludeSummaryList(v bool) *QueryMediaListRequest
	GetIncludeSummaryList() *bool
	SetMediaIds(v string) *QueryMediaListRequest
	GetMediaIds() *string
	SetOwnerAccount(v string) *QueryMediaListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryMediaListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryMediaListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryMediaListRequest
	GetResourceOwnerId() *int64
}

type QueryMediaListRequest struct {
	// Specifies whether to include media information in the returned result.
	//
	// 	- Valid values: true and false.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// true
	IncludeMediaInfo *bool `json:"IncludeMediaInfo,omitempty" xml:"IncludeMediaInfo,omitempty"`
	// Specifies whether to include playback information in the returned result.
	//
	// 	- Valid values: true and false.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// true
	IncludePlayList *bool `json:"IncludePlayList,omitempty" xml:"IncludePlayList,omitempty"`
	// Specifies whether to include snapshot information in the returned result.
	//
	// 	- Valid values: true and false.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// true
	IncludeSnapshotList *bool `json:"IncludeSnapshotList,omitempty" xml:"IncludeSnapshotList,omitempty"`
	// Specifies whether to include summaries in the returned result.
	//
	// 	- Valid values: true and false.
	//
	// 	- Default value: **false**.
	//
	// example:
	//
	// true
	IncludeSummaryList *bool `json:"IncludeSummaryList,omitempty" xml:"IncludeSummaryList,omitempty"`
	// The IDs of the media files. To obtain the ID of a media file, you can perform the following operations in the ApsaraVideo Media Processing (MPS) console: In the left-side navigation pane, choose **Media Management*	- > **Media List**. Find the required video and click Manage. The ID of the video is displayed on the Basics tab. Separate multiple IDs with commas (,). You can query up to 10 media files at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e1cd21131a94525be55acf65888****,e26cfa29e784402388463f61dbec****
	MediaIds             *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryMediaListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListRequest) GoString() string {
	return s.String()
}

func (s *QueryMediaListRequest) GetIncludeMediaInfo() *bool {
	return s.IncludeMediaInfo
}

func (s *QueryMediaListRequest) GetIncludePlayList() *bool {
	return s.IncludePlayList
}

func (s *QueryMediaListRequest) GetIncludeSnapshotList() *bool {
	return s.IncludeSnapshotList
}

func (s *QueryMediaListRequest) GetIncludeSummaryList() *bool {
	return s.IncludeSummaryList
}

func (s *QueryMediaListRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *QueryMediaListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryMediaListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryMediaListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryMediaListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryMediaListRequest) SetIncludeMediaInfo(v bool) *QueryMediaListRequest {
	s.IncludeMediaInfo = &v
	return s
}

func (s *QueryMediaListRequest) SetIncludePlayList(v bool) *QueryMediaListRequest {
	s.IncludePlayList = &v
	return s
}

func (s *QueryMediaListRequest) SetIncludeSnapshotList(v bool) *QueryMediaListRequest {
	s.IncludeSnapshotList = &v
	return s
}

func (s *QueryMediaListRequest) SetIncludeSummaryList(v bool) *QueryMediaListRequest {
	s.IncludeSummaryList = &v
	return s
}

func (s *QueryMediaListRequest) SetMediaIds(v string) *QueryMediaListRequest {
	s.MediaIds = &v
	return s
}

func (s *QueryMediaListRequest) SetOwnerAccount(v string) *QueryMediaListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryMediaListRequest) SetOwnerId(v int64) *QueryMediaListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMediaListRequest) SetResourceOwnerAccount(v string) *QueryMediaListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMediaListRequest) SetResourceOwnerId(v int64) *QueryMediaListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMediaListRequest) Validate() error {
	return dara.Validate(s)
}
