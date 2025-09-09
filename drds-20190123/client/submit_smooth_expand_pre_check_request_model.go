// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmoothExpandPreCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstType(v string) *SubmitSmoothExpandPreCheckRequest
	GetDbInstType() *string
	SetDbName(v string) *SubmitSmoothExpandPreCheckRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *SubmitSmoothExpandPreCheckRequest
	GetDrdsInstanceId() *string
}

type SubmitSmoothExpandPreCheckRequest struct {
	// The type of the database. Valid values:
	//
	// 	- RDS
	//
	// 	- POLARDB
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The name of the PolarDB-X database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*******
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s SubmitSmoothExpandPreCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmoothExpandPreCheckRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmoothExpandPreCheckRequest) GetDbInstType() *string {
	return s.DbInstType
}

func (s *SubmitSmoothExpandPreCheckRequest) GetDbName() *string {
	return s.DbName
}

func (s *SubmitSmoothExpandPreCheckRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SubmitSmoothExpandPreCheckRequest) SetDbInstType(v string) *SubmitSmoothExpandPreCheckRequest {
	s.DbInstType = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckRequest) SetDbName(v string) *SubmitSmoothExpandPreCheckRequest {
	s.DbName = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckRequest) SetDrdsInstanceId(v string) *SubmitSmoothExpandPreCheckRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitSmoothExpandPreCheckRequest) Validate() error {
	return dara.Validate(s)
}
