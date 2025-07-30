// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckJobs(v []*DescribeCheckJobsResponseBodyCheckJobs) *DescribeCheckJobsResponseBody
	GetCheckJobs() []*DescribeCheckJobsResponseBodyCheckJobs
	SetDynamicCode(v string) *DescribeCheckJobsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeCheckJobsResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeCheckJobsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeCheckJobsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeCheckJobsResponseBody
	GetHttpStatusCode() *int32
	SetPageNumber(v int32) *DescribeCheckJobsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int64) *DescribeCheckJobsResponseBody
	GetPageRecordCount() *int64
	SetRequestId(v string) *DescribeCheckJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCheckJobsResponseBody
	GetSuccess() *bool
	SetTotalRecordCount(v int64) *DescribeCheckJobsResponseBody
	GetTotalRecordCount() *int64
}

type DescribeCheckJobsResponseBody struct {
	// Item information check.
	CheckJobs []*DescribeCheckJobsResponseBodyCheckJobs `json:"CheckJobs,omitempty" xml:"CheckJobs,omitempty" type:"Repeated"`
	// Dynamic error code, this parameter will be deprecated soon.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// Dynamic error message, used to replace the **%s*	- in the **errmessage*	- return parameter. If **errmessage*	- returns **thevalueofinputparameter%sisnotvalid**, and **dynamicmessage*	- returns *[1,2,3]*, it indicates that the request parameter **dtsjobid*	- is invalid.
	//
	// example:
	//
	// present environment is not support,so skip.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// Error code returned when the call fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// Error message returned when the call fails.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of records that can be displayed on the current page.
	//
	// example:
	//
	// 20
	PageRecordCount *int64 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// Request ID.
	//
	// example:
	//
	// FC1D920B-AB89-52A9-AA5F-AA724C4205E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 100
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeCheckJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckJobsResponseBody) GetCheckJobs() []*DescribeCheckJobsResponseBodyCheckJobs {
	return s.CheckJobs
}

func (s *DescribeCheckJobsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeCheckJobsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeCheckJobsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeCheckJobsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeCheckJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeCheckJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCheckJobsResponseBody) GetPageRecordCount() *int64 {
	return s.PageRecordCount
}

func (s *DescribeCheckJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCheckJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCheckJobsResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeCheckJobsResponseBody) SetCheckJobs(v []*DescribeCheckJobsResponseBodyCheckJobs) *DescribeCheckJobsResponseBody {
	s.CheckJobs = v
	return s
}

