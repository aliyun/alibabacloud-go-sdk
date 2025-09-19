// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttestorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeAttestorsRequest
	GetCurrentPage() *int32
	SetName(v string) *DescribeAttestorsRequest
	GetName() *string
	SetPageSize(v int32) *DescribeAttestorsRequest
	GetPageSize() *int32
	SetResourceOwnerId(v int64) *DescribeAttestorsRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeAttestorsRequest
	GetSourceIp() *string
}

type DescribeAttestorsRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the witness.
	//
	// example:
	//
	// attestor-auto-ad5316
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize        *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 222.35.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAttestorsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttestorsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAttestorsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAttestorsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeAttestorsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAttestorsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAttestorsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAttestorsRequest) SetCurrentPage(v int32) *DescribeAttestorsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAttestorsRequest) SetName(v string) *DescribeAttestorsRequest {
	s.Name = &v
	return s
}

func (s *DescribeAttestorsRequest) SetPageSize(v int32) *DescribeAttestorsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAttestorsRequest) SetResourceOwnerId(v int64) *DescribeAttestorsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAttestorsRequest) SetSourceIp(v string) *DescribeAttestorsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAttestorsRequest) Validate() error {
	return dara.Validate(s)
}
