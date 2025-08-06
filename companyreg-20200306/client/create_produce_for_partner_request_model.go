// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProduceForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CreateProduceForPartnerRequest
	GetBizId() *string
	SetBizType(v string) *CreateProduceForPartnerRequest
	GetBizType() *string
	SetExtInfo(v string) *CreateProduceForPartnerRequest
	GetExtInfo() *string
}

type CreateProduceForPartnerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// P20210301102840000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esp.hightech
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s CreateProduceForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProduceForPartnerRequest) GoString() string {
	return s.String()
}

func (s *CreateProduceForPartnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *CreateProduceForPartnerRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateProduceForPartnerRequest) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *CreateProduceForPartnerRequest) SetBizId(v string) *CreateProduceForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *CreateProduceForPartnerRequest) SetBizType(v string) *CreateProduceForPartnerRequest {
	s.BizType = &v
	return s
}

func (s *CreateProduceForPartnerRequest) SetExtInfo(v string) *CreateProduceForPartnerRequest {
	s.ExtInfo = &v
	return s
}

func (s *CreateProduceForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
