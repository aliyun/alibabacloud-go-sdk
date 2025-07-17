// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDifyRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDifyRegionsRequest
	GetClientToken() *string
	SetDataRegion(v string) *DescribeDifyRegionsRequest
	GetDataRegion() *string
}

type DescribeDifyRegionsRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataRegion  *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
}

func (s DescribeDifyRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDifyRegionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDifyRegionsRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *DescribeDifyRegionsRequest) SetClientToken(v string) *DescribeDifyRegionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDifyRegionsRequest) SetDataRegion(v string) *DescribeDifyRegionsRequest {
	s.DataRegion = &v
	return s
}

func (s *DescribeDifyRegionsRequest) Validate() error {
	return dara.Validate(s)
}
