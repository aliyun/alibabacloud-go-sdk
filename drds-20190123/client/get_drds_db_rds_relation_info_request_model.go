// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDrdsDbRdsRelationInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *GetDrdsDbRdsRelationInfoRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *GetDrdsDbRdsRelationInfoRequest
	GetDrdsInstanceId() *string
}

type GetDrdsDbRdsRelationInfoRequest struct {
	// The name of the DRDS database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testDb
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbga948vbpd2
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s GetDrdsDbRdsRelationInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDrdsDbRdsRelationInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDrdsDbRdsRelationInfoRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetDrdsDbRdsRelationInfoRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *GetDrdsDbRdsRelationInfoRequest) SetDbName(v string) *GetDrdsDbRdsRelationInfoRequest {
	s.DbName = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoRequest) SetDrdsInstanceId(v string) *GetDrdsDbRdsRelationInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *GetDrdsDbRdsRelationInfoRequest) Validate() error {
	return dara.Validate(s)
}
