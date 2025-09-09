// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceDeployDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeployDetails(v []*ListServiceInstanceDeployDetailsResponseBodyDeployDetails) *ListServiceInstanceDeployDetailsResponseBody
	GetDeployDetails() []*ListServiceInstanceDeployDetailsResponseBodyDeployDetails
	SetMaxResults(v int32) *ListServiceInstanceDeployDetailsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceInstanceDeployDetailsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceInstanceDeployDetailsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListServiceInstanceDeployDetailsResponseBody
	GetTotalCount() *int32
}

type ListServiceInstanceDeployDetailsResponseBody struct {
	// The details of the service instance deployment.
	DeployDetails []*ListServiceInstanceDeployDetailsResponseBodyDeployDetails `json:"DeployDetails,omitempty" xml:"DeployDetails,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAW8kZY+u1sYOaYf5JmgmDQQ=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0036D82E-0624-5B37-B797-C460F4B02026
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstanceDeployDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceDeployDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceDeployDetailsResponseBody) GetDeployDetails() []*ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	return s.DeployDetails
}

func (s *ListServiceInstanceDeployDetailsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceInstanceDeployDetailsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceInstanceDeployDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceInstanceDeployDetailsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServiceInstanceDeployDetailsResponseBody) SetDeployDetails(v []*ListServiceInstanceDeployDetailsResponseBodyDeployDetails) *ListServiceInstanceDeployDetailsResponseBody {
	s.DeployDetails = v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBody) SetMaxResults(v int32) *ListServiceInstanceDeployDetailsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBody) SetNextToken(v string) *ListServiceInstanceDeployDetailsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBody) SetRequestId(v string) *ListServiceInstanceDeployDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBody) SetTotalCount(v int32) *ListServiceInstanceDeployDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstanceDeployDetailsResponseBodyDeployDetails struct {
	// The total number of entries that meet the specified conditions.
	//
	// example:
	//
	// 4
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the service instance was created.
	//
	// example:
	//
	// 2024-04-10T01:58:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The period over which data is aggregated.
	//
	// example:
	//
	// Month
	Cycle *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// The indicates whether the deployment was successful.
	//
	// example:
	//
	// False
	DeploySucceeded *string `json:"DeploySucceeded,omitempty" xml:"DeploySucceeded,omitempty"`
	// The error code.
	//
	// example:
	//
	// StackValidationFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error description.
	//
	// example:
	//
	// {code: StackValidationFailed, message: \\"Failed to continue create ROS stack 89e724e2-84e6-4517-a372-30a545ab4145: Resource [LinuxInstanceRunCommand]: i-wz91nfbh1fxtmfb0try4 are not running. Command invocation only support running instances. ErrorCode: StackValidationFailed\\", requestId: null}
	ErrorDetail *string `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty"`
	// The type of error that caused the deployment to fail.
	//
	// example:
	//
	// ValidationError
	ErrorType *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-c751ed91f2074af39779
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-273e8cee11d349e1803c
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The name of the service in Chinese.
	//
	// example:
	//
	// 测试服务(Test Service)
	ServiceNameChn *string `json:"ServiceNameChn,omitempty" xml:"ServiceNameChn,omitempty"`
	// The name of the service in English.
	//
	// example:
	//
	// Test Service
	ServiceNameEng *string `json:"ServiceNameEng,omitempty" xml:"ServiceNameEng,omitempty"`
	// The type of service.
	//
	// Possible values:
	//
	// - private: Deployed under the user\\"s account.
	//
	// - managed: Hosted under the service provider\\"s account.
	//
	// - operation: Managed operation service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The service version.
	//
	// example:
	//
	// 42
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The timestamp when the response is returned.
	//
	// example:
	//
	// 1723946641994
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The aliuid of user.
	//
	// example:
	//
	// 1591457835436382
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListServiceInstanceDeployDetailsResponseBodyDeployDetails) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetCount() *string {
	return s.Count
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetCycle() *string {
	return s.Cycle
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetDeploySucceeded() *string {
	return s.DeploySucceeded
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetErrorDetail() *string {
	return s.ErrorDetail
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetErrorType() *string {
	return s.ErrorType
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetServiceNameChn() *string {
	return s.ServiceNameChn
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetServiceNameEng() *string {
	return s.ServiceNameEng
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) GetUserId() *string {
	return s.UserId
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetCount(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.Count = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetCreateTime(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetCycle(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.Cycle = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetDeploySucceeded(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.DeploySucceeded = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetErrorCode(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ErrorCode = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetErrorDetail(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ErrorDetail = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetErrorType(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ErrorType = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetServiceId(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ServiceId = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetServiceInstanceId(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetServiceNameChn(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ServiceNameChn = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetServiceNameEng(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ServiceNameEng = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetServiceType(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetServiceVersion(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.ServiceVersion = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetTimestamp(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.Timestamp = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) SetUserId(v string) *ListServiceInstanceDeployDetailsResponseBodyDeployDetails {
	s.UserId = &v
	return s
}

func (s *ListServiceInstanceDeployDetailsResponseBodyDeployDetails) Validate() error {
	return dara.Validate(s)
}
