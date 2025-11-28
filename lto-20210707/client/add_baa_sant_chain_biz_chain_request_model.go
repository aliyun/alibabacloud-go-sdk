// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBaaSAntChainBizChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaaSAntChainChainId(v string) *AddBaaSAntChainBizChainRequest
	GetBaaSAntChainChainId() *string
	SetBaaSAntChainConsortiumId(v string) *AddBaaSAntChainBizChainRequest
	GetBaaSAntChainConsortiumId() *string
	SetCaCert(v string) *AddBaaSAntChainBizChainRequest
	GetCaCert() *string
	SetCaCertPassword(v string) *AddBaaSAntChainBizChainRequest
	GetCaCertPassword() *string
	SetClientCert(v string) *AddBaaSAntChainBizChainRequest
	GetClientCert() *string
	SetClientKey(v string) *AddBaaSAntChainBizChainRequest
	GetClientKey() *string
	SetClientKeyPassword(v string) *AddBaaSAntChainBizChainRequest
	GetClientKeyPassword() *string
	SetContractTemplateIdList(v string) *AddBaaSAntChainBizChainRequest
	GetContractTemplateIdList() *string
	SetName(v string) *AddBaaSAntChainBizChainRequest
	GetName() *string
	SetNodeNameList(v string) *AddBaaSAntChainBizChainRequest
	GetNodeNameList() *string
	SetRegionId(v string) *AddBaaSAntChainBizChainRequest
	GetRegionId() *string
	SetRemark(v string) *AddBaaSAntChainBizChainRequest
	GetRemark() *string
	SetUserKey(v string) *AddBaaSAntChainBizChainRequest
	GetUserKey() *string
	SetUserKeyPassword(v string) *AddBaaSAntChainBizChainRequest
	GetUserKeyPassword() *string
	SetUserName(v string) *AddBaaSAntChainBizChainRequest
	GetUserName() *string
}

type AddBaaSAntChainBizChainRequest struct {
	// This parameter is required.
	BaaSAntChainChainId *string `json:"BaaSAntChainChainId,omitempty" xml:"BaaSAntChainChainId,omitempty"`
	// This parameter is required.
	BaaSAntChainConsortiumId *string `json:"BaaSAntChainConsortiumId,omitempty" xml:"BaaSAntChainConsortiumId,omitempty"`
	// This parameter is required.
	CaCert *string `json:"CaCert,omitempty" xml:"CaCert,omitempty"`
	// This parameter is required.
	CaCertPassword *string `json:"CaCertPassword,omitempty" xml:"CaCertPassword,omitempty"`
	// This parameter is required.
	ClientCert *string `json:"ClientCert,omitempty" xml:"ClientCert,omitempty"`
	// This parameter is required.
	ClientKey *string `json:"ClientKey,omitempty" xml:"ClientKey,omitempty"`
	// This parameter is required.
	ClientKeyPassword      *string `json:"ClientKeyPassword,omitempty" xml:"ClientKeyPassword,omitempty"`
	ContractTemplateIdList *string `json:"ContractTemplateIdList,omitempty" xml:"ContractTemplateIdList,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	NodeNameList *string `json:"NodeNameList,omitempty" xml:"NodeNameList,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	UserKey *string `json:"UserKey,omitempty" xml:"UserKey,omitempty"`
	// This parameter is required.
	UserKeyPassword *string `json:"UserKeyPassword,omitempty" xml:"UserKeyPassword,omitempty"`
	// This parameter is required.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s AddBaaSAntChainBizChainRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBaaSAntChainBizChainRequest) GoString() string {
	return s.String()
}

func (s *AddBaaSAntChainBizChainRequest) GetBaaSAntChainChainId() *string {
	return s.BaaSAntChainChainId
}

func (s *AddBaaSAntChainBizChainRequest) GetBaaSAntChainConsortiumId() *string {
	return s.BaaSAntChainConsortiumId
}

func (s *AddBaaSAntChainBizChainRequest) GetCaCert() *string {
	return s.CaCert
}

func (s *AddBaaSAntChainBizChainRequest) GetCaCertPassword() *string {
	return s.CaCertPassword
}

func (s *AddBaaSAntChainBizChainRequest) GetClientCert() *string {
	return s.ClientCert
}

func (s *AddBaaSAntChainBizChainRequest) GetClientKey() *string {
	return s.ClientKey
}

func (s *AddBaaSAntChainBizChainRequest) GetClientKeyPassword() *string {
	return s.ClientKeyPassword
}

func (s *AddBaaSAntChainBizChainRequest) GetContractTemplateIdList() *string {
	return s.ContractTemplateIdList
}

func (s *AddBaaSAntChainBizChainRequest) GetName() *string {
	return s.Name
}

func (s *AddBaaSAntChainBizChainRequest) GetNodeNameList() *string {
	return s.NodeNameList
}

func (s *AddBaaSAntChainBizChainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddBaaSAntChainBizChainRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddBaaSAntChainBizChainRequest) GetUserKey() *string {
	return s.UserKey
}

func (s *AddBaaSAntChainBizChainRequest) GetUserKeyPassword() *string {
	return s.UserKeyPassword
}

func (s *AddBaaSAntChainBizChainRequest) GetUserName() *string {
	return s.UserName
}

func (s *AddBaaSAntChainBizChainRequest) SetBaaSAntChainChainId(v string) *AddBaaSAntChainBizChainRequest {
	s.BaaSAntChainChainId = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetBaaSAntChainConsortiumId(v string) *AddBaaSAntChainBizChainRequest {
	s.BaaSAntChainConsortiumId = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetCaCert(v string) *AddBaaSAntChainBizChainRequest {
	s.CaCert = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetCaCertPassword(v string) *AddBaaSAntChainBizChainRequest {
	s.CaCertPassword = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetClientCert(v string) *AddBaaSAntChainBizChainRequest {
	s.ClientCert = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetClientKey(v string) *AddBaaSAntChainBizChainRequest {
	s.ClientKey = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetClientKeyPassword(v string) *AddBaaSAntChainBizChainRequest {
	s.ClientKeyPassword = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetContractTemplateIdList(v string) *AddBaaSAntChainBizChainRequest {
	s.ContractTemplateIdList = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetName(v string) *AddBaaSAntChainBizChainRequest {
	s.Name = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetNodeNameList(v string) *AddBaaSAntChainBizChainRequest {
	s.NodeNameList = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetRegionId(v string) *AddBaaSAntChainBizChainRequest {
	s.RegionId = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetRemark(v string) *AddBaaSAntChainBizChainRequest {
	s.Remark = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetUserKey(v string) *AddBaaSAntChainBizChainRequest {
	s.UserKey = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetUserKeyPassword(v string) *AddBaaSAntChainBizChainRequest {
	s.UserKeyPassword = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) SetUserName(v string) *AddBaaSAntChainBizChainRequest {
	s.UserName = &v
	return s
}

func (s *AddBaaSAntChainBizChainRequest) Validate() error {
	return dara.Validate(s)
}
