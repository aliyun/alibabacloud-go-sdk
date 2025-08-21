// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceCaptureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ModifyDeviceCaptureRequest
	GetId() *string
	SetImage(v int32) *ModifyDeviceCaptureRequest
	GetImage() *int32
	SetOwnerId(v int64) *ModifyDeviceCaptureRequest
	GetOwnerId() *int64
	SetVideo(v int32) *ModifyDeviceCaptureRequest
	GetVideo() *int32
}

type ModifyDeviceCaptureRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	Image   *int32 `json:"Image,omitempty" xml:"Image,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	Video *int32 `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s ModifyDeviceCaptureRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceCaptureRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceCaptureRequest) GetId() *string {
	return s.Id
}

func (s *ModifyDeviceCaptureRequest) GetImage() *int32 {
	return s.Image
}

func (s *ModifyDeviceCaptureRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDeviceCaptureRequest) GetVideo() *int32 {
	return s.Video
}

func (s *ModifyDeviceCaptureRequest) SetId(v string) *ModifyDeviceCaptureRequest {
	s.Id = &v
	return s
}

func (s *ModifyDeviceCaptureRequest) SetImage(v int32) *ModifyDeviceCaptureRequest {
	s.Image = &v
	return s
}

func (s *ModifyDeviceCaptureRequest) SetOwnerId(v int64) *ModifyDeviceCaptureRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDeviceCaptureRequest) SetVideo(v int32) *ModifyDeviceCaptureRequest {
	s.Video = &v
	return s
}

func (s *ModifyDeviceCaptureRequest) Validate() error {
	return dara.Validate(s)
}
