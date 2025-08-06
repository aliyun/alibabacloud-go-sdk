// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalWatermarkExtractResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtractType(v string) *GetDigitalWatermarkExtractResultRequest
	GetExtractType() *string
	SetJobId(v string) *GetDigitalWatermarkExtractResultRequest
	GetJobId() *string
	SetMediaId(v string) *GetDigitalWatermarkExtractResultRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *GetDigitalWatermarkExtractResultRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetDigitalWatermarkExtractResultRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetDigitalWatermarkExtractResultRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetDigitalWatermarkExtractResultRequest
	GetResourceOwnerId() *string
}

type GetDigitalWatermarkExtractResultRequest struct {
	// The type of the watermark. Valid values:
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
	// The ID of the watermark extraction job.
	//
	// 	- You can obtain the ID from the response to the [SubmitDigitalWatermarkExtractJob](~~SubmitDigitalWatermarkExtractJob~~) operation.
	//
	// 	- If you specify this parameter, the result of the specified watermark extraction job is returned. If you leave this parameter empty, the results of all watermark extraction jobs submitted for the video are returned.
	//
	// example:
	//
	// 2bf4390af9e5491c09cc720ad****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the video from which you want to query the watermark content. You can specify only one ID. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). In the left-side navigation pane, choose **Media Files*	- > **Audio/Video*	- to view the video ID.
	//
	// 	- Obtain the VideoId from the response to the [SearchMedia](~~SearchMedia~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 93ab850b4f6f54b6e91d24d81d4****
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetDigitalWatermarkExtractResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalWatermarkExtractResultRequest) GoString() string {
	return s.String()
}

func (s *GetDigitalWatermarkExtractResultRequest) GetExtractType() *string {
	return s.ExtractType
}

func (s *GetDigitalWatermarkExtractResultRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetDigitalWatermarkExtractResultRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetDigitalWatermarkExtractResultRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetDigitalWatermarkExtractResultRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetDigitalWatermarkExtractResultRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetDigitalWatermarkExtractResultRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetDigitalWatermarkExtractResultRequest) SetExtractType(v string) *GetDigitalWatermarkExtractResultRequest {
	s.ExtractType = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultRequest) SetJobId(v string) *GetDigitalWatermarkExtractResultRequest {
	s.JobId = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultRequest) SetMediaId(v string) *GetDigitalWatermarkExtractResultRequest {
	s.MediaId = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultRequest) SetOwnerAccount(v string) *GetDigitalWatermarkExtractResultRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultRequest) SetOwnerId(v string) *GetDigitalWatermarkExtractResultRequest {
	s.OwnerId = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultRequest) SetResourceOwnerAccount(v string) *GetDigitalWatermarkExtractResultRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultRequest) SetResourceOwnerId(v string) *GetDigitalWatermarkExtractResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultRequest) Validate() error {
	return dara.Validate(s)
}
