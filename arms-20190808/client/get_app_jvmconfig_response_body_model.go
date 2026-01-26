// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppJVMConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAppJVMConfigResponseBody
	GetCode() *int32
	SetJvmInfoList(v []*GetAppJVMConfigResponseBodyJvmInfoList) *GetAppJVMConfigResponseBody
	GetJvmInfoList() []*GetAppJVMConfigResponseBodyJvmInfoList
	SetMessage(v string) *GetAppJVMConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAppJVMConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAppJVMConfigResponseBody
	GetSuccess() *bool
}

type GetAppJVMConfigResponseBody struct {
	// The response code. Valid values: 2XX: The request is successful. 3XX: A redirection message is returned. 4XX: The request is invalid. 5XX: A server error occurs.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of JVM information.
	JvmInfoList []*GetAppJVMConfigResponseBodyJvmInfoList `json:"JvmInfoList,omitempty" xml:"JvmInfoList,omitempty" type:"Repeated"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAppJVMConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppJVMConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppJVMConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAppJVMConfigResponseBody) GetJvmInfoList() []*GetAppJVMConfigResponseBodyJvmInfoList {
	return s.JvmInfoList
}

func (s *GetAppJVMConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAppJVMConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppJVMConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAppJVMConfigResponseBody) SetCode(v int32) *GetAppJVMConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppJVMConfigResponseBody) SetJvmInfoList(v []*GetAppJVMConfigResponseBodyJvmInfoList) *GetAppJVMConfigResponseBody {
	s.JvmInfoList = v
	return s
}

func (s *GetAppJVMConfigResponseBody) SetMessage(v string) *GetAppJVMConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppJVMConfigResponseBody) SetRequestId(v string) *GetAppJVMConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppJVMConfigResponseBody) SetSuccess(v bool) *GetAppJVMConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppJVMConfigResponseBody) Validate() error {
	if s.JvmInfoList != nil {
		for _, item := range s.JvmInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAppJVMConfigResponseBodyJvmInfoList struct {
	// The version of the agent.
	//
	// example:
	//
	// 1.7.0-SNAPSHOT_3.0.3_3756244
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The hostname.
	//
	// example:
	//
	// host_name
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 47.91.59.244
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The application ID.
	//
	// example:
	//
	// dsv9zcel92@1455182510c5369
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The process ID.
	//
	// example:
	//
	// 1
	ProcId *string `json:"ProcId,omitempty" xml:"ProcId,omitempty"`
	// The VM parameters.
	//
	// example:
	//
	// [-javaagent:/home/admin/.opt/ArmsAgent/arms-bootstrap-1.7.0-SNAPSHOT.jar, -Doneagent.plugin.arms-agent.enabled=true, -Darms.licenseKey=[******], -Darms.agent.env=K8s, -Darms.agent.podinfo.path=/etc/podinfo, -Darms.appName=productservice, -Doneagent.region=cn-hangzhou, -Dproject.name=Product]
	VmArgs *string `json:"VmArgs,omitempty" xml:"VmArgs,omitempty"`
}

func (s GetAppJVMConfigResponseBodyJvmInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetAppJVMConfigResponseBodyJvmInfoList) GoString() string {
	return s.String()
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) GetHostName() *string {
	return s.HostName
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) GetIp() *string {
	return s.Ip
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) GetPid() *string {
	return s.Pid
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) GetProcId() *string {
	return s.ProcId
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) GetVmArgs() *string {
	return s.VmArgs
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) SetAgentVersion(v string) *GetAppJVMConfigResponseBodyJvmInfoList {
	s.AgentVersion = &v
	return s
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) SetHostName(v string) *GetAppJVMConfigResponseBodyJvmInfoList {
	s.HostName = &v
	return s
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) SetIp(v string) *GetAppJVMConfigResponseBodyJvmInfoList {
	s.Ip = &v
	return s
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) SetPid(v string) *GetAppJVMConfigResponseBodyJvmInfoList {
	s.Pid = &v
	return s
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) SetProcId(v string) *GetAppJVMConfigResponseBodyJvmInfoList {
	s.ProcId = &v
	return s
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) SetVmArgs(v string) *GetAppJVMConfigResponseBodyJvmInfoList {
	s.VmArgs = &v
	return s
}

func (s *GetAppJVMConfigResponseBodyJvmInfoList) Validate() error {
	return dara.Validate(s)
}
