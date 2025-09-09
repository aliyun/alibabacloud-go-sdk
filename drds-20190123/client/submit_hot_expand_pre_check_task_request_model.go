// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotExpandPreCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstType(v string) *SubmitHotExpandPreCheckTaskRequest
	GetDbInstType() *string
	SetDbName(v string) *SubmitHotExpandPreCheckTaskRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *SubmitHotExpandPreCheckTaskRequest
	GetDrdsInstanceId() *string
	SetTableList(v []*string) *SubmitHotExpandPreCheckTaskRequest
	GetTableList() []*string
}

type SubmitHotExpandPreCheckTaskRequest struct {
	// The type of the database. Valid values:
	//
	// 	- RDS
	//
	// 	- PolarDB
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
	// drd*********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableList []*string `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Repeated"`
}

func (s SubmitHotExpandPreCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotExpandPreCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandPreCheckTaskRequest) GetDbInstType() *string {
	return s.DbInstType
}

func (s *SubmitHotExpandPreCheckTaskRequest) GetDbName() *string {
	return s.DbName
}

func (s *SubmitHotExpandPreCheckTaskRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SubmitHotExpandPreCheckTaskRequest) GetTableList() []*string {
	return s.TableList
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetDbInstType(v string) *SubmitHotExpandPreCheckTaskRequest {
	s.DbInstType = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetDbName(v string) *SubmitHotExpandPreCheckTaskRequest {
	s.DbName = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetDrdsInstanceId(v string) *SubmitHotExpandPreCheckTaskRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) SetTableList(v []*string) *SubmitHotExpandPreCheckTaskRequest {
	s.TableList = v
	return s
}

func (s *SubmitHotExpandPreCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
