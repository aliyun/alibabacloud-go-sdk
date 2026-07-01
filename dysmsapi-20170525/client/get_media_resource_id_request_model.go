// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaResourceIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtendInfo(v string) *GetMediaResourceIdRequest
	GetExtendInfo() *string
	SetFileSize(v int64) *GetMediaResourceIdRequest
	GetFileSize() *int64
	SetMemo(v string) *GetMediaResourceIdRequest
	GetMemo() *string
	SetOssKey(v string) *GetMediaResourceIdRequest
	GetOssKey() *string
	SetResourceType(v int32) *GetMediaResourceIdRequest
	GetResourceType() *int32
}

type GetMediaResourceIdRequest struct {
	// The extended field.
	//
	// > Required when the resource type is **image**.
	//
	// example:
	//
	// {\\"img_rate\\":\\"oneToOne\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The file size. Unit: bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The description of the uploaded resource.
	//
	// example:
	//
	// 图片信息
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// The address of the resource to retrieve.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://alicom-fc-media/1947741454322274/alicom-fc-media/pic/202205191526575398603697152.png
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The resource type.
	//
	// - **1**: Text
	//
	// - **2**: Image. Small images must not exceed 100 KB. Large images must not exceed 2 MB. Images must be clear. Supported formats: JPG, JPEG, PNG.
	//
	// - **3**: Audio
	//
	// - **4**: Video. Supported format: MP4.
	//
	// >
	//
	// > - If the resource type is image, **img_rate*	- is required.
	//
	// > - 1:1 ratio: oneToOne
	//
	// > - 16:9 ratio: sixteenToNine
	//
	// > - 3:1 ratio: threeToOne
	//
	// > - 48:65 ratio: fortyEightToSixtyFiv.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetMediaResourceIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaResourceIdRequest) GoString() string {
	return s.String()
}

func (s *GetMediaResourceIdRequest) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *GetMediaResourceIdRequest) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetMediaResourceIdRequest) GetMemo() *string {
	return s.Memo
}

func (s *GetMediaResourceIdRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *GetMediaResourceIdRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *GetMediaResourceIdRequest) SetExtendInfo(v string) *GetMediaResourceIdRequest {
	s.ExtendInfo = &v
	return s
}

func (s *GetMediaResourceIdRequest) SetFileSize(v int64) *GetMediaResourceIdRequest {
	s.FileSize = &v
	return s
}

func (s *GetMediaResourceIdRequest) SetMemo(v string) *GetMediaResourceIdRequest {
	s.Memo = &v
	return s
}

func (s *GetMediaResourceIdRequest) SetOssKey(v string) *GetMediaResourceIdRequest {
	s.OssKey = &v
	return s
}

func (s *GetMediaResourceIdRequest) SetResourceType(v int32) *GetMediaResourceIdRequest {
	s.ResourceType = &v
	return s
}

func (s *GetMediaResourceIdRequest) Validate() error {
	return dara.Validate(s)
}
