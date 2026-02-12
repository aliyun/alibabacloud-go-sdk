// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceBaseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceBaseInfo(v *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) *OnsInstanceBaseInfoResponseBody
	GetInstanceBaseInfo() *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo
	SetRequestId(v string) *OnsInstanceBaseInfoResponseBody
	GetRequestId() *string
}

type OnsInstanceBaseInfoResponseBody struct {
	// The information about the instance.
	InstanceBaseInfo *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo `json:"InstanceBaseInfo,omitempty" xml:"InstanceBaseInfo,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 6CC46974-65E8-4C20-AB07-D20D102E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsInstanceBaseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *OnsInstanceBaseInfoResponseBody) GetInstanceBaseInfo() *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	return s.InstanceBaseInfo
}

func (s *OnsInstanceBaseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsInstanceBaseInfoResponseBody) SetInstanceBaseInfo(v *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) *OnsInstanceBaseInfoResponseBody {
	s.InstanceBaseInfo = v
	return s
}

func (s *OnsInstanceBaseInfoResponseBody) SetRequestId(v string) *OnsInstanceBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBody) Validate() error {
	if s.InstanceBaseInfo != nil {
		if err := s.InstanceBaseInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsInstanceBaseInfoResponseBodyInstanceBaseInfo struct {
	// The time when the instance was created. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1570701259403
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The endpoints used to access ApsaraMQ for RocketMQ over different protocols.
	Endpoints *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// Indicates whether the instance uses a namespace. Valid values:
	//
	// 	- **true**: The instance uses a separate namespace. The name of each resource must be unique in the instance. The names of resources in different instances can be the same.
	//
	// 	- **false**: The instance does not use a separate namespace. The name of each resource must be globally unique within the instance and across all instances.
	//
	// example:
	//
	// true
	IndependentNaming *bool `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_138015630679****_BAAy1Hac
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// The name must be 3 to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- **0**: The instance is being deployed. This value is valid only for Enterprise Platinum Edition instances.
	//
	// 	- **2**: The instance has overdue payments. This value is valid only for Standard Edition instances.
	//
	// 	- **5**: The instance is running. This value is valid for Standard Edition instances and Enterprise Platinum Edition instances.
	//
	// 	- **7**: The instance is being upgraded and is running. This value is valid only for Enterprise Platinum Edition instances.
	//
	// example:
	//
	// 5
	InstanceStatus *int32 `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The instance type. Valid values:
	//
	// 	- **1**: Standard Edition instances that use the pay-as-you-go billing method.
	//
	// 	- **2**: Enterprise Platinum Edition instances that use the subscription billing method.
	//
	// For information about the editions and specifications of ApsaraMQ for RocketMQ instances, see [Instance editions](https://help.aliyun.com/document_detail/185261.html).
	//
	// example:
	//
	// 2
	InstanceType *int32 `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum messaging transactions per second (TPS). Valid values: 5000, 10000, 20000, 50000, 100000, 200000, 300000, 500000, 800000, and 1000000.
	//
	// You can view the details about messaging TPS on the buy page of ApsaraMQ for RocketMQ.
	//
	// > This parameter is available only to the ApsaraMQ for RocketMQ Enterprise Platinum Edition instances.
	//
	// example:
	//
	// 10000
	MaxTps *int64 `json:"MaxTps,omitempty" xml:"MaxTps,omitempty"`
	// The time when the Enterprise Platinum Edition instance expires.
	//
	// example:
	//
	// 1603555200000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// ons-cn-m7r1r5f****
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 2
	SupportClassic *int32 `json:"SupportClassic,omitempty" xml:"SupportClassic,omitempty"`
	// The maximum number of topics that can be created on the instance. Valid values: 25, 50, 100, 300, and 500.
	//
	// > This parameter is available only to the ApsaraMQ for RocketMQ Enterprise Platinum Edition instances.
	//
	// example:
	//
	// 50
	TopicCapacity *int32 `json:"TopicCapacity,omitempty" xml:"TopicCapacity,omitempty"`
	// The commodity ID of the instance.
	//
	// example:
	//
	// ons-cn-m7r1r5f****
	SpInstanceId *string `json:"spInstanceId,omitempty" xml:"spInstanceId,omitempty"`
	// The commodity type of the instance.
	//
	// example:
	//
	// 1
	SpInstanceType *int32 `json:"spInstanceType,omitempty" xml:"spInstanceType,omitempty"`
}

