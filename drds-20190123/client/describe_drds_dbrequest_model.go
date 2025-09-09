// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeDrdsDBRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeDrdsDBRequest
	GetDrdsInstanceId() *string
}

type DescribeDrdsDBRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeDrdsDBRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsDBRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsDBRequest) SetDbName(v string) *DescribeDrdsDBRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDBRequest) Validate() error {
	return dara.Validate(s)
}
