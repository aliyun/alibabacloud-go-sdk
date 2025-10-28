// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheAnalysisJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCacheAnalysisJobsResponseBody
	GetCode() *string
	SetData(v *DescribeCacheAnalysisJobsResponseBodyData) *DescribeCacheAnalysisJobsResponseBody
	GetData() *DescribeCacheAnalysisJobsResponseBodyData
	SetMessage(v string) *DescribeCacheAnalysisJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCacheAnalysisJobsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeCacheAnalysisJobsResponseBody
	GetSuccess() *string
}

type DescribeCacheAnalysisJobsResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of cache analysis tasks.
	Data *DescribeCacheAnalysisJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCacheAnalysisJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCacheAnalysisJobsResponseBody) GetData() *DescribeCacheAnalysisJobsResponseBodyData {
	return s.Data
}

func (s *DescribeCacheAnalysisJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCacheAnalysisJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCacheAnalysisJobsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeCacheAnalysisJobsResponseBody) SetCode(v string) *DescribeCacheAnalysisJobsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBody) SetData(v *DescribeCacheAnalysisJobsResponseBodyData) *DescribeCacheAnalysisJobsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBody) SetMessage(v string) *DescribeCacheAnalysisJobsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBody) SetRequestId(v string) *DescribeCacheAnalysisJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBody) SetSuccess(v string) *DescribeCacheAnalysisJobsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCacheAnalysisJobsResponseBodyData struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The ID of the data node on the instance.
	List *DescribeCacheAnalysisJobsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// The page number. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCacheAnalysisJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) GetList() *DescribeCacheAnalysisJobsResponseBodyDataList {
	return s.List
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) SetExtra(v string) *DescribeCacheAnalysisJobsResponseBodyData {
	s.Extra = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) SetList(v *DescribeCacheAnalysisJobsResponseBodyDataList) *DescribeCacheAnalysisJobsResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) SetPageNo(v int64) *DescribeCacheAnalysisJobsResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) SetPageSize(v int64) *DescribeCacheAnalysisJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) SetTotal(v int64) *DescribeCacheAnalysisJobsResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyData) Validate() error {
	if s.List != nil {
		if err := s.List.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCacheAnalysisJobsResponseBodyDataList struct {
	CacheAnalysisJob []*DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob `json:"CacheAnalysisJob,omitempty" xml:"CacheAnalysisJob,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisJobsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataList) GetCacheAnalysisJob() []*DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	return s.CacheAnalysisJob
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataList) SetCacheAnalysisJob(v []*DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) *DescribeCacheAnalysisJobsResponseBodyDataList {
	s.CacheAnalysisJob = v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataList) Validate() error {
	if s.CacheAnalysisJob != nil {
		for _, item := range s.CacheAnalysisJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob struct {
	// The details about the large keys.
	//
	// > The sub-parameters of this parameter and the content of the sub-parameters are not returned. To query the detailed information about the cache analysis tasks, call the [DescribeCacheAnalysisJob](https://help.aliyun.com/document_detail/443012.html) operation.
	BigKeys *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys `json:"BigKeys,omitempty" xml:"BigKeys,omitempty" type:"Struct"`
	// The instance ID.
	//
	// example:
	//
	// r-bp18ff4a195d****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the cache analysis task.
	//
	// example:
	//
	// sf79-sd99-sa37-****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the data node on the instance.
	//
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The state of the cache analysis task. Valid values:
	//
	// 	- **BACKUP**: The data is being backed up.
	//
	// 	- **ANALYZING**: The data is being analyzed.
	//
	// 	- **FINISHED**: The data is analyzed.
	//
	// 	- **FAILED**: An error occurred.
	//
	// example:
	//
	// BACKUP
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
}

func (s DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) GetBigKeys() *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys {
	return s.BigKeys
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) GetJobId() *string {
	return s.JobId
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) GetMessage() *string {
	return s.Message
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) GetTaskState() *string {
	return s.TaskState
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) SetBigKeys(v *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	s.BigKeys = v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) SetInstanceId(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) SetJobId(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	s.JobId = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) SetMessage(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	s.Message = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) SetNodeId(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) SetTaskState(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob {
	s.TaskState = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJob) Validate() error {
	if s.BigKeys != nil {
		if err := s.BigKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys struct {
	KeyInfo []*DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo `json:"KeyInfo,omitempty" xml:"KeyInfo,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys) GetKeyInfo() []*DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	return s.KeyInfo
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys) SetKeyInfo(v []*DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys {
	s.KeyInfo = v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeys) Validate() error {
	if s.KeyInfo != nil {
		for _, item := range s.KeyInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo struct {
	// The number of bytes that are occupied by the key.
	//
	// example:
	//
	// 12345
	Bytes *int64 `json:"Bytes,omitempty" xml:"Bytes,omitempty"`
	// The number of elements in the key.
	//
	// example:
	//
	// 127
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// 0
	Db *int32 `json:"Db,omitempty" xml:"Db,omitempty"`
	// The data type of the key.
	//
	// example:
	//
	// hashtable
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// The expiration period of the key. Unit: milliseconds. A value of 0 indicates that the key does not expire.
	//
	// example:
	//
	// 1596256542547
	ExpirationTimeMillis *int64 `json:"ExpirationTimeMillis,omitempty" xml:"ExpirationTimeMillis,omitempty"`
	// The name of the key.
	//
	// example:
	//
	// task_*****
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the data node on the instance.
	//
	// example:
	//
	// r-****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The data type of the instance.
	//
	// example:
	//
	// hash
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) GetBytes() *int64 {
	return s.Bytes
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) GetCount() *int64 {
	return s.Count
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) GetDb() *int32 {
	return s.Db
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) GetEncoding() *string {
	return s.Encoding
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) GetExpirationTimeMillis() *int64 {
	return s.ExpirationTimeMillis
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) GetKey() *string {
	return s.Key
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) GetType() *string {
	return s.Type
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetBytes(v int64) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.Bytes = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetCount(v int64) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.Count = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetDb(v int32) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.Db = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetEncoding(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.Encoding = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetExpirationTimeMillis(v int64) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.ExpirationTimeMillis = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetKey(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.Key = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetNodeId(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) SetType(v string) *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo {
	s.Type = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponseBodyDataListCacheAnalysisJobBigKeysKeyInfo) Validate() error {
	return dara.Validate(s)
}
