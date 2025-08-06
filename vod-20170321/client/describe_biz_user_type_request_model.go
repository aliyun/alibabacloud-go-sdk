// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBizUserTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *DescribeBizUserTypeRequest
	GetRegion() *string
}

type DescribeBizUserTypeRequest struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeBizUserTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBizUserTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeBizUserTypeRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeBizUserTypeRequest) SetRegion(v string) *DescribeBizUserTypeRequest {
	s.Region = &v
	return s
}

func (s *DescribeBizUserTypeRequest) Validate() error {
	return dara.Validate(s)
}
