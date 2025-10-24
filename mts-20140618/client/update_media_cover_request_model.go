// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaCoverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverURL(v string) *UpdateMediaCoverRequest
	GetCoverURL() *string
	SetMediaId(v string) *UpdateMediaCoverRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *UpdateMediaCoverRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateMediaCoverRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateMediaCoverRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateMediaCoverRequest
	GetResourceOwnerId() *int64
}

type UpdateMediaCoverRequest struct {
	// The URL of the thumbnail that you want to specify for the media file. The URL complies with RFC 2396 and is encoded in UTF-8. The URL can be up to 3,200 bytes in length.
	//
	// >  To obtain the thumbnail URL, you can find the image in the Object Storage Service (OSS) bucket and click the image to view details. In the View Details panel, copy the part before the question mark (?) from the URL field. You can enter only an HTTP URL.
	//
	// example:
	//
	// http://example-bucket-****.oss-cn-hangzhou.aliyuncs.com//example-****.mp4
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The ID of the media file whose thumbnail you want to update. To obtain the ID of a media file, you can call the [AddMedia](https://help.aliyun.com/document_detail/44458.html) operation. Alternatively, perform the following operations in the ApsaraVideo Media Processing (MPS) console: In the left-side navigation pane, choose **Media Management*	- > **Media List**. Find the required video and click **Manage**. The ID of the video is displayed on the Basics tab.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6cc3aa66d1cb4bb2adf14e726c0a****
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateMediaCoverRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaCoverRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaCoverRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *UpdateMediaCoverRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaCoverRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateMediaCoverRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateMediaCoverRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateMediaCoverRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateMediaCoverRequest) SetCoverURL(v string) *UpdateMediaCoverRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateMediaCoverRequest) SetMediaId(v string) *UpdateMediaCoverRequest {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaCoverRequest) SetOwnerAccount(v string) *UpdateMediaCoverRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateMediaCoverRequest) SetOwnerId(v int64) *UpdateMediaCoverRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMediaCoverRequest) SetResourceOwnerAccount(v string) *UpdateMediaCoverRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateMediaCoverRequest) SetResourceOwnerId(v int64) *UpdateMediaCoverRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateMediaCoverRequest) Validate() error {
	return dara.Validate(s)
}
