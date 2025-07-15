// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *DescribeSpecRequest
	GetBizRegionId() *string
	SetMatrixSpec(v string) *DescribeSpecRequest
	GetMatrixSpec() *string
	SetMaxResults(v int32) *DescribeSpecRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSpecRequest
	GetNextToken() *string
	SetSaleMode(v string) *DescribeSpecRequest
	GetSaleMode() *string
	SetSpecIds(v []*string) *DescribeSpecRequest
	GetSpecIds() []*string
	SetSpecStatus(v string) *DescribeSpecRequest
	GetSpecStatus() *string
	SetSpecType(v string) *DescribeSpecRequest
	GetSpecType() *string
}

type DescribeSpecRequest struct {
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The matrix specification.
	//
	// Valid values:
	//
	// 	- cpm.gn6.gx1
	//
	// example:
	//
	// cpm.gn6.gx1
	MatrixSpec *string `json:"MatrixSpec,omitempty" xml:"MatrixSpec,omitempty"`
	// The maximum number of items to return per page in a paginated query. The value range is 1 to 100, with a default value of 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Indicates the starting position for reading. If left empty, it starts from the beginning.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The purchase mode of cloud mobile phones.
	//
	// Valid values:
	//
	// 	- Instance (default): the instance group mode.
	//
	// 	- Node: the matrix mode [whitelisted].
	//
	// example:
	//
	// Instance
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// List of specification IDs.
	SpecIds []*string `json:"SpecIds,omitempty" xml:"SpecIds,omitempty" type:"Repeated"`
	// Specification status.
	//
	// example:
	//
	// Available
	SpecStatus *string `json:"SpecStatus,omitempty" xml:"SpecStatus,omitempty"`
	// Specification type.
	//
	// example:
	//
	// ARM
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
}

func (s DescribeSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpecRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeSpecRequest) GetMatrixSpec() *string {
	return s.MatrixSpec
}

func (s *DescribeSpecRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSpecRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSpecRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *DescribeSpecRequest) GetSpecIds() []*string {
	return s.SpecIds
}

func (s *DescribeSpecRequest) GetSpecStatus() *string {
	return s.SpecStatus
}

func (s *DescribeSpecRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *DescribeSpecRequest) SetBizRegionId(v string) *DescribeSpecRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeSpecRequest) SetMatrixSpec(v string) *DescribeSpecRequest {
	s.MatrixSpec = &v
	return s
}

func (s *DescribeSpecRequest) SetMaxResults(v int32) *DescribeSpecRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSpecRequest) SetNextToken(v string) *DescribeSpecRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSpecRequest) SetSaleMode(v string) *DescribeSpecRequest {
	s.SaleMode = &v
	return s
}

func (s *DescribeSpecRequest) SetSpecIds(v []*string) *DescribeSpecRequest {
	s.SpecIds = v
	return s
}

func (s *DescribeSpecRequest) SetSpecStatus(v string) *DescribeSpecRequest {
	s.SpecStatus = &v
	return s
}

func (s *DescribeSpecRequest) SetSpecType(v string) *DescribeSpecRequest {
	s.SpecType = &v
	return s
}

func (s *DescribeSpecRequest) Validate() error {
	return dara.Validate(s)
}
