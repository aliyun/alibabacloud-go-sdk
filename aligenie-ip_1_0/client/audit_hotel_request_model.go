// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditHotelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditHotelReq(v *AuditHotelRequestAuditHotelReq) *AuditHotelRequest
	GetAuditHotelReq() *AuditHotelRequestAuditHotelReq
}

type AuditHotelRequest struct {
	// This parameter is required.
	AuditHotelReq *AuditHotelRequestAuditHotelReq `json:"AuditHotelReq,omitempty" xml:"AuditHotelReq,omitempty" type:"Struct"`
}

func (s AuditHotelRequest) String() string {
	return dara.Prettify(s)
}

func (s AuditHotelRequest) GoString() string {
	return s.String()
}

func (s *AuditHotelRequest) GetAuditHotelReq() *AuditHotelRequestAuditHotelReq {
	return s.AuditHotelReq
}

func (s *AuditHotelRequest) SetAuditHotelReq(v *AuditHotelRequestAuditHotelReq) *AuditHotelRequest {
	s.AuditHotelReq = v
	return s
}

func (s *AuditHotelRequest) Validate() error {
	if s.AuditHotelReq != nil {
		if err := s.AuditHotelReq.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AuditHotelRequestAuditHotelReq struct {
	// example:
	//
	// 同意
	AuditOpinion *string `json:"AuditOpinion,omitempty" xml:"AuditOpinion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AuditHotelRequestAuditHotelReq) String() string {
	return dara.Prettify(s)
}

func (s AuditHotelRequestAuditHotelReq) GoString() string {
	return s.String()
}

func (s *AuditHotelRequestAuditHotelReq) GetAuditOpinion() *string {
	return s.AuditOpinion
}

func (s *AuditHotelRequestAuditHotelReq) GetHotelId() *string {
	return s.HotelId
}

func (s *AuditHotelRequestAuditHotelReq) GetStatus() *int32 {
	return s.Status
}

func (s *AuditHotelRequestAuditHotelReq) SetAuditOpinion(v string) *AuditHotelRequestAuditHotelReq {
	s.AuditOpinion = &v
	return s
}

func (s *AuditHotelRequestAuditHotelReq) SetHotelId(v string) *AuditHotelRequestAuditHotelReq {
	s.HotelId = &v
	return s
}

func (s *AuditHotelRequestAuditHotelReq) SetStatus(v int32) *AuditHotelRequestAuditHotelReq {
	s.Status = &v
	return s
}

func (s *AuditHotelRequestAuditHotelReq) Validate() error {
	return dara.Validate(s)
}
