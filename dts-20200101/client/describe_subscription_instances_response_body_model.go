// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeSubscriptionInstancesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSubscriptionInstancesResponseBody
	GetErrMessage() *string
	SetPageNumber(v int32) *DescribeSubscriptionInstancesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSubscriptionInstancesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeSubscriptionInstancesResponseBody
	GetRequestId() *string
	SetSubscriptionInstances(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstances) *DescribeSubscriptionInstancesResponseBody
	GetSubscriptionInstances() *DescribeSubscriptionInstancesResponseBodySubscriptionInstances
	SetSuccess(v string) *DescribeSubscriptionInstancesResponseBody
	GetSuccess() *string
	SetTotalRecordCount(v int64) *DescribeSubscriptionInstancesResponseBody
	GetTotalRecordCount() *int64
}

type DescribeSubscriptionInstancesResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries that can be displayed on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FC3BAAF2-74E3-4471-8EB5-96202D6A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of change tracking instances and the details of each instance.
	SubscriptionInstances *DescribeSubscriptionInstancesResponseBodySubscriptionInstances `json:"SubscriptionInstances,omitempty" xml:"SubscriptionInstances,omitempty" type:"Struct"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of change tracking instances that belong to your Alibaba Cloud account.
	//
	// example:
	//
	// 1
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSubscriptionInstancesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSubscriptionInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSubscriptionInstancesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSubscriptionInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSubscriptionInstancesResponseBody) GetSubscriptionInstances() *DescribeSubscriptionInstancesResponseBodySubscriptionInstances {
	return s.SubscriptionInstances
}

func (s *DescribeSubscriptionInstancesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSubscriptionInstancesResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeSubscriptionInstancesResponseBody) SetErrCode(v string) *DescribeSubscriptionInstancesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetErrMessage(v string) *DescribeSubscriptionInstancesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetPageNumber(v int32) *DescribeSubscriptionInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetPageRecordCount(v int32) *DescribeSubscriptionInstancesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetRequestId(v string) *DescribeSubscriptionInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetSubscriptionInstances(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstances) *DescribeSubscriptionInstancesResponseBody {
	s.SubscriptionInstances = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetSuccess(v string) *DescribeSubscriptionInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetTotalRecordCount(v int64) *DescribeSubscriptionInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) Validate() error {
	if s.SubscriptionInstances != nil {
		if err := s.SubscriptionInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstances struct {
	SubscriptionInstance []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance `json:"SubscriptionInstance,omitempty" xml:"SubscriptionInstance,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstances) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstances) GetSubscriptionInstance() []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	return s.SubscriptionInstance
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstances) SetSubscriptionInstance(v []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) *DescribeSubscriptionInstancesResponseBodySubscriptionInstances {
	s.SubscriptionInstance = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstances) Validate() error {
	if s.SubscriptionInstance != nil {
		for _, item := range s.SubscriptionInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance struct {
	// The start of the time range for change tracking. The time is displayed in the *yyyy-MM-dd*T*HH:mm:ss*Z format in UTC.
	//
	// example:
	//
	// 2021-06-21T08:25:43Z
	BeginTimestamp *string `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	// The consumption checkpoint of the change tracking instance. The time is displayed in the *yyyy-MM-dd*T*HH:mm:ss*Z format in UTC.
	//
	// example:
	//
	// 2021-06-21T09:30:31Z
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	// The downstream client information, in the following format: \\<IP address of the downstream client>:\\<Random ID generated by DTS>.
	//
	// example:
	//
	// 114.***.***.**:dts********
	ConsumptionClient *string `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	// The end of the time range for change tracking. The time is displayed in the *yyyy-MM-dd*T*HH:mm:ss*Z format in UTC.
	//
	// example:
	//
	// 2021-06-21T10:17:21Z
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The error message returned if change tracking failed.
	//
	// example:
	//
	// xxxxxxxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the change tracking instance was created. The time is displayed in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2021-06-21T02:48:20Z
	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" xml:"InstanceCreateTime,omitempty"`
	// The time when the change tracking task was created. The time is displayed in the *yyyy-MM-dd*T*HH:mm:ss*Z format in UTC.
	//
	// example:
	//
	// 2021-06-21T02:48:20Z
	JobCreateTime *string `json:"JobCreateTime,omitempty" xml:"JobCreateTime,omitempty"`
	// The billing method of the change tracking instance. Valid values:
	//
	// 	- **PrePaid**: subscription
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The status of the change tracking task. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is being prechecked.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **Starting**: The task is being started.
	//
	// 	- **Normal**: The task is running as expected.
	//
	// 	- **Abnormal**: The task is not running as expected.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The topic of the change tracking instance.
	//
	// >  This parameter is returned only if your change tracking instances are of the new version and you have called the [CreateConsumerGroup](https://help.aliyun.com/document_detail/122863.html) operation to create a consumer group.
	//
	// example:
	//
	// cn_hangzhou_rm_bp1162kryivb8****_dtstest_version2
	SubscribeTopic *string `json:"SubscribeTopic,omitempty" xml:"SubscribeTopic,omitempty"`
	// The types of operations that are tracked by the task.
	SubscriptionDataType *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	// The endpoint of the change tracking instance.
	SubscriptionHost *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost `json:"SubscriptionHost,omitempty" xml:"SubscriptionHost,omitempty" type:"Struct"`
	// The ID of the change tracking instance.
	//
	// example:
	//
	// rm-bp1162kryivb8****
	SubscriptionInstanceID *string `json:"SubscriptionInstanceID,omitempty" xml:"SubscriptionInstanceID,omitempty"`
	// The name of the change tracking instance.
	//
	// example:
	//
	// dtstest
	SubscriptionInstanceName *string `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	// The objects for change tracking.
	SubscriptionObject *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty" type:"Struct"`
	// The collection of tags.
	Tags *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetBeginTimestamp() *string {
	return s.BeginTimestamp
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetConsumptionCheckpoint() *string {
	return s.ConsumptionCheckpoint
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetConsumptionClient() *string {
	return s.ConsumptionClient
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetInstanceCreateTime() *string {
	return s.InstanceCreateTime
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetJobCreateTime() *string {
	return s.JobCreateTime
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetSourceEndpoint() *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetSubscribeTopic() *string {
	return s.SubscribeTopic
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetSubscriptionDataType() *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType {
	return s.SubscriptionDataType
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetSubscriptionHost() *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost {
	return s.SubscriptionHost
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetSubscriptionInstanceID() *string {
	return s.SubscriptionInstanceID
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetSubscriptionInstanceName() *string {
	return s.SubscriptionInstanceName
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetSubscriptionObject() *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject {
	return s.SubscriptionObject
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GetTags() *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags {
	return s.Tags
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetBeginTimestamp(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetConsumptionCheckpoint(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetConsumptionClient(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetEndTimestamp(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetErrorMessage(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetInstanceCreateTime(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.InstanceCreateTime = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetJobCreateTime(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.JobCreateTime = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetPayType(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.PayType = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSourceEndpoint(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetStatus(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.Status = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscribeTopic(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscribeTopic = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionDataType(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionDataType = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionHost(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionHost = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionInstanceID(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionInstanceID = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionObject(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionObject = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetTags(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.Tags = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) Validate() error {
	if s.SourceEndpoint != nil {
		if err := s.SourceEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.SubscriptionDataType != nil {
		if err := s.SubscriptionDataType.Validate(); err != nil {
			return err
		}
	}
	if s.SubscriptionHost != nil {
		if err := s.SubscriptionHost.Validate(); err != nil {
			return err
		}
	}
	if s.SubscriptionObject != nil {
		if err := s.SubscriptionObject.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint struct {
	// The ID of the source instance.
	//
	// example:
	//
	// rm-bp1162kryivb8****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the source instance.
	//
	// >  This parameter is returned only for change tracking instances of the new version.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) SetInstanceID(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) SetInstanceType(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType struct {
	// Indicates whether data definition language (DDL) operations are tracked. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	DDL *bool `json:"DDL,omitempty" xml:"DDL,omitempty"`
	// Indicates whether data manipulation language (DML) operations are tracked. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	DML *bool `json:"DML,omitempty" xml:"DML,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) GetDDL() *bool {
	return s.DDL
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) GetDML() *bool {
	return s.DML
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) SetDDL(v bool) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType {
	s.DDL = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) SetDML(v bool) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType {
	s.DML = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) Validate() error {
	return dara.Validate(s)
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost struct {
	// The private endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****-internal.aliyuncs.com:18002
	PrivateHost *string `json:"PrivateHost,omitempty" xml:"PrivateHost,omitempty"`
	// The public endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****.aliyuncs.com:18001
	PublicHost *string `json:"PublicHost,omitempty" xml:"PublicHost,omitempty"`
	// The virtual private cloud (VPC) endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****-vpc.aliyuncs.com:18003
	VPCHost *string `json:"VPCHost,omitempty" xml:"VPCHost,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) GetPrivateHost() *string {
	return s.PrivateHost
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) GetPublicHost() *string {
	return s.PublicHost
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) GetVPCHost() *string {
	return s.VPCHost
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) SetPrivateHost(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost {
	s.PrivateHost = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) SetPublicHost(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost {
	s.PublicHost = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) SetVPCHost(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost {
	s.VPCHost = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) Validate() error {
	return dara.Validate(s)
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject struct {
	SynchronousObject []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject `json:"SynchronousObject,omitempty" xml:"SynchronousObject,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject) GetSynchronousObject() []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject {
	return s.SynchronousObject
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject) SetSynchronousObject(v []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject {
	s.SynchronousObject = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject) Validate() error {
	if s.SynchronousObject != nil {
		for _, item := range s.SynchronousObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject struct {
	// The name of the database to which the object belongs.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The table name.
	TableList *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
	// Indicates whether the data of an entire database is tracked. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	WholeDatabase *string `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) GetTableList() *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList {
	return s.TableList
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) GetWholeDatabase() *string {
	return s.WholeDatabase
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) SetDatabaseName(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject {
	s.DatabaseName = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) SetTableList(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject {
	s.TableList = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) SetWholeDatabase(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject {
	s.WholeDatabase = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) Validate() error {
	if s.TableList != nil {
		if err := s.TableList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList) GetTable() []*string {
	return s.Table
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList) SetTable(v []*string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList {
	s.Table = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList) Validate() error {
	return dara.Validate(s)
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags struct {
	Tag []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags) GetTag() []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag {
	return s.Tag
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags) SetTag(v []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags {
	s.Tag = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that corresponds to the tag key.
	//
	// example:
	//
	// testvalue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) SetKey(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) SetValue(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) Validate() error {
	return dara.Validate(s)
}
