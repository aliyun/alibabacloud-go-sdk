// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHotBigKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHotBigKeysResponseBody
	GetCode() *string
	SetData(v *DescribeHotBigKeysResponseBodyData) *DescribeHotBigKeysResponseBody
	GetData() *DescribeHotBigKeysResponseBodyData
	SetMessage(v string) *DescribeHotBigKeysResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeHotBigKeysResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeHotBigKeysResponseBody
	GetSuccess() *string
}

type DescribeHotBigKeysResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of hot keys and large keys.
	Data *DescribeHotBigKeysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
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

func (s DescribeHotBigKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotBigKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHotBigKeysResponseBody) GetData() *DescribeHotBigKeysResponseBodyData {
	return s.Data
}

func (s *DescribeHotBigKeysResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHotBigKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHotBigKeysResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeHotBigKeysResponseBody) SetCode(v string) *DescribeHotBigKeysResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHotBigKeysResponseBody) SetData(v *DescribeHotBigKeysResponseBodyData) *DescribeHotBigKeysResponseBody {
	s.Data = v
	return s
}

func (s *DescribeHotBigKeysResponseBody) SetMessage(v string) *DescribeHotBigKeysResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHotBigKeysResponseBody) SetRequestId(v string) *DescribeHotBigKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHotBigKeysResponseBody) SetSuccess(v string) *DescribeHotBigKeysResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHotBigKeysResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHotBigKeysResponseBodyData struct {
	// The reason why the large key failed to be queried.
	//
	// example:
	//
	// current version doesn\\"t support
	BigKeyMsg *string `json:"BigKeyMsg,omitempty" xml:"BigKeyMsg,omitempty"`
	// The list of large keys.
	BigKeys           *DescribeHotBigKeysResponseBodyDataBigKeys         `json:"BigKeys,omitempty" xml:"BigKeys,omitempty" type:"Struct"`
	HighTrafficKeyMsg *string                                            `json:"HighTrafficKeyMsg,omitempty" xml:"HighTrafficKeyMsg,omitempty"`
	HighTrafficKeys   *DescribeHotBigKeysResponseBodyDataHighTrafficKeys `json:"HighTrafficKeys,omitempty" xml:"HighTrafficKeys,omitempty" type:"Struct"`
	// The reason why the hot key failed to be queried.
	//
	// example:
	//
	// current version doesn\\"t support
	HotKeyMsg *string `json:"HotKeyMsg,omitempty" xml:"HotKeyMsg,omitempty"`
	// The list of hot keys.
	HotKeys *DescribeHotBigKeysResponseBodyDataHotKeys `json:"HotKeys,omitempty" xml:"HotKeys,omitempty" type:"Struct"`
}

func (s DescribeHotBigKeysResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotBigKeysResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBodyData) GetBigKeyMsg() *string {
	return s.BigKeyMsg
}

func (s *DescribeHotBigKeysResponseBodyData) GetBigKeys() *DescribeHotBigKeysResponseBodyDataBigKeys {
	return s.BigKeys
}

func (s *DescribeHotBigKeysResponseBodyData) GetHighTrafficKeyMsg() *string {
	return s.HighTrafficKeyMsg
}

func (s *DescribeHotBigKeysResponseBodyData) GetHighTrafficKeys() *DescribeHotBigKeysResponseBodyDataHighTrafficKeys {
	return s.HighTrafficKeys
}

func (s *DescribeHotBigKeysResponseBodyData) GetHotKeyMsg() *string {
	return s.HotKeyMsg
}

func (s *DescribeHotBigKeysResponseBodyData) GetHotKeys() *DescribeHotBigKeysResponseBodyDataHotKeys {
	return s.HotKeys
}

func (s *DescribeHotBigKeysResponseBodyData) SetBigKeyMsg(v string) *DescribeHotBigKeysResponseBodyData {
	s.BigKeyMsg = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyData) SetBigKeys(v *DescribeHotBigKeysResponseBodyDataBigKeys) *DescribeHotBigKeysResponseBodyData {
	s.BigKeys = v
	return s
}

