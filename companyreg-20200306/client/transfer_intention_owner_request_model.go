// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferIntentionOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *TransferIntentionOwnerRequest
	GetBizId() *string
	SetBizType(v string) *TransferIntentionOwnerRequest
	GetBizType() *string
	SetEmployeeCode(v string) *TransferIntentionOwnerRequest
	GetEmployeeCode() *string
	SetPersonId(v int32) *TransferIntentionOwnerRequest
	GetPersonId() *int32
	SetRemark(v string) *TransferIntentionOwnerRequest
	GetRemark() *string
}

type TransferIntentionOwnerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// P20210709190452000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esp.wangwen
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	EmployeeCode *string `json:"EmployeeCode,omitempty" xml:"EmployeeCode,omitempty"`
	// example:
	//
	// 67842
	PersonId *int32  `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s TransferIntentionOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferIntentionOwnerRequest) GoString() string {
	return s.String()
}

func (s *TransferIntentionOwnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *TransferIntentionOwnerRequest) GetBizType() *string {
	return s.BizType
}

func (s *TransferIntentionOwnerRequest) GetEmployeeCode() *string {
	return s.EmployeeCode
}

func (s *TransferIntentionOwnerRequest) GetPersonId() *int32 {
	return s.PersonId
}

func (s *TransferIntentionOwnerRequest) GetRemark() *string {
	return s.Remark
}

func (s *TransferIntentionOwnerRequest) SetBizId(v string) *TransferIntentionOwnerRequest {
	s.BizId = &v
	return s
}

func (s *TransferIntentionOwnerRequest) SetBizType(v string) *TransferIntentionOwnerRequest {
	s.BizType = &v
	return s
}

func (s *TransferIntentionOwnerRequest) SetEmployeeCode(v string) *TransferIntentionOwnerRequest {
	s.EmployeeCode = &v
	return s
}

func (s *TransferIntentionOwnerRequest) SetPersonId(v int32) *TransferIntentionOwnerRequest {
	s.PersonId = &v
	return s
}

func (s *TransferIntentionOwnerRequest) SetRemark(v string) *TransferIntentionOwnerRequest {
	s.Remark = &v
	return s
}

func (s *TransferIntentionOwnerRequest) Validate() error {
	return dara.Validate(s)
}
