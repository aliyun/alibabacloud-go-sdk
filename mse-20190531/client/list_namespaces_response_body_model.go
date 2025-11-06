// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListNamespacesResponseBodyData) *ListNamespacesResponseBody
	GetData() *ListNamespacesResponseBodyData
	SetRequestId(v string) *ListNamespacesResponseBody
	GetRequestId() *string
}

type ListNamespacesResponseBody struct {
	Data *ListNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D1F1A6F3-7E03-5EAD-B3F1-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBody) GetData() *ListNamespacesResponseBodyData {
	return s.Data
}

func (s *ListNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNamespacesResponseBody) SetData(v *ListNamespacesResponseBodyData) *ListNamespacesResponseBody {
	s.Data = v
	return s
}

func (s *ListNamespacesResponseBody) SetRequestId(v string) *ListNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespacesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNamespacesResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*ListNamespacesResponseBodyDataResult  `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Results  []*ListNamespacesResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// 7
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListNamespacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNamespacesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNamespacesResponseBodyData) GetResult() []*ListNamespacesResponseBodyDataResult {
	return s.Result
}

func (s *ListNamespacesResponseBodyData) GetResults() []*ListNamespacesResponseBodyDataResults {
	return s.Results
}

func (s *ListNamespacesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListNamespacesResponseBodyData) SetPageNumber(v int32) *ListNamespacesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListNamespacesResponseBodyData) SetPageSize(v int32) *ListNamespacesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListNamespacesResponseBodyData) SetResult(v []*ListNamespacesResponseBodyDataResult) *ListNamespacesResponseBodyData {
	s.Result = v
	return s
}

func (s *ListNamespacesResponseBodyData) SetResults(v []*ListNamespacesResponseBodyDataResults) *ListNamespacesResponseBodyData {
	s.Results = v
	return s
}

func (s *ListNamespacesResponseBodyData) SetTotalSize(v int32) *ListNamespacesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListNamespacesResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNamespacesResponseBodyDataResult struct {
	// example:
	//
	// 3
	AppCount *int32 `json:"AppCount,omitempty" xml:"AppCount,omitempty"`
	// example:
	//
	// 2024-09-02T09:49:48.000+0000
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Describe   *string `json:"Describe,omitempty" xml:"Describe,omitempty"`
	// example:
	//
	// 6
	InstanceCount *int64 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// example:
	//
	// myNamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string                `json:"Region,omitempty" xml:"Region,omitempty"`
	Tags   map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 2024-09-02T09:49:48.000+0000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 178*******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListNamespacesResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyDataResult) GetAppCount() *int32 {
	return s.AppCount
}

func (s *ListNamespacesResponseBodyDataResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNamespacesResponseBodyDataResult) GetDescribe() *string {
	return s.Describe
}

func (s *ListNamespacesResponseBodyDataResult) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *ListNamespacesResponseBodyDataResult) GetNamespace() *string {
	return s.Namespace
}

func (s *ListNamespacesResponseBodyDataResult) GetRegion() *string {
	return s.Region
}

func (s *ListNamespacesResponseBodyDataResult) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListNamespacesResponseBodyDataResult) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListNamespacesResponseBodyDataResult) GetUserId() *string {
	return s.UserId
}

func (s *ListNamespacesResponseBodyDataResult) GetVersion() *int32 {
	return s.Version
}

func (s *ListNamespacesResponseBodyDataResult) SetAppCount(v int32) *ListNamespacesResponseBodyDataResult {
	s.AppCount = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResult) SetCreateTime(v int64) *ListNamespacesResponseBodyDataResult {
	s.CreateTime = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResult) SetDescribe(v string) *ListNamespacesResponseBodyDataResult {
	s.Describe = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResult) SetInstanceCount(v int64) *ListNamespacesResponseBodyDataResult {
	s.InstanceCount = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResult) SetNamespace(v string) *ListNamespacesResponseBodyDataResult {
	s.Namespace = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResult) SetRegion(v string) *ListNamespacesResponseBodyDataResult {
	s.Region = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResult) SetTags(v map[string]interface{}) *ListNamespacesResponseBodyDataResult {
	s.Tags = v
	return s
}

func (s *ListNamespacesResponseBodyDataResult) SetUpdateTime(v int64) *ListNamespacesResponseBodyDataResult {
	s.UpdateTime = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResult) SetUserId(v string) *ListNamespacesResponseBodyDataResult {
	s.UserId = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResult) SetVersion(v int32) *ListNamespacesResponseBodyDataResult {
	s.Version = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type ListNamespacesResponseBodyDataResults struct {
	// example:
	//
	// 3
	AppCount *int32 `json:"AppCount,omitempty" xml:"AppCount,omitempty"`
	// example:
	//
	// 2024-09-02T09:49:48.000+0000
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Describe   *string `json:"Describe,omitempty" xml:"Describe,omitempty"`
	// example:
	//
	// 6
	InstanceCount *int64 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// example:
	//
	// myNamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 2024-09-02T09:49:48.000+0000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 178*******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListNamespacesResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyDataResults) GetAppCount() *int32 {
	return s.AppCount
}

func (s *ListNamespacesResponseBodyDataResults) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNamespacesResponseBodyDataResults) GetDescribe() *string {
	return s.Describe
}

func (s *ListNamespacesResponseBodyDataResults) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *ListNamespacesResponseBodyDataResults) GetNamespace() *string {
	return s.Namespace
}

func (s *ListNamespacesResponseBodyDataResults) GetRegion() *string {
	return s.Region
}

func (s *ListNamespacesResponseBodyDataResults) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListNamespacesResponseBodyDataResults) GetUserId() *string {
	return s.UserId
}

func (s *ListNamespacesResponseBodyDataResults) GetVersion() *int32 {
	return s.Version
}

func (s *ListNamespacesResponseBodyDataResults) SetAppCount(v int32) *ListNamespacesResponseBodyDataResults {
	s.AppCount = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResults) SetCreateTime(v int64) *ListNamespacesResponseBodyDataResults {
	s.CreateTime = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResults) SetDescribe(v string) *ListNamespacesResponseBodyDataResults {
	s.Describe = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResults) SetInstanceCount(v int64) *ListNamespacesResponseBodyDataResults {
	s.InstanceCount = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResults) SetNamespace(v string) *ListNamespacesResponseBodyDataResults {
	s.Namespace = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResults) SetRegion(v string) *ListNamespacesResponseBodyDataResults {
	s.Region = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResults) SetUpdateTime(v int64) *ListNamespacesResponseBodyDataResults {
	s.UpdateTime = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResults) SetUserId(v string) *ListNamespacesResponseBodyDataResults {
	s.UserId = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResults) SetVersion(v int32) *ListNamespacesResponseBodyDataResults {
	s.Version = &v
	return s
}

func (s *ListNamespacesResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
