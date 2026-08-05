// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeletePolarFsRequest
	GetDBClusterId() *string
	SetPolarFsInstanceId(v string) *DeletePolarFsRequest
	GetPolarFsInstanceId() *string
}

type DeletePolarFsRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster IDs of all Data Warehouse Edition clusters in a specific region.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The PolarFs instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcs-2ze0i74ka607wck3
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s DeletePolarFsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsRequest) GoString() string {
	return s.String()
}

func (s *DeletePolarFsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeletePolarFsRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DeletePolarFsRequest) SetDBClusterId(v string) *DeletePolarFsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeletePolarFsRequest) SetPolarFsInstanceId(v string) *DeletePolarFsRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DeletePolarFsRequest) Validate() error {
	return dara.Validate(s)
}
