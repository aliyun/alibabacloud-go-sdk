// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstDbSlsInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeInstDbSlsInfoRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeInstDbSlsInfoRequest
	GetDrdsInstanceId() *string
}

type DescribeInstDbSlsInfoRequest struct {
	// The name of the database.
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

func (s DescribeInstDbSlsInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstDbSlsInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeInstDbSlsInfoRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeInstDbSlsInfoRequest) SetDbName(v string) *DescribeInstDbSlsInfoRequest {
	s.DbName = &v
	return s
}

func (s *DescribeInstDbSlsInfoRequest) SetDrdsInstanceId(v string) *DescribeInstDbSlsInfoRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeInstDbSlsInfoRequest) Validate() error {
	return dara.Validate(s)
}
