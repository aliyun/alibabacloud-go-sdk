// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllBizChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListAllBizChainRequest
	GetRegionId() *string
}

type ListAllBizChainRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAllBizChainRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllBizChainRequest) GoString() string {
	return s.String()
}

func (s *ListAllBizChainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAllBizChainRequest) SetRegionId(v string) *ListAllBizChainRequest {
	s.RegionId = &v
	return s
}

func (s *ListAllBizChainRequest) Validate() error {
	return dara.Validate(s)
}
