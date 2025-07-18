// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstancePayTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstancePayTypeRequest
	GetDBInstanceId() *string
	SetPayType(v string) *ModifyDBInstancePayTypeRequest
	GetPayType() *string
	SetPeriod(v string) *ModifyDBInstancePayTypeRequest
	GetPeriod() *string
	SetUsedTime(v int32) *ModifyDBInstancePayTypeRequest
	GetUsedTime() *int32
}

type ModifyDBInstancePayTypeRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- Postpaid: pay-as-you-go.
	//
	// 	- Prepaid: subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// This parameter must be specified only when PayType is set to Prepaid.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The subscription duration.
	//
	// 	- Valid values when Period is set to Month: 1 to 9.
	//
	// 	- Valid values when Period is set to Year: 1 to 3.
	//
	// This parameter must be specified only when PayType is set to Prepaid.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s ModifyDBInstancePayTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstancePayTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstancePayTypeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstancePayTypeRequest) GetPayType() *string {
	return s.PayType
}

func (s *ModifyDBInstancePayTypeRequest) GetPeriod() *string {
	return s.Period
}

func (s *ModifyDBInstancePayTypeRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *ModifyDBInstancePayTypeRequest) SetDBInstanceId(v string) *ModifyDBInstancePayTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstancePayTypeRequest) SetPayType(v string) *ModifyDBInstancePayTypeRequest {
	s.PayType = &v
	return s
}

func (s *ModifyDBInstancePayTypeRequest) SetPeriod(v string) *ModifyDBInstancePayTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyDBInstancePayTypeRequest) SetUsedTime(v int32) *ModifyDBInstancePayTypeRequest {
	s.UsedTime = &v
	return s
}

func (s *ModifyDBInstancePayTypeRequest) Validate() error {
	return dara.Validate(s)
}
