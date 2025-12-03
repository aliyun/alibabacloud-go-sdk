// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreSchemaDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeRestoreSchemaDetailsRequest
	GetClusterId() *string
	SetPageNumber(v int32) *DescribeRestoreSchemaDetailsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRestoreSchemaDetailsRequest
	GetPageSize() *int32
	SetRestoreRecordId(v string) *DescribeRestoreSchemaDetailsRequest
	GetRestoreRecordId() *string
}

type DescribeRestoreSchemaDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020110514xxxx
	RestoreRecordId *string `json:"RestoreRecordId,omitempty" xml:"RestoreRecordId,omitempty"`
}

func (s DescribeRestoreSchemaDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreSchemaDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSchemaDetailsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRestoreSchemaDetailsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRestoreSchemaDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreSchemaDetailsRequest) GetRestoreRecordId() *string {
	return s.RestoreRecordId
}

func (s *DescribeRestoreSchemaDetailsRequest) SetClusterId(v string) *DescribeRestoreSchemaDetailsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsRequest) SetPageNumber(v int32) *DescribeRestoreSchemaDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsRequest) SetPageSize(v int32) *DescribeRestoreSchemaDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsRequest) SetRestoreRecordId(v string) *DescribeRestoreSchemaDetailsRequest {
	s.RestoreRecordId = &v
	return s
}

func (s *DescribeRestoreSchemaDetailsRequest) Validate() error {
	return dara.Validate(s)
}
