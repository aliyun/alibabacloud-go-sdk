// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDrdsDbFailedRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *RemoveDrdsDbFailedRecordRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *RemoveDrdsDbFailedRecordRequest
	GetDrdsInstanceId() *string
}

type RemoveDrdsDbFailedRecordRequest struct {
	// The name of the DRDS database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the ApsaraDB RDS for PostgreSQL instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveDrdsDbFailedRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDrdsDbFailedRecordRequest) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbFailedRecordRequest) GetDbName() *string {
	return s.DbName
}

func (s *RemoveDrdsDbFailedRecordRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *RemoveDrdsDbFailedRecordRequest) SetDbName(v string) *RemoveDrdsDbFailedRecordRequest {
	s.DbName = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordRequest) SetDrdsInstanceId(v string) *RemoveDrdsDbFailedRecordRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordRequest) Validate() error {
	return dara.Validate(s)
}
