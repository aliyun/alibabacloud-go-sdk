// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransformOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateTransformOperationRequest
	GetDBInstanceName() *string
	SetOperation(v string) *CreateTransformOperationRequest
	GetOperation() *string
	SetRegionId(v string) *CreateTransformOperationRequest
	GetRegionId() *string
}

type CreateTransformOperationRequest struct {
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// finish
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateTransformOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransformOperationRequest) GoString() string {
	return s.String()
}

func (s *CreateTransformOperationRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateTransformOperationRequest) GetOperation() *string {
	return s.Operation
}

func (s *CreateTransformOperationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTransformOperationRequest) SetDBInstanceName(v string) *CreateTransformOperationRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateTransformOperationRequest) SetOperation(v string) *CreateTransformOperationRequest {
	s.Operation = &v
	return s
}

func (s *CreateTransformOperationRequest) SetRegionId(v string) *CreateTransformOperationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransformOperationRequest) Validate() error {
	return dara.Validate(s)
}
