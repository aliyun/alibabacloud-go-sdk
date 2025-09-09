// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHotDbListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeHotDbListRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeHotDbListRequest
	GetDrdsInstanceId() *string
}

type DescribeHotDbListRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds**********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DescribeHotDbListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotDbListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeHotDbListRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeHotDbListRequest) SetDbName(v string) *DescribeHotDbListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeHotDbListRequest) SetDrdsInstanceId(v string) *DescribeHotDbListRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeHotDbListRequest) Validate() error {
	return dara.Validate(s)
}
