// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AuthDiagnosisRequest struct {
	AutoCreateRole   *bool                            `json:"autoCreateRole,omitempty" xml:"autoCreateRole,omitempty"`
	AutoInstallAgent *bool                            `json:"autoInstallAgent,omitempty" xml:"autoInstallAgent,omitempty"`
	Instances        []*AuthDiagnosisRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s AuthDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisRequest) SetAutoCreateRole(v bool) *AuthDiagnosisRequest {
	s.AutoCreateRole = &v
	return s
}

func (s *AuthDiagnosisRequest) SetAutoInstallAgent(v bool) *AuthDiagnosisRequest {
	s.AutoInstallAgent = &v
	return s
}

func (s *AuthDiagnosisRequest) SetInstances(v []*AuthDiagnosisRequestInstances) *AuthDiagnosisRequest {
	s.Instances = v
	return s
}

type AuthDiagnosisRequestInstances struct {
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s AuthDiagnosisRequestInstances) String() string {
	return tea.Prettify(s)
}

func (s AuthDiagnosisRequestInstances) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisRequestInstances) SetInstance(v string) *AuthDiagnosisRequestInstances {
	s.Instance = &v
	return s
}

func (s *AuthDiagnosisRequestInstances) SetRegion(v string) *AuthDiagnosisRequestInstances {
	s.Region = &v
	return s
}

type AuthDiagnosisResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s AuthDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisResponseBody) SetCode(v string) *AuthDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *AuthDiagnosisResponseBody) SetData(v interface{}) *AuthDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *AuthDiagnosisResponseBody) SetMessage(v string) *AuthDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *AuthDiagnosisResponseBody) SetRequestId(v string) *AuthDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

type AuthDiagnosisResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisResponse) SetHeaders(v map[string]*string) *AuthDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *AuthDiagnosisResponse) SetStatusCode(v int32) *AuthDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthDiagnosisResponse) SetBody(v *AuthDiagnosisResponseBody) *AuthDiagnosisResponse {
	s.Body = v
	return s
}

type GenerateCopilotResponseRequest struct {
	LlmParamString *string `json:"llmParamString,omitempty" xml:"llmParamString,omitempty"`
}

func (s GenerateCopilotResponseRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateCopilotResponseRequest) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseRequest) SetLlmParamString(v string) *GenerateCopilotResponseRequest {
	s.LlmParamString = &v
	return s
}

type GenerateCopilotResponseResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// Requests for llm service failed
	Massage *string `json:"massage,omitempty" xml:"massage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateCopilotResponseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateCopilotResponseResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseResponseBody) SetCode(v string) *GenerateCopilotResponseResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetData(v string) *GenerateCopilotResponseResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetMassage(v string) *GenerateCopilotResponseResponseBody {
	s.Massage = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetRequestId(v string) *GenerateCopilotResponseResponseBody {
	s.RequestId = &v
	return s
}

type GenerateCopilotResponseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCopilotResponseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCopilotResponseResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateCopilotResponseResponse) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseResponse) SetHeaders(v map[string]*string) *GenerateCopilotResponseResponse {
	s.Headers = v
	return s
}

func (s *GenerateCopilotResponseResponse) SetStatusCode(v int32) *GenerateCopilotResponseResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCopilotResponseResponse) SetBody(v *GenerateCopilotResponseResponseBody) *GenerateCopilotResponseResponse {
	s.Body = v
	return s
}

type GetAbnormalEventsCountRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// test-pod
	Pod *string `json:"pod,omitempty" xml:"pod,omitempty"`
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetAbnormalEventsCountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAbnormalEventsCountRequest) GoString() string {
	return s.String()
}

func (s *GetAbnormalEventsCountRequest) SetCluster(v string) *GetAbnormalEventsCountRequest {
	s.Cluster = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetEnd(v float32) *GetAbnormalEventsCountRequest {
	s.End = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetInstance(v string) *GetAbnormalEventsCountRequest {
	s.Instance = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetNamespace(v string) *GetAbnormalEventsCountRequest {
	s.Namespace = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetPod(v string) *GetAbnormalEventsCountRequest {
	s.Pod = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetStart(v float32) *GetAbnormalEventsCountRequest {
	s.Start = &v
	return s
}

type GetAbnormalEventsCountResponseBody struct {
	// example:
	//
	// Success
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetAbnormalEventsCountResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// result: code=1 msg=(Request failed, status_code != 200)
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetAbnormalEventsCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAbnormalEventsCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAbnormalEventsCountResponseBody) SetCode(v string) *GetAbnormalEventsCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetAbnormalEventsCountResponseBody) SetData(v []*GetAbnormalEventsCountResponseBodyData) *GetAbnormalEventsCountResponseBody {
	s.Data = v
	return s
}

func (s *GetAbnormalEventsCountResponseBody) SetMessage(v string) *GetAbnormalEventsCountResponseBody {
	s.Message = &v
	return s
}

type GetAbnormalEventsCountResponseBodyData struct {
	// example:
	//
	// health
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAbnormalEventsCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAbnormalEventsCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAbnormalEventsCountResponseBodyData) SetType(v string) *GetAbnormalEventsCountResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetAbnormalEventsCountResponseBodyData) SetValue(v int64) *GetAbnormalEventsCountResponseBodyData {
	s.Value = &v
	return s
}

type GetAbnormalEventsCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAbnormalEventsCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAbnormalEventsCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAbnormalEventsCountResponse) GoString() string {
	return s.String()
}

func (s *GetAbnormalEventsCountResponse) SetHeaders(v map[string]*string) *GetAbnormalEventsCountResponse {
	s.Headers = v
	return s
}

func (s *GetAbnormalEventsCountResponse) SetStatusCode(v int32) *GetAbnormalEventsCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAbnormalEventsCountResponse) SetBody(v *GetAbnormalEventsCountResponseBody) *GetAbnormalEventsCountResponse {
	s.Body = v
	return s
}

type GetDiagnosisResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// quzuYl23
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetDiagnosisResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultRequest) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultRequest) SetTaskId(v string) *GetDiagnosisResultRequest {
	s.TaskId = &v
	return s
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
	return tea.Prettify(s)
}

func (s GetDiagnosisResultResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s GetDiagnosisResultResponseBodyData) GoString() string {
	return s.String()
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

type GetDiagnosisResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiagnosisResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDiagnosisResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultResponse) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponse) SetHeaders(v map[string]*string) *GetDiagnosisResultResponse {
	s.Headers = v
	return s
}

func (s *GetDiagnosisResultResponse) SetStatusCode(v int32) *GetDiagnosisResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiagnosisResultResponse) SetBody(v *GetDiagnosisResultResponseBody) *GetDiagnosisResultResponse {
	s.Body = v
	return s
}

type GetHealthPercentageRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetHealthPercentageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHealthPercentageRequest) GoString() string {
	return s.String()
}

func (s *GetHealthPercentageRequest) SetCluster(v string) *GetHealthPercentageRequest {
	s.Cluster = &v
	return s
}

func (s *GetHealthPercentageRequest) SetEnd(v float32) *GetHealthPercentageRequest {
	s.End = &v
	return s
}

func (s *GetHealthPercentageRequest) SetInstance(v string) *GetHealthPercentageRequest {
	s.Instance = &v
	return s
}

func (s *GetHealthPercentageRequest) SetStart(v float32) *GetHealthPercentageRequest {
	s.Start = &v
	return s
}

type GetHealthPercentageResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetHealthPercentageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetHealthPercentageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHealthPercentageResponseBody) GoString() string {
	return s.String()
}

func (s *GetHealthPercentageResponseBody) SetCode(v string) *GetHealthPercentageResponseBody {
	s.Code = &v
	return s
}

func (s *GetHealthPercentageResponseBody) SetData(v []*GetHealthPercentageResponseBodyData) *GetHealthPercentageResponseBody {
	s.Data = v
	return s
}

func (s *GetHealthPercentageResponseBody) SetMessage(v string) *GetHealthPercentageResponseBody {
	s.Message = &v
	return s
}

type GetHealthPercentageResponseBodyData struct {
	// example:
	//
	// health
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetHealthPercentageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHealthPercentageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHealthPercentageResponseBodyData) SetType(v string) *GetHealthPercentageResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetHealthPercentageResponseBodyData) SetValue(v int64) *GetHealthPercentageResponseBodyData {
	s.Value = &v
	return s
}

type GetHealthPercentageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHealthPercentageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHealthPercentageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHealthPercentageResponse) GoString() string {
	return s.String()
}

func (s *GetHealthPercentageResponse) SetHeaders(v map[string]*string) *GetHealthPercentageResponse {
	s.Headers = v
	return s
}

func (s *GetHealthPercentageResponse) SetStatusCode(v int32) *GetHealthPercentageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHealthPercentageResponse) SetBody(v *GetHealthPercentageResponseBody) *GetHealthPercentageResponse {
	s.Body = v
	return s
}

type InvokeDiagnosisRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cloud_assist
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "instance": "i-wz9gdv7qmdhusamc4dl01",
	//
	//     "uid": "xxxxxxxxxxxxxx",
	//
	//     "region": "cn-shenzhen"
	//
	// }
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// memgraph
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

func (s InvokeDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisRequest) SetChannel(v string) *InvokeDiagnosisRequest {
	s.Channel = &v
	return s
}

func (s *InvokeDiagnosisRequest) SetParams(v string) *InvokeDiagnosisRequest {
	s.Params = &v
	return s
}

func (s *InvokeDiagnosisRequest) SetServiceName(v string) *InvokeDiagnosisRequest {
	s.ServiceName = &v
	return s
}

type InvokeDiagnosisResponseBody struct {
	// example:
	//
	// Success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *InvokeDiagnosisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s InvokeDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisResponseBody) SetCode(v string) *InvokeDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *InvokeDiagnosisResponseBody) SetData(v *InvokeDiagnosisResponseBodyData) *InvokeDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *InvokeDiagnosisResponseBody) SetMessage(v string) *InvokeDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeDiagnosisResponseBody) SetRequestId(v string) *InvokeDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

type InvokeDiagnosisResponseBodyData struct {
	// example:
	//
	// ihqhAcrt
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s InvokeDiagnosisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InvokeDiagnosisResponseBodyData) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisResponseBodyData) SetTaskId(v string) *InvokeDiagnosisResponseBodyData {
	s.TaskId = &v
	return s
}

type InvokeDiagnosisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisResponse) SetHeaders(v map[string]*string) *InvokeDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *InvokeDiagnosisResponse) SetStatusCode(v int32) *InvokeDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeDiagnosisResponse) SetBody(v *InvokeDiagnosisResponseBody) *InvokeDiagnosisResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sysom"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 授权 SysOM 对某个机器进行诊断
//
// @param request - AuthDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthDiagnosisResponse
func (client *Client) AuthDiagnosisWithOptions(request *AuthDiagnosisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AuthDiagnosisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoCreateRole)) {
		body["autoCreateRole"] = request.AutoCreateRole
	}

	if !tea.BoolValue(util.IsUnset(request.AutoInstallAgent)) {
		body["autoInstallAgent"] = request.AutoInstallAgent
	}

	if !tea.BoolValue(util.IsUnset(request.Instances)) {
		body["instances"] = request.Instances
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthDiagnosis"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/diagnosis/auth"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授权 SysOM 对某个机器进行诊断
//
// @param request - AuthDiagnosisRequest
//
// @return AuthDiagnosisResponse
func (client *Client) AuthDiagnosis(request *AuthDiagnosisRequest) (_result *AuthDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AuthDiagnosisResponse{}
	_body, _err := client.AuthDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取copilot服务的返回结果
//
// @param request - GenerateCopilotResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCopilotResponseResponse
func (client *Client) GenerateCopilotResponseWithOptions(request *GenerateCopilotResponseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateCopilotResponseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LlmParamString)) {
		body["llmParamString"] = request.LlmParamString
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateCopilotResponse"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/copilot/generate_copilot_response"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateCopilotResponseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取copilot服务的返回结果
//
// @param request - GenerateCopilotResponseRequest
//
// @return GenerateCopilotResponseResponse
func (client *Client) GenerateCopilotResponse(request *GenerateCopilotResponseRequest) (_result *GenerateCopilotResponseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateCopilotResponseResponse{}
	_body, _err := client.GenerateCopilotResponseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节点/Pod不同等级异常事件的数量
//
// @param request - GetAbnormalEventsCountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAbnormalEventsCountResponse
func (client *Client) GetAbnormalEventsCountWithOptions(request *GetAbnormalEventsCountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAbnormalEventsCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cluster)) {
		query["cluster"] = request.Cluster
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Pod)) {
		query["pod"] = request.Pod
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAbnormalEventsCount"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/cluster_health/range/abnormaly_events_count"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAbnormalEventsCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节点/Pod不同等级异常事件的数量
//
// @param request - GetAbnormalEventsCountRequest
//
// @return GetAbnormalEventsCountResponse
func (client *Client) GetAbnormalEventsCount(request *GetAbnormalEventsCountRequest) (_result *GetAbnormalEventsCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAbnormalEventsCountResponse{}
	_body, _err := client.GetAbnormalEventsCountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取诊断结果
//
// @param request - GetDiagnosisResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDiagnosisResultResponse
func (client *Client) GetDiagnosisResultWithOptions(request *GetDiagnosisResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDiagnosisResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["task_id"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiagnosisResult"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/diagnosis/get_diagnosis_results"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiagnosisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取诊断结果
//
// @param request - GetDiagnosisResultRequest
//
// @return GetDiagnosisResultResponse
func (client *Client) GetDiagnosisResult(request *GetDiagnosisResultRequest) (_result *GetDiagnosisResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDiagnosisResultResponse{}
	_body, _err := client.GetDiagnosisResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一段时间的节点/pod健康度比例
//
// @param request - GetHealthPercentageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHealthPercentageResponse
func (client *Client) GetHealthPercentageWithOptions(request *GetHealthPercentageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHealthPercentageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cluster)) {
		query["cluster"] = request.Cluster
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHealthPercentage"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/cluster_health/range/health_percentage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHealthPercentageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一段时间的节点/pod健康度比例
//
// @param request - GetHealthPercentageRequest
//
// @return GetHealthPercentageResponse
func (client *Client) GetHealthPercentage(request *GetHealthPercentageRequest) (_result *GetHealthPercentageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHealthPercentageResponse{}
	_body, _err := client.GetHealthPercentageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发起诊断
//
// @param request - InvokeDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeDiagnosisResponse
func (client *Client) InvokeDiagnosisWithOptions(request *InvokeDiagnosisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InvokeDiagnosisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["service_name"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeDiagnosis"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/diagnosis/invoke_diagnosis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发起诊断
//
// @param request - InvokeDiagnosisRequest
//
// @return InvokeDiagnosisResponse
func (client *Client) InvokeDiagnosis(request *InvokeDiagnosisRequest) (_result *InvokeDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InvokeDiagnosisResponse{}
	_body, _err := client.InvokeDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
