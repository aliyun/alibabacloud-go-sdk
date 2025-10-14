// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachColumnarInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *AttachColumnarInstanceRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *AttachColumnarInstanceRequest
	GetRegionId() *string
}

type AttachColumnarInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachColumnarInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachColumnarInstanceRequest) GoString() string {
	return s.String()
}

func (s *AttachColumnarInstanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *AttachColumnarInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachColumnarInstanceRequest) SetDBInstanceName(v string) *AttachColumnarInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *AttachColumnarInstanceRequest) SetRegionId(v string) *AttachColumnarInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *AttachColumnarInstanceRequest) Validate() error {
	return dara.Validate(s)
}
