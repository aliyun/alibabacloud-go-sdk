// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferProduceOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *TransferProduceOwnerRequest
	GetBizId() *string
	SetBizType(v string) *TransferProduceOwnerRequest
	GetBizType() *string
	SetPersonId(v int32) *TransferProduceOwnerRequest
	GetPersonId() *int32
	SetRemark(v string) *TransferProduceOwnerRequest
	GetRemark() *string
}

type TransferProduceOwnerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// P20210208152920000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esp.companyreg_cloud
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15565
	PersonId *int32  `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s TransferProduceOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferProduceOwnerRequest) GoString() string {
	return s.String()
}

func (s *TransferProduceOwnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *TransferProduceOwnerRequest) GetBizType() *string {
	return s.BizType
}

func (s *TransferProduceOwnerRequest) GetPersonId() *int32 {
	return s.PersonId
}

func (s *TransferProduceOwnerRequest) GetRemark() *string {
	return s.Remark
}

func (s *TransferProduceOwnerRequest) SetBizId(v string) *TransferProduceOwnerRequest {
	s.BizId = &v
	return s
}

func (s *TransferProduceOwnerRequest) SetBizType(v string) *TransferProduceOwnerRequest {
	s.BizType = &v
	return s
}

func (s *TransferProduceOwnerRequest) SetPersonId(v int32) *TransferProduceOwnerRequest {
	s.PersonId = &v
	return s
}

func (s *TransferProduceOwnerRequest) SetRemark(v string) *TransferProduceOwnerRequest {
	s.Remark = &v
	return s
}

func (s *TransferProduceOwnerRequest) Validate() error {
	return dara.Validate(s)
}
