// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v *DescribePlaybookMetricsResponseBodyMetrics) *DescribePlaybookMetricsResponseBody
	GetMetrics() *DescribePlaybookMetricsResponseBodyMetrics
	SetRequestId(v string) *DescribePlaybookMetricsResponseBody
	GetRequestId() *string
}

type DescribePlaybookMetricsResponseBody struct {
	// The details of the playbook.
	Metrics *DescribePlaybookMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 567D3D0B-2153-5860-BF9A-F9DEED55FB73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookMetricsResponseBody) GetMetrics() *DescribePlaybookMetricsResponseBodyMetrics {
	return s.Metrics
}

func (s *DescribePlaybookMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlaybookMetricsResponseBody) SetMetrics(v *DescribePlaybookMetricsResponseBodyMetrics) *DescribePlaybookMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *DescribePlaybookMetricsResponseBody) SetRequestId(v string) *DescribePlaybookMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBody) Validate() error {
	if s.Metrics != nil {
		if err := s.Metrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePlaybookMetricsResponseBodyMetrics struct {
	// The status of the playbook. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	Active *int32 `json:"Active,omitempty" xml:"Active,omitempty"`
	// The description of the playbook.
	//
	// example:
	//
	// This is a playbook for waf processing
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the playbook.
	//
	// example:
	//
	// demo name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The number of the tasks that are created for the playbook and failed to run.
	//
	// example:
	//
	// 10
	FailNum *int32 `json:"FailNum,omitempty" xml:"FailNum,omitempty"`
	// The time when the playbook was created. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1655277397000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The number of historical versions of the playbook.
	//
	// example:
	//
	// 10
	HistoryMd5 *int32 `json:"HistoryMd5,omitempty" xml:"HistoryMd5,omitempty"`
	// The time when the playbook was last run. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1683526277415
	LastRuntime *int64 `json:"LastRuntime,omitempty" xml:"LastRuntime,omitempty"`
	// The type of the playbook. Valid values:
	//
	// 	- **preset**: predefined playbook
	//
	// 	- **user**: custom playbook
	//
	// example:
	//
	// user
	OwnType *string `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// 0fbc9bdb-9ae3-4ef4-a709-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The number of the tasks that are created for the playbook and were successfully run.
	//
	// example:
	//
	// 100
	SuccNum *int32 `json:"SuccNum,omitempty" xml:"SuccNum,omitempty"`
}

func (s DescribePlaybookMetricsResponseBodyMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) GetActive() *int32 {
	return s.Active
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) GetDescription() *string {
	return s.Description
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) GetFailNum() *int32 {
	return s.FailNum
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) GetHistoryMd5() *int32 {
	return s.HistoryMd5
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) GetLastRuntime() *int64 {
	return s.LastRuntime
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) GetOwnType() *string {
	return s.OwnType
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) GetSuccNum() *int32 {
	return s.SuccNum
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetActive(v int32) *DescribePlaybookMetricsResponseBodyMetrics {
	s.Active = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetDescription(v string) *DescribePlaybookMetricsResponseBodyMetrics {
	s.Description = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetDisplayName(v string) *DescribePlaybookMetricsResponseBodyMetrics {
	s.DisplayName = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetFailNum(v int32) *DescribePlaybookMetricsResponseBodyMetrics {
	s.FailNum = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetGmtCreate(v int64) *DescribePlaybookMetricsResponseBodyMetrics {
	s.GmtCreate = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetHistoryMd5(v int32) *DescribePlaybookMetricsResponseBodyMetrics {
	s.HistoryMd5 = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetLastRuntime(v int64) *DescribePlaybookMetricsResponseBodyMetrics {
	s.LastRuntime = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetOwnType(v string) *DescribePlaybookMetricsResponseBodyMetrics {
	s.OwnType = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetPlaybookUuid(v string) *DescribePlaybookMetricsResponseBodyMetrics {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) SetSuccNum(v int32) *DescribePlaybookMetricsResponseBodyMetrics {
	s.SuccNum = &v
	return s
}

func (s *DescribePlaybookMetricsResponseBodyMetrics) Validate() error {
	return dara.Validate(s)
}
