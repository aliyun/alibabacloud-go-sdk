// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDrdsDbNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *CheckDrdsDbNameRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *CheckDrdsDbNameRequest
	GetDrdsInstanceId() *string
}

type CheckDrdsDbNameRequest struct {
	// DRDS database name
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// DRDS instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// drds********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s CheckDrdsDbNameRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDrdsDbNameRequest) GoString() string {
	return s.String()
}

func (s *CheckDrdsDbNameRequest) GetDbName() *string {
	return s.DbName
}

func (s *CheckDrdsDbNameRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *CheckDrdsDbNameRequest) SetDbName(v string) *CheckDrdsDbNameRequest {
	s.DbName = &v
	return s
}

func (s *CheckDrdsDbNameRequest) SetDrdsInstanceId(v string) *CheckDrdsDbNameRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CheckDrdsDbNameRequest) Validate() error {
	return dara.Validate(s)
}
