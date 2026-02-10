// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOSSMultimodalFineTuneDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteOSSMultimodalFineTuneDatasetRequest
	GetDBClusterId() *string
	SetDatasetId(v string) *DeleteOSSMultimodalFineTuneDatasetRequest
	GetDatasetId() *string
	SetOssUrl(v string) *DeleteOSSMultimodalFineTuneDatasetRequest
	GetOssUrl() *string
}

type DeleteOSSMultimodalFineTuneDatasetRequest struct {
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

func (s DeleteOSSMultimodalFineTuneDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOSSMultimodalFineTuneDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteOSSMultimodalFineTuneDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteOSSMultimodalFineTuneDatasetRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DeleteOSSMultimodalFineTuneDatasetRequest) GetOssUrl() *string {
	return s.OssUrl
}

func (s *DeleteOSSMultimodalFineTuneDatasetRequest) SetDBClusterId(v string) *DeleteOSSMultimodalFineTuneDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteOSSMultimodalFineTuneDatasetRequest) SetDatasetId(v string) *DeleteOSSMultimodalFineTuneDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *DeleteOSSMultimodalFineTuneDatasetRequest) SetOssUrl(v string) *DeleteOSSMultimodalFineTuneDatasetRequest {
	s.OssUrl = &v
	return s
}

func (s *DeleteOSSMultimodalFineTuneDatasetRequest) Validate() error {
	return dara.Validate(s)
}
