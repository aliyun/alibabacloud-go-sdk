// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUuidVulNumClassifyStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageVul(v bool) *DescribeUuidVulNumClassifyStatisticRequest
	GetImageVul() *bool
	SetUuids(v string) *DescribeUuidVulNumClassifyStatisticRequest
	GetUuids() *string
}

type DescribeUuidVulNumClassifyStatisticRequest struct {
	// example:
	//
	// true
	ImageVul *bool `json:"ImageVul,omitempty" xml:"ImageVul,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 391abd09184cbd3743d7f5ec125d****,
	//
	// e6aeb2a5b6004479398b0bcd1160****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeUuidVulNumClassifyStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUuidVulNumClassifyStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeUuidVulNumClassifyStatisticRequest) GetImageVul() *bool {
	return s.ImageVul
}

func (s *DescribeUuidVulNumClassifyStatisticRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeUuidVulNumClassifyStatisticRequest) SetImageVul(v bool) *DescribeUuidVulNumClassifyStatisticRequest {
	s.ImageVul = &v
	return s
}

func (s *DescribeUuidVulNumClassifyStatisticRequest) SetUuids(v string) *DescribeUuidVulNumClassifyStatisticRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeUuidVulNumClassifyStatisticRequest) Validate() error {
	return dara.Validate(s)
}
