// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChainInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChainInstances(v []*ListChainInstanceResponseBodyChainInstances) *ListChainInstanceResponseBody
	GetChainInstances() []*ListChainInstanceResponseBodyChainInstances
	SetCode(v string) *ListChainInstanceResponseBody
	GetCode() *string
	SetInstanceId(v string) *ListChainInstanceResponseBody
	GetInstanceId() *string
	SetIsSuccess(v bool) *ListChainInstanceResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListChainInstanceResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListChainInstanceResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListChainInstanceResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListChainInstanceResponseBody
	GetTotalCount() *int32
}

type ListChainInstanceResponseBody struct {
	// The list of delivery chain execution records.
	ChainInstances []*ListChainInstanceResponseBodyChainInstances `json:"ChainInstances,omitempty" xml:"ChainInstances,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The page size.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 838D1602-6D8F-47FB-B60A-656645D2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChainInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChainInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListChainInstanceResponseBody) GetChainInstances() []*ListChainInstanceResponseBodyChainInstances {
	return s.ChainInstances
}

func (s *ListChainInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChainInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChainInstanceResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListChainInstanceResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListChainInstanceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChainInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChainInstanceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListChainInstanceResponseBody) SetChainInstances(v []*ListChainInstanceResponseBodyChainInstances) *ListChainInstanceResponseBody {
	s.ChainInstances = v
	return s
}

func (s *ListChainInstanceResponseBody) SetCode(v string) *ListChainInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListChainInstanceResponseBody) SetInstanceId(v string) *ListChainInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ListChainInstanceResponseBody) SetIsSuccess(v bool) *ListChainInstanceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListChainInstanceResponseBody) SetPageNo(v int32) *ListChainInstanceResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListChainInstanceResponseBody) SetPageSize(v int32) *ListChainInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChainInstanceResponseBody) SetRequestId(v string) *ListChainInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChainInstanceResponseBody) SetTotalCount(v int32) *ListChainInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChainInstanceResponseBody) Validate() error {
	if s.ChainInstances != nil {
		for _, item := range s.ChainInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChainInstanceResponseBodyChainInstances struct {
	// The delivery chain execution record.
	Chain *ListChainInstanceResponseBodyChainInstancesChain `json:"Chain,omitempty" xml:"Chain,omitempty" type:"Struct"`
	// The delivery chain instance ID.
	//
	// example:
	//
	// F4CF4DDB-BEF2-5575-****-*******
	ChainInstanceId *string `json:"ChainInstanceId,omitempty" xml:"ChainInstanceId,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1636685856000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The repository name.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace.
	//
	// example:
	//
	// test-ns
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The execution result of the delivery chain. Valid values:
	//
	// - `SUCCESS`: Succeeded.
	//
	// - `FAILED`: Failed.
	//
	// - `CANCELED`: Canceled.
	//
	// - `DENIED`: Denied.
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1636685776000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution status of the delivery chain. Valid values:
	//
	// - `RUNNING`: Running.
	//
	// - `COMPLETE`: Complete.
	//
	// - `CANCELING`: Canceling.
	//
	// - `CANCELED`: Canceled.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListChainInstanceResponseBodyChainInstances) String() string {
	return dara.Prettify(s)
}

func (s ListChainInstanceResponseBodyChainInstances) GoString() string {
	return s.String()
}

func (s *ListChainInstanceResponseBodyChainInstances) GetChain() *ListChainInstanceResponseBodyChainInstancesChain {
	return s.Chain
}

func (s *ListChainInstanceResponseBodyChainInstances) GetChainInstanceId() *string {
	return s.ChainInstanceId
}

func (s *ListChainInstanceResponseBodyChainInstances) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListChainInstanceResponseBodyChainInstances) GetRepoName() *string {
	return s.RepoName
}

func (s *ListChainInstanceResponseBodyChainInstances) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListChainInstanceResponseBodyChainInstances) GetResult() *string {
	return s.Result
}

func (s *ListChainInstanceResponseBodyChainInstances) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListChainInstanceResponseBodyChainInstances) GetStatus() *string {
	return s.Status
}

func (s *ListChainInstanceResponseBodyChainInstances) SetChain(v *ListChainInstanceResponseBodyChainInstancesChain) *ListChainInstanceResponseBodyChainInstances {
	s.Chain = v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetChainInstanceId(v string) *ListChainInstanceResponseBodyChainInstances {
	s.ChainInstanceId = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetEndTime(v int64) *ListChainInstanceResponseBodyChainInstances {
	s.EndTime = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetRepoName(v string) *ListChainInstanceResponseBodyChainInstances {
	s.RepoName = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetRepoNamespaceName(v string) *ListChainInstanceResponseBodyChainInstances {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetResult(v string) *ListChainInstanceResponseBodyChainInstances {
	s.Result = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetStartTime(v int64) *ListChainInstanceResponseBodyChainInstances {
	s.StartTime = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetStatus(v string) *ListChainInstanceResponseBodyChainInstances {
	s.Status = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) Validate() error {
	if s.Chain != nil {
		if err := s.Chain.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListChainInstanceResponseBodyChainInstancesChain struct {
	// The delivery chain ID.
	//
	// example:
	//
	// chi-m42gbku0****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The delivery chain name.
	//
	// example:
	//
	// test-chain
	ChainName *string `json:"ChainName,omitempty" xml:"ChainName,omitempty"`
	// The delivery chain version.
	//
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListChainInstanceResponseBodyChainInstancesChain) String() string {
	return dara.Prettify(s)
}

func (s ListChainInstanceResponseBodyChainInstancesChain) GoString() string {
	return s.String()
}

func (s *ListChainInstanceResponseBodyChainInstancesChain) GetChainId() *string {
	return s.ChainId
}

func (s *ListChainInstanceResponseBodyChainInstancesChain) GetChainName() *string {
	return s.ChainName
}

func (s *ListChainInstanceResponseBodyChainInstancesChain) GetVersion() *int64 {
	return s.Version
}

func (s *ListChainInstanceResponseBodyChainInstancesChain) SetChainId(v string) *ListChainInstanceResponseBodyChainInstancesChain {
	s.ChainId = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstancesChain) SetChainName(v string) *ListChainInstanceResponseBodyChainInstancesChain {
	s.ChainName = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstancesChain) SetVersion(v int64) *ListChainInstanceResponseBodyChainInstancesChain {
	s.Version = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstancesChain) Validate() error {
	return dara.Validate(s)
}
