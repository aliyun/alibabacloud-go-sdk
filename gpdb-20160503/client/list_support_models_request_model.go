// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListSupportModelsRequest
	GetRegionId() *string
}

type ListSupportModelsRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSupportModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSupportModelsRequest) GoString() string {
	return s.String()
}

func (s *ListSupportModelsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSupportModelsRequest) SetRegionId(v string) *ListSupportModelsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSupportModelsRequest) Validate() error {
	return dara.Validate(s)
}
