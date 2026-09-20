// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestorePlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateRestorePlanRequest
	GetClusterId() *string
	SetRestoreAllTable(v bool) *CreateRestorePlanRequest
	GetRestoreAllTable() *bool
	SetRestoreByCopy(v bool) *CreateRestorePlanRequest
	GetRestoreByCopy() *bool
	SetRestoreToDate(v string) *CreateRestorePlanRequest
	GetRestoreToDate() *string
	SetTables(v string) *CreateRestorePlanRequest
	GetTables() *string
	SetTargetClusterId(v string) *CreateRestorePlanRequest
	GetTargetClusterId() *string
}

type CreateRestorePlanRequest struct {
	// The ID of the ApsaraDB for HBase Performance-enhanced Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to restore all tables. Valid values:
	//
	// - **true**: Restores all tables in the ApsaraDB for HBase Performance-enhanced Edition cluster.
	//
	// - **false**: Does not restore all tables in the ApsaraDB for HBase Performance-enhanced Edition cluster.
	//
	// > If this parameter is set to **true**, the **Tables*	- parameter is invalid. If this parameter is set to **false**, the **Tables*	- parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	RestoreAllTable *bool `json:"RestoreAllTable,omitempty" xml:"RestoreAllTable,omitempty"`
	// Specifies whether to restore data by using the copy method. Set the value to **true**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	RestoreByCopy *bool `json:"RestoreByCopy,omitempty" xml:"RestoreByCopy,omitempty"`
	// The point in time to which you want to restore data. The point in time must be within the recoverable time range. You can call the [DescribeRecoverableTimeRange](https://help.aliyun.com/document_detail/188365.html) operation to query the recoverable time range.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-11-05T05:49:42Z
	RestoreToDate *string `json:"RestoreToDate,omitempty" xml:"RestoreToDate,omitempty"`
	// The table names. Specify one table name per line. Wildcards (*) are not supported.
	//
	// - To restore to the current table, use the format: `namespace:table`. Example: `default:testTable`.
	//
	// - To restore to a different table, use the format: `namespace:table/namespace:table2`. Example: `default:testTable/default:testTable2`.
	//
	// example:
	//
	// test_ns:test_table/test_ns:test_table2
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
	// The ID of the ApsaraDB for HBase Performance-enhanced Edition cluster to which data is restored. You can also restore data to the cluster that is currently backed up.
	//
	// > The specified ApsaraDB for HBase Performance-enhanced Edition cluster and the backed-up ApsaraDB for HBase Performance-enhanced Edition cluster must meet the following requirements:<ul>
	//
	// <li>They are of the same version.</li>
	//
	// <li>They are in the same region.</li>
	//
	// <li>They are associated with the BDS cluster.</li></ul>.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp169l540vc6c****
	TargetClusterId *string `json:"TargetClusterId,omitempty" xml:"TargetClusterId,omitempty"`
}

func (s CreateRestorePlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRestorePlanRequest) GoString() string {
	return s.String()
}

func (s *CreateRestorePlanRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateRestorePlanRequest) GetRestoreAllTable() *bool {
	return s.RestoreAllTable
}

func (s *CreateRestorePlanRequest) GetRestoreByCopy() *bool {
	return s.RestoreByCopy
}

func (s *CreateRestorePlanRequest) GetRestoreToDate() *string {
	return s.RestoreToDate
}

func (s *CreateRestorePlanRequest) GetTables() *string {
	return s.Tables
}

func (s *CreateRestorePlanRequest) GetTargetClusterId() *string {
	return s.TargetClusterId
}

func (s *CreateRestorePlanRequest) SetClusterId(v string) *CreateRestorePlanRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateRestorePlanRequest) SetRestoreAllTable(v bool) *CreateRestorePlanRequest {
	s.RestoreAllTable = &v
	return s
}

func (s *CreateRestorePlanRequest) SetRestoreByCopy(v bool) *CreateRestorePlanRequest {
	s.RestoreByCopy = &v
	return s
}

func (s *CreateRestorePlanRequest) SetRestoreToDate(v string) *CreateRestorePlanRequest {
	s.RestoreToDate = &v
	return s
}

func (s *CreateRestorePlanRequest) SetTables(v string) *CreateRestorePlanRequest {
	s.Tables = &v
	return s
}

func (s *CreateRestorePlanRequest) SetTargetClusterId(v string) *CreateRestorePlanRequest {
	s.TargetClusterId = &v
	return s
}

func (s *CreateRestorePlanRequest) Validate() error {
	return dara.Validate(s)
}
