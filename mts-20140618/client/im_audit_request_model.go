// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *ImAuditRequest
	GetBizType() *string
	SetContents(v string) *ImAuditRequest
	GetContents() *string
	SetImages(v string) *ImAuditRequest
	GetImages() *string
	SetOwnerId(v int64) *ImAuditRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ImAuditRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImAuditRequest
	GetResourceOwnerId() *int64
	SetScenes(v string) *ImAuditRequest
	GetScenes() *string
}

type ImAuditRequest struct {
	// The business type. By default, the public business type is used.
	//
	// example:
	//
	// 139440480445****
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The custom text entries. You can specify up to 5 text entries. The value must be a JSON array. You must specify at least one of the Images and Contents parameters.
	//
	// example:
	//
	// ["Hello","Who are you","Where am I"]
	Contents *string `json:"Contents,omitempty" xml:"Contents,omitempty"`
	// The image URLs. You can specify up to 5 image URLs. The value must be a JSON array. To view the URLs of the images, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Media Management*	- > **Media List*	- in the left-side navigation pane. You must set at least one of the Images and Contents parameters. The image to be moderated must meet the following limits. Otherwise, the moderation task may fail.
	//
	// 	- The image size cannot exceed 20 MB, the height or width of the image cannot exceed 30,000 pixels, and the image cannot exceed 0.25 billion pixels.
	//
	// 	- We recommend that you upload images of at least 256 × 256 pixels to ensure required moderation result.
	//
	// example:
	//
	// ["http://``127.66.**.**``/image.jpeg","http://``127.66.**.**``/photo.jpeg"]
	Images               *string `json:"Images,omitempty" xml:"Images,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The moderation scenarios. Separate multiple scenarios with commas (,). For example, if you specify {"porn","terrorism"} for this parameter, both pornographic content detection and terrorist content detection are performed on the images and text. Valid values:
	//
	// 	- porn: pornography
	//
	// 	- terrorism: terrorist content
	//
	// 	- ad: ad violation
	//
	// 	- qrcode: QR code
	//
	// 	- live: undesirable scene
	//
	// 	- logo: special logo
	//
	// 	- antispam: text anti-spam (valid only for text)
	//
	// This parameter is required.
	//
	// example:
	//
	// ["porn","terrorism","ad"]
	Scenes *string `json:"Scenes,omitempty" xml:"Scenes,omitempty"`
}

func (s ImAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s ImAuditRequest) GoString() string {
	return s.String()
}

func (s *ImAuditRequest) GetBizType() *string {
	return s.BizType
}

func (s *ImAuditRequest) GetContents() *string {
	return s.Contents
}

func (s *ImAuditRequest) GetImages() *string {
	return s.Images
}

func (s *ImAuditRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImAuditRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImAuditRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImAuditRequest) GetScenes() *string {
	return s.Scenes
}

func (s *ImAuditRequest) SetBizType(v string) *ImAuditRequest {
	s.BizType = &v
	return s
}

func (s *ImAuditRequest) SetContents(v string) *ImAuditRequest {
	s.Contents = &v
	return s
}

func (s *ImAuditRequest) SetImages(v string) *ImAuditRequest {
	s.Images = &v
	return s
}

func (s *ImAuditRequest) SetOwnerId(v int64) *ImAuditRequest {
	s.OwnerId = &v
	return s
}

func (s *ImAuditRequest) SetResourceOwnerAccount(v string) *ImAuditRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImAuditRequest) SetResourceOwnerId(v int64) *ImAuditRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImAuditRequest) SetScenes(v string) *ImAuditRequest {
	s.Scenes = &v
	return s
}

func (s *ImAuditRequest) Validate() error {
	return dara.Validate(s)
}
