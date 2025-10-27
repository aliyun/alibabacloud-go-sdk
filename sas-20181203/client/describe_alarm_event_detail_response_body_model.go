// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmEventDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeAlarmEventDetailResponseBodyData) *DescribeAlarmEventDetailResponseBody
	GetData() *DescribeAlarmEventDetailResponseBodyData
	SetRequestId(v string) *DescribeAlarmEventDetailResponseBody
	GetRequestId() *string
}

type DescribeAlarmEventDetailResponseBody struct {
	// The details of the alert event.
	Data *DescribeAlarmEventDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7EA50837-2F0B-5BCC-AB61-4968D88D75AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAlarmEventDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventDetailResponseBody) GetData() *DescribeAlarmEventDetailResponseBodyData {
	return s.Data
}

func (s *DescribeAlarmEventDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlarmEventDetailResponseBody) SetData(v *DescribeAlarmEventDetailResponseBodyData) *DescribeAlarmEventDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlarmEventDetailResponseBody) SetRequestId(v string) *DescribeAlarmEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAlarmEventDetailResponseBodyData struct {
	// The name of the alert event.
	//
	// example:
	//
	// Login with unusual location
	AlarmEventAliasName *string `json:"AlarmEventAliasName,omitempty" xml:"AlarmEventAliasName,omitempty"`
	// The description of the alert event.
	//
	// example:
	//
	// The detection model finds that self-mutation is running on your server. A self-mutation Trojan is a Trojan horse program with self-mutation function. It will change its hash or copy a large number of itself to different paths, and run in the background to avoid cleaning.
	AlarmEventDesc *string `json:"AlarmEventDesc,omitempty" xml:"AlarmEventDesc,omitempty"`
	// The unique identifier of the alert event.
	//
	// > To query the details of an alert event, you must provide the unique identifier of the alert event. You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to obtain the identifier.
	//
	// example:
	//
	// 9f62555666f177aa84ee1eaf465a****
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// The name of the container application.
	//
	// example:
	//
	// app:msdp-uat-service
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Indicates whether the online handling of the alert event is supported. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	CanBeDealOnLine *bool `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty"`
	// Indicates whether you can cancel marking the alert event as a false positive. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	CanCancelFault *bool `json:"CanCancelFault,omitempty" xml:"CanCancelFault,omitempty"`
	// An array consisting of the cause of the alert event, which can be used to trace the alert event.
	CauseDetails []*DescribeAlarmEventDetailResponseBodyDataCauseDetails `json:"CauseDetails,omitempty" xml:"CauseDetails,omitempty" type:"Repeated"`
	// Indicates whether the Safeguard Mode For Major Activities mode is enabled.
	//
	// example:
	//
	// true
	ContainHwMode *bool `json:"ContainHwMode,omitempty" xml:"ContainHwMode,omitempty"`
	// The ID of the container application.
	//
	// example:
	//
	// container_1606995441910_394868_01_000***
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The ID of the image to which the container belongs.
	//
	// example:
	//
	// cadb7a725641
	ContainerImageId *string `json:"ContainerImageId,omitempty" xml:"ContainerImageId,omitempty"`
	// The name of the image to which the container belongs.
	//
	// example:
	//
	// jenkins/jenkins:latest
	ContainerImageName *string `json:"ContainerImageName,omitempty" xml:"ContainerImageName,omitempty"`
	// The data source of the alert event.
	//
	// example:
	//
	// aegis_***
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The timestamp when the alert event ends. Unit: milliseconds.
	//
	// example:
	//
	// 1542366542000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// i-wz92q7m5hsbgfhdss***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the associated instance.
	//
	// example:
	//
	// 172.16.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the associated instance.
	//
	// example:
	//
	// 172.25.30.**
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The ID of the Kubernetes cluster.
	//
	// example:
	//
	// c562cf0d68e9749ee9fe544a7ab2f****
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The name of the Kubernetes cluster.
	//
	// example:
	//
	// TestK8sCluser
	K8sClusterName *string `json:"K8sClusterName,omitempty" xml:"K8sClusterName,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// example:
	//
	// sit-saic-trip
	K8sNamespace *string `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	// The ID of the Kubernetes cluster node.
	//
	// example:
	//
	// i-bp14a1ay8e0aa9t0l***
	K8sNodeId *string `json:"K8sNodeId,omitempty" xml:"K8sNodeId,omitempty"`
	// The name of the Kubernetes cluster node.
	//
	// example:
	//
	// cn-hangzhou.10.188.139.**
	K8sNodeName *string `json:"K8sNodeName,omitempty" xml:"K8sNodeName,omitempty"`
	// The name of the Kubernetes pod.
	//
	// example:
	//
	// myapp-pod
	K8sPodName *string `json:"K8sPodName,omitempty" xml:"K8sPodName,omitempty"`
	// The severity of the alert event. Valid values:
	//
	// 	- **serious**
	//
	// 	- **suspicious**
	//
	// 	- **remind**
	//
	// example:
	//
	// serious
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The solution to the alert event.
	//
	// example:
	//
	// An invalid logon source IP has been detected. If you recognize this logon attempt, we recommend that you add the current logon source IP to the valid logon source IP list to avoid future alerts. If you do not recognize this logon attempt, we recommend that you modify the password.
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// The timestamp when the alert event starts. Unit: milliseconds.
	//
	// example:
	//
	// 1542378601000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The alert type of the alert event. Valid values:
	//
	// 	- Suspicious process
	//
	// 	- Webshell
	//
	// 	- Unusual logon
	//
	// 	- Exception
	//
	// 	- Sensitive file tampering
	//
	// 	- Malicious process (cloud threat detection)
	//
	// 	- Suspicious network connection
	//
	// 	- Other
	//
	// 	- Abnormal account
	//
	// 	- Application intrusion event
	//
	// 	- Cloud threat detection
	//
	// 	- Precise defense
	//
	// 	- Application whitelist
	//
	// 	- Persistent webshell
	//
	// 	- Web application threat detection
	//
	// 	- Malicious script
	//
	// 	- Threat intelligence
	//
	// 	- Malicious network activity
	//
	// 	- Cluster exception
	//
	// 	- Webshell (on-premises threat detection)
	//
	// 	- Vulnerability exploitation
	//
	// 	- Malicious process (on-premises threat detection)
	//
	// 	- Trusted exception
	//
	// example:
	//
	// Webshell
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The instance UUID of the asset.
	//
	// example:
	//
	// 6690a46c-0edb-4663-a641-3629d1a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeAlarmEventDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmEventDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetAlarmEventAliasName() *string {
	return s.AlarmEventAliasName
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetAlarmEventDesc() *string {
	return s.AlarmEventDesc
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetCanBeDealOnLine() *bool {
	return s.CanBeDealOnLine
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetCanCancelFault() *bool {
	return s.CanCancelFault
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetCauseDetails() []*DescribeAlarmEventDetailResponseBodyDataCauseDetails {
	return s.CauseDetails
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetContainHwMode() *bool {
	return s.ContainHwMode
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetContainerImageId() *string {
	return s.ContainerImageId
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetContainerImageName() *string {
	return s.ContainerImageName
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetDataSource() *string {
	return s.DataSource
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetK8sClusterId() *string {
	return s.K8sClusterId
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetK8sClusterName() *string {
	return s.K8sClusterName
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetK8sNamespace() *string {
	return s.K8sNamespace
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetK8sNodeId() *string {
	return s.K8sNodeId
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetK8sNodeName() *string {
	return s.K8sNodeName
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetK8sPodName() *string {
	return s.K8sPodName
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetLevel() *string {
	return s.Level
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetSolution() *string {
	return s.Solution
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DescribeAlarmEventDetailResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetAlarmEventAliasName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.AlarmEventAliasName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetAlarmEventDesc(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.AlarmEventDesc = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetAlarmUniqueInfo(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetAppName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetCanBeDealOnLine(v bool) *DescribeAlarmEventDetailResponseBodyData {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetCanCancelFault(v bool) *DescribeAlarmEventDetailResponseBodyData {
	s.CanCancelFault = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetCauseDetails(v []*DescribeAlarmEventDetailResponseBodyDataCauseDetails) *DescribeAlarmEventDetailResponseBodyData {
	s.CauseDetails = v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetContainHwMode(v bool) *DescribeAlarmEventDetailResponseBodyData {
	s.ContainHwMode = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetContainerId(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.ContainerId = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetContainerImageId(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.ContainerImageId = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetContainerImageName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.ContainerImageName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetDataSource(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.DataSource = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetEndTime(v int64) *DescribeAlarmEventDetailResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetInstanceName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetInternetIp(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetIntranetIp(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetK8sClusterId(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetK8sClusterName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.K8sClusterName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetK8sNamespace(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.K8sNamespace = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetK8sNodeId(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.K8sNodeId = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetK8sNodeName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.K8sNodeName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetK8sPodName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.K8sPodName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetLevel(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.Level = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetSolution(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.Solution = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetStartTime(v int64) *DescribeAlarmEventDetailResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetType(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetUuid(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) Validate() error {
	if s.CauseDetails != nil {
		for _, item := range s.CauseDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlarmEventDetailResponseBodyDataCauseDetails struct {
	// The key that is used to trace the alert event.
	//
	// example:
	//
	// 842e314e69b1a2c45d5c1a2f88a16***
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value that is used to trace the alert event.
	Value []*DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeAlarmEventDetailResponseBodyDataCauseDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmEventDetailResponseBodyDataCauseDetails) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetails) GetKey() *string {
	return s.Key
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetails) GetValue() []*DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue {
	return s.Value
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetails) SetKey(v string) *DescribeAlarmEventDetailResponseBodyDataCauseDetails {
	s.Key = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetails) SetValue(v []*DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) *DescribeAlarmEventDetailResponseBodyDataCauseDetails {
	s.Value = v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetails) Validate() error {
	if s.Value != nil {
		for _, item := range s.Value {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue struct {
	// The name of the field that displays the tracing information.
	//
	// example:
	//
	// sshd
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the field that displays the tracing information. Valid values:
	//
	// 	- **text**
	//
	// 	- **html**
	//
	// example:
	//
	// html
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the field that displays the tracing information.
	//
	// example:
	//
	// <p>under a certain small probability, yundun may mistakenly judge the repeated attempts caused by the administrator forgetting or entering the wrong password as successful blasting. Please check according to the account number and time shown in the alarm details. Once it is confirmed that it is not the initiative of the administrator, it is recommended to immediately block the IP, and you can open it at the same time<a href="https://yundun.console.aliyun.com/?p=pam">PAM</a>, hosting host login password, improving remote connection efficiency and security control ability, and according to<a href="https://click.aliyun.com/m/1000226086/">best practice of ECS account security protection</a>Modify login password and convergence asset.</p>↵
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) GetName() *string {
	return s.Name
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) GetType() *string {
	return s.Type
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) GetValue() *string {
	return s.Value
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) SetName(v string) *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue {
	s.Name = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) SetType(v string) *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue {
	s.Type = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) SetValue(v string) *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue {
	s.Value = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) Validate() error {
	return dara.Validate(s)
}
