// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDigitalWatermarkExtractJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtractType(v string) *SubmitDigitalWatermarkExtractJobRequest
	GetExtractType() *string
	SetMediaId(v string) *SubmitDigitalWatermarkExtractJobRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *SubmitDigitalWatermarkExtractJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SubmitDigitalWatermarkExtractJobRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SubmitDigitalWatermarkExtractJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SubmitDigitalWatermarkExtractJobRequest
	GetResourceOwnerId() *string
}

type SubmitDigitalWatermarkExtractJobRequest struct {
	// The type of the watermark that you want to extract. Valid values:
	//
	// 	- **TraceMark**: user-tracing watermark
	//
	// 	- **CopyrightMark**: copyright watermark
	//
	// This parameter is required.
	//
	// example:
	//
	// TraceMark
	ExtractType *string `json:"ExtractType,omitempty" xml:"ExtractType,omitempty"`
	// The ID of the video from which you want to extract the watermark. You can specify only one ID. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Media Files*	- > **Audio/Video*	- to view the video ID.
	//
	// 	- Obtain the VideoId from the response to the [SearchMedia](~~SearchMedia~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0222e203cf80f9c22870a4d2c****
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SubmitDigitalWatermarkExtractJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDigitalWatermarkExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDigitalWatermarkExtractJobRequest) GetExtractType() *string {
	return s.ExtractType
}

func (s *SubmitDigitalWatermarkExtractJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitDigitalWatermarkExtractJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitDigitalWatermarkExtractJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SubmitDigitalWatermarkExtractJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitDigitalWatermarkExtractJobRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SubmitDigitalWatermarkExtractJobRequest) SetExtractType(v string) *SubmitDigitalWatermarkExtractJobRequest {
	s.ExtractType = &v
	return s
}

func (s *SubmitDigitalWatermarkExtractJobRequest) SetMediaId(v string) *SubmitDigitalWatermarkExtractJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitDigitalWatermarkExtractJobRequest) SetOwnerAccount(v string) *SubmitDigitalWatermarkExtractJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitDigitalWatermarkExtractJobRequest) SetOwnerId(v string) *SubmitDigitalWatermarkExtractJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitDigitalWatermarkExtractJobRequest) SetResourceOwnerAccount(v string) *SubmitDigitalWatermarkExtractJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitDigitalWatermarkExtractJobRequest) SetResourceOwnerId(v string) *SubmitDigitalWatermarkExtractJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitDigitalWatermarkExtractJobRequest) Validate() error {
	return dara.Validate(s)
}
