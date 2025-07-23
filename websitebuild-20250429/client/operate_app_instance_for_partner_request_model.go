// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAppInstanceForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtend(v string) *OperateAppInstanceForPartnerRequest
	GetExtend() *string
	SetOperateEvent(v string) *OperateAppInstanceForPartnerRequest
	GetOperateEvent() *string
}

type OperateAppInstanceForPartnerRequest struct {
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// example:
	//
	// SERVICE_DELIVERY_PROCESS
	OperateEvent *string `json:"OperateEvent,omitempty" xml:"OperateEvent,omitempty"`
}

func (s OperateAppInstanceForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateAppInstanceForPartnerRequest) GoString() string {
	return s.String()
}

func (s *OperateAppInstanceForPartnerRequest) GetExtend() *string {
	return s.Extend
}

func (s *OperateAppInstanceForPartnerRequest) GetOperateEvent() *string {
	return s.OperateEvent
}

func (s *OperateAppInstanceForPartnerRequest) SetExtend(v string) *OperateAppInstanceForPartnerRequest {
	s.Extend = &v
	return s
}

func (s *OperateAppInstanceForPartnerRequest) SetOperateEvent(v string) *OperateAppInstanceForPartnerRequest {
	s.OperateEvent = &v
	return s
}

func (s *OperateAppInstanceForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
