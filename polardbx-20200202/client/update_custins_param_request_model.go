// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustinsParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *UpdateCustinsParamRequest
	GetDBInstanceName() *string
	SetName(v string) *UpdateCustinsParamRequest
	GetName() *string
	SetRegionId(v string) *UpdateCustinsParamRequest
	GetRegionId() *string
	SetValue(v string) *UpdateCustinsParamRequest
	GetValue() *string
}

type UpdateCustinsParamRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// load_test_4009266
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateCustinsParamRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustinsParamRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustinsParamRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpdateCustinsParamRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCustinsParamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCustinsParamRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateCustinsParamRequest) SetDBInstanceName(v string) *UpdateCustinsParamRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateCustinsParamRequest) SetName(v string) *UpdateCustinsParamRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustinsParamRequest) SetRegionId(v string) *UpdateCustinsParamRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCustinsParamRequest) SetValue(v string) *UpdateCustinsParamRequest {
	s.Value = &v
	return s
}

func (s *UpdateCustinsParamRequest) Validate() error {
	return dara.Validate(s)
}
