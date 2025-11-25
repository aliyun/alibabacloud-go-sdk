// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticQpsRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticQpsList(v []*DescribeElasticQpsRecordResponseBodyElasticQpsList) *DescribeElasticQpsRecordResponseBody
	GetElasticQpsList() []*DescribeElasticQpsRecordResponseBodyElasticQpsList
	SetRequestId(v string) *DescribeElasticQpsRecordResponseBody
	GetRequestId() *string
}

type DescribeElasticQpsRecordResponseBody struct {
	// The QPS information about the instance.
	ElasticQpsList []*DescribeElasticQpsRecordResponseBodyElasticQpsList `json:"ElasticQpsList,omitempty" xml:"ElasticQpsList,omitempty" type:"Repeated"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F68B34E2-570C-508D-95FD-DFB6611D518F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticQpsRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticQpsRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticQpsRecordResponseBody) GetElasticQpsList() []*DescribeElasticQpsRecordResponseBodyElasticQpsList {
	return s.ElasticQpsList
}

func (s *DescribeElasticQpsRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticQpsRecordResponseBody) SetElasticQpsList(v []*DescribeElasticQpsRecordResponseBodyElasticQpsList) *DescribeElasticQpsRecordResponseBody {
	s.ElasticQpsList = v
	return s
}

func (s *DescribeElasticQpsRecordResponseBody) SetRequestId(v string) *DescribeElasticQpsRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticQpsRecordResponseBody) Validate() error {
	if s.ElasticQpsList != nil {
		for _, item := range s.ElasticQpsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticQpsRecordResponseBodyElasticQpsList struct {
	// The timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1688140799999
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
	// The ID of the Anti-DDoS Proxy instance.
	//
	// example:
	//
	// ddoscoo-cn-7e225i41****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the Anti-DDoS Proxy instance.
	//
	// example:
	//
	// 203.***.***.199
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The burstable QPS value. A value of 0 indicates that the burstable QPS feature is not enabled.
	//
	// example:
	//
	// 300000
	OpsElasticQps *int64 `json:"OpsElasticQps,omitempty" xml:"OpsElasticQps,omitempty"`
	// The service QPS (active).
	//
	// example:
	//
	// 1345
	OpsQps *int64 `json:"OpsQps,omitempty" xml:"OpsQps,omitempty"`
	// The service QPS (purchased).
	//
	// example:
	//
	// 1345
	OriginQps *int64 `json:"OriginQps,omitempty" xml:"OriginQps,omitempty"`
	// The daily peak 95th percentile QPS.
	//
	// example:
	//
	// 4367
	Qps *int64 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The daily peak traffic.
	//
	// example:
	//
	// 122
	QpsPeak *int64 `json:"QpsPeak,omitempty" xml:"QpsPeak,omitempty"`
	// Indicates whether the instance has expired or is released. Valid values:
	//
	// 	- **1**: The instance runs as expected.
	//
	// 	- **2**: The instance has expired.
	//
	// 	- **4**: The instance is released.
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeElasticQpsRecordResponseBodyElasticQpsList) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticQpsRecordResponseBodyElasticQpsList) GoString() string {
	return s.String()
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) GetDate() *int64 {
	return s.Date
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) GetIp() *string {
	return s.Ip
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) GetOpsElasticQps() *int64 {
	return s.OpsElasticQps
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) GetOpsQps() *int64 {
	return s.OpsQps
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) GetOriginQps() *int64 {
	return s.OriginQps
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) GetQps() *int64 {
	return s.Qps
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) GetQpsPeak() *int64 {
	return s.QpsPeak
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) SetDate(v int64) *DescribeElasticQpsRecordResponseBodyElasticQpsList {
	s.Date = &v
	return s
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) SetInstanceId(v string) *DescribeElasticQpsRecordResponseBodyElasticQpsList {
	s.InstanceId = &v
	return s
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) SetIp(v string) *DescribeElasticQpsRecordResponseBodyElasticQpsList {
	s.Ip = &v
	return s
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) SetOpsElasticQps(v int64) *DescribeElasticQpsRecordResponseBodyElasticQpsList {
	s.OpsElasticQps = &v
	return s
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) SetOpsQps(v int64) *DescribeElasticQpsRecordResponseBodyElasticQpsList {
	s.OpsQps = &v
	return s
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) SetOriginQps(v int64) *DescribeElasticQpsRecordResponseBodyElasticQpsList {
	s.OriginQps = &v
	return s
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) SetQps(v int64) *DescribeElasticQpsRecordResponseBodyElasticQpsList {
	s.Qps = &v
	return s
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) SetQpsPeak(v int64) *DescribeElasticQpsRecordResponseBodyElasticQpsList {
	s.QpsPeak = &v
	return s
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) SetStatus(v int64) *DescribeElasticQpsRecordResponseBodyElasticQpsList {
	s.Status = &v
	return s
}

func (s *DescribeElasticQpsRecordResponseBodyElasticQpsList) Validate() error {
	return dara.Validate(s)
}
