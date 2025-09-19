// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBruteForceSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceOwnerId(v int64) *DescribeBruteForceSummaryRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeBruteForceSummaryRequest
	GetSourceIp() *string
}

type DescribeBruteForceSummaryRequest struct {
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 203.119.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeBruteForceSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBruteForceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceSummaryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBruteForceSummaryRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeBruteForceSummaryRequest) SetResourceOwnerId(v int64) *DescribeBruteForceSummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBruteForceSummaryRequest) SetSourceIp(v string) *DescribeBruteForceSummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBruteForceSummaryRequest) Validate() error {
	return dara.Validate(s)
}
