// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheAnalysisJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCacheAnalysisJobResponseBody
	GetCode() *string
	SetData(v *DescribeCacheAnalysisJobResponseBodyData) *DescribeCacheAnalysisJobResponseBody
	GetData() *DescribeCacheAnalysisJobResponseBodyData
	SetMessage(v string) *DescribeCacheAnalysisJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCacheAnalysisJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeCacheAnalysisJobResponseBody
	GetSuccess() *string
}

type DescribeCacheAnalysisJobResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the cache analysis task.
	Data *DescribeCacheAnalysisJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned for the request.
	//
	// >  If the request is successful, **Successful*	- is returned. If the request fails, an error message that contains information such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCacheAnalysisJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCacheAnalysisJobResponseBody) GetData() *DescribeCacheAnalysisJobResponseBodyData {
	return s.Data
}

func (s *DescribeCacheAnalysisJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCacheAnalysisJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCacheAnalysisJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeCacheAnalysisJobResponseBody) SetCode(v string) *DescribeCacheAnalysisJobResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBody) SetData(v *DescribeCacheAnalysisJobResponseBodyData) *DescribeCacheAnalysisJobResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBody) SetMessage(v string) *DescribeCacheAnalysisJobResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBody) SetRequestId(v string) *DescribeCacheAnalysisJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBody) SetSuccess(v string) *DescribeCacheAnalysisJobResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyData struct {
	// The details of the large keys. The returned large keys are sorted in descending order based on the number of bytes occupied by the keys.
	BigKeys *DescribeCacheAnalysisJobResponseBodyDataBigKeys `json:"BigKeys,omitempty" xml:"BigKeys,omitempty" type:"Struct"`
	// The details of the large keys. The returned large keys are sorted in descending order based on the number of keys.
	BigKeysOfNum *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNum `json:"BigKeysOfNum,omitempty" xml:"BigKeysOfNum,omitempty" type:"Struct"`
	// The statistics of the keys that have expired.
	ExpiryKeysLevelCount *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCount `json:"ExpiryKeysLevelCount,omitempty" xml:"ExpiryKeysLevelCount,omitempty" type:"Struct"`
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
	// The prefixes of the keys.
	KeyPrefixes *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes `json:"KeyPrefixes,omitempty" xml:"KeyPrefixes,omitempty" type:"Struct"`
	// The message that is returned for the request.
	//
	// >  If the request is successful, **Successful*	- is returned. If the request fails, an error message that contains information such as an error code is returned.
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
	// The details of permanent keys. The returned keys are sorted in descending order based on the number of bytes occupied by the keys.
	UnexBigKeysOfBytes *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytes `json:"UnexBigKeysOfBytes,omitempty" xml:"UnexBigKeysOfBytes,omitempty" type:"Struct"`
	// The details of permanent keys. The returned keys are sorted in descending order based on the number of keys.
	UnexBigKeysOfNum *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNum `json:"UnexBigKeysOfNum,omitempty" xml:"UnexBigKeysOfNum,omitempty" type:"Struct"`
}

func (s DescribeCacheAnalysisJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyData) GetBigKeys() *DescribeCacheAnalysisJobResponseBodyDataBigKeys {
	return s.BigKeys
}

func (s *DescribeCacheAnalysisJobResponseBodyData) GetBigKeysOfNum() *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNum {
	return s.BigKeysOfNum
}

func (s *DescribeCacheAnalysisJobResponseBodyData) GetExpiryKeysLevelCount() *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCount {
	return s.ExpiryKeysLevelCount
}

func (s *DescribeCacheAnalysisJobResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCacheAnalysisJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *DescribeCacheAnalysisJobResponseBodyData) GetKeyPrefixes() *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes {
	return s.KeyPrefixes
}

func (s *DescribeCacheAnalysisJobResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DescribeCacheAnalysisJobResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeCacheAnalysisJobResponseBodyData) GetTaskState() *string {
	return s.TaskState
}

func (s *DescribeCacheAnalysisJobResponseBodyData) GetUnexBigKeysOfBytes() *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytes {
	return s.UnexBigKeysOfBytes
}

