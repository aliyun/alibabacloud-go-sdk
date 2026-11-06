// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLExplorerRetentionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigValue(v string) *DescribeSQLExplorerRetentionResponseBody
	GetConfigValue() *string
	SetDBInstanceID(v int32) *DescribeSQLExplorerRetentionResponseBody
	GetDBInstanceID() *int32
	SetDBInstanceName(v string) *DescribeSQLExplorerRetentionResponseBody
	GetDBInstanceName() *string
	SetRequestId(v string) *DescribeSQLExplorerRetentionResponseBody
	GetRequestId() *string
}

type DescribeSQLExplorerRetentionResponseBody struct {
	// The retention period of SQL Explorer logs. Valid values:
	//
	// 	- **30**: 30 days
	//
	// 	- **180**: 180 days
	//
	// 	- **365**: 1 year
	//
	// 	- **1095**: 3 years
	//
	// 	- **1825**: 5 years
	//
	// example:
	//
	// 30
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// 35707295
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// pc-****************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6A2EE5B4-CC9F-46E1-A747-E43BC9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSQLExplorerRetentionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLExplorerRetentionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLExplorerRetentionResponseBody) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *DescribeSQLExplorerRetentionResponseBody) GetDBInstanceID() *int32 {
	return s.DBInstanceID
}

func (s *DescribeSQLExplorerRetentionResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSQLExplorerRetentionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLExplorerRetentionResponseBody) SetConfigValue(v string) *DescribeSQLExplorerRetentionResponseBody {
	s.ConfigValue = &v
	return s
}

func (s *DescribeSQLExplorerRetentionResponseBody) SetDBInstanceID(v int32) *DescribeSQLExplorerRetentionResponseBody {
	s.DBInstanceID = &v
	return s
}

func (s *DescribeSQLExplorerRetentionResponseBody) SetDBInstanceName(v string) *DescribeSQLExplorerRetentionResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSQLExplorerRetentionResponseBody) SetRequestId(v string) *DescribeSQLExplorerRetentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLExplorerRetentionResponseBody) Validate() error {
	return dara.Validate(s)
}
