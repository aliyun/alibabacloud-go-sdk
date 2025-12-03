// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreFullDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeRestoreFullDetailsRequest
	GetClusterId() *string
	SetPageNumber(v int32) *DescribeRestoreFullDetailsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRestoreFullDetailsRequest
	GetPageSize() *int32
	SetRestoreRecordId(v string) *DescribeRestoreFullDetailsRequest
	GetRestoreRecordId() *string
}

type DescribeRestoreFullDetailsRequest struct {
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

func (s DescribeRestoreFullDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreFullDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreFullDetailsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRestoreFullDetailsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRestoreFullDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreFullDetailsRequest) GetRestoreRecordId() *string {
	return s.RestoreRecordId
}

func (s *DescribeRestoreFullDetailsRequest) SetClusterId(v string) *DescribeRestoreFullDetailsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreFullDetailsRequest) SetPageNumber(v int32) *DescribeRestoreFullDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreFullDetailsRequest) SetPageSize(v int32) *DescribeRestoreFullDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreFullDetailsRequest) SetRestoreRecordId(v string) *DescribeRestoreFullDetailsRequest {
	s.RestoreRecordId = &v
	return s
}

func (s *DescribeRestoreFullDetailsRequest) Validate() error {
	return dara.Validate(s)
}
