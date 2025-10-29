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
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId  *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	QueryFuseMountInfo *bool   `json:"QueryFuseMountInfo,omitempty" xml:"QueryFuseMountInfo,omitempty"`
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
