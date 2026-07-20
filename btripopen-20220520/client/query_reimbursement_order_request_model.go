// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReimbursementOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReimbOrderNo(v string) *QueryReimbursementOrderRequest
	GetReimbOrderNo() *string
	SetSubCorpId(v string) *QueryReimbursementOrderRequest
	GetSubCorpId() *string
}

type QueryReimbursementOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// RT204396
	ReimbOrderNo *string `json:"reimb_order_no,omitempty" xml:"reimb_order_no,omitempty"`
	// example:
	//
	// ding123
	SubCorpId *string `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
}

func (s QueryReimbursementOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryReimbursementOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryReimbursementOrderRequest) GetReimbOrderNo() *string {
	return s.ReimbOrderNo
}

func (s *QueryReimbursementOrderRequest) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *QueryReimbursementOrderRequest) SetReimbOrderNo(v string) *QueryReimbursementOrderRequest {
	s.ReimbOrderNo = &v
	return s
}

func (s *QueryReimbursementOrderRequest) SetSubCorpId(v string) *QueryReimbursementOrderRequest {
	s.SubCorpId = &v
	return s
}

func (s *QueryReimbursementOrderRequest) Validate() error {
	return dara.Validate(s)
}
