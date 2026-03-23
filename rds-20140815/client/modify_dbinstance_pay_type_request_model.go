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
	SetResourceOwnerId(v int64) *ModifyDBInstancePayTypeRequest
	GetResourceOwnerId() *int64
	SetUsedTime(v int32) *ModifyDBInstancePayTypeRequest
	GetUsedTime() *int32
}

type ModifyDBInstancePayTypeRequest struct {
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// This parameter is required.
	Period          *string `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UsedTime        *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
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

func (s *ModifyDBInstancePayTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *ModifyDBInstancePayTypeRequest) SetResourceOwnerId(v int64) *ModifyDBInstancePayTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstancePayTypeRequest) SetUsedTime(v int32) *ModifyDBInstancePayTypeRequest {
	s.UsedTime = &v
	return s
}

func (s *ModifyDBInstancePayTypeRequest) Validate() error {
	return dara.Validate(s)
}
