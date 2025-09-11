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
	// The extended fields.
	//
	// > If you set the ResourceType parameter to **2**, this parameter is required.
	//
	// example:
	//
	// {\\"img_rate\\":\\"oneToOne\\"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The size of the resource. Unit: bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The description of the resource.
	//
	// example:
	//
	// remark
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// The address of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://alicom-fc-media/1947741454322274/alicom-fc-media/pic/202205191526575398603697152.png
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The type of the resource.
	//
	// 	- **1**: text.
	//
	// 	- **2**: image. A small image cannot exceed 100 KB in size, and a large image cannot exceed 2 MB in size. The image must be clear. Supported format: JPG, JPEG, and PNG.
	//
	// 	- **3**: audio.
	//
	// 	- **4**: video. Supported format: MP4.
	//
	// >
	//
	// 	- If you set the ResourceType parameter to 2, the **img_rate*	- required is required. Valid values:
	//
	// 	- 1:1
	//
	// 	- 16:9
	//
	// 	- 3:1
	//
	// 	- 48:65
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
