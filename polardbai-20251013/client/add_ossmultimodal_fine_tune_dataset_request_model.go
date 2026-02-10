// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOSSMultimodalFineTuneDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *AddOSSMultimodalFineTuneDatasetRequest
	GetDBClusterId() *string
	SetDatasetId(v string) *AddOSSMultimodalFineTuneDatasetRequest
	GetDatasetId() *string
	SetOssUrl(v string) *AddOSSMultimodalFineTuneDatasetRequest
	GetOssUrl() *string
}

type AddOSSMultimodalFineTuneDatasetRequest struct {
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

func (s AddOSSMultimodalFineTuneDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s AddOSSMultimodalFineTuneDatasetRequest) GoString() string {
	return s.String()
}

func (s *AddOSSMultimodalFineTuneDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AddOSSMultimodalFineTuneDatasetRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *AddOSSMultimodalFineTuneDatasetRequest) GetOssUrl() *string {
	return s.OssUrl
}

func (s *AddOSSMultimodalFineTuneDatasetRequest) SetDBClusterId(v string) *AddOSSMultimodalFineTuneDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *AddOSSMultimodalFineTuneDatasetRequest) SetDatasetId(v string) *AddOSSMultimodalFineTuneDatasetRequest {
	s.DatasetId = &v
	return s
}

func (s *AddOSSMultimodalFineTuneDatasetRequest) SetOssUrl(v string) *AddOSSMultimodalFineTuneDatasetRequest {
	s.OssUrl = &v
	return s
}

func (s *AddOSSMultimodalFineTuneDatasetRequest) Validate() error {
	return dara.Validate(s)
}
