// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainId(v string) *UpdateBizChainRequest
	GetBizChainId() *string
	SetName(v string) *UpdateBizChainRequest
	GetName() *string
	SetRegionId(v string) *UpdateBizChainRequest
	GetRegionId() *string
	SetRemark(v string) *UpdateBizChainRequest
	GetRemark() *string
}

type UpdateBizChainRequest struct {
	// This parameter is required.
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateBizChainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizChainRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizChainRequest) GetBizChainId() *string {
	return s.BizChainId
}

func (s *UpdateBizChainRequest) GetName() *string {
	return s.Name
}

func (s *UpdateBizChainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateBizChainRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateBizChainRequest) SetBizChainId(v string) *UpdateBizChainRequest {
	s.BizChainId = &v
	return s
}

func (s *UpdateBizChainRequest) SetName(v string) *UpdateBizChainRequest {
	s.Name = &v
	return s
}

func (s *UpdateBizChainRequest) SetRegionId(v string) *UpdateBizChainRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBizChainRequest) SetRemark(v string) *UpdateBizChainRequest {
	s.Remark = &v
	return s
}

func (s *UpdateBizChainRequest) Validate() error {
	return dara.Validate(s)
}