func (s *DescribeHotBigKeysResponseBodyData) SetHighTrafficKeyMsg(v string) *DescribeHotBigKeysResponseBodyData {
	s.HighTrafficKeyMsg = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyData) SetHighTrafficKeys(v *DescribeHotBigKeysResponseBodyDataHighTrafficKeys) *DescribeHotBigKeysResponseBodyData {
	s.HighTrafficKeys = v
	return s
}

func (s *DescribeHotBigKeysResponseBodyData) SetHotKeyMsg(v string) *DescribeHotBigKeysResponseBodyData {
	s.HotKeyMsg = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyData) SetHotKeys(v *DescribeHotBigKeysResponseBodyDataHotKeys) *DescribeHotBigKeysResponseBodyData {
	s.HotKeys = v
	return s
}

func (s *DescribeHotBigKeysResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeHotBigKeysResponseBodyDataBigKeys struct {
	BigKey []*DescribeHotBigKeysResponseBodyDataBigKeysBigKey `json:"BigKey,omitempty" xml:"BigKey,omitempty" type:"Repeated"`
}

func (s DescribeHotBigKeysResponseBodyDataBigKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotBigKeysResponseBodyDataBigKeys) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeys) GetBigKey() []*DescribeHotBigKeysResponseBodyDataBigKeysBigKey {
	return s.BigKey
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeys) SetBigKey(v []*DescribeHotBigKeysResponseBodyDataBigKeysBigKey) *DescribeHotBigKeysResponseBodyDataBigKeys {
	s.BigKey = v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeys) Validate() error {
	return dara.Validate(s)
}

type DescribeHotBigKeysResponseBodyDataBigKeysBigKey struct {
	// The database in which the key is stored.
	//
	// example:
	//
	// 0
	Db *int32 `json:"Db,omitempty" xml:"Db,omitempty"`
	// The key.
	//
	// example:
	//
	// abc:def:eng
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The type of the key.
	//
	// example:
	//
	// zset
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// The ID of the data shard on the ApsaraDB for Redis instance.
	//
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The number of elements in the key.
	//
	// example:
	//
	// 2
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeHotBigKeysResponseBodyDataBigKeysBigKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotBigKeysResponseBodyDataBigKeysBigKey) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) GetDb() *int32 {
	return s.Db
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) GetKey() *string {
	return s.Key
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) GetKeyType() *string {
	return s.KeyType
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) GetSize() *int64 {
	return s.Size
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) SetDb(v int32) *DescribeHotBigKeysResponseBodyDataBigKeysBigKey {
	s.Db = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) SetKey(v string) *DescribeHotBigKeysResponseBodyDataBigKeysBigKey {
	s.Key = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) SetKeyType(v string) *DescribeHotBigKeysResponseBodyDataBigKeysBigKey {
	s.KeyType = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) SetNodeId(v string) *DescribeHotBigKeysResponseBodyDataBigKeysBigKey {
	s.NodeId = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) SetSize(v int64) *DescribeHotBigKeysResponseBodyDataBigKeysBigKey {
	s.Size = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataBigKeysBigKey) Validate() error {
	return dara.Validate(s)
}

type DescribeHotBigKeysResponseBodyDataHighTrafficKeys struct {
	HighTrafficKey []*DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey `json:"HighTrafficKey,omitempty" xml:"HighTrafficKey,omitempty" type:"Repeated"`
}

func (s DescribeHotBigKeysResponseBodyDataHighTrafficKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotBigKeysResponseBodyDataHighTrafficKeys) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeys) GetHighTrafficKey() []*DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey {
	return s.HighTrafficKey
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeys) SetHighTrafficKey(v []*DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) *DescribeHotBigKeysResponseBodyDataHighTrafficKeys {
	s.HighTrafficKey = v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeys) Validate() error {
	return dara.Validate(s)
}

type DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey struct {
	Db       *int32  `json:"Db,omitempty" xml:"Db,omitempty"`
	Hot      *string `json:"Hot,omitempty" xml:"Hot,omitempty"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	KeyType  *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	NodeId   *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Size     *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	InBytes  *int64  `json:"inBytes,omitempty" xml:"inBytes,omitempty"`
	OutBytes *int64  `json:"outBytes,omitempty" xml:"outBytes,omitempty"`
}

func (s DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) GetDb() *int32 {
	return s.Db
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) GetHot() *string {
	return s.Hot
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) GetKey() *string {
	return s.Key
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) GetKeyType() *string {
	return s.KeyType
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) GetSize() *int64 {
	return s.Size
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) SetDb(v int32) *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey {
	s.Db = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) SetHot(v string) *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey {
	s.Hot = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) SetKey(v string) *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey {
	s.Key = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) SetKeyType(v string) *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey {
	s.KeyType = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) SetNodeId(v string) *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey {
	s.NodeId = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) SetSize(v int64) *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey {
	s.Size = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) SetInBytes(v int64) *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey {
	s.InBytes = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) SetOutBytes(v int64) *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey {
	s.OutBytes = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHighTrafficKeysHighTrafficKey) Validate() error {
	return dara.Validate(s)
}

type DescribeHotBigKeysResponseBodyDataHotKeys struct {
	HotKey []*DescribeHotBigKeysResponseBodyDataHotKeysHotKey `json:"HotKey,omitempty" xml:"HotKey,omitempty" type:"Repeated"`
}

func (s DescribeHotBigKeysResponseBodyDataHotKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotBigKeysResponseBodyDataHotKeys) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeys) GetHotKey() []*DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	return s.HotKey
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeys) SetHotKey(v []*DescribeHotBigKeysResponseBodyDataHotKeysHotKey) *DescribeHotBigKeysResponseBodyDataHotKeys {
	s.HotKey = v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeys) Validate() error {
	return dara.Validate(s)
}

type DescribeHotBigKeysResponseBodyDataHotKeysHotKey struct {
	// The database in which the key is stored.
	//
	// example:
	//
	// 0
	Db *int32 `json:"Db,omitempty" xml:"Db,omitempty"`
	// The frequency at which the key is accessed, which indicates the queries per second (QPS) of the key.
	//
	// example:
	//
	// 5500~6000
	Hot *string `json:"Hot,omitempty" xml:"Hot,omitempty"`
	// The key.
	//
	// example:
	//
	// abc:def:eng
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The type of the key.
	//
	// example:
	//
	// zset
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// The statistical value that is calculated based on the least frequently used (LFU) caching algorithm.
	//
	// example:
	//
	// 253
	Lfu *int32 `json:"Lfu,omitempty" xml:"Lfu,omitempty"`
	// The ID of the data shard on the ApsaraDB for Redis instance.
	//
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Size   *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeHotBigKeysResponseBodyDataHotKeysHotKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotBigKeysResponseBodyDataHotKeysHotKey) GoString() string {
	return s.String()
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) GetDb() *int32 {
	return s.Db
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) GetHot() *string {
	return s.Hot
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) GetKey() *string {
	return s.Key
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) GetKeyType() *string {
	return s.KeyType
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) GetLfu() *int32 {
	return s.Lfu
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) GetSize() *int64 {
	return s.Size
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetDb(v int32) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.Db = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetHot(v string) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.Hot = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetKey(v string) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.Key = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetKeyType(v string) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.KeyType = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetLfu(v int32) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.Lfu = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetNodeId(v string) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.NodeId = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) SetSize(v int64) *DescribeHotBigKeysResponseBodyDataHotKeysHotKey {
	s.Size = &v
	return s
}

func (s *DescribeHotBigKeysResponseBodyDataHotKeysHotKey) Validate() error {
	return dara.Validate(s)
}
