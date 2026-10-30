// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGpuClusterListAlarmRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGpuClusterListAlarmRecordResponseBody
	GetCode() *string
	SetData(v []*GetGpuClusterListAlarmRecordResponseBodyData) *GetGpuClusterListAlarmRecordResponseBody
	GetData() []*GetGpuClusterListAlarmRecordResponseBodyData
	SetMessage(v string) *GetGpuClusterListAlarmRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGpuClusterListAlarmRecordResponseBody
	GetRequestId() *string
	SetTotal(v string) *GetGpuClusterListAlarmRecordResponseBody
	GetTotal() *string
}

type GetGpuClusterListAlarmRecordResponseBody struct {
	// The response code. Valid values: Success and Sysom.ServerError.
	//
	// example:
	//
	// Success or Sysom.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of alert data.
	Data []*GetGpuClusterListAlarmRecordResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error message returned when the code is not Success.
	//
	// example:
	//
	// server error
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// real request ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of returned records.
	//
	// example:
	//
	// 10
	Total *string `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetGpuClusterListAlarmRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGpuClusterListAlarmRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetGpuClusterListAlarmRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGpuClusterListAlarmRecordResponseBody) GetData() []*GetGpuClusterListAlarmRecordResponseBodyData {
	return s.Data
}

func (s *GetGpuClusterListAlarmRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGpuClusterListAlarmRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGpuClusterListAlarmRecordResponseBody) GetTotal() *string {
	return s.Total
}

func (s *GetGpuClusterListAlarmRecordResponseBody) SetCode(v string) *GetGpuClusterListAlarmRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBody) SetData(v []*GetGpuClusterListAlarmRecordResponseBodyData) *GetGpuClusterListAlarmRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBody) SetMessage(v string) *GetGpuClusterListAlarmRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBody) SetRequestId(v string) *GetGpuClusterListAlarmRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBody) SetTotal(v string) *GetGpuClusterListAlarmRecordResponseBody {
	s.Total = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGpuClusterListAlarmRecordResponseBodyData struct {
	// The cluster where the alert is triggered.
	//
	// example:
	//
	// cluster1
	AlarmCluster *string `json:"alarmCluster,omitempty" xml:"alarmCluster,omitempty"`
	// The extended information of the alert.
	//
	// example:
	//
	// {
	//
	// error_node_gpu,
	//
	// error_pod,
	//
	// task_pod_list,
	//
	// error_type,
	//
	// kernel,
	//
	// }
	AlarmExtInfo interface{} `json:"alarmExtInfo,omitempty" xml:"alarmExtInfo,omitempty"`
	// The unique ID of the alert.
	//
	// example:
	//
	// 746d13f8-9662-4441-b76b-a8ed8d485195
	AlarmId *string `json:"alarmId,omitempty" xml:"alarmId,omitempty"`
	// The alerting event log.
	//
	// example:
	//
	// cuda OOM：xxxxx
	AlarmLog *string `json:"alarmLog,omitempty" xml:"alarmLog,omitempty"`
	// The ID of the task where the anomaly occurred.
	//
	// example:
	//
	// task job id
	AlarmTask *string `json:"alarmTask,omitempty" xml:"alarmTask,omitempty"`
	// The date and time when the alert is triggered.
	//
	// example:
	//
	// 2025-08-22 18:16:25
	AlarmTime *string `json:"alarmTime,omitempty" xml:"alarmTime,omitempty"`
	// The topology snapshot.
	//
	// example:
	//
	// topology data，e.g.
	//
	//       "alarmTopology": {
	//
	//         "nodes": [
	//
	//           {
	//
	//             "gpu_list": [
	//
	//               {
	//
	//                 "values": {
	//
	//                   "Basic": [
	//
	//                     {
	//
	//                       "value": "1",
	//
	//                       "key": "Gpu Number\\n"
	//
	//                     }
	//
	//                   ],
	//
	//                   "Metric": [
	//
	//                     {
	//
	//                       "value": "50",
	//
	//                       "key": "gpu_util(%)\\n"
	//
	//                     }
	//
	//                   ],
	//
	//                   "Exception": [
	//
	//                     {
	//
	//                       "value": "NCCL HANG",
	//
	//                       "key": "error0"
	//
	//                     }
	//
	//                   ]
	//
	//                 },
	//
	//                 "name": "GPU1",
	//
	//                 "style": {
	//
	//                   "fill": "#e94242\\n",
	//
	//                   "opacity": "0.1"
	//
	//                 },
	//
	//                 "id": "6cc37fbc"
	//
	//               }
	//
	//             ],
	//
	//             "values": {
	//
	//               "Basic": [
	//
	//                 {
	//
	//                   "value": "i-xxxxxx",
	//
	//                   "key": "instance"
	//
	//                 }
	//
	//               ],
	//
	//               "Metric": [
	//
	//                 {
	//
	//                   "value": "value of metrics 1",
	//
	//                   "key": "name of metric 1"
	//
	//                 }
	//
	//               ],
	//
	//               "Exception": [
	//
	//                 {
	//
	//                   "value": "NCCL HANG",
	//
	//                   "key": "error0"
	//
	//                 }
	//
	//               ]
	//
	//             },
	//
	//             "name": "i-xxxxxx\\n",
	//
	//             "style": {
	//
	//               "fillOpacity": "0.5",
	//
	//               "fill": "#e94242\\n"
	//
	//             },
	//
	//             "id": "975fbcb0"
	//
	//           }
	//
	//         ],
	//
	//         "edges": [
	//
	//           {
	//
	//             "src_id": "00534a88",
	//
	//             "values": {
	//
	//               "Basic": [
	//
	//                 {
	//
	//                   "value": "v2.22.3\\n",
	//
	//                   "key": "NCCL Version\\n"
	//
	//                 }
	//
	//               ],
	//
	//               "Metric": [
	//
	//                 {
	//
	//                   "value": "10000",
	//
	//                   "key": "Recv Msg Count"
	//
	//                 }
	//
	//               ],
	//
	//               "Exception": [
	//
	//                 {
	//
	//                   "value": "NCCL HANG",
	//
	//                   "key": "error0"
	//
	//                 }
	//
	//               ]
	//
	//             },
	//
	//             "dst_id": "260525e9",
	//
	//             "style": {
	//
	//               "opacity": "0.1",
	//
	//               "lineWidth": "1",
	//
	//               "selectedStyle": "\\"{\\\\\\"opacity\\\\\\": 1, \\\\\\"lineWidth\\\\\\": 1}\\"\\n",
	//
	//               "stroke": "#e94242"
	//
	//             },
	//
	//             "type": "GPU"
	//
	//           }
	//
	//         ]
	//
	//       },
	AlarmTopology interface{} `json:"alarmTopology,omitempty" xml:"alarmTopology,omitempty"`
	// The alerting type.
	//
	// example:
	//
	// cuda OOM
	AlarmType *string `json:"alarmType,omitempty" xml:"alarmType,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// running、blocking、restart、exited、failed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The UUID of the cloud account.
	//
	// example:
	//
	// 1234123412352311
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetGpuClusterListAlarmRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGpuClusterListAlarmRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) GetAlarmCluster() *string {
	return s.AlarmCluster
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) GetAlarmExtInfo() interface{} {
	return s.AlarmExtInfo
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) GetAlarmId() *string {
	return s.AlarmId
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) GetAlarmLog() *string {
	return s.AlarmLog
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) GetAlarmTask() *string {
	return s.AlarmTask
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) GetAlarmTime() *string {
	return s.AlarmTime
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) GetAlarmTopology() interface{} {
	return s.AlarmTopology
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) GetAlarmType() *string {
	return s.AlarmType
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) SetAlarmCluster(v string) *GetGpuClusterListAlarmRecordResponseBodyData {
	s.AlarmCluster = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) SetAlarmExtInfo(v interface{}) *GetGpuClusterListAlarmRecordResponseBodyData {
	s.AlarmExtInfo = v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) SetAlarmId(v string) *GetGpuClusterListAlarmRecordResponseBodyData {
	s.AlarmId = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) SetAlarmLog(v string) *GetGpuClusterListAlarmRecordResponseBodyData {
	s.AlarmLog = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) SetAlarmTask(v string) *GetGpuClusterListAlarmRecordResponseBodyData {
	s.AlarmTask = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) SetAlarmTime(v string) *GetGpuClusterListAlarmRecordResponseBodyData {
	s.AlarmTime = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) SetAlarmTopology(v interface{}) *GetGpuClusterListAlarmRecordResponseBodyData {
	s.AlarmTopology = v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) SetAlarmType(v string) *GetGpuClusterListAlarmRecordResponseBodyData {
	s.AlarmType = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) SetStatus(v string) *GetGpuClusterListAlarmRecordResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) SetUid(v string) *GetGpuClusterListAlarmRecordResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
