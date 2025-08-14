// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletePhysicalFiles(v bool) *DeleteMediaInfosRequest
	GetDeletePhysicalFiles() *bool
	SetInputURLs(v string) *DeleteMediaInfosRequest
	GetInputURLs() *string
	SetMediaIds(v string) *DeleteMediaInfosRequest
	GetMediaIds() *string
}

type DeleteMediaInfosRequest struct {
	// Specifies whether to delete the physical file of the media asset.
	//
	// If the media asset is stored in your own OSS bucket, you must authorize the service role AliyunICEDefaultRole in advance. For more information<props="china">, see [Authorize IMS to delete recording files in OSS](https://help.aliyun.com/zh/ims/user-guide/record?spm=a2c4g.11186623.0.i8#0737d9c437bmn).
	//
	// example:
	//
	// false
	DeletePhysicalFiles *bool `json:"DeletePhysicalFiles,omitempty" xml:"DeletePhysicalFiles,omitempty"`
	// The URL of the media asset that you want to delete. The file corresponding to the URL must be registered with IMS. Separate multiple URLs with commas (,). The following two formats are supported:
	//
	// 1.  http(s)://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4?
	//
	// 2.  OSS://example-bucket/example.mp4?\\
	//
	//     In this format, it is considered by default that the region of the OSS bucket in which the media asset resides is the same as the region in which IMS is activated.
	InputURLs *string `json:"InputURLs,omitempty" xml:"InputURLs,omitempty"`
	// The ID of the media asset that you want to delete from Intelligent Media Services (IMS).
	//
	// 	- Separate multiple IDs with commas (,).
	//
	// If you leave MediaIds empty, you must specify InputURLs.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****,****15d4a4b0448391508f2cb486****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
}

func (s DeleteMediaInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaInfosRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaInfosRequest) GetDeletePhysicalFiles() *bool {
	return s.DeletePhysicalFiles
}

func (s *DeleteMediaInfosRequest) GetInputURLs() *string {
	return s.InputURLs
}

func (s *DeleteMediaInfosRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *DeleteMediaInfosRequest) SetDeletePhysicalFiles(v bool) *DeleteMediaInfosRequest {
	s.DeletePhysicalFiles = &v
	return s
}

func (s *DeleteMediaInfosRequest) SetInputURLs(v string) *DeleteMediaInfosRequest {
	s.InputURLs = &v
	return s
}

func (s *DeleteMediaInfosRequest) SetMediaIds(v string) *DeleteMediaInfosRequest {
	s.MediaIds = &v
	return s
}

func (s *DeleteMediaInfosRequest) Validate() error {
	return dara.Validate(s)
}
