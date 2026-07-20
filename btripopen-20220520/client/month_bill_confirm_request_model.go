// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillConfirmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMailBillDate(v int32) *MonthBillConfirmRequest
	GetMailBillDate() *int32
	SetUserId(v string) *MonthBillConfirmRequest
	GetUserId() *string
}

type MonthBillConfirmRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20221001
	MailBillDate *int32 `json:"mail_bill_date,omitempty" xml:"mail_bill_date,omitempty"`
	// example:
	//
	// 12345
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s MonthBillConfirmRequest) String() string {
	return dara.Prettify(s)
}

func (s MonthBillConfirmRequest) GoString() string {
	return s.String()
}

func (s *MonthBillConfirmRequest) GetMailBillDate() *int32 {
	return s.MailBillDate
}

func (s *MonthBillConfirmRequest) GetUserId() *string {
	return s.UserId
}

func (s *MonthBillConfirmRequest) SetMailBillDate(v int32) *MonthBillConfirmRequest {
	s.MailBillDate = &v
	return s
}

func (s *MonthBillConfirmRequest) SetUserId(v string) *MonthBillConfirmRequest {
	s.UserId = &v
	return s
}

func (s *MonthBillConfirmRequest) Validate() error {
	return dara.Validate(s)
}
