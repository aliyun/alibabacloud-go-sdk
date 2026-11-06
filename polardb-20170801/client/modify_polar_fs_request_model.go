// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarFsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomOssAk(v string) *ModifyPolarFsRequest
	GetCustomOssAk() *string
	SetCustomOssSk(v string) *ModifyPolarFsRequest
	GetCustomOssSk() *string
	SetDBClusterId(v string) *ModifyPolarFsRequest
	GetDBClusterId() *string
	SetPolarFsInstanceId(v string) *ModifyPolarFsRequest
	GetPolarFsInstanceId() *string
}

type ModifyPolarFsRequest struct {
	// The custom OSS AccessKey ID. You must specify this parameter together with CustomOssSk when you call the ModifyPolarFs operation. Otherwise, the InvalidCustomIdentityInfo.Malformed (400) error is returned.
	//
	// example:
	//
	// LTAI5tExample1234
	CustomOssAk *string `json:"CustomOssAk,omitempty" xml:"CustomOssAk,omitempty"`
	// The custom OSS AccessKey Secret. You must specify this parameter together with CustomOssAk when you call the ModifyPolarFs operation. Otherwise, the InvalidCustomIdentityInfo.Malformed (400) error is returned.
	//
	// example:
	//
	// example-oss-secret
	CustomOssSk *string `json:"CustomOssSk,omitempty" xml:"CustomOssSk,omitempty"`
	// The cluster ID.
	//
	// >You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query information about all clusters in the specified region, including cluster IDs.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The PolarFs instance ID.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
}

func (s ModifyPolarFsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarFsRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolarFsRequest) GetCustomOssAk() *string {
	return s.CustomOssAk
}

func (s *ModifyPolarFsRequest) GetCustomOssSk() *string {
	return s.CustomOssSk
}

func (s *ModifyPolarFsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyPolarFsRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *ModifyPolarFsRequest) SetCustomOssAk(v string) *ModifyPolarFsRequest {
	s.CustomOssAk = &v
	return s
}

func (s *ModifyPolarFsRequest) SetCustomOssSk(v string) *ModifyPolarFsRequest {
	s.CustomOssSk = &v
	return s
}

func (s *ModifyPolarFsRequest) SetDBClusterId(v string) *ModifyPolarFsRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyPolarFsRequest) SetPolarFsInstanceId(v string) *ModifyPolarFsRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *ModifyPolarFsRequest) Validate() error {
	return dara.Validate(s)
}
