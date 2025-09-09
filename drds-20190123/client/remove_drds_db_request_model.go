// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDrdsDbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *RemoveDrdsDbRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *RemoveDrdsDbRequest
	GetDrdsInstanceId() *string
}

type RemoveDrdsDbRequest struct {
	// The name of the database you want to back up.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DRDS instance to which the destination database belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s RemoveDrdsDbRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDrdsDbRequest) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbRequest) GetDbName() *string {
	return s.DbName
}

func (s *RemoveDrdsDbRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *RemoveDrdsDbRequest) SetDbName(v string) *RemoveDrdsDbRequest {
	s.DbName = &v
	return s
}

func (s *RemoveDrdsDbRequest) SetDrdsInstanceId(v string) *RemoveDrdsDbRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RemoveDrdsDbRequest) Validate() error {
	return dara.Validate(s)
}
