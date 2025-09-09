// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckExpandStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *CheckExpandStatusRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *CheckExpandStatusRequest
	GetDrdsInstanceId() *string
}

type CheckExpandStatusRequest struct {
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
	// drds***********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s CheckExpandStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckExpandStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckExpandStatusRequest) GetDbName() *string {
	return s.DbName
}

func (s *CheckExpandStatusRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *CheckExpandStatusRequest) SetDbName(v string) *CheckExpandStatusRequest {
	s.DbName = &v
	return s
}

func (s *CheckExpandStatusRequest) SetDrdsInstanceId(v string) *CheckExpandStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CheckExpandStatusRequest) Validate() error {
	return dara.Validate(s)
}
