// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableClassesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceClasses(v []*DescribeAvailableClassesResponseBodyDBInstanceClasses) *DescribeAvailableClassesResponseBody
	GetDBInstanceClasses() []*DescribeAvailableClassesResponseBodyDBInstanceClasses
	SetRequestId(v string) *DescribeAvailableClassesResponseBody
	GetRequestId() *string
}

type DescribeAvailableClassesResponseBody struct {
	// An array that consists of the instance types available for the instance.
	DBInstanceClasses []*DescribeAvailableClassesResponseBodyDBInstanceClasses `json:"DBInstanceClasses,omitempty" xml:"DBInstanceClasses,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 7E4448A6-9FE6-4474-A0C1-AA7CFC772CAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableClassesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableClassesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableClassesResponseBody) GetDBInstanceClasses() []*DescribeAvailableClassesResponseBodyDBInstanceClasses {
	return s.DBInstanceClasses
}

func (s *DescribeAvailableClassesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableClassesResponseBody) SetDBInstanceClasses(v []*DescribeAvailableClassesResponseBodyDBInstanceClasses) *DescribeAvailableClassesResponseBody {
	s.DBInstanceClasses = v
	return s
}

func (s *DescribeAvailableClassesResponseBody) SetRequestId(v string) *DescribeAvailableClassesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableClassesResponseBody) Validate() error {
	if s.DBInstanceClasses != nil {
		for _, item := range s.DBInstanceClasses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableClassesResponseBodyDBInstanceClasses struct {
	// The instance type of the instance.
	//
	// example:
	//
	// rds.mysql.c1.large
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The storage capacity range that is supported for the instance.
	DBInstanceStorageRange *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange `json:"DBInstanceStorageRange,omitempty" xml:"DBInstanceStorageRange,omitempty" type:"Struct"`
}

func (s DescribeAvailableClassesResponseBodyDBInstanceClasses) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableClassesResponseBodyDBInstanceClasses) GoString() string {
	return s.String()
}

func (s *DescribeAvailableClassesResponseBodyDBInstanceClasses) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeAvailableClassesResponseBodyDBInstanceClasses) GetDBInstanceStorageRange() *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange {
	return s.DBInstanceStorageRange
}

func (s *DescribeAvailableClassesResponseBodyDBInstanceClasses) SetDBInstanceClass(v string) *DescribeAvailableClassesResponseBodyDBInstanceClasses {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeAvailableClassesResponseBodyDBInstanceClasses) SetDBInstanceStorageRange(v *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange) *DescribeAvailableClassesResponseBodyDBInstanceClasses {
	s.DBInstanceStorageRange = v
	return s
}

func (s *DescribeAvailableClassesResponseBodyDBInstanceClasses) Validate() error {
	if s.DBInstanceStorageRange != nil {
		if err := s.DBInstanceStorageRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange struct {
	// The maximum storage capacity that is supported for the instance. Unit: GB.
	//
	// example:
	//
	// 2000
	MaxValue *int32 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum storage capacity that is supported for the instance. Unit: GB.
	//
	// example:
	//
	// 5
	MinValue *int32 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The minimum step size at which you can adjust the storage capacity of the instance. The minimum step size is 5 GB.
	//
	// example:
	//
	// 5
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange) GoString() string {
	return s.String()
}

func (s *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange) GetMaxValue() *int32 {
	return s.MaxValue
}

func (s *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange) GetMinValue() *int32 {
	return s.MinValue
}

func (s *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange) GetStep() *int32 {
	return s.Step
}

func (s *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange) SetMaxValue(v int32) *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange {
	s.MaxValue = &v
	return s
}

func (s *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange) SetMinValue(v int32) *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange {
	s.MinValue = &v
	return s
}

func (s *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange) SetStep(v int32) *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange {
	s.Step = &v
	return s
}

func (s *DescribeAvailableClassesResponseBodyDBInstanceClassesDBInstanceStorageRange) Validate() error {
	return dara.Validate(s)
}
