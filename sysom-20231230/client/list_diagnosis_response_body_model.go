// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDiagnosisResponseBody
	GetRequestId() *string
	SetCode(v string) *ListDiagnosisResponseBody
	GetCode() *string
	SetData(v []*ListDiagnosisResponseBodyData) *ListDiagnosisResponseBody
	GetData() []*ListDiagnosisResponseBodyData
	SetMessage(v string) *ListDiagnosisResponseBody
	GetMessage() *string
	SetTotal(v int64) *ListDiagnosisResponseBody
	GetTotal() *int64
}

type ListDiagnosisResponseBody struct {
	// Request ID, which can be used for end-to-end diagnosis
	//
	// example:
	//
	// 44841312-7227-55C9-AE03-D59729BFAE38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Status code
	//
	// - `code == Success` indicates successful authorization;
	//
	// - Other status codes indicate failed authorization. When authorization fails, view the `message` field to obtain detailed error information;
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned data
	Data []*ListDiagnosisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Error message
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
	// This parameter is required.
	//
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Total count
	//
	// example:
	//
	// 319
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDiagnosisResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDiagnosisResponseBody) GetData() []*ListDiagnosisResponseBodyData {
	return s.Data
}

func (s *ListDiagnosisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDiagnosisResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListDiagnosisResponseBody) SetRequestId(v string) *ListDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnosisResponseBody) SetCode(v string) *ListDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *ListDiagnosisResponseBody) SetData(v []*ListDiagnosisResponseBodyData) *ListDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *ListDiagnosisResponseBody) SetMessage(v string) *ListDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *ListDiagnosisResponseBody) SetTotal(v int64) *ListDiagnosisResponseBody {
	s.Total = &v
	return s
}

func (s *ListDiagnosisResponseBody) Validate() error {
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

type ListDiagnosisResponseBodyData struct {
	// Diagnosis error code; 0 indicates no error
	//
	// example:
	//
	// 0
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// Diagnostic command
	//
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
	Command interface{} `json:"command,omitempty" xml:"command,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 2024-12-25T15:08:19
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Error message
	//
	// example:
	//
	// Diagnosis failed
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// Diagnosis parameters
	//
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
	// Diagnosis result
	//
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
	// Diagnosis Type
	//
	// example:
	//
	// memgraph
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	// Status of the diagnostic task execution.
	//
	// Valid values:
	//
	// - **Ready**: Ready
	//
	// - **Running**: Running
	//
	// - **Success**: Succeeded
	//
	// - **Fail**: Failed
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Job ID.
	//
	// example:
	//
	// grcuU21a
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// Update Time
	//
	// example:
	//
	// 2024-12-25T15:08:19
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Diagnostic details URL
	//
	// example:
	//
	// /diagnose/detail/qe3Z34sa
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListDiagnosisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDiagnosisResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *ListDiagnosisResponseBodyData) GetCommand() interface{} {
	return s.Command
}

func (s *ListDiagnosisResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListDiagnosisResponseBodyData) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *ListDiagnosisResponseBodyData) GetParams() interface{} {
	return s.Params
}

func (s *ListDiagnosisResponseBodyData) GetResult() interface{} {
	return s.Result
}

func (s *ListDiagnosisResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListDiagnosisResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListDiagnosisResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListDiagnosisResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListDiagnosisResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *ListDiagnosisResponseBodyData) SetCode(v int32) *ListDiagnosisResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetCommand(v interface{}) *ListDiagnosisResponseBodyData {
	s.Command = v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetCreatedAt(v string) *ListDiagnosisResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetErrMsg(v string) *ListDiagnosisResponseBodyData {
	s.ErrMsg = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetParams(v interface{}) *ListDiagnosisResponseBodyData {
	s.Params = v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetResult(v interface{}) *ListDiagnosisResponseBodyData {
	s.Result = v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetServiceName(v string) *ListDiagnosisResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetStatus(v string) *ListDiagnosisResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetTaskId(v string) *ListDiagnosisResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetUpdatedAt(v string) *ListDiagnosisResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetUrl(v string) *ListDiagnosisResponseBodyData {
	s.Url = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) Validate() error {
	return dara.Validate(s)
}
