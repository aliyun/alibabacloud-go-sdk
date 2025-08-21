// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosCreditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDdosRegionId(v string) *DescribeDdosCreditRequest
	GetDdosRegionId() *string
}

type DescribeDdosCreditRequest struct {
	// The ID of the region.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/353250.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
}

func (s DescribeDdosCreditRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosCreditRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosCreditRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DescribeDdosCreditRequest) SetDdosRegionId(v string) *DescribeDdosCreditRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeDdosCreditRequest) Validate() error {
	return dara.Validate(s)
}