func (s *DescribeCheckJobsResponseBody) SetDynamicCode(v string) *DescribeCheckJobsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeCheckJobsResponseBody) SetDynamicMessage(v string) *DescribeCheckJobsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeCheckJobsResponseBody) SetErrCode(v string) *DescribeCheckJobsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeCheckJobsResponseBody) SetErrMessage(v string) *DescribeCheckJobsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeCheckJobsResponseBody) SetHttpStatusCode(v int32) *DescribeCheckJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeCheckJobsResponseBody) SetPageNumber(v int32) *DescribeCheckJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCheckJobsResponseBody) SetPageRecordCount(v int64) *DescribeCheckJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeCheckJobsResponseBody) SetRequestId(v string) *DescribeCheckJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckJobsResponseBody) SetSuccess(v bool) *DescribeCheckJobsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCheckJobsResponseBody) SetTotalRecordCount(v int64) *DescribeCheckJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeCheckJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCheckJobsResponseBodyCheckJobs struct {
	// Billing type, return values: - **POSTPAY**: Pay-as-you-go (postpaid). - **PREPAY**: Subscription (prepaid).
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// checkpoint
	//
	// example:
	//
	// 1111****
	CheckPoint *int64 `json:"CheckPoint,omitempty" xml:"CheckPoint,omitempty"`
	// Data validation method, with values:
	//
	// - **1**: Full validation. - **2**: Incremental validation.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// Number of rows with data inconsistency
	//
	// example:
	//
	// 0
	DiffCount *int64 `json:"DiffCount,omitempty" xml:"DiffCount,omitempty"`
	// Synchronization initialization progress, in percentage.
	//
	// example:
	//
	// 1
	DiffSum *int64 `json:"DiffSum,omitempty" xml:"DiffSum,omitempty"`
	// Migration, synchronization, or subscription instance ID.
	//
	// example:
	//
	// dtsz8tc99sy2158b36
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// Migration, synchronization, or subscription task ID.
	//
	// example:
	//
	// n08o6si4q338b1x
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The number of data rows in the table that have completed validation.
	//
	// example:
	//
	// 15094
	FinishCount *int64 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// DTS task ID. In most cases, there is no need to set this parameter.
	//
	// example:
	//
	// c3d12dii27t632g
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Instance specification.
	//
	// example:
	//
	// SMALL
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// Check the name of the task.
	//
	// example:
	//
	// dtstest
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// Task ID.
	//
	// example:
	//
	// as0e1ks426bq3z0
	JobStepId *string `json:"JobStepId,omitempty" xml:"JobStepId,omitempty"`
	// This parameter will be deprecated.
	//
	// example:
	//
	// ****
	ParentJobType *string `json:"ParentJobType,omitempty" xml:"ParentJobType,omitempty"`
	// Region ID to which it belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Check result, return values: -**0**: Check passed -**1**: Check failed
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Verify the total number of rows in the data.
	//
	// example:
	//
	// 159
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCheckJobsResponseBodyCheckJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckJobsResponseBodyCheckJobs) GoString() string {
	return s.String()
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetCheckPoint() *int64 {
	return s.CheckPoint
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetCheckType() *int32 {
	return s.CheckType
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetDiffCount() *int64 {
	return s.DiffCount
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetDiffSum() *int64 {
	return s.DiffSum
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetFinishCount() *int64 {
	return s.FinishCount
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetJobName() *string {
	return s.JobName
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetJobStepId() *string {
	return s.JobStepId
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetParentJobType() *string {
	return s.ParentJobType
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetChargeType(v string) *DescribeCheckJobsResponseBodyCheckJobs {
	s.ChargeType = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetCheckPoint(v int64) *DescribeCheckJobsResponseBodyCheckJobs {
	s.CheckPoint = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetCheckType(v int32) *DescribeCheckJobsResponseBodyCheckJobs {
	s.CheckType = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetDiffCount(v int64) *DescribeCheckJobsResponseBodyCheckJobs {
	s.DiffCount = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetDiffSum(v int64) *DescribeCheckJobsResponseBodyCheckJobs {
	s.DiffSum = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetDtsInstanceID(v string) *DescribeCheckJobsResponseBodyCheckJobs {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetDtsJobId(v string) *DescribeCheckJobsResponseBodyCheckJobs {
	s.DtsJobId = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetFinishCount(v int64) *DescribeCheckJobsResponseBodyCheckJobs {
	s.FinishCount = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetGroupId(v string) *DescribeCheckJobsResponseBodyCheckJobs {
	s.GroupId = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetInstanceClass(v string) *DescribeCheckJobsResponseBodyCheckJobs {
	s.InstanceClass = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetJobName(v string) *DescribeCheckJobsResponseBodyCheckJobs {
	s.JobName = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetJobStepId(v string) *DescribeCheckJobsResponseBodyCheckJobs {
	s.JobStepId = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetParentJobType(v string) *DescribeCheckJobsResponseBodyCheckJobs {
	s.ParentJobType = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetRegionId(v string) *DescribeCheckJobsResponseBodyCheckJobs {
	s.RegionId = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetStatus(v int32) *DescribeCheckJobsResponseBodyCheckJobs {
	s.Status = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) SetTotalCount(v int64) *DescribeCheckJobsResponseBodyCheckJobs {
	s.TotalCount = &v
	return s
}

func (s *DescribeCheckJobsResponseBodyCheckJobs) Validate() error {
	return dara.Validate(s)
}
