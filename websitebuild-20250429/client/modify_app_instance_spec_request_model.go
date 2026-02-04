// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *ModifyAppInstanceSpecRequest
	GetApplicationType() *string
	SetBizId(v string) *ModifyAppInstanceSpecRequest
	GetBizId() *string
	SetClientToken(v string) *ModifyAppInstanceSpecRequest
	GetClientToken() *string
	SetDeployArea(v string) *ModifyAppInstanceSpecRequest
	GetDeployArea() *string
	SetExtend(v string) *ModifyAppInstanceSpecRequest
	GetExtend() *string
	SetPaymentType(v string) *ModifyAppInstanceSpecRequest
	GetPaymentType() *string
	SetSiteVersion(v string) *ModifyAppInstanceSpecRequest
	GetSiteVersion() *string
}

type ModifyAppInstanceSpecRequest struct {
	// Application type
	//
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// Business ID
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Ensures idempotence of requests. Generate a unique value from your client to ensure it is unique across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters.
	//
	// example:
	//
	// 2108341e17661121129745384e79f9
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Deployment area
	//
	// example:
	//
	// HongKong
	DeployArea *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	// Extended information
	//
	// example:
	//
	// {\\"deliveryNodeName\\":\\"网站验收\\",\\"deliveryNodeStatus\\":\\"Reject\\",\\"deliveryOperatorRole\\":\\"Customer\\"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// Payment type
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// Site version
	//
	// example:
	//
	// 0
	SiteVersion *string `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ModifyAppInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceSpecRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *ModifyAppInstanceSpecRequest) GetBizId() *string {
	return s.BizId
}

func (s *ModifyAppInstanceSpecRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppInstanceSpecRequest) GetDeployArea() *string {
	return s.DeployArea
}

func (s *ModifyAppInstanceSpecRequest) GetExtend() *string {
	return s.Extend
}

func (s *ModifyAppInstanceSpecRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ModifyAppInstanceSpecRequest) GetSiteVersion() *string {
	return s.SiteVersion
}

func (s *ModifyAppInstanceSpecRequest) SetApplicationType(v string) *ModifyAppInstanceSpecRequest {
	s.ApplicationType = &v
	return s
}

func (s *ModifyAppInstanceSpecRequest) SetBizId(v string) *ModifyAppInstanceSpecRequest {
	s.BizId = &v
	return s
}

func (s *ModifyAppInstanceSpecRequest) SetClientToken(v string) *ModifyAppInstanceSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppInstanceSpecRequest) SetDeployArea(v string) *ModifyAppInstanceSpecRequest {
	s.DeployArea = &v
	return s
}

func (s *ModifyAppInstanceSpecRequest) SetExtend(v string) *ModifyAppInstanceSpecRequest {
	s.Extend = &v
	return s
}

func (s *ModifyAppInstanceSpecRequest) SetPaymentType(v string) *ModifyAppInstanceSpecRequest {
	s.PaymentType = &v
	return s
}

func (s *ModifyAppInstanceSpecRequest) SetSiteVersion(v string) *ModifyAppInstanceSpecRequest {
	s.SiteVersion = &v
	return s
}

func (s *ModifyAppInstanceSpecRequest) Validate() error {
	return dara.Validate(s)
}
