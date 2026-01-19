// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivyComputeSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListLivyComputeSessionsResponseBody
	GetRequestId() *string
	SetSessions(v []*ListLivyComputeSessionsResponseBodySessions) *ListLivyComputeSessionsResponseBody
	GetSessions() []*ListLivyComputeSessionsResponseBodySessions
	SetTotalCount(v int32) *ListLivyComputeSessionsResponseBody
	GetTotalCount() *int32
}

type ListLivyComputeSessionsResponseBody struct {
	// example:
	//
	// 8FAA8EEC-3026-5D15-8733-4E2A3DD970A1
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Sessions  []*ListLivyComputeSessionsResponseBodySessions `json:"sessions,omitempty" xml:"sessions,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListLivyComputeSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLivyComputeSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLivyComputeSessionsResponseBody) GetSessions() []*ListLivyComputeSessionsResponseBodySessions {
	return s.Sessions
}

func (s *ListLivyComputeSessionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLivyComputeSessionsResponseBody) SetRequestId(v string) *ListLivyComputeSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBody) SetSessions(v []*ListLivyComputeSessionsResponseBodySessions) *ListLivyComputeSessionsResponseBody {
	s.Sessions = v
	return s
}

func (s *ListLivyComputeSessionsResponseBody) SetTotalCount(v int32) *ListLivyComputeSessionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBody) Validate() error {
	if s.Sessions != nil {
		for _, item := range s.Sessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLivyComputeSessionsResponseBodySessions struct {
	// example:
	//
	// lc-xxxxxx
	ComputeId *string `json:"computeId,omitempty" xml:"computeId,omitempty"`
	// example:
	//
	// 1768213240000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 322.5
	CuHours *float64 `json:"cuHours,omitempty" xml:"cuHours,omitempty"`
	// example:
	//
	// 1768213240000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// null
	Info *string `json:"info,omitempty" xml:"info,omitempty"`
	// example:
	//
	// 1098888
	MbSeconds *int64 `json:"mbSeconds,omitempty" xml:"mbSeconds,omitempty"`
	// example:
	//
	// test_session
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// root_queue
	Queue *string `json:"queue,omitempty" xml:"queue,omitempty"`
	// example:
	//
	// livy-xxxxxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// {
	//
	//     "proxyUser": "test",
	//
	//     "conf": {
	//
	//         "spark.driver.cores": 1
	//
	//     }
	//
	// }
	SparkConf *string `json:"sparkConf,omitempty" xml:"sparkConf,omitempty"`
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// 343
	VcoreSeconds *int64 `json:"vcoreSeconds,omitempty" xml:"vcoreSeconds,omitempty"`
	// example:
	//
	// http://emr-spark-ui-cn-hangzhou.data.aliyun.com
	WebUI *string `json:"webUI,omitempty" xml:"webUI,omitempty"`
}

func (s ListLivyComputeSessionsResponseBodySessions) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeSessionsResponseBodySessions) GoString() string {
	return s.String()
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetComputeId() *string {
	return s.ComputeId
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetCuHours() *float64 {
	return s.CuHours
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetInfo() *string {
	return s.Info
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetMbSeconds() *int64 {
	return s.MbSeconds
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetName() *string {
	return s.Name
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetQueue() *string {
	return s.Queue
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetSessionId() *string {
	return s.SessionId
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetSparkConf() *string {
	return s.SparkConf
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetState() *string {
	return s.State
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetVcoreSeconds() *int64 {
	return s.VcoreSeconds
}

func (s *ListLivyComputeSessionsResponseBodySessions) GetWebUI() *string {
	return s.WebUI
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetComputeId(v string) *ListLivyComputeSessionsResponseBodySessions {
	s.ComputeId = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetCreateTime(v int64) *ListLivyComputeSessionsResponseBodySessions {
	s.CreateTime = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetCuHours(v float64) *ListLivyComputeSessionsResponseBodySessions {
	s.CuHours = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetEndTime(v int64) *ListLivyComputeSessionsResponseBodySessions {
	s.EndTime = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetInfo(v string) *ListLivyComputeSessionsResponseBodySessions {
	s.Info = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetMbSeconds(v int64) *ListLivyComputeSessionsResponseBodySessions {
	s.MbSeconds = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetName(v string) *ListLivyComputeSessionsResponseBodySessions {
	s.Name = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetQueue(v string) *ListLivyComputeSessionsResponseBodySessions {
	s.Queue = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetSessionId(v string) *ListLivyComputeSessionsResponseBodySessions {
	s.SessionId = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetSparkConf(v string) *ListLivyComputeSessionsResponseBodySessions {
	s.SparkConf = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetState(v string) *ListLivyComputeSessionsResponseBodySessions {
	s.State = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetVcoreSeconds(v int64) *ListLivyComputeSessionsResponseBodySessions {
	s.VcoreSeconds = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) SetWebUI(v string) *ListLivyComputeSessionsResponseBodySessions {
	s.WebUI = &v
	return s
}

func (s *ListLivyComputeSessionsResponseBodySessions) Validate() error {
	return dara.Validate(s)
}