func (s *DescribeCacheAnalysisJobResponseBodyData) GetUnexBigKeysOfNum() *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNum {
	return s.UnexBigKeysOfNum
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetBigKeys(v *DescribeCacheAnalysisJobResponseBodyDataBigKeys) *DescribeCacheAnalysisJobResponseBodyData {
	s.BigKeys = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetBigKeysOfNum(v *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNum) *DescribeCacheAnalysisJobResponseBodyData {
	s.BigKeysOfNum = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetExpiryKeysLevelCount(v *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCount) *DescribeCacheAnalysisJobResponseBodyData {
	s.ExpiryKeysLevelCount = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetInstanceId(v string) *DescribeCacheAnalysisJobResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetJobId(v string) *DescribeCacheAnalysisJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetKeyPrefixes(v *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes) *DescribeCacheAnalysisJobResponseBodyData {
	s.KeyPrefixes = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetMessage(v string) *DescribeCacheAnalysisJobResponseBodyData {
	s.Message = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetNodeId(v string) *DescribeCacheAnalysisJobResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetTaskState(v string) *DescribeCacheAnalysisJobResponseBodyData {
	s.TaskState = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetUnexBigKeysOfBytes(v *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytes) *DescribeCacheAnalysisJobResponseBodyData {
	s.UnexBigKeysOfBytes = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) SetUnexBigKeysOfNum(v *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNum) *DescribeCacheAnalysisJobResponseBodyData {
	s.UnexBigKeysOfNum = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyDataBigKeys struct {
	KeyInfo []*DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo `json:"KeyInfo,omitempty" xml:"KeyInfo,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataBigKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataBigKeys) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeys) GetKeyInfo() []*DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	return s.KeyInfo
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeys) SetKeyInfo(v []*DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) *DescribeCacheAnalysisJobResponseBodyDataBigKeys {
	s.KeyInfo = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeys) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo struct {
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
	// The database name.
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
	// The time when the key expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. A value of 0 indicates that the key never expires.
	//
	// example:
	//
	// 1596256542547
	ExpirationTimeMillis *int64 `json:"ExpirationTimeMillis,omitempty" xml:"ExpirationTimeMillis,omitempty"`
	// The key name.
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
	// The data type of the instance.
	//
	// example:
	//
	// hash
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetBytes() *int64 {
	return s.Bytes
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetCount() *int64 {
	return s.Count
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetDb() *int32 {
	return s.Db
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetEncoding() *string {
	return s.Encoding
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetExpirationTimeMillis() *int64 {
	return s.ExpirationTimeMillis
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetKey() *string {
	return s.Key
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) GetType() *string {
	return s.Type
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetBytes(v int64) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Bytes = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetCount(v int64) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Count = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetDb(v int32) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Db = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetEncoding(v string) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Encoding = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetExpirationTimeMillis(v int64) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.ExpirationTimeMillis = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetKey(v string) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Key = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetNodeId(v string) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) SetType(v string) *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo {
	s.Type = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysKeyInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNum struct {
	KeyInfo []*DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo `json:"KeyInfo,omitempty" xml:"KeyInfo,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNum) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNum) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNum) GetKeyInfo() []*DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo {
	return s.KeyInfo
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNum) SetKeyInfo(v []*DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNum {
	s.KeyInfo = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNum) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo struct {
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
	// The database name.
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
	// The time when the key expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. A value of 0 indicates that the key never expires.
	//
	// example:
	//
	// 1596256542547
	ExpirationTimeMillis *int64 `json:"ExpirationTimeMillis,omitempty" xml:"ExpirationTimeMillis,omitempty"`
	// The key name.
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
	// The data type of the instance.
	//
	// example:
	//
	// hash
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) GetBytes() *int64 {
	return s.Bytes
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) GetCount() *int64 {
	return s.Count
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) GetDb() *int32 {
	return s.Db
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) GetEncoding() *string {
	return s.Encoding
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) GetExpirationTimeMillis() *int64 {
	return s.ExpirationTimeMillis
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) GetKey() *string {
	return s.Key
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) GetType() *string {
	return s.Type
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) SetBytes(v int64) *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo {
	s.Bytes = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) SetCount(v int64) *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo {
	s.Count = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) SetDb(v int32) *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo {
	s.Db = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) SetEncoding(v string) *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo {
	s.Encoding = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) SetExpirationTimeMillis(v int64) *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo {
	s.ExpirationTimeMillis = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) SetKey(v string) *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo {
	s.Key = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) SetNodeId(v string) *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) SetType(v string) *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo {
	s.Type = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataBigKeysOfNumKeyInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCount struct {
	ExpiryLevel []*DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel `json:"ExpiryLevel,omitempty" xml:"ExpiryLevel,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCount) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCount) GetExpiryLevel() []*DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel {
	return s.ExpiryLevel
}

func (s *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCount) SetExpiryLevel(v []*DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel) *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCount {
	s.ExpiryLevel = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCount) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel struct {
	// The time when the cache analysis task was complete. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1717469347000
	AnalysisTs *int64 `json:"AnalysisTs,omitempty" xml:"AnalysisTs,omitempty"`
	// The expiration level. Valid values:
	//
	// 	- **0**: The key never expires.
	//
	// 	- **1**: The key has expired.
	//
	// 	- **2**: The key has expired for 0 to 1 hour.
	//
	// 	- **3**: The key has expired for 1 to 3 hours.
	//
	// 	- **4**: The key has expired for 3 to 12 hours.
	//
	// 	- **5**: The key has expired for 12 to 24 hours.
	//
	// 	- **6**: The key has expired for one to two days.
	//
	// 	- **7**: The key has expired for three to seven days.
	//
	// 	- **8**: The key has expired for more than seven days.
	//
	// example:
	//
	// 0
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The number of bytes occupied by the keys that have expired.
	//
	// example:
	//
	// 8064
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// The total number of the keys that have expired.
	//
	// example:
	//
	// 62
	TotalKeys *int64 `json:"TotalKeys,omitempty" xml:"TotalKeys,omitempty"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel) GetAnalysisTs() *int64 {
	return s.AnalysisTs
}

func (s *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel) GetLevel() *int32 {
	return s.Level
}

func (s *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel) GetTotalKeys() *int64 {
	return s.TotalKeys
}

func (s *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel) SetAnalysisTs(v int64) *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel {
	s.AnalysisTs = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel) SetLevel(v int32) *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel {
	s.Level = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel) SetTotalBytes(v int64) *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel {
	s.TotalBytes = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel) SetTotalKeys(v int64) *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel {
	s.TotalKeys = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataExpiryKeysLevelCountExpiryLevel) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes struct {
	Prefix []*DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix `json:"Prefix,omitempty" xml:"Prefix,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes) GetPrefix() []*DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix {
	return s.Prefix
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes) SetPrefix(v []*DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes {
	s.Prefix = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixes) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix struct {
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
	// The number of keys that contain the prefix.
	//
	// example:
	//
	// 123
	KeyNum *int64 `json:"KeyNum,omitempty" xml:"KeyNum,omitempty"`
	// The prefix of the key.
	//
	// example:
	//
	// task_
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The data type of the instance.
	//
	// example:
	//
	// hash
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) GetBytes() *int64 {
	return s.Bytes
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) GetCount() *int64 {
	return s.Count
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) GetKeyNum() *int64 {
	return s.KeyNum
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) GetPrefix() *string {
	return s.Prefix
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) GetType() *string {
	return s.Type
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) SetBytes(v int64) *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix {
	s.Bytes = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) SetCount(v int64) *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix {
	s.Count = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) SetKeyNum(v int64) *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix {
	s.KeyNum = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) SetPrefix(v string) *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix {
	s.Prefix = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) SetType(v string) *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix {
	s.Type = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataKeyPrefixesPrefix) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytes struct {
	KeyInfo []*DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo `json:"KeyInfo,omitempty" xml:"KeyInfo,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytes) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytes) GetKeyInfo() []*DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo {
	return s.KeyInfo
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytes) SetKeyInfo(v []*DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytes {
	s.KeyInfo = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytes) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo struct {
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
	// The database name.
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
	// The time when the key expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. A value of 0 indicates that the key never expires.
	//
	// example:
	//
	// 1596256542547
	ExpirationTimeMillis *int64 `json:"ExpirationTimeMillis,omitempty" xml:"ExpirationTimeMillis,omitempty"`
	// The key name.
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
	// The data type of the instance.
	//
	// example:
	//
	// hash
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) GetBytes() *int64 {
	return s.Bytes
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) GetCount() *int64 {
	return s.Count
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) GetDb() *int32 {
	return s.Db
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) GetEncoding() *string {
	return s.Encoding
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) GetExpirationTimeMillis() *int64 {
	return s.ExpirationTimeMillis
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) GetKey() *string {
	return s.Key
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) GetType() *string {
	return s.Type
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) SetBytes(v int64) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo {
	s.Bytes = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) SetCount(v int64) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo {
	s.Count = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) SetDb(v int32) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo {
	s.Db = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) SetEncoding(v string) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo {
	s.Encoding = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) SetExpirationTimeMillis(v int64) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo {
	s.ExpirationTimeMillis = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) SetKey(v string) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo {
	s.Key = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) SetNodeId(v string) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) SetType(v string) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo {
	s.Type = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfBytesKeyInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNum struct {
	KeyInfo []*DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo `json:"KeyInfo,omitempty" xml:"KeyInfo,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNum) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNum) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNum) GetKeyInfo() []*DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo {
	return s.KeyInfo
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNum) SetKeyInfo(v []*DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNum {
	s.KeyInfo = v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNum) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo struct {
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
	// The database name.
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
	// The time when the key expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. A value of 0 indicates that the key never expires.
	//
	// example:
	//
	// 1596256542547
	ExpirationTimeMillis *int64 `json:"ExpirationTimeMillis,omitempty" xml:"ExpirationTimeMillis,omitempty"`
	// The key name.
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
	// The data type of the instance.
	//
	// example:
	//
	// hash
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) GetBytes() *int64 {
	return s.Bytes
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) GetCount() *int64 {
	return s.Count
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) GetDb() *int32 {
	return s.Db
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) GetEncoding() *string {
	return s.Encoding
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) GetExpirationTimeMillis() *int64 {
	return s.ExpirationTimeMillis
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) GetKey() *string {
	return s.Key
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) GetType() *string {
	return s.Type
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) SetBytes(v int64) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo {
	s.Bytes = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) SetCount(v int64) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo {
	s.Count = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) SetDb(v int32) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo {
	s.Db = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) SetEncoding(v string) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo {
	s.Encoding = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) SetExpirationTimeMillis(v int64) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo {
	s.ExpirationTimeMillis = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) SetKey(v string) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo {
	s.Key = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) SetNodeId(v string) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) SetType(v string) *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo {
	s.Type = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponseBodyDataUnexBigKeysOfNumKeyInfo) Validate() error {
	return dara.Validate(s)
}
