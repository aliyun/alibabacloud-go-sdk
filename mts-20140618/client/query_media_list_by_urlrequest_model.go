// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaListByURLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileURLs(v string) *QueryMediaListByURLRequest
	GetFileURLs() *string
	SetIncludeMediaInfo(v bool) *QueryMediaListByURLRequest
	GetIncludeMediaInfo() *bool
	SetIncludePlayList(v bool) *QueryMediaListByURLRequest
	GetIncludePlayList() *bool
	SetIncludeSnapshotList(v bool) *QueryMediaListByURLRequest
	GetIncludeSnapshotList() *bool
	SetIncludeSummaryList(v bool) *QueryMediaListByURLRequest
	GetIncludeSummaryList() *bool
	SetOwnerAccount(v string) *QueryMediaListByURLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryMediaListByURLRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryMediaListByURLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryMediaListByURLRequest
	GetResourceOwnerId() *int64
}

type QueryMediaListByURLRequest struct {
	// The OSS URLs of the media files. To obtain the OSS URL of a media file, you can perform the following operations in the ApsaraVideo Media Processing (MPS) console: In the left-side navigation pane, choose **Media Management*	- > **Media List**. Find the media file whose OSS URL you want to view and click **Manage*	- in the **Actions*	- column. The OSS URL of the media file is displayed on the **Obtain Encoding URL*	- tab. Separate multiple URLs with commas (,). You can query up to 10 media files at a time.
	//
	// 	- The URL complies with RFC 3986 and is encoded in UTF-8, with reserved characters being percent-encoded. The value can be up to 3,200 bytes in size. For more information, see [URL encoding](https://help.aliyun.com/document_detail/423796.html).
	//
	// 	- Only OSS HTTP URLs are supported. Alibaba Cloud CDN URLs and HTTPS URLs are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example-bucket-****.oss-cn-shanghai.aliyuncs.com/example.mp4
	FileURLs *string `json:"FileURLs,omitempty" xml:"FileURLs,omitempty"`
	// Specifies whether to include media information in the returned result.
	//
	// 	- Valid values: true and false.
	//
	// 	- Default value: **false**.
	//
	// > To obtain detailed information about the media files, set this parameter to true.
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
	IncludeSummaryList   *bool   `json:"IncludeSummaryList,omitempty" xml:"IncludeSummaryList,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryMediaListByURLRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListByURLRequest) GoString() string {
	return s.String()
}

func (s *QueryMediaListByURLRequest) GetFileURLs() *string {
	return s.FileURLs
}

func (s *QueryMediaListByURLRequest) GetIncludeMediaInfo() *bool {
	return s.IncludeMediaInfo
}

func (s *QueryMediaListByURLRequest) GetIncludePlayList() *bool {
	return s.IncludePlayList
}

func (s *QueryMediaListByURLRequest) GetIncludeSnapshotList() *bool {
	return s.IncludeSnapshotList
}

func (s *QueryMediaListByURLRequest) GetIncludeSummaryList() *bool {
	return s.IncludeSummaryList
}

func (s *QueryMediaListByURLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryMediaListByURLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryMediaListByURLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryMediaListByURLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryMediaListByURLRequest) SetFileURLs(v string) *QueryMediaListByURLRequest {
	s.FileURLs = &v
	return s
}

func (s *QueryMediaListByURLRequest) SetIncludeMediaInfo(v bool) *QueryMediaListByURLRequest {
	s.IncludeMediaInfo = &v
	return s
}

func (s *QueryMediaListByURLRequest) SetIncludePlayList(v bool) *QueryMediaListByURLRequest {
	s.IncludePlayList = &v
	return s
}

func (s *QueryMediaListByURLRequest) SetIncludeSnapshotList(v bool) *QueryMediaListByURLRequest {
	s.IncludeSnapshotList = &v
	return s
}

func (s *QueryMediaListByURLRequest) SetIncludeSummaryList(v bool) *QueryMediaListByURLRequest {
	s.IncludeSummaryList = &v
	return s
}

func (s *QueryMediaListByURLRequest) SetOwnerAccount(v string) *QueryMediaListByURLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryMediaListByURLRequest) SetOwnerId(v int64) *QueryMediaListByURLRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMediaListByURLRequest) SetResourceOwnerAccount(v string) *QueryMediaListByURLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryMediaListByURLRequest) SetResourceOwnerId(v int64) *QueryMediaListByURLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryMediaListByURLRequest) Validate() error {
	return dara.Validate(s)
}
