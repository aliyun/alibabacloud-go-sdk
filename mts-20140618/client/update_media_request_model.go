// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCateId(v int64) *UpdateMediaRequest
	GetCateId() *int64
	SetCoverURL(v string) *UpdateMediaRequest
	GetCoverURL() *string
	SetDescription(v string) *UpdateMediaRequest
	GetDescription() *string
	SetMediaId(v string) *UpdateMediaRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *UpdateMediaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateMediaRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateMediaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateMediaRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *UpdateMediaRequest
	GetTags() *string
	SetTitle(v string) *UpdateMediaRequest
	GetTitle() *string
}

type UpdateMediaRequest struct {
	// The ID of the category to which the media file belongs. The value must be an integer.
	//
	// 	- If you do not specify this parameter, the value is NULL.
	//
	// 	- The value cannot be negative.
	//
	// example:
	//
	// 1
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The URL of the thumbnail. This parameter is used to specify the storage location of the thumbnail. To obtain the URL, you can log on to the **MPS console*	- and choose **Workflows*	- > **Media Buckets*	- in the left-side navigation pane. Alternatively, you can log on to the **OSS console*	- and click **Buckets*	- in the left-side navigation pane.
	//
	// 	- The value can be up to 3,200 bytes in length.
	//
	// 	- The URL complies with RFC 2396 and is encoded in UTF-8, with reserved characters being percent-encoded. For more information, see [URL encoding](https://help.aliyun.com/document_detail/423796.html).
	//
	// example:
	//
	// http://example-bucket-****.oss-cn-hangzhou.aliyuncs.com/test****.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The description of the media file. Multiple character types, such as letters and digits, are supported.
	//
	// 	- If you do not specify this parameter, the value is NULL.
	//
	// 	- The value is encoded in UTF-8 and can be up to 1,024 bytes in length.
	//
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the media file whose basic information you want to update. To obtain the ID of the media file, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Media Management*	- > **Media List*	- in the left-side navigation pane.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e1cd21131a94525be55acf65888****
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags that you want to add to the media file.
	//
	// 	- You can specify up to 16 tags for a media file. Separate multiple tags with commas (,).
	//
	// 	- Each tag can be up to 32 bytes in length.
	//
	// 	- The value is encoded in UTF-8.
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the media file. Multiple character types, such as letters and digits, are supported.
	//
	// 	- If you do not specify this parameter, the value is NULL.
	//
	// 	- The value is encoded in UTF-8 and can be up to 128 bytes in length.
	//
	// example:
	//
	// hello
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *UpdateMediaRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *UpdateMediaRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMediaRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateMediaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateMediaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateMediaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateMediaRequest) GetTags() *string {
	return s.Tags
}

func (s *UpdateMediaRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateMediaRequest) SetCateId(v int64) *UpdateMediaRequest {
	s.CateId = &v
	return s
}

func (s *UpdateMediaRequest) SetCoverURL(v string) *UpdateMediaRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateMediaRequest) SetDescription(v string) *UpdateMediaRequest {
	s.Description = &v
	return s
}

func (s *UpdateMediaRequest) SetMediaId(v string) *UpdateMediaRequest {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaRequest) SetOwnerAccount(v string) *UpdateMediaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateMediaRequest) SetOwnerId(v int64) *UpdateMediaRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMediaRequest) SetResourceOwnerAccount(v string) *UpdateMediaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateMediaRequest) SetResourceOwnerId(v int64) *UpdateMediaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateMediaRequest) SetTags(v string) *UpdateMediaRequest {
	s.Tags = &v
	return s
}

func (s *UpdateMediaRequest) SetTitle(v string) *UpdateMediaRequest {
	s.Title = &v
	return s
}

func (s *UpdateMediaRequest) Validate() error {
	return dara.Validate(s)
}
