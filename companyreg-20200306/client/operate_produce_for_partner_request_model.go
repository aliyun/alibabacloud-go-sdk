// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateProduceForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *OperateProduceForPartnerRequest
	GetBizId() *string
	SetBizType(v string) *OperateProduceForPartnerRequest
	GetBizType() *string
	SetExtInfo(v string) *OperateProduceForPartnerRequest
	GetExtInfo() *string
	SetOperateType(v string) *OperateProduceForPartnerRequest
	GetOperateType() *string
}

type OperateProduceForPartnerRequest struct {
	// example:
	//
	// P20210930105636000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// esp.beian_assist
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// {\\"beianServiceNumber\\":\\"9969c666-0935-4c5b-8042-926ff546e39f\\"}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	// example:
	//
	// CERT_MATERIAL_SUBMITTED
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s OperateProduceForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateProduceForPartnerRequest) GoString() string {
	return s.String()
}

func (s *OperateProduceForPartnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *OperateProduceForPartnerRequest) GetBizType() *string {
	return s.BizType
}

func (s *OperateProduceForPartnerRequest) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *OperateProduceForPartnerRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *OperateProduceForPartnerRequest) SetBizId(v string) *OperateProduceForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *OperateProduceForPartnerRequest) SetBizType(v string) *OperateProduceForPartnerRequest {
	s.BizType = &v
	return s
}

func (s *OperateProduceForPartnerRequest) SetExtInfo(v string) *OperateProduceForPartnerRequest {
	s.ExtInfo = &v
	return s
}

func (s *OperateProduceForPartnerRequest) SetOperateType(v string) *OperateProduceForPartnerRequest {
	s.OperateType = &v
	return s
}

func (s *OperateProduceForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
