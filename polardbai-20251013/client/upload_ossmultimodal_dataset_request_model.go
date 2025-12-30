// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadOSSMultimodalDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UploadOSSMultimodalDatasetRequest
	GetDBClusterId() *string
	SetDatasetId(v string) *UploadOSSMultimodalDatasetRequest
	GetDatasetId() *string
	SetOssUrl(v string) *UploadOSSMultimodalDatasetRequest
	GetOssUrl() *string
}

type UploadOSSMultimodalDatasetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// ds-*****ab0
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// https://bucket-name.oss-cn-beijing.aliyuncs.com/005ed960-6a07-11ef-ab81-f32551c4afe8
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
}

func (s UploadOSSMultimodalDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadOSSMultimodalDatasetRequest) GoString() string {
	return s.String()
}

func (s *UploadOSSMultimodalDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UploadOSSMultimodalDatasetRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *UploadOSSMultimodalDatasetRequest) GetOssUrl() *string {
	return s.OssUrl
}

func (s *UploadOSSMultimodalDatasetRequest) SetDBClusterId(v string) *UploadOSSMultimodalDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *UploadOSSMultimodalDatasetRequest) SetDatasetId(v string) *UploadOSSMultimodalDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *UploadOSSMultimodalDatasetRequest) SetOssUrl(v string) *UploadOSSMultimodalDatasetRequest {
	s.OssUrl = &v
	return s
}

func (s *UploadOSSMultimodalDatasetRequest) Validate() error {
	return dara.Validate(s)
}
