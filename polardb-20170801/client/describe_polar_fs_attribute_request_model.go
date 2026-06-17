// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribePolarFsAttributeRequest
	GetDBClusterId() *string
	SetPolarFsInstanceId(v string) *DescribePolarFsAttributeRequest
	GetPolarFsInstanceId() *string
	SetQueryFuseMountInfo(v bool) *DescribePolarFsAttributeRequest
	GetQueryFuseMountInfo() *bool
}

type DescribePolarFsAttributeRequest struct {
	// The cluster ID.
	//
	// > You can find cluster IDs by calling the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation.
	//
	// example:
	//
	// pc-bp1q76364ird*****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the PolarFS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// Specifies whether to query the FUSE mount information. Valid values:
	//
	// - **true**: Queries the FUSE mount information.
	//
	// - **false**: Does not query the FUSE mount information. This is the default value.
	//
	// example:
	//
	// false
	QueryFuseMountInfo *bool `json:"QueryFuseMountInfo,omitempty" xml:"QueryFuseMountInfo,omitempty"`
}

func (s DescribePolarFsAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarFsAttributeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePolarFsAttributeRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribePolarFsAttributeRequest) GetQueryFuseMountInfo() *bool {
	return s.QueryFuseMountInfo
}

func (s *DescribePolarFsAttributeRequest) SetDBClusterId(v string) *DescribePolarFsAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribePolarFsAttributeRequest) SetPolarFsInstanceId(v string) *DescribePolarFsAttributeRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribePolarFsAttributeRequest) SetQueryFuseMountInfo(v bool) *DescribePolarFsAttributeRequest {
	s.QueryFuseMountInfo = &v
	return s
}

func (s *DescribePolarFsAttributeRequest) Validate() error {
	return dara.Validate(s)
}
