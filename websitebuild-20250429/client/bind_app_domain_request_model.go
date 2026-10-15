// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAppDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *BindAppDomainRequest
	GetBizId() *string
	SetDeployType(v string) *BindAppDomainRequest
	GetDeployType() *string
	SetDomainName(v string) *BindAppDomainRequest
	GetDomainName() *string
	SetDomainType(v string) *BindAppDomainRequest
	GetDomainType() *string
	SetExtend(v string) *BindAppDomainRequest
	GetExtend() *string
	SetOperateType(v string) *BindAppDomainRequest
	GetOperateType() *string
}

type BindAppDomainRequest struct {
	// The business ID.
	//
	// example:
	//
	// WD20250820143531000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The domain hosting method. Valid values:
	//
	// - CUSTOMER_ESA: Uses the customer\\"s own Edge Security Acceleration (ESA). Only supported when DomainType is set to CUSTOM.
	//
	// If this parameter is not specified, the original hosting process is used.
	//
	// example:
	//
	// CDN
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The domain name.
	//
	// example:
	//
	// xiaoniu.link
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The domain management type. Valid values:
	//
	// - CUSTOM
	//
	// - PLATFORM_PREFIX
	//
	// Default value: CUSTOM.
	//
	// example:
	//
	// CUSTOM
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The extended information (OverwriteExistingRecord).
	//
	// example:
	//
	// {\\"deliveryNodeName\\":\\"Website acceptance\\",\\"deliveryNodeStatus\\":\\"Reject\\",\\"deliveryOperatorRole\\":\\"Customer\\"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The operation type for domain binding.
	//
	// example:
	//
	// vul_fix
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s BindAppDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAppDomainRequest) GoString() string {
	return s.String()
}

func (s *BindAppDomainRequest) GetBizId() *string {
	return s.BizId
}

func (s *BindAppDomainRequest) GetDeployType() *string {
	return s.DeployType
}

func (s *BindAppDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *BindAppDomainRequest) GetDomainType() *string {
	return s.DomainType
}

func (s *BindAppDomainRequest) GetExtend() *string {
	return s.Extend
}

func (s *BindAppDomainRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *BindAppDomainRequest) SetBizId(v string) *BindAppDomainRequest {
	s.BizId = &v
	return s
}

func (s *BindAppDomainRequest) SetDeployType(v string) *BindAppDomainRequest {
	s.DeployType = &v
	return s
}

func (s *BindAppDomainRequest) SetDomainName(v string) *BindAppDomainRequest {
	s.DomainName = &v
	return s
}

func (s *BindAppDomainRequest) SetDomainType(v string) *BindAppDomainRequest {
	s.DomainType = &v
	return s
}

func (s *BindAppDomainRequest) SetExtend(v string) *BindAppDomainRequest {
	s.Extend = &v
	return s
}

func (s *BindAppDomainRequest) SetOperateType(v string) *BindAppDomainRequest {
	s.OperateType = &v
	return s
}

func (s *BindAppDomainRequest) Validate() error {
	return dara.Validate(s)
}
