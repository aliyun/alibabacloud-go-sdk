// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListAllMemberRequest
	GetRegionId() *string
}

type ListAllMemberRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAllMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllMemberRequest) GoString() string {
	return s.String()
}

func (s *ListAllMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAllMemberRequest) SetRegionId(v string) *ListAllMemberRequest {
	s.RegionId = &v
	return s
}

func (s *ListAllMemberRequest) Validate() error {
	return dara.Validate(s)
}
