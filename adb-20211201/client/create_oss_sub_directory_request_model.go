// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssSubDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateOssSubDirectoryRequest
	GetDBClusterId() *string
	SetPath(v string) *CreateOssSubDirectoryRequest
	GetPath() *string
}

type CreateOssSubDirectoryRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the information about all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp149vz49b36t****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The OSS path where you want to create a subdirectory.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://testBucketName/das_lakehouse
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateOssSubDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOssSubDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateOssSubDirectoryRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateOssSubDirectoryRequest) GetPath() *string {
	return s.Path
}

func (s *CreateOssSubDirectoryRequest) SetDBClusterId(v string) *CreateOssSubDirectoryRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateOssSubDirectoryRequest) SetPath(v string) *CreateOssSubDirectoryRequest {
	s.Path = &v
	return s
}

func (s *CreateOssSubDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
