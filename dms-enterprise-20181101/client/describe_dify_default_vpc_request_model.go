// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDifyDefaultVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDifyDefaultVpcRequest
	GetClientToken() *string
	SetDataRegion(v string) *DescribeDifyDefaultVpcRequest
	GetDataRegion() *string
	SetWorkspaceId(v string) *DescribeDifyDefaultVpcRequest
	GetWorkspaceId() *string
}

type DescribeDifyDefaultVpcRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DataRegion  *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeDifyDefaultVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyDefaultVpcRequest) GoString() string {
	return s.String()
}

func (s *DescribeDifyDefaultVpcRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDifyDefaultVpcRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *DescribeDifyDefaultVpcRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeDifyDefaultVpcRequest) SetClientToken(v string) *DescribeDifyDefaultVpcRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDifyDefaultVpcRequest) SetDataRegion(v string) *DescribeDifyDefaultVpcRequest {
	s.DataRegion = &v
	return s
}

func (s *DescribeDifyDefaultVpcRequest) SetWorkspaceId(v string) *DescribeDifyDefaultVpcRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeDifyDefaultVpcRequest) Validate() error {
	return dara.Validate(s)
}
