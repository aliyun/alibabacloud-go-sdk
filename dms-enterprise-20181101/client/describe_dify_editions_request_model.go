// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDifyEditionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDifyEditionsRequest
	GetClientToken() *string
	SetDataRegion(v string) *DescribeDifyEditionsRequest
	GetDataRegion() *string
}

type DescribeDifyEditionsRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataRegion  *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
}

func (s DescribeDifyEditionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyEditionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDifyEditionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDifyEditionsRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *DescribeDifyEditionsRequest) SetClientToken(v string) *DescribeDifyEditionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDifyEditionsRequest) SetDataRegion(v string) *DescribeDifyEditionsRequest {
	s.DataRegion = &v
	return s
}

func (s *DescribeDifyEditionsRequest) Validate() error {
	return dara.Validate(s)
}
