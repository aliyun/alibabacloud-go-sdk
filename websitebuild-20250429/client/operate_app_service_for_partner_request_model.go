// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAppServiceForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *OperateAppServiceForPartnerRequest
	GetBizId() *string
	SetExtend(v string) *OperateAppServiceForPartnerRequest
	GetExtend() *string
	SetOperateEvent(v string) *OperateAppServiceForPartnerRequest
	GetOperateEvent() *string
	SetServiceType(v string) *OperateAppServiceForPartnerRequest
	GetServiceType() *string
}

type OperateAppServiceForPartnerRequest struct {
	// Business ID.
	//
	// example:
	//
	// WS00001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Additional extension information in JSON structure, facilitating future parameter extensions.
	//
	// example:
	//
	// {\\"user_device_id\\":\\"6bef45cb0c76de284d24de074c088b73\\"}\\n
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// Operation event:
	//
	// SERVICE_FINISH: Service completed
	//
	// example:
	//
	// SERVICE_FINISH
	OperateEvent *string `json:"OperateEvent,omitempty" xml:"OperateEvent,omitempty"`
	// Service Type.
	//
	// Valid values:
	//
	// - private: Deployed under the user\\"s account.
	//
	// - managed: Hosted under the service provider\\"s account.
	//
	// - operation: Alibaba Cloud Managed Services.
	//
	// example:
	//
	// WEBSITE_DESIGN
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s OperateAppServiceForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateAppServiceForPartnerRequest) GoString() string {
	return s.String()
}

func (s *OperateAppServiceForPartnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *OperateAppServiceForPartnerRequest) GetExtend() *string {
	return s.Extend
}

func (s *OperateAppServiceForPartnerRequest) GetOperateEvent() *string {
	return s.OperateEvent
}

func (s *OperateAppServiceForPartnerRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *OperateAppServiceForPartnerRequest) SetBizId(v string) *OperateAppServiceForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *OperateAppServiceForPartnerRequest) SetExtend(v string) *OperateAppServiceForPartnerRequest {
	s.Extend = &v
	return s
}

func (s *OperateAppServiceForPartnerRequest) SetOperateEvent(v string) *OperateAppServiceForPartnerRequest {
	s.OperateEvent = &v
	return s
}

func (s *OperateAppServiceForPartnerRequest) SetServiceType(v string) *OperateAppServiceForPartnerRequest {
	s.ServiceType = &v
	return s
}

func (s *OperateAppServiceForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
