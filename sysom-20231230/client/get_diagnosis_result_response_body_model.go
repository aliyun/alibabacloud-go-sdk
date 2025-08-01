// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiagnosisResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDiagnosisResultResponseBody
	GetCode() *string
	SetData(v *GetDiagnosisResultResponseBodyData) *GetDiagnosisResultResponseBody
	GetData() *GetDiagnosisResultResponseBodyData
	SetMessage(v string) *GetDiagnosisResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDiagnosisResultResponseBody
	GetRequestId() *string
}

type GetDiagnosisResultResponseBody struct {
	// example:
	//
	// Success
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetDiagnosisResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 9515E5A0-8905-59B0-9BBF-5F0BE568C3A0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s GetDiagnosisResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDiagnosisResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDiagnosisResultResponseBody) GetData() *GetDiagnosisResultResponseBodyData {
	return s.Data
}

func (s *GetDiagnosisResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDiagnosisResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDiagnosisResultResponseBody) SetCode(v string) *GetDiagnosisResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetData(v *GetDiagnosisResultResponseBodyData) *GetDiagnosisResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetMessage(v string) *GetDiagnosisResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetRequestId(v string) *GetDiagnosisResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDiagnosisResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDiagnosisResultResponseBodyData struct {
	// example:
	//
	// 0
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {
	//
	//     "jobs":[
	//
	//         {
	//
	//             "cmd":"mkdir -p /var/log/sysak && sysak podmem -r 100  -a -j /var/log/sysak/podmem.json > /dev/null 2>&1 && cat /var/log/sysak/podmem.json",
	//
	//             "instance":"172.20.12.174",
	//
	//             "fetch_file_list":[
	//
	//             ]
	//
	//         }
	//
	//     ],
	//
	//     "in_order":true,
	//
	//     "offline_mode":false,
	//
	//     "offline_results":[
	//
	//     ]
	//
	// }
	Command   interface{} `json:"command,omitempty" xml:"command,omitempty"`
	CreatedAt *string     `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// example:
	//
	// Diagnosis failed
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// example:
	//
	// {
	//
	//     "type":"all",
	//
	//     "value":"",
	//
	//     "channel":"ssh",
	//
	//     "instance":"172.1.2.174",
	//
	//     "service_name":"filecache"
	//
	// }
	Params interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// {
	//
	//     "summary":"  memory cgroup leak",
	//
	//     "dataMemEvent":{
	//
	//         "data":[
	//
	//             {
	//
	//                 "key":"Util",
	//
	//                 "value":20
	//
	//             },
	//
	//             {
	//
	//                 "key":"MemLeak",
	//
	//                 "value":"OK"
	//
	//             },
	//
	//             {
	//
	//                 "key":"MemcgLeak",
	//
	//                 "value":"NG"
	//
	//             },
	//
	//             {
	//
	//                 "key":"MemFrag",
	//
	//                 "value":"OK"
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "dataMemOverView":{
	//
	//         "data":[
	//
	//             {
	//
	//                 "key":"app",
	//
	//                 "value":10937332
	//
	//             },
	//
	//             {
	//
	//                 "key":"free",
	//
	//                 "value":806800
	//
	//             },
	//
	//             {
	//
	//                 "key":"kernel",
	//
	//                 "value":4527660
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "dataKerMem":{
	//
	//         "data":[
	//
	//             {
	//
	//                 "key":"SReclaimable",
	//
	//                 "value":3411292
	//
	//             },
	//
	//             {
	//
	//                 "key":"VmallocUsed",
	//
	//                 "value":30980
	//
	//             },
	//
	//             {
	//
	//                 "key":"allocPage",
	//
	//                 "value":177732
	//
	//             },
	//
	//             {
	//
	//                 "key":"KernelStack",
	//
	//                 "value":9280
	//
	//             },
	//
	//             {
	//
	//                 "key":"PageTables",
	//
	//                 "value":38056
	//
	//             },
	//
	//             {
	//
	//                 "key":"SUnreclaim",
	//
	//                 "value":170248
	//
	//             },
	//
	//             {
	//
	//                 "key":"reserved",
	//
	//                 "value":690072
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "dataUserMem":{
	//
	//         "data":[
	//
	//             {
	//
	//                 "key":"filecache",
	//
	//                 "value":8010008
	//
	//             },
	//
	//             {
	//
	//                 "key":"anon",
	//
	//                 "value":2468608
	//
	//             },
	//
	//             {
	//
	//                 "key":"mlock",
	//
	//                 "value":0
	//
	//             },
	//
	//             {
	//
	//                 "key":"huge1G",
	//
	//                 "value":0
	//
	//             },
	//
	//             {
	//
	//                 "key":"huge2M",
	//
	//                 "value":0
	//
	//             },
	//
	//             {
	//
	//                 "key":"buffers",
	//
	//                 "value":458608
	//
	//             },
	//
	//             {
	//
	//                 "key":"shmem",
	//
	//                 "value":2284
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "dataCacheList":{
	//
	//         "data":[
	//
	//             {
	//
	//                 "key":0,
	//
	//                 "Name":"/var/lib/mysql/sysom/sys_handler_log.ibd",
	//
	//                 "cached":576764,
	//
	//                 "Task":"mysqld_78575 "
	//
	//             },
	//
	//             {
	//
	//                 "key":1,
	//
	//                 "Name":"/var/log/sysom/sysom-migration-access.log",
	//
	//                 "cached":276688,
	//
	//                 "Task":"gunicorn_33647 ,gunicorn_460836 ,gunicorn_559934 ,gunicorn_731758 ,gunicorn_2362682 "
	//
	//             },
	//
	//             {
	//
	//                 "key":2,
	//
	//                 "Name":"/var/log/sysom/sysom-rtdemo-access.log",
	//
	//                 "cached":229404,
	//
	//                 "Task":"gunicorn_60718 ,gunicorn_720734 ,gunicorn_722168 "
	//
	//             },
	//
	//             {
	//
	//                 "key":3,
	//
	//                 "Name":"/var/log/sysom/sysom-monitor-server-access.log",
	//
	//                 "cached":197368,
	//
	//                 "Task":"gunicorn_33682 ,gunicorn_671155 ,gunicorn_714998 "
	//
	//             },
	//
	//             {
	//
	//                 "key":4,
	//
	//                 "Name":"/var/log/sysom/sysom-channel-access.log",
	//
	//                 "cached":180276,
	//
	//                 "Task":"gunicorn_33233 ,gunicorn_499735 ,gunicorn_725497 "
	//
	//             },
	//
	//             {
	//
	//                 "key":5,
	//
	//                 "Name":"total cached of close file",
	//
	//                 "cached":3729668,
	//
	//                 "Task":""
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "dataProcMemList":{
	//
	//         "data":[
	//
	//             {
	//
	//                 "key":0,
	//
	//                 "task":"mysqld",
	//
	//                 "MemTotal":240856,
	//
	//                 "RssAnon":218248,
	//
	//                 "RssFile":22608
	//
	//             },
	//
	//             {
	//
	//                 "key":1,
	//
	//                 "task":"systemd-journal",
	//
	//                 "MemTotal":150248,
	//
	//                 "RssAnon":74300,
	//
	//                 "RssFile":75944
	//
	//             },
	//
	//             {
	//
	//                 "key":2,
	//
	//                 "task":"gunicorn",
	//
	//                 "MemTotal":144640,
	//
	//                 "RssAnon":114200,
	//
	//                 "RssFile":30440
	//
	//             },
	//
	//             {
	//
	//                 "key":3,
	//
	//                 "task":"gunicorn",
	//
	//                 "MemTotal":141480,
	//
	//                 "RssAnon":111040,
	//
	//                 "RssFile":30440
	//
	//             },
	//
	//             {
	//
	//                 "key":4,
	//
	//                 "task":"grafana-server",
	//
	//                 "MemTotal":103660,
	//
	//                 "RssAnon":42732,
	//
	//                 "RssFile":60928
	//
	//             },
	//
	//             {
	//
	//                 "key":5,
	//
	//                 "task":"gunicorn",
	//
	//                 "MemTotal":97444,
	//
	//                 "RssAnon":76256,
	//
	//                 "RssFile":21188
	//
	//             },
	//
	//             {
	//
	//                 "key":6,
	//
	//                 "task":"gunicorn",
	//
	//                 "MemTotal":97260,
	//
	//                 "RssAnon":76072,
	//
	//                 "RssFile":21188
	//
	//             },
	//
	//             {
	//
	//                 "key":7,
	//
	//                 "task":"prometheus",
	//
	//                 "MemTotal":95356,
	//
	//                 "RssAnon":45376,
	//
	//                 "RssFile":49980
	//
	//             },
	//
	//             {
	//
	//                 "key":8,
	//
	//                 "task":"gunicorn",
	//
	//                 "MemTotal":90144,
	//
	//                 "RssAnon":76456,
	//
	//                 "RssFile":13688
	//
	//             },
	//
	//             {
	//
	//                 "key":9,
	//
	//                 "task":"gunicorn",
	//
	//                 "MemTotal":89796,
	//
	//                 "RssAnon":76108,
	//
	//                 "RssFile":13688
	//
	//             }
	//
	//         ]
	//
	//     }
	//
	// }
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// memgraph
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// grcuU21a
	TaskId    *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// example:
	//
	// /diagnose/detail/qe3Z34sa
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDiagnosisResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDiagnosisResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *GetDiagnosisResultResponseBodyData) GetCommand() interface{} {
	return s.Command
}

func (s *GetDiagnosisResultResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetDiagnosisResultResponseBodyData) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *GetDiagnosisResultResponseBodyData) GetParams() interface{} {
	return s.Params
}

func (s *GetDiagnosisResultResponseBodyData) GetResult() interface{} {
	return s.Result
}

func (s *GetDiagnosisResultResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetDiagnosisResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDiagnosisResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDiagnosisResultResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetDiagnosisResultResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetDiagnosisResultResponseBodyData) SetCode(v int32) *GetDiagnosisResultResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetCommand(v interface{}) *GetDiagnosisResultResponseBodyData {
	s.Command = v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetCreatedAt(v string) *GetDiagnosisResultResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetErrMsg(v string) *GetDiagnosisResultResponseBodyData {
	s.ErrMsg = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetParams(v interface{}) *GetDiagnosisResultResponseBodyData {
	s.Params = v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetResult(v interface{}) *GetDiagnosisResultResponseBodyData {
	s.Result = v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetServiceName(v string) *GetDiagnosisResultResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetStatus(v string) *GetDiagnosisResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetTaskId(v string) *GetDiagnosisResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetUpdatedAt(v string) *GetDiagnosisResultResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetUrl(v string) *GetDiagnosisResultResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
