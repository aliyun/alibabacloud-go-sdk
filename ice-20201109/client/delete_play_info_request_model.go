// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlayInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletePhysicalFiles(v bool) *DeletePlayInfoRequest
	GetDeletePhysicalFiles() *bool
	SetFileURLs(v string) *DeletePlayInfoRequest
	GetFileURLs() *string
	SetMediaId(v string) *DeletePlayInfoRequest
	GetMediaId() *string
}

type DeletePlayInfoRequest struct {
	// Specifies whether to delete the physical file of the media stream.
	//
	// If the media asset is stored in your own Object Storage Service (OSS) bucket, you must authorize the service role AliyunICEDefaultRole in advance. <props="china">For more information, see [Authorize IMS to delete recording files in OSS](https://help.aliyun.com/document_detail/449331.html#p-ko2-wc7-iad).
	//
	// You can delete only the physical files of transcoded streams, but not the physical files of source files.
	//
	// example:
	//
	// false
	DeletePhysicalFiles *bool `json:"DeletePhysicalFiles,omitempty" xml:"DeletePhysicalFiles,omitempty"`
	// The URL of the media stream file that you want to delete. Separate multiple URLs with commas (,).
	//
	// example:
	//
	// https://ice-test001.oss-cn-shanghai.aliyuncs.com/%E6%8E%A5%E5%8F%A3%E6%B5%8B%E8%AF%95/%E5%B0%8F%E7%8C%AA%E4%BD%A9%E5%A5%87640*360.mp4
	FileURLs *string `json:"FileURLs,omitempty" xml:"FileURLs,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 1d3518e0027d71ed80cd909598416303
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s DeletePlayInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePlayInfoRequest) GoString() string {
	return s.String()
}

func (s *DeletePlayInfoRequest) GetDeletePhysicalFiles() *bool {
	return s.DeletePhysicalFiles
}

func (s *DeletePlayInfoRequest) GetFileURLs() *string {
	return s.FileURLs
}

func (s *DeletePlayInfoRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *DeletePlayInfoRequest) SetDeletePhysicalFiles(v bool) *DeletePlayInfoRequest {
	s.DeletePhysicalFiles = &v
	return s
}

func (s *DeletePlayInfoRequest) SetFileURLs(v string) *DeletePlayInfoRequest {
	s.FileURLs = &v
	return s
}

func (s *DeletePlayInfoRequest) SetMediaId(v string) *DeletePlayInfoRequest {
	s.MediaId = &v
	return s
}

func (s *DeletePlayInfoRequest) Validate() error {
	return dara.Validate(s)
}
