// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCacheAnalysisJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCacheAnalysisJobResponseBody
	GetCode() *string
	SetData(v *CreateCacheAnalysisJobResponseBodyData) *CreateCacheAnalysisJobResponseBody
	GetData() *CreateCacheAnalysisJobResponseBodyData
	SetMessage(v string) *CreateCacheAnalysisJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCacheAnalysisJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateCacheAnalysisJobResponseBody
	GetSuccess() *string
}

type CreateCacheAnalysisJobResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information.
	Data *CreateCacheAnalysisJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateCacheAnalysisJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheAnalysisJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCacheAnalysisJobResponseBody) GetData() *CreateCacheAnalysisJobResponseBodyData {
	return s.Data
}

func (s *CreateCacheAnalysisJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCacheAnalysisJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCacheAnalysisJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateCacheAnalysisJobResponseBody) SetCode(v string) *CreateCacheAnalysisJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBody) SetData(v *CreateCacheAnalysisJobResponseBodyData) *CreateCacheAnalysisJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateCacheAnalysisJobResponseBody) SetMessage(v string) *CreateCacheAnalysisJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBody) SetRequestId(v string) *CreateCacheAnalysisJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBody) SetSuccess(v string) *CreateCacheAnalysisJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCacheAnalysisJobResponseBodyData struct {
	// The number of elements in the key.
	BigKeys *CreateCacheAnalysisJobResponseBodyDataBigKeys `json:"BigKeys,omitempty" xml:"BigKeys,omitempty" type:"Struct"`
	// The instance ID.
	//
	// example:
	//
	// r-bp18ff4a195d****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the cache analysis task.
	//
	// >  This parameter can be used to query a specific cache analysis task. When you call the CreateCacheAnalysisJob operation, it takes some time to create a cache analysis task. As a result, the analysis results cannot be immediately returned. You can call the [DescribeCacheAnalysisJob](https://help.aliyun.com/document_detail/180983.html) operation to query the analysis results of the specified cache analysis task.
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

func (s CreateCacheAnalysisJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheAnalysisJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisJobResponseBodyData) GetBigKeys() *CreateCacheAnalysisJobResponseBodyDataBigKeys {
	return s.BigKeys
}

func (s *CreateCacheAnalysisJobResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCacheAnalysisJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *CreateCacheAnalysisJobResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CreateCacheAnalysisJobResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateCacheAnalysisJobResponseBodyData) GetTaskState() *string {
	return s.TaskState
}

func (s *CreateCacheAnalysisJobResponseBodyData) SetBigKeys(v *CreateCacheAnalysisJobResponseBodyDataBigKeys) *CreateCacheAnalysisJobResponseBodyData {
	s.BigKeys = v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyData) SetInstanceId(v string) *CreateCacheAnalysisJobResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyData) SetJobId(v string) *CreateCacheAnalysisJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyData) SetMessage(v string) *CreateCacheAnalysisJobResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyData) SetNodeId(v string) *CreateCacheAnalysisJobResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyData) SetTaskState(v string) *CreateCacheAnalysisJobResponseBodyData {
	s.TaskState = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyData) Validate() error {
	if s.BigKeys != nil {
		if err := s.BigKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCacheAnalysisJobResponseBodyDataBigKeys struct {
	KeyInfo []*CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo `json:"KeyInfo,omitempty" xml:"KeyInfo,omitempty" type:"Repeated"`
}

func (s CreateCacheAnalysisJobResponseBodyDataBigKeys) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheAnalysisJobResponseBodyDataBigKeys) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeys) GetKeyInfo() []*CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	return s.KeyInfo
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeys) SetKeyInfo(v []*CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) *CreateCacheAnalysisJobResponseBodyDataBigKeys {
	s.KeyInfo = v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeys) Validate() error {
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

type CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo struct {
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
	// task_x****
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the data node on the instance.
	//
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The data type of the ApsaraDB for Redis instance.
	//
	// example:
	//
	// hash
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GoString() string {
	return s.String()
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetBytes() *int64 {
	return s.Bytes
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetCount() *int64 {
	return s.Count
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetDb() *int32 {
	return s.Db
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetEncoding() *string {
	return s.Encoding
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetExpirationTimeMillis() *int64 {
	return s.ExpirationTimeMillis
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetKey() *string {
	return s.Key
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetType() *string {
	return s.Type
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetBytes(v int64) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Bytes = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetCount(v int64) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Count = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetDb(v int32) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Db = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetEncoding(v string) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Encoding = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetExpirationTimeMillis(v int64) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.ExpirationTimeMillis = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetKey(v string) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Key = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetNodeId(v string) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.NodeId = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetType(v string) *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Type = &v
	return s
}

func (s *CreateCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) Validate() error {
	return dara.Validate(s)
}
