// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainFeeCalculateRefundRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDistributeOrderId(v string) *TrainFeeCalculateRefundRequest
	GetDistributeOrderId() *string
	SetOrderId(v string) *TrainFeeCalculateRefundRequest
	GetOrderId() *string
	SetRefundTrainInfos(v []*TrainFeeCalculateRefundRequestRefundTrainInfos) *TrainFeeCalculateRefundRequest
	GetRefundTrainInfos() []*TrainFeeCalculateRefundRequestRefundTrainInfos
}

type TrainFeeCalculateRefundRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	DistributeOrderId *string `json:"distribute_order_id,omitempty" xml:"distribute_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1002145190081005400
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// This parameter is required.
	RefundTrainInfos []*TrainFeeCalculateRefundRequestRefundTrainInfos `json:"refund_train_infos,omitempty" xml:"refund_train_infos,omitempty" type:"Repeated"`
}

func (s TrainFeeCalculateRefundRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateRefundRequest) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateRefundRequest) GetDistributeOrderId() *string {
	return s.DistributeOrderId
}

func (s *TrainFeeCalculateRefundRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainFeeCalculateRefundRequest) GetRefundTrainInfos() []*TrainFeeCalculateRefundRequestRefundTrainInfos {
	return s.RefundTrainInfos
}

func (s *TrainFeeCalculateRefundRequest) SetDistributeOrderId(v string) *TrainFeeCalculateRefundRequest {
	s.DistributeOrderId = &v
	return s
}

func (s *TrainFeeCalculateRefundRequest) SetOrderId(v string) *TrainFeeCalculateRefundRequest {
	s.OrderId = &v
	return s
}

func (s *TrainFeeCalculateRefundRequest) SetRefundTrainInfos(v []*TrainFeeCalculateRefundRequestRefundTrainInfos) *TrainFeeCalculateRefundRequest {
	s.RefundTrainInfos = v
	return s
}

func (s *TrainFeeCalculateRefundRequest) Validate() error {
	return dara.Validate(s)
}

type TrainFeeCalculateRefundRequestRefundTrainInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// BDC
	ArrStationCode *string `json:"arr_station_code,omitempty" xml:"arr_station_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BTC
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-06 15:19:01
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// This parameter is required.
	RefundPassengerInfos []*TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos `json:"refund_passenger_infos,omitempty" xml:"refund_passenger_infos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// K1234
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainFeeCalculateRefundRequestRefundTrainInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateRefundRequestRefundTrainInfos) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfos) GetArrStationCode() *string {
	return s.ArrStationCode
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfos) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfos) GetRefundPassengerInfos() []*TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos {
	return s.RefundPassengerInfos
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfos) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfos) SetArrStationCode(v string) *TrainFeeCalculateRefundRequestRefundTrainInfos {
	s.ArrStationCode = &v
	return s
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfos) SetDepStationCode(v string) *TrainFeeCalculateRefundRequestRefundTrainInfos {
	s.DepStationCode = &v
	return s
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfos) SetDepTime(v string) *TrainFeeCalculateRefundRequestRefundTrainInfos {
	s.DepTime = &v
	return s
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfos) SetRefundPassengerInfos(v []*TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos) *TrainFeeCalculateRefundRequestRefundTrainInfos {
	s.RefundPassengerInfos = v
	return s
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfos) SetTrainNo(v string) *TrainFeeCalculateRefundRequestRefundTrainInfos {
	s.TrainNo = &v
	return s
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfos) Validate() error {
	return dara.Validate(s)
}

type TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos struct {
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
	// 123456
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 949c9f34f677a0e5d249dfc94f5e62cc7
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
}

func (s TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos) GetPassengerCertNo() *string {
	return s.PassengerCertNo
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos) GetPassengerCertType() *string {
	return s.PassengerCertType
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos) SetPassengerCertNo(v string) *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos {
	s.PassengerCertNo = &v
	return s
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos) SetPassengerCertType(v string) *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos {
	s.PassengerCertType = &v
	return s
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos) SetPassengerId(v string) *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos {
	s.PassengerId = &v
	return s
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos) SetPassengerName(v string) *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos {
	s.PassengerName = &v
	return s
}

func (s *TrainFeeCalculateRefundRequestRefundTrainInfosRefundPassengerInfos) Validate() error {
	return dara.Validate(s)
}
