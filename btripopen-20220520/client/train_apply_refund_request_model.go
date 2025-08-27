// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainApplyRefundRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *TrainApplyRefundRequest
	GetOrderId() *string
	SetOutOrderId(v string) *TrainApplyRefundRequest
	GetOutOrderId() *string
	SetOutRefundId(v string) *TrainApplyRefundRequest
	GetOutRefundId() *string
	SetRefundTrainInfos(v []*TrainApplyRefundRequestRefundTrainInfos) *TrainApplyRefundRequest
	GetRefundTrainInfos() []*TrainApplyRefundRequestRefundTrainInfos
}

type TrainApplyRefundRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1002145190081005400
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456778
	OutRefundId *string `json:"out_refund_id,omitempty" xml:"out_refund_id,omitempty"`
	// This parameter is required.
	RefundTrainInfos []*TrainApplyRefundRequestRefundTrainInfos `json:"refund_train_infos,omitempty" xml:"refund_train_infos,omitempty" type:"Repeated"`
}

func (s TrainApplyRefundRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyRefundRequest) GoString() string {
	return s.String()
}

func (s *TrainApplyRefundRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainApplyRefundRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainApplyRefundRequest) GetOutRefundId() *string {
	return s.OutRefundId
}

func (s *TrainApplyRefundRequest) GetRefundTrainInfos() []*TrainApplyRefundRequestRefundTrainInfos {
	return s.RefundTrainInfos
}

func (s *TrainApplyRefundRequest) SetOrderId(v string) *TrainApplyRefundRequest {
	s.OrderId = &v
	return s
}

func (s *TrainApplyRefundRequest) SetOutOrderId(v string) *TrainApplyRefundRequest {
	s.OutOrderId = &v
	return s
}

func (s *TrainApplyRefundRequest) SetOutRefundId(v string) *TrainApplyRefundRequest {
	s.OutRefundId = &v
	return s
}

func (s *TrainApplyRefundRequest) SetRefundTrainInfos(v []*TrainApplyRefundRequestRefundTrainInfos) *TrainApplyRefundRequest {
	s.RefundTrainInfos = v
	return s
}

func (s *TrainApplyRefundRequest) Validate() error {
	return dara.Validate(s)
}

type TrainApplyRefundRequestRefundTrainInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-06 15:19:01
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// This parameter is required.
	RefundPassengerInfos []*TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos `json:"refund_passenger_infos,omitempty" xml:"refund_passenger_infos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// K234
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainApplyRefundRequestRefundTrainInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyRefundRequestRefundTrainInfos) GoString() string {
	return s.String()
}

func (s *TrainApplyRefundRequestRefundTrainInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainApplyRefundRequestRefundTrainInfos) GetRefundPassengerInfos() []*TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos {
	return s.RefundPassengerInfos
}

func (s *TrainApplyRefundRequestRefundTrainInfos) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainApplyRefundRequestRefundTrainInfos) SetDepTime(v string) *TrainApplyRefundRequestRefundTrainInfos {
	s.DepTime = &v
	return s
}

func (s *TrainApplyRefundRequestRefundTrainInfos) SetRefundPassengerInfos(v []*TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos) *TrainApplyRefundRequestRefundTrainInfos {
	s.RefundPassengerInfos = v
	return s
}

func (s *TrainApplyRefundRequestRefundTrainInfos) SetTrainNo(v string) *TrainApplyRefundRequestRefundTrainInfos {
	s.TrainNo = &v
	return s
}

func (s *TrainApplyRefundRequestRefundTrainInfos) Validate() error {
	return dara.Validate(s)
}

type TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// 291487e553c5abde3b611aae283e2526f0d733ab55094aadc0b5ba587222a233c
	PassengerCertNo *string `json:"passenger_cert_no,omitempty" xml:"passenger_cert_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 170d9ac6f8807f9ec603c688f45f78a41
	PassengerCertType *string `json:"passenger_cert_type,omitempty" xml:"passenger_cert_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12334
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 949c9f34f677a0e5d249dfc94f5e62cc7
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
}

func (s TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos) GoString() string {
	return s.String()
}

func (s *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos) GetPassengerCertNo() *string {
	return s.PassengerCertNo
}

func (s *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos) GetPassengerCertType() *string {
	return s.PassengerCertType
}

func (s *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos) SetPassengerCertNo(v string) *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos {
	s.PassengerCertNo = &v
	return s
}

func (s *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos) SetPassengerCertType(v string) *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos {
	s.PassengerCertType = &v
	return s
}

func (s *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos) SetPassengerId(v string) *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos {
	s.PassengerId = &v
	return s
}

func (s *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos) SetPassengerName(v string) *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos {
	s.PassengerName = &v
	return s
}

func (s *TrainApplyRefundRequestRefundTrainInfosRefundPassengerInfos) Validate() error {
	return dara.Validate(s)
}