func (s OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GoString() string {
	return s.String()
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetEndpoints() *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	return s.Endpoints
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetIndependentNaming() *bool {
	return s.IndependentNaming
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetInstanceName() *string {
	return s.InstanceName
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetInstanceStatus() *int32 {
	return s.InstanceStatus
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetInstanceType() *int32 {
	return s.InstanceType
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetMaxTps() *int64 {
	return s.MaxTps
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetRemark() *string {
	return s.Remark
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetSupportClassic() *int32 {
	return s.SupportClassic
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetTopicCapacity() *int32 {
	return s.TopicCapacity
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetSpInstanceId() *string {
	return s.SpInstanceId
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) GetSpInstanceType() *int32 {
	return s.SpInstanceType
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetCreateTime(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.CreateTime = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetEndpoints(v *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.Endpoints = v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetIndependentNaming(v bool) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.IndependentNaming = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetInstanceId(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.InstanceId = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetInstanceName(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.InstanceName = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetInstanceStatus(v int32) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.InstanceStatus = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetInstanceType(v int32) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.InstanceType = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetMaxTps(v int64) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.MaxTps = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetReleaseTime(v int64) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.ReleaseTime = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetRemark(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.Remark = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetSupportClassic(v int32) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.SupportClassic = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetTopicCapacity(v int32) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.TopicCapacity = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetSpInstanceId(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.SpInstanceId = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) SetSpInstanceType(v int32) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo {
	s.SpInstanceType = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfo) Validate() error {
	if s.Endpoints != nil {
		if err := s.Endpoints.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints struct {
	// The private HTTP endpoint of the instance.
	//
	// example:
	//
	// http://138015630679****.mqrest.cn-chengdu-internal.aliyuncs.com
	HttpInternalEndpoint *string `json:"HttpInternalEndpoint,omitempty" xml:"HttpInternalEndpoint,omitempty"`
	// The public HTTP endpoint of the instance.
	//
	// example:
	//
	// http://138015630679****.mqrest.cn-chengdu.aliyuncs.com
	HttpInternetEndpoint *string `json:"HttpInternetEndpoint,omitempty" xml:"HttpInternetEndpoint,omitempty"`
	// The public HTTPS endpoint of the instance.
	//
	// example:
	//
	// https://138015630679****.mqrest.cn-chengdu.aliyuncs.com
	HttpInternetSecureEndpoint *string `json:"HttpInternetSecureEndpoint,omitempty" xml:"HttpInternetSecureEndpoint,omitempty"`
	// The private TCP endpoint of the instance.
	//
	// example:
	//
	// http://MQ_INST_138015630679****_BAAy1Hac.cn-chengdu.mq-internal.aliyuncs.com:8080
	TcpEndpoint *string `json:"TcpEndpoint,omitempty" xml:"TcpEndpoint,omitempty"`
	// The public TCP endpoint of the instance.
	//
	// 	- Only instances that are deployed in the China (Chengdu), China (Qingdao), or China (Shenzhen) region can be accessed by using public TCP endpoints.
	//
	// 	- Before you use a public TCP endpoint, make sure that your client SDK meets the following requirements:
	//
	//     	- TCP client SDK for Java: V2.0.0.Final or later For more information, see [Release notes for the SDK for Java](https://help.aliyun.com/document_detail/325569.html).
	//
	//     	- TCP client SDK for C++: V3.0.0 or later For more information, see [Release notes for the SDK for C++](https://help.aliyun.com/document_detail/325570.html).
	//
	// 	- You are charged for Internet traffic when you use a public TCP endpoint. For more information, see [Internet traffic fee](https://help.aliyun.com/document_detail/325571.html).
	//
	// example:
	//
	// http://MQ_INST_138015630679****_BAAy1Hac.mq.cn-chengdu.aliyuncs.com:80
	TcpInternetEndpoint *string `json:"TcpInternetEndpoint,omitempty" xml:"TcpInternetEndpoint,omitempty"`
}

func (s OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) GoString() string {
	return s.String()
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) GetHttpInternalEndpoint() *string {
	return s.HttpInternalEndpoint
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) GetHttpInternetEndpoint() *string {
	return s.HttpInternetEndpoint
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) GetHttpInternetSecureEndpoint() *string {
	return s.HttpInternetSecureEndpoint
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) GetTcpEndpoint() *string {
	return s.TcpEndpoint
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) GetTcpInternetEndpoint() *string {
	return s.TcpInternetEndpoint
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetHttpInternalEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.HttpInternalEndpoint = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetHttpInternetEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.HttpInternetEndpoint = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetHttpInternetSecureEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.HttpInternetSecureEndpoint = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetTcpEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.TcpEndpoint = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) SetTcpInternetEndpoint(v string) *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints {
	s.TcpInternetEndpoint = &v
	return s
}

func (s *OnsInstanceBaseInfoResponseBodyInstanceBaseInfoEndpoints) Validate() error {
	return dara.Validate(s)
}
