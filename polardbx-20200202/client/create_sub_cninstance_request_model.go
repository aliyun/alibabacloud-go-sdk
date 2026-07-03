// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCNInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateSubCNInstanceRequest
	GetDBInstanceName() *string
	SetIsAutoCreate(v bool) *CreateSubCNInstanceRequest
	GetIsAutoCreate() *bool
	SetReadType(v string) *CreateSubCNInstanceRequest
	GetReadType() *string
	SetRegionId(v string) *CreateSubCNInstanceRequest
	GetRegionId() *string
}

type CreateSubCNInstanceRequest struct {
	// The instance ID. > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the details of all instances in the specified region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzravgpt8q****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// Specifies whether to automatically calculate compute resource parameters. Valid values:
	//
	// - **true**:
	IsAutoCreate *bool `json:"IsAutoCreate,omitempty" xml:"IsAutoCreate,omitempty"`
	// The read/write type. Valid values:
	//
	// - ReadWrite: row store read/write.
	//
	// - ColumnarRead: column store read-only.
	//
	// example:
	//
	// ReadWrite
	ReadType *string `json:"ReadType,omitempty" xml:"ReadType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateSubCNInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCNInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateSubCNInstanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateSubCNInstanceRequest) GetIsAutoCreate() *bool {
	return s.IsAutoCreate
}

func (s *CreateSubCNInstanceRequest) GetReadType() *string {
	return s.ReadType
}

func (s *CreateSubCNInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSubCNInstanceRequest) SetDBInstanceName(v string) *CreateSubCNInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateSubCNInstanceRequest) SetIsAutoCreate(v bool) *CreateSubCNInstanceRequest {
	s.IsAutoCreate = &v
	return s
}

func (s *CreateSubCNInstanceRequest) SetReadType(v string) *CreateSubCNInstanceRequest {
	s.ReadType = &v
	return s
}

func (s *CreateSubCNInstanceRequest) SetRegionId(v string) *CreateSubCNInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSubCNInstanceRequest) Validate() error {
	return dara.Validate(s)
}
