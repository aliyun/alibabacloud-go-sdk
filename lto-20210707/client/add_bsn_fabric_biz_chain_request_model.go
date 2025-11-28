// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBsnFabricBizChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *AddBsnFabricBizChainRequest
	GetAppCode() *string
	SetName(v string) *AddBsnFabricBizChainRequest
	GetName() *string
	SetNodeList(v string) *AddBsnFabricBizChainRequest
	GetNodeList() *string
	SetRegionId(v string) *AddBsnFabricBizChainRequest
	GetRegionId() *string
	SetRemark(v string) *AddBsnFabricBizChainRequest
	GetRemark() *string
	SetUserCode(v string) *AddBsnFabricBizChainRequest
	GetUserCode() *string
}

type AddBsnFabricBizChainRequest struct {
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	NodeList *string `json:"NodeList,omitempty" xml:"NodeList,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	UserCode *string `json:"UserCode,omitempty" xml:"UserCode,omitempty"`
}

func (s AddBsnFabricBizChainRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBsnFabricBizChainRequest) GoString() string {
	return s.String()
}

func (s *AddBsnFabricBizChainRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *AddBsnFabricBizChainRequest) GetName() *string {
	return s.Name
}

func (s *AddBsnFabricBizChainRequest) GetNodeList() *string {
	return s.NodeList
}

func (s *AddBsnFabricBizChainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddBsnFabricBizChainRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddBsnFabricBizChainRequest) GetUserCode() *string {
	return s.UserCode
}

func (s *AddBsnFabricBizChainRequest) SetAppCode(v string) *AddBsnFabricBizChainRequest {
	s.AppCode = &v
	return s
}

func (s *AddBsnFabricBizChainRequest) SetName(v string) *AddBsnFabricBizChainRequest {
	s.Name = &v
	return s
}

func (s *AddBsnFabricBizChainRequest) SetNodeList(v string) *AddBsnFabricBizChainRequest {
	s.NodeList = &v
	return s
}

func (s *AddBsnFabricBizChainRequest) SetRegionId(v string) *AddBsnFabricBizChainRequest {
	s.RegionId = &v
	return s
}

func (s *AddBsnFabricBizChainRequest) SetRemark(v string) *AddBsnFabricBizChainRequest {
	s.Remark = &v
	return s
}

func (s *AddBsnFabricBizChainRequest) SetUserCode(v string) *AddBsnFabricBizChainRequest {
	s.UserCode = &v
	return s
}

func (s *AddBsnFabricBizChainRequest) Validate() error {
	return dara.Validate(s)
}
