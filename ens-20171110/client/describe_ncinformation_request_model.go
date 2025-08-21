// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNCInformationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeNCInformationRequest
	GetEnsRegionId() *string
	SetResourceId(v string) *DescribeNCInformationRequest
	GetResourceId() *string
}

type DescribeNCInformationRequest struct {
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	ResourceId  *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeNCInformationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNCInformationRequest) GoString() string {
	return s.String()
}

func (s *DescribeNCInformationRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeNCInformationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNCInformationRequest) SetEnsRegionId(v string) *DescribeNCInformationRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNCInformationRequest) SetResourceId(v string) *DescribeNCInformationRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeNCInformationRequest) Validate() error {
	return dara.Validate(s)
}
