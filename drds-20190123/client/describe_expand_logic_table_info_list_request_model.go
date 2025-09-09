// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpandLogicTableInfoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeExpandLogicTableInfoListRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeExpandLogicTableInfoListRequest
	GetDrdsInstanceId() *string
}

type DescribeExpandLogicTableInfoListRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_flashback
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeExpandLogicTableInfoListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeExpandLogicTableInfoListRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeExpandLogicTableInfoListRequest) SetDbName(v string) *DescribeExpandLogicTableInfoListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListRequest) SetDrdsInstanceId(v string) *DescribeExpandLogicTableInfoListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListRequest) Validate() error {
	return dara.Validate(s)
}
