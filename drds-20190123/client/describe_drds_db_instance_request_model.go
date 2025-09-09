// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDbInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstanceId(v string) *DescribeDrdsDbInstanceRequest
	GetDbInstanceId() *string
	SetDbName(v string) *DescribeDrdsDbInstanceRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeDrdsDbInstanceRequest
	GetDrdsInstanceId() *string
}

type DescribeDrdsDbInstanceRequest struct {
	// The ID of the custom ApsaraDB RDS for MySQL instance that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1t1mk5a5bdj****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The name of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbga1138****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsDbInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeDrdsDbInstanceRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsDbInstanceRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsDbInstanceRequest) SetDbInstanceId(v string) *DescribeDrdsDbInstanceRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceRequest) SetDbName(v string) *DescribeDrdsDbInstanceRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDbInstanceRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceRequest) Validate() error {
	return dara.Validate(s)
}
