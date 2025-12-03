// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeRestoreTablesRequest
	GetClusterId() *string
	SetRestoreRecordId(v string) *DescribeRestoreTablesRequest
	GetRestoreRecordId() *string
}

type DescribeRestoreTablesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020110514xxxx
	RestoreRecordId *string `json:"RestoreRecordId,omitempty" xml:"RestoreRecordId,omitempty"`
}

func (s DescribeRestoreTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTablesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRestoreTablesRequest) GetRestoreRecordId() *string {
	return s.RestoreRecordId
}

func (s *DescribeRestoreTablesRequest) SetClusterId(v string) *DescribeRestoreTablesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreTablesRequest) SetRestoreRecordId(v string) *DescribeRestoreTablesRequest {
	s.RestoreRecordId = &v
	return s
}

func (s *DescribeRestoreTablesRequest) Validate() error {
	return dara.Validate(s)
}
