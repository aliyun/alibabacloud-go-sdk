// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainFeeCalculateChangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeTrainDetails(v []*TrainFeeCalculateChangeRequestChangeTrainDetails) *TrainFeeCalculateChangeRequest
	GetChangeTrainDetails() []*TrainFeeCalculateChangeRequestChangeTrainDetails
	SetDistributeOrderId(v string) *TrainFeeCalculateChangeRequest
	GetDistributeOrderId() *string
	SetOrderId(v string) *TrainFeeCalculateChangeRequest
	GetOrderId() *string
}

type TrainFeeCalculateChangeRequest struct {
	// This parameter is required.
	ChangeTrainDetails []*TrainFeeCalculateChangeRequestChangeTrainDetails `json:"change_train_details,omitempty" xml:"change_train_details,omitempty" type:"Repeated"`
	// This parameter is required.
	DistributeOrderId *string `json:"distribute_order_id,omitempty" xml:"distribute_order_id,omitempty"`
	// This parameter is required.
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s TrainFeeCalculateChangeRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateChangeRequest) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateChangeRequest) GetChangeTrainDetails() []*TrainFeeCalculateChangeRequestChangeTrainDetails {
	return s.ChangeTrainDetails
}

func (s *TrainFeeCalculateChangeRequest) GetDistributeOrderId() *string {
	return s.DistributeOrderId
}

func (s *TrainFeeCalculateChangeRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainFeeCalculateChangeRequest) SetChangeTrainDetails(v []*TrainFeeCalculateChangeRequestChangeTrainDetails) *TrainFeeCalculateChangeRequest {
	s.ChangeTrainDetails = v
	return s
}

func (s *TrainFeeCalculateChangeRequest) SetDistributeOrderId(v string) *TrainFeeCalculateChangeRequest {
	s.DistributeOrderId = &v
	return s
}

func (s *TrainFeeCalculateChangeRequest) SetOrderId(v string) *TrainFeeCalculateChangeRequest {
	s.OrderId = &v
	return s
}

func (s *TrainFeeCalculateChangeRequest) Validate() error {
	if s.ChangeTrainDetails != nil {
		for _, item := range s.ChangeTrainDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TrainFeeCalculateChangeRequestChangeTrainDetails struct {
	// This parameter is required.
	ArrStationCode *string `json:"arr_station_code,omitempty" xml:"arr_station_code,omitempty"`
	// This parameter is required.
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	// This parameter is required.
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// This parameter is required.
	OriginalDepTime *string `json:"original_dep_time,omitempty" xml:"original_dep_time,omitempty"`
	// This parameter is required.
	OriginalTrainNo *string `json:"original_train_no,omitempty" xml:"original_train_no,omitempty"`
	// This parameter is required.
	PassengerInfo *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo `json:"passenger_info,omitempty" xml:"passenger_info,omitempty" type:"Struct"`
	// This parameter is required.
	SeatType *string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// This parameter is required.
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainFeeCalculateChangeRequestChangeTrainDetails) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateChangeRequestChangeTrainDetails) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) GetArrStationCode() *string {
	return s.ArrStationCode
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) GetOriginalDepTime() *string {
	return s.OriginalDepTime
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) GetOriginalTrainNo() *string {
	return s.OriginalTrainNo
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) GetPassengerInfo() *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo {
	return s.PassengerInfo
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) SetArrStationCode(v string) *TrainFeeCalculateChangeRequestChangeTrainDetails {
	s.ArrStationCode = &v
	return s
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) SetDepStationCode(v string) *TrainFeeCalculateChangeRequestChangeTrainDetails {
	s.DepStationCode = &v
	return s
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) SetDepTime(v string) *TrainFeeCalculateChangeRequestChangeTrainDetails {
	s.DepTime = &v
	return s
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) SetOriginalDepTime(v string) *TrainFeeCalculateChangeRequestChangeTrainDetails {
	s.OriginalDepTime = &v
	return s
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) SetOriginalTrainNo(v string) *TrainFeeCalculateChangeRequestChangeTrainDetails {
	s.OriginalTrainNo = &v
	return s
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) SetPassengerInfo(v *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo) *TrainFeeCalculateChangeRequestChangeTrainDetails {
	s.PassengerInfo = v
	return s
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) SetSeatType(v string) *TrainFeeCalculateChangeRequestChangeTrainDetails {
	s.SeatType = &v
	return s
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) SetTrainNo(v string) *TrainFeeCalculateChangeRequestChangeTrainDetails {
	s.TrainNo = &v
	return s
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetails) Validate() error {
	if s.PassengerInfo != nil {
		if err := s.PassengerInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo struct {
	// This parameter is required.
	PassengerCertNo *string `json:"passenger_cert_no,omitempty" xml:"passenger_cert_no,omitempty"`
	// This parameter is required.
	PassengerCertType *string `json:"passenger_cert_type,omitempty" xml:"passenger_cert_type,omitempty"`
	// This parameter is required.
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
}

func (s TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo) GetPassengerCertNo() *string {
	return s.PassengerCertNo
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo) GetPassengerCertType() *string {
	return s.PassengerCertType
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo) SetPassengerCertNo(v string) *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo {
	s.PassengerCertNo = &v
	return s
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo) SetPassengerCertType(v string) *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo {
	s.PassengerCertType = &v
	return s
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo) SetPassengerName(v string) *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo {
	s.PassengerName = &v
	return s
}

func (s *TrainFeeCalculateChangeRequestChangeTrainDetailsPassengerInfo) Validate() error {
	return dara.Validate(s)
}
