// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstDbLogInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeInstDbLogInfoRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeInstDbLogInfoRequest
	GetDrdsInstanceId() *string
}

type DescribeInstDbLogInfoRequest struct {
	// The name of the DRDS database.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeInstDbLogInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstDbLogInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeInstDbLogInfoRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeInstDbLogInfoRequest) SetDbName(v string) *DescribeInstDbLogInfoRequest {
	s.DbName = &v
	return s
}

func (s *DescribeInstDbLogInfoRequest) SetDrdsInstanceId(v string) *DescribeInstDbLogInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeInstDbLogInfoRequest) Validate() error {
	return dara.Validate(s)
}
