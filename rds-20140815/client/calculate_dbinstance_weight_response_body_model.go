// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCalculateDBInstanceWeightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *CalculateDBInstanceWeightResponseBodyItems) *CalculateDBInstanceWeightResponseBody
	GetItems() *CalculateDBInstanceWeightResponseBodyItems
	SetRequestId(v string) *CalculateDBInstanceWeightResponseBody
	GetRequestId() *string
}

type CalculateDBInstanceWeightResponseBody struct {
	Items *CalculateDBInstanceWeightResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C816A4BF-A6EC-4722-95F9-2055859CCFD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CalculateDBInstanceWeightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CalculateDBInstanceWeightResponseBody) GoString() string {
	return s.String()
}

func (s *CalculateDBInstanceWeightResponseBody) GetItems() *CalculateDBInstanceWeightResponseBodyItems {
	return s.Items
}

func (s *CalculateDBInstanceWeightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CalculateDBInstanceWeightResponseBody) SetItems(v *CalculateDBInstanceWeightResponseBodyItems) *CalculateDBInstanceWeightResponseBody {
	s.Items = v
	return s
}

func (s *CalculateDBInstanceWeightResponseBody) SetRequestId(v string) *CalculateDBInstanceWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *CalculateDBInstanceWeightResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CalculateDBInstanceWeightResponseBodyItems struct {
	DBInstanceWeight []*CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight `json:"DBInstanceWeight,omitempty" xml:"DBInstanceWeight,omitempty" type:"Repeated"`
}

func (s CalculateDBInstanceWeightResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s CalculateDBInstanceWeightResponseBodyItems) GoString() string {
	return s.String()
}

func (s *CalculateDBInstanceWeightResponseBodyItems) GetDBInstanceWeight() []*CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight {
	return s.DBInstanceWeight
}

func (s *CalculateDBInstanceWeightResponseBodyItems) SetDBInstanceWeight(v []*CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight) *CalculateDBInstanceWeightResponseBodyItems {
	s.DBInstanceWeight = v
	return s
}

func (s *CalculateDBInstanceWeightResponseBodyItems) Validate() error {
	if s.DBInstanceWeight != nil {
		for _, item := range s.DBInstanceWeight {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight struct {
	DBInstanceId                   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceType                 *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	ReadonlyInstanceSQLDelayedTime *string `json:"ReadonlyInstanceSQLDelayedTime,omitempty" xml:"ReadonlyInstanceSQLDelayedTime,omitempty"`
	Weight                         *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight) String() string {
	return dara.Prettify(s)
}

func (s CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight) GoString() string {
	return s.String()
}

func (s *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight) GetReadonlyInstanceSQLDelayedTime() *string {
	return s.ReadonlyInstanceSQLDelayedTime
}

func (s *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight) GetWeight() *string {
	return s.Weight
}

func (s *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight) SetDBInstanceId(v string) *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight {
	s.DBInstanceId = &v
	return s
}

func (s *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight) SetDBInstanceType(v string) *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight {
	s.DBInstanceType = &v
	return s
}

func (s *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight) SetReadonlyInstanceSQLDelayedTime(v string) *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight {
	s.ReadonlyInstanceSQLDelayedTime = &v
	return s
}

func (s *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight) SetWeight(v string) *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight {
	s.Weight = &v
	return s
}

func (s *CalculateDBInstanceWeightResponseBodyItemsDBInstanceWeight) Validate() error {
	return dara.Validate(s)
}
