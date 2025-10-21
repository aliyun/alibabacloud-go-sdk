// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceConfigChangeLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstanceConfigChangeLogResponseBodyData) *DescribeDBInstanceConfigChangeLogResponseBody
	GetData() *DescribeDBInstanceConfigChangeLogResponseBodyData
	SetRequestId(v string) *DescribeDBInstanceConfigChangeLogResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceConfigChangeLogResponseBody struct {
	Data *DescribeDBInstanceConfigChangeLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 780DE414-*********-88BE-A2E21B862B57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceConfigChangeLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigChangeLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigChangeLogResponseBody) GetData() *DescribeDBInstanceConfigChangeLogResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstanceConfigChangeLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceConfigChangeLogResponseBody) SetData(v *DescribeDBInstanceConfigChangeLogResponseBodyData) *DescribeDBInstanceConfigChangeLogResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponseBody) SetRequestId(v string) *DescribeDBInstanceConfigChangeLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceConfigChangeLogResponseBodyData struct {
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId    *string                                                             `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ParamChangeLogs []*DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs `json:"ParamChangeLogs,omitempty" xml:"ParamChangeLogs,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceConfigChangeLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigChangeLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyData) GetParamChangeLogs() []*DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs {
	return s.ParamChangeLogs
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyData) SetDBInstanceId(v string) *DescribeDBInstanceConfigChangeLogResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyData) SetParamChangeLogs(v []*DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) *DescribeDBInstanceConfigChangeLogResponseBodyData {
	s.ParamChangeLogs = v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyData) Validate() error {
	if s.ParamChangeLogs != nil {
		for _, item := range s.ParamChangeLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs struct {
	Applied *bool `json:"Applied,omitempty" xml:"Applied,omitempty"`
	// example:
	//
	// 2025-06-25 13:46:06
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2025-06-25 13:46:06
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1
	ID *int64 `json:"ID,omitempty" xml:"ID,omitempty"`
	// example:
	//
	// max_concurrent_queries
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 100
	NewValue *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	// example:
	//
	// 50
	OldValue *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
}

func (s DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) GetApplied() *bool {
	return s.Applied
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) GetID() *int64 {
	return s.ID
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) GetName() *string {
	return s.Name
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) GetNewValue() *string {
	return s.NewValue
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) GetOldValue() *string {
	return s.OldValue
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) SetApplied(v bool) *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs {
	s.Applied = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) SetGmtCreated(v string) *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) SetGmtModified(v string) *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs {
	s.GmtModified = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) SetID(v int64) *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs {
	s.ID = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) SetName(v string) *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs {
	s.Name = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) SetNewValue(v string) *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs {
	s.NewValue = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) SetOldValue(v string) *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs {
	s.OldValue = &v
	return s
}

func (s *DescribeDBInstanceConfigChangeLogResponseBodyDataParamChangeLogs) Validate() error {
	return dara.Validate(s)
}
