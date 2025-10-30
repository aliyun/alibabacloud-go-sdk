// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIMVInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeIMVInfosResponseBody
	GetDBInstanceId() *string
	SetImvInfos(v []*DescribeIMVInfosResponseBodyImvInfos) *DescribeIMVInfosResponseBody
	GetImvInfos() []*DescribeIMVInfosResponseBodyImvInfos
	SetRequestId(v string) *DescribeIMVInfosResponseBody
	GetRequestId() *string
}

type DescribeIMVInfosResponseBody struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The queried materialized views.
	ImvInfos []*DescribeIMVInfosResponseBodyImvInfos `json:"ImvInfos,omitempty" xml:"ImvInfos,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIMVInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIMVInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIMVInfosResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeIMVInfosResponseBody) GetImvInfos() []*DescribeIMVInfosResponseBodyImvInfos {
	return s.ImvInfos
}

func (s *DescribeIMVInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIMVInfosResponseBody) SetDBInstanceId(v string) *DescribeIMVInfosResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeIMVInfosResponseBody) SetImvInfos(v []*DescribeIMVInfosResponseBodyImvInfos) *DescribeIMVInfosResponseBody {
	s.ImvInfos = v
	return s
}

func (s *DescribeIMVInfosResponseBody) SetRequestId(v string) *DescribeIMVInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIMVInfosResponseBody) Validate() error {
	if s.ImvInfos != nil {
		for _, item := range s.ImvInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIMVInfosResponseBodyImvInfos struct {
	// The name of the table based on which the materialized view is created.
	//
	// example:
	//
	// "public."t2"
	Base *string `json:"Base,omitempty" xml:"Base,omitempty"`
	// The dependency between the materialized view and the base table and all metric values, which can be used to build a lineage graph.
	//
	// example:
	//
	// {\\"maintenance_calls\\" : 1, \\"avg_apply_time\\" : 2, \\"avg_calc_rows\\" : 1, \\"avg_calc_time\\" : 11, \\"avg_delta_rows\\" : 1, \\"avg_maintenance_total_time\\" : 14, \\"avg_maintenance_total_time_total\\" : 14, \\"max_apply_time\\" : 2, \\"max_calc_rows\\" : 1, \\"max_calc_time\\" : 11, \\"max_delta_rows\\" : 1, \\"max_maintenance_total_time\\" : 14, \\"max_maintenance_total_time_total\\" : 14, \\"min_apply_time\\" : 2, \\"min_calc_rows\\" : 1, \\"min_calc_time\\" : 11, \\"min_delta_rows\\" : 1, \\"min_maintenance_total_time\\" : 14, \\"min_maintenance_total_time_total\\" : 14, \\"max_outerjoin_apply_time\\" : null, \\"max_outerjoin_calc_rows\\" : null, \\"max_outerjoin_calc_time\\" : null, \\"max_outerjoin_delta_rows\\" : null, \\"avg_outerjoin_apply_time\\" : null, \\"avg_outerjoin_calc_rows\\" : null, \\"avg_outerjoin_calc_time\\" : null, \\"avg_outerjoin_delta_rows\\" : null, \\"min_outerjoin_apply_time\\" : null, \\"min_outerjoin_calc_rows\\" : null, \\"min_outerjoin_calc_time\\" : null, \\"min_outerjoin_delta_rows\\" : null, \\"create_rows\\" : null, \\"create_time\\" : null, \\"direct_visited\\" : null, \\"indirect_visited\\" : null, \\"max_refresh_rows\\" : null, \\"max_refresh_time\\" : null, \\"avg_refresh_rows\\" : null, \\"avg_refresh_time\\" : null, \\"min_refresh_rows\\" : null, \\"min_refresh_time\\" : null, \\"refresh_calls\\" : null, \\"avg_wait_lock_time\\" : null, \\"max_wait_lock_time\\" : null, \\"min_wait_lock_time\\" : null, \\"latest_maintenance_time\\" : \\"2023-08-09T07:39:14.753252+00:00\\"}
	DetailInfo *string `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty"`
	// The name of the materialized view.
	//
	// example:
	//
	// public."mv1"
	MV *string `json:"MV,omitempty" xml:"MV,omitempty"`
}

func (s DescribeIMVInfosResponseBodyImvInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeIMVInfosResponseBodyImvInfos) GoString() string {
	return s.String()
}

func (s *DescribeIMVInfosResponseBodyImvInfos) GetBase() *string {
	return s.Base
}

func (s *DescribeIMVInfosResponseBodyImvInfos) GetDetailInfo() *string {
	return s.DetailInfo
}

func (s *DescribeIMVInfosResponseBodyImvInfos) GetMV() *string {
	return s.MV
}

func (s *DescribeIMVInfosResponseBodyImvInfos) SetBase(v string) *DescribeIMVInfosResponseBodyImvInfos {
	s.Base = &v
	return s
}

func (s *DescribeIMVInfosResponseBodyImvInfos) SetDetailInfo(v string) *DescribeIMVInfosResponseBodyImvInfos {
	s.DetailInfo = &v
	return s
}

func (s *DescribeIMVInfosResponseBodyImvInfos) SetMV(v string) *DescribeIMVInfosResponseBodyImvInfos {
	s.MV = &v
	return s
}

func (s *DescribeIMVInfosResponseBodyImvInfos) Validate() error {
	return dara.Validate(s)
}
